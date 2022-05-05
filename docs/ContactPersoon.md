# ContactPersoon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | [readonly] 
**Voornaam** | Pointer to **NullableString** |  | [optional] 
**Tussenvoegsel** | Pointer to **NullableString** |  | [optional] 
**Achternaam** | **string** |  | 
**IsTechnischContact** | Pointer to **bool** |  | [optional] 
**Telefoonnummers** | [**[]Telefoonnummer**](Telefoonnummer.md) |  | [readonly] 
**Email** | Pointer to **NullableString** | contactpersoon t.b.v. informatie voor de repressieve dienst (verplicht voor een technisch contact) | [optional] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 

## Methods

### NewContactPersoon

`func NewContactPersoon(id int32, aansluiting int32, achternaam string, telefoonnummers []Telefoonnummer, created time.Time, modified time.Time, ) *ContactPersoon`

NewContactPersoon instantiates a new ContactPersoon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactPersoonWithDefaults

`func NewContactPersoonWithDefaults() *ContactPersoon`

NewContactPersoonWithDefaults instantiates a new ContactPersoon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactPersoon) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactPersoon) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactPersoon) SetId(v int32)`

SetId sets Id field to given value.


### GetAansluiting

`func (o *ContactPersoon) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *ContactPersoon) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *ContactPersoon) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.


### GetVoornaam

`func (o *ContactPersoon) GetVoornaam() string`

GetVoornaam returns the Voornaam field if non-nil, zero value otherwise.

### GetVoornaamOk

`func (o *ContactPersoon) GetVoornaamOk() (*string, bool)`

GetVoornaamOk returns a tuple with the Voornaam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoornaam

`func (o *ContactPersoon) SetVoornaam(v string)`

SetVoornaam sets Voornaam field to given value.

### HasVoornaam

`func (o *ContactPersoon) HasVoornaam() bool`

HasVoornaam returns a boolean if a field has been set.

### SetVoornaamNil

`func (o *ContactPersoon) SetVoornaamNil(b bool)`

 SetVoornaamNil sets the value for Voornaam to be an explicit nil

### UnsetVoornaam
`func (o *ContactPersoon) UnsetVoornaam()`

UnsetVoornaam ensures that no value is present for Voornaam, not even an explicit nil
### GetTussenvoegsel

`func (o *ContactPersoon) GetTussenvoegsel() string`

GetTussenvoegsel returns the Tussenvoegsel field if non-nil, zero value otherwise.

### GetTussenvoegselOk

`func (o *ContactPersoon) GetTussenvoegselOk() (*string, bool)`

GetTussenvoegselOk returns a tuple with the Tussenvoegsel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTussenvoegsel

`func (o *ContactPersoon) SetTussenvoegsel(v string)`

SetTussenvoegsel sets Tussenvoegsel field to given value.

### HasTussenvoegsel

`func (o *ContactPersoon) HasTussenvoegsel() bool`

HasTussenvoegsel returns a boolean if a field has been set.

### SetTussenvoegselNil

`func (o *ContactPersoon) SetTussenvoegselNil(b bool)`

 SetTussenvoegselNil sets the value for Tussenvoegsel to be an explicit nil

### UnsetTussenvoegsel
`func (o *ContactPersoon) UnsetTussenvoegsel()`

UnsetTussenvoegsel ensures that no value is present for Tussenvoegsel, not even an explicit nil
### GetAchternaam

`func (o *ContactPersoon) GetAchternaam() string`

GetAchternaam returns the Achternaam field if non-nil, zero value otherwise.

### GetAchternaamOk

`func (o *ContactPersoon) GetAchternaamOk() (*string, bool)`

GetAchternaamOk returns a tuple with the Achternaam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchternaam

`func (o *ContactPersoon) SetAchternaam(v string)`

SetAchternaam sets Achternaam field to given value.


### GetIsTechnischContact

`func (o *ContactPersoon) GetIsTechnischContact() bool`

GetIsTechnischContact returns the IsTechnischContact field if non-nil, zero value otherwise.

### GetIsTechnischContactOk

`func (o *ContactPersoon) GetIsTechnischContactOk() (*bool, bool)`

GetIsTechnischContactOk returns a tuple with the IsTechnischContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTechnischContact

`func (o *ContactPersoon) SetIsTechnischContact(v bool)`

SetIsTechnischContact sets IsTechnischContact field to given value.

### HasIsTechnischContact

`func (o *ContactPersoon) HasIsTechnischContact() bool`

HasIsTechnischContact returns a boolean if a field has been set.

### GetTelefoonnummers

`func (o *ContactPersoon) GetTelefoonnummers() []Telefoonnummer`

GetTelefoonnummers returns the Telefoonnummers field if non-nil, zero value otherwise.

### GetTelefoonnummersOk

`func (o *ContactPersoon) GetTelefoonnummersOk() (*[]Telefoonnummer, bool)`

GetTelefoonnummersOk returns a tuple with the Telefoonnummers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelefoonnummers

`func (o *ContactPersoon) SetTelefoonnummers(v []Telefoonnummer)`

SetTelefoonnummers sets Telefoonnummers field to given value.


### GetEmail

`func (o *ContactPersoon) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactPersoon) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactPersoon) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactPersoon) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ContactPersoon) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ContactPersoon) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCreated

`func (o *ContactPersoon) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ContactPersoon) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ContactPersoon) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *ContactPersoon) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ContactPersoon) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ContactPersoon) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


