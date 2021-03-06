// Code generated by protoc-gen-go.
// source: store.proto
// DO NOT EDIT!

/*
Package store is a generated protocol buffer package.

It is generated from these files:
	store.proto

It has these top-level messages:
	UniqueStorageValue
	MessageValue
	UniqueStorage
*/
package store

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

type MessageValue_OPERATION int32

const (
	MessageValue_CAS MessageValue_OPERATION = 0
)

var MessageValue_OPERATION_name = map[int32]string{
	0: "CAS",
}
var MessageValue_OPERATION_value = map[string]int32{
	"CAS": 0,
}

func (x MessageValue_OPERATION) String() string {
	return proto.EnumName(MessageValue_OPERATION_name, int32(x))
}
func (MessageValue_OPERATION) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type UniqueStorageValue struct {
	Value      string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Expiration uint32 `protobuf:"varint,2,opt,name=expiration" json:"expiration,omitempty"`
}

func (m *UniqueStorageValue) Reset()                    { *m = UniqueStorageValue{} }
func (m *UniqueStorageValue) String() string            { return proto.CompactTextString(m) }
func (*UniqueStorageValue) ProtoMessage()               {}
func (*UniqueStorageValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UniqueStorageValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *UniqueStorageValue) GetExpiration() uint32 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type MessageValue struct {
	Operation  MessageValue_OPERATION `protobuf:"varint,1,opt,name=operation,enum=store.MessageValue_OPERATION" json:"operation,omitempty"`
	Key        string                 `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value      string                 `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Expiration uint32                 `protobuf:"varint,4,opt,name=expiration" json:"expiration,omitempty"`
}

func (m *MessageValue) Reset()                    { *m = MessageValue{} }
func (m *MessageValue) String() string            { return proto.CompactTextString(m) }
func (*MessageValue) ProtoMessage()               {}
func (*MessageValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MessageValue) GetOperation() MessageValue_OPERATION {
	if m != nil {
		return m.Operation
	}
	return MessageValue_CAS
}

func (m *MessageValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MessageValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *MessageValue) GetExpiration() uint32 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type UniqueStorage struct {
	Items map[string]*UniqueStorageValue `protobuf:"bytes,1,rep,name=items" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UniqueStorage) Reset()                    { *m = UniqueStorage{} }
func (m *UniqueStorage) String() string            { return proto.CompactTextString(m) }
func (*UniqueStorage) ProtoMessage()               {}
func (*UniqueStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UniqueStorage) GetItems() map[string]*UniqueStorageValue {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*UniqueStorageValue)(nil), "store.UniqueStorageValue")
	proto.RegisterType((*MessageValue)(nil), "store.MessageValue")
	proto.RegisterType((*UniqueStorage)(nil), "store.UniqueStorage")
	proto.RegisterEnum("store.MessageValue_OPERATION", MessageValue_OPERATION_name, MessageValue_OPERATION_value)
}

func init() { proto.RegisterFile("store.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2e, 0xc9, 0x2f,
	0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xbc, 0xb8, 0x84, 0x42,
	0xf3, 0x32, 0x0b, 0x4b, 0x53, 0x83, 0x4b, 0xf2, 0x8b, 0x12, 0xd3, 0x53, 0xc3, 0x12, 0x73, 0x4a,
	0x53, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x40, 0x0c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08,
	0x47, 0x48, 0x8e, 0x8b, 0x2b, 0xb5, 0xa2, 0x20, 0xb3, 0x28, 0xb1, 0x24, 0x33, 0x3f, 0x4f, 0x82,
	0x49, 0x81, 0x51, 0x83, 0x37, 0x08, 0x49, 0x44, 0x69, 0x25, 0x23, 0x17, 0x8f, 0x6f, 0x6a, 0x71,
	0x31, 0xdc, 0x18, 0x6b, 0x2e, 0xce, 0xfc, 0x82, 0x54, 0xa8, 0x7a, 0x90, 0x51, 0x7c, 0x46, 0xb2,
	0x7a, 0x10, 0x47, 0x20, 0xab, 0xd3, 0xf3, 0x0f, 0x70, 0x0d, 0x72, 0x0c, 0xf1, 0xf4, 0xf7, 0x0b,
	0x42, 0xa8, 0x17, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x04, 0x5b, 0xc3, 0x19, 0x04, 0x62, 0x22,
	0x5c, 0xc5, 0x8c, 0xdb, 0x55, 0x2c, 0x18, 0xae, 0x12, 0xe1, 0xe2, 0x84, 0x9b, 0x2f, 0xc4, 0xce,
	0xc5, 0xec, 0xec, 0x18, 0x2c, 0xc0, 0xa0, 0x34, 0x9b, 0x91, 0x8b, 0x17, 0xc5, 0xe3, 0x42, 0xa6,
	0x5c, 0xac, 0x99, 0x25, 0xa9, 0xb9, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xf2, 0x50,
	0x87, 0xa2, 0x28, 0xd2, 0xf3, 0x04, 0xa9, 0x70, 0xcd, 0x2b, 0x29, 0xaa, 0x0c, 0x82, 0xa8, 0x96,
	0x0a, 0xe6, 0xe2, 0x42, 0x08, 0xc2, 0x1c, 0xcd, 0x88, 0x70, 0xb4, 0x3e, 0xcc, 0xd1, 0x20, 0x8f,
	0x70, 0x1b, 0x49, 0x62, 0x33, 0x16, 0x1c, 0x0a, 0x50, 0xff, 0x58, 0x31, 0x59, 0x30, 0x26, 0xb1,
	0x81, 0xe3, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xe5, 0xfa, 0x6e, 0xb2, 0x01, 0x00,
	0x00,
}
