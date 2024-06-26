/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
)

// checks if the ForwardZoneConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForwardZoneConfig{}

// ForwardZoneConfig struct for ForwardZoneConfig
type ForwardZoneConfig struct {
	// Optional. External DNS servers to forward to. Order is not significant.
	ExternalForwarders []Forwarder `json:"external_forwarders,omitempty"`
	// The resource identifier.
	Hosts []string `json:"hosts,omitempty"`
	// The resource identifier.
	InternalForwarders []string `json:"internal_forwarders,omitempty"`
	// The resource identifier.
	Nsgs                 []string `json:"nsgs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ForwardZoneConfig ForwardZoneConfig

// NewForwardZoneConfig instantiates a new ForwardZoneConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForwardZoneConfig() *ForwardZoneConfig {
	this := ForwardZoneConfig{}
	return &this
}

// NewForwardZoneConfigWithDefaults instantiates a new ForwardZoneConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForwardZoneConfigWithDefaults() *ForwardZoneConfig {
	this := ForwardZoneConfig{}
	return &this
}

// GetExternalForwarders returns the ExternalForwarders field value if set, zero value otherwise.
func (o *ForwardZoneConfig) GetExternalForwarders() []Forwarder {
	if o == nil || IsNil(o.ExternalForwarders) {
		var ret []Forwarder
		return ret
	}
	return o.ExternalForwarders
}

// GetExternalForwardersOk returns a tuple with the ExternalForwarders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZoneConfig) GetExternalForwardersOk() ([]Forwarder, bool) {
	if o == nil || IsNil(o.ExternalForwarders) {
		return nil, false
	}
	return o.ExternalForwarders, true
}

// HasExternalForwarders returns a boolean if a field has been set.
func (o *ForwardZoneConfig) HasExternalForwarders() bool {
	if o != nil && !IsNil(o.ExternalForwarders) {
		return true
	}

	return false
}

// SetExternalForwarders gets a reference to the given []Forwarder and assigns it to the ExternalForwarders field.
func (o *ForwardZoneConfig) SetExternalForwarders(v []Forwarder) {
	o.ExternalForwarders = v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *ForwardZoneConfig) GetHosts() []string {
	if o == nil || IsNil(o.Hosts) {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZoneConfig) GetHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *ForwardZoneConfig) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *ForwardZoneConfig) SetHosts(v []string) {
	o.Hosts = v
}

// GetInternalForwarders returns the InternalForwarders field value if set, zero value otherwise.
func (o *ForwardZoneConfig) GetInternalForwarders() []string {
	if o == nil || IsNil(o.InternalForwarders) {
		var ret []string
		return ret
	}
	return o.InternalForwarders
}

// GetInternalForwardersOk returns a tuple with the InternalForwarders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZoneConfig) GetInternalForwardersOk() ([]string, bool) {
	if o == nil || IsNil(o.InternalForwarders) {
		return nil, false
	}
	return o.InternalForwarders, true
}

// HasInternalForwarders returns a boolean if a field has been set.
func (o *ForwardZoneConfig) HasInternalForwarders() bool {
	if o != nil && !IsNil(o.InternalForwarders) {
		return true
	}

	return false
}

// SetInternalForwarders gets a reference to the given []string and assigns it to the InternalForwarders field.
func (o *ForwardZoneConfig) SetInternalForwarders(v []string) {
	o.InternalForwarders = v
}

// GetNsgs returns the Nsgs field value if set, zero value otherwise.
func (o *ForwardZoneConfig) GetNsgs() []string {
	if o == nil || IsNil(o.Nsgs) {
		var ret []string
		return ret
	}
	return o.Nsgs
}

// GetNsgsOk returns a tuple with the Nsgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZoneConfig) GetNsgsOk() ([]string, bool) {
	if o == nil || IsNil(o.Nsgs) {
		return nil, false
	}
	return o.Nsgs, true
}

// HasNsgs returns a boolean if a field has been set.
func (o *ForwardZoneConfig) HasNsgs() bool {
	if o != nil && !IsNil(o.Nsgs) {
		return true
	}

	return false
}

// SetNsgs gets a reference to the given []string and assigns it to the Nsgs field.
func (o *ForwardZoneConfig) SetNsgs(v []string) {
	o.Nsgs = v
}

func (o ForwardZoneConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForwardZoneConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalForwarders) {
		toSerialize["external_forwarders"] = o.ExternalForwarders
	}
	if !IsNil(o.Hosts) {
		toSerialize["hosts"] = o.Hosts
	}
	if !IsNil(o.InternalForwarders) {
		toSerialize["internal_forwarders"] = o.InternalForwarders
	}
	if !IsNil(o.Nsgs) {
		toSerialize["nsgs"] = o.Nsgs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ForwardZoneConfig) UnmarshalJSON(data []byte) (err error) {
	varForwardZoneConfig := _ForwardZoneConfig{}

	err = json.Unmarshal(data, &varForwardZoneConfig)

	if err != nil {
		return err
	}

	*o = ForwardZoneConfig(varForwardZoneConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "external_forwarders")
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "internal_forwarders")
		delete(additionalProperties, "nsgs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableForwardZoneConfig struct {
	value *ForwardZoneConfig
	isSet bool
}

func (v NullableForwardZoneConfig) Get() *ForwardZoneConfig {
	return v.value
}

func (v *NullableForwardZoneConfig) Set(val *ForwardZoneConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableForwardZoneConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableForwardZoneConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForwardZoneConfig(val *ForwardZoneConfig) *NullableForwardZoneConfig {
	return &NullableForwardZoneConfig{value: val, isSet: true}
}

func (v NullableForwardZoneConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForwardZoneConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
