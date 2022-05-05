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

// AfsluitingZaakTransitieRequest struct for AfsluitingZaakTransitieRequest
type AfsluitingZaakTransitieRequest struct {
	Description string `json:"description"`
}

// NewAfsluitingZaakTransitieRequest instantiates a new AfsluitingZaakTransitieRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfsluitingZaakTransitieRequest(description string) *AfsluitingZaakTransitieRequest {
	this := AfsluitingZaakTransitieRequest{}
	this.Description = description
	return &this
}

// NewAfsluitingZaakTransitieRequestWithDefaults instantiates a new AfsluitingZaakTransitieRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfsluitingZaakTransitieRequestWithDefaults() *AfsluitingZaakTransitieRequest {
	this := AfsluitingZaakTransitieRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *AfsluitingZaakTransitieRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AfsluitingZaakTransitieRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AfsluitingZaakTransitieRequest) SetDescription(v string) {
	o.Description = v
}

func (o AfsluitingZaakTransitieRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAfsluitingZaakTransitieRequest struct {
	value *AfsluitingZaakTransitieRequest
	isSet bool
}

func (v NullableAfsluitingZaakTransitieRequest) Get() *AfsluitingZaakTransitieRequest {
	return v.value
}

func (v *NullableAfsluitingZaakTransitieRequest) Set(val *AfsluitingZaakTransitieRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAfsluitingZaakTransitieRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAfsluitingZaakTransitieRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfsluitingZaakTransitieRequest(val *AfsluitingZaakTransitieRequest) *NullableAfsluitingZaakTransitieRequest {
	return &NullableAfsluitingZaakTransitieRequest{value: val, isSet: true}
}

func (v NullableAfsluitingZaakTransitieRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfsluitingZaakTransitieRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


