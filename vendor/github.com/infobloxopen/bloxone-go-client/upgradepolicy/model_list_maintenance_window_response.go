/*
Schedule Software/Config Updates

Infoblox by default does automatic software updates when they become available. Updates are applied to all on-prem hosts, physical or virtual. However, you can override and schedule the software updates. You can also defer the updates to a later date and time. You can configure up to a total of 50 deferrals (scheduled and deferred software updates), which means you have the flexibility to create up to 50 update groups across different on-prem hosts by mapping with appropriate tags. Tags are be used to associate deferrals (scheduled or deferred) with a specific or group of onprem-hosts. Apart from software update deferrals, config update deferrals also can be configured using these overrides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package upgradepolicy

import (
	"encoding/json"
)

// checks if the ListMaintenanceWindowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMaintenanceWindowResponse{}

// ListMaintenanceWindowResponse struct for ListMaintenanceWindowResponse
type ListMaintenanceWindowResponse struct {
	Result               []MaintenanceWindow `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListMaintenanceWindowResponse ListMaintenanceWindowResponse

// NewListMaintenanceWindowResponse instantiates a new ListMaintenanceWindowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMaintenanceWindowResponse() *ListMaintenanceWindowResponse {
	this := ListMaintenanceWindowResponse{}
	return &this
}

// NewListMaintenanceWindowResponseWithDefaults instantiates a new ListMaintenanceWindowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMaintenanceWindowResponseWithDefaults() *ListMaintenanceWindowResponse {
	this := ListMaintenanceWindowResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListMaintenanceWindowResponse) GetResult() []MaintenanceWindow {
	if o == nil || IsNil(o.Result) {
		var ret []MaintenanceWindow
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMaintenanceWindowResponse) GetResultOk() ([]MaintenanceWindow, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListMaintenanceWindowResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []MaintenanceWindow and assigns it to the Result field.
func (o *ListMaintenanceWindowResponse) SetResult(v []MaintenanceWindow) {
	o.Result = v
}

func (o ListMaintenanceWindowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMaintenanceWindowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListMaintenanceWindowResponse) UnmarshalJSON(data []byte) (err error) {
	varListMaintenanceWindowResponse := _ListMaintenanceWindowResponse{}

	err = json.Unmarshal(data, &varListMaintenanceWindowResponse)

	if err != nil {
		return err
	}

	*o = ListMaintenanceWindowResponse(varListMaintenanceWindowResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListMaintenanceWindowResponse struct {
	value *ListMaintenanceWindowResponse
	isSet bool
}

func (v NullableListMaintenanceWindowResponse) Get() *ListMaintenanceWindowResponse {
	return v.value
}

func (v *NullableListMaintenanceWindowResponse) Set(val *ListMaintenanceWindowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMaintenanceWindowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMaintenanceWindowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMaintenanceWindowResponse(val *ListMaintenanceWindowResponse) *NullableListMaintenanceWindowResponse {
	return &NullableListMaintenanceWindowResponse{value: val, isSet: true}
}

func (v NullableListMaintenanceWindowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMaintenanceWindowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
