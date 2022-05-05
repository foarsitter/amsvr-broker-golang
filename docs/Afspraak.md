# Afspraak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**StartTijd** | **time.Time** |  | 
**EindTijd** | Pointer to **time.Time** |  | [optional] 
**Titel** | **string** |  | 
**Title** | **string** |  | [readonly] 
**Start** | **time.Time** |  | [readonly] 
**End** | **time.Time** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Aansluiting** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAfspraak

`func NewAfspraak(id int32, startTijd time.Time, titel string, title string, start time.Time, end time.Time, url string, ) *Afspraak`

NewAfspraak instantiates a new Afspraak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfspraakWithDefaults

`func NewAfspraakWithDefaults() *Afspraak`

NewAfspraakWithDefaults instantiates a new Afspraak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Afspraak) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Afspraak) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Afspraak) SetId(v int32)`

SetId sets Id field to given value.


### GetStartTijd

`func (o *Afspraak) GetStartTijd() time.Time`

GetStartTijd returns the StartTijd field if non-nil, zero value otherwise.

### GetStartTijdOk

`func (o *Afspraak) GetStartTijdOk() (*time.Time, bool)`

GetStartTijdOk returns a tuple with the StartTijd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTijd

`func (o *Afspraak) SetStartTijd(v time.Time)`

SetStartTijd sets StartTijd field to given value.


### GetEindTijd

`func (o *Afspraak) GetEindTijd() time.Time`

GetEindTijd returns the EindTijd field if non-nil, zero value otherwise.

### GetEindTijdOk

`func (o *Afspraak) GetEindTijdOk() (*time.Time, bool)`

GetEindTijdOk returns a tuple with the EindTijd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEindTijd

`func (o *Afspraak) SetEindTijd(v time.Time)`

SetEindTijd sets EindTijd field to given value.

### HasEindTijd

`func (o *Afspraak) HasEindTijd() bool`

HasEindTijd returns a boolean if a field has been set.

### GetTitel

`func (o *Afspraak) GetTitel() string`

GetTitel returns the Titel field if non-nil, zero value otherwise.

### GetTitelOk

`func (o *Afspraak) GetTitelOk() (*string, bool)`

GetTitelOk returns a tuple with the Titel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitel

`func (o *Afspraak) SetTitel(v string)`

SetTitel sets Titel field to given value.


### GetTitle

`func (o *Afspraak) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Afspraak) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Afspraak) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStart

`func (o *Afspraak) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Afspraak) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Afspraak) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *Afspraak) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Afspraak) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Afspraak) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetUrl

`func (o *Afspraak) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Afspraak) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Afspraak) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAansluiting

`func (o *Afspraak) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *Afspraak) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *Afspraak) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.

### HasAansluiting

`func (o *Afspraak) HasAansluiting() bool`

HasAansluiting returns a boolean if a field has been set.

### SetAansluitingNil

`func (o *Afspraak) SetAansluitingNil(b bool)`

 SetAansluitingNil sets the value for Aansluiting to be an explicit nil

### UnsetAansluiting
`func (o *Afspraak) UnsetAansluiting()`

UnsetAansluiting ensures that no value is present for Aansluiting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


