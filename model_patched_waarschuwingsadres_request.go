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

// PatchedWaarschuwingsadresRequest struct for PatchedWaarschuwingsadresRequest
type PatchedWaarschuwingsadresRequest struct {
	Telefoonnummer *int32 `json:"telefoonnummer,omitempty"`
	// De contactpersoon met de hoogste prioriteit wordt als eerste opgeroepen
	Prioriteit *int32 `json:"prioriteit,omitempty"`
}

// NewPatchedWaarschuwingsadresRequest instantiates a new PatchedWaarschuwingsadresRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWaarschuwingsadresRequest() *PatchedWaarschuwingsadresRequest {
	this := PatchedWaarschuwingsadresRequest{}
	return &this
}

// NewPatchedWaarschuwingsadresRequestWithDefaults instantiates a new PatchedWaarschuwingsadresRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWaarschuwingsadresRequestWithDefaults() *PatchedWaarschuwingsadresRequest {
	this := PatchedWaarschuwingsadresRequest{}
	return &this
}

// GetTelefoonnummer returns the Telefoonnummer field value if set, zero value otherwise.
func (o *PatchedWaarschuwingsadresRequest) GetTelefoonnummer() int32 {
	if o == nil || o.Telefoonnummer == nil {
		var ret int32
		return ret
	}
	return *o.Telefoonnummer
}

// GetTelefoonnummerOk returns a tuple with the Telefoonnummer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWaarschuwingsadresRequest) GetTelefoonnummerOk() (*int32, bool) {
	if o == nil || o.Telefoonnummer == nil {
		return nil, false
	}
	return o.Telefoonnummer, true
}

// HasTelefoonnummer returns a boolean if a field has been set.
func (o *PatchedWaarschuwingsadresRequest) HasTelefoonnummer() bool {
	if o != nil && o.Telefoonnummer != nil {
		return true
	}

	return false
}

// SetTelefoonnummer gets a reference to the given int32 and assigns it to the Telefoonnummer field.
func (o *PatchedWaarschuwingsadresRequest) SetTelefoonnummer(v int32) {
	o.Telefoonnummer = &v
}

// GetPrioriteit returns the Prioriteit field value if set, zero value otherwise.
func (o *PatchedWaarschuwingsadresRequest) GetPrioriteit() int32 {
	if o == nil || o.Prioriteit == nil {
		var ret int32
		return ret
	}
	return *o.Prioriteit
}

// GetPrioriteitOk returns a tuple with the Prioriteit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWaarschuwingsadresRequest) GetPrioriteitOk() (*int32, bool) {
	if o == nil || o.Prioriteit == nil {
		return nil, false
	}
	return o.Prioriteit, true
}

// HasPrioriteit returns a boolean if a field has been set.
func (o *PatchedWaarschuwingsadresRequest) HasPrioriteit() bool {
	if o != nil && o.Prioriteit != nil {
		return true
	}

	return false
}

// SetPrioriteit gets a reference to the given int32 and assigns it to the Prioriteit field.
func (o *PatchedWaarschuwingsadresRequest) SetPrioriteit(v int32) {
	o.Prioriteit = &v
}

func (o PatchedWaarschuwingsadresRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Telefoonnummer != nil {
		toSerialize["telefoonnummer"] = o.Telefoonnummer
	}
	if o.Prioriteit != nil {
		toSerialize["prioriteit"] = o.Prioriteit
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWaarschuwingsadresRequest struct {
	value *PatchedWaarschuwingsadresRequest
	isSet bool
}

func (v NullablePatchedWaarschuwingsadresRequest) Get() *PatchedWaarschuwingsadresRequest {
	return v.value
}

func (v *NullablePatchedWaarschuwingsadresRequest) Set(val *PatchedWaarschuwingsadresRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWaarschuwingsadresRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWaarschuwingsadresRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWaarschuwingsadresRequest(val *PatchedWaarschuwingsadresRequest) *NullablePatchedWaarschuwingsadresRequest {
	return &NullablePatchedWaarschuwingsadresRequest{value: val, isSet: true}
}

func (v NullablePatchedWaarschuwingsadresRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWaarschuwingsadresRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


