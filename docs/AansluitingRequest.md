# AansluitingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Naam** | **string** |  | 
**Aansluitnummer** | **string** |  | 
**Atsp** | **int32** |  | 
**Gemeente** | Pointer to **NullableInt32** |  | [optional] 
**Straat** | Pointer to **string** |  | [optional] 
**Huisnummer** | Pointer to **string** |  | [optional] 
**Huisletter** | Pointer to **string** |  | [optional] 
**Toevoeging** | Pointer to **string** |  | [optional] 
**Postcode** | Pointer to **string** |  | [optional] 
**Plaats** | Pointer to **string** |  | [optional] 
**BagLookupId** | Pointer to **string** |  | [optional] 
**BagNietBeschikbaar** | Pointer to **NullableBool** |  | [optional] 
**BagComment** | Pointer to **string** |  | [optional] 
**RijksdriehoekstelX** | Pointer to **NullableFloat64** | https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel | [optional] 
**RijksdriehoekstelY** | Pointer to **NullableFloat64** | https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel | [optional] 
**Gebruiksfunctie** | Pointer to [**NullableGebruiksfunctieEnum**](GebruiksfunctieEnum.md) |  | [optional] 
**Bijzonderheden** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAansluitingRequest

`func NewAansluitingRequest(naam string, aansluitnummer string, atsp int32, ) *AansluitingRequest`

NewAansluitingRequest instantiates a new AansluitingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAansluitingRequestWithDefaults

`func NewAansluitingRequestWithDefaults() *AansluitingRequest`

NewAansluitingRequestWithDefaults instantiates a new AansluitingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNaam

`func (o *AansluitingRequest) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *AansluitingRequest) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *AansluitingRequest) SetNaam(v string)`

SetNaam sets Naam field to given value.


### GetAansluitnummer

`func (o *AansluitingRequest) GetAansluitnummer() string`

GetAansluitnummer returns the Aansluitnummer field if non-nil, zero value otherwise.

### GetAansluitnummerOk

`func (o *AansluitingRequest) GetAansluitnummerOk() (*string, bool)`

GetAansluitnummerOk returns a tuple with the Aansluitnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluitnummer

`func (o *AansluitingRequest) SetAansluitnummer(v string)`

SetAansluitnummer sets Aansluitnummer field to given value.


### GetAtsp

`func (o *AansluitingRequest) GetAtsp() int32`

GetAtsp returns the Atsp field if non-nil, zero value otherwise.

### GetAtspOk

`func (o *AansluitingRequest) GetAtspOk() (*int32, bool)`

GetAtspOk returns a tuple with the Atsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsp

`func (o *AansluitingRequest) SetAtsp(v int32)`

SetAtsp sets Atsp field to given value.


### GetGemeente

`func (o *AansluitingRequest) GetGemeente() int32`

GetGemeente returns the Gemeente field if non-nil, zero value otherwise.

### GetGemeenteOk

`func (o *AansluitingRequest) GetGemeenteOk() (*int32, bool)`

GetGemeenteOk returns a tuple with the Gemeente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemeente

`func (o *AansluitingRequest) SetGemeente(v int32)`

SetGemeente sets Gemeente field to given value.

### HasGemeente

`func (o *AansluitingRequest) HasGemeente() bool`

HasGemeente returns a boolean if a field has been set.

### SetGemeenteNil

`func (o *AansluitingRequest) SetGemeenteNil(b bool)`

 SetGemeenteNil sets the value for Gemeente to be an explicit nil

### UnsetGemeente
`func (o *AansluitingRequest) UnsetGemeente()`

UnsetGemeente ensures that no value is present for Gemeente, not even an explicit nil
### GetStraat

`func (o *AansluitingRequest) GetStraat() string`

GetStraat returns the Straat field if non-nil, zero value otherwise.

### GetStraatOk

`func (o *AansluitingRequest) GetStraatOk() (*string, bool)`

GetStraatOk returns a tuple with the Straat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStraat

`func (o *AansluitingRequest) SetStraat(v string)`

SetStraat sets Straat field to given value.

### HasStraat

`func (o *AansluitingRequest) HasStraat() bool`

HasStraat returns a boolean if a field has been set.

### GetHuisnummer

`func (o *AansluitingRequest) GetHuisnummer() string`

GetHuisnummer returns the Huisnummer field if non-nil, zero value otherwise.

### GetHuisnummerOk

`func (o *AansluitingRequest) GetHuisnummerOk() (*string, bool)`

GetHuisnummerOk returns a tuple with the Huisnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisnummer

`func (o *AansluitingRequest) SetHuisnummer(v string)`

SetHuisnummer sets Huisnummer field to given value.

### HasHuisnummer

`func (o *AansluitingRequest) HasHuisnummer() bool`

