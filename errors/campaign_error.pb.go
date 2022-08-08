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
// source: google/ads/googleads/v11/errors/campaign_error.proto

package errors

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

// Enum describing possible campaign errors.
type CampaignErrorEnum_CampaignError int32

const (
	// Enum unspecified.
	CampaignErrorEnum_UNSPECIFIED CampaignErrorEnum_CampaignError = 0
	// The received error code is not known in this version.
	CampaignErrorEnum_UNKNOWN CampaignErrorEnum_CampaignError = 1
	// Cannot target content network.
	CampaignErrorEnum_CANNOT_TARGET_CONTENT_NETWORK CampaignErrorEnum_CampaignError = 3
	// Cannot target search network.
	CampaignErrorEnum_CANNOT_TARGET_SEARCH_NETWORK CampaignErrorEnum_CampaignError = 4
	// Cannot cover search network without google search network.
	CampaignErrorEnum_CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH CampaignErrorEnum_CampaignError = 5
	// Cannot target Google Search network for a CPM campaign.
	CampaignErrorEnum_CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN CampaignErrorEnum_CampaignError = 6
	// Must target at least one network.
	CampaignErrorEnum_CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK CampaignErrorEnum_CampaignError = 7
	// Only some Google partners are allowed to target partner search network.
	CampaignErrorEnum_CANNOT_TARGET_PARTNER_SEARCH_NETWORK CampaignErrorEnum_CampaignError = 8
	// Cannot target content network only as campaign has criteria-level bidding
	// strategy.
	CampaignErrorEnum_CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY CampaignErrorEnum_CampaignError = 9
	// Cannot modify the start or end date such that the campaign duration would
	// not contain the durations of all runnable trials.
	CampaignErrorEnum_CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS CampaignErrorEnum_CampaignError = 10
	// Cannot modify dates, budget or status of a trial campaign.
	CampaignErrorEnum_CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN CampaignErrorEnum_CampaignError = 11
	// Trying to modify the name of an active or paused campaign, where the name
	// is already assigned to another active or paused campaign.
	CampaignErrorEnum_DUPLICATE_CAMPAIGN_NAME CampaignErrorEnum_CampaignError = 12
	// Two fields are in conflicting modes.
	CampaignErrorEnum_INCOMPATIBLE_CAMPAIGN_FIELD CampaignErrorEnum_CampaignError = 13
	// Campaign name cannot be used.
	CampaignErrorEnum_INVALID_CAMPAIGN_NAME CampaignErrorEnum_CampaignError = 14
	// Given status is invalid.
	CampaignErrorEnum_INVALID_AD_SERVING_OPTIMIZATION_STATUS CampaignErrorEnum_CampaignError = 15
	// Error in the campaign level tracking URL.
	CampaignErrorEnum_INVALID_TRACKING_URL CampaignErrorEnum_CampaignError = 16
	// Cannot set both tracking URL template and tracking setting. A user has
	// to clear legacy tracking setting in order to add tracking URL template.
	CampaignErrorEnum_CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING CampaignErrorEnum_CampaignError = 17
	// The maximum number of impressions for Frequency Cap should be an integer
	// greater than 0.
	CampaignErrorEnum_MAX_IMPRESSIONS_NOT_IN_RANGE CampaignErrorEnum_CampaignError = 18
	// Only the Day, Week and Month time units are supported.
	CampaignErrorEnum_TIME_UNIT_NOT_SUPPORTED CampaignErrorEnum_CampaignError = 19
	// Operation not allowed on a campaign whose serving status has ended
	CampaignErrorEnum_INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED CampaignErrorEnum_CampaignError = 20
	// This budget is exclusively linked to a Campaign that is using experiments
	// so it cannot be shared.
	CampaignErrorEnum_BUDGET_CANNOT_BE_SHARED CampaignErrorEnum_CampaignError = 21
	// Campaigns using experiments cannot use a shared budget.
	CampaignErrorEnum_CAMPAIGN_CANNOT_USE_SHARED_BUDGET CampaignErrorEnum_CampaignError = 22
	// A different budget cannot be assigned to a campaign when there are
	// running or scheduled trials.
	CampaignErrorEnum_CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS CampaignErrorEnum_CampaignError = 23
	// No link found between the campaign and the label.
	CampaignErrorEnum_CAMPAIGN_LABEL_DOES_NOT_EXIST CampaignErrorEnum_CampaignError = 24
	// The label has already been attached to the campaign.
	CampaignErrorEnum_CAMPAIGN_LABEL_ALREADY_EXISTS CampaignErrorEnum_CampaignError = 25
	// A ShoppingSetting was not found when creating a shopping campaign.
	CampaignErrorEnum_MISSING_SHOPPING_SETTING CampaignErrorEnum_CampaignError = 26
	// The country in shopping setting is not an allowed country.
	CampaignErrorEnum_INVALID_SHOPPING_SALES_COUNTRY CampaignErrorEnum_CampaignError = 27
	// The requested channel type is not available according to the customer's
	// account setting.
	CampaignErrorEnum_ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE CampaignErrorEnum_CampaignError = 31
	// The AdvertisingChannelSubType is not a valid subtype of the primary
	// channel type.
	CampaignErrorEnum_INVALID_ADVERTISING_CHANNEL_SUB_TYPE CampaignErrorEnum_CampaignError = 32
	// At least one conversion must be selected.
	CampaignErrorEnum_AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED CampaignErrorEnum_CampaignError = 33
	// Setting ad rotation mode for a campaign is not allowed. Ad rotation mode
	// at campaign is deprecated.
	CampaignErrorEnum_CANNOT_SET_AD_ROTATION_MODE CampaignErrorEnum_CampaignError = 34
	// Trying to change start date on a campaign that has started.
	CampaignErrorEnum_CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED CampaignErrorEnum_CampaignError = 35
	// Trying to modify a date into the past.
	CampaignErrorEnum_CANNOT_SET_DATE_TO_PAST CampaignErrorEnum_CampaignError = 36
	// Hotel center id in the hotel setting does not match any customer links.
	CampaignErrorEnum_MISSING_HOTEL_CUSTOMER_LINK CampaignErrorEnum_CampaignError = 37
	// Hotel center id in the hotel setting must match an active customer link.
	CampaignErrorEnum_INVALID_HOTEL_CUSTOMER_LINK CampaignErrorEnum_CampaignError = 38
	// Hotel setting was not found when creating a hotel ads campaign.
	CampaignErrorEnum_MISSING_HOTEL_SETTING CampaignErrorEnum_CampaignError = 39
	// A Campaign cannot use shared campaign budgets and be part of a campaign
	// group.
	CampaignErrorEnum_CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP CampaignErrorEnum_CampaignError = 40
	// The app ID was not found.
	CampaignErrorEnum_APP_NOT_FOUND CampaignErrorEnum_CampaignError = 41
	// Campaign.shopping_setting.enable_local is not supported for the specified
	// campaign type.
	CampaignErrorEnum_SHOPPING_ENABLE_LOCAL_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE CampaignErrorEnum_CampaignError = 42
	// The merchant does not support the creation of campaigns for Shopping
	// Comparison Listing Ads.
	CampaignErrorEnum_MERCHANT_NOT_ALLOWED_FOR_COMPARISON_LISTING_ADS CampaignErrorEnum_CampaignError = 43
	// The App campaign for engagement cannot be created because there aren't
	// enough installs.
	CampaignErrorEnum_INSUFFICIENT_APP_INSTALLS_COUNT CampaignErrorEnum_CampaignError = 44
	// The App campaign for engagement cannot be created because the app is
	// sensitive.
	CampaignErrorEnum_SENSITIVE_CATEGORY_APP CampaignErrorEnum_CampaignError = 45
	// Customers with Housing, Employment, or Credit ads must accept updated
	// personalized ads policy to continue creating campaigns.
	CampaignErrorEnum_HEC_AGREEMENT_REQUIRED CampaignErrorEnum_CampaignError = 46
	// The field is not compatible with view through conversion optimization.
	CampaignErrorEnum_NOT_COMPATIBLE_WITH_VIEW_THROUGH_CONVERSION_OPTIMIZATION CampaignErrorEnum_CampaignError = 49
	// The field type cannot be excluded because an active campaign-asset link
	// of this type exists.
	CampaignErrorEnum_INVALID_EXCLUDED_PARENT_ASSET_FIELD_TYPE CampaignErrorEnum_CampaignError = 48
	// The app pre-registration campaign cannot be created for non-Android
	// applications.
	CampaignErrorEnum_CANNOT_CREATE_APP_PRE_REGISTRATION_FOR_NON_ANDROID_APP CampaignErrorEnum_CampaignError = 50
	// The campaign cannot be created since the app is not available for
	// pre-registration in any country.
	CampaignErrorEnum_APP_NOT_AVAILABLE_TO_CREATE_APP_PRE_REGISTRATION_CAMPAIGN CampaignErrorEnum_CampaignError = 51
	// The type of the Budget is not compatible with this Campaign.
	CampaignErrorEnum_INCOMPATIBLE_BUDGET_TYPE CampaignErrorEnum_CampaignError = 52
	// Category bid list in the local services campaign setting contains
	// multiple bids for the same category ID.
	CampaignErrorEnum_LOCAL_SERVICES_DUPLICATE_CATEGORY_BID CampaignErrorEnum_CampaignError = 53
	// Category bid list in the local services campaign setting contains
	// a bid for an invalid category ID.
	CampaignErrorEnum_LOCAL_SERVICES_INVALID_CATEGORY_BID CampaignErrorEnum_CampaignError = 54
	// Category bid list in the local services campaign setting is missing a
	// bid for a category ID that must be present.
	CampaignErrorEnum_LOCAL_SERVICES_MISSING_CATEGORY_BID CampaignErrorEnum_CampaignError = 55
)

