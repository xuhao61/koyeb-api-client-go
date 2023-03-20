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

// DeploymentDefinition struct for DeploymentDefinition
type DeploymentDefinition struct {
	Name *string `json:"name,omitempty"`
	Type *DeploymentDefinitionType `json:"type,omitempty"`
	Routes []DeploymentRoute `json:"routes,omitempty"`
	Ports []DeploymentPort `json:"ports,omitempty"`
	Env []DeploymentEnv `json:"env,omitempty"`
	Regions []string `json:"regions,omitempty"`
	Scalings []DeploymentScaling `json:"scalings,omitempty"`
	InstanceTypes []DeploymentInstanceType `json:"instance_types,omitempty"`
	HealthChecks []DeploymentHealthCheck `json:"health_checks,omitempty"`
	Docker *DockerSource `json:"docker,omitempty"`
	Git *GitSource `json:"git,omitempty"`
}

// NewDeploymentDefinition instantiates a new DeploymentDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentDefinition() *DeploymentDefinition {
	this := DeploymentDefinition{}
	var type_ DeploymentDefinitionType = DEPLOYMENTDEFINITIONTYPE_INVALID
	this.Type = &type_
	return &this
}

// NewDeploymentDefinitionWithDefaults instantiates a new DeploymentDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentDefinitionWithDefaults() *DeploymentDefinition {
	this := DeploymentDefinition{}
	var type_ DeploymentDefinitionType = DEPLOYMENTDEFINITIONTYPE_INVALID
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentDefinition) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetType() DeploymentDefinitionType {
	if o == nil || isNil(o.Type) {
		var ret DeploymentDefinitionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetTypeOk() (*DeploymentDefinitionType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DeploymentDefinitionType and assigns it to the Type field.
func (o *DeploymentDefinition) SetType(v DeploymentDefinitionType) {
	o.Type = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetRoutes() []DeploymentRoute {
	if o == nil || isNil(o.Routes) {
		var ret []DeploymentRoute
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetRoutesOk() ([]DeploymentRoute, bool) {
	if o == nil || isNil(o.Routes) {
    return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasRoutes() bool {
	if o != nil && !isNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []DeploymentRoute and assigns it to the Routes field.
func (o *DeploymentDefinition) SetRoutes(v []DeploymentRoute) {
	o.Routes = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetPorts() []DeploymentPort {
	if o == nil || isNil(o.Ports) {
		var ret []DeploymentPort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetPortsOk() ([]DeploymentPort, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []DeploymentPort and assigns it to the Ports field.
func (o *DeploymentDefinition) SetPorts(v []DeploymentPort) {
	o.Ports = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetEnv() []DeploymentEnv {
	if o == nil || isNil(o.Env) {
		var ret []DeploymentEnv
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetEnvOk() ([]DeploymentEnv, bool) {
	if o == nil || isNil(o.Env) {
    return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasEnv() bool {
	if o != nil && !isNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []DeploymentEnv and assigns it to the Env field.
func (o *DeploymentDefinition) SetEnv(v []DeploymentEnv) {
	o.Env = v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetRegions() []string {
	if o == nil || isNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetRegionsOk() ([]string, bool) {
	if o == nil || isNil(o.Regions) {
    return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasRegions() bool {
	if o != nil && !isNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *DeploymentDefinition) SetRegions(v []string) {
	o.Regions = v
}

// GetScalings returns the Scalings field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetScalings() []DeploymentScaling {
	if o == nil || isNil(o.Scalings) {
		var ret []DeploymentScaling
		return ret
	}
	return o.Scalings
}

// GetScalingsOk returns a tuple with the Scalings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetScalingsOk() ([]DeploymentScaling, bool) {
	if o == nil || isNil(o.Scalings) {
    return nil, false
	}
	return o.Scalings, true
}

// HasScalings returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasScalings() bool {
	if o != nil && !isNil(o.Scalings) {
		return true
	}

	return false
}

// SetScalings gets a reference to the given []DeploymentScaling and assigns it to the Scalings field.
func (o *DeploymentDefinition) SetScalings(v []DeploymentScaling) {
	o.Scalings = v
}

// GetInstanceTypes returns the InstanceTypes field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetInstanceTypes() []DeploymentInstanceType {
	if o == nil || isNil(o.InstanceTypes) {
		var ret []DeploymentInstanceType
		return ret
	}
	return o.InstanceTypes
}

// GetInstanceTypesOk returns a tuple with the InstanceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetInstanceTypesOk() ([]DeploymentInstanceType, bool) {
	if o == nil || isNil(o.InstanceTypes) {
    return nil, false
	}
	return o.InstanceTypes, true
}

// HasInstanceTypes returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasInstanceTypes() bool {
	if o != nil && !isNil(o.InstanceTypes) {
		return true
	}

	return false
}

// SetInstanceTypes gets a reference to the given []DeploymentInstanceType and assigns it to the InstanceTypes field.
func (o *DeploymentDefinition) SetInstanceTypes(v []DeploymentInstanceType) {
	o.InstanceTypes = v
}

// GetHealthChecks returns the HealthChecks field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetHealthChecks() []DeploymentHealthCheck {
	if o == nil || isNil(o.HealthChecks) {
		var ret []DeploymentHealthCheck
		return ret
	}
	return o.HealthChecks
}

// GetHealthChecksOk returns a tuple with the HealthChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetHealthChecksOk() ([]DeploymentHealthCheck, bool) {
	if o == nil || isNil(o.HealthChecks) {
    return nil, false
	}
	return o.HealthChecks, true
}

// HasHealthChecks returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasHealthChecks() bool {
	if o != nil && !isNil(o.HealthChecks) {
		return true
	}

	return false
}

// SetHealthChecks gets a reference to the given []DeploymentHealthCheck and assigns it to the HealthChecks field.
func (o *DeploymentDefinition) SetHealthChecks(v []DeploymentHealthCheck) {
	o.HealthChecks = v
}

// GetDocker returns the Docker field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetDocker() DockerSource {
	if o == nil || isNil(o.Docker) {
		var ret DockerSource
		return ret
	}
	return *o.Docker
}

// GetDockerOk returns a tuple with the Docker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetDockerOk() (*DockerSource, bool) {
	if o == nil || isNil(o.Docker) {
    return nil, false
	}
	return o.Docker, true
}

// HasDocker returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasDocker() bool {
	if o != nil && !isNil(o.Docker) {
		return true
	}

	return false
}

// SetDocker gets a reference to the given DockerSource and assigns it to the Docker field.
func (o *DeploymentDefinition) SetDocker(v DockerSource) {
	o.Docker = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetGit() GitSource {
	if o == nil || isNil(o.Git) {
		var ret GitSource
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetGitOk() (*GitSource, bool) {
	if o == nil || isNil(o.Git) {
    return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasGit() bool {
	if o != nil && !isNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given GitSource and assigns it to the Git field.
func (o *DeploymentDefinition) SetGit(v GitSource) {
	o.Git = &v
}

func (o DeploymentDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !isNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !isNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !isNil(o.Scalings) {
		toSerialize["scalings"] = o.Scalings
	}
	if !isNil(o.InstanceTypes) {
		toSerialize["instance_types"] = o.InstanceTypes
	}
	if !isNil(o.HealthChecks) {
		toSerialize["health_checks"] = o.HealthChecks
	}
	if !isNil(o.Docker) {
		toSerialize["docker"] = o.Docker
	}
	if !isNil(o.Git) {
		toSerialize["git"] = o.Git
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentDefinition struct {
	value *DeploymentDefinition
	isSet bool
}

func (v NullableDeploymentDefinition) Get() *DeploymentDefinition {
	return v.value
}

func (v *NullableDeploymentDefinition) Set(val *DeploymentDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentDefinition(val *DeploymentDefinition) *NullableDeploymentDefinition {
	return &NullableDeploymentDefinition{value: val, isSet: true}
}

func (v NullableDeploymentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


