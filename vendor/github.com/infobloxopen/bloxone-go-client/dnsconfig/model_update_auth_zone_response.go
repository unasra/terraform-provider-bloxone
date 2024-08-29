/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
)

// checks if the UpdateAuthZoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAuthZoneResponse{}

// UpdateAuthZoneResponse The Authoritative Zone object update response format.
type UpdateAuthZoneResponse struct {
	// The updated AuthZone object.
	Result               *AuthZone `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateAuthZoneResponse UpdateAuthZoneResponse

// NewUpdateAuthZoneResponse instantiates a new UpdateAuthZoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthZoneResponse() *UpdateAuthZoneResponse {
	this := UpdateAuthZoneResponse{}
	return &this
}

// NewUpdateAuthZoneResponseWithDefaults instantiates a new UpdateAuthZoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthZoneResponseWithDefaults() *UpdateAuthZoneResponse {
	this := UpdateAuthZoneResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateAuthZoneResponse) GetResult() AuthZone {
	if o == nil || IsNil(o.Result) {
		var ret AuthZone
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthZoneResponse) GetResultOk() (*AuthZone, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateAuthZoneResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AuthZone and assigns it to the Result field.
func (o *UpdateAuthZoneResponse) SetResult(v AuthZone) {
	o.Result = &v
}

func (o UpdateAuthZoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAuthZoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateAuthZoneResponse) UnmarshalJSON(data []byte) (err error) {
	varUpdateAuthZoneResponse := _UpdateAuthZoneResponse{}

	err = json.Unmarshal(data, &varUpdateAuthZoneResponse)

	if err != nil {
		return err
	}

	*o = UpdateAuthZoneResponse(varUpdateAuthZoneResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateAuthZoneResponse struct {
	value *UpdateAuthZoneResponse
	isSet bool
}

func (v NullableUpdateAuthZoneResponse) Get() *UpdateAuthZoneResponse {
	return v.value
}

func (v *NullableUpdateAuthZoneResponse) Set(val *UpdateAuthZoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthZoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthZoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthZoneResponse(val *UpdateAuthZoneResponse) *NullableUpdateAuthZoneResponse {
	return &NullableUpdateAuthZoneResponse{value: val, isSet: true}
}

func (v NullableUpdateAuthZoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthZoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
