# Sector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | [readonly] 
**Nummer** | **int32** |  | 
**Naam** | **string** |  | 
**Straat** | Pointer to **string** |  | [optional] 
**Huisnummer** | Pointer to **string** |  | [optional] 
**Huisletter** | Pointer to **string** |  | [optional] 
**Toevoeging** | Pointer to **string** |  | [optional] 
**Postcode** | Pointer to **string** |  | [optional] 
**Plaats** | Pointer to **string** |  | [optional] 
**RijksdriehoekstelX** | Pointer to **NullableFloat64** | https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel | [optional] 
**RijksdriehoekstelY** | Pointer to **NullableFloat64** | https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel | [optional] 
**Waarschuwingsadressen** | [**[]Waarschuwingsadres**](Waarschuwingsadres.md) |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**Toestand** | Pointer to [**NullableMeldingToestand**](MeldingToestand.md) | 0 &#x3D; Test, 1 &#x3D; Actief, 2 &#x3D; Passief | [optional] 
**Gebruiksfunctie** | Pointer to [**NullableGebruiksfunctieEnum**](GebruiksfunctieEnum.md) |  | [optional] 
**BagId** | **string** |  | [readonly] 
**BagLookupId** | Pointer to **string** |  | [optional] 
**Storingscodes** | Pointer to [**[]StoringscodesEnum**](StoringscodesEnum.md) |  | [optional] 

## Methods

### NewSector

`func NewSector(id int32, aansluiting int32, nummer int32, naam string, waarschuwingsadressen []Waarschuwingsadres, created time.Time, modified time.Time, bagId string, ) *Sector`

NewSector instantiates a new Sector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectorWithDefaults

`func NewSectorWithDefaults() *Sector`

NewSectorWithDefaults instantiates a new Sector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Sector) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sector) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sector) SetId(v int32)`

SetId sets Id field to given value.


### GetAansluiting

`func (o *Sector) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *Sector) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *Sector) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.


### GetNummer

`func (o *Sector) GetNummer() int32`

GetNummer returns the Nummer field if non-nil, zero value otherwise.

### GetNummerOk

`func (o *Sector) GetNummerOk() (*int32, bool)`

GetNummerOk returns a tuple with the Nummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNummer

`func (o *Sector) SetNummer(v int32)`

SetNummer sets Nummer field to given value.


### GetNaam

`func (o *Sector) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *Sector) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *Sector) SetNaam(v string)`

SetNaam sets Naam field to given value.


### GetStraat

`func (o *Sector) GetStraat() string`

GetStraat returns the Straat field if non-nil, zero value otherwise.

### GetStraatOk

`func (o *Sector) GetStraatOk() (*string, bool)`

GetStraatOk returns a tuple with the Straat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStraat

`func (o *Sector) SetStraat(v string)`

SetStraat sets Straat field to given value.

### HasStraat

`func (o *Sector) HasStraat() bool`

HasStraat returns a boolean if a field has been set.

### GetHuisnummer

`func (o *Sector) GetHuisnummer() string`

GetHuisnummer returns the Huisnummer field if non-nil, zero value otherwise.

### GetHuisnummerOk

`func (o *Sector) GetHuisnummerOk() (*string, bool)`

GetHuisnummerOk returns a tuple with the Huisnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisnummer

`func (o *Sector) SetHuisnummer(v string)`

SetHuisnummer sets Huisnummer field to given value.

### HasHuisnummer

`func (o *Sector) HasHuisnummer() bool`

HasHuisnummer returns a boolean if a field has been set.

### GetHuisletter

`func (o *Sector) GetHuisletter() string`

GetHuisletter returns the Huisletter field if non-nil, zero value otherwise.

### GetHuisletterOk

`func (o *Sector) GetHuisletterOk() (*string, bool)`

GetHuisletterOk returns a tuple with the Huisletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisletter

`func (o *Sector) SetHuisletter(v string)`

SetHuisletter sets Huisletter field to given value.

### HasHuisletter

`func (o *Sector) HasHuisletter() bool`

HasHuisletter returns a boolean if a field has been set.

### GetToevoeging

`func (o *Sector) GetToevoeging() string`

