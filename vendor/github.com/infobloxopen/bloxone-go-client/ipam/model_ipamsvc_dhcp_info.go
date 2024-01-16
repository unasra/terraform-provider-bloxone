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

// checks if the IpamsvcDHCPInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcDHCPInfo{}

// IpamsvcDHCPInfo The __DHCPInfo__ object represents the DHCP information associated with an address object.
type IpamsvcDHCPInfo struct {
	// The DHCP host name associated with this client.
	ClientHostname *string `json:"client_hostname,omitempty"`
	// The hardware address associated with this client.
	ClientHwaddr *string `json:"client_hwaddr,omitempty"`
	// The ID associated with this client.
	ClientId *string `json:"client_id,omitempty"`
	// The timestamp at which the _state_, when set to _leased_, will be changed to _free_.
	End *time.Time `json:"end,omitempty"`
	// The DHCP fingerprint for the associated lease.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// Identity Association Identifier (IAID) of the lease. Applicable only for DHCPv6.
	Iaid *int64 `json:"iaid,omitempty"`
	// Lease type. Applicable only for address under DHCP control. The value can be empty for address not under DHCP control.  Valid values are: * _DHCPv6NonTemporaryAddress_: DHCPv6 non-temporary address (NA) * _DHCPv6TemporaryAddress_: DHCPv6 temporary address (TA) * _DHCPv6PrefixDelegation_: DHCPv6 prefix delegation (PD) * _DHCPv4_: DHCPv4 lease
	LeaseType *string `json:"lease_type,omitempty"`
	// The length of time that a valid address is preferred (i.e., the time until deprecation). When the preferred lifetime expires, the address becomes deprecated on the client. It is still considered leased on the server. Applicable only for DHCPv6.
	PreferredLifetime *time.Time `json:"preferred_lifetime,omitempty"`
	// The remaining time, in seconds, until which the _state_, when set to _leased_, will remain in that state.
	Remain *int64 `json:"remain,omitempty"`
	// The timestamp at which _state_ was first set to _leased_.
	Start *time.Time `json:"start,omitempty"`
	// Indicates the status of this IP address from a DHCP protocol standpoint as:   * _none_: Address is not under DHCP control.   * _free_: Address is under DHCP control but has no lease currently assigned.   * _leased_: Address is under DHCP control and has a lease currently assigned. The lease details are contained in the matching _dhcp/lease_ resource.
	State *string `json:"state,omitempty"`
	// The timestamp at which the _state_ was last reported.
	StateTs *time.Time `json:"state_ts,omitempty"`
}

// NewIpamsvcDHCPInfo instantiates a new IpamsvcDHCPInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcDHCPInfo() *IpamsvcDHCPInfo {
	this := IpamsvcDHCPInfo{}
	return &this
}

// NewIpamsvcDHCPInfoWithDefaults instantiates a new IpamsvcDHCPInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcDHCPInfoWithDefaults() *IpamsvcDHCPInfo {
	this := IpamsvcDHCPInfo{}
	return &this
}

// GetClientHostname returns the ClientHostname field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetClientHostname() string {
	if o == nil || IsNil(o.ClientHostname) {
		var ret string
		return ret
	}
	return *o.ClientHostname
}

// GetClientHostnameOk returns a tuple with the ClientHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetClientHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientHostname) {
		return nil, false
	}
	return o.ClientHostname, true
}

// HasClientHostname returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasClientHostname() bool {
	if o != nil && !IsNil(o.ClientHostname) {
		return true
	}

	return false
}

// SetClientHostname gets a reference to the given string and assigns it to the ClientHostname field.
func (o *IpamsvcDHCPInfo) SetClientHostname(v string) {
	o.ClientHostname = &v
}

// GetClientHwaddr returns the ClientHwaddr field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetClientHwaddr() string {
	if o == nil || IsNil(o.ClientHwaddr) {
		var ret string
		return ret
	}
	return *o.ClientHwaddr
}

// GetClientHwaddrOk returns a tuple with the ClientHwaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetClientHwaddrOk() (*string, bool) {
	if o == nil || IsNil(o.ClientHwaddr) {
		return nil, false
	}
	return o.ClientHwaddr, true
}

// HasClientHwaddr returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasClientHwaddr() bool {
	if o != nil && !IsNil(o.ClientHwaddr) {
		return true
	}

	return false
}

// SetClientHwaddr gets a reference to the given string and assigns it to the ClientHwaddr field.
func (o *IpamsvcDHCPInfo) SetClientHwaddr(v string) {
	o.ClientHwaddr = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IpamsvcDHCPInfo) SetClientId(v string) {
	o.ClientId = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetEnd() time.Time {
	if o == nil || IsNil(o.End) {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *IpamsvcDHCPInfo) SetEnd(v time.Time) {
	o.End = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *IpamsvcDHCPInfo) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetIaid returns the Iaid field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetIaid() int64 {
	if o == nil || IsNil(o.Iaid) {
		var ret int64
		return ret
	}
	return *o.Iaid
}

// GetIaidOk returns a tuple with the Iaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetIaidOk() (*int64, bool) {
	if o == nil || IsNil(o.Iaid) {
		return nil, false
	}
	return o.Iaid, true
}

// HasIaid returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasIaid() bool {
	if o != nil && !IsNil(o.Iaid) {
		return true
	}

	return false
}

