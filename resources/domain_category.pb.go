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
// source: google/ads/googleads/v11/resources/domain_category.proto

package resources

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A category generated automatically by crawling a domain. If a campaign uses
// the DynamicSearchAdsSetting, then domain categories will be generated for
// the domain. The categories can be targeted using WebpageConditionInfo.
// See: https://support.google.com/google-ads/answer/2471185
type DomainCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the domain category.
	// Domain category resource names have the form:
	//
	// `customers/{customer_id}/domainCategories/{campaign_id}~{category_base64}~{language_code}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The campaign this category is recommended for.
	Campaign *string `protobuf:"bytes,10,opt,name=campaign,proto3,oneof" json:"campaign,omitempty"`
	// Output only. Recommended category for the website domain. e.g. if you have a website
	// about electronics, the categories could be "cameras", "televisions", etc.
	Category *string `protobuf:"bytes,11,opt,name=category,proto3,oneof" json:"category,omitempty"`
	// Output only. The language code specifying the language of the website. e.g. "en" for
	// English. The language can be specified in the DynamicSearchAdsSetting
	// required for dynamic search ads. This is the language of the pages from
	// your website that you want Google Ads to find, create ads for,
	// and match searches with.
	LanguageCode *string `protobuf:"bytes,12,opt,name=language_code,json=languageCode,proto3,oneof" json:"language_code,omitempty"`
	// Output only. The domain for the website. The domain can be specified in the
	// DynamicSearchAdsSetting required for dynamic search ads.
	Domain *string `protobuf:"bytes,13,opt,name=domain,proto3,oneof" json:"domain,omitempty"`
	// Output only. Fraction of pages on your site that this category matches.
	CoverageFraction *float64 `protobuf:"fixed64,14,opt,name=coverage_fraction,json=coverageFraction,proto3,oneof" json:"coverage_fraction,omitempty"`
	// Output only. The position of this category in the set of categories. Lower numbers
	// indicate a better match for the domain. null indicates not recommended.
	CategoryRank *int64 `protobuf:"varint,15,opt,name=category_rank,json=categoryRank,proto3,oneof" json:"category_rank,omitempty"`
	// Output only. Indicates whether this category has sub-categories.
	HasChildren *bool `protobuf:"varint,16,opt,name=has_children,json=hasChildren,proto3,oneof" json:"has_children,omitempty"`
	// Output only. The recommended cost per click for the category.
	RecommendedCpcBidMicros *int64 `protobuf:"varint,17,opt,name=recommended_cpc_bid_micros,json=recommendedCpcBidMicros,proto3,oneof" json:"recommended_cpc_bid_micros,omitempty"`
}

func (x *DomainCategory) Reset() {
	*x = DomainCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_domain_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainCategory) ProtoMessage() {}

func (x *DomainCategory) ProtoReflect() protoreflect.Message {
	mi := &file_resources_domain_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainCategory.ProtoReflect.Descriptor instead.
func (*DomainCategory) Descriptor() ([]byte, []int) {
	return file_resources_domain_category_proto_rawDescGZIP(), []int{0}
}

func (x *DomainCategory) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *DomainCategory) GetCampaign() string {
	if x != nil && x.Campaign != nil {
		return *x.Campaign
	}
	return ""
}

func (x *DomainCategory) GetCategory() string {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return ""
}

func (x *DomainCategory) GetLanguageCode() string {
	if x != nil && x.LanguageCode != nil {
		return *x.LanguageCode
	}
	return ""
}

func (x *DomainCategory) GetDomain() string {
	if x != nil && x.Domain != nil {
		return *x.Domain
	}
	return ""
}

func (x *DomainCategory) GetCoverageFraction() float64 {
	if x != nil && x.CoverageFraction != nil {
		return *x.CoverageFraction
	}
	return 0
}

func (x *DomainCategory) GetCategoryRank() int64 {
	if x != nil && x.CategoryRank != nil {
		return *x.CategoryRank
	}
	return 0
}

func (x *DomainCategory) GetHasChildren() bool {
	if x != nil && x.HasChildren != nil {
		return *x.HasChildren
	}
	return false
}

func (x *DomainCategory) GetRecommendedCpcBidMicros() int64 {
	if x != nil && x.RecommendedCpcBidMicros != nil {
		return *x.RecommendedCpcBidMicros
	}
	return 0
}

var File_resources_domain_category_proto protoreflect.FileDescriptor

var file_resources_domain_category_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x06, 0x0a, 0x0e, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x54, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x48, 0x00, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x24, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x02, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x11, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x01, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52, 0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a,
	0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x05, 0x52, 0x0c, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0c,
	0x68, 0x61, 0x73, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x06, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x1a, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x70, 0x63, 0x5f, 0x62, 0x69, 0x64,
	0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x48, 0x07, 0x52, 0x17, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x43, 0x70, 0x63, 0x42, 0x69, 0x64, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x88, 0x01, 0x01,
	0x3a, 0x87, 0x01, 0xea, 0x41, 0x83, 0x01, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x58, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x7d, 0x7e, 0x7b, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x7d, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x66,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x68, 0x61,
	0x73, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x70, 0x63, 0x5f, 0x62,
	0x69, 0x64, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x42, 0x85, 0x02, 0x0a, 0x26, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x42, 0x13, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_domain_category_proto_rawDescOnce sync.Once
	file_resources_domain_category_proto_rawDescData = file_resources_domain_category_proto_rawDesc
)

func file_resources_domain_category_proto_rawDescGZIP() []byte {
	file_resources_domain_category_proto_rawDescOnce.Do(func() {
		file_resources_domain_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_domain_category_proto_rawDescData)
	})
	return file_resources_domain_category_proto_rawDescData
}

var file_resources_domain_category_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_domain_category_proto_goTypes = []interface{}{
	(*DomainCategory)(nil), // 0: google.ads.googleads.v11.resources.DomainCategory
}
var file_resources_domain_category_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resources_domain_category_proto_init() }
func file_resources_domain_category_proto_init() {
	if File_resources_domain_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_domain_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainCategory); i {
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
	file_resources_domain_category_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_domain_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_domain_category_proto_goTypes,
		DependencyIndexes: file_resources_domain_category_proto_depIdxs,
		MessageInfos:      file_resources_domain_category_proto_msgTypes,
	}.Build()
	File_resources_domain_category_proto = out.File
	file_resources_domain_category_proto_rawDesc = nil
	file_resources_domain_category_proto_goTypes = nil
	file_resources_domain_category_proto_depIdxs = nil
}
