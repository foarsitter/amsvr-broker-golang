# SectorWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Toestand** | Pointer to [**NullableMeldingToestand**](MeldingToestand.md) | 0 &#x3D; Test, 1 &#x3D; Actief, 2 &#x3D; Passief | [optional] 
**Gebruiksfunctie** | Pointer to [**NullableGebruiksfunctieEnum**](GebruiksfunctieEnum.md) |  | [optional] 
**BagLookupId** | Pointer to **string** |  | [optional] 
**Storingscodes** | Pointer to [**[]StoringscodesEnum**](StoringscodesEnum.md) |  | [optional] 
**Telefoonnummers** | **[]int32** | List of telefoonnummer id&#39;s to turn into a list of Waarschuwingsadres. The order of the list is mapped to priority, starting by 1 | 

## Methods

### NewSectorWriteRequest

`func NewSectorWriteRequest(nummer int32, naam string, telefoonnummers []int32, ) *SectorWriteRequest`

NewSectorWriteRequest instantiates a new SectorWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectorWriteRequestWithDefaults

`func NewSectorWriteRequestWithDefaults() *SectorWriteRequest`

NewSectorWriteRequestWithDefaults instantiates a new SectorWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNummer

`func (o *SectorWriteRequest) GetNummer() int32`

GetNummer returns the Nummer field if non-nil, zero value otherwise.

### GetNummerOk

`func (o *SectorWriteRequest) GetNummerOk() (*int32, bool)`

GetNummerOk returns a tuple with the Nummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNummer

`func (o *SectorWriteRequest) SetNummer(v int32)`

SetNummer sets Nummer field to given value.


### GetNaam

`func (o *SectorWriteRequest) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *SectorWriteRequest) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *SectorWriteRequest) SetNaam(v string)`

SetNaam sets Naam field to given value.


### GetStraat

`func (o *SectorWriteRequest) GetStraat() string`

GetStraat returns the Straat field if non-nil, zero value otherwise.

### GetStraatOk

`func (o *SectorWriteRequest) GetStraatOk() (*string, bool)`

GetStraatOk returns a tuple with the Straat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStraat

`func (o *SectorWriteRequest) SetStraat(v string)`

SetStraat sets Straat field to given value.

### HasStraat

`func (o *SectorWriteRequest) HasStraat() bool`

HasStraat returns a boolean if a field has been set.

### GetHuisnummer

`func (o *SectorWriteRequest) GetHuisnummer() string`

GetHuisnummer returns the Huisnummer field if non-nil, zero value otherwise.

### GetHuisnummerOk

`func (o *SectorWriteRequest) GetHuisnummerOk() (*string, bool)`

GetHuisnummerOk returns a tuple with the Huisnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisnummer

`func (o *SectorWriteRequest) SetHuisnummer(v string)`

SetHuisnummer sets Huisnummer field to given value.

### HasHuisnummer

`func (o *SectorWriteRequest) HasHuisnummer() bool`

HasHuisnummer returns a boolean if a field has been set.

### GetHuisletter

`func (o *SectorWriteRequest) GetHuisletter() string`

GetHuisletter returns the Huisletter field if non-nil, zero value otherwise.

### GetHuisletterOk

`func (o *SectorWriteRequest) GetHuisletterOk() (*string, bool)`

GetHuisletterOk returns a tuple with the Huisletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisletter

`func (o *SectorWriteRequest) SetHuisletter(v string)`

SetHuisletter sets Huisletter field to given value.

### HasHuisletter

`func (o *SectorWriteRequest) HasHuisletter() bool`

HasHuisletter returns a boolean if a field has been set.

### GetToevoeging

`func (o *SectorWriteRequest) GetToevoeging() string`

GetToevoeging returns the Toevoeging field if non-nil, zero value otherwise.

### GetToevoegingOk

`func (o *SectorWriteRequest) GetToevoegingOk() (*string, bool)`

GetToevoegingOk returns a tuple with the Toevoeging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToevoeging

`func (o *SectorWriteRequest) SetToevoeging(v string)`

