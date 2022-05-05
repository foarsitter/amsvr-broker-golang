# PatchedContactPersoonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Voornaam** | Pointer to **NullableString** |  | [optional] 
**Tussenvoegsel** | Pointer to **NullableString** |  | [optional] 
**Achternaam** | Pointer to **string** |  | [optional] 
**IsTechnischContact** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **NullableString** | contactpersoon t.b.v. informatie voor de repressieve dienst (verplicht voor een technisch contact) | [optional] 

## Methods

### NewPatchedContactPersoonRequest

`func NewPatchedContactPersoonRequest() *PatchedContactPersoonRequest`

NewPatchedContactPersoonRequest instantiates a new PatchedContactPersoonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedContactPersoonRequestWithDefaults

`func NewPatchedContactPersoonRequestWithDefaults() *PatchedContactPersoonRequest`

NewPatchedContactPersoonRequestWithDefaults instantiates a new PatchedContactPersoonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoornaam

`func (o *PatchedContactPersoonRequest) GetVoornaam() string`

GetVoornaam returns the Voornaam field if non-nil, zero value otherwise.

### GetVoornaamOk

`func (o *PatchedContactPersoonRequest) GetVoornaamOk() (*string, bool)`

GetVoornaamOk returns a tuple with the Voornaam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoornaam

`func (o *PatchedContactPersoonRequest) SetVoornaam(v string)`

SetVoornaam sets Voornaam field to given value.

### HasVoornaam

`func (o *PatchedContactPersoonRequest) HasVoornaam() bool`

HasVoornaam returns a boolean if a field has been set.

### SetVoornaamNil

`func (o *PatchedContactPersoonRequest) SetVoornaamNil(b bool)`

 SetVoornaamNil sets the value for Voornaam to be an explicit nil

### UnsetVoornaam
`func (o *PatchedContactPersoonRequest) UnsetVoornaam()`

UnsetVoornaam ensures that no value is present for Voornaam, not even an explicit nil
### GetTussenvoegsel

`func (o *PatchedContactPersoonRequest) GetTussenvoegsel() string`

GetTussenvoegsel returns the Tussenvoegsel field if non-nil, zero value otherwise.

### GetTussenvoegselOk

`func (o *PatchedContactPersoonRequest) GetTussenvoegselOk() (*string, bool)`

GetTussenvoegselOk returns a tuple with the Tussenvoegsel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTussenvoegsel

`func (o *PatchedContactPersoonRequest) SetTussenvoegsel(v string)`

SetTussenvoegsel sets Tussenvoegsel field to given value.

### HasTussenvoegsel

`func (o *PatchedContactPersoonRequest) HasTussenvoegsel() bool`

HasTussenvoegsel returns a boolean if a field has been set.

### SetTussenvoegselNil

`func (o *PatchedContactPersoonRequest) SetTussenvoegselNil(b bool)`

 SetTussenvoegselNil sets the value for Tussenvoegsel to be an explicit nil

### UnsetTussenvoegsel
`func (o *PatchedContactPersoonRequest) UnsetTussenvoegsel()`

UnsetTussenvoegsel ensures that no value is present for Tussenvoegsel, not even an explicit nil
### GetAchternaam

`func (o *PatchedContactPersoonRequest) GetAchternaam() string`

GetAchternaam returns the Achternaam field if non-nil, zero value otherwise.

### GetAchternaamOk

`func (o *PatchedContactPersoonRequest) GetAchternaamOk() (*string, bool)`

GetAchternaamOk returns a tuple with the Achternaam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchternaam

`func (o *PatchedContactPersoonRequest) SetAchternaam(v string)`

SetAchternaam sets Achternaam field to given value.

### HasAchternaam

`func (o *PatchedContactPersoonRequest) HasAchternaam() bool`

HasAchternaam returns a boolean if a field has been set.

### GetIsTechnischContact

`func (o *PatchedContactPersoonRequest) GetIsTechnischContact() bool`

GetIsTechnischContact returns the IsTechnischContact field if non-nil, zero value otherwise.

### GetIsTechnischContactOk

`func (o *PatchedContactPersoonRequest) GetIsTechnischContactOk() (*bool, bool)`

GetIsTechnischContactOk returns a tuple with the IsTechnischContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTechnischContact

`func (o *PatchedContactPersoonRequest) SetIsTechnischContact(v bool)`

SetIsTechnischContact sets IsTechnischContact field to given value.

### HasIsTechnischContact

`func (o *PatchedContactPersoonRequest) HasIsTechnischContact() bool`

HasIsTechnischContact returns a boolean if a field has been set.

### GetEmail

`func (o *PatchedContactPersoonRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedContactPersoonRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedContactPersoonRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedContactPersoonRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PatchedContactPersoonRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PatchedContactPersoonRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


