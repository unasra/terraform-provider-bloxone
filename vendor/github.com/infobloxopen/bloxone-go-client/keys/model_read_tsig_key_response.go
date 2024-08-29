/*
DDI Keys API

The DDI Keys application is a BloxOne DDI service for managing TSIG keys and GSS-TSIG (Kerberos) keys which are used by other BloxOne DDI applications. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keys

import (
	"encoding/json"
)

// checks if the ReadTSIGKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadTSIGKeyResponse{}

// ReadTSIGKeyResponse The response format to retrieve the __TSIGKey__ object.
type ReadTSIGKeyResponse struct {
	// The TSIGKey object.
	Result               *TSIGKey `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadTSIGKeyResponse ReadTSIGKeyResponse

// NewReadTSIGKeyResponse instantiates a new ReadTSIGKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadTSIGKeyResponse() *ReadTSIGKeyResponse {
	this := ReadTSIGKeyResponse{}
	return &this
}

// NewReadTSIGKeyResponseWithDefaults instantiates a new ReadTSIGKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadTSIGKeyResponseWithDefaults() *ReadTSIGKeyResponse {
	this := ReadTSIGKeyResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadTSIGKeyResponse) GetResult() TSIGKey {
	if o == nil || IsNil(o.Result) {
		var ret TSIGKey
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadTSIGKeyResponse) GetResultOk() (*TSIGKey, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadTSIGKeyResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given TSIGKey and assigns it to the Result field.
func (o *ReadTSIGKeyResponse) SetResult(v TSIGKey) {
	o.Result = &v
}

func (o ReadTSIGKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadTSIGKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReadTSIGKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varReadTSIGKeyResponse := _ReadTSIGKeyResponse{}

	err = json.Unmarshal(data, &varReadTSIGKeyResponse)

	if err != nil {
		return err
	}

	*o = ReadTSIGKeyResponse(varReadTSIGKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadTSIGKeyResponse struct {
	value *ReadTSIGKeyResponse
	isSet bool
}

func (v NullableReadTSIGKeyResponse) Get() *ReadTSIGKeyResponse {
	return v.value
}

func (v *NullableReadTSIGKeyResponse) Set(val *ReadTSIGKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadTSIGKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadTSIGKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadTSIGKeyResponse(val *ReadTSIGKeyResponse) *NullableReadTSIGKeyResponse {
	return &NullableReadTSIGKeyResponse{value: val, isSet: true}
}

func (v NullableReadTSIGKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadTSIGKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
