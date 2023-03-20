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

// CreateAccountRequest Create new account
type CreateAccountRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Name *string `json:"name,omitempty"`
	Captcha *string `json:"captcha,omitempty"`
	UseV2 *bool `json:"use_v2,omitempty"`
}

// NewCreateAccountRequest instantiates a new CreateAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountRequest(email string, password string) *CreateAccountRequest {
	this := CreateAccountRequest{}
	this.Email = email
	this.Password = password
	return &this
}

// NewCreateAccountRequestWithDefaults instantiates a new CreateAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountRequestWithDefaults() *CreateAccountRequest {
	this := CreateAccountRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *CreateAccountRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateAccountRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *CreateAccountRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateAccountRequest) SetPassword(v string) {
	o.Password = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateAccountRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateAccountRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateAccountRequest) SetName(v string) {
	o.Name = &v
}

// GetCaptcha returns the Captcha field value if set, zero value otherwise.
func (o *CreateAccountRequest) GetCaptcha() string {
	if o == nil || isNil(o.Captcha) {
		var ret string
		return ret
	}
	return *o.Captcha
}

// GetCaptchaOk returns a tuple with the Captcha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetCaptchaOk() (*string, bool) {
	if o == nil || isNil(o.Captcha) {
    return nil, false
	}
	return o.Captcha, true
}

// HasCaptcha returns a boolean if a field has been set.
func (o *CreateAccountRequest) HasCaptcha() bool {
	if o != nil && !isNil(o.Captcha) {
		return true
	}

	return false
}

// SetCaptcha gets a reference to the given string and assigns it to the Captcha field.
func (o *CreateAccountRequest) SetCaptcha(v string) {
	o.Captcha = &v
}

// GetUseV2 returns the UseV2 field value if set, zero value otherwise.
func (o *CreateAccountRequest) GetUseV2() bool {
	if o == nil || isNil(o.UseV2) {
		var ret bool
		return ret
	}
	return *o.UseV2
}

// GetUseV2Ok returns a tuple with the UseV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetUseV2Ok() (*bool, bool) {
	if o == nil || isNil(o.UseV2) {
    return nil, false
	}
	return o.UseV2, true
}

// HasUseV2 returns a boolean if a field has been set.
func (o *CreateAccountRequest) HasUseV2() bool {
	if o != nil && !isNil(o.UseV2) {
		return true
	}

	return false
}

// SetUseV2 gets a reference to the given bool and assigns it to the UseV2 field.
func (o *CreateAccountRequest) SetUseV2(v bool) {
	o.UseV2 = &v
}

func (o CreateAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Captcha) {
		toSerialize["captcha"] = o.Captcha
	}
	if !isNil(o.UseV2) {
		toSerialize["use_v2"] = o.UseV2
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAccountRequest struct {
	value *CreateAccountRequest
	isSet bool
}

func (v NullableCreateAccountRequest) Get() *CreateAccountRequest {
	return v.value
}

func (v *NullableCreateAccountRequest) Set(val *CreateAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountRequest(val *CreateAccountRequest) *NullableCreateAccountRequest {
	return &NullableCreateAccountRequest{value: val, isSet: true}
}

func (v NullableCreateAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


