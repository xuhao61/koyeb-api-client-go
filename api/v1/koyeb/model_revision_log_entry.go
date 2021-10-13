/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// RevisionLogEntry struct for RevisionLogEntry
type RevisionLogEntry struct {
	RevisionId *string `json:"revision_id,omitempty"`
	InstanceId *string `json:"instance_id,omitempty"`
	Stream *string `json:"stream,omitempty"`
	Msg *string `json:"msg,omitempty"`
	Offset *string `json:"offset,omitempty"`
}

// NewRevisionLogEntry instantiates a new RevisionLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionLogEntry() *RevisionLogEntry {
	this := RevisionLogEntry{}
	return &this
}

// NewRevisionLogEntryWithDefaults instantiates a new RevisionLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionLogEntryWithDefaults() *RevisionLogEntry {
	this := RevisionLogEntry{}
	return &this
}

// GetRevisionId returns the RevisionId field value if set, zero value otherwise.
func (o *RevisionLogEntry) GetRevisionId() string {
	if o == nil || o.RevisionId == nil {
		var ret string
		return ret
	}
	return *o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionLogEntry) GetRevisionIdOk() (*string, bool) {
	if o == nil || o.RevisionId == nil {
		return nil, false
	}
	return o.RevisionId, true
}

// HasRevisionId returns a boolean if a field has been set.
func (o *RevisionLogEntry) HasRevisionId() bool {
	if o != nil && o.RevisionId != nil {
		return true
	}

	return false
}

// SetRevisionId gets a reference to the given string and assigns it to the RevisionId field.
func (o *RevisionLogEntry) SetRevisionId(v string) {
	o.RevisionId = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *RevisionLogEntry) GetInstanceId() string {
	if o == nil || o.InstanceId == nil {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionLogEntry) GetInstanceIdOk() (*string, bool) {
	if o == nil || o.InstanceId == nil {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *RevisionLogEntry) HasInstanceId() bool {
	if o != nil && o.InstanceId != nil {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *RevisionLogEntry) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *RevisionLogEntry) GetStream() string {
	if o == nil || o.Stream == nil {
		var ret string
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionLogEntry) GetStreamOk() (*string, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *RevisionLogEntry) HasStream() bool {
	if o != nil && o.Stream != nil {
		return true
	}

	return false
}

// SetStream gets a reference to the given string and assigns it to the Stream field.
func (o *RevisionLogEntry) SetStream(v string) {
	o.Stream = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *RevisionLogEntry) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionLogEntry) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *RevisionLogEntry) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *RevisionLogEntry) SetMsg(v string) {
	o.Msg = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *RevisionLogEntry) GetOffset() string {
	if o == nil || o.Offset == nil {
		var ret string
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionLogEntry) GetOffsetOk() (*string, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *RevisionLogEntry) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given string and assigns it to the Offset field.
func (o *RevisionLogEntry) SetOffset(v string) {
	o.Offset = &v
}

func (o RevisionLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RevisionId != nil {
		toSerialize["revision_id"] = o.RevisionId
	}
	if o.InstanceId != nil {
		toSerialize["instance_id"] = o.InstanceId
	}
	if o.Stream != nil {
		toSerialize["stream"] = o.Stream
	}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	return json.Marshal(toSerialize)
}

type NullableRevisionLogEntry struct {
	value *RevisionLogEntry
	isSet bool
}

func (v NullableRevisionLogEntry) Get() *RevisionLogEntry {
	return v.value
}

func (v *NullableRevisionLogEntry) Set(val *RevisionLogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionLogEntry(val *RevisionLogEntry) *NullableRevisionLogEntry {
	return &NullableRevisionLogEntry{value: val, isSet: true}
}

func (v NullableRevisionLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


