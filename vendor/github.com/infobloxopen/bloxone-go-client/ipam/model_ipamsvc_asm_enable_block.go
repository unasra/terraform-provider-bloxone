/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"time"
)

// checks if the IpamsvcAsmEnableBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcAsmEnableBlock{}

// IpamsvcAsmEnableBlock ASM enable group of fields.
type IpamsvcAsmEnableBlock struct {
	// Indicates whether Automated Scope Management is enabled or not.
	Enable *bool `json:"enable,omitempty"`
	// Indicates whether sending notifications to the users is enabled or not.
	EnableNotification *bool `json:"enable_notification,omitempty"`
	// The date at which notifications will be re-enabled automatically.
	ReenableDate *time.Time `json:"reenable_date,omitempty"`
}

// NewIpamsvcAsmEnableBlock instantiates a new IpamsvcAsmEnableBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcAsmEnableBlock() *IpamsvcAsmEnableBlock {
	this := IpamsvcAsmEnableBlock{}
	return &this
}

// NewIpamsvcAsmEnableBlockWithDefaults instantiates a new IpamsvcAsmEnableBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcAsmEnableBlockWithDefaults() *IpamsvcAsmEnableBlock {
	this := IpamsvcAsmEnableBlock{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *IpamsvcAsmEnableBlock) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAsmEnableBlock) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *IpamsvcAsmEnableBlock) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *IpamsvcAsmEnableBlock) SetEnable(v bool) {
	o.Enable = &v
}

// GetEnableNotification returns the EnableNotification field value if set, zero value otherwise.
func (o *IpamsvcAsmEnableBlock) GetEnableNotification() bool {
	if o == nil || IsNil(o.EnableNotification) {
		var ret bool
		return ret
	}
	return *o.EnableNotification
}

// GetEnableNotificationOk returns a tuple with the EnableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAsmEnableBlock) GetEnableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableNotification) {
		return nil, false
	}
	return o.EnableNotification, true
}

// HasEnableNotification returns a boolean if a field has been set.
func (o *IpamsvcAsmEnableBlock) HasEnableNotification() bool {
	if o != nil && !IsNil(o.EnableNotification) {
		return true
	}

	return false
}

// SetEnableNotification gets a reference to the given bool and assigns it to the EnableNotification field.
func (o *IpamsvcAsmEnableBlock) SetEnableNotification(v bool) {
	o.EnableNotification = &v
}

// GetReenableDate returns the ReenableDate field value if set, zero value otherwise.
func (o *IpamsvcAsmEnableBlock) GetReenableDate() time.Time {
	if o == nil || IsNil(o.ReenableDate) {
		var ret time.Time
		return ret
	}
	return *o.ReenableDate
}

// GetReenableDateOk returns a tuple with the ReenableDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAsmEnableBlock) GetReenableDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReenableDate) {
		return nil, false
	}
	return o.ReenableDate, true
}

// HasReenableDate returns a boolean if a field has been set.
func (o *IpamsvcAsmEnableBlock) HasReenableDate() bool {
	if o != nil && !IsNil(o.ReenableDate) {
		return true
	}

	return false
}

// SetReenableDate gets a reference to the given time.Time and assigns it to the ReenableDate field.
func (o *IpamsvcAsmEnableBlock) SetReenableDate(v time.Time) {
	o.ReenableDate = &v
}

func (o IpamsvcAsmEnableBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcAsmEnableBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.EnableNotification) {
		toSerialize["enable_notification"] = o.EnableNotification
	}
	if !IsNil(o.ReenableDate) {
		toSerialize["reenable_date"] = o.ReenableDate
	}
	return toSerialize, nil
}

type NullableIpamsvcAsmEnableBlock struct {
	value *IpamsvcAsmEnableBlock
	isSet bool
}

func (v NullableIpamsvcAsmEnableBlock) Get() *IpamsvcAsmEnableBlock {
	return v.value
}

func (v *NullableIpamsvcAsmEnableBlock) Set(val *IpamsvcAsmEnableBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcAsmEnableBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcAsmEnableBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcAsmEnableBlock(val *IpamsvcAsmEnableBlock) *NullableIpamsvcAsmEnableBlock {
	return &NullableIpamsvcAsmEnableBlock{value: val, isSet: true}
}

func (v NullableIpamsvcAsmEnableBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcAsmEnableBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}