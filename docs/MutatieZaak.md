# MutatieZaak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | 
**Status** | [**NullableMutatieZaakStatusEnum**](MutatieZaakStatusEnum.md) | mutatie &#x3D; Mutatie, ingediend &#x3D; Ingediend [ATSP], afgewezen &#x3D; Afgewezen [Risicobeheer], goedgekeurd &#x3D; Goedgekeurd [Risicobeheer], aangepast_broker &#x3D; Aangepast in de Broker [ATSP], gms_toegevoegd &#x3D; Toevoegingen doorgevoerd in het GMS [GMS], ingepland &#x3D; Live test ingepland [ATSP], aanpassingen_geverifieerd &#x3D; Aanpassingen goedgekeurd [AMS-Servicedesk], gms_verwijderd &#x3D; Criteria verwijderd in het GMS [AMS-Servicedesk], geannuleerd &#x3D; Geannuleerd | [readonly] 
**StatusChanged** | **time.Time** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**User** | **NullableInt32** |  | [readonly] 
**Gesloten** | **bool** |  | [readonly] 
**GeslotenDoor** | **NullableInt32** |  | [readonly] 
**Type** | Pointer to [**MutatieZaakTypeEnum**](MutatieZaakTypeEnum.md) |  | [optional] 
**Reden** | Pointer to **NullableString** |  | [optional] 
**OmschrijvingVanWijziging** | **string** |  | 
**BouwdelenToevoegen** | Pointer to **bool** |  | [optional] 
**MogelijkeTransities** | **[]string** |  | [readonly] 

## Methods

### NewMutatieZaak

`func NewMutatieZaak(id int32, aansluiting int32, status NullableMutatieZaakStatusEnum, statusChanged time.Time, created time.Time, modified time.Time, user NullableInt32, gesloten bool, geslotenDoor NullableInt32, omschrijvingVanWijziging string, mogelijkeTransities []string, ) *MutatieZaak`

NewMutatieZaak instantiates a new MutatieZaak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutatieZaakWithDefaults

`func NewMutatieZaakWithDefaults() *MutatieZaak`

NewMutatieZaakWithDefaults instantiates a new MutatieZaak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MutatieZaak) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MutatieZaak) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MutatieZaak) SetId(v int32)`

SetId sets Id field to given value.


### GetAansluiting

`func (o *MutatieZaak) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *MutatieZaak) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *MutatieZaak) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.


### GetStatus

`func (o *MutatieZaak) GetStatus() MutatieZaakStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MutatieZaak) GetStatusOk() (*MutatieZaakStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MutatieZaak) SetStatus(v MutatieZaakStatusEnum)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *MutatieZaak) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MutatieZaak) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusChanged

`func (o *MutatieZaak) GetStatusChanged() time.Time`

GetStatusChanged returns the StatusChanged field if non-nil, zero value otherwise.

### GetStatusChangedOk

`func (o *MutatieZaak) GetStatusChangedOk() (*time.Time, bool)`

GetStatusChangedOk returns a tuple with the StatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChanged

`func (o *MutatieZaak) SetStatusChanged(v time.Time)`

SetStatusChanged sets StatusChanged field to given value.


### GetCreated

`func (o *MutatieZaak) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MutatieZaak) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MutatieZaak) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *MutatieZaak) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *MutatieZaak) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *MutatieZaak) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetUser

`func (o *MutatieZaak) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MutatieZaak) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MutatieZaak) SetUser(v int32)`

SetUser sets User field to given value.


### SetUserNil

