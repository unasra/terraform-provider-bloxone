/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the ReadHAGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadHAGroupResponse{}

// ReadHAGroupResponse The response format to retrieve the __HAGroup__ object.
type ReadHAGroupResponse struct {
	// The HAGroup object.
	Result               *HAGroup `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadHAGroupResponse ReadHAGroupResponse

// NewReadHAGroupResponse instantiates a new ReadHAGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadHAGroupResponse() *ReadHAGroupResponse {
	this := ReadHAGroupResponse{}
	return &this
}

// NewReadHAGroupResponseWithDefaults instantiates a new ReadHAGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadHAGroupResponseWithDefaults() *ReadHAGroupResponse {
	this := ReadHAGroupResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadHAGroupResponse) GetResult() HAGroup {
	if o == nil || IsNil(o.Result) {
		var ret HAGroup
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadHAGroupResponse) GetResultOk() (*HAGroup, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadHAGroupResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given HAGroup and assigns it to the Result field.
func (o *ReadHAGroupResponse) SetResult(v HAGroup) {
	o.Result = &v
}

func (o ReadHAGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadHAGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReadHAGroupResponse) UnmarshalJSON(data []byte) (err error) {
	varReadHAGroupResponse := _ReadHAGroupResponse{}

	err = json.Unmarshal(data, &varReadHAGroupResponse)

	if err != nil {
		return err
	}

	*o = ReadHAGroupResponse(varReadHAGroupResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadHAGroupResponse struct {
	value *ReadHAGroupResponse
	isSet bool
}

func (v NullableReadHAGroupResponse) Get() *ReadHAGroupResponse {
	return v.value
}

func (v *NullableReadHAGroupResponse) Set(val *ReadHAGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadHAGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadHAGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadHAGroupResponse(val *ReadHAGroupResponse) *NullableReadHAGroupResponse {
	return &NullableReadHAGroupResponse{value: val, isSet: true}
}

func (v NullableReadHAGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadHAGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
