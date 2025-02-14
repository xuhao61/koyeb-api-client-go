/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"time"
)

// DeploymentProvisioningInfoStageBuildAttemptBuildStep struct for DeploymentProvisioningInfoStageBuildAttemptBuildStep
type DeploymentProvisioningInfoStageBuildAttemptBuildStep struct {
	Name *string `json:"name,omitempty"`
	Status *DeploymentProvisioningInfoStageStatus `json:"status,omitempty"`
	Messages []string `json:"messages,omitempty"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
}

// NewDeploymentProvisioningInfoStageBuildAttemptBuildStep instantiates a new DeploymentProvisioningInfoStageBuildAttemptBuildStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentProvisioningInfoStageBuildAttemptBuildStep() *DeploymentProvisioningInfoStageBuildAttemptBuildStep {
	this := DeploymentProvisioningInfoStageBuildAttemptBuildStep{}
	var status DeploymentProvisioningInfoStageStatus = DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewDeploymentProvisioningInfoStageBuildAttemptBuildStepWithDefaults instantiates a new DeploymentProvisioningInfoStageBuildAttemptBuildStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentProvisioningInfoStageBuildAttemptBuildStepWithDefaults() *DeploymentProvisioningInfoStageBuildAttemptBuildStep {
	this := DeploymentProvisioningInfoStageBuildAttemptBuildStep{}
	var status DeploymentProvisioningInfoStageStatus = DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetStatus() DeploymentProvisioningInfoStageStatus {
	if o == nil || isNil(o.Status) {
		var ret DeploymentProvisioningInfoStageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetStatusOk() (*DeploymentProvisioningInfoStageStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentProvisioningInfoStageStatus and assigns it to the Status field.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) SetStatus(v DeploymentProvisioningInfoStageStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetMessages() []string {
	if o == nil || isNil(o.Messages) {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetMessagesOk() ([]string, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) SetMessages(v []string) {
	o.Messages = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetStartedAt() time.Time {
	if o == nil || isNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartedAt) {
    return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) HasStartedAt() bool {
	if o != nil && !isNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetFinishedAt() time.Time {
	if o == nil || isNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.FinishedAt) {
    return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) HasFinishedAt() bool {
	if o != nil && !isNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

func (o DeploymentProvisioningInfoStageBuildAttemptBuildStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !isNil(o.FinishedAt) {
		toSerialize["finished_at"] = o.FinishedAt
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentProvisioningInfoStageBuildAttemptBuildStep struct {
	value *DeploymentProvisioningInfoStageBuildAttemptBuildStep
	isSet bool
}

func (v NullableDeploymentProvisioningInfoStageBuildAttemptBuildStep) Get() *DeploymentProvisioningInfoStageBuildAttemptBuildStep {
	return v.value
}

func (v *NullableDeploymentProvisioningInfoStageBuildAttemptBuildStep) Set(val *DeploymentProvisioningInfoStageBuildAttemptBuildStep) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentProvisioningInfoStageBuildAttemptBuildStep) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentProvisioningInfoStageBuildAttemptBuildStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentProvisioningInfoStageBuildAttemptBuildStep(val *DeploymentProvisioningInfoStageBuildAttemptBuildStep) *NullableDeploymentProvisioningInfoStageBuildAttemptBuildStep {
	return &NullableDeploymentProvisioningInfoStageBuildAttemptBuildStep{value: val, isSet: true}
}

func (v NullableDeploymentProvisioningInfoStageBuildAttemptBuildStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentProvisioningInfoStageBuildAttemptBuildStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


