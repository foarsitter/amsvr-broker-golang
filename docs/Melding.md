# Melding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Criterium** | **int32** | The naam field of the Criterium | 
**Tijd** | **time.Time** |  | 
**Aansluitnummer** | **string** |  | [readonly] 
**Sector** | Pointer to **NullableInt32** | Reference to the number of the sector | [optional] 
**Status** | Pointer to [**NullableMeldingStatus**](MeldingStatus.md) | 0 &#x3D; Rust, 1 &#x3D; Alarm | [optional] 
**Toestand** | [**NullableMeldingToestand**](MeldingToestand.md) | 0 &#x3D; Actief, 1 &#x3D; Test, 2 &#x3D; Passief | 
**Data** | Pointer to **string** |  | [optional] 
**Actueel** | Pointer to **bool** |  | [optional] 
**GmsId** | Pointer to **NullableString** |  | [optional] 
**Oorzaak** | Pointer to **NullableString** |  | [optional] 
**MeldingId** | Pointer to **NullableString** |  | [optional] 
**Prio** | Pointer to [**NullablePrioEnum**](PrioEnum.md) |  | [optional] 

## Methods

### NewMelding

`func NewMelding(id int32, criterium int32, tijd time.Time, aansluitnummer string, toestand NullableMeldingToestand, ) *Melding`

NewMelding instantiates a new Melding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeldingWithDefaults

`func NewMeldingWithDefaults() *Melding`

NewMeldingWithDefaults instantiates a new Melding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Melding) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Melding) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Melding) SetId(v int32)`

SetId sets Id field to given value.


### GetCriterium

`func (o *Melding) GetCriterium() int32`

GetCriterium returns the Criterium field if non-nil, zero value otherwise.

### GetCriteriumOk

`func (o *Melding) GetCriteriumOk() (*int32, bool)`

GetCriteriumOk returns a tuple with the Criterium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriterium

`func (o *Melding) SetCriterium(v int32)`

SetCriterium sets Criterium field to given value.


### GetTijd

`func (o *Melding) GetTijd() time.Time`

GetTijd returns the Tijd field if non-nil, zero value otherwise.

### GetTijdOk

`func (o *Melding) GetTijdOk() (*time.Time, bool)`

GetTijdOk returns a tuple with the Tijd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTijd

`func (o *Melding) SetTijd(v time.Time)`

SetTijd sets Tijd field to given value.


### GetAansluitnummer

`func (o *Melding) GetAansluitnummer() string`

GetAansluitnummer returns the Aansluitnummer field if non-nil, zero value otherwise.

### GetAansluitnummerOk

`func (o *Melding) GetAansluitnummerOk() (*string, bool)`

GetAansluitnummerOk returns a tuple with the Aansluitnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluitnummer

`func (o *Melding) SetAansluitnummer(v string)`

SetAansluitnummer sets Aansluitnummer field to given value.


### GetSector

`func (o *Melding) GetSector() int32`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *Melding) GetSectorOk() (*int32, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *Melding) SetSector(v int32)`

SetSector sets Sector field to given value.

### HasSector

`func (o *Melding) HasSector() bool`

HasSector returns a boolean if a field has been set.

### SetSectorNil

`func (o *Melding) SetSectorNil(b bool)`

 SetSectorNil sets the value for Sector to be an explicit nil

### UnsetSector
`func (o *Melding) UnsetSector()`

UnsetSector ensures that no value is present for Sector, not even an explicit nil
### GetStatus

