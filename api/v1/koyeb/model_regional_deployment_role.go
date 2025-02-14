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

// RegionalDeploymentRole the model 'RegionalDeploymentRole'
type RegionalDeploymentRole string

// List of RegionalDeployment.Role
const (
	REGIONALDEPLOYMENTROLE_INVALID RegionalDeploymentRole = "INVALID"
	REGIONALDEPLOYMENTROLE_ACTIVE RegionalDeploymentRole = "ACTIVE"
	REGIONALDEPLOYMENTROLE_UPCOMING RegionalDeploymentRole = "UPCOMING"
	REGIONALDEPLOYMENTROLE_CURRENT RegionalDeploymentRole = "CURRENT"
)

// All allowed values of RegionalDeploymentRole enum
var AllowedRegionalDeploymentRoleEnumValues = []RegionalDeploymentRole{
	"INVALID",
	"ACTIVE",
	"UPCOMING",
	"CURRENT",
}

func (v *RegionalDeploymentRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegionalDeploymentRole(value)
	for _, existing := range AllowedRegionalDeploymentRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegionalDeploymentRole", value)
}

// NewRegionalDeploymentRoleFromValue returns a pointer to a valid RegionalDeploymentRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegionalDeploymentRoleFromValue(v string) (*RegionalDeploymentRole, error) {
	ev := RegionalDeploymentRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegionalDeploymentRole: valid values are %v", v, AllowedRegionalDeploymentRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegionalDeploymentRole) IsValid() bool {
	for _, existing := range AllowedRegionalDeploymentRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RegionalDeployment.Role value
func (v RegionalDeploymentRole) Ptr() *RegionalDeploymentRole {
	return &v
}

type NullableRegionalDeploymentRole struct {
	value *RegionalDeploymentRole
	isSet bool
}

func (v NullableRegionalDeploymentRole) Get() *RegionalDeploymentRole {
	return v.value
}

func (v *NullableRegionalDeploymentRole) Set(val *RegionalDeploymentRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalDeploymentRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalDeploymentRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalDeploymentRole(val *RegionalDeploymentRole) *NullableRegionalDeploymentRole {
	return &NullableRegionalDeploymentRole{value: val, isSet: true}
}

func (v NullableRegionalDeploymentRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalDeploymentRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

