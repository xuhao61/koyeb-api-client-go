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

// PersistentVolumeStatus the model 'PersistentVolumeStatus'
type PersistentVolumeStatus string

// List of PersistentVolumeStatus
const (
	PERSISTENTVOLUMESTATUS_INVALID PersistentVolumeStatus = "PERSISTENT_VOLUME_STATUS_INVALID"
	PERSISTENTVOLUMESTATUS_ATTACHED PersistentVolumeStatus = "PERSISTENT_VOLUME_STATUS_ATTACHED"
	PERSISTENTVOLUMESTATUS_DETACHED PersistentVolumeStatus = "PERSISTENT_VOLUME_STATUS_DETACHED"
)

// All allowed values of PersistentVolumeStatus enum
var AllowedPersistentVolumeStatusEnumValues = []PersistentVolumeStatus{
	"PERSISTENT_VOLUME_STATUS_INVALID",
	"PERSISTENT_VOLUME_STATUS_ATTACHED",
	"PERSISTENT_VOLUME_STATUS_DETACHED",
}

func (v *PersistentVolumeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PersistentVolumeStatus(value)
	for _, existing := range AllowedPersistentVolumeStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PersistentVolumeStatus", value)
}

// NewPersistentVolumeStatusFromValue returns a pointer to a valid PersistentVolumeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPersistentVolumeStatusFromValue(v string) (*PersistentVolumeStatus, error) {
	ev := PersistentVolumeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PersistentVolumeStatus: valid values are %v", v, AllowedPersistentVolumeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PersistentVolumeStatus) IsValid() bool {
	for _, existing := range AllowedPersistentVolumeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersistentVolumeStatus value
func (v PersistentVolumeStatus) Ptr() *PersistentVolumeStatus {
	return &v
}

type NullablePersistentVolumeStatus struct {
	value *PersistentVolumeStatus
	isSet bool
}

func (v NullablePersistentVolumeStatus) Get() *PersistentVolumeStatus {
	return v.value
}

func (v *NullablePersistentVolumeStatus) Set(val *PersistentVolumeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePersistentVolumeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePersistentVolumeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersistentVolumeStatus(val *PersistentVolumeStatus) *NullablePersistentVolumeStatus {
	return &NullablePersistentVolumeStatus{value: val, isSet: true}
}

func (v NullablePersistentVolumeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersistentVolumeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

