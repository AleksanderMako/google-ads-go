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
// source: google/ads/googleads/v11/enums/user_list_string_rule_item_operator.proto

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

// Enum describing possible user list string rule item operators.
type UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator int32

const (
	// Not specified.
	UserListStringRuleItemOperatorEnum_UNSPECIFIED UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 0
	// Used for return value only. Represents value unknown in this version.
	UserListStringRuleItemOperatorEnum_UNKNOWN UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 1
	// Contains.
	UserListStringRuleItemOperatorEnum_CONTAINS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 2
	// Equals.
	UserListStringRuleItemOperatorEnum_EQUALS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 3
	// Starts with.
	UserListStringRuleItemOperatorEnum_STARTS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 4
	// Ends with.
	UserListStringRuleItemOperatorEnum_ENDS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 5
	// Not equals.
	UserListStringRuleItemOperatorEnum_NOT_EQUALS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 6
	// Not contains.
	UserListStringRuleItemOperatorEnum_NOT_CONTAINS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 7
	// Not starts with.
	UserListStringRuleItemOperatorEnum_NOT_STARTS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 8
	// Not ends with.
	UserListStringRuleItemOperatorEnum_NOT_ENDS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 9
)

// Enum value maps for UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator.
var (
	UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "CONTAINS",
		3: "EQUALS",
		4: "STARTS_WITH",
		5: "ENDS_WITH",
		6: "NOT_EQUALS",
		7: "NOT_CONTAINS",
		8: "NOT_STARTS_WITH",
		9: "NOT_ENDS_WITH",
	}
	UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_value = map[string]int32{
		"UNSPECIFIED":     0,
		"UNKNOWN":         1,
		"CONTAINS":        2,
		"EQUALS":          3,
		"STARTS_WITH":     4,
		"ENDS_WITH":       5,
		"NOT_EQUALS":      6,
		"NOT_CONTAINS":    7,
		"NOT_STARTS_WITH": 8,
		"NOT_ENDS_WITH":   9,
	}
)

func (x UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator) Enum() *UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator {
	p := new(UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator)
	*p = x
	return p
}

func (x UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_user_list_string_rule_item_operator_proto_enumTypes[0].Descriptor()
}

func (UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator) Type() protoreflect.EnumType {
	return &file_enums_user_list_string_rule_item_operator_proto_enumTypes[0]
}

func (x UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator.Descriptor instead.
func (UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator) EnumDescriptor() ([]byte, []int) {
	return file_enums_user_list_string_rule_item_operator_proto_rawDescGZIP(), []int{0, 0}
}

// Supported rule operator for string type.
type UserListStringRuleItemOperatorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserListStringRuleItemOperatorEnum) Reset() {
	*x = UserListStringRuleItemOperatorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_user_list_string_rule_item_operator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListStringRuleItemOperatorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListStringRuleItemOperatorEnum) ProtoMessage() {}

func (x *UserListStringRuleItemOperatorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_user_list_string_rule_item_operator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListStringRuleItemOperatorEnum.ProtoReflect.Descriptor instead.
func (*UserListStringRuleItemOperatorEnum) Descriptor() ([]byte, []int) {
	return file_enums_user_list_string_rule_item_operator_proto_rawDescGZIP(), []int{0}
}

var File_enums_user_list_string_rule_item_operator_proto protoreflect.FileDescriptor

var file_enums_user_list_string_rule_item_operator_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xe9, 0x01, 0x0a, 0x22, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0xc2, 0x01, 0x0a, 0x1e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x04, 0x12, 0x0d, 0x0a,
	0x09, 0x45, 0x4e, 0x44, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a,
	0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c,
	0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x07, 0x12, 0x13,
	0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x53, 0x5f, 0x57, 0x49, 0x54,
	0x48, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x53, 0x5f,
	0x57, 0x49, 0x54, 0x48, 0x10, 0x09, 0x42, 0xfd, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x23, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_user_list_string_rule_item_operator_proto_rawDescOnce sync.Once
	file_enums_user_list_string_rule_item_operator_proto_rawDescData = file_enums_user_list_string_rule_item_operator_proto_rawDesc
)

func file_enums_user_list_string_rule_item_operator_proto_rawDescGZIP() []byte {
	file_enums_user_list_string_rule_item_operator_proto_rawDescOnce.Do(func() {
		file_enums_user_list_string_rule_item_operator_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_user_list_string_rule_item_operator_proto_rawDescData)
	})
	return file_enums_user_list_string_rule_item_operator_proto_rawDescData
}

var file_enums_user_list_string_rule_item_operator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_user_list_string_rule_item_operator_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_user_list_string_rule_item_operator_proto_goTypes = []interface{}{
	(UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator)(0), // 0: google.ads.googleads.v11.enums.UserListStringRuleItemOperatorEnum.UserListStringRuleItemOperator
	(*UserListStringRuleItemOperatorEnum)(nil),                             // 1: google.ads.googleads.v11.enums.UserListStringRuleItemOperatorEnum
}
var file_enums_user_list_string_rule_item_operator_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_user_list_string_rule_item_operator_proto_init() }
func file_enums_user_list_string_rule_item_operator_proto_init() {
	if File_enums_user_list_string_rule_item_operator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_user_list_string_rule_item_operator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListStringRuleItemOperatorEnum); i {
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
			RawDescriptor: file_enums_user_list_string_rule_item_operator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_user_list_string_rule_item_operator_proto_goTypes,
		DependencyIndexes: file_enums_user_list_string_rule_item_operator_proto_depIdxs,
		EnumInfos:         file_enums_user_list_string_rule_item_operator_proto_enumTypes,
		MessageInfos:      file_enums_user_list_string_rule_item_operator_proto_msgTypes,
	}.Build()
	File_enums_user_list_string_rule_item_operator_proto = out.File
	file_enums_user_list_string_rule_item_operator_proto_rawDesc = nil
	file_enums_user_list_string_rule_item_operator_proto_goTypes = nil
	file_enums_user_list_string_rule_item_operator_proto_depIdxs = nil
}
