# AanwezigCriterium

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Sector** | **int32** |  | [readonly] 
**Criterium** | **int32** |  | 
**LaatstGemeld** | **NullableTime** |  | [readonly] 
**LaatsteToestand** | [**NullableMeldingToestand**](MeldingToestand.md) |  | [readonly] 
**LaatsteStatus** | [**NullableMeldingStatus**](MeldingStatus.md) |  | [readonly] 

## Methods

### NewAanwezigCriterium

`func NewAanwezigCriterium(id int32, sector int32, criterium int32, laatstGemeld NullableTime, laatsteToestand NullableMeldingToestand, laatsteStatus NullableMeldingStatus, ) *AanwezigCriterium`

NewAanwezigCriterium instantiates a new AanwezigCriterium object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAanwezigCriteriumWithDefaults

`func NewAanwezigCriteriumWithDefaults() *AanwezigCriterium`

NewAanwezigCriteriumWithDefaults instantiates a new AanwezigCriterium object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AanwezigCriterium) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AanwezigCriterium) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AanwezigCriterium) SetId(v int32)`

SetId sets Id field to given value.


### GetSector

`func (o *AanwezigCriterium) GetSector() int32`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *AanwezigCriterium) GetSectorOk() (*int32, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *AanwezigCriterium) SetSector(v int32)`

SetSector sets Sector field to given value.


### GetCriterium

`func (o *AanwezigCriterium) GetCriterium() int32`

GetCriterium returns the Criterium field if non-nil, zero value otherwise.

### GetCriteriumOk

`func (o *AanwezigCriterium) GetCriteriumOk() (*int32, bool)`

GetCriteriumOk returns a tuple with the Criterium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriterium

`func (o *AanwezigCriterium) SetCriterium(v int32)`

SetCriterium sets Criterium field to given value.


### GetLaatstGemeld

`func (o *AanwezigCriterium) GetLaatstGemeld() time.Time`

GetLaatstGemeld returns the LaatstGemeld field if non-nil, zero value otherwise.

### GetLaatstGemeldOk

`func (o *AanwezigCriterium) GetLaatstGemeldOk() (*time.Time, bool)`

GetLaatstGemeldOk returns a tuple with the LaatstGemeld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaatstGemeld

`func (o *AanwezigCriterium) SetLaatstGemeld(v time.Time)`

SetLaatstGemeld sets LaatstGemeld field to given value.


### SetLaatstGemeldNil

`func (o *AanwezigCriterium) SetLaatstGemeldNil(b bool)`

 SetLaatstGemeldNil sets the value for LaatstGemeld to be an explicit nil

### UnsetLaatstGemeld
`func (o *AanwezigCriterium) UnsetLaatstGemeld()`

UnsetLaatstGemeld ensures that no value is present for LaatstGemeld, not even an explicit nil
### GetLaatsteToestand

`func (o *AanwezigCriterium) GetLaatsteToestand() MeldingToestand`

GetLaatsteToestand returns the LaatsteToestand field if non-nil, zero value otherwise.

### GetLaatsteToestandOk

`func (o *AanwezigCriterium) GetLaatsteToestandOk() (*MeldingToestand, bool)`

GetLaatsteToestandOk returns a tuple with the LaatsteToestand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaatsteToestand

`func (o *AanwezigCriterium) SetLaatsteToestand(v MeldingToestand)`

SetLaatsteToestand sets LaatsteToestand field to given value.


### SetLaatsteToestandNil

`func (o *AanwezigCriterium) SetLaatsteToestandNil(b bool)`

 SetLaatsteToestandNil sets the value for LaatsteToestand to be an explicit nil

### UnsetLaatsteToestand
`func (o *AanwezigCriterium) UnsetLaatsteToestand()`

UnsetLaatsteToestand ensures that no value is present for LaatsteToestand, not even an explicit nil
### GetLaatsteStatus

`func (o *AanwezigCriterium) GetLaatsteStatus() MeldingStatus`

GetLaatsteStatus returns the LaatsteStatus field if non-nil, zero value otherwise.

### GetLaatsteStatusOk

`func (o *AanwezigCriterium) GetLaatsteStatusOk() (*MeldingStatus, bool)`

GetLaatsteStatusOk returns a tuple with the LaatsteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaatsteStatus

`func (o *AanwezigCriterium) SetLaatsteStatus(v MeldingStatus)`

SetLaatsteStatus sets LaatsteStatus field to given value.


### SetLaatsteStatusNil

`func (o *AanwezigCriterium) SetLaatsteStatusNil(b bool)`

 SetLaatsteStatusNil sets the value for LaatsteStatus to be an explicit nil

### UnsetLaatsteStatus
`func (o *AanwezigCriterium) UnsetLaatsteStatus()`

UnsetLaatsteStatus ensures that no value is present for LaatsteStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