HasHuisnummer returns a boolean if a field has been set.

### GetHuisletter

`func (o *AansluitingRequest) GetHuisletter() string`

GetHuisletter returns the Huisletter field if non-nil, zero value otherwise.

### GetHuisletterOk

`func (o *AansluitingRequest) GetHuisletterOk() (*string, bool)`

GetHuisletterOk returns a tuple with the Huisletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisletter

`func (o *AansluitingRequest) SetHuisletter(v string)`

SetHuisletter sets Huisletter field to given value.

### HasHuisletter

`func (o *AansluitingRequest) HasHuisletter() bool`

HasHuisletter returns a boolean if a field has been set.

### GetToevoeging

`func (o *AansluitingRequest) GetToevoeging() string`

GetToevoeging returns the Toevoeging field if non-nil, zero value otherwise.

### GetToevoegingOk

`func (o *AansluitingRequest) GetToevoegingOk() (*string, bool)`

GetToevoegingOk returns a tuple with the Toevoeging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToevoeging

`func (o *AansluitingRequest) SetToevoeging(v string)`

SetToevoeging sets Toevoeging field to given value.

### HasToevoeging

`func (o *AansluitingRequest) HasToevoeging() bool`

HasToevoeging returns a boolean if a field has been set.

### GetPostcode

