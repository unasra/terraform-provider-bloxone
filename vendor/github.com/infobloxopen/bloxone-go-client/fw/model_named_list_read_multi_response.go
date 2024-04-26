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

<<<<<<<< HEAD:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_security_policy_rule_multi_response.go
// checks if the SecurityPolicyRuleMultiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityPolicyRuleMultiResponse{}

// SecurityPolicyRuleMultiResponse The Security Policy Rule list response.
type SecurityPolicyRuleMultiResponse struct {
	// The list of Security Policy Rule objects.
	Results []SecurityPolicyRule `json:"results,omitempty"`
}

// NewSecurityPolicyRuleMultiResponse instantiates a new SecurityPolicyRuleMultiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPolicyRuleMultiResponse() *SecurityPolicyRuleMultiResponse {
	this := SecurityPolicyRuleMultiResponse{}
	return &this
}

// NewSecurityPolicyRuleMultiResponseWithDefaults instantiates a new SecurityPolicyRuleMultiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPolicyRuleMultiResponseWithDefaults() *SecurityPolicyRuleMultiResponse {
	this := SecurityPolicyRuleMultiResponse{}
========
// checks if the NamedListReadMultiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamedListReadMultiResponse{}

// NamedListReadMultiResponse The Named List list response.
type NamedListReadMultiResponse struct {
	// The list of Named List objects.
	Results              []NamedListRead `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NamedListReadMultiResponse NamedListReadMultiResponse

// NewNamedListReadMultiResponse instantiates a new NamedListReadMultiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamedListReadMultiResponse() *NamedListReadMultiResponse {
	this := NamedListReadMultiResponse{}
	return &this
}

// NewNamedListReadMultiResponseWithDefaults instantiates a new NamedListReadMultiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamedListReadMultiResponseWithDefaults() *NamedListReadMultiResponse {
	this := NamedListReadMultiResponse{}
>>>>>>>> 2be6b8d0d1f652a60c6afe42a36d891d3d0d27d7:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_named_list_read_multi_response.go
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
<<<<<<<< HEAD:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_security_policy_rule_multi_response.go
func (o *SecurityPolicyRuleMultiResponse) GetResults() []SecurityPolicyRule {
	if o == nil || IsNil(o.Results) {
		var ret []SecurityPolicyRule
========
func (o *NamedListReadMultiResponse) GetResults() []NamedListRead {
	if o == nil || IsNil(o.Results) {
		var ret []NamedListRead
>>>>>>>> 2be6b8d0d1f652a60c6afe42a36d891d3d0d27d7:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_named_list_read_multi_response.go
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<<< HEAD:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_security_policy_rule_multi_response.go
func (o *SecurityPolicyRuleMultiResponse) GetResultsOk() ([]SecurityPolicyRule, bool) {
========
func (o *NamedListReadMultiResponse) GetResultsOk() ([]NamedListRead, bool) {
>>>>>>>> 2be6b8d0d1f652a60c6afe42a36d891d3d0d27d7:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_named_list_read_multi_response.go
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
<<<<<<<< HEAD:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_security_policy_rule_multi_response.go
func (o *SecurityPolicyRuleMultiResponse) HasResults() bool {
========
func (o *NamedListReadMultiResponse) HasResults() bool {
>>>>>>>> 2be6b8d0d1f652a60c6afe42a36d891d3d0d27d7:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_named_list_read_multi_response.go
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

<<<<<<<< HEAD:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_security_policy_rule_multi_response.go
// SetResults gets a reference to the given []SecurityPolicyRule and assigns it to the Results field.
func (o *SecurityPolicyRuleMultiResponse) SetResults(v []SecurityPolicyRule) {
	o.Results = v
}

func (o SecurityPolicyRuleMultiResponse) MarshalJSON() ([]byte, error) {
========
// SetResults gets a reference to the given []NamedListRead and assigns it to the Results field.
func (o *NamedListReadMultiResponse) SetResults(v []NamedListRead) {
	o.Results = v
}

func (o NamedListReadMultiResponse) MarshalJSON() ([]byte, error) {
>>>>>>>> 2be6b8d0d1f652a60c6afe42a36d891d3d0d27d7:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_named_list_read_multi_response.go
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

<<<<<<<< HEAD:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_security_policy_rule_multi_response.go
func (o SecurityPolicyRuleMultiResponse) ToMap() (map[string]interface{}, error) {
========
func (o NamedListReadMultiResponse) ToMap() (map[string]interface{}, error) {
>>>>>>>> 2be6b8d0d1f652a60c6afe42a36d891d3d0d27d7:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_named_list_read_multi_response.go
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

<<<<<<<< HEAD:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_security_policy_rule_multi_response.go
type NullableSecurityPolicyRuleMultiResponse struct {
	value *SecurityPolicyRuleMultiResponse
	isSet bool
}

func (v NullableSecurityPolicyRuleMultiResponse) Get() *SecurityPolicyRuleMultiResponse {
	return v.value
}

func (v *NullableSecurityPolicyRuleMultiResponse) Set(val *SecurityPolicyRuleMultiResponse) {
========
func (o *NamedListReadMultiResponse) UnmarshalJSON(data []byte) (err error) {
	varNamedListReadMultiResponse := _NamedListReadMultiResponse{}

	err = json.Unmarshal(data, &varNamedListReadMultiResponse)

	if err != nil {
		return err
	}

	*o = NamedListReadMultiResponse(varNamedListReadMultiResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNamedListReadMultiResponse struct {
	value *NamedListReadMultiResponse
	isSet bool
}

func (v NullableNamedListReadMultiResponse) Get() *NamedListReadMultiResponse {
	return v.value
}

func (v *NullableNamedListReadMultiResponse) Set(val *NamedListReadMultiResponse) {
>>>>>>>> 2be6b8d0d1f652a60c6afe42a36d891d3d0d27d7:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_named_list_read_multi_response.go
	v.value = val
	v.isSet = true
}

<<<<<<<< HEAD:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_security_policy_rule_multi_response.go
func (v NullableSecurityPolicyRuleMultiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPolicyRuleMultiResponse) Unset() {
========
func (v NullableNamedListReadMultiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedListReadMultiResponse) Unset() {
>>>>>>>> 2be6b8d0d1f652a60c6afe42a36d891d3d0d27d7:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_named_list_read_multi_response.go
	v.value = nil
	v.isSet = false
}

<<<<<<<< HEAD:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_security_policy_rule_multi_response.go
func NewNullableSecurityPolicyRuleMultiResponse(val *SecurityPolicyRuleMultiResponse) *NullableSecurityPolicyRuleMultiResponse {
	return &NullableSecurityPolicyRuleMultiResponse{value: val, isSet: true}
}

func (v NullableSecurityPolicyRuleMultiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPolicyRuleMultiResponse) UnmarshalJSON(src []byte) error {
========
func NewNullableNamedListReadMultiResponse(val *NamedListReadMultiResponse) *NullableNamedListReadMultiResponse {
	return &NullableNamedListReadMultiResponse{value: val, isSet: true}
}

func (v NullableNamedListReadMultiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedListReadMultiResponse) UnmarshalJSON(src []byte) error {
>>>>>>>> 2be6b8d0d1f652a60c6afe42a36d891d3d0d27d7:vendor/github.com/infobloxopen/bloxone-go-client/fw/model_named_list_read_multi_response.go
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
