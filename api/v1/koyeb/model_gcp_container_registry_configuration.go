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

// GCPContainerRegistryConfiguration struct for GCPContainerRegistryConfiguration
type GCPContainerRegistryConfiguration struct {
	KeyfileContent *string `json:"keyfile_content,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewGCPContainerRegistryConfiguration instantiates a new GCPContainerRegistryConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPContainerRegistryConfiguration() *GCPContainerRegistryConfiguration {
	this := GCPContainerRegistryConfiguration{}
	return &this
}

// NewGCPContainerRegistryConfigurationWithDefaults instantiates a new GCPContainerRegistryConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPContainerRegistryConfigurationWithDefaults() *GCPContainerRegistryConfiguration {
	this := GCPContainerRegistryConfiguration{}
	return &this
}

// GetKeyfileContent returns the KeyfileContent field value if set, zero value otherwise.
func (o *GCPContainerRegistryConfiguration) GetKeyfileContent() string {
	if o == nil || o.KeyfileContent == nil {
		var ret string
		return ret
	}
	return *o.KeyfileContent
}

// GetKeyfileContentOk returns a tuple with the KeyfileContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPContainerRegistryConfiguration) GetKeyfileContentOk() (*string, bool) {
	if o == nil || o.KeyfileContent == nil {
		return nil, false
	}
	return o.KeyfileContent, true
}

// HasKeyfileContent returns a boolean if a field has been set.
func (o *GCPContainerRegistryConfiguration) HasKeyfileContent() bool {
	if o != nil && o.KeyfileContent != nil {
		return true
	}

	return false
}

// SetKeyfileContent gets a reference to the given string and assigns it to the KeyfileContent field.
func (o *GCPContainerRegistryConfiguration) SetKeyfileContent(v string) {
	o.KeyfileContent = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GCPContainerRegistryConfiguration) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPContainerRegistryConfiguration) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GCPContainerRegistryConfiguration) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GCPContainerRegistryConfiguration) SetUrl(v string) {
	o.Url = &v
}

func (o GCPContainerRegistryConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyfileContent != nil {
		toSerialize["keyfile_content"] = o.KeyfileContent
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableGCPContainerRegistryConfiguration struct {
	value *GCPContainerRegistryConfiguration
	isSet bool
}

func (v NullableGCPContainerRegistryConfiguration) Get() *GCPContainerRegistryConfiguration {
	return v.value
}

func (v *NullableGCPContainerRegistryConfiguration) Set(val *GCPContainerRegistryConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPContainerRegistryConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPContainerRegistryConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPContainerRegistryConfiguration(val *GCPContainerRegistryConfiguration) *NullableGCPContainerRegistryConfiguration {
	return &NullableGCPContainerRegistryConfiguration{value: val, isSet: true}
}

func (v NullableGCPContainerRegistryConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPContainerRegistryConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


