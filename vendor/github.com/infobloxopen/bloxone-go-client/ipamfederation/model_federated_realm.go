/*
IPAM Federation API

The DDI IPAM Federation application enables a SaaS administrator to manage multiple IPAM systems from one central control point CSP.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipamfederation

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the FederatedRealm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederatedRealm{}

// FederatedRealm A __FederatedRealm__ object is a unique set of federated blocks per realm.
type FederatedRealm struct {
	// The aggregate of all Federated Blocks within the Realm.
	AllocationV4 *Allocation `json:"allocation_v4,omitempty"`
	// The description of the federated realm. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The name of the federated realm. May contain 1 to 256 characters; can include UTF-8.
	Name string `json:"name"`
	// The tags for the federated realm in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederatedRealm FederatedRealm

// NewFederatedRealm instantiates a new FederatedRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederatedRealm(name string) *FederatedRealm {
	this := FederatedRealm{}
	this.Name = name
	return &this
}

// NewFederatedRealmWithDefaults instantiates a new FederatedRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederatedRealmWithDefaults() *FederatedRealm {
	this := FederatedRealm{}
	return &this
}

// GetAllocationV4 returns the AllocationV4 field value if set, zero value otherwise.
func (o *FederatedRealm) GetAllocationV4() Allocation {
	if o == nil || IsNil(o.AllocationV4) {
		var ret Allocation
		return ret
	}
	return *o.AllocationV4
}

// GetAllocationV4Ok returns a tuple with the AllocationV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedRealm) GetAllocationV4Ok() (*Allocation, bool) {
	if o == nil || IsNil(o.AllocationV4) {
		return nil, false
	}
	return o.AllocationV4, true
}

// HasAllocationV4 returns a boolean if a field has been set.
func (o *FederatedRealm) HasAllocationV4() bool {
	if o != nil && !IsNil(o.AllocationV4) {
		return true
	}

	return false
}

// SetAllocationV4 gets a reference to the given Allocation and assigns it to the AllocationV4 field.
func (o *FederatedRealm) SetAllocationV4(v Allocation) {
	o.AllocationV4 = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *FederatedRealm) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedRealm) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *FederatedRealm) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *FederatedRealm) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FederatedRealm) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedRealm) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FederatedRealm) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FederatedRealm) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FederatedRealm) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedRealm) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FederatedRealm) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FederatedRealm) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *FederatedRealm) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FederatedRealm) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FederatedRealm) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FederatedRealm) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedRealm) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FederatedRealm) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *FederatedRealm) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FederatedRealm) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedRealm) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FederatedRealm) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FederatedRealm) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o FederatedRealm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederatedRealm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllocationV4) {
		toSerialize["allocation_v4"] = o.AllocationV4
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederatedRealm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFederatedRealm := _FederatedRealm{}

	err = json.Unmarshal(data, &varFederatedRealm)

	if err != nil {
		return err
	}

	*o = FederatedRealm(varFederatedRealm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allocation_v4")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederatedRealm struct {
	value *FederatedRealm
	isSet bool
}

func (v NullableFederatedRealm) Get() *FederatedRealm {
	return v.value
}

func (v *NullableFederatedRealm) Set(val *FederatedRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableFederatedRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableFederatedRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederatedRealm(val *FederatedRealm) *NullableFederatedRealm {
	return &NullableFederatedRealm{value: val, isSet: true}
}

func (v NullableFederatedRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederatedRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
