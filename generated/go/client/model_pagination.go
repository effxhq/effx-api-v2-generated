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

// Pagination struct for Pagination
type Pagination struct {
	// The offset to be used for the next request. When there are no more records, this value will be null.
	NextOffset NullableInt32 `json:"nextOffset,omitempty"`
	// The total number of records.
	TotalResults *int32 `json:"totalResults,omitempty"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination() *Pagination {
	this := Pagination{}
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	return &this
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pagination) GetNextOffset() int32 {
	if o == nil || o.NextOffset.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NextOffset.Get()
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pagination) GetNextOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextOffset.Get(), o.NextOffset.IsSet()
}

// HasNextOffset returns a boolean if a field has been set.
func (o *Pagination) HasNextOffset() bool {
	if o != nil && o.NextOffset.IsSet() {
		return true
	}

	return false
}

// SetNextOffset gets a reference to the given NullableInt32 and assigns it to the NextOffset field.
func (o *Pagination) SetNextOffset(v int32) {
	o.NextOffset.Set(&v)
}
// SetNextOffsetNil sets the value for NextOffset to be an explicit nil
func (o *Pagination) SetNextOffsetNil() {
	o.NextOffset.Set(nil)
}

// UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil
func (o *Pagination) UnsetNextOffset() {
	o.NextOffset.Unset()
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *Pagination) GetTotalResults() int32 {
	if o == nil || o.TotalResults == nil {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetTotalResultsOk() (*int32, bool) {
	if o == nil || o.TotalResults == nil {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *Pagination) HasTotalResults() bool {
	if o != nil && o.TotalResults != nil {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *Pagination) SetTotalResults(v int32) {
	o.TotalResults = &v
}

func (o Pagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NextOffset.IsSet() {
		toSerialize["nextOffset"] = o.NextOffset.Get()
	}
	if o.TotalResults != nil {
		toSerialize["totalResults"] = o.TotalResults
	}
	return json.Marshal(toSerialize)
}

type NullablePagination struct {
	value *Pagination
	isSet bool
}

func (v NullablePagination) Get() *Pagination {
	return v.value
}

func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


