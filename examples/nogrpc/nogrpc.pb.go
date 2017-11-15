// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/stevvooe/protobuild/examples/nogrpc/nogrpc.proto

/*
Package nogrpc is a generated protocol buffer package.

It is generated from these files:
	github.com/stevvooe/protobuild/examples/nogrpc/nogrpc.proto

It has these top-level messages:
	Thing
*/
package nogrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Thing struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Thing) Reset()                    { *m = Thing{} }
func (m *Thing) String() string            { return proto.CompactTextString(m) }
func (*Thing) ProtoMessage()               {}
func (*Thing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Thing) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Thing)(nil), "protobuild.example.nogrpc.Thing")
}

func init() {
	proto.RegisterFile("github.com/stevvooe/protobuild/examples/nogrpc/nogrpc.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x2e, 0x49, 0x2d, 0x2b, 0xcb, 0xcf, 0x4f, 0xd5,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0xcd, 0xcc, 0x49, 0xd1, 0x4f, 0xad, 0x48, 0xcc, 0x2d,
	0xc8, 0x49, 0x2d, 0xd6, 0xcf, 0xcb, 0x4f, 0x2f, 0x2a, 0x48, 0x86, 0x52, 0x7a, 0x60, 0x15, 0x42,
	0x92, 0x08, 0x85, 0x7a, 0x50, 0x85, 0x7a, 0x10, 0x05, 0x52, 0xf2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39,
	0x70, 0xa3, 0xd2, 0xf4, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x20, 0x7a, 0x95,
	0x1c, 0xb9, 0x58, 0x43, 0x32, 0x32, 0xf3, 0xd2, 0x85, 0x2c, 0xb8, 0x38, 0xe1, 0x72, 0x12, 0xcc,
	0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x52, 0x7a, 0x10, 0xdd, 0x7a, 0x30, 0xdd, 0x7a, 0x21, 0x30, 0x15,
	0x41, 0x08, 0xc5, 0x4e, 0x16, 0x51, 0x66, 0x24, 0xba, 0xde, 0x1a, 0x42, 0x27, 0xb1, 0x81, 0xd5,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xf5, 0xe5, 0x70, 0xfe, 0x00, 0x00, 0x00,
}