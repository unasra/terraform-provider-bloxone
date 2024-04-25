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

// checks if the ApplicationFiltersDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationFiltersDeleteRequest{}

// ApplicationFiltersDeleteRequest The Application Filter delete request.
type ApplicationFiltersDeleteRequest struct {
	// The list of Application Filter object identifiers.
	Ids []int32 `json:"ids,omitempty"`
}

// NewApplicationFiltersDeleteRequest instantiates a new ApplicationFiltersDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFiltersDeleteRequest() *ApplicationFiltersDeleteRequest {
	this := ApplicationFiltersDeleteRequest{}
	return &this
}

// NewApplicationFiltersDeleteRequestWithDefaults instantiates a new ApplicationFiltersDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFiltersDeleteRequestWithDefaults() *ApplicationFiltersDeleteRequest {
	this := ApplicationFiltersDeleteRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ApplicationFiltersDeleteRequest) GetIds() []int32 {
	if o == nil || IsNil(o.Ids) {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFiltersDeleteRequest) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ApplicationFiltersDeleteRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *ApplicationFiltersDeleteRequest) SetIds(v []int32) {
	o.Ids = v
}

func (o ApplicationFiltersDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationFiltersDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableApplicationFiltersDeleteRequest struct {
	value *ApplicationFiltersDeleteRequest
	isSet bool
}

func (v NullableApplicationFiltersDeleteRequest) Get() *ApplicationFiltersDeleteRequest {
	return v.value
}

func (v *NullableApplicationFiltersDeleteRequest) Set(val *ApplicationFiltersDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFiltersDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFiltersDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFiltersDeleteRequest(val *ApplicationFiltersDeleteRequest) *NullableApplicationFiltersDeleteRequest {
	return &NullableApplicationFiltersDeleteRequest{value: val, isSet: true}
}

func (v NullableApplicationFiltersDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFiltersDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}