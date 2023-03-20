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

// CreateService struct for CreateService
type CreateService struct {
	AppId *string `json:"app_id,omitempty"`
	Definition *DeploymentDefinition `json:"definition,omitempty"`
}

// NewCreateService instantiates a new CreateService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateService() *CreateService {
	this := CreateService{}
	return &this
}

// NewCreateServiceWithDefaults instantiates a new CreateService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceWithDefaults() *CreateService {
	this := CreateService{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *CreateService) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateService) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *CreateService) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *CreateService) SetAppId(v string) {
	o.AppId = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *CreateService) GetDefinition() DeploymentDefinition {
	if o == nil || isNil(o.Definition) {
		var ret DeploymentDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateService) GetDefinitionOk() (*DeploymentDefinition, bool) {
	if o == nil || isNil(o.Definition) {
    return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *CreateService) HasDefinition() bool {
	if o != nil && !isNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given DeploymentDefinition and assigns it to the Definition field.
func (o *CreateService) SetDefinition(v DeploymentDefinition) {
	o.Definition = &v
}

func (o CreateService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	return json.Marshal(toSerialize)
}

type NullableCreateService struct {
	value *CreateService
	isSet bool
}

func (v NullableCreateService) Get() *CreateService {
	return v.value
}

func (v *NullableCreateService) Set(val *CreateService) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateService) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateService(val *CreateService) *NullableCreateService {
	return &NullableCreateService{value: val, isSet: true}
}

func (v NullableCreateService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


