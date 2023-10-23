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

// checks if the ConfigDelegation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigDelegation{}

// ConfigDelegation DNS zone delegation.
type ConfigDelegation struct {
	// Optional. Comment for zone delegation.
	Comment *string `json:"comment,omitempty"`
	// Required. DNS zone delegation servers. Order is not significant.
	DelegationServers []ConfigDelegationServer `json:"delegation_servers,omitempty"`
	// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating resource records.
	Disabled *bool `json:"disabled,omitempty"`
	// Delegation FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.
	Fqdn string `json:"fqdn"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The resource identifier.
	Parent *string `json:"parent,omitempty"`
	// Delegation FQDN in punycode.
	ProtocolFqdn *string `json:"protocol_fqdn,omitempty"`
	// Tagging specifics.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// The resource identifier.
	View *string `json:"view,omitempty"`
}

// NewConfigDelegation instantiates a new ConfigDelegation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigDelegation(fqdn string) *ConfigDelegation {
	this := ConfigDelegation{}
	this.Fqdn = fqdn
	return &this
}

// NewConfigDelegationWithDefaults instantiates a new ConfigDelegation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigDelegationWithDefaults() *ConfigDelegation {
	this := ConfigDelegation{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ConfigDelegation) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelegation) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ConfigDelegation) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ConfigDelegation) SetComment(v string) {
	o.Comment = &v
}

// GetDelegationServers returns the DelegationServers field value if set, zero value otherwise.
func (o *ConfigDelegation) GetDelegationServers() []ConfigDelegationServer {
	if o == nil || IsNil(o.DelegationServers) {
		var ret []ConfigDelegationServer
		return ret
	}
	return o.DelegationServers
}

// GetDelegationServersOk returns a tuple with the DelegationServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelegation) GetDelegationServersOk() ([]ConfigDelegationServer, bool) {
	if o == nil || IsNil(o.DelegationServers) {
		return nil, false
	}
	return o.DelegationServers, true
}

// HasDelegationServers returns a boolean if a field has been set.
func (o *ConfigDelegation) HasDelegationServers() bool {
	if o != nil && !IsNil(o.DelegationServers) {
		return true
	}

	return false
}

// SetDelegationServers gets a reference to the given []ConfigDelegationServer and assigns it to the DelegationServers field.
func (o *ConfigDelegation) SetDelegationServers(v []ConfigDelegationServer) {
	o.DelegationServers = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ConfigDelegation) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelegation) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ConfigDelegation) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ConfigDelegation) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetFqdn returns the Fqdn field value
func (o *ConfigDelegation) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *ConfigDelegation) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *ConfigDelegation) SetFqdn(v string) {
	o.Fqdn = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigDelegation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelegation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigDelegation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigDelegation) SetId(v string) {
	o.Id = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ConfigDelegation) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelegation) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ConfigDelegation) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *ConfigDelegation) SetParent(v string) {
	o.Parent = &v
}

// GetProtocolFqdn returns the ProtocolFqdn field value if set, zero value otherwise.
func (o *ConfigDelegation) GetProtocolFqdn() string {
	if o == nil || IsNil(o.ProtocolFqdn) {
		var ret string
		return ret
	}
	return *o.ProtocolFqdn
}

// GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelegation) GetProtocolFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolFqdn) {
		return nil, false
	}
	return o.ProtocolFqdn, true
}

// HasProtocolFqdn returns a boolean if a field has been set.
func (o *ConfigDelegation) HasProtocolFqdn() bool {
	if o != nil && !IsNil(o.ProtocolFqdn) {
		return true
	}

	return false
}

// SetProtocolFqdn gets a reference to the given string and assigns it to the ProtocolFqdn field.
func (o *ConfigDelegation) SetProtocolFqdn(v string) {
	o.ProtocolFqdn = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfigDelegation) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelegation) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfigDelegation) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *ConfigDelegation) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *ConfigDelegation) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelegation) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *ConfigDelegation) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *ConfigDelegation) SetView(v string) {
	o.View = &v
}

func (o ConfigDelegation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigDelegation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.DelegationServers) {
		toSerialize["delegation_servers"] = o.DelegationServers
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	toSerialize["fqdn"] = o.Fqdn
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.ProtocolFqdn) {
		toSerialize["protocol_fqdn"] = o.ProtocolFqdn
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	return toSerialize, nil
}

type NullableConfigDelegation struct {
	value *ConfigDelegation
	isSet bool
}

func (v NullableConfigDelegation) Get() *ConfigDelegation {
	return v.value
}

func (v *NullableConfigDelegation) Set(val *ConfigDelegation) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigDelegation) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigDelegation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigDelegation(val *ConfigDelegation) *NullableConfigDelegation {
	return &NullableConfigDelegation{value: val, isSet: true}
}

func (v NullableConfigDelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigDelegation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