// Enum value maps for CampaignErrorEnum_CampaignError.
var (
	CampaignErrorEnum_CampaignError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		3:  "CANNOT_TARGET_CONTENT_NETWORK",
		4:  "CANNOT_TARGET_SEARCH_NETWORK",
		5:  "CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH",
		6:  "CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN",
		7:  "CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK",
		8:  "CANNOT_TARGET_PARTNER_SEARCH_NETWORK",
		9:  "CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY",
		10: "CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS",
		11: "CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN",
		12: "DUPLICATE_CAMPAIGN_NAME",
		13: "INCOMPATIBLE_CAMPAIGN_FIELD",
		14: "INVALID_CAMPAIGN_NAME",
		15: "INVALID_AD_SERVING_OPTIMIZATION_STATUS",
		16: "INVALID_TRACKING_URL",
		17: "CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING",
		18: "MAX_IMPRESSIONS_NOT_IN_RANGE",
		19: "TIME_UNIT_NOT_SUPPORTED",
		20: "INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED",
		21: "BUDGET_CANNOT_BE_SHARED",
		22: "CAMPAIGN_CANNOT_USE_SHARED_BUDGET",
		23: "CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS",
		24: "CAMPAIGN_LABEL_DOES_NOT_EXIST",
		25: "CAMPAIGN_LABEL_ALREADY_EXISTS",
		26: "MISSING_SHOPPING_SETTING",
		27: "INVALID_SHOPPING_SALES_COUNTRY",
		31: "ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE",
		32: "INVALID_ADVERTISING_CHANNEL_SUB_TYPE",
		33: "AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED",
		34: "CANNOT_SET_AD_ROTATION_MODE",
		35: "CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED",
		36: "CANNOT_SET_DATE_TO_PAST",
		37: "MISSING_HOTEL_CUSTOMER_LINK",
		38: "INVALID_HOTEL_CUSTOMER_LINK",
		39: "MISSING_HOTEL_SETTING",
		40: "CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP",
		41: "APP_NOT_FOUND",
		42: "SHOPPING_ENABLE_LOCAL_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE",
		43: "MERCHANT_NOT_ALLOWED_FOR_COMPARISON_LISTING_ADS",
		44: "INSUFFICIENT_APP_INSTALLS_COUNT",
		45: "SENSITIVE_CATEGORY_APP",
		46: "HEC_AGREEMENT_REQUIRED",
		49: "NOT_COMPATIBLE_WITH_VIEW_THROUGH_CONVERSION_OPTIMIZATION",
		48: "INVALID_EXCLUDED_PARENT_ASSET_FIELD_TYPE",
		50: "CANNOT_CREATE_APP_PRE_REGISTRATION_FOR_NON_ANDROID_APP",
		51: "APP_NOT_AVAILABLE_TO_CREATE_APP_PRE_REGISTRATION_CAMPAIGN",
		52: "INCOMPATIBLE_BUDGET_TYPE",
		53: "LOCAL_SERVICES_DUPLICATE_CATEGORY_BID",
		54: "LOCAL_SERVICES_INVALID_CATEGORY_BID",
		55: "LOCAL_SERVICES_MISSING_CATEGORY_BID",
	}
	CampaignErrorEnum_CampaignError_value = map[string]int32{
		"UNSPECIFIED":                   0,
		"UNKNOWN":                       1,
		"CANNOT_TARGET_CONTENT_NETWORK": 3,
		"CANNOT_TARGET_SEARCH_NETWORK":  4,
		"CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH":                      5,
		"CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN":                            6,
		"CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK":                               7,
		"CANNOT_TARGET_PARTNER_SEARCH_NETWORK":                                    8,
		"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY": 9,
		"CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS":                      10,
		"CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN":                                        11,
		"DUPLICATE_CAMPAIGN_NAME":                                                 12,
		"INCOMPATIBLE_CAMPAIGN_FIELD":                                             13,
		"INVALID_CAMPAIGN_NAME":                                                   14,
		"INVALID_AD_SERVING_OPTIMIZATION_STATUS":                                  15,
		"INVALID_TRACKING_URL":                                                    16,
		"CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING":              17,
		"MAX_IMPRESSIONS_NOT_IN_RANGE":                                            18,
		"TIME_UNIT_NOT_SUPPORTED":                                                 19,
		"INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED":                           20,
		"BUDGET_CANNOT_BE_SHARED":                                                 21,
		"CAMPAIGN_CANNOT_USE_SHARED_BUDGET":                                       22,
		"CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS":                            23,
		"CAMPAIGN_LABEL_DOES_NOT_EXIST":                                           24,
		"CAMPAIGN_LABEL_ALREADY_EXISTS":                                           25,
		"MISSING_SHOPPING_SETTING":                                                26,
		"INVALID_SHOPPING_SALES_COUNTRY":                                          27,
		"ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE":                 31,
		"INVALID_ADVERTISING_CHANNEL_SUB_TYPE":                                    32,
		"AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED":                                33,
		"CANNOT_SET_AD_ROTATION_MODE":                                             34,
		"CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED":                             35,
		"CANNOT_SET_DATE_TO_PAST":                                                 36,
		"MISSING_HOTEL_CUSTOMER_LINK":                                             37,
		"INVALID_HOTEL_CUSTOMER_LINK":                                             38,
		"MISSING_HOTEL_SETTING":                                                   39,
		"CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP":          40,
		"APP_NOT_FOUND":                                                           41,
		"SHOPPING_ENABLE_LOCAL_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE":                   42,
		"MERCHANT_NOT_ALLOWED_FOR_COMPARISON_LISTING_ADS":                         43,
		"INSUFFICIENT_APP_INSTALLS_COUNT":                                         44,
		"SENSITIVE_CATEGORY_APP":                                                  45,
		"HEC_AGREEMENT_REQUIRED":                                                  46,
		"NOT_COMPATIBLE_WITH_VIEW_THROUGH_CONVERSION_OPTIMIZATION":                49,
		"INVALID_EXCLUDED_PARENT_ASSET_FIELD_TYPE":                                48,
		"CANNOT_CREATE_APP_PRE_REGISTRATION_FOR_NON_ANDROID_APP":                  50,
		"APP_NOT_AVAILABLE_TO_CREATE_APP_PRE_REGISTRATION_CAMPAIGN":               51,
		"INCOMPATIBLE_BUDGET_TYPE":                                                52,
		"LOCAL_SERVICES_DUPLICATE_CATEGORY_BID":                                   53,
		"LOCAL_SERVICES_INVALID_CATEGORY_BID":                                     54,
		"LOCAL_SERVICES_MISSING_CATEGORY_BID":                                     55,
	}
)

