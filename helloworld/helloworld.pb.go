// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	BackupRequest
	HelloReply
	BeatRequest
	BeatReply
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
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Ismapper   bool   `protobuf:"varint,2,opt,name=ismapper" json:"ismapper,omitempty"`
	ShardSize  int64  `protobuf:"varint,3,opt,name=shard_size,json=shardSize" json:"shard_size,omitempty"`
	Offset     int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	NoReducers int64  `protobuf:"varint,5,opt,name=no_reducers,json=noReducers" json:"no_reducers,omitempty"`
	Fail       int64  `protobuf:"varint,6,opt,name=fail" json:"fail,omitempty"`
	FileNumber int64  `protobuf:"varint,7,opt,name=file_number,json=fileNumber" json:"file_number,omitempty"`
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

func (m *HelloRequest) GetShardSize() int64 {
	if m != nil {
		return m.ShardSize
	}
	return 0
}

func (m *HelloRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *HelloRequest) GetNoReducers() int64 {
	if m != nil {
		return m.NoReducers
	}
	return 0
}

func (m *HelloRequest) GetFail() int64 {
	if m != nil {
		return m.Fail
	}
	return 0
}

func (m *HelloRequest) GetFileNumber() int64 {
	if m != nil {
		return m.FileNumber
	}
	return 0
}

type BackupRequest struct {
	BackupRequest []*HelloRequest `protobuf:"bytes,1,rep,name=backup_request,json=backupRequest" json:"backup_request,omitempty"`
}

