/*
AMSVR Broker API

# Introductie     Voor u heeft u de online documentatie van de AMSVR Broker API. Omdat het een Nederlands domein      betreft worden er Engelse en Nederlandse terminologie door elkaar heen gebruikt.  Liefhebbers van de Swagger UI kunnen [hier terecht](/swagger-ui). Daarnaast is de API eveneens beschikbaar in de op endpoint niveau zoals bijvoorbeeld [hier](/api/aansluitingen).  ##        

API version: release-0.6.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package broker

import (
	"encoding/json"
	"fmt"
)

// MeldingToestand the model 'MeldingToestand'
type MeldingToestand int32

// List of MeldingToestand
const (
	MELDINGTOESTAND__0 MeldingToestand = 0
	MELDINGTOESTAND__1 MeldingToestand = 1
	MELDINGTOESTAND__2 MeldingToestand = 2
)

// All allowed values of MeldingToestand enum
var AllowedMeldingToestandEnumValues = []MeldingToestand{
	0,
	1,
	2,
}

func (v *MeldingToestand) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MeldingToestand(value)
	for _, existing := range AllowedMeldingToestandEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MeldingToestand", value)
}

// NewMeldingToestandFromValue returns a pointer to a valid MeldingToestand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMeldingToestandFromValue(v int32) (*MeldingToestand, error) {
	ev := MeldingToestand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MeldingToestand: valid values are %v", v, AllowedMeldingToestandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MeldingToestand) IsValid() bool {
	for _, existing := range AllowedMeldingToestandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MeldingToestand value
func (v MeldingToestand) Ptr() *MeldingToestand {
	return &v
}

type NullableMeldingToestand struct {
	value *MeldingToestand
	isSet bool
}

func (v NullableMeldingToestand) Get() *MeldingToestand {
	return v.value
}

func (v *NullableMeldingToestand) Set(val *MeldingToestand) {
	v.value = val
	v.isSet = true
}

func (v NullableMeldingToestand) IsSet() bool {
	return v.isSet
}

func (v *NullableMeldingToestand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeldingToestand(val *MeldingToestand) *NullableMeldingToestand {
	return &NullableMeldingToestand{value: val, isSet: true}
}

func (v NullableMeldingToestand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeldingToestand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

