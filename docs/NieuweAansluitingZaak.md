# NieuweAansluitingZaak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | 
**Status** | [**NullableNieuweAansluitingZaakStatusEnum**](NieuweAansluitingZaakStatusEnum.md) | nieuw &#x3D; Nieuw, ingediend &#x3D; Ingediend ter aanvraag [ATSP], afgekeurd &#x3D; Aanvraag afgekeurd [Risicobeheer], realisatie &#x3D; Aanvraag goedgekeurd [Risicobeheer], gms_doorgevoerd &#x3D; Doorgevoerd [GMS], technisch_gereed &#x3D; Live test ingepland [ATSP], goedgekeurd &#x3D; Goedgekeurd [AMS-Servicedesk], opgeleverd &#x3D; Actief gezet [Risicobeheer], geannuleerd &#x3D; Geannuleerd | [readonly] 
**StatusChanged** | **time.Time** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**User** | **NullableInt32** |  | [readonly] 
**Gesloten** | **bool** |  | [readonly] 
**GeslotenDoor** | **NullableInt32** |  | [readonly] 
**MogelijkeTransities** | **[]string** |  | [readonly] 

## Methods

### NewNieuweAansluitingZaak

`func NewNieuweAansluitingZaak(id int32, aansluiting int32, status NullableNieuweAansluitingZaakStatusEnum, statusChanged time.Time, created time.Time, modified time.Time, user NullableInt32, gesloten bool, geslotenDoor NullableInt32, mogelijkeTransities []string, ) *NieuweAansluitingZaak`

NewNieuweAansluitingZaak instantiates a new NieuweAansluitingZaak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNieuweAansluitingZaakWithDefaults

`func NewNieuweAansluitingZaakWithDefaults() *NieuweAansluitingZaak`

NewNieuweAansluitingZaakWithDefaults instantiates a new NieuweAansluitingZaak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NieuweAansluitingZaak) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NieuweAansluitingZaak) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NieuweAansluitingZaak) SetId(v int32)`

SetId sets Id field to given value.


### GetAansluiting

`func (o *NieuweAansluitingZaak) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *NieuweAansluitingZaak) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *NieuweAansluitingZaak) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.


### GetStatus

`func (o *NieuweAansluitingZaak) GetStatus() NieuweAansluitingZaakStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NieuweAansluitingZaak) GetStatusOk() (*NieuweAansluitingZaakStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NieuweAansluitingZaak) SetStatus(v NieuweAansluitingZaakStatusEnum)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *NieuweAansluitingZaak) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *NieuweAansluitingZaak) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusChanged

`func (o *NieuweAansluitingZaak) GetStatusChanged() time.Time`

GetStatusChanged returns the StatusChanged field if non-nil, zero value otherwise.

### GetStatusChangedOk

`func (o *NieuweAansluitingZaak) GetStatusChangedOk() (*time.Time, bool)`

GetStatusChangedOk returns a tuple with the StatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChanged

`func (o *NieuweAansluitingZaak) SetStatusChanged(v time.Time)`

SetStatusChanged sets StatusChanged field to given value.


### GetCreated

`func (o *NieuweAansluitingZaak) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NieuweAansluitingZaak) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NieuweAansluitingZaak) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *NieuweAansluitingZaak) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NieuweAansluitingZaak) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NieuweAansluitingZaak) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetUser

`func (o *NieuweAansluitingZaak) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NieuweAansluitingZaak) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NieuweAansluitingZaak) SetUser(v int32)`

SetUser sets User field to given value.


### SetUserNil

`func (o *NieuweAansluitingZaak) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *NieuweAansluitingZaak) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetGesloten

`func (o *NieuweAansluitingZaak) GetGesloten() bool`

GetGesloten returns the Gesloten field if non-nil, zero value otherwise.

### GetGeslotenOk

`func (o *NieuweAansluitingZaak) GetGeslotenOk() (*bool, bool)`

GetGeslotenOk returns a tuple with the Gesloten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGesloten

`func (o *NieuweAansluitingZaak) SetGesloten(v bool)`

SetGesloten sets Gesloten field to given value.


### GetGeslotenDoor

`func (o *NieuweAansluitingZaak) GetGeslotenDoor() int32`

GetGeslotenDoor returns the GeslotenDoor field if non-nil, zero value otherwise.

### GetGeslotenDoorOk

`func (o *NieuweAansluitingZaak) GetGeslotenDoorOk() (*int32, bool)`

GetGeslotenDoorOk returns a tuple with the GeslotenDoor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeslotenDoor

`func (o *NieuweAansluitingZaak) SetGeslotenDoor(v int32)`

SetGeslotenDoor sets GeslotenDoor field to given value.


### SetGeslotenDoorNil

`func (o *NieuweAansluitingZaak) SetGeslotenDoorNil(b bool)`

 SetGeslotenDoorNil sets the value for GeslotenDoor to be an explicit nil

### UnsetGeslotenDoor
`func (o *NieuweAansluitingZaak) UnsetGeslotenDoor()`

UnsetGeslotenDoor ensures that no value is present for GeslotenDoor, not even an explicit nil
### GetMogelijkeTransities

`func (o *NieuweAansluitingZaak) GetMogelijkeTransities() []string`

GetMogelijkeTransities returns the MogelijkeTransities field if non-nil, zero value otherwise.

### GetMogelijkeTransitiesOk

`func (o *NieuweAansluitingZaak) GetMogelijkeTransitiesOk() (*[]string, bool)`

GetMogelijkeTransitiesOk returns a tuple with the MogelijkeTransities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMogelijkeTransities

`func (o *NieuweAansluitingZaak) SetMogelijkeTransities(v []string)`

SetMogelijkeTransities sets MogelijkeTransities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


