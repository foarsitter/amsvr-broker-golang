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

// KiezerVervangenZaakIndienenStatusEnum the model 'KiezerVervangenZaakIndienenStatusEnum'
type KiezerVervangenZaakIndienenStatusEnum string

// List of KiezerVervangenZaakIndienenStatusEnum
const (
	KIEZERVERVANGENZAAKINDIENENSTATUSENUM_NIEUW KiezerVervangenZaakIndienenStatusEnum = "nieuw"
	KIEZERVERVANGENZAAKINDIENENSTATUSENUM_INGEDIEND KiezerVervangenZaakIndienenStatusEnum = "ingediend"
	KIEZERVERVANGENZAAKINDIENENSTATUSENUM_GOEDGEKEURD KiezerVervangenZaakIndienenStatusEnum = "goedgekeurd"
	KIEZERVERVANGENZAAKINDIENENSTATUSENUM_INGEPLAND KiezerVervangenZaakIndienenStatusEnum = "ingepland"
	KIEZERVERVANGENZAAKINDIENENSTATUSENUM_DOORGEVOERD KiezerVervangenZaakIndienenStatusEnum = "doorgevoerd"
	KIEZERVERVANGENZAAKINDIENENSTATUSENUM_GEANNULEERD KiezerVervangenZaakIndienenStatusEnum = "geannuleerd"
)

// All allowed values of KiezerVervangenZaakIndienenStatusEnum enum
var AllowedKiezerVervangenZaakIndienenStatusEnumEnumValues = []KiezerVervangenZaakIndienenStatusEnum{
	"nieuw",
	"ingediend",
	"goedgekeurd",
	"ingepland",
	"doorgevoerd",
	"geannuleerd",
}

func (v *KiezerVervangenZaakIndienenStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KiezerVervangenZaakIndienenStatusEnum(value)
	for _, existing := range AllowedKiezerVervangenZaakIndienenStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KiezerVervangenZaakIndienenStatusEnum", value)
}

// NewKiezerVervangenZaakIndienenStatusEnumFromValue returns a pointer to a valid KiezerVervangenZaakIndienenStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKiezerVervangenZaakIndienenStatusEnumFromValue(v string) (*KiezerVervangenZaakIndienenStatusEnum, error) {
	ev := KiezerVervangenZaakIndienenStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KiezerVervangenZaakIndienenStatusEnum: valid values are %v", v, AllowedKiezerVervangenZaakIndienenStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KiezerVervangenZaakIndienenStatusEnum) IsValid() bool {
	for _, existing := range AllowedKiezerVervangenZaakIndienenStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KiezerVervangenZaakIndienenStatusEnum value
func (v KiezerVervangenZaakIndienenStatusEnum) Ptr() *KiezerVervangenZaakIndienenStatusEnum {
	return &v
}

type NullableKiezerVervangenZaakIndienenStatusEnum struct {
	value *KiezerVervangenZaakIndienenStatusEnum
	isSet bool
}

func (v NullableKiezerVervangenZaakIndienenStatusEnum) Get() *KiezerVervangenZaakIndienenStatusEnum {
	return v.value
}

func (v *NullableKiezerVervangenZaakIndienenStatusEnum) Set(val *KiezerVervangenZaakIndienenStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableKiezerVervangenZaakIndienenStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableKiezerVervangenZaakIndienenStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKiezerVervangenZaakIndienenStatusEnum(val *KiezerVervangenZaakIndienenStatusEnum) *NullableKiezerVervangenZaakIndienenStatusEnum {
	return &NullableKiezerVervangenZaakIndienenStatusEnum{value: val, isSet: true}
}

func (v NullableKiezerVervangenZaakIndienenStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKiezerVervangenZaakIndienenStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