GetToevoeging returns the Toevoeging field if non-nil, zero value otherwise.

### GetToevoegingOk

`func (o *Sector) GetToevoegingOk() (*string, bool)`

GetToevoegingOk returns a tuple with the Toevoeging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToevoeging

`func (o *Sector) SetToevoeging(v string)`

SetToevoeging sets Toevoeging field to given value.

### HasToevoeging

`func (o *Sector) HasToevoeging() bool`

HasToevoeging returns a boolean if a field has been set.

### GetPostcode

`func (o *Sector) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *Sector) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *Sector) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *Sector) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetPlaats

`func (o *Sector) GetPlaats() string`

GetPlaats returns the Plaats field if non-nil, zero value otherwise.

### GetPlaatsOk

`func (o *Sector) GetPlaatsOk() (*string, bool)`

GetPlaatsOk returns a tuple with the Plaats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaats

`func (o *Sector) SetPlaats(v string)`

SetPlaats sets Plaats field to given value.

### HasPlaats

`func (o *Sector) HasPlaats() bool`

HasPlaats returns a boolean if a field has been set.

### GetRijksdriehoekstelX

`func (o *Sector) GetRijksdriehoekstelX() float64`

GetRijksdriehoekstelX returns the RijksdriehoekstelX field if non-nil, zero value otherwise.

### GetRijksdriehoekstelXOk

`func (o *Sector) GetRijksdriehoekstelXOk() (*float64, bool)`

GetRijksdriehoekstelXOk returns a tuple with the RijksdriehoekstelX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRijksdriehoekstelX

`func (o *Sector) SetRijksdriehoekstelX(v float64)`

SetRijksdriehoekstelX sets RijksdriehoekstelX field to given value.

### HasRijksdriehoekstelX

`func (o *Sector) HasRijksdriehoekstelX() bool`

HasRijksdriehoekstelX returns a boolean if a field has been set.

### SetRijksdriehoekstelXNil

`func (o *Sector) SetRijksdriehoekstelXNil(b bool)`

 SetRijksdriehoekstelXNil sets the value for RijksdriehoekstelX to be an explicit nil

### UnsetRijksdriehoekstelX
`func (o *Sector) UnsetRijksdriehoekstelX()`

UnsetRijksdriehoekstelX ensures that no value is present for RijksdriehoekstelX, not even an explicit nil
### GetRijksdriehoekstelY

`func (o *Sector) GetRijksdriehoekstelY() float64`

GetRijksdriehoekstelY returns the RijksdriehoekstelY field if non-nil, zero value otherwise.

### GetRijksdriehoekstelYOk

`func (o *Sector) GetRijksdriehoekstelYOk() (*float64, bool)`

GetRijksdriehoekstelYOk returns a tuple with the RijksdriehoekstelY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRijksdriehoekstelY

`func (o *Sector) SetRijksdriehoekstelY(v float64)`

SetRijksdriehoekstelY sets RijksdriehoekstelY field to given value.

### HasRijksdriehoekstelY

`func (o *Sector) HasRijksdriehoekstelY() bool`

HasRijksdriehoekstelY returns a boolean if a field has been set.

### SetRijksdriehoekstelYNil

`func (o *Sector) SetRijksdriehoekstelYNil(b bool)`

 SetRijksdriehoekstelYNil sets the value for RijksdriehoekstelY to be an explicit nil

### UnsetRijksdriehoekstelY
`func (o *Sector) UnsetRijksdriehoekstelY()`

UnsetRijksdriehoekstelY ensures that no value is present for RijksdriehoekstelY, not even an explicit nil
### GetWaarschuwingsadressen

`func (o *Sector) GetWaarschuwingsadressen() []Waarschuwingsadres`

GetWaarschuwingsadressen returns the Waarschuwingsadressen field if non-nil, zero value otherwise.

### GetWaarschuwingsadressenOk

`func (o *Sector) GetWaarschuwingsadressenOk() (*[]Waarschuwingsadres, bool)`

GetWaarschuwingsadressenOk returns a tuple with the Waarschuwingsadressen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaarschuwingsadressen

`func (o *Sector) SetWaarschuwingsadressen(v []Waarschuwingsadres)`

SetWaarschuwingsadressen sets Waarschuwingsadressen field to given value.


### GetCreated

`func (o *Sector) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Sector) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Sector) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Sector) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Sector) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Sector) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetToestand

