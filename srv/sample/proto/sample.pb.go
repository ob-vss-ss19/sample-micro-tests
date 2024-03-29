// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/sample/proto/sample.proto

package sample

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Item struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_1501173291aba795, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1501173291aba795, []int{1}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResult struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResult) Reset()         { *m = ListResult{} }
func (m *ListResult) String() string { return proto.CompactTextString(m) }
func (*ListResult) ProtoMessage()    {}
func (*ListResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1501173291aba795, []int{2}
}

func (m *ListResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResult.Unmarshal(m, b)
}
func (m *ListResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResult.Marshal(b, m, deterministic)
}
func (m *ListResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResult.Merge(m, src)
}
func (m *ListResult) XXX_Size() int {
	return xxx_messageInfo_ListResult.Size(m)
}
func (m *ListResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResult.DiscardUnknown(m)
}

var xxx_messageInfo_ListResult proto.InternalMessageInfo

func (m *ListResult) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "sample.Item")
	proto.RegisterType((*ListRequest)(nil), "sample.ListRequest")
	proto.RegisterType((*ListResult)(nil), "sample.ListResult")
}

func init() { proto.RegisterFile("srv/sample/proto/sample.proto", fileDescriptor_1501173291aba795) }

var fileDescriptor_1501173291aba795 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x2e, 0x2a, 0xd3,
	0x2f, 0x4e, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0x72, 0xf4, 0xc0,
	0x1c, 0x21, 0x36, 0x08, 0x4f, 0x49, 0x8b, 0x8b, 0xc5, 0xb3, 0x24, 0x35, 0x57, 0x88, 0x8f, 0x8b,
	0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b,
	0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0xe2, 0xe5,
	0xe2, 0xf6, 0xc9, 0x2c, 0x2e, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x32, 0xe0, 0xe2,
	0x82, 0x70, 0x8b, 0x4b, 0x73, 0x4a, 0x84, 0x94, 0xb8, 0x58, 0x33, 0x4b, 0x52, 0x73, 0x8b, 0x25,
	0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x78, 0xf4, 0xa0, 0xd6, 0x81, 0x4c, 0x0f, 0x82, 0x48, 0x19,
	0x59, 0x72, 0xb1, 0x05, 0x83, 0x45, 0x85, 0xf4, 0xb9, 0x58, 0x40, 0x7a, 0x85, 0x84, 0x61, 0xca,
	0x90, 0x0c, 0x96, 0x12, 0x42, 0x15, 0x04, 0x19, 0x9f, 0xc4, 0x06, 0x76, 0xb6, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x1c, 0x2b, 0xac, 0x00, 0xd7, 0x00, 0x00, 0x00,
}
