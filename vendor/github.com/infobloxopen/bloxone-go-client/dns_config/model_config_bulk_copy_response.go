/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
)

// checks if the ConfigBulkCopyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigBulkCopyResponse{}

// ConfigBulkCopyResponse struct for ConfigBulkCopyResponse
type ConfigBulkCopyResponse struct {
	Errors  []ConfigBulkCopyError `json:"errors,omitempty"`
	Results []ConfigCopyResponse  `json:"results,omitempty"`
}

// NewConfigBulkCopyResponse instantiates a new ConfigBulkCopyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigBulkCopyResponse() *ConfigBulkCopyResponse {
	this := ConfigBulkCopyResponse{}
	return &this
}

// NewConfigBulkCopyResponseWithDefaults instantiates a new ConfigBulkCopyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigBulkCopyResponseWithDefaults() *ConfigBulkCopyResponse {
	this := ConfigBulkCopyResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ConfigBulkCopyResponse) GetErrors() []ConfigBulkCopyError {
	if o == nil || IsNil(o.Errors) {
		var ret []ConfigBulkCopyError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigBulkCopyResponse) GetErrorsOk() ([]ConfigBulkCopyError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ConfigBulkCopyResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ConfigBulkCopyError and assigns it to the Errors field.
func (o *ConfigBulkCopyResponse) SetErrors(v []ConfigBulkCopyError) {
	o.Errors = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ConfigBulkCopyResponse) GetResults() []ConfigCopyResponse {
	if o == nil || IsNil(o.Results) {
		var ret []ConfigCopyResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigBulkCopyResponse) GetResultsOk() ([]ConfigCopyResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ConfigBulkCopyResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ConfigCopyResponse and assigns it to the Results field.
func (o *ConfigBulkCopyResponse) SetResults(v []ConfigCopyResponse) {
	o.Results = v
}

func (o ConfigBulkCopyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigBulkCopyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableConfigBulkCopyResponse struct {
	value *ConfigBulkCopyResponse
	isSet bool
}

func (v NullableConfigBulkCopyResponse) Get() *ConfigBulkCopyResponse {
	return v.value
}

func (v *NullableConfigBulkCopyResponse) Set(val *ConfigBulkCopyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigBulkCopyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigBulkCopyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigBulkCopyResponse(val *ConfigBulkCopyResponse) *NullableConfigBulkCopyResponse {
	return &NullableConfigBulkCopyResponse{value: val, isSet: true}
}

func (v NullableConfigBulkCopyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigBulkCopyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
