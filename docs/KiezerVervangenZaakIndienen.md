# KiezerVervangenZaakIndienen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | 
**Status** | [**NullableKiezerVervangenZaakIndienenStatusEnum**](KiezerVervangenZaakIndienenStatusEnum.md) | nieuw &#x3D; Nieuw, ingediend &#x3D; Ingediend [ATSP], goedgekeurd &#x3D; Goedgekeurd [AMS-Servicedesk], ingepland &#x3D; Live test ingepland [ATSP], doorgevoerd &#x3D; Doorgevoerd [GMS], geannuleerd &#x3D; Geannuleerd | [readonly] 
**StatusChanged** | **time.Time** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**User** | **NullableInt32** |  | [readonly] 
**Gesloten** | **bool** |  | [readonly] 
**GeslotenDoor** | **NullableInt32** |  | [readonly] 
**MogelijkeTransities** | **[]string** |  | [readonly] 

## Methods

### NewKiezerVervangenZaakIndienen

`func NewKiezerVervangenZaakIndienen(id int32, aansluiting int32, status NullableKiezerVervangenZaakIndienenStatusEnum, statusChanged time.Time, created time.Time, modified time.Time, user NullableInt32, gesloten bool, geslotenDoor NullableInt32, mogelijkeTransities []string, ) *KiezerVervangenZaakIndienen`

NewKiezerVervangenZaakIndienen instantiates a new KiezerVervangenZaakIndienen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKiezerVervangenZaakIndienenWithDefaults

`func NewKiezerVervangenZaakIndienenWithDefaults() *KiezerVervangenZaakIndienen`

NewKiezerVervangenZaakIndienenWithDefaults instantiates a new KiezerVervangenZaakIndienen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KiezerVervangenZaakIndienen) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KiezerVervangenZaakIndienen) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KiezerVervangenZaakIndienen) SetId(v int32)`

SetId sets Id field to given value.


### GetAansluiting

`func (o *KiezerVervangenZaakIndienen) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *KiezerVervangenZaakIndienen) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *KiezerVervangenZaakIndienen) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.


### GetStatus

`func (o *KiezerVervangenZaakIndienen) GetStatus() KiezerVervangenZaakIndienenStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KiezerVervangenZaakIndienen) GetStatusOk() (*KiezerVervangenZaakIndienenStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KiezerVervangenZaakIndienen) SetStatus(v KiezerVervangenZaakIndienenStatusEnum)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *KiezerVervangenZaakIndienen) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *KiezerVervangenZaakIndienen) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusChanged

`func (o *KiezerVervangenZaakIndienen) GetStatusChanged() time.Time`

GetStatusChanged returns the StatusChanged field if non-nil, zero value otherwise.

### GetStatusChangedOk

`func (o *KiezerVervangenZaakIndienen) GetStatusChangedOk() (*time.Time, bool)`

GetStatusChangedOk returns a tuple with the StatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChanged

`func (o *KiezerVervangenZaakIndienen) SetStatusChanged(v time.Time)`

SetStatusChanged sets StatusChanged field to given value.


### GetCreated

`func (o *KiezerVervangenZaakIndienen) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *KiezerVervangenZaakIndienen) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *KiezerVervangenZaakIndienen) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *KiezerVervangenZaakIndienen) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *KiezerVervangenZaakIndienen) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *KiezerVervangenZaakIndienen) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetUser

`func (o *KiezerVervangenZaakIndienen) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *KiezerVervangenZaakIndienen) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *KiezerVervangenZaakIndienen) SetUser(v int32)`

SetUser sets User field to given value.


### SetUserNil

`func (o *KiezerVervangenZaakIndienen) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *KiezerVervangenZaakIndienen) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetGesloten

`func (o *KiezerVervangenZaakIndienen) GetGesloten() bool`

GetGesloten returns the Gesloten field if non-nil, zero value otherwise.

### GetGeslotenOk

`func (o *KiezerVervangenZaakIndienen) GetGeslotenOk() (*bool, bool)`

GetGeslotenOk returns a tuple with the Gesloten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGesloten

`func (o *KiezerVervangenZaakIndienen) SetGesloten(v bool)`

SetGesloten sets Gesloten field to given value.


### GetGeslotenDoor

`func (o *KiezerVervangenZaakIndienen) GetGeslotenDoor() int32`

GetGeslotenDoor returns the GeslotenDoor field if non-nil, zero value otherwise.

### GetGeslotenDoorOk

`func (o *KiezerVervangenZaakIndienen) GetGeslotenDoorOk() (*int32, bool)`

GetGeslotenDoorOk returns a tuple with the GeslotenDoor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeslotenDoor

`func (o *KiezerVervangenZaakIndienen) SetGeslotenDoor(v int32)`

SetGeslotenDoor sets GeslotenDoor field to given value.


### SetGeslotenDoorNil

`func (o *KiezerVervangenZaakIndienen) SetGeslotenDoorNil(b bool)`

 SetGeslotenDoorNil sets the value for GeslotenDoor to be an explicit nil

### UnsetGeslotenDoor
`func (o *KiezerVervangenZaakIndienen) UnsetGeslotenDoor()`

UnsetGeslotenDoor ensures that no value is present for GeslotenDoor, not even an explicit nil
### GetMogelijkeTransities

`func (o *KiezerVervangenZaakIndienen) GetMogelijkeTransities() []string`

GetMogelijkeTransities returns the MogelijkeTransities field if non-nil, zero value otherwise.

### GetMogelijkeTransitiesOk

`func (o *KiezerVervangenZaakIndienen) GetMogelijkeTransitiesOk() (*[]string, bool)`

GetMogelijkeTransitiesOk returns a tuple with the MogelijkeTransities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMogelijkeTransities

`func (o *KiezerVervangenZaakIndienen) SetMogelijkeTransities(v []string)`

SetMogelijkeTransities sets MogelijkeTransities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


