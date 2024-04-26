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

// checks if the CreateTSIGKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTSIGKeyResponse{}

// CreateTSIGKeyResponse The response format to create a __TSIGKey__ object.
type CreateTSIGKeyResponse struct {
	Result *TSIGKey `json:"result,omitempty"`
}

// NewCreateTSIGKeyResponse instantiates a new CreateTSIGKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTSIGKeyResponse() *CreateTSIGKeyResponse {
	this := CreateTSIGKeyResponse{}
	return &this
}

// NewCreateTSIGKeyResponseWithDefaults instantiates a new CreateTSIGKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTSIGKeyResponseWithDefaults() *CreateTSIGKeyResponse {
	this := CreateTSIGKeyResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateTSIGKeyResponse) GetResult() TSIGKey {
	if o == nil || IsNil(o.Result) {
		var ret TSIGKey
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTSIGKeyResponse) GetResultOk() (*TSIGKey, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateTSIGKeyResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given TSIGKey and assigns it to the Result field.
func (o *CreateTSIGKeyResponse) SetResult(v TSIGKey) {
	o.Result = &v
}

func (o CreateTSIGKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTSIGKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCreateTSIGKeyResponse struct {
	value *CreateTSIGKeyResponse
	isSet bool
}

func (v NullableCreateTSIGKeyResponse) Get() *CreateTSIGKeyResponse {
	return v.value
}

func (v *NullableCreateTSIGKeyResponse) Set(val *CreateTSIGKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTSIGKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTSIGKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTSIGKeyResponse(val *CreateTSIGKeyResponse) *NullableCreateTSIGKeyResponse {
	return &NullableCreateTSIGKeyResponse{value: val, isSet: true}
}

func (v NullableCreateTSIGKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTSIGKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
