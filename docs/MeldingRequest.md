# MeldingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criterium** | **int32** | The naam field of the Criterium | 
**Tijd** | **time.Time** |  | 
**Sector** | Pointer to **NullableInt32** | Reference to the number of the sector | [optional] 
**Status** | Pointer to [**NullableMeldingStatus**](MeldingStatus.md) | 0 &#x3D; Rust, 1 &#x3D; Alarm | [optional] 
**Toestand** | [**NullableMeldingToestand**](MeldingToestand.md) | 0 &#x3D; Actief, 1 &#x3D; Test, 2 &#x3D; Passief | 
**Data** | Pointer to **string** |  | [optional] 
**Actueel** | Pointer to **bool** |  | [optional] 
**GmsId** | Pointer to **NullableString** |  | [optional] 
**Oorzaak** | Pointer to **NullableString** |  | [optional] 
**FepNr** | **int32** |  | 
**VolgNr** | **int32** |  | 
**MeldingId** | Pointer to **NullableString** |  | [optional] 
**Prio** | Pointer to [**NullablePrioEnum**](PrioEnum.md) |  | [optional] 

## Methods

### NewMeldingRequest

`func NewMeldingRequest(criterium int32, tijd time.Time, toestand NullableMeldingToestand, fepNr int32, volgNr int32, ) *MeldingRequest`

NewMeldingRequest instantiates a new MeldingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeldingRequestWithDefaults

`func NewMeldingRequestWithDefaults() *MeldingRequest`

NewMeldingRequestWithDefaults instantiates a new MeldingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriterium

`func (o *MeldingRequest) GetCriterium() int32`

GetCriterium returns the Criterium field if non-nil, zero value otherwise.

### GetCriteriumOk

`func (o *MeldingRequest) GetCriteriumOk() (*int32, bool)`

GetCriteriumOk returns a tuple with the Criterium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriterium

`func (o *MeldingRequest) SetCriterium(v int32)`

SetCriterium sets Criterium field to given value.


### GetTijd

`func (o *MeldingRequest) GetTijd() time.Time`

GetTijd returns the Tijd field if non-nil, zero value otherwise.

### GetTijdOk

`func (o *MeldingRequest) GetTijdOk() (*time.Time, bool)`

GetTijdOk returns a tuple with the Tijd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTijd

`func (o *MeldingRequest) SetTijd(v time.Time)`

SetTijd sets Tijd field to given value.


### GetSector

