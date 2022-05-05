# Waarschuwingsadres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Telefoonnummer** | **int32** |  | 
**Prioriteit** | Pointer to **int32** | De contactpersoon met de hoogste prioriteit wordt als eerste opgeroepen | [optional] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 

## Methods

### NewWaarschuwingsadres

`func NewWaarschuwingsadres(id int32, telefoonnummer int32, created time.Time, modified time.Time, ) *Waarschuwingsadres`

NewWaarschuwingsadres instantiates a new Waarschuwingsadres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaarschuwingsadresWithDefaults

`func NewWaarschuwingsadresWithDefaults() *Waarschuwingsadres`

NewWaarschuwingsadresWithDefaults instantiates a new Waarschuwingsadres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Waarschuwingsadres) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Waarschuwingsadres) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Waarschuwingsadres) SetId(v int32)`

SetId sets Id field to given value.


### GetTelefoonnummer

`func (o *Waarschuwingsadres) GetTelefoonnummer() int32`

GetTelefoonnummer returns the Telefoonnummer field if non-nil, zero value otherwise.

### GetTelefoonnummerOk

`func (o *Waarschuwingsadres) GetTelefoonnummerOk() (*int32, bool)`

GetTelefoonnummerOk returns a tuple with the Telefoonnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelefoonnummer

`func (o *Waarschuwingsadres) SetTelefoonnummer(v int32)`

SetTelefoonnummer sets Telefoonnummer field to given value.


### GetPrioriteit

`func (o *Waarschuwingsadres) GetPrioriteit() int32`

GetPrioriteit returns the Prioriteit field if non-nil, zero value otherwise.

### GetPrioriteitOk

`func (o *Waarschuwingsadres) GetPrioriteitOk() (*int32, bool)`

GetPrioriteitOk returns a tuple with the Prioriteit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioriteit

`func (o *Waarschuwingsadres) SetPrioriteit(v int32)`

SetPrioriteit sets Prioriteit field to given value.

### HasPrioriteit

`func (o *Waarschuwingsadres) HasPrioriteit() bool`

HasPrioriteit returns a boolean if a field has been set.

### GetCreated

`func (o *Waarschuwingsadres) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Waarschuwingsadres) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Waarschuwingsadres) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Waarschuwingsadres) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Waarschuwingsadres) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Waarschuwingsadres) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


