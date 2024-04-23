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

// checks if the InternalDomainListsInternalDomainsItemsPartialUpdate400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalDomainListsInternalDomainsItemsPartialUpdate400Response{}

// InternalDomainListsInternalDomainsItemsPartialUpdate400Response struct for InternalDomainListsInternalDomainsItemsPartialUpdate400Response
type InternalDomainListsInternalDomainsItemsPartialUpdate400Response struct {
	Error *InternalDomainListsInternalDomainsItemsPartialUpdate400ResponseError `json:"error,omitempty"`
}

// NewInternalDomainListsInternalDomainsItemsPartialUpdate400Response instantiates a new InternalDomainListsInternalDomainsItemsPartialUpdate400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalDomainListsInternalDomainsItemsPartialUpdate400Response() *InternalDomainListsInternalDomainsItemsPartialUpdate400Response {
	this := InternalDomainListsInternalDomainsItemsPartialUpdate400Response{}
	return &this
}

// NewInternalDomainListsInternalDomainsItemsPartialUpdate400ResponseWithDefaults instantiates a new InternalDomainListsInternalDomainsItemsPartialUpdate400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalDomainListsInternalDomainsItemsPartialUpdate400ResponseWithDefaults() *InternalDomainListsInternalDomainsItemsPartialUpdate400Response {
	this := InternalDomainListsInternalDomainsItemsPartialUpdate400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InternalDomainListsInternalDomainsItemsPartialUpdate400Response) GetError() InternalDomainListsInternalDomainsItemsPartialUpdate400ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret InternalDomainListsInternalDomainsItemsPartialUpdate400ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalDomainListsInternalDomainsItemsPartialUpdate400Response) GetErrorOk() (*InternalDomainListsInternalDomainsItemsPartialUpdate400ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InternalDomainListsInternalDomainsItemsPartialUpdate400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given InternalDomainListsInternalDomainsItemsPartialUpdate400ResponseError and assigns it to the Error field.
func (o *InternalDomainListsInternalDomainsItemsPartialUpdate400Response) SetError(v InternalDomainListsInternalDomainsItemsPartialUpdate400ResponseError) {
	o.Error = &v
}

func (o InternalDomainListsInternalDomainsItemsPartialUpdate400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalDomainListsInternalDomainsItemsPartialUpdate400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableInternalDomainListsInternalDomainsItemsPartialUpdate400Response struct {
	value *InternalDomainListsInternalDomainsItemsPartialUpdate400Response
	isSet bool
}

func (v NullableInternalDomainListsInternalDomainsItemsPartialUpdate400Response) Get() *InternalDomainListsInternalDomainsItemsPartialUpdate400Response {
	return v.value
}

func (v *NullableInternalDomainListsInternalDomainsItemsPartialUpdate400Response) Set(val *InternalDomainListsInternalDomainsItemsPartialUpdate400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalDomainListsInternalDomainsItemsPartialUpdate400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalDomainListsInternalDomainsItemsPartialUpdate400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalDomainListsInternalDomainsItemsPartialUpdate400Response(val *InternalDomainListsInternalDomainsItemsPartialUpdate400Response) *NullableInternalDomainListsInternalDomainsItemsPartialUpdate400Response {
	return &NullableInternalDomainListsInternalDomainsItemsPartialUpdate400Response{value: val, isSet: true}
}

func (v NullableInternalDomainListsInternalDomainsItemsPartialUpdate400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalDomainListsInternalDomainsItemsPartialUpdate400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
