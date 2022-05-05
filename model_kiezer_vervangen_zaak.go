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

// KiezerVervangenZaak struct for KiezerVervangenZaak
type KiezerVervangenZaak struct {
	Description string `json:"description"`
}

// NewKiezerVervangenZaak instantiates a new KiezerVervangenZaak object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKiezerVervangenZaak(description string) *KiezerVervangenZaak {
	this := KiezerVervangenZaak{}
	this.Description = description
	return &this
}

// NewKiezerVervangenZaakWithDefaults instantiates a new KiezerVervangenZaak object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKiezerVervangenZaakWithDefaults() *KiezerVervangenZaak {
	this := KiezerVervangenZaak{}
	return &this
}

// GetDescription returns the Description field value
func (o *KiezerVervangenZaak) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *KiezerVervangenZaak) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *KiezerVervangenZaak) SetDescription(v string) {
	o.Description = v
}

func (o KiezerVervangenZaak) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableKiezerVervangenZaak struct {
	value *KiezerVervangenZaak
	isSet bool
}

func (v NullableKiezerVervangenZaak) Get() *KiezerVervangenZaak {
	return v.value
}

func (v *NullableKiezerVervangenZaak) Set(val *KiezerVervangenZaak) {
	v.value = val
	v.isSet = true
}

func (v NullableKiezerVervangenZaak) IsSet() bool {
	return v.isSet
}

func (v *NullableKiezerVervangenZaak) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKiezerVervangenZaak(val *KiezerVervangenZaak) *NullableKiezerVervangenZaak {
	return &NullableKiezerVervangenZaak{value: val, isSet: true}
}

func (v NullableKiezerVervangenZaak) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKiezerVervangenZaak) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