`func (o *Sector) GetToestand() MeldingToestand`

GetToestand returns the Toestand field if non-nil, zero value otherwise.

### GetToestandOk

`func (o *Sector) GetToestandOk() (*MeldingToestand, bool)`

GetToestandOk returns a tuple with the Toestand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToestand

`func (o *Sector) SetToestand(v MeldingToestand)`

SetToestand sets Toestand field to given value.

### HasToestand

`func (o *Sector) HasToestand() bool`

HasToestand returns a boolean if a field has been set.

### SetToestandNil

`func (o *Sector) SetToestandNil(b bool)`

 SetToestandNil sets the value for Toestand to be an explicit nil

### UnsetToestand
`func (o *Sector) UnsetToestand()`

UnsetToestand ensures that no value is present for Toestand, not even an explicit nil
### GetGebruiksfunctie

`func (o *Sector) GetGebruiksfunctie() GebruiksfunctieEnum`

GetGebruiksfunctie returns the Gebruiksfunctie field if non-nil, zero value otherwise.

### GetGebruiksfunctieOk

`func (o *Sector) GetGebruiksfunctieOk() (*GebruiksfunctieEnum, bool)`

GetGebruiksfunctieOk returns a tuple with the Gebruiksfunctie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGebruiksfunctie

`func (o *Sector) SetGebruiksfunctie(v GebruiksfunctieEnum)`

SetGebruiksfunctie sets Gebruiksfunctie field to given value.

### HasGebruiksfunctie

`func (o *Sector) HasGebruiksfunctie() bool`

HasGebruiksfunctie returns a boolean if a field has been set.

### SetGebruiksfunctieNil

`func (o *Sector) SetGebruiksfunctieNil(b bool)`

 SetGebruiksfunctieNil sets the value for Gebruiksfunctie to be an explicit nil

### UnsetGebruiksfunctie
`func (o *Sector) UnsetGebruiksfunctie()`

UnsetGebruiksfunctie ensures that no value is present for Gebruiksfunctie, not even an explicit nil
### GetBagId

`func (o *Sector) GetBagId() string`

GetBagId returns the BagId field if non-nil, zero value otherwise.

### GetBagIdOk

`func (o *Sector) GetBagIdOk() (*string, bool)`

GetBagIdOk returns a tuple with the BagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagId

`func (o *Sector) SetBagId(v string)`

SetBagId sets BagId field to given value.


### GetBagLookupId

`func (o *Sector) GetBagLookupId() string`

GetBagLookupId returns the BagLookupId field if non-nil, zero value otherwise.

### GetBagLookupIdOk

`func (o *Sector) GetBagLookupIdOk() (*string, bool)`

GetBagLookupIdOk returns a tuple with the BagLookupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagLookupId

`func (o *Sector) SetBagLookupId(v string)`

SetBagLookupId sets BagLookupId field to given value.

### HasBagLookupId

`func (o *Sector) HasBagLookupId() bool`

HasBagLookupId returns a boolean if a field has been set.

### GetStoringscodes

`func (o *Sector) GetStoringscodes() []StoringscodesEnum`

GetStoringscodes returns the Storingscodes field if non-nil, zero value otherwise.

### GetStoringscodesOk

`func (o *Sector) GetStoringscodesOk() (*[]StoringscodesEnum, bool)`

GetStoringscodesOk returns a tuple with the Storingscodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoringscodes

`func (o *Sector) SetStoringscodes(v []StoringscodesEnum)`

SetStoringscodes sets Storingscodes field to given value.

### HasStoringscodes

`func (o *Sector) HasStoringscodes() bool`

HasStoringscodes returns a boolean if a field has been set.

### SetStoringscodesNil

`func (o *Sector) SetStoringscodesNil(b bool)`

 SetStoringscodesNil sets the value for Storingscodes to be an explicit nil

### UnsetStoringscodes
`func (o *Sector) UnsetStoringscodes()`

UnsetStoringscodes ensures that no value is present for Storingscodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


