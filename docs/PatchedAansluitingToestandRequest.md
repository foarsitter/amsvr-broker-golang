# PatchedAansluitingToestandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Toestand** | Pointer to [**NullableMeldingStatus**](MeldingStatus.md) | 0 &#x3D; Passief, 1 &#x3D; Actief | [optional] 

## Methods

### NewPatchedAansluitingToestandRequest

`func NewPatchedAansluitingToestandRequest() *PatchedAansluitingToestandRequest`

NewPatchedAansluitingToestandRequest instantiates a new PatchedAansluitingToestandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAansluitingToestandRequestWithDefaults

`func NewPatchedAansluitingToestandRequestWithDefaults() *PatchedAansluitingToestandRequest`

NewPatchedAansluitingToestandRequestWithDefaults instantiates a new PatchedAansluitingToestandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToestand

`func (o *PatchedAansluitingToestandRequest) GetToestand() MeldingStatus`

GetToestand returns the Toestand field if non-nil, zero value otherwise.

### GetToestandOk

`func (o *PatchedAansluitingToestandRequest) GetToestandOk() (*MeldingStatus, bool)`

GetToestandOk returns a tuple with the Toestand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToestand

`func (o *PatchedAansluitingToestandRequest) SetToestand(v MeldingStatus)`

SetToestand sets Toestand field to given value.

### HasToestand

`func (o *PatchedAansluitingToestandRequest) HasToestand() bool`

HasToestand returns a boolean if a field has been set.

### SetToestandNil

`func (o *PatchedAansluitingToestandRequest) SetToestandNil(b bool)`

 SetToestandNil sets the value for Toestand to be an explicit nil

### UnsetToestand
`func (o *PatchedAansluitingToestandRequest) UnsetToestand()`

UnsetToestand ensures that no value is present for Toestand, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