`func (o *AansluitingRequest) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *AansluitingRequest) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *AansluitingRequest) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *AansluitingRequest) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetPlaats

`func (o *AansluitingRequest) GetPlaats() string`

GetPlaats returns the Plaats field if non-nil, zero value otherwise.

### GetPlaatsOk

`func (o *AansluitingRequest) GetPlaatsOk() (*string, bool)`

GetPlaatsOk returns a tuple with the Plaats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaats

`func (o *AansluitingRequest) SetPlaats(v string)`

SetPlaats sets Plaats field to given value.

### HasPlaats

`func (o *AansluitingRequest) HasPlaats() bool`

HasPlaats returns a boolean if a field has been set.

### GetBagLookupId

`func (o *AansluitingRequest) GetBagLookupId() string`

GetBagLookupId returns the BagLookupId field if non-nil, zero value otherwise.

### GetBagLookupIdOk

`func (o *AansluitingRequest) GetBagLookupIdOk() (*string, bool)`

GetBagLookupIdOk returns a tuple with the BagLookupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagLookupId

`func (o *AansluitingRequest) SetBagLookupId(v string)`

SetBagLookupId sets BagLookupId field to given value.

### HasBagLookupId

`func (o *AansluitingRequest) HasBagLookupId() bool`

HasBagLookupId returns a boolean if a field has been set.

### GetBagNietBeschikbaar

`func (o *AansluitingRequest) GetBagNietBeschikbaar() bool`

GetBagNietBeschikbaar returns the BagNietBeschikbaar field if non-nil, zero value otherwise.

### GetBagNietBeschikbaarOk

`func (o *AansluitingRequest) GetBagNietBeschikbaarOk() (*bool, bool)`

GetBagNietBeschikbaarOk returns a tuple with the BagNietBeschikbaar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagNietBeschikbaar

`func (o *AansluitingRequest) SetBagNietBeschikbaar(v bool)`

SetBagNietBeschikbaar sets BagNietBeschikbaar field to given value.

### HasBagNietBeschikbaar

`func (o *AansluitingRequest) HasBagNietBeschikbaar() bool`

HasBagNietBeschikbaar returns a boolean if a field has been set.

### SetBagNietBeschikbaarNil

`func (o *AansluitingRequest) SetBagNietBeschikbaarNil(b bool)`

 SetBagNietBeschikbaarNil sets the value for BagNietBeschikbaar to be an explicit nil

### UnsetBagNietBeschikbaar
`func (o *AansluitingRequest) UnsetBagNietBeschikbaar()`

UnsetBagNietBeschikbaar ensures that no value is present for BagNietBeschikbaar, not even an explicit nil
### GetBagComment

`func (o *AansluitingRequest) GetBagComment() string`

GetBagComment returns the BagComment field if non-nil, zero value otherwise.

### GetBagCommentOk

`func (o *AansluitingRequest) GetBagCommentOk() (*string, bool)`

GetBagCommentOk returns a tuple with the BagComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagComment

`func (o *AansluitingRequest) SetBagComment(v string)`

SetBagComment sets BagComment field to given value.

### HasBagComment

`func (o *AansluitingRequest) HasBagComment() bool`

HasBagComment returns a boolean if a field has been set.

### GetRijksdriehoekstelX

`func (o *AansluitingRequest) GetRijksdriehoekstelX() float64`

GetRijksdriehoekstelX returns the RijksdriehoekstelX field if non-nil, zero value otherwise.

### GetRijksdriehoekstelXOk

`func (o *AansluitingRequest) GetRijksdriehoekstelXOk() (*float64, bool)`

GetRijksdriehoekstelXOk returns a tuple with the RijksdriehoekstelX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRijksdriehoekstelX

`func (o *AansluitingRequest) SetRijksdriehoekstelX(v float64)`

SetRijksdriehoekstelX sets RijksdriehoekstelX field to given value.

### HasRijksdriehoekstelX

`func (o *AansluitingRequest) HasRijksdriehoekstelX() bool`

HasRijksdriehoekstelX returns a boolean if a field has been set.

### SetRijksdriehoekstelXNil

`func (o *AansluitingRequest) SetRijksdriehoekstelXNil(b bool)`

 SetRijksdriehoekstelXNil sets the value for RijksdriehoekstelX to be an explicit nil

### UnsetRijksdriehoekstelX
`func (o *AansluitingRequest) UnsetRijksdriehoekstelX()`

UnsetRijksdriehoekstelX ensures that no value is present for RijksdriehoekstelX, not even an explicit nil
### GetRijksdriehoekstelY

`func (o *AansluitingRequest) GetRijksdriehoekstelY() float64`

GetRijksdriehoekstelY returns the RijksdriehoekstelY field if non-nil, zero value otherwise.

### GetRijksdriehoekstelYOk

`func (o *AansluitingRequest) GetRijksdriehoekstelYOk() (*float64, bool)`

GetRijksdriehoekstelYOk returns a tuple with the RijksdriehoekstelY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRijksdriehoekstelY

`func (o *AansluitingRequest) SetRijksdriehoekstelY(v float64)`

SetRijksdriehoekstelY sets RijksdriehoekstelY field to given value.

### HasRijksdriehoekstelY

`func (o *AansluitingRequest) HasRijksdriehoekstelY() bool`

HasRijksdriehoekstelY returns a boolean if a field has been set.

### SetRijksdriehoekstelYNil

`func (o *AansluitingRequest) SetRijksdriehoekstelYNil(b bool)`

 SetRijksdriehoekstelYNil sets the value for RijksdriehoekstelY to be an explicit nil

### UnsetRijksdriehoekstelY
`func (o *AansluitingRequest) UnsetRijksdriehoekstelY()`

UnsetRijksdriehoekstelY ensures that no value is present for RijksdriehoekstelY, not even an explicit nil
### GetGebruiksfunctie

`func (o *AansluitingRequest) GetGebruiksfunctie() GebruiksfunctieEnum`

GetGebruiksfunctie returns the Gebruiksfunctie field if non-nil, zero value otherwise.

### GetGebruiksfunctieOk

`func (o *AansluitingRequest) GetGebruiksfunctieOk() (*GebruiksfunctieEnum, bool)`

GetGebruiksfunctieOk returns a tuple with the Gebruiksfunctie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGebruiksfunctie

`func (o *AansluitingRequest) SetGebruiksfunctie(v GebruiksfunctieEnum)`

SetGebruiksfunctie sets Gebruiksfunctie field to given value.

### HasGebruiksfunctie

`func (o *AansluitingRequest) HasGebruiksfunctie() bool`

HasGebruiksfunctie returns a boolean if a field has been set.

### SetGebruiksfunctieNil

`func (o *AansluitingRequest) SetGebruiksfunctieNil(b bool)`

 SetGebruiksfunctieNil sets the value for Gebruiksfunctie to be an explicit nil

### UnsetGebruiksfunctie
`func (o *AansluitingRequest) UnsetGebruiksfunctie()`

UnsetGebruiksfunctie ensures that no value is present for Gebruiksfunctie, not even an explicit nil
### GetBijzonderheden

`func (o *AansluitingRequest) GetBijzonderheden() string`

GetBijzonderheden returns the Bijzonderheden field if non-nil, zero value otherwise.

### GetBijzonderhedenOk

`func (o *AansluitingRequest) GetBijzonderhedenOk() (*string, bool)`

GetBijzonderhedenOk returns a tuple with the Bijzonderheden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBijzonderheden

`func (o *AansluitingRequest) SetBijzonderheden(v string)`

SetBijzonderheden sets Bijzonderheden field to given value.

### HasBijzonderheden

`func (o *AansluitingRequest) HasBijzonderheden() bool`

HasBijzonderheden returns a boolean if a field has been set.

### SetBijzonderhedenNil

`func (o *AansluitingRequest) SetBijzonderhedenNil(b bool)`

 SetBijzonderhedenNil sets the value for Bijzonderheden to be an explicit nil

### UnsetBijzonderheden
`func (o *AansluitingRequest) UnsetBijzonderheden()`

UnsetBijzonderheden ensures that no value is present for Bijzonderheden, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


