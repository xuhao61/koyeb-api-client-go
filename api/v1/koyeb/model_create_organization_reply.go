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

// CreateOrganizationReply struct for CreateOrganizationReply
type CreateOrganizationReply struct {
	Organization *Organization `json:"organization,omitempty"`
}

// NewCreateOrganizationReply instantiates a new CreateOrganizationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationReply() *CreateOrganizationReply {
	this := CreateOrganizationReply{}
	return &this
}

// NewCreateOrganizationReplyWithDefaults instantiates a new CreateOrganizationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationReplyWithDefaults() *CreateOrganizationReply {
	this := CreateOrganizationReply{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CreateOrganizationReply) GetOrganization() Organization {
	if o == nil || isNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationReply) GetOrganizationOk() (*Organization, bool) {
	if o == nil || isNil(o.Organization) {
    return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CreateOrganizationReply) HasOrganization() bool {
	if o != nil && !isNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *CreateOrganizationReply) SetOrganization(v Organization) {
	o.Organization = &v
}

func (o CreateOrganizationReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationReply struct {
	value *CreateOrganizationReply
	isSet bool
}

func (v NullableCreateOrganizationReply) Get() *CreateOrganizationReply {
	return v.value
}

func (v *NullableCreateOrganizationReply) Set(val *CreateOrganizationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationReply(val *CreateOrganizationReply) *NullableCreateOrganizationReply {
	return &NullableCreateOrganizationReply{value: val, isSet: true}
}

func (v NullableCreateOrganizationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


