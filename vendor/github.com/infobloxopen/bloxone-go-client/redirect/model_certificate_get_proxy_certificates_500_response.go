/*
BloxOne Redirect API

You can configure BloxOne Threat Defense Cloud to redirect traffic to the Infoblox server that displays the default or customized redirect page. You can redirect traffic to a custom destination using custom redirects.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redirect

import (
	"encoding/json"
)

// checks if the CertificateGetProxyCertificates500Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateGetProxyCertificates500Response{}

// CertificateGetProxyCertificates500Response struct for CertificateGetProxyCertificates500Response
type CertificateGetProxyCertificates500Response struct {
	Error                *CertificateGetProxyCertificates500ResponseError `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateGetProxyCertificates500Response CertificateGetProxyCertificates500Response

// NewCertificateGetProxyCertificates500Response instantiates a new CertificateGetProxyCertificates500Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateGetProxyCertificates500Response() *CertificateGetProxyCertificates500Response {
	this := CertificateGetProxyCertificates500Response{}
	return &this
}

// NewCertificateGetProxyCertificates500ResponseWithDefaults instantiates a new CertificateGetProxyCertificates500Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateGetProxyCertificates500ResponseWithDefaults() *CertificateGetProxyCertificates500Response {
	this := CertificateGetProxyCertificates500Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CertificateGetProxyCertificates500Response) GetError() CertificateGetProxyCertificates500ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret CertificateGetProxyCertificates500ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateGetProxyCertificates500Response) GetErrorOk() (*CertificateGetProxyCertificates500ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CertificateGetProxyCertificates500Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CertificateGetProxyCertificates500ResponseError and assigns it to the Error field.
func (o *CertificateGetProxyCertificates500Response) SetError(v CertificateGetProxyCertificates500ResponseError) {
	o.Error = &v
}

func (o CertificateGetProxyCertificates500Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateGetProxyCertificates500Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateGetProxyCertificates500Response) UnmarshalJSON(data []byte) (err error) {
	varCertificateGetProxyCertificates500Response := _CertificateGetProxyCertificates500Response{}

	err = json.Unmarshal(data, &varCertificateGetProxyCertificates500Response)

	if err != nil {
		return err
	}

	*o = CertificateGetProxyCertificates500Response(varCertificateGetProxyCertificates500Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateGetProxyCertificates500Response struct {
	value *CertificateGetProxyCertificates500Response
	isSet bool
}

func (v NullableCertificateGetProxyCertificates500Response) Get() *CertificateGetProxyCertificates500Response {
	return v.value
}

func (v *NullableCertificateGetProxyCertificates500Response) Set(val *CertificateGetProxyCertificates500Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateGetProxyCertificates500Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateGetProxyCertificates500Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateGetProxyCertificates500Response(val *CertificateGetProxyCertificates500Response) *NullableCertificateGetProxyCertificates500Response {
	return &NullableCertificateGetProxyCertificates500Response{value: val, isSet: true}
}

func (v NullableCertificateGetProxyCertificates500Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateGetProxyCertificates500Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