func (x CampaignErrorEnum_CampaignError) Enum() *CampaignErrorEnum_CampaignError {
	p := new(CampaignErrorEnum_CampaignError)
	*p = x
	return p
}

func (x CampaignErrorEnum_CampaignError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignErrorEnum_CampaignError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_campaign_error_proto_enumTypes[0].Descriptor()
}

func (CampaignErrorEnum_CampaignError) Type() protoreflect.EnumType {
	return &file_errors_campaign_error_proto_enumTypes[0]
}

func (x CampaignErrorEnum_CampaignError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignErrorEnum_CampaignError.Descriptor instead.
func (CampaignErrorEnum_CampaignError) EnumDescriptor() ([]byte, []int) {
	return file_errors_campaign_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible campaign errors.
type CampaignErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CampaignErrorEnum) Reset() {
	*x = CampaignErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_campaign_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignErrorEnum) ProtoMessage() {}

func (x *CampaignErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_campaign_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignErrorEnum.ProtoReflect.Descriptor instead.
func (*CampaignErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_campaign_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_campaign_error_proto protoreflect.FileDescriptor

var file_errors_campaign_error_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xcc, 0x10, 0x0a, 0x11, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xb6, 0x10,
	0x0a, 0x0d, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a,
	0x1d, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x43,
	0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x03,
	0x12, 0x20, 0x0a, 0x1c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b,
	0x10, 0x04, 0x12, 0x36, 0x0a, 0x32, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52,
	0x47, 0x45, 0x54, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f,
	0x52, 0x4b, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c,
	0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x05, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x41,
	0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x47, 0x4f, 0x4f, 0x47,
	0x4c, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x50,
	0x4d, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x06, 0x12, 0x2d, 0x0a, 0x29,
	0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x54, 0x41,
	0x52, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x54, 0x5f, 0x4c, 0x45, 0x41, 0x53, 0x54, 0x5f, 0x4f, 0x4e,
	0x45, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x07, 0x12, 0x28, 0x0a, 0x24, 0x43,
	0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x50, 0x41, 0x52,
	0x54, 0x4e, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x4e, 0x45, 0x54, 0x57,
	0x4f, 0x52, 0x4b, 0x10, 0x08, 0x12, 0x4b, 0x0a, 0x47, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f,
	0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4e,
	0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x57, 0x49, 0x54, 0x48,
	0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x41, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59,
	0x10, 0x09, 0x12, 0x36, 0x0a, 0x32, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x44,
	0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x41, 0x49, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x41, 0x42, 0x4c,
	0x45, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x53, 0x10, 0x0a, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x41,
	0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x0b,
	0x12, 0x1b, 0x0a, 0x17, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41,
	0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0c, 0x12, 0x1f, 0x0a,
	0x1b, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x41,
	0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x0d, 0x12, 0x19,
	0x0a, 0x15, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49,
	0x47, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0e, 0x12, 0x2a, 0x0a, 0x26, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f,
	0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x10, 0x0f, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x10, 0x12,
	0x3e, 0x0a, 0x3a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x42, 0x4f,
	0x54, 0x48, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x5f,
	0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x52, 0x41,
	0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x11, 0x12,
	0x20, 0x0a, 0x1c, 0x4d, 0x41, 0x58, 0x5f, 0x49, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10,
	0x12, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x13, 0x12, 0x31,
	0x0a, 0x2d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x46, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10,
	0x14, 0x12, 0x1b, 0x0a, 0x17, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x10, 0x15, 0x12, 0x25,
	0x0a, 0x21, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x44,
	0x47, 0x45, 0x54, 0x10, 0x16, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x4f, 0x4e,
	0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x54,
	0x52, 0x49, 0x41, 0x4c, 0x53, 0x10, 0x17, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x41, 0x4d, 0x50, 0x41,
	0x49, 0x47, 0x4e, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f, 0x44, 0x4f, 0x45, 0x53, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x18, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x41,
	0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f, 0x41, 0x4c, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x19, 0x12, 0x1c, 0x0a,
	0x18, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x1a, 0x12, 0x22, 0x0a, 0x1e, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f,
	0x53, 0x41, 0x4c, 0x45, 0x53, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x1b, 0x12,
	0x3b, 0x0a, 0x37, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x1f, 0x12, 0x28, 0x0a, 0x24,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53,
	0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x53, 0x55, 0x42, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x20, 0x12, 0x2c, 0x0a, 0x28, 0x41, 0x54, 0x5f, 0x4c, 0x45, 0x41,
	0x53, 0x54, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x21, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x45, 0x54, 0x5f, 0x41, 0x44, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x10, 0x22, 0x12, 0x2f, 0x0a, 0x2b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f,
	0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x49, 0x46, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x45, 0x44, 0x10, 0x23, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x53, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x50, 0x41, 0x53,
	0x54, 0x10, 0x24, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x48,
	0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4c, 0x49,
	0x4e, 0x4b, 0x10, 0x25, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x48, 0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4c,
	0x49, 0x4e, 0x4b, 0x10, 0x26, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47,
	0x5f, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x27,
	0x12, 0x42, 0x0a, 0x3e, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x53,
	0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42,
	0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x57, 0x48, 0x49, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54,
	0x5f, 0x4f, 0x46, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x10, 0x28, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x50, 0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x29, 0x12, 0x39, 0x0a, 0x35, 0x53, 0x48, 0x4f, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46,
	0x4f, 0x52, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x2a, 0x12, 0x33, 0x0a, 0x2f, 0x4d, 0x45, 0x52, 0x43, 0x48, 0x41, 0x4e, 0x54, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x41, 0x44, 0x53, 0x10, 0x2b, 0x12, 0x23, 0x0a, 0x1f, 0x49, 0x4e, 0x53, 0x55, 0x46,
	0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x54,
	0x41, 0x4c, 0x4c, 0x53, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x2c, 0x12, 0x1a, 0x0a, 0x16,
	0x53, 0x45, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x2d, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x45, 0x43, 0x5f,
	0x41, 0x47, 0x52, 0x45, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52,
	0x45, 0x44, 0x10, 0x2e, 0x12, 0x3c, 0x0a, 0x38, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x56, 0x49, 0x45, 0x57,
	0x5f, 0x54, 0x48, 0x52, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x31, 0x12, 0x2c, 0x0a, 0x28, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x45, 0x58,
	0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x53,
	0x53, 0x45, 0x54, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x30,
	0x12, 0x3a, 0x0a, 0x36, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x52, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x41,
	0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x32, 0x12, 0x3d, 0x0a, 0x39,
	0x41, 0x50, 0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c,
	0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x5f,
	0x50, 0x52, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x33, 0x12, 0x1c, 0x0a, 0x18, 0x49,
	0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x42, 0x55, 0x44, 0x47,
	0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x34, 0x12, 0x29, 0x0a, 0x25, 0x4c, 0x4f, 0x43,
	0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x53, 0x5f, 0x44, 0x55, 0x50, 0x4c,
	0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x42,
	0x49, 0x44, 0x10, 0x35, 0x12, 0x27, 0x0a, 0x23, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x42, 0x49, 0x44, 0x10, 0x36, 0x12, 0x27, 0x0a,
	0x23, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x53, 0x5f,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x42, 0x49, 0x44, 0x10, 0x37, 0x42, 0xf2, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x12,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x31, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_errors_campaign_error_proto_rawDescOnce sync.Once
	file_errors_campaign_error_proto_rawDescData = file_errors_campaign_error_proto_rawDesc
)

func file_errors_campaign_error_proto_rawDescGZIP() []byte {
	file_errors_campaign_error_proto_rawDescOnce.Do(func() {
		file_errors_campaign_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_campaign_error_proto_rawDescData)
	})
	return file_errors_campaign_error_proto_rawDescData
}

var file_errors_campaign_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_campaign_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_campaign_error_proto_goTypes = []interface{}{
	(CampaignErrorEnum_CampaignError)(0), // 0: google.ads.googleads.v11.errors.CampaignErrorEnum.CampaignError
	(*CampaignErrorEnum)(nil),            // 1: google.ads.googleads.v11.errors.CampaignErrorEnum
}
var file_errors_campaign_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_campaign_error_proto_init() }
func file_errors_campaign_error_proto_init() {
	if File_errors_campaign_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_campaign_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignErrorEnum); i {
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
			RawDescriptor: file_errors_campaign_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_campaign_error_proto_goTypes,
		DependencyIndexes: file_errors_campaign_error_proto_depIdxs,
		EnumInfos:         file_errors_campaign_error_proto_enumTypes,
		MessageInfos:      file_errors_campaign_error_proto_msgTypes,
	}.Build()
	File_errors_campaign_error_proto = out.File
	file_errors_campaign_error_proto_rawDesc = nil
	file_errors_campaign_error_proto_goTypes = nil
	file_errors_campaign_error_proto_depIdxs = nil
}
