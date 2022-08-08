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
// source: google/ads/googleads/v11/enums/seasonality_event_scope.proto

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

// The possible scopes of a Seasonality Event.
type SeasonalityEventScopeEnum_SeasonalityEventScope int32

const (
	// No value has been specified.
	SeasonalityEventScopeEnum_UNSPECIFIED SeasonalityEventScopeEnum_SeasonalityEventScope = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	SeasonalityEventScopeEnum_UNKNOWN SeasonalityEventScopeEnum_SeasonalityEventScope = 1
	// The seasonality event is applied to all the customer's traffic for
	// supported advertising channel types and device types. The CUSTOMER scope
	// cannot be used in mutates.
	SeasonalityEventScopeEnum_CUSTOMER SeasonalityEventScopeEnum_SeasonalityEventScope = 2
	// The seasonality event is applied to all specified campaigns.
	SeasonalityEventScopeEnum_CAMPAIGN SeasonalityEventScopeEnum_SeasonalityEventScope = 4
	// The seasonality event is applied to all campaigns that belong to
	// specified channel types.
	SeasonalityEventScopeEnum_CHANNEL SeasonalityEventScopeEnum_SeasonalityEventScope = 5
)

// Enum value maps for SeasonalityEventScopeEnum_SeasonalityEventScope.
var (
	SeasonalityEventScopeEnum_SeasonalityEventScope_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "CUSTOMER",
		4: "CAMPAIGN",
		5: "CHANNEL",
	}
	SeasonalityEventScopeEnum_SeasonalityEventScope_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"CUSTOMER":    2,
		"CAMPAIGN":    4,
		"CHANNEL":     5,
	}
)

func (x SeasonalityEventScopeEnum_SeasonalityEventScope) Enum() *SeasonalityEventScopeEnum_SeasonalityEventScope {
	p := new(SeasonalityEventScopeEnum_SeasonalityEventScope)
	*p = x
	return p
}

func (x SeasonalityEventScopeEnum_SeasonalityEventScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SeasonalityEventScopeEnum_SeasonalityEventScope) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_seasonality_event_scope_proto_enumTypes[0].Descriptor()
}

func (SeasonalityEventScopeEnum_SeasonalityEventScope) Type() protoreflect.EnumType {
	return &file_enums_seasonality_event_scope_proto_enumTypes[0]
}

func (x SeasonalityEventScopeEnum_SeasonalityEventScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SeasonalityEventScopeEnum_SeasonalityEventScope.Descriptor instead.
func (SeasonalityEventScopeEnum_SeasonalityEventScope) EnumDescriptor() ([]byte, []int) {
	return file_enums_seasonality_event_scope_proto_rawDescGZIP(), []int{0, 0}
}

// Message describing seasonality event scopes. The two types of seasonality
// events are BiddingSeasonalityAdjustments and BiddingDataExclusions.
type SeasonalityEventScopeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SeasonalityEventScopeEnum) Reset() {
	*x = SeasonalityEventScopeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_seasonality_event_scope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeasonalityEventScopeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeasonalityEventScopeEnum) ProtoMessage() {}

func (x *SeasonalityEventScopeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_seasonality_event_scope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeasonalityEventScopeEnum.ProtoReflect.Descriptor instead.
func (*SeasonalityEventScopeEnum) Descriptor() ([]byte, []int) {
	return file_enums_seasonality_event_scope_proto_rawDescGZIP(), []int{0}
}

var File_enums_seasonality_event_scope_proto protoreflect.FileDescriptor

var file_enums_seasonality_event_scope_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x7b,
	0x0a, 0x19, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x5e, 0x0a, 0x15, 0x53,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x10, 0x05, 0x42, 0xf4, 0x01, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x42, 0x1a, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_seasonality_event_scope_proto_rawDescOnce sync.Once
	file_enums_seasonality_event_scope_proto_rawDescData = file_enums_seasonality_event_scope_proto_rawDesc
)

func file_enums_seasonality_event_scope_proto_rawDescGZIP() []byte {
	file_enums_seasonality_event_scope_proto_rawDescOnce.Do(func() {
		file_enums_seasonality_event_scope_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_seasonality_event_scope_proto_rawDescData)
	})
	return file_enums_seasonality_event_scope_proto_rawDescData
}

var file_enums_seasonality_event_scope_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_seasonality_event_scope_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_seasonality_event_scope_proto_goTypes = []interface{}{
	(SeasonalityEventScopeEnum_SeasonalityEventScope)(0), // 0: google.ads.googleads.v11.enums.SeasonalityEventScopeEnum.SeasonalityEventScope
	(*SeasonalityEventScopeEnum)(nil),                    // 1: google.ads.googleads.v11.enums.SeasonalityEventScopeEnum
}
var file_enums_seasonality_event_scope_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_seasonality_event_scope_proto_init() }
func file_enums_seasonality_event_scope_proto_init() {
	if File_enums_seasonality_event_scope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_seasonality_event_scope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeasonalityEventScopeEnum); i {
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
			RawDescriptor: file_enums_seasonality_event_scope_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_seasonality_event_scope_proto_goTypes,
		DependencyIndexes: file_enums_seasonality_event_scope_proto_depIdxs,
		EnumInfos:         file_enums_seasonality_event_scope_proto_enumTypes,
		MessageInfos:      file_enums_seasonality_event_scope_proto_msgTypes,
	}.Build()
	File_enums_seasonality_event_scope_proto = out.File
	file_enums_seasonality_event_scope_proto_rawDesc = nil
	file_enums_seasonality_event_scope_proto_goTypes = nil
	file_enums_seasonality_event_scope_proto_depIdxs = nil
}
