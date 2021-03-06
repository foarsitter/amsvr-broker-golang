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

// AansluitingBestandRequest struct for AansluitingBestandRequest
type AansluitingBestandRequest struct {
	Naam         *string        `json:"naam,omitempty"`
	Omschrijving NullableString `json:"omschrijving,omitempty"`
	// overige = Overige, pve = PVE, upd = UPD, nva = NVA (nota van aanvulling), nvw = NVW (nota van wijzigingen), gms_mutatierapport = GMS mutatierapport
	Type NullableAansluitingBestandTypeEnum `json:"type,omitempty"`
	File NullableString                     `json:"file,omitempty"`
}

// NewAansluitingBestandRequest instantiates a new AansluitingBestandRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAansluitingBestandRequest() *AansluitingBestandRequest {
	this := AansluitingBestandRequest{}
	return &this
}

// NewAansluitingBestandRequestWithDefaults instantiates a new AansluitingBestandRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAansluitingBestandRequestWithDefaults() *AansluitingBestandRequest {
	this := AansluitingBestandRequest{}
	return &this
}

// GetNaam returns the Naam field value if set, zero value otherwise.
func (o *AansluitingBestandRequest) GetNaam() string {
	if o == nil || o.Naam == nil {
		var ret string
		return ret
	}
	return *o.Naam
}

// GetNaamOk returns a tuple with the Naam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AansluitingBestandRequest) GetNaamOk() (*string, bool) {
	if o == nil || o.Naam == nil {
		return nil, false
	}
	return o.Naam, true
}

// HasNaam returns a boolean if a field has been set.
func (o *AansluitingBestandRequest) HasNaam() bool {
	if o != nil && o.Naam != nil {
		return true
	}

	return false
}

// SetNaam gets a reference to the given string and assigns it to the Naam field.
func (o *AansluitingBestandRequest) SetNaam(v string) {
	o.Naam = &v
}

// GetOmschrijving returns the Omschrijving field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AansluitingBestandRequest) GetOmschrijving() string {
	if o == nil || o.Omschrijving.Get() == nil {
		var ret string
		return ret
	}
	return *o.Omschrijving.Get()
}

// GetOmschrijvingOk returns a tuple with the Omschrijving field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AansluitingBestandRequest) GetOmschrijvingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Omschrijving.Get(), o.Omschrijving.IsSet()
}

// HasOmschrijving returns a boolean if a field has been set.
func (o *AansluitingBestandRequest) HasOmschrijving() bool {
	if o != nil && o.Omschrijving.IsSet() {
		return true
	}

	return false
}

// SetOmschrijving gets a reference to the given NullableString and assigns it to the Omschrijving field.
func (o *AansluitingBestandRequest) SetOmschrijving(v string) {
	o.Omschrijving.Set(&v)
}

// SetOmschrijvingNil sets the value for Omschrijving to be an explicit nil
func (o *AansluitingBestandRequest) SetOmschrijvingNil() {
	o.Omschrijving.Set(nil)
}

// UnsetOmschrijving ensures that no value is present for Omschrijving, not even an explicit nil
func (o *AansluitingBestandRequest) UnsetOmschrijving() {
	o.Omschrijving.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AansluitingBestandRequest) GetType() AansluitingBestandTypeEnum {
	if o == nil || o.Type.Get() == nil {
		var ret AansluitingBestandTypeEnum
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AansluitingBestandRequest) GetTypeOk() (*AansluitingBestandTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *AansluitingBestandRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableAansluitingBestandTypeEnum and assigns it to the Type field.
func (o *AansluitingBestandRequest) SetType(v AansluitingBestandTypeEnum) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *AansluitingBestandRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *AansluitingBestandRequest) UnsetType() {
	o.Type.Unset()
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AansluitingBestandRequest) GetFile() *string {
	if o == nil || o.File.Get() == nil {
		var ret *string
		return ret
	}
	return o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AansluitingBestandRequest) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *AansluitingBestandRequest) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given Nullable*os.File and assigns it to the File field.
func (o *AansluitingBestandRequest) SetFile(v *string) {
	o.File.Set(v)
}

// SetFileNil sets the value for File to be an explicit nil
func (o *AansluitingBestandRequest) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *AansluitingBestandRequest) UnsetFile() {
	o.File.Unset()
}

func (o AansluitingBestandRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Naam != nil {
		toSerialize["naam"] = o.Naam
	}
	if o.Omschrijving.IsSet() {
		toSerialize["omschrijving"] = o.Omschrijving.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.File.IsSet() {
		toSerialize["file"] = o.File.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAansluitingBestandRequest struct {
	value *AansluitingBestandRequest
	isSet bool
}

func (v NullableAansluitingBestandRequest) Get() *AansluitingBestandRequest {
	return v.value
}

func (v *NullableAansluitingBestandRequest) Set(val *AansluitingBestandRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAansluitingBestandRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAansluitingBestandRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAansluitingBestandRequest(val *AansluitingBestandRequest) *NullableAansluitingBestandRequest {
	return &NullableAansluitingBestandRequest{value: val, isSet: true}
}

func (v NullableAansluitingBestandRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAansluitingBestandRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
