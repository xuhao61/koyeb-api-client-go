/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"fmt"
)

// RegionalDeploymentStatus the model 'RegionalDeploymentStatus'
type RegionalDeploymentStatus string

// List of RegionalDeployment.Status
const (
	REGIONALDEPLOYMENTSTATUS_PENDING RegionalDeploymentStatus = "PENDING"
	REGIONALDEPLOYMENTSTATUS_PROVISIONING RegionalDeploymentStatus = "PROVISIONING"
	REGIONALDEPLOYMENTSTATUS_SCHEDULED RegionalDeploymentStatus = "SCHEDULED"
	REGIONALDEPLOYMENTSTATUS_CANCELING RegionalDeploymentStatus = "CANCELING"
	REGIONALDEPLOYMENTSTATUS_CANCELED RegionalDeploymentStatus = "CANCELED"
	REGIONALDEPLOYMENTSTATUS_ALLOCATING RegionalDeploymentStatus = "ALLOCATING"
	REGIONALDEPLOYMENTSTATUS_STARTING RegionalDeploymentStatus = "STARTING"
	REGIONALDEPLOYMENTSTATUS_HEALTHY RegionalDeploymentStatus = "HEALTHY"
	REGIONALDEPLOYMENTSTATUS_DEGRADED RegionalDeploymentStatus = "DEGRADED"
	REGIONALDEPLOYMENTSTATUS_UNHEALTHY RegionalDeploymentStatus = "UNHEALTHY"
	REGIONALDEPLOYMENTSTATUS_STOPPING RegionalDeploymentStatus = "STOPPING"
	REGIONALDEPLOYMENTSTATUS_STOPPED RegionalDeploymentStatus = "STOPPED"
	REGIONALDEPLOYMENTSTATUS_ERRORING RegionalDeploymentStatus = "ERRORING"
	REGIONALDEPLOYMENTSTATUS_ERROR RegionalDeploymentStatus = "ERROR"
)

// All allowed values of RegionalDeploymentStatus enum
var AllowedRegionalDeploymentStatusEnumValues = []RegionalDeploymentStatus{
	"PENDING",
	"PROVISIONING",
	"SCHEDULED",
	"CANCELING",
	"CANCELED",
	"ALLOCATING",
	"STARTING",
	"HEALTHY",
	"DEGRADED",
	"UNHEALTHY",
	"STOPPING",
	"STOPPED",
	"ERRORING",
	"ERROR",
}

func (v *RegionalDeploymentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegionalDeploymentStatus(value)
	for _, existing := range AllowedRegionalDeploymentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegionalDeploymentStatus", value)
}

// NewRegionalDeploymentStatusFromValue returns a pointer to a valid RegionalDeploymentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegionalDeploymentStatusFromValue(v string) (*RegionalDeploymentStatus, error) {
	ev := RegionalDeploymentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegionalDeploymentStatus: valid values are %v", v, AllowedRegionalDeploymentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegionalDeploymentStatus) IsValid() bool {
	for _, existing := range AllowedRegionalDeploymentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RegionalDeployment.Status value
func (v RegionalDeploymentStatus) Ptr() *RegionalDeploymentStatus {
	return &v
}

type NullableRegionalDeploymentStatus struct {
	value *RegionalDeploymentStatus
	isSet bool
}

func (v NullableRegionalDeploymentStatus) Get() *RegionalDeploymentStatus {
	return v.value
}

func (v *NullableRegionalDeploymentStatus) Set(val *RegionalDeploymentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalDeploymentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalDeploymentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalDeploymentStatus(val *RegionalDeploymentStatus) *NullableRegionalDeploymentStatus {
	return &NullableRegionalDeploymentStatus{value: val, isSet: true}
}

func (v NullableRegionalDeploymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalDeploymentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