func (m *BackupRequest) Reset()                    { *m = BackupRequest{} }
func (m *BackupRequest) String() string            { return proto.CompactTextString(m) }
func (*BackupRequest) ProtoMessage()               {}
func (*BackupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BackupRequest) GetBackupRequest() []*HelloRequest {
	if m != nil {
		return m.BackupRequest
	}
	return nil
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The request message containing the user's name.
type BeatRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *BeatRequest) Reset()                    { *m = BeatRequest{} }
func (m *BeatRequest) String() string            { return proto.CompactTextString(m) }
func (*BeatRequest) ProtoMessage()               {}
func (*BeatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BeatRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type BeatReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *BeatReply) Reset()                    { *m = BeatReply{} }
func (m *BeatReply) String() string            { return proto.CompactTextString(m) }
func (*BeatReply) ProtoMessage()               {}
func (*BeatReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BeatReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*BackupRequest)(nil), "helloworld.BackupRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*BeatRequest)(nil), "helloworld.BeatRequest")
	proto.RegisterType((*BeatReply)(nil), "helloworld.BeatReply")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x74, 0xe9, 0xc7, 0x94, 0x5d, 0x90, 0x25, 0x16, 0x53, 0x84, 0x28, 0x91, 0x40, 0x3d,
	0xa5, 0xd2, 0x22, 0xae, 0x20, 0xc2, 0x81, 0x3d, 0xa1, 0x2a, 0x3d, 0x20, 0x71, 0x89, 0x9c, 0x64,
	0xe2, 0x46, 0xeb, 0xc4, 0xc6, 0x4e, 0x04, 0xdd, 0xbf, 0xc8, 0x85, 0x9f, 0x84, 0xe2, 0x34, 0xac,
	0x11, 0xbb, 0xbd, 0xcd, 0xbc, 0x8f, 0xd1, 0xf3, 0x93, 0xe1, 0xf1, 0x0e, 0xa5, 0x54, 0x3f, 0x94,
	0x91, 0x79, 0xa8, 0x8d, 0x6a, 0x14, 0x85, 0x1b, 0x24, 0xf8, 0x45, 0xe0, 0xe1, 0x65, 0xb7, 0xc6,
	0xf8, 0xbd, 0x45, 0xdb, 0x50, 0x0a, 0x27, 0x35, 0xaf, 0x90, 0x91, 0x25, 0x59, 0xcd, 0x62, 0x37,
	0xd3, 0x05, 0x4c, 0x4b, 0x5b, 0x71, 0xad, 0xd1, 0xb0, 0xfb, 0x4b, 0xb2, 0x9a, 0xc6, 0x7f, 0x77,
	0xfa, 0x02, 0xc0, 0xee, 0xb8, 0xc9, 0x13, 0x5b, 0x5e, 0x23, 0x1b, 0x2d, 0xc9, 0x6a, 0x14, 0xcf,
	0x1c, 0xb2, 0x2d, 0xaf, 0x91, 0x9e, 0xc3, 0x58, 0x15, 0x85, 0xc5, 0x86, 0x9d, 0x38, 0xea, 0xb0,
	0xd1, 0x97, 0x30, 0xaf, 0x55, 0x62, 0x30, 0x6f, 0x33, 0x34, 0x96, 0x3d, 0x70, 0x24, 0xd4, 0x2a,
	0x3e, 0x20, 0x5d, 0x8e, 0x82, 0x97, 0x92, 0x8d, 0x1d, 0xe3, 0xe6, 0xce, 0x54, 0x94, 0x12, 0x93,
	0xba, 0xad, 0x52, 0x34, 0x6c, 0xd2, 0x9b, 0x3a, 0xe8, 0x8b, 0x43, 0x82, 0x0d, 0x9c, 0x46, 0x3c,
	0xbb, 0x6a, 0xf5, 0xf0, 0x9a, 0x0f, 0x70, 0x96, 0x3a, 0x20, 0x31, 0x3d, 0xc2, 0xc8, 0x72, 0xb4,
	0x9a, 0x5f, 0xb0, 0xd0, 0x6b, 0xc5, 0x7f, 0x7f, 0x7c, 0x9a, 0xfa, 0x07, 0x82, 0x37, 0x00, 0x07,
	0x5a, 0xcb, 0x3d, 0x65, 0x30, 0xa9, 0xd0, 0x5a, 0x2e, 0x86, 0x7e, 0x86, 0x35, 0x78, 0x05, 0xf3,
	0x08, 0x79, 0x73, 0xa4, 0xc5, 0xe0, 0x35, 0xcc, 0x7a, 0xc9, 0xd1, 0x4b, 0x17, 0xbf, 0x09, 0x4c,
	0x3e, 0x1b, 0xc4, 0x06, 0x0d, 0x7d, 0x0f, 0xd3, 0x2d, 0xdf, 0xbb, 0x00, 0xf4, 0xce, 0xc8, 0x8b,
	0xf3, 0x5b, 0x18, 0x2d, 0xf7, 0xc1, 0x3d, 0x1a, 0xc1, 0xd9, 0xa7, 0x1d, 0x66, 0x57, 0x97, 0xc8,
	0x4d, 0x93, 0x22, 0x6f, 0xe8, 0x53, 0x5f, 0xeb, 0x25, 0x5e, 0x3c, 0xf9, 0x9f, 0xe8, 0x6f, 0x7c,
	0x04, 0xd8, 0x62, 0x9d, 0xf7, 0xbd, 0xd2, 0x67, 0xff, 0xc8, 0xfc, 0xaa, 0xee, 0x8e, 0x11, 0x09,
	0x78, 0x5e, 0xaa, 0x50, 0x18, 0x9d, 0x85, 0xf8, 0x93, 0x57, 0x5a, 0xa2, 0xf5, 0xb4, 0xd1, 0x23,
	0x27, 0xfe, 0xda, 0xcd, 0x9b, 0xee, 0x83, 0x6e, 0xc8, 0xb7, 0x77, 0x42, 0x29, 0x21, 0x31, 0x14,
	0x4a, 0xf2, 0x5a, 0x84, 0xca, 0x88, 0x75, 0x67, 0x5f, 0x0f, 0xf6, 0xf5, 0x8d, 0xdd, 0x1b, 0xd3,
	0xb1, 0xfb, 0xe0, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xf3, 0xd8, 0x7b, 0xf4, 0x02,
	0x00, 0x00,
}
