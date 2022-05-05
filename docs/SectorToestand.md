# SectorToestand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Toestand** | Pointer to [**NullableMeldingToestand**](MeldingToestand.md) | 0 &#x3D; Test, 1 &#x3D; Actief, 2 &#x3D; Passief | [optional] 

## Methods

### NewSectorToestand

`func NewSectorToestand() *SectorToestand`

NewSectorToestand instantiates a new SectorToestand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectorToestandWithDefaults

`func NewSectorToestandWithDefaults() *SectorToestand`

NewSectorToestandWithDefaults instantiates a new SectorToestand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToestand

`func (o *SectorToestand) GetToestand() MeldingToestand`

GetToestand returns the Toestand field if non-nil, zero value otherwise.

### GetToestandOk

`func (o *SectorToestand) GetToestandOk() (*MeldingToestand, bool)`

GetToestandOk returns a tuple with the Toestand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToestand

`func (o *SectorToestand) SetToestand(v MeldingToestand)`

SetToestand sets Toestand field to given value.

### HasToestand

`func (o *SectorToestand) HasToestand() bool`

HasToestand returns a boolean if a field has been set.

### SetToestandNil

`func (o *SectorToestand) SetToestandNil(b bool)`

 SetToestandNil sets the value for Toestand to be an explicit nil

### UnsetToestand
`func (o *SectorToestand) UnsetToestand()`

UnsetToestand ensures that no value is present for Toestand, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


