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

// checks if the Inheritance2InheritedBool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Inheritance2InheritedBool{}

// Inheritance2InheritedBool The inheritance configuration for a field of type _Bool_.
type Inheritance2InheritedBool struct {
	// The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
	Action *string `json:"action,omitempty"`
	// The human-readable display name for the object referred to by _source_.
	DisplayName *string `json:"display_name,omitempty"`
	// The resource identifier.
	Source *string `json:"source,omitempty"`
	// The inherited value.
	Value                *bool `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Inheritance2InheritedBool Inheritance2InheritedBool

// NewInheritance2InheritedBool instantiates a new Inheritance2InheritedBool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInheritance2InheritedBool() *Inheritance2InheritedBool {
	this := Inheritance2InheritedBool{}
	return &this
}

// NewInheritance2InheritedBoolWithDefaults instantiates a new Inheritance2InheritedBool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInheritance2InheritedBoolWithDefaults() *Inheritance2InheritedBool {
	this := Inheritance2InheritedBool{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Inheritance2InheritedBool) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inheritance2InheritedBool) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Inheritance2InheritedBool) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Inheritance2InheritedBool) SetAction(v string) {
	o.Action = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Inheritance2InheritedBool) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inheritance2InheritedBool) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Inheritance2InheritedBool) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Inheritance2InheritedBool) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Inheritance2InheritedBool) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inheritance2InheritedBool) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Inheritance2InheritedBool) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *Inheritance2InheritedBool) SetSource(v string) {
	o.Source = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Inheritance2InheritedBool) GetValue() bool {
	if o == nil || IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inheritance2InheritedBool) GetValueOk() (*bool, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Inheritance2InheritedBool) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *Inheritance2InheritedBool) SetValue(v bool) {
	o.Value = &v
}

func (o Inheritance2InheritedBool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Inheritance2InheritedBool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Inheritance2InheritedBool) UnmarshalJSON(data []byte) (err error) {
	varInheritance2InheritedBool := _Inheritance2InheritedBool{}

	err = json.Unmarshal(data, &varInheritance2InheritedBool)

	if err != nil {
		return err
	}

	*o = Inheritance2InheritedBool(varInheritance2InheritedBool)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "source")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInheritance2InheritedBool struct {
	value *Inheritance2InheritedBool
	isSet bool
}

func (v NullableInheritance2InheritedBool) Get() *Inheritance2InheritedBool {
	return v.value
}

func (v *NullableInheritance2InheritedBool) Set(val *Inheritance2InheritedBool) {
	v.value = val
	v.isSet = true
}

func (v NullableInheritance2InheritedBool) IsSet() bool {
	return v.isSet
}

func (v *NullableInheritance2InheritedBool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInheritance2InheritedBool(val *Inheritance2InheritedBool) *NullableInheritance2InheritedBool {
	return &NullableInheritance2InheritedBool{value: val, isSet: true}
}

func (v NullableInheritance2InheritedBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInheritance2InheritedBool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
