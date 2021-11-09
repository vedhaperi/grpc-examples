// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld/helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld/helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld

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

// The request message containing the user's name.
type HelloRequest struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Ismapper  bool   `protobuf:"varint,2,opt,name=ismapper" json:"ismapper,omitempty"`
	ShardSize int32  `protobuf:"varint,3,opt,name=shard_size,json=shardSize" json:"shard_size,omitempty"`
	Offset    int32  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetIsmapper() bool {
	if m != nil {
		return m.Ismapper
	}
	return false
}

func (m *HelloRequest) GetShardSize() int32 {
	if m != nil {
		return m.ShardSize
	}
	return 0
}

func (m *HelloRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("helloworld/helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0xd6, 0x36, 0x1d, 0x04, 0x61, 0x0e, 0x65, 0x69, 0x11, 0x42, 0x0e, 0x92, 0xd3,
	0x06, 0x14, 0xaf, 0x1e, 0x7a, 0x51, 0x6f, 0x25, 0x3d, 0x08, 0x5e, 0x64, 0xb5, 0xd3, 0x6d, 0x60,
	0xd3, 0x59, 0x77, 0x53, 0x34, 0xfd, 0xf5, 0x92, 0xa5, 0x69, 0x73, 0xe8, 0xed, 0x7d, 0x33, 0xf3,
	0x78, 0xbb, 0x0f, 0x66, 0x1b, 0x32, 0x86, 0x7f, 0xd9, 0x99, 0x55, 0x7e, 0x92, 0xd2, 0x3a, 0xae,
	0x19, 0xe1, 0x34, 0x49, 0x77, 0x70, 0xf3, 0xda, 0x52, 0x41, 0x3f, 0x3b, 0xf2, 0x35, 0x22, 0x0c,
	0xb6, 0xaa, 0x22, 0x11, 0x25, 0x51, 0x36, 0x2e, 0x82, 0xc6, 0x29, 0xc4, 0xa5, 0xaf, 0x94, 0xb5,
	0xe4, 0xc4, 0x65, 0x12, 0x65, 0x71, 0x71, 0x64, 0xbc, 0x03, 0xf0, 0x1b, 0xe5, 0x56, 0x9f, 0xbe,
	0xdc, 0x93, 0xb8, 0x4a, 0xa2, 0xec, 0xba, 0x18, 0x87, 0xc9, 0xb2, 0xdc, 0x13, 0x4e, 0x60, 0xc8,
	0xeb, 0xb5, 0xa7, 0x5a, 0x0c, 0xc2, 0xea, 0x40, 0xe9, 0x3d, 0xc0, 0x21, 0xd6, 0x9a, 0x06, 0x05,
	0x8c, 0x2a, 0xf2, 0x5e, 0xe9, 0x2e, 0xb7, 0xc3, 0x87, 0x37, 0x18, 0xbd, 0x38, 0xa2, 0x9a, 0x1c,
	0x3e, 0x43, 0xbc, 0x54, 0x4d, 0x70, 0xa1, 0x90, 0xbd, 0x4f, 0xf5, 0xdf, 0x3f, 0x9d, 0x9c, 0xd9,
	0x58, 0xd3, 0xa4, 0x17, 0x73, 0x0d, 0xb3, 0x92, 0xa5, 0x76, 0xf6, 0x5b, 0xd2, 0x9f, 0xaa, 0xac,
	0x21, 0xdf, 0xbb, 0x9d, 0xdf, 0x86, 0xe3, 0xf7, 0x56, 0x2f, 0xda, 0x96, 0x16, 0xd1, 0xc7, 0x93,
	0x66, 0xd6, 0x86, 0xa4, 0x66, 0xa3, 0xb6, 0x5a, 0xb2, 0xd3, 0x79, 0x6b, 0xcf, 0x3b, 0x7b, 0x7e,
	0xb6, 0xe4, 0xaf, 0x61, 0x68, 0xf9, 0xf1, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x57, 0xa0, 0x56, 0x44,
	0x84, 0x01, 0x00, 0x00,
}
