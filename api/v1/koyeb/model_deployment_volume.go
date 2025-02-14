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

// DeploymentVolume struct for DeploymentVolume
type DeploymentVolume struct {
	Id *string `json:"id,omitempty"`
	Path *string `json:"path,omitempty"`
	ReplicaIndex *int64 `json:"replica_index,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

// NewDeploymentVolume instantiates a new DeploymentVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentVolume() *DeploymentVolume {
	this := DeploymentVolume{}
	return &this
}

// NewDeploymentVolumeWithDefaults instantiates a new DeploymentVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentVolumeWithDefaults() *DeploymentVolume {
	this := DeploymentVolume{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeploymentVolume) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentVolume) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentVolume) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeploymentVolume) SetId(v string) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DeploymentVolume) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentVolume) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DeploymentVolume) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DeploymentVolume) SetPath(v string) {
	o.Path = &v
}

// GetReplicaIndex returns the ReplicaIndex field value if set, zero value otherwise.
func (o *DeploymentVolume) GetReplicaIndex() int64 {
	if o == nil || isNil(o.ReplicaIndex) {
		var ret int64
		return ret
	}
	return *o.ReplicaIndex
}

// GetReplicaIndexOk returns a tuple with the ReplicaIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentVolume) GetReplicaIndexOk() (*int64, bool) {
	if o == nil || isNil(o.ReplicaIndex) {
    return nil, false
	}
	return o.ReplicaIndex, true
}

// HasReplicaIndex returns a boolean if a field has been set.
func (o *DeploymentVolume) HasReplicaIndex() bool {
	if o != nil && !isNil(o.ReplicaIndex) {
		return true
	}

	return false
}

// SetReplicaIndex gets a reference to the given int64 and assigns it to the ReplicaIndex field.
func (o *DeploymentVolume) SetReplicaIndex(v int64) {
	o.ReplicaIndex = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *DeploymentVolume) GetScopes() []string {
	if o == nil || isNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentVolume) GetScopesOk() ([]string, bool) {
	if o == nil || isNil(o.Scopes) {
    return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *DeploymentVolume) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *DeploymentVolume) SetScopes(v []string) {
	o.Scopes = v
}

func (o DeploymentVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.ReplicaIndex) {
		toSerialize["replica_index"] = o.ReplicaIndex
	}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentVolume struct {
	value *DeploymentVolume
	isSet bool
}

func (v NullableDeploymentVolume) Get() *DeploymentVolume {
	return v.value
}

func (v *NullableDeploymentVolume) Set(val *DeploymentVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentVolume(val *DeploymentVolume) *NullableDeploymentVolume {
	return &NullableDeploymentVolume{value: val, isSet: true}
}

func (v NullableDeploymentVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


