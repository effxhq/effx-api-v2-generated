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

// ServiceConfiguration struct for ServiceConfiguration
type ServiceConfiguration struct {
	Version string `json:"version"`
	Kind string `json:"kind"`
	Spec ServiceConfigurationSpec `json:"spec"`
}

// NewServiceConfiguration instantiates a new ServiceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceConfiguration(version string, kind string, spec ServiceConfigurationSpec) *ServiceConfiguration {
	this := ServiceConfiguration{}
	this.Version = version
	this.Kind = kind
	this.Spec = spec
	return &this
}

// NewServiceConfigurationWithDefaults instantiates a new ServiceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceConfigurationWithDefaults() *ServiceConfiguration {
	this := ServiceConfiguration{}
	return &this
}

// GetVersion returns the Version field value
func (o *ServiceConfiguration) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ServiceConfiguration) SetVersion(v string) {
	o.Version = v
}

// GetKind returns the Kind field value
func (o *ServiceConfiguration) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ServiceConfiguration) SetKind(v string) {
	o.Kind = v
}

// GetSpec returns the Spec field value
func (o *ServiceConfiguration) GetSpec() ServiceConfigurationSpec {
	if o == nil {
		var ret ServiceConfigurationSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *ServiceConfiguration) GetSpecOk() (*ServiceConfigurationSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *ServiceConfiguration) SetSpec(v ServiceConfigurationSpec) {
	o.Spec = v
}

func (o ServiceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableServiceConfiguration struct {
	value *ServiceConfiguration
	isSet bool
}

func (v NullableServiceConfiguration) Get() *ServiceConfiguration {
	return v.value
}

func (v *NullableServiceConfiguration) Set(val *ServiceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceConfiguration(val *ServiceConfiguration) *NullableServiceConfiguration {
	return &NullableServiceConfiguration{value: val, isSet: true}
}

func (v NullableServiceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

