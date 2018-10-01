// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/enums/campaign_serving_status.proto

package enums

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible serving statuses of a campaign.
type CampaignServingStatusEnum_CampaignServingStatus int32

const (
	// No value has been specified.
	CampaignServingStatusEnum_UNSPECIFIED CampaignServingStatusEnum_CampaignServingStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	CampaignServingStatusEnum_UNKNOWN CampaignServingStatusEnum_CampaignServingStatus = 1
	// Serving.
	CampaignServingStatusEnum_SERVING CampaignServingStatusEnum_CampaignServingStatus = 2
	// None.
	CampaignServingStatusEnum_NONE CampaignServingStatusEnum_CampaignServingStatus = 3
	// Ended.
	CampaignServingStatusEnum_ENDED CampaignServingStatusEnum_CampaignServingStatus = 4
	// Pending.
	CampaignServingStatusEnum_PENDING CampaignServingStatusEnum_CampaignServingStatus = 5
	// Suspended.
	CampaignServingStatusEnum_SUSPENDED CampaignServingStatusEnum_CampaignServingStatus = 6
)

var CampaignServingStatusEnum_CampaignServingStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SERVING",
	3: "NONE",
	4: "ENDED",
	5: "PENDING",
	6: "SUSPENDED",
}

var CampaignServingStatusEnum_CampaignServingStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"SERVING":     2,
	"NONE":        3,
	"ENDED":       4,
	"PENDING":     5,
	"SUSPENDED":   6,
}

func (x CampaignServingStatusEnum_CampaignServingStatus) String() string {
	return proto.EnumName(CampaignServingStatusEnum_CampaignServingStatus_name, int32(x))
}

func (CampaignServingStatusEnum_CampaignServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_782a10ac4e85e3ac, []int{0, 0}
}

// Message describing Campaign serving statuses.
type CampaignServingStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignServingStatusEnum) Reset()         { *m = CampaignServingStatusEnum{} }
func (m *CampaignServingStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignServingStatusEnum) ProtoMessage()    {}
func (*CampaignServingStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_782a10ac4e85e3ac, []int{0}
}
func (m *CampaignServingStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignServingStatusEnum.Unmarshal(m, b)
}
func (m *CampaignServingStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignServingStatusEnum.Marshal(b, m, deterministic)
}
func (m *CampaignServingStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignServingStatusEnum.Merge(m, src)
}
func (m *CampaignServingStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignServingStatusEnum.Size(m)
}
func (m *CampaignServingStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignServingStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignServingStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CampaignServingStatusEnum)(nil), "google.ads.googleads.v0.enums.CampaignServingStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CampaignServingStatusEnum_CampaignServingStatus", CampaignServingStatusEnum_CampaignServingStatus_name, CampaignServingStatusEnum_CampaignServingStatus_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/enums/campaign_serving_status.proto", fileDescriptor_782a10ac4e85e3ac)
}

var fileDescriptor_782a10ac4e85e3ac = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xe4, 0xc4, 0xdc, 0x82, 0xc4, 0xcc, 0xf4, 0xbc, 0xf8, 0xe2, 0xd4, 0xa2,
	0xb2, 0xcc, 0xbc, 0xf4, 0xf8, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0xbd, 0x82, 0xa2, 0xfc, 0x92,
	0x7c, 0x21, 0x59, 0x88, 0x0e, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x66, 0xbd, 0x32, 0x03, 0x3d,
	0xb0, 0x66, 0xa5, 0x09, 0x8c, 0x5c, 0x92, 0xce, 0x50, 0x03, 0x82, 0x21, 0xfa, 0x83, 0xc1, 0xda,
	0x5d, 0xf3, 0x4a, 0x73, 0x95, 0x8a, 0xb9, 0x44, 0xb1, 0x4a, 0x0a, 0xf1, 0x73, 0x71, 0x87, 0xfa,
	0x05, 0x07, 0xb8, 0x3a, 0x7b, 0xba, 0x79, 0xba, 0xba, 0x08, 0x30, 0x08, 0x71, 0x73, 0xb1, 0x87,
	0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0x82, 0x38, 0xc1, 0xae, 0x41, 0x61, 0x9e, 0x7e,
	0xee, 0x02, 0x4c, 0x42, 0x1c, 0x5c, 0x2c, 0x7e, 0xfe, 0x7e, 0xae, 0x02, 0xcc, 0x42, 0x9c, 0x5c,
	0xac, 0xae, 0x7e, 0x2e, 0xae, 0x2e, 0x02, 0x2c, 0x20, 0x15, 0x01, 0xae, 0x7e, 0x2e, 0x20, 0x15,
	0xac, 0x42, 0xbc, 0x5c, 0x9c, 0xc1, 0xa1, 0xc1, 0x01, 0x10, 0x39, 0x36, 0xa7, 0xe5, 0x8c, 0x5c,
	0x8a, 0xc9, 0xf9, 0xb9, 0x7a, 0x78, 0x1d, 0xee, 0x24, 0x85, 0xd5, 0x61, 0x01, 0x20, 0x3f, 0x07,
	0x30, 0x2e, 0x62, 0x62, 0x76, 0x77, 0x74, 0x5c, 0xc5, 0x24, 0xeb, 0x0e, 0x31, 0xc3, 0x31, 0xa5,
	0x58, 0x0f, 0xc2, 0x04, 0xb1, 0xc2, 0x0c, 0xf4, 0x40, 0xbe, 0x2b, 0x3e, 0x05, 0x93, 0x8f, 0x71,
	0x4c, 0x29, 0x8e, 0x81, 0xcb, 0xc7, 0x84, 0x19, 0xc4, 0x80, 0xe5, 0x1f, 0x11, 0x90, 0x4f, 0x62,
	0x03, 0x07, 0xb1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x5a, 0x96, 0x3a, 0xa1, 0x01, 0x00,
	0x00,
}