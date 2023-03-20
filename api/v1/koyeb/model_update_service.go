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

// UpdateService struct for UpdateService
type UpdateService struct {
	Definition *DeploymentDefinition `json:"definition,omitempty"`
}

// NewUpdateService instantiates a new UpdateService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateService() *UpdateService {
	this := UpdateService{}
	return &this
}

// NewUpdateServiceWithDefaults instantiates a new UpdateService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceWithDefaults() *UpdateService {
	this := UpdateService{}
	return &this
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *UpdateService) GetDefinition() DeploymentDefinition {
	if o == nil || isNil(o.Definition) {
		var ret DeploymentDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateService) GetDefinitionOk() (*DeploymentDefinition, bool) {
	if o == nil || isNil(o.Definition) {
    return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *UpdateService) HasDefinition() bool {
	if o != nil && !isNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given DeploymentDefinition and assigns it to the Definition field.
func (o *UpdateService) SetDefinition(v DeploymentDefinition) {
	o.Definition = &v
}

func (o UpdateService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateService struct {
	value *UpdateService
	isSet bool
}

func (v NullableUpdateService) Get() *UpdateService {
	return v.value
}

func (v *NullableUpdateService) Set(val *UpdateService) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateService) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateService(val *UpdateService) *NullableUpdateService {
	return &NullableUpdateService{value: val, isSet: true}
}

func (v NullableUpdateService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


