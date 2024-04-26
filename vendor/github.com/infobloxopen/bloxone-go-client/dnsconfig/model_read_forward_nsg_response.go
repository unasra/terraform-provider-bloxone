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

// checks if the ReadForwardNSGResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadForwardNSGResponse{}

// ReadForwardNSGResponse The ForwardNSG object read response format.
type ReadForwardNSGResponse struct {
	Result               *ForwardNSG `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadForwardNSGResponse ReadForwardNSGResponse

// NewReadForwardNSGResponse instantiates a new ReadForwardNSGResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadForwardNSGResponse() *ReadForwardNSGResponse {
	this := ReadForwardNSGResponse{}
	return &this
}

// NewReadForwardNSGResponseWithDefaults instantiates a new ReadForwardNSGResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadForwardNSGResponseWithDefaults() *ReadForwardNSGResponse {
	this := ReadForwardNSGResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadForwardNSGResponse) GetResult() ForwardNSG {
	if o == nil || IsNil(o.Result) {
		var ret ForwardNSG
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadForwardNSGResponse) GetResultOk() (*ForwardNSG, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadForwardNSGResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ForwardNSG and assigns it to the Result field.
func (o *ReadForwardNSGResponse) SetResult(v ForwardNSG) {
	o.Result = &v
}

func (o ReadForwardNSGResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadForwardNSGResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReadForwardNSGResponse) UnmarshalJSON(data []byte) (err error) {
	varReadForwardNSGResponse := _ReadForwardNSGResponse{}

	err = json.Unmarshal(data, &varReadForwardNSGResponse)

	if err != nil {
		return err
	}

	*o = ReadForwardNSGResponse(varReadForwardNSGResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadForwardNSGResponse struct {
	value *ReadForwardNSGResponse
	isSet bool
}

func (v NullableReadForwardNSGResponse) Get() *ReadForwardNSGResponse {
	return v.value
}

func (v *NullableReadForwardNSGResponse) Set(val *ReadForwardNSGResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadForwardNSGResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadForwardNSGResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadForwardNSGResponse(val *ReadForwardNSGResponse) *NullableReadForwardNSGResponse {
	return &NullableReadForwardNSGResponse{value: val, isSet: true}
}

func (v NullableReadForwardNSGResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadForwardNSGResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
