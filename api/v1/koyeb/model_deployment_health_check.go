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

// DeploymentHealthCheck struct for DeploymentHealthCheck
type DeploymentHealthCheck struct {
	GracePeriod *int64 `json:"grace_period,omitempty"`
	Interval *int64 `json:"interval,omitempty"`
	RestartLimit *int64 `json:"restart_limit,omitempty"`
	Timeout *int64 `json:"timeout,omitempty"`
	Tcp *TCPHealthCheck `json:"tcp,omitempty"`
	Http *HTTPHealthCheck `json:"http,omitempty"`
}

// NewDeploymentHealthCheck instantiates a new DeploymentHealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHealthCheck() *DeploymentHealthCheck {
	this := DeploymentHealthCheck{}
	return &this
}

// NewDeploymentHealthCheckWithDefaults instantiates a new DeploymentHealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHealthCheckWithDefaults() *DeploymentHealthCheck {
	this := DeploymentHealthCheck{}
	return &this
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetGracePeriod() int64 {
	if o == nil || isNil(o.GracePeriod) {
		var ret int64
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetGracePeriodOk() (*int64, bool) {
	if o == nil || isNil(o.GracePeriod) {
    return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasGracePeriod() bool {
	if o != nil && !isNil(o.GracePeriod) {
		return true
	}

	return false
}

// SetGracePeriod gets a reference to the given int64 and assigns it to the GracePeriod field.
func (o *DeploymentHealthCheck) SetGracePeriod(v int64) {
	o.GracePeriod = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetInterval() int64 {
	if o == nil || isNil(o.Interval) {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetIntervalOk() (*int64, bool) {
	if o == nil || isNil(o.Interval) {
    return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasInterval() bool {
	if o != nil && !isNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *DeploymentHealthCheck) SetInterval(v int64) {
	o.Interval = &v
}

// GetRestartLimit returns the RestartLimit field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetRestartLimit() int64 {
	if o == nil || isNil(o.RestartLimit) {
		var ret int64
		return ret
	}
	return *o.RestartLimit
}

// GetRestartLimitOk returns a tuple with the RestartLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetRestartLimitOk() (*int64, bool) {
	if o == nil || isNil(o.RestartLimit) {
    return nil, false
	}
	return o.RestartLimit, true
}

// HasRestartLimit returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasRestartLimit() bool {
	if o != nil && !isNil(o.RestartLimit) {
		return true
	}

	return false
}

// SetRestartLimit gets a reference to the given int64 and assigns it to the RestartLimit field.
func (o *DeploymentHealthCheck) SetRestartLimit(v int64) {
	o.RestartLimit = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetTimeout() int64 {
	if o == nil || isNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetTimeoutOk() (*int64, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *DeploymentHealthCheck) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetTcp returns the Tcp field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetTcp() TCPHealthCheck {
	if o == nil || isNil(o.Tcp) {
		var ret TCPHealthCheck
		return ret
	}
	return *o.Tcp
}

// GetTcpOk returns a tuple with the Tcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetTcpOk() (*TCPHealthCheck, bool) {
	if o == nil || isNil(o.Tcp) {
    return nil, false
	}
	return o.Tcp, true
}

// HasTcp returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasTcp() bool {
	if o != nil && !isNil(o.Tcp) {
		return true
	}

	return false
}

// SetTcp gets a reference to the given TCPHealthCheck and assigns it to the Tcp field.
func (o *DeploymentHealthCheck) SetTcp(v TCPHealthCheck) {
	o.Tcp = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetHttp() HTTPHealthCheck {
	if o == nil || isNil(o.Http) {
		var ret HTTPHealthCheck
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetHttpOk() (*HTTPHealthCheck, bool) {
	if o == nil || isNil(o.Http) {
    return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasHttp() bool {
	if o != nil && !isNil(o.Http) {
		return true
	}

	return false
}

// SetHttp gets a reference to the given HTTPHealthCheck and assigns it to the Http field.
func (o *DeploymentHealthCheck) SetHttp(v HTTPHealthCheck) {
	o.Http = &v
}

func (o DeploymentHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GracePeriod) {
		toSerialize["grace_period"] = o.GracePeriod
	}
	if !isNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !isNil(o.RestartLimit) {
		toSerialize["restart_limit"] = o.RestartLimit
	}
	if !isNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !isNil(o.Tcp) {
		toSerialize["tcp"] = o.Tcp
	}
	if !isNil(o.Http) {
		toSerialize["http"] = o.Http
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHealthCheck struct {
	value *DeploymentHealthCheck
	isSet bool
}

func (v NullableDeploymentHealthCheck) Get() *DeploymentHealthCheck {
	return v.value
}

func (v *NullableDeploymentHealthCheck) Set(val *DeploymentHealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHealthCheck(val *DeploymentHealthCheck) *NullableDeploymentHealthCheck {
	return &NullableDeploymentHealthCheck{value: val, isSet: true}
}

func (v NullableDeploymentHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


