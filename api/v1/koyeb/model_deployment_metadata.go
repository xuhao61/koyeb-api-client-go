/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// DeploymentMetadata struct for DeploymentMetadata
type DeploymentMetadata struct {
	Trigger *TriggerDeploymentMetadata `json:"trigger,omitempty"`
}

// NewDeploymentMetadata instantiates a new DeploymentMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentMetadata() *DeploymentMetadata {
	this := DeploymentMetadata{}
	return &this
}

// NewDeploymentMetadataWithDefaults instantiates a new DeploymentMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentMetadataWithDefaults() *DeploymentMetadata {
	this := DeploymentMetadata{}
	return &this
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *DeploymentMetadata) GetTrigger() TriggerDeploymentMetadata {
	if o == nil || o.Trigger == nil {
		var ret TriggerDeploymentMetadata
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentMetadata) GetTriggerOk() (*TriggerDeploymentMetadata, bool) {
	if o == nil || o.Trigger == nil {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *DeploymentMetadata) HasTrigger() bool {
	if o != nil && o.Trigger != nil {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given TriggerDeploymentMetadata and assigns it to the Trigger field.
func (o *DeploymentMetadata) SetTrigger(v TriggerDeploymentMetadata) {
	o.Trigger = &v
}

func (o DeploymentMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Trigger != nil {
		toSerialize["trigger"] = o.Trigger
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentMetadata struct {
	value *DeploymentMetadata
	isSet bool
}

func (v NullableDeploymentMetadata) Get() *DeploymentMetadata {
	return v.value
}

func (v *NullableDeploymentMetadata) Set(val *DeploymentMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentMetadata(val *DeploymentMetadata) *NullableDeploymentMetadata {
	return &NullableDeploymentMetadata{value: val, isSet: true}
}

func (v NullableDeploymentMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


