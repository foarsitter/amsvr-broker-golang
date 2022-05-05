# AfsluitingZaak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | 
**Status** | [**NullableAfsluitingZaakStatusEnum**](AfsluitingZaakStatusEnum.md) | nieuw &#x3D; Afsluiting, ingediend &#x3D; Ingediend [ATSP], afgewezen &#x3D; Afgewezen [Risicobeheer], goedgekeurd &#x3D; Goedgekeurd [Risicobeheer], bevestigd &#x3D; Bevestigd [ATSP], uitgevoerd &#x3D; Uitgevoerd [AMS-Servicedesk], verwijderd &#x3D; Verwijderd [GMS], geannuleerd &#x3D; Geannuleerd | [readonly] 
**StatusChanged** | **time.Time** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**User** | **NullableInt32** |  | [readonly] 
**Gesloten** | **bool** |  | [readonly] 
**GeslotenDoor** | **NullableInt32** |  | [readonly] 
**Reden** | **string** |  | 
**RedenVanAfwijzing** | Pointer to **NullableString** |  | [optional] 
**GeplandeDatum** | **string** | Voor wanneer staat de afsluiting gepland? | 
**DefinitieveDatum** | **NullableString** | Wanneer wordt de aansluiting definitief verwijderd? | [readonly] 
**MogelijkeTransities** | **[]string** |  | [readonly] 

## Methods

### NewAfsluitingZaak

`func NewAfsluitingZaak(id int32, aansluiting int32, status NullableAfsluitingZaakStatusEnum, statusChanged time.Time, created time.Time, modified time.Time, user NullableInt32, gesloten bool, geslotenDoor NullableInt32, reden string, geplandeDatum string, definitieveDatum NullableString, mogelijkeTransities []string, ) *AfsluitingZaak`

NewAfsluitingZaak instantiates a new AfsluitingZaak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfsluitingZaakWithDefaults

`func NewAfsluitingZaakWithDefaults() *AfsluitingZaak`

NewAfsluitingZaakWithDefaults instantiates a new AfsluitingZaak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AfsluitingZaak) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AfsluitingZaak) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AfsluitingZaak) SetId(v int32)`

SetId sets Id field to given value.


### GetAansluiting

`func (o *AfsluitingZaak) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *AfsluitingZaak) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *AfsluitingZaak) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.


### GetStatus

`func (o *AfsluitingZaak) GetStatus() AfsluitingZaakStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AfsluitingZaak) GetStatusOk() (*AfsluitingZaakStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AfsluitingZaak) SetStatus(v AfsluitingZaakStatusEnum)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *AfsluitingZaak) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AfsluitingZaak) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusChanged

`func (o *AfsluitingZaak) GetStatusChanged() time.Time`

GetStatusChanged returns the StatusChanged field if non-nil, zero value otherwise.

### GetStatusChangedOk

`func (o *AfsluitingZaak) GetStatusChangedOk() (*time.Time, bool)`

GetStatusChangedOk returns a tuple with the StatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChanged

`func (o *AfsluitingZaak) SetStatusChanged(v time.Time)`

SetStatusChanged sets StatusChanged field to given value.


### GetCreated

`func (o *AfsluitingZaak) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AfsluitingZaak) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AfsluitingZaak) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *AfsluitingZaak) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AfsluitingZaak) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AfsluitingZaak) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetUser

`func (o *AfsluitingZaak) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AfsluitingZaak) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AfsluitingZaak) SetUser(v int32)`

SetUser sets User field to given value.


### SetUserNil

