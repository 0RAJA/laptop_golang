// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filter_message.proto

package pb

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

//寻找的computer的条件
type Filter struct {
	MaxPriceUsd          float64  `protobuf:"fixed64,1,opt,name=max_price_usd,json=maxPriceUsd,proto3" json:"max_price_usd,omitempty"`
	MinCpuCores          uint32   `protobuf:"varint,2,opt,name=min_cpu_cores,json=minCpuCores,proto3" json:"min_cpu_cores,omitempty"`
	MinCpuGhz            float64  `protobuf:"fixed64,3,opt,name=min_cpu_ghz,json=minCpuGhz,proto3" json:"min_cpu_ghz,omitempty"`
	MinRam               *Memory  `protobuf:"bytes,4,opt,name=min_ram,json=minRam,proto3" json:"min_ram,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_02dd12c5821a9fa1, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetMaxPriceUsd() float64 {
	if m != nil {
		return m.MaxPriceUsd
	}
	return 0
}

func (m *Filter) GetMinCpuCores() uint32 {
	if m != nil {
		return m.MinCpuCores
	}
	return 0
}

func (m *Filter) GetMinCpuGhz() float64 {
	if m != nil {
		return m.MinCpuGhz
	}
	return 0
}

func (m *Filter) GetMinRam() *Memory {
	if m != nil {
		return m.MinRam
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "rpc.proto.Filter")
}

func init() { proto.RegisterFile("filter_message.proto", fileDescriptor_02dd12c5821a9fa1) }

var fileDescriptor_02dd12c5821a9fa1 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xcb, 0xcc, 0x29,
	0x49, 0x2d, 0x8a, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x2c, 0x2a, 0x48, 0x86, 0x30, 0xa5, 0x44, 0x72, 0x53, 0x73, 0xf3, 0x8b, 0x2a, 0x51,
	0x15, 0x28, 0xcd, 0x61, 0xe4, 0x62, 0x73, 0x03, 0xeb, 0x14, 0x52, 0xe2, 0xe2, 0xcd, 0x4d, 0xac,
	0x88, 0x2f, 0x28, 0xca, 0x4c, 0x4e, 0x8d, 0x2f, 0x2d, 0x4e, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x0c, 0xe2, 0xce, 0x4d, 0xac, 0x08, 0x00, 0x89, 0x85, 0x16, 0xa7, 0x80, 0xd5, 0x64, 0xe6, 0xc5,
	0x27, 0x17, 0x94, 0xc6, 0x27, 0xe7, 0x17, 0xa5, 0x16, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x06,
	0x71, 0xe7, 0x66, 0xe6, 0x39, 0x17, 0x94, 0x3a, 0x83, 0x84, 0x84, 0xe4, 0xb8, 0xb8, 0x61, 0x6a,
	0xd2, 0x33, 0xaa, 0x24, 0x98, 0xc1, 0xa6, 0x70, 0x42, 0x54, 0xb8, 0x67, 0x54, 0x09, 0x69, 0x71,
	0xb1, 0x83, 0xe4, 0x8b, 0x12, 0x73, 0x25, 0x58, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x04, 0xf5, 0xe0,
	0xae, 0xd4, 0xf3, 0x05, 0x3b, 0x32, 0x88, 0x2d, 0x37, 0x33, 0x2f, 0x28, 0x31, 0xd7, 0x49, 0x85,
	0x4b, 0x26, 0x39, 0x3f, 0x57, 0x2f, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x28, 0xc0, 0x59,
	0x2f, 0x1d, 0xae, 0x36, 0xa9, 0x34, 0x2d, 0x80, 0x31, 0x8a, 0xa9, 0x20, 0x29, 0x89, 0x0d, 0xcc,
	0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x68, 0x0c, 0x7b, 0x25, 0x04, 0x01, 0x00, 0x00,
}
