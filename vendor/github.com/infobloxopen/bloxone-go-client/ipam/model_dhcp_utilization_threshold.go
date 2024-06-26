/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
)

// checks if the DHCPUtilizationThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DHCPUtilizationThreshold{}

// DHCPUtilizationThreshold A __DHCPUtilizationThreshold__ object represents threshold settings for DHCP utilization.
type DHCPUtilizationThreshold struct {
	// Indicates whether the DHCP utilization threshold is enabled or not.
	Enabled bool `json:"enabled"`
	// The high threshold value for DHCP utilization in percentage.
	High int64 `json:"high"`
	// The low threshold value for DHCP utilization in percentage.
	Low                  int64 `json:"low"`
	AdditionalProperties map[string]interface{}
}

type _DHCPUtilizationThreshold DHCPUtilizationThreshold

// NewDHCPUtilizationThreshold instantiates a new DHCPUtilizationThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDHCPUtilizationThreshold(enabled bool, high int64, low int64) *DHCPUtilizationThreshold {
	this := DHCPUtilizationThreshold{}
	this.Enabled = enabled
	this.High = high
	this.Low = low
	return &this
}

// NewDHCPUtilizationThresholdWithDefaults instantiates a new DHCPUtilizationThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDHCPUtilizationThresholdWithDefaults() *DHCPUtilizationThreshold {
	this := DHCPUtilizationThreshold{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DHCPUtilizationThreshold) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DHCPUtilizationThreshold) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DHCPUtilizationThreshold) SetEnabled(v bool) {
	o.Enabled = v
}

// GetHigh returns the High field value
func (o *DHCPUtilizationThreshold) GetHigh() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.High
}

// GetHighOk returns a tuple with the High field value
// and a boolean to check if the value has been set.
func (o *DHCPUtilizationThreshold) GetHighOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.High, true
}

// SetHigh sets field value
func (o *DHCPUtilizationThreshold) SetHigh(v int64) {
	o.High = v
}

// GetLow returns the Low field value
func (o *DHCPUtilizationThreshold) GetLow() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Low
}

// GetLowOk returns a tuple with the Low field value
// and a boolean to check if the value has been set.
func (o *DHCPUtilizationThreshold) GetLowOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Low, true
}

// SetLow sets field value
func (o *DHCPUtilizationThreshold) SetLow(v int64) {
	o.Low = v
}

func (o DHCPUtilizationThreshold) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DHCPUtilizationThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["high"] = o.High
	toSerialize["low"] = o.Low

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DHCPUtilizationThreshold) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"high",
		"low",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDHCPUtilizationThreshold := _DHCPUtilizationThreshold{}

	err = json.Unmarshal(data, &varDHCPUtilizationThreshold)

	if err != nil {
		return err
	}

	*o = DHCPUtilizationThreshold(varDHCPUtilizationThreshold)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "high")
		delete(additionalProperties, "low")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDHCPUtilizationThreshold struct {
	value *DHCPUtilizationThreshold
	isSet bool
}

func (v NullableDHCPUtilizationThreshold) Get() *DHCPUtilizationThreshold {
	return v.value
}

func (v *NullableDHCPUtilizationThreshold) Set(val *DHCPUtilizationThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableDHCPUtilizationThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableDHCPUtilizationThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDHCPUtilizationThreshold(val *DHCPUtilizationThreshold) *NullableDHCPUtilizationThreshold {
	return &NullableDHCPUtilizationThreshold{value: val, isSet: true}
}

func (v NullableDHCPUtilizationThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDHCPUtilizationThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
