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

// DeploymentGPUInfo struct for DeploymentGPUInfo
type DeploymentGPUInfo struct {
	Name *string `json:"name,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewDeploymentGPUInfo instantiates a new DeploymentGPUInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentGPUInfo() *DeploymentGPUInfo {
	this := DeploymentGPUInfo{}
	return &this
}

// NewDeploymentGPUInfoWithDefaults instantiates a new DeploymentGPUInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentGPUInfoWithDefaults() *DeploymentGPUInfo {
	this := DeploymentGPUInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentGPUInfo) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGPUInfo) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentGPUInfo) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentGPUInfo) SetName(v string) {
	o.Name = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DeploymentGPUInfo) GetCount() int64 {
	if o == nil || isNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGPUInfo) GetCountOk() (*int64, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DeploymentGPUInfo) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *DeploymentGPUInfo) SetCount(v int64) {
	o.Count = &v
}

func (o DeploymentGPUInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentGPUInfo struct {
	value *DeploymentGPUInfo
	isSet bool
}

func (v NullableDeploymentGPUInfo) Get() *DeploymentGPUInfo {
	return v.value
}

func (v *NullableDeploymentGPUInfo) Set(val *DeploymentGPUInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentGPUInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentGPUInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentGPUInfo(val *DeploymentGPUInfo) *NullableDeploymentGPUInfo {
	return &NullableDeploymentGPUInfo{value: val, isSet: true}
}

func (v NullableDeploymentGPUInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentGPUInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


