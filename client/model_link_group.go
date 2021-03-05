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

// LinkGroup struct for LinkGroup
type LinkGroup struct {
	// A value to display that the top of the group of links
	Label string `json:"label"`
	Links []Link `json:"links"`
}

// NewLinkGroup instantiates a new LinkGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkGroup(label string, links []Link) *LinkGroup {
	this := LinkGroup{}
	this.Label = label
	this.Links = links
	return &this
}

// NewLinkGroupWithDefaults instantiates a new LinkGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkGroupWithDefaults() *LinkGroup {
	this := LinkGroup{}
	return &this
}

// GetLabel returns the Label field value
func (o *LinkGroup) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *LinkGroup) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *LinkGroup) SetLabel(v string) {
	o.Label = v
}

// GetLinks returns the Links field value
func (o *LinkGroup) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *LinkGroup) GetLinksOk() (*[]Link, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *LinkGroup) SetLinks(v []Link) {
	o.Links = v
}

func (o LinkGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableLinkGroup struct {
	value *LinkGroup
	isSet bool
}

func (v NullableLinkGroup) Get() *LinkGroup {
	return v.value
}

func (v *NullableLinkGroup) Set(val *LinkGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkGroup(val *LinkGroup) *NullableLinkGroup {
	return &NullableLinkGroup{value: val, isSet: true}
}

func (v NullableLinkGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

