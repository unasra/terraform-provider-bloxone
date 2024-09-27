/*
Discovery Configuration API V2

The Discovery configuration service is a BloxOne Service that provides configuration for accessing and syncing the Cloud assets   Base Paths:  1. provider: **_/api/cloud_discovery/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clouddiscovery

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the DiscoveryConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveryConfig{}

// DiscoveryConfig Discovery configuration
type DiscoveryConfig struct {
	// Account preference. For ex.: single, multiple, auto-discover-multiple.
	AccountPreference string `json:"account_preference"`
	// Additional configuration. Ex.: '{    \"excluded_object_types\": [],    \"exclusion_account_list\": [],    \"zone_forwarding\": \"true\" or \"false\" }'.
	AdditionalConfig *AdditionalConfig `json:"additional_config,omitempty"`
	// Timestamp when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Credential preference. Ex.: '{    \"type\": \"static\" or \"delegated\",    \"access_identifier_type\": \"role_arn\" or \"tenant_id\" or \"project_id\"  }'.
	CredentialPreference *CredentialPreference `json:"credential_preference,omitempty"`
	// Timestamp when the object has been deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Description of the discovery config. Optional.
	Description *string `json:"description,omitempty"`
	// Desired state. Default is \"enabled\".
	DesiredState *string `json:"desired_state,omitempty"`
	// Destinations types enabled: Ex.: DNS, IPAM and ACCOUNT.
	DestinationTypesEnabled []string `json:"destination_types_enabled,omitempty"`
	// Destinations.
	Destinations []Destination `json:"destinations,omitempty"`
	// Auto-generated unique discovery config ID. Format BloxID.
	Id *string `json:"id,omitempty"`
	// Last sync timestamp.
	LastSync *time.Time `json:"last_sync,omitempty"`
	// Name of the discovery config.
	Name string `json:"name"`
	// Provider type. Ex.: Amazon Web Services, Google Cloud Platform, Microsoft Azure.
	ProviderType string `json:"provider_type"`
	// Source configs.
	SourceConfigs []SourceConfig `json:"source_configs,omitempty"`
	// Status of the sync operation. In single account case, Its the combined status of account & all the destinations statuses In auto discover case, Its the status of the account discovery only.
	Status *string `json:"status,omitempty"`
	// Aggregate status message of the sync operation.
	StatusMessage *string `json:"status_message,omitempty"`
	SyncInterval  *string `json:"sync_interval,omitempty"`
	// Tagging specifics.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Timestamp when the object has been updated.
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveryConfig DiscoveryConfig

// NewDiscoveryConfig instantiates a new DiscoveryConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryConfig(accountPreference string, name string, providerType string) *DiscoveryConfig {
	this := DiscoveryConfig{}
	this.AccountPreference = accountPreference
	this.Name = name
	this.ProviderType = providerType
	return &this
}

// NewDiscoveryConfigWithDefaults instantiates a new DiscoveryConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryConfigWithDefaults() *DiscoveryConfig {
	this := DiscoveryConfig{}
	return &this
}

// GetAccountPreference returns the AccountPreference field value
func (o *DiscoveryConfig) GetAccountPreference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountPreference
}

// GetAccountPreferenceOk returns a tuple with the AccountPreference field value
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetAccountPreferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountPreference, true
}

// SetAccountPreference sets field value
func (o *DiscoveryConfig) SetAccountPreference(v string) {
	o.AccountPreference = v
}

// GetAdditionalConfig returns the AdditionalConfig field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetAdditionalConfig() AdditionalConfig {
	if o == nil || IsNil(o.AdditionalConfig) {
		var ret AdditionalConfig
		return ret
	}
	return *o.AdditionalConfig
}

// GetAdditionalConfigOk returns a tuple with the AdditionalConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetAdditionalConfigOk() (*AdditionalConfig, bool) {
	if o == nil || IsNil(o.AdditionalConfig) {
		return nil, false
	}
	return o.AdditionalConfig, true
}

// HasAdditionalConfig returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasAdditionalConfig() bool {
	if o != nil && !IsNil(o.AdditionalConfig) {
		return true
	}

	return false
}

// SetAdditionalConfig gets a reference to the given AdditionalConfig and assigns it to the AdditionalConfig field.
func (o *DiscoveryConfig) SetAdditionalConfig(v AdditionalConfig) {
	o.AdditionalConfig = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DiscoveryConfig) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCredentialPreference returns the CredentialPreference field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetCredentialPreference() CredentialPreference {
	if o == nil || IsNil(o.CredentialPreference) {
		var ret CredentialPreference
		return ret
	}
	return *o.CredentialPreference
}

// GetCredentialPreferenceOk returns a tuple with the CredentialPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetCredentialPreferenceOk() (*CredentialPreference, bool) {
	if o == nil || IsNil(o.CredentialPreference) {
		return nil, false
	}
	return o.CredentialPreference, true
}

// HasCredentialPreference returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasCredentialPreference() bool {
	if o != nil && !IsNil(o.CredentialPreference) {
		return true
	}

	return false
}

// SetCredentialPreference gets a reference to the given CredentialPreference and assigns it to the CredentialPreference field.
func (o *DiscoveryConfig) SetCredentialPreference(v CredentialPreference) {
	o.CredentialPreference = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *DiscoveryConfig) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DiscoveryConfig) SetDescription(v string) {
	o.Description = &v
}

// GetDesiredState returns the DesiredState field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetDesiredState() string {
	if o == nil || IsNil(o.DesiredState) {
		var ret string
		return ret
	}
	return *o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetDesiredStateOk() (*string, bool) {
	if o == nil || IsNil(o.DesiredState) {
		return nil, false
	}
	return o.DesiredState, true
}

// HasDesiredState returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasDesiredState() bool {
	if o != nil && !IsNil(o.DesiredState) {
		return true
	}

	return false
}

// SetDesiredState gets a reference to the given string and assigns it to the DesiredState field.
func (o *DiscoveryConfig) SetDesiredState(v string) {
	o.DesiredState = &v
}

// GetDestinationTypesEnabled returns the DestinationTypesEnabled field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetDestinationTypesEnabled() []string {
	if o == nil || IsNil(o.DestinationTypesEnabled) {
		var ret []string
		return ret
	}
	return o.DestinationTypesEnabled
}

// GetDestinationTypesEnabledOk returns a tuple with the DestinationTypesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetDestinationTypesEnabledOk() ([]string, bool) {
	if o == nil || IsNil(o.DestinationTypesEnabled) {
		return nil, false
	}
	return o.DestinationTypesEnabled, true
}

// HasDestinationTypesEnabled returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasDestinationTypesEnabled() bool {
	if o != nil && !IsNil(o.DestinationTypesEnabled) {
		return true
	}

	return false
}

// SetDestinationTypesEnabled gets a reference to the given []string and assigns it to the DestinationTypesEnabled field.
func (o *DiscoveryConfig) SetDestinationTypesEnabled(v []string) {
	o.DestinationTypesEnabled = v
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetDestinations() []Destination {
	if o == nil || IsNil(o.Destinations) {
		var ret []Destination
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetDestinationsOk() ([]Destination, bool) {
	if o == nil || IsNil(o.Destinations) {
		return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasDestinations() bool {
	if o != nil && !IsNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []Destination and assigns it to the Destinations field.
func (o *DiscoveryConfig) SetDestinations(v []Destination) {
	o.Destinations = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiscoveryConfig) SetId(v string) {
	o.Id = &v
}

// GetLastSync returns the LastSync field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetLastSync() time.Time {
	if o == nil || IsNil(o.LastSync) {
		var ret time.Time
		return ret
	}
	return *o.LastSync
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetLastSyncOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSync) {
		return nil, false
	}
	return o.LastSync, true
}

// HasLastSync returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasLastSync() bool {
	if o != nil && !IsNil(o.LastSync) {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given time.Time and assigns it to the LastSync field.
func (o *DiscoveryConfig) SetLastSync(v time.Time) {
	o.LastSync = &v
}

// GetName returns the Name field value
func (o *DiscoveryConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DiscoveryConfig) SetName(v string) {
	o.Name = v
}

// GetProviderType returns the ProviderType field value
func (o *DiscoveryConfig) GetProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *DiscoveryConfig) SetProviderType(v string) {
	o.ProviderType = v
}

// GetSourceConfigs returns the SourceConfigs field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetSourceConfigs() []SourceConfig {
	if o == nil || IsNil(o.SourceConfigs) {
		var ret []SourceConfig
		return ret
	}
	return o.SourceConfigs
}

// GetSourceConfigsOk returns a tuple with the SourceConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetSourceConfigsOk() ([]SourceConfig, bool) {
	if o == nil || IsNil(o.SourceConfigs) {
		return nil, false
	}
	return o.SourceConfigs, true
}

// HasSourceConfigs returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasSourceConfigs() bool {
	if o != nil && !IsNil(o.SourceConfigs) {
		return true
	}

	return false
}

// SetSourceConfigs gets a reference to the given []SourceConfig and assigns it to the SourceConfigs field.
func (o *DiscoveryConfig) SetSourceConfigs(v []SourceConfig) {
	o.SourceConfigs = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DiscoveryConfig) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *DiscoveryConfig) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetSyncInterval returns the SyncInterval field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetSyncInterval() string {
	if o == nil || IsNil(o.SyncInterval) {
		var ret string
		return ret
	}
	return *o.SyncInterval
}

// GetSyncIntervalOk returns a tuple with the SyncInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetSyncIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.SyncInterval) {
		return nil, false
	}
	return o.SyncInterval, true
}

// HasSyncInterval returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasSyncInterval() bool {
	if o != nil && !IsNil(o.SyncInterval) {
		return true
	}

	return false
}

// SetSyncInterval gets a reference to the given string and assigns it to the SyncInterval field.
func (o *DiscoveryConfig) SetSyncInterval(v string) {
	o.SyncInterval = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *DiscoveryConfig) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DiscoveryConfig) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryConfig) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DiscoveryConfig) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DiscoveryConfig) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DiscoveryConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_preference"] = o.AccountPreference
	if !IsNil(o.AdditionalConfig) {
		toSerialize["additional_config"] = o.AdditionalConfig
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CredentialPreference) {
		toSerialize["credential_preference"] = o.CredentialPreference
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DesiredState) {
		toSerialize["desired_state"] = o.DesiredState
	}
	if !IsNil(o.DestinationTypesEnabled) {
		toSerialize["destination_types_enabled"] = o.DestinationTypesEnabled
	}
	if !IsNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastSync) {
		toSerialize["last_sync"] = o.LastSync
	}
	toSerialize["name"] = o.Name
	toSerialize["provider_type"] = o.ProviderType
	if !IsNil(o.SourceConfigs) {
		toSerialize["source_configs"] = o.SourceConfigs
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["status_message"] = o.StatusMessage
	}
	if !IsNil(o.SyncInterval) {
		toSerialize["sync_interval"] = o.SyncInterval
	}
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

func (o *DiscoveryConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_preference",
		"name",
		"provider_type",
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

	varDiscoveryConfig := _DiscoveryConfig{}

	err = json.Unmarshal(data, &varDiscoveryConfig)

	if err != nil {
		return err
	}

	*o = DiscoveryConfig(varDiscoveryConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_preference")
		delete(additionalProperties, "additional_config")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "credential_preference")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "desired_state")
		delete(additionalProperties, "destination_types_enabled")
		delete(additionalProperties, "destinations")
		delete(additionalProperties, "id")
		delete(additionalProperties, "last_sync")
		delete(additionalProperties, "name")
		delete(additionalProperties, "provider_type")
		delete(additionalProperties, "source_configs")
		delete(additionalProperties, "status")
		delete(additionalProperties, "status_message")
		delete(additionalProperties, "sync_interval")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveryConfig struct {
	value *DiscoveryConfig
	isSet bool
}

func (v NullableDiscoveryConfig) Get() *DiscoveryConfig {
	return v.value
}

func (v *NullableDiscoveryConfig) Set(val *DiscoveryConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryConfig(val *DiscoveryConfig) *NullableDiscoveryConfig {
	return &NullableDiscoveryConfig{value: val, isSet: true}
}

func (v NullableDiscoveryConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
