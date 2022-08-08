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
// source: google/ads/googleads/v11/resources/campaign_experiment.proto

package resources

import (
	enums "github.com/kritzware/google-ads-go/enums"
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

// An A/B experiment that compares the performance of the base campaign
// (the control) and a variation of that campaign (the experiment).
type CampaignExperiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the campaign experiment.
	// Campaign experiment resource names have the form:
	//
	// `customers/{customer_id}/campaignExperiments/{campaign_experiment_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the campaign experiment.
	//
	// This field is read-only.
	Id *int64 `protobuf:"varint,13,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Immutable. The campaign draft with staged changes to the base campaign.
	CampaignDraft *string `protobuf:"bytes,14,opt,name=campaign_draft,json=campaignDraft,proto3,oneof" json:"campaign_draft,omitempty"`
	// The name of the campaign experiment.
	//
	// This field is required when creating new campaign experiments
	// and must not conflict with the name of another non-removed
	// campaign experiment or campaign.
	//
	// It must not contain any null (code point 0x0), NL line feed
	// (code point 0xA) or carriage return (code point 0xD) characters.
	Name *string `protobuf:"bytes,15,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// The description of the experiment.
	Description *string `protobuf:"bytes,16,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Immutable. Share of traffic directed to experiment as a percent (must be between 1 and
	// 99 inclusive. Base campaign receives the remainder of the traffic
	// (100 - traffic_split_percent). Required for create.
	TrafficSplitPercent *int64 `protobuf:"varint,17,opt,name=traffic_split_percent,json=trafficSplitPercent,proto3,oneof" json:"traffic_split_percent,omitempty"`
	// Immutable. Determines the behavior of the traffic split.
	TrafficSplitType enums.CampaignExperimentTrafficSplitTypeEnum_CampaignExperimentTrafficSplitType `protobuf:"varint,7,opt,name=traffic_split_type,json=trafficSplitType,proto3,enum=google.ads.googleads.v11.enums.CampaignExperimentTrafficSplitTypeEnum_CampaignExperimentTrafficSplitType" json:"traffic_split_type,omitempty"`
	// Output only. The experiment campaign, as opposed to the base campaign.
	ExperimentCampaign *string `protobuf:"bytes,18,opt,name=experiment_campaign,json=experimentCampaign,proto3,oneof" json:"experiment_campaign,omitempty"`
	// Output only. The status of the campaign experiment. This field is read-only.
	Status enums.CampaignExperimentStatusEnum_CampaignExperimentStatus `protobuf:"varint,9,opt,name=status,proto3,enum=google.ads.googleads.v11.enums.CampaignExperimentStatusEnum_CampaignExperimentStatus" json:"status,omitempty"`
	// Output only. The resource name of the long-running operation that can be used to poll
	// for completion of experiment create or promote. The most recent long
	// running operation is returned.
	LongRunningOperation *string `protobuf:"bytes,19,opt,name=long_running_operation,json=longRunningOperation,proto3,oneof" json:"long_running_operation,omitempty"`
	// Date when the campaign experiment starts. By default, the experiment starts
	// now or on the campaign's start date, whichever is later. If this field is
	// set, then the experiment starts at the beginning of the specified date in
	// the customer's time zone. Cannot be changed once the experiment starts.
	//
	// Format: YYYY-MM-DD
	// Example: 2019-03-14
	StartDate *string `protobuf:"bytes,20,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	// The last day of the campaign experiment. By default, the experiment ends on
	// the campaign's end date. If this field is set, then the experiment ends at
	// the end of the specified date in the customer's time zone.
	//
	// Format: YYYY-MM-DD
	// Example: 2019-04-18
	EndDate *string `protobuf:"bytes,21,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
}

func (x *CampaignExperiment) Reset() {
	*x = CampaignExperiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_campaign_experiment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignExperiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignExperiment) ProtoMessage() {}

func (x *CampaignExperiment) ProtoReflect() protoreflect.Message {
	mi := &file_resources_campaign_experiment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignExperiment.ProtoReflect.Descriptor instead.
func (*CampaignExperiment) Descriptor() ([]byte, []int) {
	return file_resources_campaign_experiment_proto_rawDescGZIP(), []int{0}
}

func (x *CampaignExperiment) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CampaignExperiment) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *CampaignExperiment) GetCampaignDraft() string {
	if x != nil && x.CampaignDraft != nil {
		return *x.CampaignDraft
	}
	return ""
}

func (x *CampaignExperiment) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CampaignExperiment) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CampaignExperiment) GetTrafficSplitPercent() int64 {
	if x != nil && x.TrafficSplitPercent != nil {
		return *x.TrafficSplitPercent
	}
	return 0
}

func (x *CampaignExperiment) GetTrafficSplitType() enums.CampaignExperimentTrafficSplitTypeEnum_CampaignExperimentTrafficSplitType {
	if x != nil {
		return x.TrafficSplitType
	}
	return enums.CampaignExperimentTrafficSplitTypeEnum_CampaignExperimentTrafficSplitType(0)
}

func (x *CampaignExperiment) GetExperimentCampaign() string {
	if x != nil && x.ExperimentCampaign != nil {
		return *x.ExperimentCampaign
	}
	return ""
}

func (x *CampaignExperiment) GetStatus() enums.CampaignExperimentStatusEnum_CampaignExperimentStatus {
	if x != nil {
		return x.Status
	}
	return enums.CampaignExperimentStatusEnum_CampaignExperimentStatus(0)
}

func (x *CampaignExperiment) GetLongRunningOperation() string {
	if x != nil && x.LongRunningOperation != nil {
		return *x.LongRunningOperation
	}
	return ""
}

func (x *CampaignExperiment) GetStartDate() string {
	if x != nil && x.StartDate != nil {
		return *x.StartDate
	}
	return ""
}

func (x *CampaignExperiment) GetEndDate() string {
	if x != nil && x.EndDate != nil {
		return *x.EndDate
	}
	return ""
}

var File_resources_campaign_experiment_proto protoreflect.FileDescriptor

var file_resources_campaign_experiment_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f,
	0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x08, 0x0a,
	0x12, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x05, 0xfa,
	0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x5a, 0x0a, 0x0e, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2e, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x28, 0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x48,
	0x01, 0x52, 0x0d, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x73,
	0x70, 0x6c, 0x69, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x04, 0x52, 0x13, 0x74, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x9c, 0x01, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x69,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x53, 0x70, 0x6c, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x10,
	0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x5f, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0,
	0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x48, 0x05, 0x52, 0x12, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x72, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x55, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x16, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x72, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x06, 0x52, 0x14, 0x6c, 0x6f,
	0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x76, 0xea, 0x41, 0x73, 0x0a, 0x2b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x7d, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x16,
	0x0a, 0x14, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x5f,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x89, 0x02,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x17, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_resources_campaign_experiment_proto_rawDescOnce sync.Once
	file_resources_campaign_experiment_proto_rawDescData = file_resources_campaign_experiment_proto_rawDesc
)

func file_resources_campaign_experiment_proto_rawDescGZIP() []byte {
	file_resources_campaign_experiment_proto_rawDescOnce.Do(func() {
		file_resources_campaign_experiment_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_campaign_experiment_proto_rawDescData)
	})
	return file_resources_campaign_experiment_proto_rawDescData
}

var file_resources_campaign_experiment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_campaign_experiment_proto_goTypes = []interface{}{
	(*CampaignExperiment)(nil), // 0: google.ads.googleads.v11.resources.CampaignExperiment
	(enums.CampaignExperimentTrafficSplitTypeEnum_CampaignExperimentTrafficSplitType)(0), // 1: google.ads.googleads.v11.enums.CampaignExperimentTrafficSplitTypeEnum.CampaignExperimentTrafficSplitType
	(enums.CampaignExperimentStatusEnum_CampaignExperimentStatus)(0),                     // 2: google.ads.googleads.v11.enums.CampaignExperimentStatusEnum.CampaignExperimentStatus
}
var file_resources_campaign_experiment_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v11.resources.CampaignExperiment.traffic_split_type:type_name -> google.ads.googleads.v11.enums.CampaignExperimentTrafficSplitTypeEnum.CampaignExperimentTrafficSplitType
	2, // 1: google.ads.googleads.v11.resources.CampaignExperiment.status:type_name -> google.ads.googleads.v11.enums.CampaignExperimentStatusEnum.CampaignExperimentStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_resources_campaign_experiment_proto_init() }
func file_resources_campaign_experiment_proto_init() {
	if File_resources_campaign_experiment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_campaign_experiment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignExperiment); i {
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
	file_resources_campaign_experiment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_campaign_experiment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_campaign_experiment_proto_goTypes,
		DependencyIndexes: file_resources_campaign_experiment_proto_depIdxs,
		MessageInfos:      file_resources_campaign_experiment_proto_msgTypes,
	}.Build()
	File_resources_campaign_experiment_proto = out.File
	file_resources_campaign_experiment_proto_rawDesc = nil
	file_resources_campaign_experiment_proto_goTypes = nil
	file_resources_campaign_experiment_proto_depIdxs = nil
}
