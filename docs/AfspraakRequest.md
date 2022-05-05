# AfspraakRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTijd** | **time.Time** |  | 
**EindTijd** | Pointer to **time.Time** |  | [optional] 
**Titel** | **string** |  | 
**Aansluiting** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAfspraakRequest

`func NewAfspraakRequest(startTijd time.Time, titel string, ) *AfspraakRequest`

NewAfspraakRequest instantiates a new AfspraakRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfspraakRequestWithDefaults

`func NewAfspraakRequestWithDefaults() *AfspraakRequest`

NewAfspraakRequestWithDefaults instantiates a new AfspraakRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTijd

`func (o *AfspraakRequest) GetStartTijd() time.Time`

GetStartTijd returns the StartTijd field if non-nil, zero value otherwise.

### GetStartTijdOk

`func (o *AfspraakRequest) GetStartTijdOk() (*time.Time, bool)`

GetStartTijdOk returns a tuple with the StartTijd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTijd

`func (o *AfspraakRequest) SetStartTijd(v time.Time)`

SetStartTijd sets StartTijd field to given value.


### GetEindTijd

`func (o *AfspraakRequest) GetEindTijd() time.Time`

GetEindTijd returns the EindTijd field if non-nil, zero value otherwise.

### GetEindTijdOk

`func (o *AfspraakRequest) GetEindTijdOk() (*time.Time, bool)`

GetEindTijdOk returns a tuple with the EindTijd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEindTijd

`func (o *AfspraakRequest) SetEindTijd(v time.Time)`

SetEindTijd sets EindTijd field to given value.

### HasEindTijd

`func (o *AfspraakRequest) HasEindTijd() bool`

HasEindTijd returns a boolean if a field has been set.

### GetTitel

`func (o *AfspraakRequest) GetTitel() string`

GetTitel returns the Titel field if non-nil, zero value otherwise.

### GetTitelOk

`func (o *AfspraakRequest) GetTitelOk() (*string, bool)`

GetTitelOk returns a tuple with the Titel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitel

`func (o *AfspraakRequest) SetTitel(v string)`

SetTitel sets Titel field to given value.


### GetAansluiting

`func (o *AfspraakRequest) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *AfspraakRequest) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *AfspraakRequest) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.

### HasAansluiting

`func (o *AfspraakRequest) HasAansluiting() bool`

HasAansluiting returns a boolean if a field has been set.

### SetAansluitingNil

`func (o *AfspraakRequest) SetAansluitingNil(b bool)`

 SetAansluitingNil sets the value for Aansluiting to be an explicit nil

### UnsetAansluiting
`func (o *AfspraakRequest) UnsetAansluiting()`

UnsetAansluiting ensures that no value is present for Aansluiting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


