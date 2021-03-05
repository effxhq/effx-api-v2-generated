/*
 * effx API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConfigurationFile struct for ConfigurationFile
type ConfigurationFile struct {
	// An Effx Yaml file serialized into a string
	FileContents string `json:"fileContents"`
	Tags *map[string]string `json:"tags,omitempty"`
	// Attach metadata to resources like teams and services
	Annotations *map[string]string `json:"annotations,omitempty"`
}

// NewConfigurationFile instantiates a new ConfigurationFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationFile(fileContents string) *ConfigurationFile {
	this := ConfigurationFile{}
	this.FileContents = fileContents
	return &this
}

// NewConfigurationFileWithDefaults instantiates a new ConfigurationFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationFileWithDefaults() *ConfigurationFile {
	this := ConfigurationFile{}
	return &this
}

// GetFileContents returns the FileContents field value
func (o *ConfigurationFile) GetFileContents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileContents
}

// GetFileContentsOk returns a tuple with the FileContents field value
// and a boolean to check if the value has been set.
func (o *ConfigurationFile) GetFileContentsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FileContents, true
}

// SetFileContents sets field value
func (o *ConfigurationFile) SetFileContents(v string) {
	o.FileContents = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfigurationFile) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationFile) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfigurationFile) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *ConfigurationFile) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ConfigurationFile) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationFile) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ConfigurationFile) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *ConfigurationFile) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

func (o ConfigurationFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fileContents"] = o.FileContents
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationFile struct {
	value *ConfigurationFile
	isSet bool
}

func (v NullableConfigurationFile) Get() *ConfigurationFile {
	return v.value
}

func (v *NullableConfigurationFile) Set(val *ConfigurationFile) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationFile) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationFile(val *ConfigurationFile) *NullableConfigurationFile {
	return &NullableConfigurationFile{value: val, isSet: true}
}

func (v NullableConfigurationFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