`func (o *MutatieZaak) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *MutatieZaak) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetGesloten

`func (o *MutatieZaak) GetGesloten() bool`

GetGesloten returns the Gesloten field if non-nil, zero value otherwise.

### GetGeslotenOk

`func (o *MutatieZaak) GetGeslotenOk() (*bool, bool)`

GetGeslotenOk returns a tuple with the Gesloten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGesloten

`func (o *MutatieZaak) SetGesloten(v bool)`

SetGesloten sets Gesloten field to given value.


### GetGeslotenDoor

`func (o *MutatieZaak) GetGeslotenDoor() int32`

GetGeslotenDoor returns the GeslotenDoor field if non-nil, zero value otherwise.

### GetGeslotenDoorOk

`func (o *MutatieZaak) GetGeslotenDoorOk() (*int32, bool)`

GetGeslotenDoorOk returns a tuple with the GeslotenDoor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeslotenDoor

`func (o *MutatieZaak) SetGeslotenDoor(v int32)`

SetGeslotenDoor sets GeslotenDoor field to given value.


### SetGeslotenDoorNil

`func (o *MutatieZaak) SetGeslotenDoorNil(b bool)`

 SetGeslotenDoorNil sets the value for GeslotenDoor to be an explicit nil

### UnsetGeslotenDoor
`func (o *MutatieZaak) UnsetGeslotenDoor()`

UnsetGeslotenDoor ensures that no value is present for GeslotenDoor, not even an explicit nil
### GetType

`func (o *MutatieZaak) GetType() MutatieZaakTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MutatieZaak) GetTypeOk() (*MutatieZaakTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MutatieZaak) SetType(v MutatieZaakTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *MutatieZaak) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReden

`func (o *MutatieZaak) GetReden() string`

GetReden returns the Reden field if non-nil, zero value otherwise.

### GetRedenOk

`func (o *MutatieZaak) GetRedenOk() (*string, bool)`

GetRedenOk returns a tuple with the Reden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReden

`func (o *MutatieZaak) SetReden(v string)`

SetReden sets Reden field to given value.

### HasReden

`func (o *MutatieZaak) HasReden() bool`

HasReden returns a boolean if a field has been set.

### SetRedenNil

`func (o *MutatieZaak) SetRedenNil(b bool)`

 SetRedenNil sets the value for Reden to be an explicit nil

### UnsetReden
`func (o *MutatieZaak) UnsetReden()`

UnsetReden ensures that no value is present for Reden, not even an explicit nil
### GetOmschrijvingVanWijziging

`func (o *MutatieZaak) GetOmschrijvingVanWijziging() string`

GetOmschrijvingVanWijziging returns the OmschrijvingVanWijziging field if non-nil, zero value otherwise.

### GetOmschrijvingVanWijzigingOk

`func (o *MutatieZaak) GetOmschrijvingVanWijzigingOk() (*string, bool)`

GetOmschrijvingVanWijzigingOk returns a tuple with the OmschrijvingVanWijziging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmschrijvingVanWijziging

`func (o *MutatieZaak) SetOmschrijvingVanWijziging(v string)`

SetOmschrijvingVanWijziging sets OmschrijvingVanWijziging field to given value.


### GetBouwdelenToevoegen

`func (o *MutatieZaak) GetBouwdelenToevoegen() bool`

GetBouwdelenToevoegen returns the BouwdelenToevoegen field if non-nil, zero value otherwise.

### GetBouwdelenToevoegenOk

`func (o *MutatieZaak) GetBouwdelenToevoegenOk() (*bool, bool)`

GetBouwdelenToevoegenOk returns a tuple with the BouwdelenToevoegen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBouwdelenToevoegen

`func (o *MutatieZaak) SetBouwdelenToevoegen(v bool)`

SetBouwdelenToevoegen sets BouwdelenToevoegen field to given value.

### HasBouwdelenToevoegen

`func (o *MutatieZaak) HasBouwdelenToevoegen() bool`

HasBouwdelenToevoegen returns a boolean if a field has been set.

### GetMogelijkeTransities

`func (o *MutatieZaak) GetMogelijkeTransities() []string`

GetMogelijkeTransities returns the MogelijkeTransities field if non-nil, zero value otherwise.

### GetMogelijkeTransitiesOk

`func (o *MutatieZaak) GetMogelijkeTransitiesOk() (*[]string, bool)`

GetMogelijkeTransitiesOk returns a tuple with the MogelijkeTransities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMogelijkeTransities

`func (o *MutatieZaak) SetMogelijkeTransities(v []string)`

SetMogelijkeTransities sets MogelijkeTransities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


