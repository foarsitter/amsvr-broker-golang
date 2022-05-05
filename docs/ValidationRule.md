# ValidationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NullableValidationRuleTypeEnum**](ValidationRuleTypeEnum.md) |  | [readonly] 
**Melding** | **string** |  | [readonly] 

## Methods

### NewValidationRule

`func NewValidationRule(type_ NullableValidationRuleTypeEnum, melding string, ) *ValidationRule`

NewValidationRule instantiates a new ValidationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationRuleWithDefaults

`func NewValidationRuleWithDefaults() *ValidationRule`

NewValidationRuleWithDefaults instantiates a new ValidationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ValidationRule) GetType() ValidationRuleTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidationRule) GetTypeOk() (*ValidationRuleTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidationRule) SetType(v ValidationRuleTypeEnum)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ValidationRule) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ValidationRule) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMelding

`func (o *ValidationRule) GetMelding() string`

GetMelding returns the Melding field if non-nil, zero value otherwise.

### GetMeldingOk

`func (o *ValidationRule) GetMeldingOk() (*string, bool)`

GetMeldingOk returns a tuple with the Melding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMelding

`func (o *ValidationRule) SetMelding(v string)`

SetMelding sets Melding field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


