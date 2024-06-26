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

// checks if the ApplicationFiltersDeleteSingleApplicationFilters400ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationFiltersDeleteSingleApplicationFilters400ResponseError{}

// ApplicationFiltersDeleteSingleApplicationFilters400ResponseError struct for ApplicationFiltersDeleteSingleApplicationFilters400ResponseError
type ApplicationFiltersDeleteSingleApplicationFilters400ResponseError struct {
	Code                 *string `json:"code,omitempty"`
	Message              *string `json:"message,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationFiltersDeleteSingleApplicationFilters400ResponseError ApplicationFiltersDeleteSingleApplicationFilters400ResponseError

// NewApplicationFiltersDeleteSingleApplicationFilters400ResponseError instantiates a new ApplicationFiltersDeleteSingleApplicationFilters400ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFiltersDeleteSingleApplicationFilters400ResponseError() *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError {
	this := ApplicationFiltersDeleteSingleApplicationFilters400ResponseError{}
	return &this
}

// NewApplicationFiltersDeleteSingleApplicationFilters400ResponseErrorWithDefaults instantiates a new ApplicationFiltersDeleteSingleApplicationFilters400ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFiltersDeleteSingleApplicationFilters400ResponseErrorWithDefaults() *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError {
	this := ApplicationFiltersDeleteSingleApplicationFilters400ResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) SetStatus(v string) {
	o.Status = &v
}

func (o ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) UnmarshalJSON(data []byte) (err error) {
	varApplicationFiltersDeleteSingleApplicationFilters400ResponseError := _ApplicationFiltersDeleteSingleApplicationFilters400ResponseError{}

	err = json.Unmarshal(data, &varApplicationFiltersDeleteSingleApplicationFilters400ResponseError)

	if err != nil {
		return err
	}

	*o = ApplicationFiltersDeleteSingleApplicationFilters400ResponseError(varApplicationFiltersDeleteSingleApplicationFilters400ResponseError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationFiltersDeleteSingleApplicationFilters400ResponseError struct {
	value *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError
	isSet bool
}

func (v NullableApplicationFiltersDeleteSingleApplicationFilters400ResponseError) Get() *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError {
	return v.value
}

func (v *NullableApplicationFiltersDeleteSingleApplicationFilters400ResponseError) Set(val *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFiltersDeleteSingleApplicationFilters400ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFiltersDeleteSingleApplicationFilters400ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFiltersDeleteSingleApplicationFilters400ResponseError(val *ApplicationFiltersDeleteSingleApplicationFilters400ResponseError) *NullableApplicationFiltersDeleteSingleApplicationFilters400ResponseError {
	return &NullableApplicationFiltersDeleteSingleApplicationFilters400ResponseError{value: val, isSet: true}
}

func (v NullableApplicationFiltersDeleteSingleApplicationFilters400ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFiltersDeleteSingleApplicationFilters400ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
