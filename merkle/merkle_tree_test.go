package merkle

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var AllZeros = strings.Repeat("0", 256)

func TestBitString(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"0", AllZeros},
	}

	for _, test := range tests {
		if got, want := BitString(test.input), test.output; got != want {
			t.Errorf("BitString(%v)=%v, want %v", test.input, got, want)
		}
	}
}

func TestAddRoot(t *testing.T) {
	m := New()
	tests := []struct {
		epoch Epoch
		code  codes.Code
	}{
		{10, codes.OK},
		{10, codes.OK},
		{11, codes.OK},
		{10, codes.FailedPrecondition},
		{12, codes.OK},
	}
	for _, test := range tests {
		_, err := m.addRoot(test.epoch)
		if got, want := grpc.Code(err), test.code; got != want {
			t.Errorf("addRoot(%v)=%v, want %v", test.epoch, got, want)
		}
	}
}

func TestAddLeaf(t *testing.T) {
	m := New()
	tests := []struct {
		epoch Epoch
		index string
		code  codes.Code
	}{
		// First insert
		{0, "0000000000000000000000000000000000000000000000000000000000000000", codes.OK},
		// Inserting a duplicate in the same epoch should fail.
		{0, "0000000000000000000000000000000000000000000000000000000000000000", codes.AlreadyExists},
		// Insert a leaf node with a long shared prefix. Should increase tree depth to max.
		{0, "0000000000000000000000000000000000000000000000000000000000000001", codes.OK},
		// Insert a leaf node with a short shared prefix. Should be placed near the root.
		{0, "8000000000000000000000000000000000000000000000000000000000000001", codes.OK},
		// Update a leaf node in the next epoch. Should be placed at the same level as the previous epoch.
		{1, "8000000000000000000000000000000000000000000000000000000000000001", codes.OK},
		{1, "0000000000000000000000000000000000000000000000000000000000000001", codes.OK},
	}
	for i, test := range tests {
		err := m.AddLeaf([]byte{}, test.epoch, test.index)
		if got, want := grpc.Code(err), test.code; got != want {
			t.Errorf("%v: AddLeaf(_, %v, %v)=%v, want %v, %v",
				i, test.epoch, test.index, got, want, err)
		}
	}
}

var letters = []rune("01234567890abcdef")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func BenchmarkAddLeaf(b *testing.B) {
	m := New()
	var epoch Epoch
	for i := 0; i < b.N; i++ {
		index := randSeq(64)
		err := m.AddLeaf([]byte{}, epoch, index)
		if got, want := grpc.Code(err), codes.OK; got != want {
			b.Errorf("%v: AddLeaf(_, %v, %v)=%v, want %v",
				i, epoch, index, got, want)
		}
	}
}

func BenchmarkAddLeafAdvanceEpoch(b *testing.B) {
	m := New()
	var epoch Epoch
	for i := 0; i < b.N; i++ {
		index := randSeq(64)
		epoch++
		err := m.AddLeaf([]byte{}, epoch, index)
		if got, want := grpc.Code(err), codes.OK; got != want {
			b.Errorf("%v: AddLeaf(_, %v, %v)=%v, want %v",
				i, epoch, index, got, want)
		}
	}
}

func BenchmarkAudit(b *testing.B) {
	m := New()
	var epoch Epoch
	items := make([]string, 0, b.N)
	for i := 0; i < b.N; i++ {
		index := randSeq(64)
		items = append(items, index)
		err := m.AddLeaf([]byte{}, epoch, index)
		if got, want := grpc.Code(err), codes.OK; got != want {
			b.Errorf("%v: AddLeaf(_, %v, %v)=%v, want %v",
				i, epoch, index, got, want)
		}
	}
	for _, v := range items {
		m.AuditPath(epoch, v)
	}
}

func TestPushDown(t *testing.T) {
	n := &node{index: BitString(AllZeros)}
	if !n.leaf() {
		t.Errorf("node without children was a leaf")
	}
	n.pushDown()
	if n.leaf() {
		t.Errorf("node was still a leaf after push")
	}
	if !n.left.leaf() {
		t.Errorf("new child was not a leaf after push")
	}
}

