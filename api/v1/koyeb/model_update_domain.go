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

// UpdateDomain struct for UpdateDomain
type UpdateDomain struct {
	// To attach or detach from an app for custom domain.
	AppId *string `json:"app_id,omitempty"`
	// To change subdomain for auto-assigned domain.
	Subdomain *string `json:"subdomain,omitempty"`
}

// NewUpdateDomain instantiates a new UpdateDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDomain() *UpdateDomain {
	this := UpdateDomain{}
	return &this
}

// NewUpdateDomainWithDefaults instantiates a new UpdateDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDomainWithDefaults() *UpdateDomain {
	this := UpdateDomain{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *UpdateDomain) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomain) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *UpdateDomain) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *UpdateDomain) SetAppId(v string) {
	o.AppId = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *UpdateDomain) GetSubdomain() string {
	if o == nil || isNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomain) GetSubdomainOk() (*string, bool) {
	if o == nil || isNil(o.Subdomain) {
    return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *UpdateDomain) HasSubdomain() bool {
	if o != nil && !isNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *UpdateDomain) SetSubdomain(v string) {
	o.Subdomain = &v
}

func (o UpdateDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDomain struct {
	value *UpdateDomain
	isSet bool
}

func (v NullableUpdateDomain) Get() *UpdateDomain {
	return v.value
}

func (v *NullableUpdateDomain) Set(val *UpdateDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDomain(val *UpdateDomain) *NullableUpdateDomain {
	return &NullableUpdateDomain{value: val, isSet: true}
}

func (v NullableUpdateDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


