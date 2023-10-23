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

// checks if the ConfigCopyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigCopyResponse{}

// ConfigCopyResponse struct for ConfigCopyResponse
type ConfigCopyResponse struct {
	// The description of the resource that was requested to be copied.
	Description *string `json:"description,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// An Unique Id to identify copy operation.
	JobId *string `json:"job_id,omitempty"`
}

// NewConfigCopyResponse instantiates a new ConfigCopyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigCopyResponse() *ConfigCopyResponse {
	this := ConfigCopyResponse{}
	return &this
}

// NewConfigCopyResponseWithDefaults instantiates a new ConfigCopyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigCopyResponseWithDefaults() *ConfigCopyResponse {
	this := ConfigCopyResponse{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigCopyResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCopyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigCopyResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigCopyResponse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigCopyResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCopyResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigCopyResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigCopyResponse) SetId(v string) {
	o.Id = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *ConfigCopyResponse) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCopyResponse) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *ConfigCopyResponse) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *ConfigCopyResponse) SetJobId(v string) {
	o.JobId = &v
}

func (o ConfigCopyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigCopyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	return toSerialize, nil
}

type NullableConfigCopyResponse struct {
	value *ConfigCopyResponse
	isSet bool
}

func (v NullableConfigCopyResponse) Get() *ConfigCopyResponse {
	return v.value
}

func (v *NullableConfigCopyResponse) Set(val *ConfigCopyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigCopyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigCopyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigCopyResponse(val *ConfigCopyResponse) *NullableConfigCopyResponse {
	return &NullableConfigCopyResponse{value: val, isSet: true}
}

func (v NullableConfigCopyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigCopyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
