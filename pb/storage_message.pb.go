// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage_message.proto

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

type Storage_Driver int32

const (
	Storage_UNKNOWN Storage_Driver = 0
	Storage_HDD     Storage_Driver = 1
	Storage_SSD     Storage_Driver = 2
)

var Storage_Driver_name = map[int32]string{
	0: "UNKNOWN",
	1: "HDD",
	2: "SSD",
}

var Storage_Driver_value = map[string]int32{
	"UNKNOWN": 0,
	"HDD":     1,
	"SSD":     2,
}

func (x Storage_Driver) String() string {
	return proto.EnumName(Storage_Driver_name, int32(x))
}

func (Storage_Driver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_170f09d838bd8a04, []int{0, 0}
}

//存储消息
type Storage struct {
	Driver               Storage_Driver `protobuf:"varint,1,opt,name=driver,proto3,enum=rpc.proto.Storage_Driver" json:"driver,omitempty"`
	Memory               *Memory        `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_170f09d838bd8a04, []int{0}
}

func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetDriver() Storage_Driver {
	if m != nil {
		return m.Driver
	}
	return Storage_UNKNOWN
}

func (m *Storage) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterEnum("rpc.proto.Storage_Driver", Storage_Driver_name, Storage_Driver_value)
	proto.RegisterType((*Storage)(nil), "rpc.proto.Storage")
}

func init() { proto.RegisterFile("storage_message.proto", fileDescriptor_170f09d838bd8a04) }

var fileDescriptor_170f09d838bd8a04 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0x8d, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x2c, 0x2a, 0x48, 0x86, 0x30, 0xa5, 0x44, 0x72, 0x53, 0x73, 0xf3, 0x8b, 0x2a,
	0x51, 0x15, 0x28, 0x4d, 0x60, 0xe4, 0x62, 0x0f, 0x86, 0x68, 0x15, 0x32, 0xe4, 0x62, 0x4b, 0x29,
	0xca, 0x2c, 0x4b, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x92, 0xd4, 0x83, 0xeb, 0xd6,
	0x83, 0xaa, 0xd1, 0x73, 0x01, 0x2b, 0x08, 0x82, 0x2a, 0x14, 0xd2, 0xe4, 0x62, 0x83, 0x18, 0x2b,
	0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0x88, 0xa4, 0xc5, 0x17, 0x2c, 0x11, 0x04, 0x55, 0xa0,
	0xa4, 0xce, 0xc5, 0x06, 0xd1, 0x2c, 0xc4, 0xcd, 0xc5, 0x1e, 0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee,
	0x27, 0xc0, 0x20, 0xc4, 0xce, 0xc5, 0xec, 0xe1, 0xe2, 0x22, 0xc0, 0x08, 0x62, 0x04, 0x07, 0xbb,
	0x08, 0x30, 0x39, 0xa9, 0x70, 0xc9, 0x24, 0xe7, 0xe7, 0xea, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x05, 0x05, 0x38, 0xeb, 0xa5, 0xc3, 0x0d, 0x4d, 0x2a, 0x4d, 0x0b, 0x60, 0x8c, 0x62, 0x2a,
	0x48, 0x4a, 0x62, 0x03, 0xf3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0x8b, 0x8d, 0x90,
	0xf9, 0x00, 0x00, 0x00,
}
