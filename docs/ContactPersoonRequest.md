# ContactPersoonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Voornaam** | Pointer to **NullableString** |  | [optional] 
**Tussenvoegsel** | Pointer to **NullableString** |  | [optional] 
**Achternaam** | **string** |  | 
**IsTechnischContact** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **NullableString** | contactpersoon t.b.v. informatie voor de repressieve dienst (verplicht voor een technisch contact) | [optional] 

## Methods

### NewContactPersoonRequest

`func NewContactPersoonRequest(achternaam string, ) *ContactPersoonRequest`

NewContactPersoonRequest instantiates a new ContactPersoonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactPersoonRequestWithDefaults

`func NewContactPersoonRequestWithDefaults() *ContactPersoonRequest`

NewContactPersoonRequestWithDefaults instantiates a new ContactPersoonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoornaam

`func (o *ContactPersoonRequest) GetVoornaam() string`

GetVoornaam returns the Voornaam field if non-nil, zero value otherwise.

### GetVoornaamOk

`func (o *ContactPersoonRequest) GetVoornaamOk() (*string, bool)`

GetVoornaamOk returns a tuple with the Voornaam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoornaam

`func (o *ContactPersoonRequest) SetVoornaam(v string)`

SetVoornaam sets Voornaam field to given value.

### HasVoornaam

`func (o *ContactPersoonRequest) HasVoornaam() bool`

HasVoornaam returns a boolean if a field has been set.

### SetVoornaamNil

`func (o *ContactPersoonRequest) SetVoornaamNil(b bool)`

 SetVoornaamNil sets the value for Voornaam to be an explicit nil

### UnsetVoornaam
`func (o *ContactPersoonRequest) UnsetVoornaam()`

UnsetVoornaam ensures that no value is present for Voornaam, not even an explicit nil
### GetTussenvoegsel

`func (o *ContactPersoonRequest) GetTussenvoegsel() string`

GetTussenvoegsel returns the Tussenvoegsel field if non-nil, zero value otherwise.

### GetTussenvoegselOk

`func (o *ContactPersoonRequest) GetTussenvoegselOk() (*string, bool)`

GetTussenvoegselOk returns a tuple with the Tussenvoegsel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTussenvoegsel

`func (o *ContactPersoonRequest) SetTussenvoegsel(v string)`

SetTussenvoegsel sets Tussenvoegsel field to given value.

### HasTussenvoegsel

`func (o *ContactPersoonRequest) HasTussenvoegsel() bool`

HasTussenvoegsel returns a boolean if a field has been set.

### SetTussenvoegselNil

`func (o *ContactPersoonRequest) SetTussenvoegselNil(b bool)`

 SetTussenvoegselNil sets the value for Tussenvoegsel to be an explicit nil

### UnsetTussenvoegsel
`func (o *ContactPersoonRequest) UnsetTussenvoegsel()`

UnsetTussenvoegsel ensures that no value is present for Tussenvoegsel, not even an explicit nil
### GetAchternaam

`func (o *ContactPersoonRequest) GetAchternaam() string`

GetAchternaam returns the Achternaam field if non-nil, zero value otherwise.

### GetAchternaamOk

`func (o *ContactPersoonRequest) GetAchternaamOk() (*string, bool)`

GetAchternaamOk returns a tuple with the Achternaam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchternaam

`func (o *ContactPersoonRequest) SetAchternaam(v string)`

SetAchternaam sets Achternaam field to given value.


### GetIsTechnischContact

`func (o *ContactPersoonRequest) GetIsTechnischContact() bool`

GetIsTechnischContact returns the IsTechnischContact field if non-nil, zero value otherwise.

### GetIsTechnischContactOk

`func (o *ContactPersoonRequest) GetIsTechnischContactOk() (*bool, bool)`

GetIsTechnischContactOk returns a tuple with the IsTechnischContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTechnischContact

`func (o *ContactPersoonRequest) SetIsTechnischContact(v bool)`

SetIsTechnischContact sets IsTechnischContact field to given value.

### HasIsTechnischContact

`func (o *ContactPersoonRequest) HasIsTechnischContact() bool`

HasIsTechnischContact returns a boolean if a field has been set.

### GetEmail

`func (o *ContactPersoonRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactPersoonRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactPersoonRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactPersoonRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ContactPersoonRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ContactPersoonRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


