# PatchedAfspraakRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTijd** | Pointer to **time.Time** |  | [optional] 
**EindTijd** | Pointer to **time.Time** |  | [optional] 
**Titel** | Pointer to **string** |  | [optional] 
**Aansluiting** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPatchedAfspraakRequest

`func NewPatchedAfspraakRequest() *PatchedAfspraakRequest`

NewPatchedAfspraakRequest instantiates a new PatchedAfspraakRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAfspraakRequestWithDefaults

`func NewPatchedAfspraakRequestWithDefaults() *PatchedAfspraakRequest`

NewPatchedAfspraakRequestWithDefaults instantiates a new PatchedAfspraakRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTijd

`func (o *PatchedAfspraakRequest) GetStartTijd() time.Time`

GetStartTijd returns the StartTijd field if non-nil, zero value otherwise.

### GetStartTijdOk

`func (o *PatchedAfspraakRequest) GetStartTijdOk() (*time.Time, bool)`

GetStartTijdOk returns a tuple with the StartTijd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTijd

`func (o *PatchedAfspraakRequest) SetStartTijd(v time.Time)`

SetStartTijd sets StartTijd field to given value.

### HasStartTijd

`func (o *PatchedAfspraakRequest) HasStartTijd() bool`

HasStartTijd returns a boolean if a field has been set.

### GetEindTijd

`func (o *PatchedAfspraakRequest) GetEindTijd() time.Time`

GetEindTijd returns the EindTijd field if non-nil, zero value otherwise.

### GetEindTijdOk

`func (o *PatchedAfspraakRequest) GetEindTijdOk() (*time.Time, bool)`

GetEindTijdOk returns a tuple with the EindTijd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEindTijd

`func (o *PatchedAfspraakRequest) SetEindTijd(v time.Time)`

SetEindTijd sets EindTijd field to given value.

### HasEindTijd

`func (o *PatchedAfspraakRequest) HasEindTijd() bool`

HasEindTijd returns a boolean if a field has been set.

### GetTitel

`func (o *PatchedAfspraakRequest) GetTitel() string`

GetTitel returns the Titel field if non-nil, zero value otherwise.

### GetTitelOk

`func (o *PatchedAfspraakRequest) GetTitelOk() (*string, bool)`

GetTitelOk returns a tuple with the Titel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitel

`func (o *PatchedAfspraakRequest) SetTitel(v string)`

SetTitel sets Titel field to given value.

### HasTitel

`func (o *PatchedAfspraakRequest) HasTitel() bool`

HasTitel returns a boolean if a field has been set.

### GetAansluiting

`func (o *PatchedAfspraakRequest) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *PatchedAfspraakRequest) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *PatchedAfspraakRequest) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.

### HasAansluiting

`func (o *PatchedAfspraakRequest) HasAansluiting() bool`

HasAansluiting returns a boolean if a field has been set.

### SetAansluitingNil

`func (o *PatchedAfspraakRequest) SetAansluitingNil(b bool)`

 SetAansluitingNil sets the value for Aansluiting to be an explicit nil

### UnsetAansluiting
`func (o *PatchedAfspraakRequest) UnsetAansluiting()`

UnsetAansluiting ensures that no value is present for Aansluiting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


