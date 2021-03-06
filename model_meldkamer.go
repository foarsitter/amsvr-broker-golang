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

// Meldkamer struct for Meldkamer
type Meldkamer struct {
	Id int32 `json:"id"`
	Naam string `json:"naam"`
	Afkorting NullableString `json:"afkorting,omitempty"`
}

// NewMeldkamer instantiates a new Meldkamer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeldkamer(id int32, naam string) *Meldkamer {
	this := Meldkamer{}
	this.Id = id
	this.Naam = naam
	return &this
}

// NewMeldkamerWithDefaults instantiates a new Meldkamer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeldkamerWithDefaults() *Meldkamer {
	this := Meldkamer{}
	return &this
}

// GetId returns the Id field value
func (o *Meldkamer) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Meldkamer) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Meldkamer) SetId(v int32) {
	o.Id = v
}

// GetNaam returns the Naam field value
func (o *Meldkamer) GetNaam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Naam
}

// GetNaamOk returns a tuple with the Naam field value
// and a boolean to check if the value has been set.
func (o *Meldkamer) GetNaamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Naam, true
}

// SetNaam sets field value
func (o *Meldkamer) SetNaam(v string) {
	o.Naam = v
}

// GetAfkorting returns the Afkorting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Meldkamer) GetAfkorting() string {
	if o == nil || o.Afkorting.Get() == nil {
		var ret string
		return ret
	}
	return *o.Afkorting.Get()
}

// GetAfkortingOk returns a tuple with the Afkorting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Meldkamer) GetAfkortingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Afkorting.Get(), o.Afkorting.IsSet()
}

// HasAfkorting returns a boolean if a field has been set.
func (o *Meldkamer) HasAfkorting() bool {
	if o != nil && o.Afkorting.IsSet() {
		return true
	}

	return false
}

// SetAfkorting gets a reference to the given NullableString and assigns it to the Afkorting field.
func (o *Meldkamer) SetAfkorting(v string) {
	o.Afkorting.Set(&v)
}
// SetAfkortingNil sets the value for Afkorting to be an explicit nil
func (o *Meldkamer) SetAfkortingNil() {
	o.Afkorting.Set(nil)
}

// UnsetAfkorting ensures that no value is present for Afkorting, not even an explicit nil
func (o *Meldkamer) UnsetAfkorting() {
	o.Afkorting.Unset()
}

func (o Meldkamer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["naam"] = o.Naam
	}
	if o.Afkorting.IsSet() {
		toSerialize["afkorting"] = o.Afkorting.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMeldkamer struct {
	value *Meldkamer
	isSet bool
}

func (v NullableMeldkamer) Get() *Meldkamer {
	return v.value
}

func (v *NullableMeldkamer) Set(val *Meldkamer) {
	v.value = val
	v.isSet = true
}

func (v NullableMeldkamer) IsSet() bool {
	return v.isSet
}

func (v *NullableMeldkamer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeldkamer(val *Meldkamer) *NullableMeldkamer {
	return &NullableMeldkamer{value: val, isSet: true}
}

func (v NullableMeldkamer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeldkamer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


