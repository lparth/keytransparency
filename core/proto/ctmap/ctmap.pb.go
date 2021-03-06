// Code generated by protoc-gen-go.
// source: ctmap.proto
// DO NOT EDIT!

/*
Package ctmap is a generated protocol buffer package.

It is generated from these files:
	ctmap.proto

It has these top-level messages:
	MapHead
	SignedMapHead
*/
package ctmap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import signature "github.com/google/keytransparency/core/proto/signature"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MapHead is the head node of the Merkle Tree as well as additional metadata
// for the tree.
type MapHead struct {
	// realm is the domain identifier for the transparent map.
	Realm string `protobuf:"bytes,1,opt,name=realm" json:"realm,omitempty"`
	// epoch is the epoch number of this map head.
	Epoch int64 `protobuf:"varint,2,opt,name=epoch" json:"epoch,omitempty"`
	// root is the value of the root node of the Merkle tree.
	Root []byte `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	// issue_time is the time when this epoch was created. Monotonically
	// increasing.
	IssueTime *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=issue_time,json=issueTime" json:"issue_time,omitempty"`
}

func (m *MapHead) Reset()                    { *m = MapHead{} }
func (m *MapHead) String() string            { return proto.CompactTextString(m) }
func (*MapHead) ProtoMessage()               {}
func (*MapHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MapHead) GetIssueTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.IssueTime
	}
	return nil
}

// SignedMapHead represents a signed state of the Merkel Tree.
type SignedMapHead struct {
	// map_head contains the head node of the Merkle Tree along with other
	// metadata.
	MapHead *MapHead `protobuf:"bytes,1,opt,name=map_head,json=mapHead" json:"map_head,omitempty"`
	// signatures is a set of map_head signatures. Each signature is identified by
	// the first 64 bits of the public key that verifies it.
	Signatures map[string]*signature.DigitallySigned `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SignedMapHead) Reset()                    { *m = SignedMapHead{} }
func (m *SignedMapHead) String() string            { return proto.CompactTextString(m) }
func (*SignedMapHead) ProtoMessage()               {}
func (*SignedMapHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignedMapHead) GetMapHead() *MapHead {
	if m != nil {
		return m.MapHead
	}
	return nil
}

func (m *SignedMapHead) GetSignatures() map[string]*signature.DigitallySigned {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func init() {
	proto.RegisterType((*MapHead)(nil), "ctmap.MapHead")
	proto.RegisterType((*SignedMapHead)(nil), "ctmap.SignedMapHead")
}

func init() { proto.RegisterFile("ctmap.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4f, 0xeb, 0x30,
	0x14, 0x85, 0xe5, 0xa6, 0x7d, 0x7d, 0xbd, 0x79, 0x0f, 0x90, 0xc5, 0x10, 0x65, 0x21, 0xaa, 0x18,
	0xc2, 0x92, 0xa0, 0xb2, 0x00, 0x73, 0x41, 0x2c, 0x2c, 0x86, 0x85, 0xa9, 0x72, 0xd3, 0x4b, 0x6a,
	0x35, 0x8e, 0x23, 0xc7, 0x41, 0xca, 0xce, 0xbf, 0xe4, 0xcf, 0xa0, 0xd8, 0x2e, 0x05, 0xb6, 0x7b,
	0x6e, 0xbe, 0x5c, 0x9f, 0x73, 0x20, 0x2c, 0x8c, 0xe4, 0x4d, 0xd6, 0x68, 0x65, 0x14, 0x9d, 0x58,
	0x11, 0x9f, 0x95, 0x4a, 0x95, 0x15, 0xe6, 0x76, 0xb9, 0xee, 0x5e, 0x73, 0x23, 0x24, 0xb6, 0x86,
	0x4b, 0xcf, 0xc5, 0xf7, 0xa5, 0x30, 0xdb, 0x6e, 0x9d, 0x15, 0x4a, 0xe6, 0x9e, 0xdd, 0x61, 0x6f,
	0x34, 0xaf, 0xdb, 0x86, 0x6b, 0xac, 0x8b, 0x3e, 0x2f, 0x94, 0xf6, 0x07, 0xf2, 0x56, 0x94, 0x35,
	0x37, 0x9d, 0xc6, 0xc3, 0xe4, 0xee, 0xcc, 0xdf, 0x09, 0x4c, 0x1f, 0x79, 0xf3, 0x80, 0x7c, 0x43,
	0x4f, 0x61, 0xa2, 0x91, 0x57, 0x32, 0x22, 0x09, 0x49, 0x67, 0xcc, 0x89, 0x61, 0x8b, 0x8d, 0x2a,
	0xb6, 0xd1, 0x28, 0x21, 0x69, 0xc0, 0x9c, 0xa0, 0x14, 0xc6, 0x5a, 0x29, 0x13, 0x05, 0x09, 0x49,
	0xff, 0x31, 0x3b, 0xd3, 0x1b, 0x00, 0xd1, 0xb6, 0x1d, 0xae, 0x06, 0xb3, 0xd1, 0x38, 0x21, 0x69,
	0xb8, 0x88, 0x33, 0xe7, 0x2e, 0xdb, 0x27, 0xc9, 0x9e, 0xf7, 0x49, 0xd8, 0xcc, 0xd2, 0x83, 0x9e,
	0x7f, 0x10, 0xf8, 0xff, 0x24, 0xca, 0x1a, 0x37, 0x7b, 0x33, 0x17, 0xf0, 0x57, 0xf2, 0x66, 0xb5,
	0x45, 0xbe, 0xb1, 0x7e, 0xc2, 0xc5, 0x51, 0xe6, 0x8a, 0xf2, 0x04, 0x9b, 0x4a, 0x8f, 0x2e, 0x01,
	0xbe, 0x62, 0xb5, 0xd1, 0x28, 0x09, 0xd2, 0x70, 0x71, 0xee, 0xe1, 0x1f, 0x47, 0xad, 0x72, 0xd8,
	0x5d, 0x6d, 0x74, 0xcf, 0xbe, 0xfd, 0x17, 0xbf, 0xc0, 0xf1, 0xaf, 0xcf, 0xf4, 0x04, 0x82, 0x1d,
	0xf6, 0xbe, 0x8e, 0x61, 0xa4, 0x97, 0x30, 0x79, 0xe3, 0x55, 0x87, 0xb6, 0x8c, 0x21, 0xdd, 0xa1,
	0xcf, 0xa5, 0x28, 0x85, 0xe1, 0x55, 0xd5, 0xbb, 0x27, 0x99, 0x03, 0x6f, 0x47, 0xd7, 0x64, 0xfd,
	0xc7, 0x86, 0xbf, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x30, 0x45, 0x8e, 0x6b, 0xea, 0x01, 0x00,
	0x00,
}
