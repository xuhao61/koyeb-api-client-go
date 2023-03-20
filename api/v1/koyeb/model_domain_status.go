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

// DomainStatus the model 'DomainStatus'
type DomainStatus string

// List of Domain.Status
const (
	DOMAINSTATUS_PENDING DomainStatus = "PENDING"
	DOMAINSTATUS_ACTIVE DomainStatus = "ACTIVE"
	DOMAINSTATUS_ERROR DomainStatus = "ERROR"
	DOMAINSTATUS_DELETING DomainStatus = "DELETING"
	DOMAINSTATUS_DELETED DomainStatus = "DELETED"
)

// All allowed values of DomainStatus enum
var AllowedDomainStatusEnumValues = []DomainStatus{
	"PENDING",
	"ACTIVE",
	"ERROR",
	"DELETING",
	"DELETED",
}

func (v *DomainStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DomainStatus(value)
	for _, existing := range AllowedDomainStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DomainStatus", value)
}

// NewDomainStatusFromValue returns a pointer to a valid DomainStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDomainStatusFromValue(v string) (*DomainStatus, error) {
	ev := DomainStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DomainStatus: valid values are %v", v, AllowedDomainStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DomainStatus) IsValid() bool {
	for _, existing := range AllowedDomainStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Domain.Status value
func (v DomainStatus) Ptr() *DomainStatus {
	return &v
}

type NullableDomainStatus struct {
	value *DomainStatus
	isSet bool
}

func (v NullableDomainStatus) Get() *DomainStatus {
	return v.value
}

func (v *NullableDomainStatus) Set(val *DomainStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainStatus(val *DomainStatus) *NullableDomainStatus {
	return &NullableDomainStatus{value: val, isSet: true}
}

func (v NullableDomainStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

