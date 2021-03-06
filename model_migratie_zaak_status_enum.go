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

// MigratieZaakStatusEnum the model 'MigratieZaakStatusEnum'
type MigratieZaakStatusEnum string

// List of MigratieZaakStatusEnum
const (
	MIGRATIEZAAKSTATUSENUM_NIEUW MigratieZaakStatusEnum = "nieuw"
	MIGRATIEZAAKSTATUSENUM_INGEDIEND MigratieZaakStatusEnum = "ingediend"
	MIGRATIEZAAKSTATUSENUM_GMS_DOORGEVOERD MigratieZaakStatusEnum = "gms_doorgevoerd"
	MIGRATIEZAAKSTATUSENUM_INGEPLAND MigratieZaakStatusEnum = "ingepland"
	MIGRATIEZAAKSTATUSENUM_GOEDGEKEURD MigratieZaakStatusEnum = "goedgekeurd"
	MIGRATIEZAAKSTATUSENUM_GEANNULEERD MigratieZaakStatusEnum = "geannuleerd"
	MIGRATIEZAAKSTATUSENUM_GMS_VERWIJDERD MigratieZaakStatusEnum = "gms_verwijderd"
)

// All allowed values of MigratieZaakStatusEnum enum
var AllowedMigratieZaakStatusEnumEnumValues = []MigratieZaakStatusEnum{
	"nieuw",
	"ingediend",
	"gms_doorgevoerd",
	"ingepland",
	"goedgekeurd",
	"geannuleerd",
	"gms_verwijderd",
}

func (v *MigratieZaakStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MigratieZaakStatusEnum(value)
	for _, existing := range AllowedMigratieZaakStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MigratieZaakStatusEnum", value)
}

// NewMigratieZaakStatusEnumFromValue returns a pointer to a valid MigratieZaakStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMigratieZaakStatusEnumFromValue(v string) (*MigratieZaakStatusEnum, error) {
	ev := MigratieZaakStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MigratieZaakStatusEnum: valid values are %v", v, AllowedMigratieZaakStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MigratieZaakStatusEnum) IsValid() bool {
	for _, existing := range AllowedMigratieZaakStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MigratieZaakStatusEnum value
func (v MigratieZaakStatusEnum) Ptr() *MigratieZaakStatusEnum {
	return &v
}

type NullableMigratieZaakStatusEnum struct {
	value *MigratieZaakStatusEnum
	isSet bool
}

func (v NullableMigratieZaakStatusEnum) Get() *MigratieZaakStatusEnum {
	return v.value
}

func (v *NullableMigratieZaakStatusEnum) Set(val *MigratieZaakStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMigratieZaakStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMigratieZaakStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigratieZaakStatusEnum(val *MigratieZaakStatusEnum) *NullableMigratieZaakStatusEnum {
	return &NullableMigratieZaakStatusEnum{value: val, isSet: true}
}

func (v NullableMigratieZaakStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigratieZaakStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

