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

// checks if the CategoryFiltersCreateCategoryFilter400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryFiltersCreateCategoryFilter400Response{}

// CategoryFiltersCreateCategoryFilter400Response struct for CategoryFiltersCreateCategoryFilter400Response
type CategoryFiltersCreateCategoryFilter400Response struct {
	Error *CategoryFiltersCreateCategoryFilter400ResponseError `json:"error,omitempty"`
}

// NewCategoryFiltersCreateCategoryFilter400Response instantiates a new CategoryFiltersCreateCategoryFilter400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryFiltersCreateCategoryFilter400Response() *CategoryFiltersCreateCategoryFilter400Response {
	this := CategoryFiltersCreateCategoryFilter400Response{}
	return &this
}

// NewCategoryFiltersCreateCategoryFilter400ResponseWithDefaults instantiates a new CategoryFiltersCreateCategoryFilter400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryFiltersCreateCategoryFilter400ResponseWithDefaults() *CategoryFiltersCreateCategoryFilter400Response {
	this := CategoryFiltersCreateCategoryFilter400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CategoryFiltersCreateCategoryFilter400Response) GetError() CategoryFiltersCreateCategoryFilter400ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret CategoryFiltersCreateCategoryFilter400ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryFiltersCreateCategoryFilter400Response) GetErrorOk() (*CategoryFiltersCreateCategoryFilter400ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CategoryFiltersCreateCategoryFilter400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CategoryFiltersCreateCategoryFilter400ResponseError and assigns it to the Error field.
func (o *CategoryFiltersCreateCategoryFilter400Response) SetError(v CategoryFiltersCreateCategoryFilter400ResponseError) {
	o.Error = &v
}

func (o CategoryFiltersCreateCategoryFilter400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryFiltersCreateCategoryFilter400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableCategoryFiltersCreateCategoryFilter400Response struct {
	value *CategoryFiltersCreateCategoryFilter400Response
	isSet bool
}

func (v NullableCategoryFiltersCreateCategoryFilter400Response) Get() *CategoryFiltersCreateCategoryFilter400Response {
	return v.value
}

func (v *NullableCategoryFiltersCreateCategoryFilter400Response) Set(val *CategoryFiltersCreateCategoryFilter400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryFiltersCreateCategoryFilter400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryFiltersCreateCategoryFilter400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryFiltersCreateCategoryFilter400Response(val *CategoryFiltersCreateCategoryFilter400Response) *NullableCategoryFiltersCreateCategoryFilter400Response {
	return &NullableCategoryFiltersCreateCategoryFilter400Response{value: val, isSet: true}
}

func (v NullableCategoryFiltersCreateCategoryFilter400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryFiltersCreateCategoryFilter400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