`func (o *MeldingRequest) GetSector() int32`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *MeldingRequest) GetSectorOk() (*int32, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *MeldingRequest) SetSector(v int32)`

SetSector sets Sector field to given value.

### HasSector

`func (o *MeldingRequest) HasSector() bool`

HasSector returns a boolean if a field has been set.

### SetSectorNil

`func (o *MeldingRequest) SetSectorNil(b bool)`

 SetSectorNil sets the value for Sector to be an explicit nil

### UnsetSector
`func (o *MeldingRequest) UnsetSector()`

UnsetSector ensures that no value is present for Sector, not even an explicit nil
### GetStatus

`func (o *MeldingRequest) GetStatus() MeldingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeldingRequest) GetStatusOk() (*MeldingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeldingRequest) SetStatus(v MeldingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MeldingRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MeldingRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MeldingRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetToestand

`func (o *MeldingRequest) GetToestand() MeldingToestand`

GetToestand returns the Toestand field if non-nil, zero value otherwise.

### GetToestandOk

`func (o *MeldingRequest) GetToestandOk() (*MeldingToestand, bool)`

GetToestandOk returns a tuple with the Toestand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToestand

`func (o *MeldingRequest) SetToestand(v MeldingToestand)`

SetToestand sets Toestand field to given value.


### SetToestandNil

`func (o *MeldingRequest) SetToestandNil(b bool)`

 SetToestandNil sets the value for Toestand to be an explicit nil

### UnsetToestand
`func (o *MeldingRequest) UnsetToestand()`

UnsetToestand ensures that no value is present for Toestand, not even an explicit nil
### GetData

`func (o *MeldingRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MeldingRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MeldingRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *MeldingRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetActueel

`func (o *MeldingRequest) GetActueel() bool`

GetActueel returns the Actueel field if non-nil, zero value otherwise.

### GetActueelOk

`func (o *MeldingRequest) GetActueelOk() (*bool, bool)`

GetActueelOk returns a tuple with the Actueel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActueel

`func (o *MeldingRequest) SetActueel(v bool)`

SetActueel sets Actueel field to given value.

### HasActueel

`func (o *MeldingRequest) HasActueel() bool`

HasActueel returns a boolean if a field has been set.

### GetGmsId

`func (o *MeldingRequest) GetGmsId() string`

GetGmsId returns the GmsId field if non-nil, zero value otherwise.

### GetGmsIdOk

`func (o *MeldingRequest) GetGmsIdOk() (*string, bool)`

GetGmsIdOk returns a tuple with the GmsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmsId

`func (o *MeldingRequest) SetGmsId(v string)`

SetGmsId sets GmsId field to given value.

### HasGmsId

`func (o *MeldingRequest) HasGmsId() bool`

HasGmsId returns a boolean if a field has been set.

### SetGmsIdNil

`func (o *MeldingRequest) SetGmsIdNil(b bool)`

 SetGmsIdNil sets the value for GmsId to be an explicit nil

### UnsetGmsId
`func (o *MeldingRequest) UnsetGmsId()`

UnsetGmsId ensures that no value is present for GmsId, not even an explicit nil
### GetOorzaak

`func (o *MeldingRequest) GetOorzaak() string`

GetOorzaak returns the Oorzaak field if non-nil, zero value otherwise.

### GetOorzaakOk

`func (o *MeldingRequest) GetOorzaakOk() (*string, bool)`

GetOorzaakOk returns a tuple with the Oorzaak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOorzaak

`func (o *MeldingRequest) SetOorzaak(v string)`

SetOorzaak sets Oorzaak field to given value.

### HasOorzaak

`func (o *MeldingRequest) HasOorzaak() bool`

HasOorzaak returns a boolean if a field has been set.

### SetOorzaakNil

`func (o *MeldingRequest) SetOorzaakNil(b bool)`

 SetOorzaakNil sets the value for Oorzaak to be an explicit nil

### UnsetOorzaak
`func (o *MeldingRequest) UnsetOorzaak()`

UnsetOorzaak ensures that no value is present for Oorzaak, not even an explicit nil
### GetFepNr

`func (o *MeldingRequest) GetFepNr() int32`

GetFepNr returns the FepNr field if non-nil, zero value otherwise.

### GetFepNrOk

`func (o *MeldingRequest) GetFepNrOk() (*int32, bool)`

GetFepNrOk returns a tuple with the FepNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFepNr

`func (o *MeldingRequest) SetFepNr(v int32)`

SetFepNr sets FepNr field to given value.


### GetVolgNr

`func (o *MeldingRequest) GetVolgNr() int32`

GetVolgNr returns the VolgNr field if non-nil, zero value otherwise.

### GetVolgNrOk

`func (o *MeldingRequest) GetVolgNrOk() (*int32, bool)`

GetVolgNrOk returns a tuple with the VolgNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolgNr

`func (o *MeldingRequest) SetVolgNr(v int32)`

SetVolgNr sets VolgNr field to given value.


### GetMeldingId

`func (o *MeldingRequest) GetMeldingId() string`

GetMeldingId returns the MeldingId field if non-nil, zero value otherwise.

### GetMeldingIdOk

`func (o *MeldingRequest) GetMeldingIdOk() (*string, bool)`

GetMeldingIdOk returns a tuple with the MeldingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeldingId

`func (o *MeldingRequest) SetMeldingId(v string)`

SetMeldingId sets MeldingId field to given value.

### HasMeldingId

`func (o *MeldingRequest) HasMeldingId() bool`

HasMeldingId returns a boolean if a field has been set.

### SetMeldingIdNil

`func (o *MeldingRequest) SetMeldingIdNil(b bool)`

 SetMeldingIdNil sets the value for MeldingId to be an explicit nil

### UnsetMeldingId
`func (o *MeldingRequest) UnsetMeldingId()`

UnsetMeldingId ensures that no value is present for MeldingId, not even an explicit nil
### GetPrio

`func (o *MeldingRequest) GetPrio() PrioEnum`

GetPrio returns the Prio field if non-nil, zero value otherwise.

### GetPrioOk

`func (o *MeldingRequest) GetPrioOk() (*PrioEnum, bool)`

GetPrioOk returns a tuple with the Prio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrio

`func (o *MeldingRequest) SetPrio(v PrioEnum)`

SetPrio sets Prio field to given value.

### HasPrio

`func (o *MeldingRequest) HasPrio() bool`

HasPrio returns a boolean if a field has been set.

### SetPrioNil

`func (o *MeldingRequest) SetPrioNil(b bool)`

 SetPrioNil sets the value for Prio to be an explicit nil

### UnsetPrio
`func (o *MeldingRequest) UnsetPrio()`

UnsetPrio ensures that no value is present for Prio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