// SetIaid gets a reference to the given int64 and assigns it to the Iaid field.
func (o *IpamsvcDHCPInfo) SetIaid(v int64) {
	o.Iaid = &v
}

// GetLeaseType returns the LeaseType field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetLeaseType() string {
	if o == nil || IsNil(o.LeaseType) {
		var ret string
		return ret
	}
	return *o.LeaseType
}

// GetLeaseTypeOk returns a tuple with the LeaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetLeaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LeaseType) {
		return nil, false
	}
	return o.LeaseType, true
}

// HasLeaseType returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasLeaseType() bool {
	if o != nil && !IsNil(o.LeaseType) {
		return true
	}

	return false
}

// SetLeaseType gets a reference to the given string and assigns it to the LeaseType field.
func (o *IpamsvcDHCPInfo) SetLeaseType(v string) {
	o.LeaseType = &v
}

// GetPreferredLifetime returns the PreferredLifetime field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetPreferredLifetime() time.Time {
	if o == nil || IsNil(o.PreferredLifetime) {
		var ret time.Time
		return ret
	}
	return *o.PreferredLifetime
}

// GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetPreferredLifetimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PreferredLifetime) {
		return nil, false
	}
	return o.PreferredLifetime, true
}

// HasPreferredLifetime returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasPreferredLifetime() bool {
	if o != nil && !IsNil(o.PreferredLifetime) {
		return true
	}

	return false
}

// SetPreferredLifetime gets a reference to the given time.Time and assigns it to the PreferredLifetime field.
func (o *IpamsvcDHCPInfo) SetPreferredLifetime(v time.Time) {
	o.PreferredLifetime = &v
}

// GetRemain returns the Remain field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetRemain() int64 {
	if o == nil || IsNil(o.Remain) {
		var ret int64
		return ret
	}
	return *o.Remain
}

// GetRemainOk returns a tuple with the Remain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetRemainOk() (*int64, bool) {
	if o == nil || IsNil(o.Remain) {
		return nil, false
	}
	return o.Remain, true
}

// HasRemain returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasRemain() bool {
	if o != nil && !IsNil(o.Remain) {
		return true
	}

	return false
}

// SetRemain gets a reference to the given int64 and assigns it to the Remain field.
func (o *IpamsvcDHCPInfo) SetRemain(v int64) {
	o.Remain = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetStart() time.Time {
	if o == nil || IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *IpamsvcDHCPInfo) SetStart(v time.Time) {
	o.Start = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IpamsvcDHCPInfo) SetState(v string) {
	o.State = &v
}

// GetStateTs returns the StateTs field value if set, zero value otherwise.
func (o *IpamsvcDHCPInfo) GetStateTs() time.Time {
	if o == nil || IsNil(o.StateTs) {
		var ret time.Time
		return ret
	}
	return *o.StateTs
}

// GetStateTsOk returns a tuple with the StateTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPInfo) GetStateTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StateTs) {
		return nil, false
	}
	return o.StateTs, true
}

// HasStateTs returns a boolean if a field has been set.
func (o *IpamsvcDHCPInfo) HasStateTs() bool {
	if o != nil && !IsNil(o.StateTs) {
		return true
	}

	return false
}

// SetStateTs gets a reference to the given time.Time and assigns it to the StateTs field.
func (o *IpamsvcDHCPInfo) SetStateTs(v time.Time) {
	o.StateTs = &v
}

func (o IpamsvcDHCPInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcDHCPInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientHostname) {
		toSerialize["client_hostname"] = o.ClientHostname
	}
	if !IsNil(o.ClientHwaddr) {
		toSerialize["client_hwaddr"] = o.ClientHwaddr
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if !IsNil(o.Iaid) {
		toSerialize["iaid"] = o.Iaid
	}
	if !IsNil(o.LeaseType) {
		toSerialize["lease_type"] = o.LeaseType
	}
	if !IsNil(o.PreferredLifetime) {
		toSerialize["preferred_lifetime"] = o.PreferredLifetime
	}
	if !IsNil(o.Remain) {
		toSerialize["remain"] = o.Remain
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StateTs) {
		toSerialize["state_ts"] = o.StateTs
	}
	return toSerialize, nil
}

type NullableIpamsvcDHCPInfo struct {
	value *IpamsvcDHCPInfo
	isSet bool
}

func (v NullableIpamsvcDHCPInfo) Get() *IpamsvcDHCPInfo {
	return v.value
}

func (v *NullableIpamsvcDHCPInfo) Set(val *IpamsvcDHCPInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcDHCPInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcDHCPInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcDHCPInfo(val *IpamsvcDHCPInfo) *NullableIpamsvcDHCPInfo {
	return &NullableIpamsvcDHCPInfo{value: val, isSet: true}
}

func (v NullableIpamsvcDHCPInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcDHCPInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
