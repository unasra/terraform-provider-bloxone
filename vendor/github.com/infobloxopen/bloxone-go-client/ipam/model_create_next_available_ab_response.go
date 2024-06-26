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

// checks if the CreateNextAvailableABResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNextAvailableABResponse{}

// CreateNextAvailableABResponse The Next Available Address Block object create response format.
type CreateNextAvailableABResponse struct {
	// The list of Next Available Address Block objects.
	Results              []AddressBlock `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateNextAvailableABResponse CreateNextAvailableABResponse

// NewCreateNextAvailableABResponse instantiates a new CreateNextAvailableABResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNextAvailableABResponse() *CreateNextAvailableABResponse {
	this := CreateNextAvailableABResponse{}
	return &this
}

// NewCreateNextAvailableABResponseWithDefaults instantiates a new CreateNextAvailableABResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNextAvailableABResponseWithDefaults() *CreateNextAvailableABResponse {
	this := CreateNextAvailableABResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CreateNextAvailableABResponse) GetResults() []AddressBlock {
	if o == nil || IsNil(o.Results) {
		var ret []AddressBlock
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNextAvailableABResponse) GetResultsOk() ([]AddressBlock, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CreateNextAvailableABResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []AddressBlock and assigns it to the Results field.
func (o *CreateNextAvailableABResponse) SetResults(v []AddressBlock) {
	o.Results = v
}

func (o CreateNextAvailableABResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNextAvailableABResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateNextAvailableABResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateNextAvailableABResponse := _CreateNextAvailableABResponse{}

	err = json.Unmarshal(data, &varCreateNextAvailableABResponse)

	if err != nil {
		return err
	}

	*o = CreateNextAvailableABResponse(varCreateNextAvailableABResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateNextAvailableABResponse struct {
	value *CreateNextAvailableABResponse
	isSet bool
}

func (v NullableCreateNextAvailableABResponse) Get() *CreateNextAvailableABResponse {
	return v.value
}

func (v *NullableCreateNextAvailableABResponse) Set(val *CreateNextAvailableABResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNextAvailableABResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNextAvailableABResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNextAvailableABResponse(val *CreateNextAvailableABResponse) *NullableCreateNextAvailableABResponse {
	return &NullableCreateNextAvailableABResponse{value: val, isSet: true}
}

func (v NullableCreateNextAvailableABResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNextAvailableABResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
