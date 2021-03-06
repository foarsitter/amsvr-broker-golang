/*
AMSVR Broker API

# Introductie     Voor u heeft u de online documentatie van de AMSVR Broker API. Omdat het een Nederlands domein      betreft worden er Engelse en Nederlandse terminologie door elkaar heen gebruikt.  Liefhebbers van de Swagger UI kunnen [hier terecht](/swagger-ui). Daarnaast is de API eveneens beschikbaar in de op endpoint niveau zoals bijvoorbeeld [hier](/api/aansluitingen).  ##        

API version: release-0.6.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package broker

import (
	"encoding/json"
)

// PaginatedAansluitingBestandList struct for PaginatedAansluitingBestandList
type PaginatedAansluitingBestandList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []AansluitingBestand `json:"results,omitempty"`
}

// NewPaginatedAansluitingBestandList instantiates a new PaginatedAansluitingBestandList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAansluitingBestandList() *PaginatedAansluitingBestandList {
	this := PaginatedAansluitingBestandList{}
	return &this
}

// NewPaginatedAansluitingBestandListWithDefaults instantiates a new PaginatedAansluitingBestandList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAansluitingBestandListWithDefaults() *PaginatedAansluitingBestandList {
	this := PaginatedAansluitingBestandList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedAansluitingBestandList) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedAansluitingBestandList) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedAansluitingBestandList) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedAansluitingBestandList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedAansluitingBestandList) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedAansluitingBestandList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedAansluitingBestandList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedAansluitingBestandList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedAansluitingBestandList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedAansluitingBestandList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedAansluitingBestandList) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedAansluitingBestandList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedAansluitingBestandList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedAansluitingBestandList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedAansluitingBestandList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedAansluitingBestandList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedAansluitingBestandList) GetResults() []AansluitingBestand {
	if o == nil || o.Results == nil {
		var ret []AansluitingBestand
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedAansluitingBestandList) GetResultsOk() ([]AansluitingBestand, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedAansluitingBestandList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []AansluitingBestand and assigns it to the Results field.
func (o *PaginatedAansluitingBestandList) SetResults(v []AansluitingBestand) {
	o.Results = v
}

func (o PaginatedAansluitingBestandList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedAansluitingBestandList struct {
	value *PaginatedAansluitingBestandList
	isSet bool
}

func (v NullablePaginatedAansluitingBestandList) Get() *PaginatedAansluitingBestandList {
	return v.value
}

func (v *NullablePaginatedAansluitingBestandList) Set(val *PaginatedAansluitingBestandList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAansluitingBestandList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAansluitingBestandList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAansluitingBestandList(val *PaginatedAansluitingBestandList) *NullablePaginatedAansluitingBestandList {
	return &NullablePaginatedAansluitingBestandList{value: val, isSet: true}
}

func (v NullablePaginatedAansluitingBestandList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAansluitingBestandList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


