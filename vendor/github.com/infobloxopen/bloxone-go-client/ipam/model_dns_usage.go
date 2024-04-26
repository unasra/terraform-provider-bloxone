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

// checks if the DNSUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNSUsage{}

// DNSUsage The __DNSUsage__ object tracks DNS usage of a resource record on an address.
type DNSUsage struct {
	// The absolute name of the resource record in associated zone.
	AbsoluteName *string `json:"absolute_name,omitempty"`
	// The address of the referenced record.
	Address *string `json:"address,omitempty"`
	// The DNS rdata of the referenced record.
	DnsRdata *string `json:"dns_rdata,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The name in zone of the referenced record.
	Name *string `json:"name,omitempty"`
	// The resource identifier.
	Record *string `json:"record,omitempty"`
	// The resource identifier.
	Space *string `json:"space,omitempty"`
	// The type of the referenced record.
	Type *string `json:"type,omitempty"`
	// The resource identifier.
	View *string `json:"view,omitempty"`
	// The resource identifier.
	Zone                 *string `json:"zone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DNSUsage DNSUsage

// NewDNSUsage instantiates a new DNSUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSUsage() *DNSUsage {
	this := DNSUsage{}
	return &this
}

// NewDNSUsageWithDefaults instantiates a new DNSUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSUsageWithDefaults() *DNSUsage {
	this := DNSUsage{}
	return &this
}

// GetAbsoluteName returns the AbsoluteName field value if set, zero value otherwise.
func (o *DNSUsage) GetAbsoluteName() string {
	if o == nil || IsNil(o.AbsoluteName) {
		var ret string
		return ret
	}
	return *o.AbsoluteName
}

// GetAbsoluteNameOk returns a tuple with the AbsoluteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSUsage) GetAbsoluteNameOk() (*string, bool) {
	if o == nil || IsNil(o.AbsoluteName) {
		return nil, false
	}
	return o.AbsoluteName, true
}

// HasAbsoluteName returns a boolean if a field has been set.
func (o *DNSUsage) HasAbsoluteName() bool {
	if o != nil && !IsNil(o.AbsoluteName) {
		return true
	}

	return false
}

// SetAbsoluteName gets a reference to the given string and assigns it to the AbsoluteName field.
func (o *DNSUsage) SetAbsoluteName(v string) {
	o.AbsoluteName = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DNSUsage) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSUsage) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DNSUsage) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DNSUsage) SetAddress(v string) {
	o.Address = &v
}

// GetDnsRdata returns the DnsRdata field value if set, zero value otherwise.
func (o *DNSUsage) GetDnsRdata() string {
	if o == nil || IsNil(o.DnsRdata) {
		var ret string
		return ret
	}
	return *o.DnsRdata
}

// GetDnsRdataOk returns a tuple with the DnsRdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSUsage) GetDnsRdataOk() (*string, bool) {
	if o == nil || IsNil(o.DnsRdata) {
		return nil, false
	}
	return o.DnsRdata, true
}

// HasDnsRdata returns a boolean if a field has been set.
func (o *DNSUsage) HasDnsRdata() bool {
	if o != nil && !IsNil(o.DnsRdata) {
		return true
	}

	return false
}

// SetDnsRdata gets a reference to the given string and assigns it to the DnsRdata field.
func (o *DNSUsage) SetDnsRdata(v string) {
	o.DnsRdata = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DNSUsage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSUsage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DNSUsage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DNSUsage) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DNSUsage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSUsage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DNSUsage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DNSUsage) SetName(v string) {
	o.Name = &v
}

// GetRecord returns the Record field value if set, zero value otherwise.
func (o *DNSUsage) GetRecord() string {
	if o == nil || IsNil(o.Record) {
		var ret string
		return ret
	}
	return *o.Record
}

// GetRecordOk returns a tuple with the Record field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSUsage) GetRecordOk() (*string, bool) {
	if o == nil || IsNil(o.Record) {
		return nil, false
	}
	return o.Record, true
}

// HasRecord returns a boolean if a field has been set.
func (o *DNSUsage) HasRecord() bool {
	if o != nil && !IsNil(o.Record) {
		return true
	}

	return false
}

// SetRecord gets a reference to the given string and assigns it to the Record field.
func (o *DNSUsage) SetRecord(v string) {
	o.Record = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *DNSUsage) GetSpace() string {
	if o == nil || IsNil(o.Space) {
		var ret string
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSUsage) GetSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *DNSUsage) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given string and assigns it to the Space field.
func (o *DNSUsage) SetSpace(v string) {
	o.Space = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DNSUsage) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSUsage) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DNSUsage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DNSUsage) SetType(v string) {
	o.Type = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *DNSUsage) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSUsage) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *DNSUsage) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *DNSUsage) SetView(v string) {
	o.View = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *DNSUsage) GetZone() string {
	if o == nil || IsNil(o.Zone) {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSUsage) GetZoneOk() (*string, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *DNSUsage) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *DNSUsage) SetZone(v string) {
	o.Zone = &v
}

func (o DNSUsage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbsoluteName) {
		toSerialize["absolute_name"] = o.AbsoluteName
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.DnsRdata) {
		toSerialize["dns_rdata"] = o.DnsRdata
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Record) {
		toSerialize["record"] = o.Record
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DNSUsage) UnmarshalJSON(data []byte) (err error) {
	varDNSUsage := _DNSUsage{}

	err = json.Unmarshal(data, &varDNSUsage)

	if err != nil {
		return err
	}

	*o = DNSUsage(varDNSUsage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "absolute_name")
		delete(additionalProperties, "address")
		delete(additionalProperties, "dns_rdata")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "record")
		delete(additionalProperties, "space")
		delete(additionalProperties, "type")
		delete(additionalProperties, "view")
		delete(additionalProperties, "zone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDNSUsage struct {
	value *DNSUsage
	isSet bool
}

func (v NullableDNSUsage) Get() *DNSUsage {
	return v.value
}

func (v *NullableDNSUsage) Set(val *DNSUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSUsage(val *DNSUsage) *NullableDNSUsage {
	return &NullableDNSUsage{value: val, isSet: true}
}

func (v NullableDNSUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
