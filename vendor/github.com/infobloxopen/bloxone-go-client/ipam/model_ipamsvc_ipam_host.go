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

// checks if the IpamsvcIpamHost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcIpamHost{}

// IpamsvcIpamHost The __IpamHost__ object (_ipam/host_) represents any network connected equipment that is assigned one or more IP addresses.
type IpamsvcIpamHost struct {
	// The list of all addresses associated with the IPAM host, which may be in different IP spaces.
	Addresses []IpamsvcHostAddress `json:"addresses,omitempty"`
	// This flag specifies if resource records have to be auto generated for the host.
	AutoGenerateRecords *bool `json:"auto_generate_records,omitempty"`
	// The description for the IPAM host. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The name records to be generated for the host.  This field is required if _auto_generate_records_ is true.
	HostNames []IpamsvcHostName `json:"host_names,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The name of the IPAM host. Must contain 1 to 256 characters. Can include UTF-8.
	Name string `json:"name"`
	// The tags for the IPAM host in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewIpamsvcIpamHost instantiates a new IpamsvcIpamHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcIpamHost(name string) *IpamsvcIpamHost {
	this := IpamsvcIpamHost{}
	this.Name = name
	return &this
}

// NewIpamsvcIpamHostWithDefaults instantiates a new IpamsvcIpamHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcIpamHostWithDefaults() *IpamsvcIpamHost {
	this := IpamsvcIpamHost{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *IpamsvcIpamHost) GetAddresses() []IpamsvcHostAddress {
	if o == nil || IsNil(o.Addresses) {
		var ret []IpamsvcHostAddress
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcIpamHost) GetAddressesOk() ([]IpamsvcHostAddress, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *IpamsvcIpamHost) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []IpamsvcHostAddress and assigns it to the Addresses field.
func (o *IpamsvcIpamHost) SetAddresses(v []IpamsvcHostAddress) {
	o.Addresses = v
}

// GetAutoGenerateRecords returns the AutoGenerateRecords field value if set, zero value otherwise.
func (o *IpamsvcIpamHost) GetAutoGenerateRecords() bool {
	if o == nil || IsNil(o.AutoGenerateRecords) {
		var ret bool
		return ret
	}
	return *o.AutoGenerateRecords
}

// GetAutoGenerateRecordsOk returns a tuple with the AutoGenerateRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcIpamHost) GetAutoGenerateRecordsOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoGenerateRecords) {
		return nil, false
	}
	return o.AutoGenerateRecords, true
}

// HasAutoGenerateRecords returns a boolean if a field has been set.
func (o *IpamsvcIpamHost) HasAutoGenerateRecords() bool {
	if o != nil && !IsNil(o.AutoGenerateRecords) {
		return true
	}

	return false
}

// SetAutoGenerateRecords gets a reference to the given bool and assigns it to the AutoGenerateRecords field.
func (o *IpamsvcIpamHost) SetAutoGenerateRecords(v bool) {
	o.AutoGenerateRecords = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IpamsvcIpamHost) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcIpamHost) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IpamsvcIpamHost) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *IpamsvcIpamHost) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IpamsvcIpamHost) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcIpamHost) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IpamsvcIpamHost) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IpamsvcIpamHost) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHostNames returns the HostNames field value if set, zero value otherwise.
func (o *IpamsvcIpamHost) GetHostNames() []IpamsvcHostName {
	if o == nil || IsNil(o.HostNames) {
		var ret []IpamsvcHostName
		return ret
	}
	return o.HostNames
}

// GetHostNamesOk returns a tuple with the HostNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcIpamHost) GetHostNamesOk() ([]IpamsvcHostName, bool) {
	if o == nil || IsNil(o.HostNames) {
		return nil, false
	}
	return o.HostNames, true
}

// HasHostNames returns a boolean if a field has been set.
func (o *IpamsvcIpamHost) HasHostNames() bool {
	if o != nil && !IsNil(o.HostNames) {
		return true
	}

	return false
}

// SetHostNames gets a reference to the given []IpamsvcHostName and assigns it to the HostNames field.
func (o *IpamsvcIpamHost) SetHostNames(v []IpamsvcHostName) {
	o.HostNames = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpamsvcIpamHost) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcIpamHost) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpamsvcIpamHost) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IpamsvcIpamHost) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *IpamsvcIpamHost) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IpamsvcIpamHost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IpamsvcIpamHost) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IpamsvcIpamHost) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcIpamHost) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IpamsvcIpamHost) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *IpamsvcIpamHost) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IpamsvcIpamHost) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcIpamHost) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IpamsvcIpamHost) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IpamsvcIpamHost) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o IpamsvcIpamHost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcIpamHost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.AutoGenerateRecords) {
		toSerialize["auto_generate_records"] = o.AutoGenerateRecords
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.HostNames) {
		toSerialize["host_names"] = o.HostNames
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableIpamsvcIpamHost struct {
	value *IpamsvcIpamHost
	isSet bool
}

func (v NullableIpamsvcIpamHost) Get() *IpamsvcIpamHost {
	return v.value
}

func (v *NullableIpamsvcIpamHost) Set(val *IpamsvcIpamHost) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcIpamHost) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcIpamHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcIpamHost(val *IpamsvcIpamHost) *NullableIpamsvcIpamHost {
	return &NullableIpamsvcIpamHost{value: val, isSet: true}
}

func (v NullableIpamsvcIpamHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcIpamHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}