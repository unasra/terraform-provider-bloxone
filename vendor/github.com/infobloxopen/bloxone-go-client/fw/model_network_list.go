/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
	"time"
)

// checks if the NetworkList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkList{}

// NetworkList The Network List object.  Before you can apply security policies, you must first define the networks that you want to protect from malicious attacks. The first step in configuring BloxOne Cloud is to set up DNS Firewall by defining your remote networks. You identify these external networks by their IP addresses. A network can contain a group of IPv4 or IPv6 addresses or blocks.
type NetworkList struct {
	// The time when this Network List object was created.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// The brief description for the network list.
	Description *string `json:"description,omitempty"`
	// The Network List object identifier.
	Id *int32 `json:"id,omitempty"`
	// The list of networks' CIDRs that are subject for malicious attacks protection.
	Items []string `json:"items,omitempty"`
	// The name of the network list.
	Name *string `json:"name,omitempty"`
	// The identifier of the security policy with which the network list is associated.
	PolicyId *int32 `json:"policy_id,omitempty"`
	// The time when this Network List object was last updated.
	UpdatedTime          *time.Time `json:"updated_time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkList NetworkList

// NewNetworkList instantiates a new NetworkList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkList() *NetworkList {
	this := NetworkList{}
	return &this
}

// NewNetworkListWithDefaults instantiates a new NetworkList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkListWithDefaults() *NetworkList {
	this := NetworkList{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *NetworkList) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkList) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *NetworkList) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *NetworkList) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkList) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkList) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkList) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkList) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkList) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkList) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkList) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NetworkList) SetId(v int32) {
	o.Id = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *NetworkList) GetItems() []string {
	if o == nil || IsNil(o.Items) {
		var ret []string
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkList) GetItemsOk() ([]string, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *NetworkList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []string and assigns it to the Items field.
func (o *NetworkList) SetItems(v []string) {
	o.Items = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkList) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkList) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkList) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkList) SetName(v string) {
	o.Name = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *NetworkList) GetPolicyId() int32 {
	if o == nil || IsNil(o.PolicyId) {
		var ret int32
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkList) GetPolicyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *NetworkList) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given int32 and assigns it to the PolicyId field.
func (o *NetworkList) SetPolicyId(v int32) {
	o.PolicyId = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *NetworkList) GetUpdatedTime() time.Time {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkList) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *NetworkList) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given time.Time and assigns it to the UpdatedTime field.
func (o *NetworkList) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = &v
}

func (o NetworkList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PolicyId) {
		toSerialize["policy_id"] = o.PolicyId
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updated_time"] = o.UpdatedTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkList) UnmarshalJSON(data []byte) (err error) {
	varNetworkList := _NetworkList{}

	err = json.Unmarshal(data, &varNetworkList)

	if err != nil {
		return err
	}

	*o = NetworkList(varNetworkList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "items")
		delete(additionalProperties, "name")
		delete(additionalProperties, "policy_id")
		delete(additionalProperties, "updated_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkList struct {
	value *NetworkList
	isSet bool
}

func (v NullableNetworkList) Get() *NetworkList {
	return v.value
}

func (v *NullableNetworkList) Set(val *NetworkList) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkList) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkList(val *NetworkList) *NullableNetworkList {
	return &NullableNetworkList{value: val, isSet: true}
}

func (v NullableNetworkList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
