// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data/msg.proto

package data

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TestMsg struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMsg) Reset()         { *m = TestMsg{} }
func (m *TestMsg) String() string { return proto.CompactTextString(m) }
func (*TestMsg) ProtoMessage()    {}
func (*TestMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_0453764a1c6a3baa, []int{0}
}

func (m *TestMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMsg.Unmarshal(m, b)
}
func (m *TestMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMsg.Marshal(b, m, deterministic)
}
func (m *TestMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMsg.Merge(m, src)
}
func (m *TestMsg) XXX_Size() int {
	return xxx_messageInfo_TestMsg.Size(m)
}
func (m *TestMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TestMsg proto.InternalMessageInfo

func (m *TestMsg) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *TestMsg) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type TestMsg2 struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMsg2) Reset()         { *m = TestMsg2{} }
func (m *TestMsg2) String() string { return proto.CompactTextString(m) }
func (*TestMsg2) ProtoMessage()    {}
func (*TestMsg2) Descriptor() ([]byte, []int) {
	return fileDescriptor_0453764a1c6a3baa, []int{1}
}

func (m *TestMsg2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMsg2.Unmarshal(m, b)
}
func (m *TestMsg2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMsg2.Marshal(b, m, deterministic)
}
func (m *TestMsg2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMsg2.Merge(m, src)
}
func (m *TestMsg2) XXX_Size() int {
	return xxx_messageInfo_TestMsg2.Size(m)
}
func (m *TestMsg2) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMsg2.DiscardUnknown(m)
}

var xxx_messageInfo_TestMsg2 proto.InternalMessageInfo

func (m *TestMsg2) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *TestMsg2) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*TestMsg)(nil), "data.TestMsg")
	proto.RegisterType((*TestMsg2)(nil), "data.TestMsg2")
}

func init() { proto.RegisterFile("data/msg.proto", fileDescriptor_0453764a1c6a3baa) }

var fileDescriptor_0453764a1c6a3baa = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x49, 0x2c, 0x49,
	0xd4, 0xcf, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xf1, 0x95, 0x0c,
	0xb9, 0xd8, 0x43, 0x52, 0x8b, 0x4b, 0x7c, 0x8b, 0xd3, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b,
	0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0xb0, 0x58, 0x65, 0x41, 0xaa, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x98, 0xad, 0x64, 0xc4, 0xc5, 0x01, 0xd5, 0x62, 0x44, 0xac,
	0x9e, 0x24, 0x36, 0xb0, 0x9d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x98, 0x09, 0x53,
	0x85, 0x00, 0x00, 0x00,
}