func TestCreateBranch(t *testing.T) {
	n := &node{index: BitString(AllZeros)}
	n.createBranch("0")
	if n.left == nil {
		t.Errorf("nil branch after create")
	}
}

func TestCreateBranchCOW(t *testing.T) {
	la := &node{epoch: 0, index: "0", depth: 1}
	lb := &node{epoch: 0, index: "1", depth: 1}
	r0 := &node{epoch: 0, index: "", left: la, right: lb}
	r1 := &node{epoch: 1, index: "", left: la, right: lb}

	var e0 Epoch
	var e1 Epoch = 1

	r1.createBranch("0")
	if got, want := r1.left.epoch, e1; got != want {
		t.Errorf("r1.left.epoch = %v, want %v", got, want)
	}
	if got, want := r0.left.epoch, e0; got != want {
		t.Errorf("r0.left.epoch = %v, want %v", got, want)
	}
}

func TestAuditDepth(t *testing.T) {
	m := New()
	tests := []struct {
		epoch Epoch
		index string
		depth int
	}{
		{0, "0000000000000000000000000000000000000000000000000000000000000000", 257},
		{0, "0000000000000000000000000000000000000000000000000000000000000001", 257},
		{0, "8000000000000000000000000000000000000000000000000000000000000001", 2},
		{1, "8000000000000000000000000000000000000000000000000000000000000001", 2},
		{1, "0000000000000000000000000000000000000000000000000000000000000001", 257},
	}
	for i, test := range tests {
		err := m.AddLeaf([]byte{}, test.epoch, test.index)
		if got, want := grpc.Code(err), codes.OK; got != want {
			t.Errorf("%v: AddLeaf(_, %v, %v)=%v, want %v",
				i, test.epoch, test.index, got, want)
		}
	}

	for i, test := range tests {
		audit, err := m.AuditPath(test.epoch, test.index)
		if got, want := grpc.Code(err), codes.OK; got != want {
			t.Errorf("%v: AuditPath(_, %v, %v)=%v, want %v",
				i, test.epoch, test.index, got, want)
		}
		if got, want := len(audit), test.depth; got != want {
			for j, a := range audit {
				fmt.Println(j, ": ", a)
			}
			t.Errorf("len(audit(%v, %v))=%v, want %v", test.epoch, test.index, got, want)
		}
	}
}

func TestAuditNeighors(t *testing.T) {
	m := New()
	tests := []struct {
		epoch         Epoch
		index         string
		emptyNeighors []bool
	}{
		{0, "0000000000000000000000000000000000000000000000000000000000000000", []bool{}},
		{0, "F000000000000000000000000000000000000000000000000000000000000000", []bool{false}},
		{0, "2000000000000000000000000000000000000000000000000000000000000000", []bool{false, true, false}},
		{0, "C000000000000000000000000000000000000000000000000000000000000000", []bool{false, true, false}},
	}
	for i, test := range tests {
		// Insert.
		err := m.AddLeaf([]byte{}, test.epoch, test.index)
		if got, want := grpc.Code(err), codes.OK; got != want {
			t.Errorf("%v: AddLeaf(_, %v, %v)=%v, want %v",
				i, test.epoch, test.index, got, want)
		}
		// Verify audit path.
		audit, err := m.AuditPath(test.epoch, test.index)
		if got, want := grpc.Code(err), codes.OK; got != want {
			t.Errorf("%v: AuditPath(_, %v, %v)=%v, want %v",
				i, test.epoch, test.index, got, want)
		}
		if got, want := len(audit), len(test.emptyNeighors)+1; got != want {
			for j, a := range audit {
				fmt.Println(j, ": ", a)
			}
			t.Errorf("len(audit(%v, %v))=%v, want %v", test.epoch, test.index, got, want)
		}
		for j, v := range test.emptyNeighors {
			// Starting from the leaf, going to the root. Skipping leaf.
			depth := len(audit) - 1 - j
			nstr := neighborOf(BitString(test.index), depth)
			if got, want := bytes.Equal(audit[j+1], EmptyValue(nstr)), v; got != want {
				t.Errorf("AuditPath(%v)[%v]=%v, want %v", test.index, j, got, want)
			}
		}
	}
}

func neighborOf(index string, depth int) string {
	return index[:depth-1] + string(neighbor(index[depth-1]))
}