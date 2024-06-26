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

// checks if the CustomRedirectsDeleteCustomRedirect400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomRedirectsDeleteCustomRedirect400Response{}

// CustomRedirectsDeleteCustomRedirect400Response struct for CustomRedirectsDeleteCustomRedirect400Response
type CustomRedirectsDeleteCustomRedirect400Response struct {
	Error                *CustomRedirectsDeleteCustomRedirect400ResponseError `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomRedirectsDeleteCustomRedirect400Response CustomRedirectsDeleteCustomRedirect400Response

// NewCustomRedirectsDeleteCustomRedirect400Response instantiates a new CustomRedirectsDeleteCustomRedirect400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomRedirectsDeleteCustomRedirect400Response() *CustomRedirectsDeleteCustomRedirect400Response {
	this := CustomRedirectsDeleteCustomRedirect400Response{}
	return &this
}

// NewCustomRedirectsDeleteCustomRedirect400ResponseWithDefaults instantiates a new CustomRedirectsDeleteCustomRedirect400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomRedirectsDeleteCustomRedirect400ResponseWithDefaults() *CustomRedirectsDeleteCustomRedirect400Response {
	this := CustomRedirectsDeleteCustomRedirect400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CustomRedirectsDeleteCustomRedirect400Response) GetError() CustomRedirectsDeleteCustomRedirect400ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret CustomRedirectsDeleteCustomRedirect400ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRedirectsDeleteCustomRedirect400Response) GetErrorOk() (*CustomRedirectsDeleteCustomRedirect400ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CustomRedirectsDeleteCustomRedirect400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CustomRedirectsDeleteCustomRedirect400ResponseError and assigns it to the Error field.
func (o *CustomRedirectsDeleteCustomRedirect400Response) SetError(v CustomRedirectsDeleteCustomRedirect400ResponseError) {
	o.Error = &v
}

func (o CustomRedirectsDeleteCustomRedirect400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomRedirectsDeleteCustomRedirect400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomRedirectsDeleteCustomRedirect400Response) UnmarshalJSON(data []byte) (err error) {
	varCustomRedirectsDeleteCustomRedirect400Response := _CustomRedirectsDeleteCustomRedirect400Response{}

	err = json.Unmarshal(data, &varCustomRedirectsDeleteCustomRedirect400Response)

	if err != nil {
		return err
	}

	*o = CustomRedirectsDeleteCustomRedirect400Response(varCustomRedirectsDeleteCustomRedirect400Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomRedirectsDeleteCustomRedirect400Response struct {
	value *CustomRedirectsDeleteCustomRedirect400Response
	isSet bool
}

func (v NullableCustomRedirectsDeleteCustomRedirect400Response) Get() *CustomRedirectsDeleteCustomRedirect400Response {
	return v.value
}

func (v *NullableCustomRedirectsDeleteCustomRedirect400Response) Set(val *CustomRedirectsDeleteCustomRedirect400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomRedirectsDeleteCustomRedirect400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomRedirectsDeleteCustomRedirect400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomRedirectsDeleteCustomRedirect400Response(val *CustomRedirectsDeleteCustomRedirect400Response) *NullableCustomRedirectsDeleteCustomRedirect400Response {
	return &NullableCustomRedirectsDeleteCustomRedirect400Response{value: val, isSet: true}
}

func (v NullableCustomRedirectsDeleteCustomRedirect400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomRedirectsDeleteCustomRedirect400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
