/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// GetOrganizationUsageDetailsReply struct for GetOrganizationUsageDetailsReply
type GetOrganizationUsageDetailsReply struct {
	UsageDetails []UsageDetails `json:"usage_details,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Count *int64 `json:"count,omitempty"`
	Order *string `json:"order,omitempty"`
}

// NewGetOrganizationUsageDetailsReply instantiates a new GetOrganizationUsageDetailsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationUsageDetailsReply() *GetOrganizationUsageDetailsReply {
	this := GetOrganizationUsageDetailsReply{}
	return &this
}

// NewGetOrganizationUsageDetailsReplyWithDefaults instantiates a new GetOrganizationUsageDetailsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationUsageDetailsReplyWithDefaults() *GetOrganizationUsageDetailsReply {
	this := GetOrganizationUsageDetailsReply{}
	return &this
}

// GetUsageDetails returns the UsageDetails field value if set, zero value otherwise.
func (o *GetOrganizationUsageDetailsReply) GetUsageDetails() []UsageDetails {
	if o == nil || isNil(o.UsageDetails) {
		var ret []UsageDetails
		return ret
	}
	return o.UsageDetails
}

// GetUsageDetailsOk returns a tuple with the UsageDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationUsageDetailsReply) GetUsageDetailsOk() ([]UsageDetails, bool) {
	if o == nil || isNil(o.UsageDetails) {
    return nil, false
	}
	return o.UsageDetails, true
}

// HasUsageDetails returns a boolean if a field has been set.
func (o *GetOrganizationUsageDetailsReply) HasUsageDetails() bool {
	if o != nil && !isNil(o.UsageDetails) {
		return true
	}

	return false
}

// SetUsageDetails gets a reference to the given []UsageDetails and assigns it to the UsageDetails field.
func (o *GetOrganizationUsageDetailsReply) SetUsageDetails(v []UsageDetails) {
	o.UsageDetails = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetOrganizationUsageDetailsReply) GetLimit() int64 {
	if o == nil || isNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationUsageDetailsReply) GetLimitOk() (*int64, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetOrganizationUsageDetailsReply) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *GetOrganizationUsageDetailsReply) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetOrganizationUsageDetailsReply) GetOffset() int64 {
	if o == nil || isNil(o.Offset) {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationUsageDetailsReply) GetOffsetOk() (*int64, bool) {
	if o == nil || isNil(o.Offset) {
    return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetOrganizationUsageDetailsReply) HasOffset() bool {
	if o != nil && !isNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *GetOrganizationUsageDetailsReply) SetOffset(v int64) {
	o.Offset = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetOrganizationUsageDetailsReply) GetCount() int64 {
	if o == nil || isNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationUsageDetailsReply) GetCountOk() (*int64, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetOrganizationUsageDetailsReply) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *GetOrganizationUsageDetailsReply) SetCount(v int64) {
	o.Count = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GetOrganizationUsageDetailsReply) GetOrder() string {
	if o == nil || isNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationUsageDetailsReply) GetOrderOk() (*string, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GetOrganizationUsageDetailsReply) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *GetOrganizationUsageDetailsReply) SetOrder(v string) {
	o.Order = &v
}

func (o GetOrganizationUsageDetailsReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UsageDetails) {
		toSerialize["usage_details"] = o.UsageDetails
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationUsageDetailsReply struct {
	value *GetOrganizationUsageDetailsReply
	isSet bool
}

func (v NullableGetOrganizationUsageDetailsReply) Get() *GetOrganizationUsageDetailsReply {
	return v.value
}

func (v *NullableGetOrganizationUsageDetailsReply) Set(val *GetOrganizationUsageDetailsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationUsageDetailsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationUsageDetailsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationUsageDetailsReply(val *GetOrganizationUsageDetailsReply) *NullableGetOrganizationUsageDetailsReply {
	return &NullableGetOrganizationUsageDetailsReply{value: val, isSet: true}
}

func (v NullableGetOrganizationUsageDetailsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationUsageDetailsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


