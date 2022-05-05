# SectorRequest

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

## Methods

### NewSectorRequest

`func NewSectorRequest(nummer int32, naam string, ) *SectorRequest`

NewSectorRequest instantiates a new SectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectorRequestWithDefaults

`func NewSectorRequestWithDefaults() *SectorRequest`

NewSectorRequestWithDefaults instantiates a new SectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNummer

`func (o *SectorRequest) GetNummer() int32`

GetNummer returns the Nummer field if non-nil, zero value otherwise.

### GetNummerOk

`func (o *SectorRequest) GetNummerOk() (*int32, bool)`

GetNummerOk returns a tuple with the Nummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNummer

`func (o *SectorRequest) SetNummer(v int32)`

SetNummer sets Nummer field to given value.


### GetNaam

`func (o *SectorRequest) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *SectorRequest) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *SectorRequest) SetNaam(v string)`

SetNaam sets Naam field to given value.


### GetStraat

`func (o *SectorRequest) GetStraat() string`

GetStraat returns the Straat field if non-nil, zero value otherwise.

### GetStraatOk

`func (o *SectorRequest) GetStraatOk() (*string, bool)`

GetStraatOk returns a tuple with the Straat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStraat

`func (o *SectorRequest) SetStraat(v string)`

SetStraat sets Straat field to given value.

### HasStraat

`func (o *SectorRequest) HasStraat() bool`

HasStraat returns a boolean if a field has been set.

### GetHuisnummer

`func (o *SectorRequest) GetHuisnummer() string`

GetHuisnummer returns the Huisnummer field if non-nil, zero value otherwise.

### GetHuisnummerOk

`func (o *SectorRequest) GetHuisnummerOk() (*string, bool)`

GetHuisnummerOk returns a tuple with the Huisnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisnummer

`func (o *SectorRequest) SetHuisnummer(v string)`

SetHuisnummer sets Huisnummer field to given value.

### HasHuisnummer

`func (o *SectorRequest) HasHuisnummer() bool`

HasHuisnummer returns a boolean if a field has been set.

### GetHuisletter

`func (o *SectorRequest) GetHuisletter() string`

GetHuisletter returns the Huisletter field if non-nil, zero value otherwise.

### GetHuisletterOk

`func (o *SectorRequest) GetHuisletterOk() (*string, bool)`

GetHuisletterOk returns a tuple with the Huisletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisletter

`func (o *SectorRequest) SetHuisletter(v string)`

SetHuisletter sets Huisletter field to given value.

### HasHuisletter

`func (o *SectorRequest) HasHuisletter() bool`

HasHuisletter returns a boolean if a field has been set.

### GetToevoeging

`func (o *SectorRequest) GetToevoeging() string`

GetToevoeging returns the Toevoeging field if non-nil, zero value otherwise.

### GetToevoegingOk

`func (o *SectorRequest) GetToevoegingOk() (*string, bool)`

GetToevoegingOk returns a tuple with the Toevoeging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToevoeging

`func (o *SectorRequest) SetToevoeging(v string)`

SetToevoeging sets Toevoeging field to given value.

### HasToevoeging

`func (o *SectorRequest) HasToevoeging() bool`

HasToevoeging returns a boolean if a field has been set.

### GetPostcode

