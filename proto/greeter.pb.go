// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

/*
Package greeter is a generated protocol buffer package.

It is generated from these files:
	greeter.proto

It has these top-level messages:
	GreetRequest
	GreetResponse
*/
package greeter

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

type GreetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GreetRequest) Reset()                    { *m = GreetRequest{} }
func (m *GreetRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()               {}
func (*GreetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GreetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetResponse struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetResponse) Reset()                    { *m = GreetResponse{} }
func (m *GreetResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()               {}
func (*GreetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GreetResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "sh.kono.examples.greeter.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "sh.kono.examples.greeter.GreetResponse")
}

func init() { proto.RegisterFile("greeter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x28, 0xce, 0xd0, 0xcb, 0xce,
	0xcf, 0xcb, 0xd7, 0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x83, 0xca, 0x2b, 0x29,
	0x71, 0xf1, 0xb8, 0x83, 0x98, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c,
	0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x36, 0x17,
	0x2f, 0x54, 0x4d, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x14, 0x17, 0x07, 0x58, 0x7f, 0x66,
	0x5e, 0x3a, 0x54, 0x21, 0x9c, 0x6f, 0x94, 0xcc, 0xc5, 0xee, 0x0e, 0x31, 0x5b, 0x28, 0x82, 0x8b,
	0x15, 0xcc, 0x14, 0x52, 0xd3, 0xc3, 0x65, 0xbf, 0x1e, 0xb2, 0xe5, 0x52, 0xea, 0x04, 0xd5, 0x41,
	0x1c, 0xe0, 0xc4, 0x19, 0xc5, 0x0e, 0x95, 0x48, 0x62, 0x03, 0xfb, 0xd0, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x72, 0xa8, 0xfb, 0xc1, 0xf2, 0x00, 0x00, 0x00,
}