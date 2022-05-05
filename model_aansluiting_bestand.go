/*
AMSVR Broker API

# Introductie     Voor u heeft u de online documentatie van de AMSVR Broker API. Omdat het een Nederlands domein      betreft worden er Engelse en Nederlandse terminologie door elkaar heen gebruikt.  Liefhebbers van de Swagger UI kunnen [hier terecht](/swagger-ui). Daarnaast is de API eveneens beschikbaar in de op endpoint niveau zoals bijvoorbeeld [hier](/api/aansluitingen).  ##        

API version: release-0.6.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package broker

import (
	"encoding/json"
	"time"
)

// AansluitingBestand struct for AansluitingBestand
type AansluitingBestand struct {
	Id int32 `json:"id"`
	Naam *string `json:"naam,omitempty"`
	Omschrijving NullableString `json:"omschrijving,omitempty"`
	// overige = Overige, pve = PVE, upd = UPD, nva = NVA (nota van aanvulling), nvw = NVW (nota van wijzigingen), gms_mutatierapport = GMS mutatierapport
	Type NullableAansluitingBestandTypeEnum `json:"type,omitempty"`
	File NullableString `json:"file,omitempty"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

// NewAansluitingBestand instantiates a new AansluitingBestand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAansluitingBestand(id int32, created time.Time, modified time.Time) *AansluitingBestand {
	this := AansluitingBestand{}
	this.Id = id
	this.Created = created
	this.Modified = modified
	return &this
}

// NewAansluitingBestandWithDefaults instantiates a new AansluitingBestand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAansluitingBestandWithDefaults() *AansluitingBestand {
	this := AansluitingBestand{}
	return &this
}

// GetId returns the Id field value
func (o *AansluitingBestand) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AansluitingBestand) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AansluitingBestand) SetId(v int32) {
	o.Id = v
}

// GetNaam returns the Naam field value if set, zero value otherwise.
func (o *AansluitingBestand) GetNaam() string {
	if o == nil || o.Naam == nil {
		var ret string
		return ret
	}
	return *o.Naam
}

// GetNaamOk returns a tuple with the Naam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AansluitingBestand) GetNaamOk() (*string, bool) {
	if o == nil || o.Naam == nil {
		return nil, false
	}
	return o.Naam, true
}

// HasNaam returns a boolean if a field has been set.
func (o *AansluitingBestand) HasNaam() bool {
	if o != nil && o.Naam != nil {
		return true
	}

	return false
}

// SetNaam gets a reference to the given string and assigns it to the Naam field.
func (o *AansluitingBestand) SetNaam(v string) {
	o.Naam = &v
}

// GetOmschrijving returns the Omschrijving field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AansluitingBestand) GetOmschrijving() string {
	if o == nil || o.Omschrijving.Get() == nil {
		var ret string
		return ret
	}
	return *o.Omschrijving.Get()
}

// GetOmschrijvingOk returns a tuple with the Omschrijving field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AansluitingBestand) GetOmschrijvingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Omschrijving.Get(), o.Omschrijving.IsSet()
}

// HasOmschrijving returns a boolean if a field has been set.
func (o *AansluitingBestand) HasOmschrijving() bool {
	if o != nil && o.Omschrijving.IsSet() {
		return true
	}

	return false
}

// SetOmschrijving gets a reference to the given NullableString and assigns it to the Omschrijving field.
func (o *AansluitingBestand) SetOmschrijving(v string) {
	o.Omschrijving.Set(&v)
}
// SetOmschrijvingNil sets the value for Omschrijving to be an explicit nil
func (o *AansluitingBestand) SetOmschrijvingNil() {
	o.Omschrijving.Set(nil)
}

// UnsetOmschrijving ensures that no value is present for Omschrijving, not even an explicit nil
func (o *AansluitingBestand) UnsetOmschrijving() {
	o.Omschrijving.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AansluitingBestand) GetType() AansluitingBestandTypeEnum {
	if o == nil || o.Type.Get() == nil {
		var ret AansluitingBestandTypeEnum
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AansluitingBestand) GetTypeOk() (*AansluitingBestandTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *AansluitingBestand) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableAansluitingBestandTypeEnum and assigns it to the Type field.
func (o *AansluitingBestand) SetType(v AansluitingBestandTypeEnum) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *AansluitingBestand) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *AansluitingBestand) UnsetType() {
	o.Type.Unset()
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AansluitingBestand) GetFile() string {
	if o == nil || o.File.Get() == nil {
		var ret string
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AansluitingBestand) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *AansluitingBestand) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableString and assigns it to the File field.
func (o *AansluitingBestand) SetFile(v string) {
	o.File.Set(&v)
}
// SetFileNil sets the value for File to be an explicit nil
func (o *AansluitingBestand) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *AansluitingBestand) UnsetFile() {
	o.File.Unset()
}

// GetCreated returns the Created field value
func (o *AansluitingBestand) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AansluitingBestand) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AansluitingBestand) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *AansluitingBestand) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *AansluitingBestand) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *AansluitingBestand) SetModified(v time.Time) {
	o.Modified = v
}

func (o AansluitingBestand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
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
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["modified"] = o.Modified
	}
	return json.Marshal(toSerialize)
}

type NullableAansluitingBestand struct {
	value *AansluitingBestand
	isSet bool
}

func (v NullableAansluitingBestand) Get() *AansluitingBestand {
	return v.value
}

func (v *NullableAansluitingBestand) Set(val *AansluitingBestand) {
	v.value = val
	v.isSet = true
}

func (v NullableAansluitingBestand) IsSet() bool {
	return v.isSet
}

func (v *NullableAansluitingBestand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAansluitingBestand(val *AansluitingBestand) *NullableAansluitingBestand {
	return &NullableAansluitingBestand{value: val, isSet: true}
}

func (v NullableAansluitingBestand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAansluitingBestand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


