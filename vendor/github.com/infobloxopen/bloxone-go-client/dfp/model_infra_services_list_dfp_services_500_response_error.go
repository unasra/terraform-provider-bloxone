/*
DFP API

BloxOne Cloud is a SaaS offering designed to provide protection to devices on and off-premises, including roaming, remote, and branch offices. It provides visibility into infected and compromised devices, prevents DNS-based data exfiltration, and automatically stops device communications with command-and-control servers (C&Cs) and botnets, in addition to providing recursive DNS services in the cloud. You can access the services by deploying the BloxOne Endpoint agent or the DNS forwarding proxy.  For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries.  By implementing the DNS forwarding proxy, you can rest assured that BloxOne Cloud effectively enforces DNS client-based security policies at your remote sites. On-premises devices that send DNS queries reveal their actual client IP addresses (instead of their NAT IP address), which allows BloxOne Cloud to apply the security policies applicable to the respective endpoints and identify infected clients.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfp

import (
	"encoding/json"
)

// checks if the InfraServicesListDfpServices500ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraServicesListDfpServices500ResponseError{}

// InfraServicesListDfpServices500ResponseError struct for InfraServicesListDfpServices500ResponseError
type InfraServicesListDfpServices500ResponseError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *string `json:"status,omitempty"`
}

// NewInfraServicesListDfpServices500ResponseError instantiates a new InfraServicesListDfpServices500ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraServicesListDfpServices500ResponseError() *InfraServicesListDfpServices500ResponseError {
	this := InfraServicesListDfpServices500ResponseError{}
	return &this
}

// NewInfraServicesListDfpServices500ResponseErrorWithDefaults instantiates a new InfraServicesListDfpServices500ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraServicesListDfpServices500ResponseErrorWithDefaults() *InfraServicesListDfpServices500ResponseError {
	this := InfraServicesListDfpServices500ResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InfraServicesListDfpServices500ResponseError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraServicesListDfpServices500ResponseError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InfraServicesListDfpServices500ResponseError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InfraServicesListDfpServices500ResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InfraServicesListDfpServices500ResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraServicesListDfpServices500ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InfraServicesListDfpServices500ResponseError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InfraServicesListDfpServices500ResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InfraServicesListDfpServices500ResponseError) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraServicesListDfpServices500ResponseError) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InfraServicesListDfpServices500ResponseError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InfraServicesListDfpServices500ResponseError) SetStatus(v string) {
	o.Status = &v
}

func (o InfraServicesListDfpServices500ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraServicesListDfpServices500ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableInfraServicesListDfpServices500ResponseError struct {
	value *InfraServicesListDfpServices500ResponseError
	isSet bool
}

func (v NullableInfraServicesListDfpServices500ResponseError) Get() *InfraServicesListDfpServices500ResponseError {
	return v.value
}

func (v *NullableInfraServicesListDfpServices500ResponseError) Set(val *InfraServicesListDfpServices500ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraServicesListDfpServices500ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraServicesListDfpServices500ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraServicesListDfpServices500ResponseError(val *InfraServicesListDfpServices500ResponseError) *NullableInfraServicesListDfpServices500ResponseError {
	return &NullableInfraServicesListDfpServices500ResponseError{value: val, isSet: true}
}

func (v NullableInfraServicesListDfpServices500ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraServicesListDfpServices500ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}