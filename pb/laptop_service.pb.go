// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_service.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateLaptopRequest struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopRequest) Reset()         { *m = CreateLaptopRequest{} }
func (m *CreateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopRequest) ProtoMessage()    {}
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{0}
}

func (m *CreateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopRequest.Unmarshal(m, b)
}
func (m *CreateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *CreateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopRequest.Merge(m, src)
}
func (m *CreateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopRequest.Size(m)
}
func (m *CreateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopRequest proto.InternalMessageInfo

func (m *CreateLaptopRequest) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopResponse) Reset()         { *m = CreateLaptopResponse{} }
func (m *CreateLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopResponse) ProtoMessage()    {}
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{1}
}

func (m *CreateLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopResponse.Unmarshal(m, b)
}
func (m *CreateLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopResponse.Marshal(b, m, deterministic)
}
func (m *CreateLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopResponse.Merge(m, src)
}
func (m *CreateLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopResponse.Size(m)
}
func (m *CreateLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopResponse proto.InternalMessageInfo

func (m *CreateLaptopResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchLaptopRequest struct {
	Filter               *Filter  `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchLaptopRequest) Reset()         { *m = SearchLaptopRequest{} }
func (m *SearchLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*SearchLaptopRequest) ProtoMessage()    {}
func (*SearchLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{2}
}

func (m *SearchLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchLaptopRequest.Unmarshal(m, b)
}
func (m *SearchLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchLaptopRequest.Marshal(b, m, deterministic)
}
func (m *SearchLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchLaptopRequest.Merge(m, src)
}
func (m *SearchLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_SearchLaptopRequest.Size(m)
}
func (m *SearchLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchLaptopRequest proto.InternalMessageInfo

func (m *SearchLaptopRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type SearchLaptopResponse struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchLaptopResponse) Reset()         { *m = SearchLaptopResponse{} }
func (m *SearchLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*SearchLaptopResponse) ProtoMessage()    {}
func (*SearchLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{3}
}

func (m *SearchLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchLaptopResponse.Unmarshal(m, b)
}
func (m *SearchLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchLaptopResponse.Marshal(b, m, deterministic)
}
func (m *SearchLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchLaptopResponse.Merge(m, src)
}
func (m *SearchLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_SearchLaptopResponse.Size(m)
}
func (m *SearchLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchLaptopResponse proto.InternalMessageInfo

func (m *SearchLaptopResponse) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type UploadLaptopRequest struct {
	// Types that are valid to be assigned to Data:
	//	*UploadLaptopRequest_Info
	//	*UploadLaptopRequest_ChunkData
	Data                 isUploadLaptopRequest_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *UploadLaptopRequest) Reset()         { *m = UploadLaptopRequest{} }
func (m *UploadLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*UploadLaptopRequest) ProtoMessage()    {}
func (*UploadLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{4}
}

func (m *UploadLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadLaptopRequest.Unmarshal(m, b)
}
func (m *UploadLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadLaptopRequest.Marshal(b, m, deterministic)
}
func (m *UploadLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadLaptopRequest.Merge(m, src)
}
func (m *UploadLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_UploadLaptopRequest.Size(m)
}
func (m *UploadLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadLaptopRequest proto.InternalMessageInfo

type isUploadLaptopRequest_Data interface {
	isUploadLaptopRequest_Data()
}

type UploadLaptopRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadLaptopRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadLaptopRequest_Info) isUploadLaptopRequest_Data() {}

func (*UploadLaptopRequest_ChunkData) isUploadLaptopRequest_Data() {}

func (m *UploadLaptopRequest) GetData() isUploadLaptopRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UploadLaptopRequest) GetInfo() *ImageInfo {
	if x, ok := m.GetData().(*UploadLaptopRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (m *UploadLaptopRequest) GetChunkData() []byte {
	if x, ok := m.GetData().(*UploadLaptopRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UploadLaptopRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UploadLaptopRequest_Info)(nil),
		(*UploadLaptopRequest_ChunkData)(nil),
	}
}

type ImageInfo struct {
	LaptopId             string   `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	ImageType            string   `protobuf:"bytes,2,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageInfo) Reset()         { *m = ImageInfo{} }
func (m *ImageInfo) String() string { return proto.CompactTextString(m) }
func (*ImageInfo) ProtoMessage()    {}
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{5}
}

func (m *ImageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageInfo.Unmarshal(m, b)
}
func (m *ImageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageInfo.Marshal(b, m, deterministic)
}
func (m *ImageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageInfo.Merge(m, src)
}
func (m *ImageInfo) XXX_Size() int {
	return xxx_messageInfo_ImageInfo.Size(m)
}
func (m *ImageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ImageInfo proto.InternalMessageInfo

func (m *ImageInfo) GetLaptopId() string {
	if m != nil {
		return m.LaptopId
	}
	return ""
}

func (m *ImageInfo) GetImageType() string {
	if m != nil {
		return m.ImageType
	}
	return ""
}

type UploadLaptopResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size                 uint64   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadLaptopResponse) Reset()         { *m = UploadLaptopResponse{} }
func (m *UploadLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*UploadLaptopResponse) ProtoMessage()    {}
func (*UploadLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{6}
}

func (m *UploadLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadLaptopResponse.Unmarshal(m, b)
}
func (m *UploadLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadLaptopResponse.Marshal(b, m, deterministic)
}
func (m *UploadLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadLaptopResponse.Merge(m, src)
}
func (m *UploadLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_UploadLaptopResponse.Size(m)
}
func (m *UploadLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadLaptopResponse proto.InternalMessageInfo

func (m *UploadLaptopResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UploadLaptopResponse) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type RateLaptopRequest struct {
	LaptopId             string   `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	Score                float64  `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLaptopRequest) Reset()         { *m = RateLaptopRequest{} }
func (m *RateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*RateLaptopRequest) ProtoMessage()    {}
func (*RateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{7}
}

func (m *RateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLaptopRequest.Unmarshal(m, b)
}
func (m *RateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *RateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLaptopRequest.Merge(m, src)
}
func (m *RateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_RateLaptopRequest.Size(m)
}
func (m *RateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLaptopRequest proto.InternalMessageInfo

func (m *RateLaptopRequest) GetLaptopId() string {
	if m != nil {
		return m.LaptopId
	}
	return ""
}

func (m *RateLaptopRequest) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}

type RateLaptopResponse struct {
	LaptopId             string   `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	RateCount            uint32   `protobuf:"varint,2,opt,name=rate_count,json=rateCount,proto3" json:"rate_count,omitempty"`
	AverageScore         float64  `protobuf:"fixed64,3,opt,name=average_score,json=averageScore,proto3" json:"average_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLaptopResponse) Reset()         { *m = RateLaptopResponse{} }
func (m *RateLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*RateLaptopResponse) ProtoMessage()    {}
func (*RateLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{8}
}

func (m *RateLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLaptopResponse.Unmarshal(m, b)
}
func (m *RateLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLaptopResponse.Marshal(b, m, deterministic)
}
func (m *RateLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLaptopResponse.Merge(m, src)
}
func (m *RateLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_RateLaptopResponse.Size(m)
}
func (m *RateLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLaptopResponse proto.InternalMessageInfo

func (m *RateLaptopResponse) GetLaptopId() string {
	if m != nil {
		return m.LaptopId
	}
	return ""
}

func (m *RateLaptopResponse) GetRateCount() uint32 {
	if m != nil {
		return m.RateCount
	}
	return 0
}

func (m *RateLaptopResponse) GetAverageScore() float64 {
	if m != nil {
		return m.AverageScore
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateLaptopRequest)(nil), "rpc.proto.CreateLaptopRequest")
	proto.RegisterType((*CreateLaptopResponse)(nil), "rpc.proto.CreateLaptopResponse")
	proto.RegisterType((*SearchLaptopRequest)(nil), "rpc.proto.SearchLaptopRequest")
	proto.RegisterType((*SearchLaptopResponse)(nil), "rpc.proto.SearchLaptopResponse")
	proto.RegisterType((*UploadLaptopRequest)(nil), "rpc.proto.UploadLaptopRequest")
	proto.RegisterType((*ImageInfo)(nil), "rpc.proto.ImageInfo")
	proto.RegisterType((*UploadLaptopResponse)(nil), "rpc.proto.UploadLaptopResponse")
	proto.RegisterType((*RateLaptopRequest)(nil), "rpc.proto.RateLaptopRequest")
	proto.RegisterType((*RateLaptopResponse)(nil), "rpc.proto.RateLaptopResponse")
}

func init() { proto.RegisterFile("laptop_service.proto", fileDescriptor_240c60d9fb227e71) }

var fileDescriptor_240c60d9fb227e71 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe7, 0xae, 0x54, 0xe4, 0xd1, 0x82, 0xea, 0x46, 0x68, 0x94, 0x8e, 0x4d, 0x06, 0xa1,
	0xd2, 0x43, 0x3a, 0xc6, 0xad, 0x27, 0x58, 0xd1, 0x58, 0x25, 0x0e, 0x93, 0x0b, 0x17, 0x2e, 0x95,
	0x9b, 0xba, 0xa9, 0x21, 0x8d, 0x43, 0xec, 0x54, 0x1a, 0x47, 0xbe, 0x02, 0x5f, 0x8a, 0x3b, 0x5f,
	0x81, 0x0f, 0x82, 0x62, 0x67, 0x25, 0xa1, 0x19, 0x88, 0x5b, 0xf3, 0x7f, 0x7f, 0xff, 0x7f, 0xcf,
	0x7e, 0xaf, 0xe0, 0x86, 0x2c, 0xd6, 0x32, 0x9e, 0x29, 0x9e, 0x6c, 0x84, 0xcf, 0xbd, 0x38, 0x91,
	0x5a, 0x62, 0x27, 0x89, 0x7d, 0xfb, 0xb3, 0x7b, 0x6d, 0x58, 0x73, 0xa5, 0x58, 0xc0, 0xaf, 0xd5,
	0xa5, 0x08, 0x35, 0x4f, 0xfe, 0x50, 0x7b, 0x81, 0x94, 0x41, 0xc8, 0x87, 0x2c, 0x16, 0x43, 0x16,
	0x45, 0x52, 0x33, 0x2d, 0x64, 0xa4, 0x6c, 0x95, 0xbc, 0x84, 0xce, 0x38, 0xe1, 0x4c, 0xf3, 0xb7,
	0x26, 0x91, 0xf2, 0xcf, 0x29, 0x57, 0x1a, 0x3f, 0x83, 0x86, 0x45, 0x1c, 0xa0, 0x63, 0xd4, 0xbf,
	0x73, 0xda, 0xf6, 0xb6, 0x70, 0x2f, 0x77, 0xe6, 0x06, 0xf2, 0x14, 0xdc, 0x72, 0x82, 0x8a, 0x65,
	0xa4, 0x38, 0xbe, 0x0b, 0x35, 0xb1, 0x30, 0xc7, 0x1d, 0x5a, 0x13, 0x8b, 0x8c, 0x34, 0xe5, 0x2c,
	0xf1, 0x57, 0x3b, 0x24, 0xdb, 0x76, 0x05, 0xe9, 0xdc, 0x14, 0x68, 0x6e, 0x20, 0xaf, 0xc0, 0x2d,
	0x27, 0xe4, 0xa4, 0xff, 0x68, 0xf6, 0x23, 0x74, 0xde, 0xc7, 0xa1, 0x64, 0x8b, 0x72, 0x13, 0x03,
	0xa8, 0x8b, 0x68, 0x29, 0xf3, 0xf3, 0x6e, 0xe1, 0xfc, 0x64, 0xcd, 0x02, 0x3e, 0x89, 0x96, 0xf2,
	0x62, 0x8f, 0x1a, 0x0f, 0x3e, 0x02, 0xf0, 0x57, 0x69, 0xf4, 0x69, 0xb6, 0x60, 0x9a, 0x1d, 0xd4,
	0x8e, 0x51, 0xbf, 0x79, 0xb1, 0x47, 0x1d, 0xa3, 0xbd, 0x66, 0x9a, 0x9d, 0x35, 0xa0, 0x9e, 0x95,
	0xc8, 0x1b, 0x70, 0xb6, 0xa7, 0xf1, 0x43, 0x70, 0xf2, 0x99, 0x6d, 0x1f, 0xe5, 0xb6, 0x15, 0x26,
	0x0b, 0x7c, 0x08, 0x20, 0x32, 0xe7, 0x4c, 0x5f, 0xc5, 0xdc, 0x44, 0x3a, 0xd4, 0x31, 0xca, 0xbb,
	0xab, 0x98, 0x93, 0x11, 0xb8, 0xe5, 0xa6, 0xab, 0x5f, 0x18, 0x63, 0xa8, 0x2b, 0xf1, 0xc5, 0x06,
	0xd4, 0xa9, 0xf9, 0x4d, 0xce, 0xa1, 0x4d, 0x77, 0xa6, 0xfb, 0xd7, 0x66, 0x5c, 0xb8, 0xa5, 0x7c,
	0x99, 0xd8, 0x18, 0x44, 0xed, 0x07, 0x49, 0x01, 0xd3, 0xdd, 0x19, 0xff, 0xeb, 0x56, 0x09, 0xd3,
	0x7c, 0xe6, 0xcb, 0x34, 0xd2, 0x26, 0xad, 0x45, 0x9d, 0x4c, 0x19, 0x67, 0x02, 0x7e, 0x0c, 0x2d,
	0xb6, 0xe1, 0x49, 0x76, 0x6d, 0xcb, 0xdb, 0x37, 0xbc, 0x66, 0x2e, 0x4e, 0x33, 0xed, 0xf4, 0xfb,
	0x3e, 0xb4, 0x2c, 0x73, 0x6a, 0xff, 0x0b, 0x78, 0x0d, 0xcd, 0xe2, 0xba, 0xe1, 0x47, 0x85, 0x61,
	0x55, 0x6c, 0x72, 0xf7, 0xe8, 0xc6, 0xba, 0xbd, 0x03, 0xe9, 0x7d, 0xfd, 0xf1, 0xf3, 0x5b, 0xed,
	0x3e, 0x69, 0x0f, 0x37, 0xcf, 0x87, 0xb6, 0xf9, 0xa1, 0x6f, 0x8c, 0x23, 0x34, 0xc0, 0x21, 0x34,
	0x8b, 0x3b, 0x57, 0xc2, 0x55, 0xac, 0x73, 0x09, 0x57, 0xb5, 0xac, 0xe4, 0x81, 0xc1, 0x75, 0x70,
	0x11, 0xa7, 0x8c, 0xf1, 0x04, 0x61, 0x09, 0xcd, 0xe2, 0xa4, 0x4b, 0xb4, 0x8a, 0xbd, 0x2d, 0xd1,
	0xaa, 0x56, 0xa4, 0xf2, 0x72, 0xa9, 0x31, 0x8e, 0xd0, 0xa0, 0x8f, 0xb0, 0x00, 0xf8, 0x3d, 0x56,
	0xdc, 0x2b, 0xc4, 0xed, 0x6c, 0x4d, 0xf7, 0xf0, 0x86, 0x6a, 0x8e, 0xea, 0x1a, 0x94, 0x4b, 0xee,
	0x15, 0x50, 0x89, 0x7d, 0xc5, 0x3e, 0x3a, 0x41, 0x67, 0x4f, 0xa0, 0xe7, 0xcb, 0xb5, 0x17, 0x08,
	0xbd, 0x4a, 0xe7, 0x1e, 0xbd, 0x1c, 0x7b, 0xc1, 0x36, 0x6f, 0x9e, 0x2e, 0x2f, 0xd1, 0x87, 0x5a,
	0x3c, 0x9f, 0x37, 0xcc, 0xf7, 0x8b, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x62, 0x41, 0x6b, 0x2c,
	0x03, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LaptopServiceClient is the client API for LaptopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaptopServiceClient interface {
	//一元RPC 创建电脑
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
	//服务器流式RPC 检索电脑
	SearchLaptop(ctx context.Context, in *SearchLaptopRequest, opts ...grpc.CallOption) (LaptopService_SearchLaptopClient, error)
	//客户端流式RPC 上传图片
	UploadLaptop(ctx context.Context, opts ...grpc.CallOption) (LaptopService_UploadLaptopClient, error)
	//双向流式RPC 评分
	RateLaptop(ctx context.Context, opts ...grpc.CallOption) (LaptopService_RateLaptopClient, error)
}

type laptopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLaptopServiceClient(cc grpc.ClientConnInterface) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/rpc.proto.LaptopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laptopServiceClient) SearchLaptop(ctx context.Context, in *SearchLaptopRequest, opts ...grpc.CallOption) (LaptopService_SearchLaptopClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[0], "/rpc.proto.LaptopService/SearchLaptop", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceSearchLaptopClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LaptopService_SearchLaptopClient interface {
	Recv() (*SearchLaptopResponse, error)
	grpc.ClientStream
}

type laptopServiceSearchLaptopClient struct {
	grpc.ClientStream
}

func (x *laptopServiceSearchLaptopClient) Recv() (*SearchLaptopResponse, error) {
	m := new(SearchLaptopResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *laptopServiceClient) UploadLaptop(ctx context.Context, opts ...grpc.CallOption) (LaptopService_UploadLaptopClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[1], "/rpc.proto.LaptopService/UploadLaptop", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceUploadLaptopClient{stream}
	return x, nil
}

type LaptopService_UploadLaptopClient interface {
	Send(*UploadLaptopRequest) error
	CloseAndRecv() (*UploadLaptopResponse, error)
	grpc.ClientStream
}

type laptopServiceUploadLaptopClient struct {
	grpc.ClientStream
}

func (x *laptopServiceUploadLaptopClient) Send(m *UploadLaptopRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *laptopServiceUploadLaptopClient) CloseAndRecv() (*UploadLaptopResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadLaptopResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *laptopServiceClient) RateLaptop(ctx context.Context, opts ...grpc.CallOption) (LaptopService_RateLaptopClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[2], "/rpc.proto.LaptopService/RateLaptop", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceRateLaptopClient{stream}
	return x, nil
}

type LaptopService_RateLaptopClient interface {
	Send(*RateLaptopRequest) error
	Recv() (*RateLaptopResponse, error)
	grpc.ClientStream
}

type laptopServiceRateLaptopClient struct {
	grpc.ClientStream
}

func (x *laptopServiceRateLaptopClient) Send(m *RateLaptopRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *laptopServiceRateLaptopClient) Recv() (*RateLaptopResponse, error) {
	m := new(RateLaptopResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LaptopServiceServer is the server API for LaptopService service.
type LaptopServiceServer interface {
	//一元RPC 创建电脑
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
	//服务器流式RPC 检索电脑
	SearchLaptop(*SearchLaptopRequest, LaptopService_SearchLaptopServer) error
	//客户端流式RPC 上传图片
	UploadLaptop(LaptopService_UploadLaptopServer) error
	//双向流式RPC 评分
	RateLaptop(LaptopService_RateLaptopServer) error
}

// UnimplementedLaptopServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLaptopServiceServer struct {
}

func (*UnimplementedLaptopServiceServer) CreateLaptop(ctx context.Context, req *CreateLaptopRequest) (*CreateLaptopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaptop not implemented")
}
func (*UnimplementedLaptopServiceServer) SearchLaptop(req *SearchLaptopRequest, srv LaptopService_SearchLaptopServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchLaptop not implemented")
}
func (*UnimplementedLaptopServiceServer) UploadLaptop(srv LaptopService_UploadLaptopServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadLaptop not implemented")
}
func (*UnimplementedLaptopServiceServer) RateLaptop(srv LaptopService_RateLaptopServer) error {
	return status.Errorf(codes.Unimplemented, "method RateLaptop not implemented")
}

func RegisterLaptopServiceServer(s *grpc.Server, srv LaptopServiceServer) {
	s.RegisterService(&_LaptopService_serviceDesc, srv)
}

func _LaptopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.proto.LaptopService/CreateLaptop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaptopService_SearchLaptop_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchLaptopRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LaptopServiceServer).SearchLaptop(m, &laptopServiceSearchLaptopServer{stream})
}

type LaptopService_SearchLaptopServer interface {
	Send(*SearchLaptopResponse) error
	grpc.ServerStream
}

type laptopServiceSearchLaptopServer struct {
	grpc.ServerStream
}

func (x *laptopServiceSearchLaptopServer) Send(m *SearchLaptopResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LaptopService_UploadLaptop_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LaptopServiceServer).UploadLaptop(&laptopServiceUploadLaptopServer{stream})
}

type LaptopService_UploadLaptopServer interface {
	SendAndClose(*UploadLaptopResponse) error
	Recv() (*UploadLaptopRequest, error)
	grpc.ServerStream
}

type laptopServiceUploadLaptopServer struct {
	grpc.ServerStream
}

func (x *laptopServiceUploadLaptopServer) SendAndClose(m *UploadLaptopResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *laptopServiceUploadLaptopServer) Recv() (*UploadLaptopRequest, error) {
	m := new(UploadLaptopRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LaptopService_RateLaptop_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LaptopServiceServer).RateLaptop(&laptopServiceRateLaptopServer{stream})
}

type LaptopService_RateLaptopServer interface {
	Send(*RateLaptopResponse) error
	Recv() (*RateLaptopRequest, error)
	grpc.ServerStream
}

type laptopServiceRateLaptopServer struct {
	grpc.ServerStream
}

func (x *laptopServiceRateLaptopServer) Send(m *RateLaptopResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *laptopServiceRateLaptopServer) Recv() (*RateLaptopRequest, error) {
	m := new(RateLaptopRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LaptopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.proto.LaptopService",
	HandlerType: (*LaptopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LaptopService_CreateLaptop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchLaptop",
			Handler:       _LaptopService_SearchLaptop_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadLaptop",
			Handler:       _LaptopService_UploadLaptop_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RateLaptop",
			Handler:       _LaptopService_RateLaptop_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "laptop_service.proto",
}
