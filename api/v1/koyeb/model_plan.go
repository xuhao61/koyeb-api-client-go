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

// Plan the model 'Plan'
type Plan string

// List of Plan
const (
	PLAN_HOBBY Plan = "hobby"
	PLAN_STARTER Plan = "starter"
	PLAN_STARTUP Plan = "startup"
	PLAN_BUSINESS Plan = "business"
	PLAN_ENTERPRISE Plan = "enterprise"
)

// All allowed values of Plan enum
var AllowedPlanEnumValues = []Plan{
	"hobby",
	"starter",
	"startup",
	"business",
	"enterprise",
}

func (v *Plan) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Plan(value)
	for _, existing := range AllowedPlanEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Plan", value)
}

// NewPlanFromValue returns a pointer to a valid Plan
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlanFromValue(v string) (*Plan, error) {
	ev := Plan(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Plan: valid values are %v", v, AllowedPlanEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Plan) IsValid() bool {
	for _, existing := range AllowedPlanEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Plan value
func (v Plan) Ptr() *Plan {
	return &v
}

type NullablePlan struct {
	value *Plan
	isSet bool
}

func (v NullablePlan) Get() *Plan {
	return v.value
}

func (v *NullablePlan) Set(val *Plan) {
	v.value = val
	v.isSet = true
}

func (v NullablePlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlan(val *Plan) *NullablePlan {
	return &NullablePlan{value: val, isSet: true}
}

func (v NullablePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

