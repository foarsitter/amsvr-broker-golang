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

// ValidationRule struct for ValidationRule
type ValidationRule struct {
	Type NullableValidationRuleTypeEnum `json:"type"`
	Melding string `json:"melding"`
}

// NewValidationRule instantiates a new ValidationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationRule(type_ NullableValidationRuleTypeEnum, melding string) *ValidationRule {
	this := ValidationRule{}
	this.Type = type_
	this.Melding = melding
	return &this
}

// NewValidationRuleWithDefaults instantiates a new ValidationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationRuleWithDefaults() *ValidationRule {
	this := ValidationRule{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for ValidationRuleTypeEnum will be returned
func (o *ValidationRule) GetType() ValidationRuleTypeEnum {
	if o == nil || o.Type.Get() == nil {
		var ret ValidationRuleTypeEnum
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValidationRule) GetTypeOk() (*ValidationRuleTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *ValidationRule) SetType(v ValidationRuleTypeEnum) {
	o.Type.Set(&v)
}

// GetMelding returns the Melding field value
func (o *ValidationRule) GetMelding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Melding
}

// GetMeldingOk returns a tuple with the Melding field value
// and a boolean to check if the value has been set.
func (o *ValidationRule) GetMeldingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Melding, true
}

// SetMelding sets field value
func (o *ValidationRule) SetMelding(v string) {
	o.Melding = v
}

func (o ValidationRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if true {
		toSerialize["melding"] = o.Melding
	}
	return json.Marshal(toSerialize)
}

type NullableValidationRule struct {
	value *ValidationRule
	isSet bool
}

func (v NullableValidationRule) Get() *ValidationRule {
	return v.value
}

func (v *NullableValidationRule) Set(val *ValidationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationRule(val *ValidationRule) *NullableValidationRule {
	return &NullableValidationRule{value: val, isSet: true}
}

func (v NullableValidationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