`func (o *SectorRequest) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *SectorRequest) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *SectorRequest) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *SectorRequest) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetPlaats

`func (o *SectorRequest) GetPlaats() string`

GetPlaats returns the Plaats field if non-nil, zero value otherwise.

### GetPlaatsOk

`func (o *SectorRequest) GetPlaatsOk() (*string, bool)`

GetPlaatsOk returns a tuple with the Plaats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaats

`func (o *SectorRequest) SetPlaats(v string)`

SetPlaats sets Plaats field to given value.

### HasPlaats

`func (o *SectorRequest) HasPlaats() bool`

HasPlaats returns a boolean if a field has been set.

### GetRijksdriehoekstelX

`func (o *SectorRequest) GetRijksdriehoekstelX() float64`

GetRijksdriehoekstelX returns the RijksdriehoekstelX field if non-nil, zero value otherwise.

### GetRijksdriehoekstelXOk

`func (o *SectorRequest) GetRijksdriehoekstelXOk() (*float64, bool)`

GetRijksdriehoekstelXOk returns a tuple with the RijksdriehoekstelX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRijksdriehoekstelX

`func (o *SectorRequest) SetRijksdriehoekstelX(v float64)`

SetRijksdriehoekstelX sets RijksdriehoekstelX field to given value.

### HasRijksdriehoekstelX

`func (o *SectorRequest) HasRijksdriehoekstelX() bool`

HasRijksdriehoekstelX returns a boolean if a field has been set.

### SetRijksdriehoekstelXNil

`func (o *SectorRequest) SetRijksdriehoekstelXNil(b bool)`

 SetRijksdriehoekstelXNil sets the value for RijksdriehoekstelX to be an explicit nil

### UnsetRijksdriehoekstelX
`func (o *SectorRequest) UnsetRijksdriehoekstelX()`

UnsetRijksdriehoekstelX ensures that no value is present for RijksdriehoekstelX, not even an explicit nil
### GetRijksdriehoekstelY

`func (o *SectorRequest) GetRijksdriehoekstelY() float64`

GetRijksdriehoekstelY returns the RijksdriehoekstelY field if non-nil, zero value otherwise.

### GetRijksdriehoekstelYOk

`func (o *SectorRequest) GetRijksdriehoekstelYOk() (*float64, bool)`

GetRijksdriehoekstelYOk returns a tuple with the RijksdriehoekstelY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRijksdriehoekstelY

`func (o *SectorRequest) SetRijksdriehoekstelY(v float64)`

SetRijksdriehoekstelY sets RijksdriehoekstelY field to given value.

### HasRijksdriehoekstelY

`func (o *SectorRequest) HasRijksdriehoekstelY() bool`

HasRijksdriehoekstelY returns a boolean if a field has been set.

### SetRijksdriehoekstelYNil

`func (o *SectorRequest) SetRijksdriehoekstelYNil(b bool)`

 SetRijksdriehoekstelYNil sets the value for RijksdriehoekstelY to be an explicit nil

### UnsetRijksdriehoekstelY
`func (o *SectorRequest) UnsetRijksdriehoekstelY()`

UnsetRijksdriehoekstelY ensures that no value is present for RijksdriehoekstelY, not even an explicit nil
### GetToestand

`func (o *SectorRequest) GetToestand() MeldingToestand`

GetToestand returns the Toestand field if non-nil, zero value otherwise.

### GetToestandOk

`func (o *SectorRequest) GetToestandOk() (*MeldingToestand, bool)`

GetToestandOk returns a tuple with the Toestand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToestand

`func (o *SectorRequest) SetToestand(v MeldingToestand)`

SetToestand sets Toestand field to given value.

### HasToestand

`func (o *SectorRequest) HasToestand() bool`

HasToestand returns a boolean if a field has been set.

### SetToestandNil

`func (o *SectorRequest) SetToestandNil(b bool)`

 SetToestandNil sets the value for Toestand to be an explicit nil

### UnsetToestand
`func (o *SectorRequest) UnsetToestand()`

UnsetToestand ensures that no value is present for Toestand, not even an explicit nil
### GetGebruiksfunctie

`func (o *SectorRequest) GetGebruiksfunctie() GebruiksfunctieEnum`

GetGebruiksfunctie returns the Gebruiksfunctie field if non-nil, zero value otherwise.

### GetGebruiksfunctieOk

`func (o *SectorRequest) GetGebruiksfunctieOk() (*GebruiksfunctieEnum, bool)`

GetGebruiksfunctieOk returns a tuple with the Gebruiksfunctie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGebruiksfunctie

`func (o *SectorRequest) SetGebruiksfunctie(v GebruiksfunctieEnum)`

SetGebruiksfunctie sets Gebruiksfunctie field to given value.

### HasGebruiksfunctie

`func (o *SectorRequest) HasGebruiksfunctie() bool`

HasGebruiksfunctie returns a boolean if a field has been set.

### SetGebruiksfunctieNil

`func (o *SectorRequest) SetGebruiksfunctieNil(b bool)`

 SetGebruiksfunctieNil sets the value for Gebruiksfunctie to be an explicit nil

### UnsetGebruiksfunctie
`func (o *SectorRequest) UnsetGebruiksfunctie()`

UnsetGebruiksfunctie ensures that no value is present for Gebruiksfunctie, not even an explicit nil
### GetBagLookupId

`func (o *SectorRequest) GetBagLookupId() string`

GetBagLookupId returns the BagLookupId field if non-nil, zero value otherwise.

### GetBagLookupIdOk

`func (o *SectorRequest) GetBagLookupIdOk() (*string, bool)`

GetBagLookupIdOk returns a tuple with the BagLookupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagLookupId

`func (o *SectorRequest) SetBagLookupId(v string)`

SetBagLookupId sets BagLookupId field to given value.

### HasBagLookupId

`func (o *SectorRequest) HasBagLookupId() bool`

HasBagLookupId returns a boolean if a field has been set.

### GetStoringscodes

`func (o *SectorRequest) GetStoringscodes() []StoringscodesEnum`

GetStoringscodes returns the Storingscodes field if non-nil, zero value otherwise.

### GetStoringscodesOk

`func (o *SectorRequest) GetStoringscodesOk() (*[]StoringscodesEnum, bool)`

GetStoringscodesOk returns a tuple with the Storingscodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoringscodes

`func (o *SectorRequest) SetStoringscodes(v []StoringscodesEnum)`

SetStoringscodes sets Storingscodes field to given value.

### HasStoringscodes

`func (o *SectorRequest) HasStoringscodes() bool`

HasStoringscodes returns a boolean if a field has been set.

### SetStoringscodesNil

`func (o *SectorRequest) SetStoringscodesNil(b bool)`

 SetStoringscodesNil sets the value for Storingscodes to be an explicit nil

### UnsetStoringscodes
`func (o *SectorRequest) UnsetStoringscodes()`

UnsetStoringscodes ensures that no value is present for Storingscodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


