// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: google/ads/googleads/v11/enums/campaign_criterion_status.proto

package enums

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The possible statuses of a CampaignCriterion.
type CampaignCriterionStatusEnum_CampaignCriterionStatus int32

const (
	// No value has been specified.
	CampaignCriterionStatusEnum_UNSPECIFIED CampaignCriterionStatusEnum_CampaignCriterionStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	CampaignCriterionStatusEnum_UNKNOWN CampaignCriterionStatusEnum_CampaignCriterionStatus = 1
	// The campaign criterion is enabled.
	CampaignCriterionStatusEnum_ENABLED CampaignCriterionStatusEnum_CampaignCriterionStatus = 2
	// The campaign criterion is paused.
	CampaignCriterionStatusEnum_PAUSED CampaignCriterionStatusEnum_CampaignCriterionStatus = 3
	// The campaign criterion is removed.
	CampaignCriterionStatusEnum_REMOVED CampaignCriterionStatusEnum_CampaignCriterionStatus = 4
)

// Enum value maps for CampaignCriterionStatusEnum_CampaignCriterionStatus.
var (
	CampaignCriterionStatusEnum_CampaignCriterionStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ENABLED",
		3: "PAUSED",
		4: "REMOVED",
	}
	CampaignCriterionStatusEnum_CampaignCriterionStatus_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"ENABLED":     2,
		"PAUSED":      3,
		"REMOVED":     4,
	}
)

func (x CampaignCriterionStatusEnum_CampaignCriterionStatus) Enum() *CampaignCriterionStatusEnum_CampaignCriterionStatus {
	p := new(CampaignCriterionStatusEnum_CampaignCriterionStatus)
	*p = x
	return p
}

func (x CampaignCriterionStatusEnum_CampaignCriterionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignCriterionStatusEnum_CampaignCriterionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_campaign_criterion_status_proto_enumTypes[0].Descriptor()
}

func (CampaignCriterionStatusEnum_CampaignCriterionStatus) Type() protoreflect.EnumType {
	return &file_enums_campaign_criterion_status_proto_enumTypes[0]
}

func (x CampaignCriterionStatusEnum_CampaignCriterionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignCriterionStatusEnum_CampaignCriterionStatus.Descriptor instead.
func (CampaignCriterionStatusEnum_CampaignCriterionStatus) EnumDescriptor() ([]byte, []int) {
	return file_enums_campaign_criterion_status_proto_rawDescGZIP(), []int{0, 0}
}

// Message describing CampaignCriterion statuses.
type CampaignCriterionStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CampaignCriterionStatusEnum) Reset() {
	*x = CampaignCriterionStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_campaign_criterion_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignCriterionStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignCriterionStatusEnum) ProtoMessage() {}

func (x *CampaignCriterionStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_campaign_criterion_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignCriterionStatusEnum.ProtoReflect.Descriptor instead.
func (*CampaignCriterionStatusEnum) Descriptor() ([]byte, []int) {
	return file_enums_campaign_criterion_status_proto_rawDescGZIP(), []int{0}
}

var File_enums_campaign_criterion_status_proto protoreflect.FileDescriptor

var file_enums_campaign_criterion_status_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x22, 0x7c, 0x0a, 0x1b, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22,
	0x5d, 0x0a, 0x17, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x04, 0x42, 0xf6,
	0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31,
	0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_campaign_criterion_status_proto_rawDescOnce sync.Once
	file_enums_campaign_criterion_status_proto_rawDescData = file_enums_campaign_criterion_status_proto_rawDesc
)

func file_enums_campaign_criterion_status_proto_rawDescGZIP() []byte {
	file_enums_campaign_criterion_status_proto_rawDescOnce.Do(func() {
		file_enums_campaign_criterion_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_campaign_criterion_status_proto_rawDescData)
	})
	return file_enums_campaign_criterion_status_proto_rawDescData
}

var file_enums_campaign_criterion_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_campaign_criterion_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_campaign_criterion_status_proto_goTypes = []interface{}{
	(CampaignCriterionStatusEnum_CampaignCriterionStatus)(0), // 0: google.ads.googleads.v11.enums.CampaignCriterionStatusEnum.CampaignCriterionStatus
	(*CampaignCriterionStatusEnum)(nil),                      // 1: google.ads.googleads.v11.enums.CampaignCriterionStatusEnum
}
var file_enums_campaign_criterion_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_campaign_criterion_status_proto_init() }
func file_enums_campaign_criterion_status_proto_init() {
	if File_enums_campaign_criterion_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_campaign_criterion_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignCriterionStatusEnum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_campaign_criterion_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_campaign_criterion_status_proto_goTypes,
		DependencyIndexes: file_enums_campaign_criterion_status_proto_depIdxs,
		EnumInfos:         file_enums_campaign_criterion_status_proto_enumTypes,
		MessageInfos:      file_enums_campaign_criterion_status_proto_msgTypes,
	}.Build()
	File_enums_campaign_criterion_status_proto = out.File
	file_enums_campaign_criterion_status_proto_rawDesc = nil
	file_enums_campaign_criterion_status_proto_goTypes = nil
	file_enums_campaign_criterion_status_proto_depIdxs = nil
}
