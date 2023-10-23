/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
)

// checks if the ConfigDNSSECValidationBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigDNSSECValidationBlock{}

// ConfigDNSSECValidationBlock Block for fields: _dnssec_enabled_, _dnssec_enable_validation_, _dnssec_validate_expiry_, _dnssec_trust_anchors_.
type ConfigDNSSECValidationBlock struct {
	// Optional. Field config for _dnssec_enable_validation_ field.
	DnssecEnableValidation *bool `json:"dnssec_enable_validation,omitempty"`
	// Optional. Field config for _dnssec_enabled_ field.
	DnssecEnabled *bool `json:"dnssec_enabled,omitempty"`
	// Optional. Field config for _dnssec_trust_anchors_ field.
	DnssecTrustAnchors []ConfigTrustAnchor `json:"dnssec_trust_anchors,omitempty"`
	// Optional. Field config for _dnssec_validate_expiry_ field.
	DnssecValidateExpiry *bool `json:"dnssec_validate_expiry,omitempty"`
}

// NewConfigDNSSECValidationBlock instantiates a new ConfigDNSSECValidationBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigDNSSECValidationBlock() *ConfigDNSSECValidationBlock {
	this := ConfigDNSSECValidationBlock{}
	return &this
}

// NewConfigDNSSECValidationBlockWithDefaults instantiates a new ConfigDNSSECValidationBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigDNSSECValidationBlockWithDefaults() *ConfigDNSSECValidationBlock {
	this := ConfigDNSSECValidationBlock{}
	return &this
}

// GetDnssecEnableValidation returns the DnssecEnableValidation field value if set, zero value otherwise.
func (o *ConfigDNSSECValidationBlock) GetDnssecEnableValidation() bool {
	if o == nil || IsNil(o.DnssecEnableValidation) {
		var ret bool
		return ret
	}
	return *o.DnssecEnableValidation
}

// GetDnssecEnableValidationOk returns a tuple with the DnssecEnableValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDNSSECValidationBlock) GetDnssecEnableValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.DnssecEnableValidation) {
		return nil, false
	}
	return o.DnssecEnableValidation, true
}

// HasDnssecEnableValidation returns a boolean if a field has been set.
func (o *ConfigDNSSECValidationBlock) HasDnssecEnableValidation() bool {
	if o != nil && !IsNil(o.DnssecEnableValidation) {
		return true
	}

	return false
}

// SetDnssecEnableValidation gets a reference to the given bool and assigns it to the DnssecEnableValidation field.
func (o *ConfigDNSSECValidationBlock) SetDnssecEnableValidation(v bool) {
	o.DnssecEnableValidation = &v
}

// GetDnssecEnabled returns the DnssecEnabled field value if set, zero value otherwise.
func (o *ConfigDNSSECValidationBlock) GetDnssecEnabled() bool {
	if o == nil || IsNil(o.DnssecEnabled) {
		var ret bool
		return ret
	}
	return *o.DnssecEnabled
}

// GetDnssecEnabledOk returns a tuple with the DnssecEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDNSSECValidationBlock) GetDnssecEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DnssecEnabled) {
		return nil, false
	}
	return o.DnssecEnabled, true
}

// HasDnssecEnabled returns a boolean if a field has been set.
func (o *ConfigDNSSECValidationBlock) HasDnssecEnabled() bool {
	if o != nil && !IsNil(o.DnssecEnabled) {
		return true
	}

	return false
}

// SetDnssecEnabled gets a reference to the given bool and assigns it to the DnssecEnabled field.
func (o *ConfigDNSSECValidationBlock) SetDnssecEnabled(v bool) {
	o.DnssecEnabled = &v
}

// GetDnssecTrustAnchors returns the DnssecTrustAnchors field value if set, zero value otherwise.
func (o *ConfigDNSSECValidationBlock) GetDnssecTrustAnchors() []ConfigTrustAnchor {
	if o == nil || IsNil(o.DnssecTrustAnchors) {
		var ret []ConfigTrustAnchor
		return ret
	}
	return o.DnssecTrustAnchors
}

// GetDnssecTrustAnchorsOk returns a tuple with the DnssecTrustAnchors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDNSSECValidationBlock) GetDnssecTrustAnchorsOk() ([]ConfigTrustAnchor, bool) {
	if o == nil || IsNil(o.DnssecTrustAnchors) {
		return nil, false
	}
	return o.DnssecTrustAnchors, true
}

// HasDnssecTrustAnchors returns a boolean if a field has been set.
func (o *ConfigDNSSECValidationBlock) HasDnssecTrustAnchors() bool {
	if o != nil && !IsNil(o.DnssecTrustAnchors) {
		return true
	}

	return false
}

// SetDnssecTrustAnchors gets a reference to the given []ConfigTrustAnchor and assigns it to the DnssecTrustAnchors field.
func (o *ConfigDNSSECValidationBlock) SetDnssecTrustAnchors(v []ConfigTrustAnchor) {
	o.DnssecTrustAnchors = v
}

// GetDnssecValidateExpiry returns the DnssecValidateExpiry field value if set, zero value otherwise.
func (o *ConfigDNSSECValidationBlock) GetDnssecValidateExpiry() bool {
	if o == nil || IsNil(o.DnssecValidateExpiry) {
		var ret bool
		return ret
	}
	return *o.DnssecValidateExpiry
}

// GetDnssecValidateExpiryOk returns a tuple with the DnssecValidateExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDNSSECValidationBlock) GetDnssecValidateExpiryOk() (*bool, bool) {
	if o == nil || IsNil(o.DnssecValidateExpiry) {
		return nil, false
	}
	return o.DnssecValidateExpiry, true
}

// HasDnssecValidateExpiry returns a boolean if a field has been set.
func (o *ConfigDNSSECValidationBlock) HasDnssecValidateExpiry() bool {
	if o != nil && !IsNil(o.DnssecValidateExpiry) {
		return true
	}

	return false
}

// SetDnssecValidateExpiry gets a reference to the given bool and assigns it to the DnssecValidateExpiry field.
func (o *ConfigDNSSECValidationBlock) SetDnssecValidateExpiry(v bool) {
	o.DnssecValidateExpiry = &v
}

func (o ConfigDNSSECValidationBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigDNSSECValidationBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DnssecEnableValidation) {
		toSerialize["dnssec_enable_validation"] = o.DnssecEnableValidation
	}
	if !IsNil(o.DnssecEnabled) {
		toSerialize["dnssec_enabled"] = o.DnssecEnabled
	}
	if !IsNil(o.DnssecTrustAnchors) {
		toSerialize["dnssec_trust_anchors"] = o.DnssecTrustAnchors
	}
	if !IsNil(o.DnssecValidateExpiry) {
		toSerialize["dnssec_validate_expiry"] = o.DnssecValidateExpiry
	}
	return toSerialize, nil
}

type NullableConfigDNSSECValidationBlock struct {
	value *ConfigDNSSECValidationBlock
	isSet bool
}

func (v NullableConfigDNSSECValidationBlock) Get() *ConfigDNSSECValidationBlock {
	return v.value
}

func (v *NullableConfigDNSSECValidationBlock) Set(val *ConfigDNSSECValidationBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigDNSSECValidationBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigDNSSECValidationBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigDNSSECValidationBlock(val *ConfigDNSSECValidationBlock) *NullableConfigDNSSECValidationBlock {
	return &NullableConfigDNSSECValidationBlock{value: val, isSet: true}
}

func (v NullableConfigDNSSECValidationBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigDNSSECValidationBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}