// Code generated by protoc-gen-go.
// source: snapshot.proto
// DO NOT EDIT!

/*
Package storage is a generated protocol buffer package.

It is generated from these files:
	snapshot.proto

It has these top-level messages:
	Snapshot
*/
package storage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Snapshot represents a snapshot of the blockchain, including the state
// tree and issuance memory.
type Snapshot struct {
	// Nodes contains every node within the state tree, including interior nodes.
	// The nodes are ordered according to a pre-order traversal.
	Nodes []*Snapshot_StateTreeNode `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	// Issuances contains the record of recent issuances for ensuring uniqueness
	// of issuances.
	Issuances []*Snapshot_Issuance `protobuf:"bytes,2,rep,name=issuances" json:"issuances,omitempty"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (m *Snapshot) String() string            { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Snapshot) GetNodes() []*Snapshot_StateTreeNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *Snapshot) GetIssuances() []*Snapshot_Issuance {
	if m != nil {
		return m.Issuances
	}
	return nil
}

type Snapshot_Issuance struct {
	Hash     []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ExpiryMs uint64 `protobuf:"varint,2,opt,name=expiry_ms,json=expiryMs" json:"expiry_ms,omitempty"`
}

func (m *Snapshot_Issuance) Reset()                    { *m = Snapshot_Issuance{} }
func (m *Snapshot_Issuance) String() string            { return proto.CompactTextString(m) }
func (*Snapshot_Issuance) ProtoMessage()               {}
func (*Snapshot_Issuance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Snapshot_StateTreeNode struct {
	Key  []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Leaf bool   `protobuf:"varint,2,opt,name=leaf" json:"leaf,omitempty"`
	Hash []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *Snapshot_StateTreeNode) Reset()                    { *m = Snapshot_StateTreeNode{} }
func (m *Snapshot_StateTreeNode) String() string            { return proto.CompactTextString(m) }
func (*Snapshot_StateTreeNode) ProtoMessage()               {}
func (*Snapshot_StateTreeNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func init() {
	proto.RegisterType((*Snapshot)(nil), "chain.core.txdb.internal.storage.Snapshot")
	proto.RegisterType((*Snapshot_Issuance)(nil), "chain.core.txdb.internal.storage.Snapshot.Issuance")
	proto.RegisterType((*Snapshot_StateTreeNode)(nil), "chain.core.txdb.internal.storage.Snapshot.StateTreeNode")
}

func init() { proto.RegisterFile("snapshot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x90, 0x3f, 0x4b, 0x03, 0x41,
	0x10, 0xc5, 0xb9, 0x5c, 0xd4, 0xbb, 0xf1, 0x0f, 0xb2, 0x95, 0xc4, 0x26, 0x58, 0x5d, 0x35, 0x85,
	0x69, 0x04, 0x3b, 0xbb, 0x14, 0x06, 0xdc, 0x58, 0xd9, 0xc8, 0xe4, 0x32, 0x7a, 0x8b, 0x71, 0xf7,
	0xd8, 0x59, 0x21, 0xf9, 0x60, 0x7e, 0x3f, 0xef, 0x36, 0x1b, 0x83, 0x95, 0xa4, 0x7b, 0x0c, 0xfc,
	0x7e, 0xef, 0x31, 0x70, 0x21, 0x96, 0x5a, 0x69, 0x5c, 0xc0, 0xd6, 0xbb, 0xe0, 0xd4, 0xb8, 0x6e,
	0xc8, 0x58, 0xac, 0x9d, 0x67, 0x0c, 0xeb, 0xe5, 0x02, 0x8d, 0x0d, 0xec, 0x2d, 0xad, 0x50, 0x82,
	0xf3, 0xf4, 0xce, 0x37, 0xdf, 0x03, 0x28, 0xe6, 0x09, 0x52, 0x33, 0x38, 0xb2, 0x6e, 0xc9, 0x72,
	0x95, 0x8d, 0xf3, 0xea, 0xf4, 0xf6, 0x0e, 0xff, 0xc3, 0x71, 0x87, 0xe2, 0x3c, 0x50, 0xe0, 0x67,
	0xcf, 0x3c, 0xeb, 0x04, 0x7a, 0xab, 0x51, 0x4f, 0x50, 0x1a, 0x91, 0x2f, 0xb2, 0x75, 0xe7, 0x1c,
	0x44, 0xe7, 0xe4, 0x00, 0xe7, 0x34, 0xb1, 0x7a, 0x6f, 0x19, 0xdd, 0x43, 0xb1, 0x3b, 0x2b, 0x05,
	0xc3, 0x86, 0xa4, 0xe9, 0xd6, 0x66, 0xd5, 0x99, 0x8e, 0x59, 0x5d, 0x43, 0xc9, 0xeb, 0xd6, 0xf8,
	0xcd, 0xeb, 0x67, 0x5f, 0x99, 0x55, 0x43, 0x5d, 0x6c, 0x0f, 0x8f, 0x32, 0x9a, 0xc2, 0xf9, 0x9f,
	0x9d, 0xea, 0x12, 0xf2, 0x0f, 0xde, 0x24, 0x41, 0x1f, 0x7b, 0xe7, 0x8a, 0xe9, 0x2d, 0xa2, 0x85,
	0x8e, 0xf9, 0xb7, 0x27, 0xdf, 0xf7, 0x3c, 0x94, 0x2f, 0x27, 0x69, 0xef, 0xe2, 0x38, 0xfe, 0x7a,
	0xf2, 0x13, 0x00, 0x00, 0xff, 0xff, 0x84, 0x5f, 0x0a, 0x54, 0x7d, 0x01, 0x00, 0x00,
}