`func (o *Melding) GetStatus() MeldingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Melding) GetStatusOk() (*MeldingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Melding) SetStatus(v MeldingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Melding) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Melding) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Melding) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetToestand

`func (o *Melding) GetToestand() MeldingToestand`

GetToestand returns the Toestand field if non-nil, zero value otherwise.

### GetToestandOk

`func (o *Melding) GetToestandOk() (*MeldingToestand, bool)`

GetToestandOk returns a tuple with the Toestand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToestand

`func (o *Melding) SetToestand(v MeldingToestand)`

SetToestand sets Toestand field to given value.


### SetToestandNil

`func (o *Melding) SetToestandNil(b bool)`

 SetToestandNil sets the value for Toestand to be an explicit nil

### UnsetToestand
`func (o *Melding) UnsetToestand()`

UnsetToestand ensures that no value is present for Toestand, not even an explicit nil
### GetData

`func (o *Melding) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Melding) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Melding) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Melding) HasData() bool`

HasData returns a boolean if a field has been set.

### GetActueel

`func (o *Melding) GetActueel() bool`

GetActueel returns the Actueel field if non-nil, zero value otherwise.

### GetActueelOk

`func (o *Melding) GetActueelOk() (*bool, bool)`

GetActueelOk returns a tuple with the Actueel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActueel

`func (o *Melding) SetActueel(v bool)`

SetActueel sets Actueel field to given value.

### HasActueel

`func (o *Melding) HasActueel() bool`

HasActueel returns a boolean if a field has been set.

### GetGmsId

`func (o *Melding) GetGmsId() string`

GetGmsId returns the GmsId field if non-nil, zero value otherwise.

### GetGmsIdOk

`func (o *Melding) GetGmsIdOk() (*string, bool)`

GetGmsIdOk returns a tuple with the GmsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmsId

`func (o *Melding) SetGmsId(v string)`

SetGmsId sets GmsId field to given value.

### HasGmsId

`func (o *Melding) HasGmsId() bool`

HasGmsId returns a boolean if a field has been set.

### SetGmsIdNil

`func (o *Melding) SetGmsIdNil(b bool)`

 SetGmsIdNil sets the value for GmsId to be an explicit nil

### UnsetGmsId
`func (o *Melding) UnsetGmsId()`

UnsetGmsId ensures that no value is present for GmsId, not even an explicit nil
### GetOorzaak

`func (o *Melding) GetOorzaak() string`

GetOorzaak returns the Oorzaak field if non-nil, zero value otherwise.

### GetOorzaakOk

`func (o *Melding) GetOorzaakOk() (*string, bool)`

GetOorzaakOk returns a tuple with the Oorzaak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOorzaak

`func (o *Melding) SetOorzaak(v string)`

SetOorzaak sets Oorzaak field to given value.

### HasOorzaak

`func (o *Melding) HasOorzaak() bool`

HasOorzaak returns a boolean if a field has been set.

### SetOorzaakNil

`func (o *Melding) SetOorzaakNil(b bool)`

 SetOorzaakNil sets the value for Oorzaak to be an explicit nil

### UnsetOorzaak
`func (o *Melding) UnsetOorzaak()`

UnsetOorzaak ensures that no value is present for Oorzaak, not even an explicit nil
### GetMeldingId

`func (o *Melding) GetMeldingId() string`

GetMeldingId returns the MeldingId field if non-nil, zero value otherwise.

### GetMeldingIdOk

`func (o *Melding) GetMeldingIdOk() (*string, bool)`

GetMeldingIdOk returns a tuple with the MeldingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeldingId

`func (o *Melding) SetMeldingId(v string)`

SetMeldingId sets MeldingId field to given value.

### HasMeldingId

`func (o *Melding) HasMeldingId() bool`

HasMeldingId returns a boolean if a field has been set.

### SetMeldingIdNil

`func (o *Melding) SetMeldingIdNil(b bool)`

 SetMeldingIdNil sets the value for MeldingId to be an explicit nil

### UnsetMeldingId
`func (o *Melding) UnsetMeldingId()`

UnsetMeldingId ensures that no value is present for MeldingId, not even an explicit nil
### GetPrio

`func (o *Melding) GetPrio() PrioEnum`

GetPrio returns the Prio field if non-nil, zero value otherwise.

### GetPrioOk

`func (o *Melding) GetPrioOk() (*PrioEnum, bool)`

GetPrioOk returns a tuple with the Prio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrio

`func (o *Melding) SetPrio(v PrioEnum)`

SetPrio sets Prio field to given value.

### HasPrio

`func (o *Melding) HasPrio() bool`

HasPrio returns a boolean if a field has been set.

### SetPrioNil

`func (o *Melding) SetPrioNil(b bool)`

 SetPrioNil sets the value for Prio to be an explicit nil

### UnsetPrio
`func (o *Melding) UnsetPrio()`

UnsetPrio ensures that no value is present for Prio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


