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

// checks if the IpamsvcOptionGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcOptionGroup{}

// IpamsvcOptionGroup An __OptionGroup__ object (_dhcp/option_group_) is a named collection of options.
type IpamsvcOptionGroup struct {
	// The description for the option group. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The list of DHCP options for the option group. May be either a specific option or a group of options.
	DhcpOptions []IpamsvcOptionItem `json:"dhcp_options,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The name of the option group. Must contain 1 to 256 characters. Can include UTF-8.
	Name string `json:"name"`
	// The type of protocol (_ip4_ or _ip6_).
	Protocol *string `json:"protocol,omitempty"`
	// The tags for the option group in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewIpamsvcOptionGroup instantiates a new IpamsvcOptionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcOptionGroup(name string) *IpamsvcOptionGroup {
	this := IpamsvcOptionGroup{}
	this.Name = name
	return &this
}

// NewIpamsvcOptionGroupWithDefaults instantiates a new IpamsvcOptionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcOptionGroupWithDefaults() *IpamsvcOptionGroup {
	this := IpamsvcOptionGroup{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IpamsvcOptionGroup) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionGroup) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IpamsvcOptionGroup) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *IpamsvcOptionGroup) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IpamsvcOptionGroup) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionGroup) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IpamsvcOptionGroup) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IpamsvcOptionGroup) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *IpamsvcOptionGroup) GetDhcpOptions() []IpamsvcOptionItem {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret []IpamsvcOptionItem
		return ret
	}
	return o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionGroup) GetDhcpOptionsOk() ([]IpamsvcOptionItem, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *IpamsvcOptionGroup) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []IpamsvcOptionItem and assigns it to the DhcpOptions field.
func (o *IpamsvcOptionGroup) SetDhcpOptions(v []IpamsvcOptionItem) {
	o.DhcpOptions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpamsvcOptionGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpamsvcOptionGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IpamsvcOptionGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *IpamsvcOptionGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IpamsvcOptionGroup) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *IpamsvcOptionGroup) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionGroup) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *IpamsvcOptionGroup) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *IpamsvcOptionGroup) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IpamsvcOptionGroup) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionGroup) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IpamsvcOptionGroup) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *IpamsvcOptionGroup) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IpamsvcOptionGroup) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcOptionGroup) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IpamsvcOptionGroup) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IpamsvcOptionGroup) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o IpamsvcOptionGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcOptionGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcp_options"] = o.DhcpOptions
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableIpamsvcOptionGroup struct {
	value *IpamsvcOptionGroup
	isSet bool
}

func (v NullableIpamsvcOptionGroup) Get() *IpamsvcOptionGroup {
	return v.value
}

func (v *NullableIpamsvcOptionGroup) Set(val *IpamsvcOptionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcOptionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcOptionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcOptionGroup(val *IpamsvcOptionGroup) *NullableIpamsvcOptionGroup {
	return &NullableIpamsvcOptionGroup{value: val, isSet: true}
}

func (v NullableIpamsvcOptionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcOptionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