SetToevoeging sets Toevoeging field to given value.

### HasToevoeging

`func (o *SectorWriteRequest) HasToevoeging() bool`

HasToevoeging returns a boolean if a field has been set.

### GetPostcode

`func (o *SectorWriteRequest) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *SectorWriteRequest) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *SectorWriteRequest) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *SectorWriteRequest) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetPlaats

`func (o *SectorWriteRequest) GetPlaats() string`

GetPlaats returns the Plaats field if non-nil, zero value otherwise.

### GetPlaatsOk

`func (o *SectorWriteRequest) GetPlaatsOk() (*string, bool)`

GetPlaatsOk returns a tuple with the Plaats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaats

`func (o *SectorWriteRequest) SetPlaats(v string)`

SetPlaats sets Plaats field to given value.

### HasPlaats

`func (o *SectorWriteRequest) HasPlaats() bool`

HasPlaats returns a boolean if a field has been set.

### GetRijksdriehoekstelX

`func (o *SectorWriteRequest) GetRijksdriehoekstelX() float64`

GetRijksdriehoekstelX returns the RijksdriehoekstelX field if non-nil, zero value otherwise.

### GetRijksdriehoekstelXOk

`func (o *SectorWriteRequest) GetRijksdriehoekstelXOk() (*float64, bool)`

GetRijksdriehoekstelXOk returns a tuple with the RijksdriehoekstelX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRijksdriehoekstelX

`func (o *SectorWriteRequest) SetRijksdriehoekstelX(v float64)`

SetRijksdriehoekstelX sets RijksdriehoekstelX field to given value.

### HasRijksdriehoekstelX

`func (o *SectorWriteRequest) HasRijksdriehoekstelX() bool`

HasRijksdriehoekstelX returns a boolean if a field has been set.

### SetRijksdriehoekstelXNil

`func (o *SectorWriteRequest) SetRijksdriehoekstelXNil(b bool)`

 SetRijksdriehoekstelXNil sets the value for RijksdriehoekstelX to be an explicit nil

### UnsetRijksdriehoekstelX
`func (o *SectorWriteRequest) UnsetRijksdriehoekstelX()`

UnsetRijksdriehoekstelX ensures that no value is present for RijksdriehoekstelX, not even an explicit nil
### GetRijksdriehoekstelY

`func (o *SectorWriteRequest) GetRijksdriehoekstelY() float64`

GetRijksdriehoekstelY returns the RijksdriehoekstelY field if non-nil, zero value otherwise.

### GetRijksdriehoekstelYOk

`func (o *SectorWriteRequest) GetRijksdriehoekstelYOk() (*float64, bool)`

GetRijksdriehoekstelYOk returns a tuple with the RijksdriehoekstelY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRijksdriehoekstelY

`func (o *SectorWriteRequest) SetRijksdriehoekstelY(v float64)`

SetRijksdriehoekstelY sets RijksdriehoekstelY field to given value.

### HasRijksdriehoekstelY

`func (o *SectorWriteRequest) HasRijksdriehoekstelY() bool`

HasRijksdriehoekstelY returns a boolean if a field has been set.

### SetRijksdriehoekstelYNil

`func (o *SectorWriteRequest) SetRijksdriehoekstelYNil(b bool)`

 SetRijksdriehoekstelYNil sets the value for RijksdriehoekstelY to be an explicit nil

### UnsetRijksdriehoekstelY
`func (o *SectorWriteRequest) UnsetRijksdriehoekstelY()`

UnsetRijksdriehoekstelY ensures that no value is present for RijksdriehoekstelY, not even an explicit nil
### GetToestand

`func (o *SectorWriteRequest) GetToestand() MeldingToestand`

GetToestand returns the Toestand field if non-nil, zero value otherwise.

### GetToestandOk

`func (o *SectorWriteRequest) GetToestandOk() (*MeldingToestand, bool)`

GetToestandOk returns a tuple with the Toestand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToestand

`func (o *SectorWriteRequest) SetToestand(v MeldingToestand)`

SetToestand sets Toestand field to given value.

### HasToestand

`func (o *SectorWriteRequest) HasToestand() bool`

HasToestand returns a boolean if a field has been set.

### SetToestandNil

`func (o *SectorWriteRequest) SetToestandNil(b bool)`

 SetToestandNil sets the value for Toestand to be an explicit nil

### UnsetToestand
`func (o *SectorWriteRequest) UnsetToestand()`

UnsetToestand ensures that no value is present for Toestand, not even an explicit nil
### GetGebruiksfunctie

`func (o *SectorWriteRequest) GetGebruiksfunctie() GebruiksfunctieEnum`

GetGebruiksfunctie returns the Gebruiksfunctie field if non-nil, zero value otherwise.

### GetGebruiksfunctieOk

`func (o *SectorWriteRequest) GetGebruiksfunctieOk() (*GebruiksfunctieEnum, bool)`

GetGebruiksfunctieOk returns a tuple with the Gebruiksfunctie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGebruiksfunctie

`func (o *SectorWriteRequest) SetGebruiksfunctie(v GebruiksfunctieEnum)`

SetGebruiksfunctie sets Gebruiksfunctie field to given value.

### HasGebruiksfunctie

`func (o *SectorWriteRequest) HasGebruiksfunctie() bool`

HasGebruiksfunctie returns a boolean if a field has been set.

### SetGebruiksfunctieNil

`func (o *SectorWriteRequest) SetGebruiksfunctieNil(b bool)`

 SetGebruiksfunctieNil sets the value for Gebruiksfunctie to be an explicit nil

### UnsetGebruiksfunctie
`func (o *SectorWriteRequest) UnsetGebruiksfunctie()`

UnsetGebruiksfunctie ensures that no value is present for Gebruiksfunctie, not even an explicit nil
### GetBagLookupId

`func (o *SectorWriteRequest) GetBagLookupId() string`

GetBagLookupId returns the BagLookupId field if non-nil, zero value otherwise.

### GetBagLookupIdOk

`func (o *SectorWriteRequest) GetBagLookupIdOk() (*string, bool)`

GetBagLookupIdOk returns a tuple with the BagLookupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagLookupId

`func (o *SectorWriteRequest) SetBagLookupId(v string)`

SetBagLookupId sets BagLookupId field to given value.

### HasBagLookupId

`func (o *SectorWriteRequest) HasBagLookupId() bool`

HasBagLookupId returns a boolean if a field has been set.

### GetStoringscodes

`func (o *SectorWriteRequest) GetStoringscodes() []StoringscodesEnum`

GetStoringscodes returns the Storingscodes field if non-nil, zero value otherwise.

### GetStoringscodesOk

`func (o *SectorWriteRequest) GetStoringscodesOk() (*[]StoringscodesEnum, bool)`

GetStoringscodesOk returns a tuple with the Storingscodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoringscodes

`func (o *SectorWriteRequest) SetStoringscodes(v []StoringscodesEnum)`

SetStoringscodes sets Storingscodes field to given value.

### HasStoringscodes

`func (o *SectorWriteRequest) HasStoringscodes() bool`

HasStoringscodes returns a boolean if a field has been set.

### SetStoringscodesNil

`func (o *SectorWriteRequest) SetStoringscodesNil(b bool)`

 SetStoringscodesNil sets the value for Storingscodes to be an explicit nil

### UnsetStoringscodes
`func (o *SectorWriteRequest) UnsetStoringscodes()`

UnsetStoringscodes ensures that no value is present for Storingscodes, not even an explicit nil
### GetTelefoonnummers

`func (o *SectorWriteRequest) GetTelefoonnummers() []int32`

GetTelefoonnummers returns the Telefoonnummers field if non-nil, zero value otherwise.

### GetTelefoonnummersOk

`func (o *SectorWriteRequest) GetTelefoonnummersOk() (*[]int32, bool)`

GetTelefoonnummersOk returns a tuple with the Telefoonnummers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelefoonnummers

`func (o *SectorWriteRequest) SetTelefoonnummers(v []int32)`

SetTelefoonnummers sets Telefoonnummers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


