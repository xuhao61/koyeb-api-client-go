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

// CreateArchive struct for CreateArchive
type CreateArchive struct {
	// How much space to provision for the archive, in bytes.
	Size *string `json:"size,omitempty"`
}

// NewCreateArchive instantiates a new CreateArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateArchive() *CreateArchive {
	this := CreateArchive{}
	return &this
}

// NewCreateArchiveWithDefaults instantiates a new CreateArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateArchiveWithDefaults() *CreateArchive {
	this := CreateArchive{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CreateArchive) GetSize() string {
	if o == nil || isNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArchive) GetSizeOk() (*string, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CreateArchive) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *CreateArchive) SetSize(v string) {
	o.Size = &v
}

func (o CreateArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableCreateArchive struct {
	value *CreateArchive
	isSet bool
}

func (v NullableCreateArchive) Get() *CreateArchive {
	return v.value
}

func (v *NullableCreateArchive) Set(val *CreateArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateArchive(val *CreateArchive) *NullableCreateArchive {
	return &NullableCreateArchive{value: val, isSet: true}
}

func (v NullableCreateArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