`func (o *AfsluitingZaak) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *AfsluitingZaak) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetGesloten

`func (o *AfsluitingZaak) GetGesloten() bool`

GetGesloten returns the Gesloten field if non-nil, zero value otherwise.

### GetGeslotenOk

`func (o *AfsluitingZaak) GetGeslotenOk() (*bool, bool)`

GetGeslotenOk returns a tuple with the Gesloten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGesloten

`func (o *AfsluitingZaak) SetGesloten(v bool)`

SetGesloten sets Gesloten field to given value.


### GetGeslotenDoor

`func (o *AfsluitingZaak) GetGeslotenDoor() int32`

GetGeslotenDoor returns the GeslotenDoor field if non-nil, zero value otherwise.

### GetGeslotenDoorOk

`func (o *AfsluitingZaak) GetGeslotenDoorOk() (*int32, bool)`

GetGeslotenDoorOk returns a tuple with the GeslotenDoor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeslotenDoor

`func (o *AfsluitingZaak) SetGeslotenDoor(v int32)`

SetGeslotenDoor sets GeslotenDoor field to given value.


### SetGeslotenDoorNil

`func (o *AfsluitingZaak) SetGeslotenDoorNil(b bool)`

 SetGeslotenDoorNil sets the value for GeslotenDoor to be an explicit nil

### UnsetGeslotenDoor
`func (o *AfsluitingZaak) UnsetGeslotenDoor()`

UnsetGeslotenDoor ensures that no value is present for GeslotenDoor, not even an explicit nil
### GetReden

`func (o *AfsluitingZaak) GetReden() string`

GetReden returns the Reden field if non-nil, zero value otherwise.

### GetRedenOk

`func (o *AfsluitingZaak) GetRedenOk() (*string, bool)`

GetRedenOk returns a tuple with the Reden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReden

`func (o *AfsluitingZaak) SetReden(v string)`

SetReden sets Reden field to given value.


### GetRedenVanAfwijzing

`func (o *AfsluitingZaak) GetRedenVanAfwijzing() string`

GetRedenVanAfwijzing returns the RedenVanAfwijzing field if non-nil, zero value otherwise.

### GetRedenVanAfwijzingOk

`func (o *AfsluitingZaak) GetRedenVanAfwijzingOk() (*string, bool)`

GetRedenVanAfwijzingOk returns a tuple with the RedenVanAfwijzing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedenVanAfwijzing

`func (o *AfsluitingZaak) SetRedenVanAfwijzing(v string)`

SetRedenVanAfwijzing sets RedenVanAfwijzing field to given value.

### HasRedenVanAfwijzing

`func (o *AfsluitingZaak) HasRedenVanAfwijzing() bool`

HasRedenVanAfwijzing returns a boolean if a field has been set.

### SetRedenVanAfwijzingNil

`func (o *AfsluitingZaak) SetRedenVanAfwijzingNil(b bool)`

 SetRedenVanAfwijzingNil sets the value for RedenVanAfwijzing to be an explicit nil

### UnsetRedenVanAfwijzing
`func (o *AfsluitingZaak) UnsetRedenVanAfwijzing()`

UnsetRedenVanAfwijzing ensures that no value is present for RedenVanAfwijzing, not even an explicit nil
### GetGeplandeDatum

`func (o *AfsluitingZaak) GetGeplandeDatum() string`

GetGeplandeDatum returns the GeplandeDatum field if non-nil, zero value otherwise.

### GetGeplandeDatumOk

`func (o *AfsluitingZaak) GetGeplandeDatumOk() (*string, bool)`

GetGeplandeDatumOk returns a tuple with the GeplandeDatum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeplandeDatum

`func (o *AfsluitingZaak) SetGeplandeDatum(v string)`

SetGeplandeDatum sets GeplandeDatum field to given value.


### GetDefinitieveDatum

`func (o *AfsluitingZaak) GetDefinitieveDatum() string`

GetDefinitieveDatum returns the DefinitieveDatum field if non-nil, zero value otherwise.

### GetDefinitieveDatumOk

`func (o *AfsluitingZaak) GetDefinitieveDatumOk() (*string, bool)`

GetDefinitieveDatumOk returns a tuple with the DefinitieveDatum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitieveDatum

`func (o *AfsluitingZaak) SetDefinitieveDatum(v string)`

SetDefinitieveDatum sets DefinitieveDatum field to given value.


### SetDefinitieveDatumNil

`func (o *AfsluitingZaak) SetDefinitieveDatumNil(b bool)`

 SetDefinitieveDatumNil sets the value for DefinitieveDatum to be an explicit nil

### UnsetDefinitieveDatum
`func (o *AfsluitingZaak) UnsetDefinitieveDatum()`

UnsetDefinitieveDatum ensures that no value is present for DefinitieveDatum, not even an explicit nil
### GetMogelijkeTransities

`func (o *AfsluitingZaak) GetMogelijkeTransities() []string`

GetMogelijkeTransities returns the MogelijkeTransities field if non-nil, zero value otherwise.

### GetMogelijkeTransitiesOk

`func (o *AfsluitingZaak) GetMogelijkeTransitiesOk() (*[]string, bool)`

GetMogelijkeTransitiesOk returns a tuple with the MogelijkeTransities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMogelijkeTransities

`func (o *AfsluitingZaak) SetMogelijkeTransities(v []string)`

SetMogelijkeTransities sets MogelijkeTransities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


