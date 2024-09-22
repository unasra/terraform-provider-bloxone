/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the CPSubnet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CPSubnet{}

// CPSubnet struct for CPSubnet
type CPSubnet struct {
	Address *string `json:"address,omitempty"`
	Cidr    *int64  `json:"cidr,omitempty"`
	Comment *string `json:"comment,omitempty"`
	// The resource identifier.
	Id          *string `json:"id,omitempty"`
	IpSpaceName *string `json:"ip_space_name,omitempty"`
	// The resource identifier.
	IpSpaceRef *string `json:"ip_space_ref,omitempty"`
	Name       *string `json:"name,omitempty"`
	// The resource identifier.
	Parent               *string `json:"parent,omitempty"`
	Protocol             *string `json:"protocol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CPSubnet CPSubnet

// NewCPSubnet instantiates a new CPSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCPSubnet() *CPSubnet {
	this := CPSubnet{}
	return &this
}

// NewCPSubnetWithDefaults instantiates a new CPSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCPSubnetWithDefaults() *CPSubnet {
	this := CPSubnet{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CPSubnet) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPSubnet) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CPSubnet) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *CPSubnet) SetAddress(v string) {
	o.Address = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *CPSubnet) GetCidr() int64 {
	if o == nil || IsNil(o.Cidr) {
		var ret int64
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPSubnet) GetCidrOk() (*int64, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *CPSubnet) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int64 and assigns it to the Cidr field.
func (o *CPSubnet) SetCidr(v int64) {
	o.Cidr = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CPSubnet) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPSubnet) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CPSubnet) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CPSubnet) SetComment(v string) {
	o.Comment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CPSubnet) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPSubnet) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CPSubnet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CPSubnet) SetId(v string) {
	o.Id = &v
}

// GetIpSpaceName returns the IpSpaceName field value if set, zero value otherwise.
func (o *CPSubnet) GetIpSpaceName() string {
	if o == nil || IsNil(o.IpSpaceName) {
		var ret string
		return ret
	}
	return *o.IpSpaceName
}

// GetIpSpaceNameOk returns a tuple with the IpSpaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPSubnet) GetIpSpaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.IpSpaceName) {
		return nil, false
	}
	return o.IpSpaceName, true
}

// HasIpSpaceName returns a boolean if a field has been set.
func (o *CPSubnet) HasIpSpaceName() bool {
	if o != nil && !IsNil(o.IpSpaceName) {
		return true
	}

	return false
}

// SetIpSpaceName gets a reference to the given string and assigns it to the IpSpaceName field.
func (o *CPSubnet) SetIpSpaceName(v string) {
	o.IpSpaceName = &v
}

// GetIpSpaceRef returns the IpSpaceRef field value if set, zero value otherwise.
func (o *CPSubnet) GetIpSpaceRef() string {
	if o == nil || IsNil(o.IpSpaceRef) {
		var ret string
		return ret
	}
	return *o.IpSpaceRef
}

// GetIpSpaceRefOk returns a tuple with the IpSpaceRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPSubnet) GetIpSpaceRefOk() (*string, bool) {
	if o == nil || IsNil(o.IpSpaceRef) {
		return nil, false
	}
	return o.IpSpaceRef, true
}

// HasIpSpaceRef returns a boolean if a field has been set.
func (o *CPSubnet) HasIpSpaceRef() bool {
	if o != nil && !IsNil(o.IpSpaceRef) {
		return true
	}

	return false
}

// SetIpSpaceRef gets a reference to the given string and assigns it to the IpSpaceRef field.
func (o *CPSubnet) SetIpSpaceRef(v string) {
	o.IpSpaceRef = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CPSubnet) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPSubnet) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CPSubnet) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CPSubnet) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *CPSubnet) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPSubnet) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *CPSubnet) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *CPSubnet) SetParent(v string) {
	o.Parent = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *CPSubnet) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPSubnet) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *CPSubnet) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *CPSubnet) SetProtocol(v string) {
	o.Protocol = &v
}

func (o CPSubnet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CPSubnet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpSpaceName) {
		toSerialize["ip_space_name"] = o.IpSpaceName
	}
	if !IsNil(o.IpSpaceRef) {
		toSerialize["ip_space_ref"] = o.IpSpaceRef
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CPSubnet) UnmarshalJSON(data []byte) (err error) {
	varCPSubnet := _CPSubnet{}

	err = json.Unmarshal(data, &varCPSubnet)

	if err != nil {
		return err
	}

	*o = CPSubnet(varCPSubnet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "cidr")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ip_space_name")
		delete(additionalProperties, "ip_space_ref")
		delete(additionalProperties, "name")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "protocol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCPSubnet struct {
	value *CPSubnet
	isSet bool
}

func (v NullableCPSubnet) Get() *CPSubnet {
	return v.value
}

func (v *NullableCPSubnet) Set(val *CPSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableCPSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableCPSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCPSubnet(val *CPSubnet) *NullableCPSubnet {
	return &NullableCPSubnet{value: val, isSet: true}
}

func (v NullableCPSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCPSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
