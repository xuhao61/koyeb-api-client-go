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

// CatalogInstanceListItem struct for CatalogInstanceListItem
type CatalogInstanceListItem struct {
	Id *string `json:"id,omitempty"`
	Description *string `json:"description,omitempty"`
	Vcpu *int64 `json:"vcpu,omitempty"`
	Memory *string `json:"memory,omitempty"`
	Disk *string `json:"disk,omitempty"`
	PriceHourly *string `json:"price_hourly,omitempty"`
	PriceMonthly *string `json:"price_monthly,omitempty"`
	Regions []string `json:"regions,omitempty"`
	Status *string `json:"status,omitempty"`
	RequirePlan []string `json:"require_plan,omitempty"`
}

// NewCatalogInstanceListItem instantiates a new CatalogInstanceListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogInstanceListItem() *CatalogInstanceListItem {
	this := CatalogInstanceListItem{}
	return &this
}

// NewCatalogInstanceListItemWithDefaults instantiates a new CatalogInstanceListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogInstanceListItemWithDefaults() *CatalogInstanceListItem {
	this := CatalogInstanceListItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatalogInstanceListItem) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInstanceListItem) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatalogInstanceListItem) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CatalogInstanceListItem) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogInstanceListItem) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInstanceListItem) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogInstanceListItem) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogInstanceListItem) SetDescription(v string) {
	o.Description = &v
}

// GetVcpu returns the Vcpu field value if set, zero value otherwise.
func (o *CatalogInstanceListItem) GetVcpu() int64 {
	if o == nil || isNil(o.Vcpu) {
		var ret int64
		return ret
	}
	return *o.Vcpu
}

// GetVcpuOk returns a tuple with the Vcpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInstanceListItem) GetVcpuOk() (*int64, bool) {
	if o == nil || isNil(o.Vcpu) {
    return nil, false
	}
	return o.Vcpu, true
}

// HasVcpu returns a boolean if a field has been set.
func (o *CatalogInstanceListItem) HasVcpu() bool {
	if o != nil && !isNil(o.Vcpu) {
		return true
	}

	return false
}

// SetVcpu gets a reference to the given int64 and assigns it to the Vcpu field.
func (o *CatalogInstanceListItem) SetVcpu(v int64) {
	o.Vcpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *CatalogInstanceListItem) GetMemory() string {
	if o == nil || isNil(o.Memory) {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInstanceListItem) GetMemoryOk() (*string, bool) {
	if o == nil || isNil(o.Memory) {
    return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *CatalogInstanceListItem) HasMemory() bool {
	if o != nil && !isNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *CatalogInstanceListItem) SetMemory(v string) {
	o.Memory = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *CatalogInstanceListItem) GetDisk() string {
	if o == nil || isNil(o.Disk) {
		var ret string
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInstanceListItem) GetDiskOk() (*string, bool) {
	if o == nil || isNil(o.Disk) {
    return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *CatalogInstanceListItem) HasDisk() bool {
	if o != nil && !isNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given string and assigns it to the Disk field.
func (o *CatalogInstanceListItem) SetDisk(v string) {
	o.Disk = &v
}

// GetPriceHourly returns the PriceHourly field value if set, zero value otherwise.
func (o *CatalogInstanceListItem) GetPriceHourly() string {
	if o == nil || isNil(o.PriceHourly) {
		var ret string
		return ret
	}
	return *o.PriceHourly
}

// GetPriceHourlyOk returns a tuple with the PriceHourly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInstanceListItem) GetPriceHourlyOk() (*string, bool) {
	if o == nil || isNil(o.PriceHourly) {
    return nil, false
	}
	return o.PriceHourly, true
}

// HasPriceHourly returns a boolean if a field has been set.
func (o *CatalogInstanceListItem) HasPriceHourly() bool {
	if o != nil && !isNil(o.PriceHourly) {
		return true
	}

	return false
}

// SetPriceHourly gets a reference to the given string and assigns it to the PriceHourly field.
func (o *CatalogInstanceListItem) SetPriceHourly(v string) {
	o.PriceHourly = &v
}

// GetPriceMonthly returns the PriceMonthly field value if set, zero value otherwise.
func (o *CatalogInstanceListItem) GetPriceMonthly() string {
	if o == nil || isNil(o.PriceMonthly) {
		var ret string
		return ret
	}
	return *o.PriceMonthly
}

// GetPriceMonthlyOk returns a tuple with the PriceMonthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInstanceListItem) GetPriceMonthlyOk() (*string, bool) {
	if o == nil || isNil(o.PriceMonthly) {
    return nil, false
	}
	return o.PriceMonthly, true
}

// HasPriceMonthly returns a boolean if a field has been set.
func (o *CatalogInstanceListItem) HasPriceMonthly() bool {
	if o != nil && !isNil(o.PriceMonthly) {
		return true
	}

	return false
}

// SetPriceMonthly gets a reference to the given string and assigns it to the PriceMonthly field.
func (o *CatalogInstanceListItem) SetPriceMonthly(v string) {
	o.PriceMonthly = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *CatalogInstanceListItem) GetRegions() []string {
	if o == nil || isNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInstanceListItem) GetRegionsOk() ([]string, bool) {
	if o == nil || isNil(o.Regions) {
    return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *CatalogInstanceListItem) HasRegions() bool {
	if o != nil && !isNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *CatalogInstanceListItem) SetRegions(v []string) {
	o.Regions = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CatalogInstanceListItem) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInstanceListItem) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CatalogInstanceListItem) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CatalogInstanceListItem) SetStatus(v string) {
	o.Status = &v
}

// GetRequirePlan returns the RequirePlan field value if set, zero value otherwise.
func (o *CatalogInstanceListItem) GetRequirePlan() []string {
	if o == nil || isNil(o.RequirePlan) {
		var ret []string
		return ret
	}
	return o.RequirePlan
}

// GetRequirePlanOk returns a tuple with the RequirePlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInstanceListItem) GetRequirePlanOk() ([]string, bool) {
	if o == nil || isNil(o.RequirePlan) {
    return nil, false
	}
	return o.RequirePlan, true
}

// HasRequirePlan returns a boolean if a field has been set.
func (o *CatalogInstanceListItem) HasRequirePlan() bool {
	if o != nil && !isNil(o.RequirePlan) {
		return true
	}

	return false
}

// SetRequirePlan gets a reference to the given []string and assigns it to the RequirePlan field.
func (o *CatalogInstanceListItem) SetRequirePlan(v []string) {
	o.RequirePlan = v
}

func (o CatalogInstanceListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Vcpu) {
		toSerialize["vcpu"] = o.Vcpu
	}
	if !isNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !isNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !isNil(o.PriceHourly) {
		toSerialize["price_hourly"] = o.PriceHourly
	}
	if !isNil(o.PriceMonthly) {
		toSerialize["price_monthly"] = o.PriceMonthly
	}
	if !isNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.RequirePlan) {
		toSerialize["require_plan"] = o.RequirePlan
	}
	return json.Marshal(toSerialize)
}

type NullableCatalogInstanceListItem struct {
	value *CatalogInstanceListItem
	isSet bool
}

func (v NullableCatalogInstanceListItem) Get() *CatalogInstanceListItem {
	return v.value
}

func (v *NullableCatalogInstanceListItem) Set(val *CatalogInstanceListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogInstanceListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogInstanceListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogInstanceListItem(val *CatalogInstanceListItem) *NullableCatalogInstanceListItem {
	return &NullableCatalogInstanceListItem{value: val, isSet: true}
}

func (v NullableCatalogInstanceListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogInstanceListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


