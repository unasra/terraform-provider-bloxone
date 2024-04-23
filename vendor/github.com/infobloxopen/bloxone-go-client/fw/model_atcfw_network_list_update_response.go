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

// checks if the AtcfwNetworkListUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtcfwNetworkListUpdateResponse{}

// AtcfwNetworkListUpdateResponse The Network List update response.
type AtcfwNetworkListUpdateResponse struct {
	Results *AtcfwNetworkList `json:"results,omitempty"`
}

// NewAtcfwNetworkListUpdateResponse instantiates a new AtcfwNetworkListUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtcfwNetworkListUpdateResponse() *AtcfwNetworkListUpdateResponse {
	this := AtcfwNetworkListUpdateResponse{}
	return &this
}

// NewAtcfwNetworkListUpdateResponseWithDefaults instantiates a new AtcfwNetworkListUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtcfwNetworkListUpdateResponseWithDefaults() *AtcfwNetworkListUpdateResponse {
	this := AtcfwNetworkListUpdateResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AtcfwNetworkListUpdateResponse) GetResults() AtcfwNetworkList {
	if o == nil || IsNil(o.Results) {
		var ret AtcfwNetworkList
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwNetworkListUpdateResponse) GetResultsOk() (*AtcfwNetworkList, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AtcfwNetworkListUpdateResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given AtcfwNetworkList and assigns it to the Results field.
func (o *AtcfwNetworkListUpdateResponse) SetResults(v AtcfwNetworkList) {
	o.Results = &v
}

func (o AtcfwNetworkListUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtcfwNetworkListUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableAtcfwNetworkListUpdateResponse struct {
	value *AtcfwNetworkListUpdateResponse
	isSet bool
}

func (v NullableAtcfwNetworkListUpdateResponse) Get() *AtcfwNetworkListUpdateResponse {
	return v.value
}

func (v *NullableAtcfwNetworkListUpdateResponse) Set(val *AtcfwNetworkListUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAtcfwNetworkListUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAtcfwNetworkListUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtcfwNetworkListUpdateResponse(val *AtcfwNetworkListUpdateResponse) *NullableAtcfwNetworkListUpdateResponse {
	return &NullableAtcfwNetworkListUpdateResponse{value: val, isSet: true}
}

func (v NullableAtcfwNetworkListUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtcfwNetworkListUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
