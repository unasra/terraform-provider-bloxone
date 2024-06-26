/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
)

// checks if the PoPRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoPRegion{}

// PoPRegion PoPRegion message for a Point of Presence (PoP) region
type PoPRegion struct {
	// PoP Region's IP addresses
	Addresses []string `json:"addresses,omitempty"`
	// The PoP Region's serial, unique ID
	Id *int32 `json:"id,omitempty"`
	// PoP Region's location
	Location *string `json:"location,omitempty"`
	// PoP Region's name
	Region               *string `json:"region,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoPRegion PoPRegion

// NewPoPRegion instantiates a new PoPRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoPRegion() *PoPRegion {
	this := PoPRegion{}
	return &this
}

// NewPoPRegionWithDefaults instantiates a new PoPRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoPRegionWithDefaults() *PoPRegion {
	this := PoPRegion{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *PoPRegion) GetAddresses() []string {
	if o == nil || IsNil(o.Addresses) {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoPRegion) GetAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *PoPRegion) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *PoPRegion) SetAddresses(v []string) {
	o.Addresses = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoPRegion) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoPRegion) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoPRegion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PoPRegion) SetId(v int32) {
	o.Id = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PoPRegion) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoPRegion) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PoPRegion) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *PoPRegion) SetLocation(v string) {
	o.Location = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PoPRegion) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoPRegion) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PoPRegion) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PoPRegion) SetRegion(v string) {
	o.Region = &v
}

func (o PoPRegion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoPRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PoPRegion) UnmarshalJSON(data []byte) (err error) {
	varPoPRegion := _PoPRegion{}

	err = json.Unmarshal(data, &varPoPRegion)

	if err != nil {
		return err
	}

	*o = PoPRegion(varPoPRegion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "addresses")
		delete(additionalProperties, "id")
		delete(additionalProperties, "location")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoPRegion struct {
	value *PoPRegion
	isSet bool
}

func (v NullablePoPRegion) Get() *PoPRegion {
	return v.value
}

func (v *NullablePoPRegion) Set(val *PoPRegion) {
	v.value = val
	v.isSet = true
}

func (v NullablePoPRegion) IsSet() bool {
	return v.isSet
}

func (v *NullablePoPRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoPRegion(val *PoPRegion) *NullablePoPRegion {
	return &NullablePoPRegion{value: val, isSet: true}
}

func (v NullablePoPRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoPRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
