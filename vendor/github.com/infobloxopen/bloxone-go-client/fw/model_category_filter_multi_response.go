/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
)

// checks if the CategoryFilterMultiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryFilterMultiResponse{}

// CategoryFilterMultiResponse The Category Filter list response.
type CategoryFilterMultiResponse struct {
	// The list of Category Filter objects.
	Results              []CategoryFilter `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CategoryFilterMultiResponse CategoryFilterMultiResponse

// NewCategoryFilterMultiResponse instantiates a new CategoryFilterMultiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryFilterMultiResponse() *CategoryFilterMultiResponse {
	this := CategoryFilterMultiResponse{}
	return &this
}

// NewCategoryFilterMultiResponseWithDefaults instantiates a new CategoryFilterMultiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryFilterMultiResponseWithDefaults() *CategoryFilterMultiResponse {
	this := CategoryFilterMultiResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CategoryFilterMultiResponse) GetResults() []CategoryFilter {
	if o == nil || IsNil(o.Results) {
		var ret []CategoryFilter
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryFilterMultiResponse) GetResultsOk() ([]CategoryFilter, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CategoryFilterMultiResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CategoryFilter and assigns it to the Results field.
func (o *CategoryFilterMultiResponse) SetResults(v []CategoryFilter) {
	o.Results = v
}

func (o CategoryFilterMultiResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryFilterMultiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CategoryFilterMultiResponse) UnmarshalJSON(data []byte) (err error) {
	varCategoryFilterMultiResponse := _CategoryFilterMultiResponse{}

	err = json.Unmarshal(data, &varCategoryFilterMultiResponse)

	if err != nil {
		return err
	}

	*o = CategoryFilterMultiResponse(varCategoryFilterMultiResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCategoryFilterMultiResponse struct {
	value *CategoryFilterMultiResponse
	isSet bool
}

func (v NullableCategoryFilterMultiResponse) Get() *CategoryFilterMultiResponse {
	return v.value
}

func (v *NullableCategoryFilterMultiResponse) Set(val *CategoryFilterMultiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryFilterMultiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryFilterMultiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryFilterMultiResponse(val *CategoryFilterMultiResponse) *NullableCategoryFilterMultiResponse {
	return &NullableCategoryFilterMultiResponse{value: val, isSet: true}
}

func (v NullableCategoryFilterMultiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryFilterMultiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
