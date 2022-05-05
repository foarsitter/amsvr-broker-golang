# Aansluiting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Naam** | **string** |  | 
**Aansluitnummer** | **string** |  | 
**Status** | [**NullableAansluitingStatusEnum**](AansluitingStatusEnum.md) | 0 &#x3D; Nieuw, 10 &#x3D; Ingediend ter aanvraag [ATSP], 20 &#x3D; Aanvraag afgekeurd [Risicobeheer], 30 &#x3D; Aanvraag goedgekeurd [Risicobeheer], 50 &#x3D; Live test ingepland [ATSP], 55 &#x3D; Goedgekeurd [AMS-servicedesk], 60 &#x3D; Actief gezet [Risicobeheer], 5 &#x3D; Migratie gestart [ATSP], 61 &#x3D; Migratie goedgekeurd [AMS-Servicedesk] | [readonly] 
**Atsp** | **int32** |  | 
**Gemeente** | Pointer to **NullableInt32** |  | [optional] 
**Straat** | Pointer to **string** |  | [optional] 
**Huisnummer** | Pointer to **string** |  | [optional] 
**Huisletter** | Pointer to **string** |  | [optional] 
**Toevoeging** | Pointer to **string** |  | [optional] 
**Postcode** | Pointer to **string** |  | [optional] 
**Plaats** | Pointer to **string** |  | [optional] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**Toestand** | [**NullableMeldingStatus**](MeldingStatus.md) | 0 &#x3D; Passief, 1 &#x3D; Actief | [readonly] 
**BagLookupId** | Pointer to **string** |  | [optional] 
**BagId** | **string** |  | [readonly] 
**BagNietBeschikbaar** | Pointer to **NullableBool** |  | [optional] 
**BagComment** | Pointer to **string** |  | [optional] 
**RijksdriehoekstelX** | Pointer to **NullableFloat64** | https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel | [optional] 
**RijksdriehoekstelY** | Pointer to **NullableFloat64** | https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel | [optional] 
**Gebruiksfunctie** | Pointer to [**NullableGebruiksfunctieEnum**](GebruiksfunctieEnum.md) |  | [optional] 
**Bijzonderheden** | Pointer to **NullableString** |  | [optional] 
**Bron** | [**NullableBronEnum**](BronEnum.md) |  | [readonly] 

## Methods

### NewAansluiting

`func NewAansluiting(id int32, naam string, aansluitnummer string, status NullableAansluitingStatusEnum, atsp int32, created time.Time, modified time.Time, toestand NullableMeldingStatus, bagId string, bron NullableBronEnum, ) *Aansluiting`

NewAansluiting instantiates a new Aansluiting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAansluitingWithDefaults

`func NewAansluitingWithDefaults() *Aansluiting`

NewAansluitingWithDefaults instantiates a new Aansluiting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Aansluiting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Aansluiting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Aansluiting) SetId(v int32)`

SetId sets Id field to given value.


### GetNaam

`func (o *Aansluiting) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *Aansluiting) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *Aansluiting) SetNaam(v string)`

SetNaam sets Naam field to given value.


### GetAansluitnummer

`func (o *Aansluiting) GetAansluitnummer() string`

GetAansluitnummer returns the Aansluitnummer field if non-nil, zero value otherwise.

### GetAansluitnummerOk

`func (o *Aansluiting) GetAansluitnummerOk() (*string, bool)`

GetAansluitnummerOk returns a tuple with the Aansluitnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluitnummer

`func (o *Aansluiting) SetAansluitnummer(v string)`

SetAansluitnummer sets Aansluitnummer field to given value.


### GetStatus

`func (o *Aansluiting) GetStatus() AansluitingStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Aansluiting) GetStatusOk() (*AansluitingStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Aansluiting) SetStatus(v AansluitingStatusEnum)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *Aansluiting) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Aansluiting) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAtsp

`func (o *Aansluiting) GetAtsp() int32`

GetAtsp returns the Atsp field if non-nil, zero value otherwise.

### GetAtspOk

`func (o *Aansluiting) GetAtspOk() (*int32, bool)`

GetAtspOk returns a tuple with the Atsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsp

`func (o *Aansluiting) SetAtsp(v int32)`

SetAtsp sets Atsp field to given value.


### GetGemeente

`func (o *Aansluiting) GetGemeente() int32`

GetGemeente returns the Gemeente field if non-nil, zero value otherwise.

### GetGemeenteOk

`func (o *Aansluiting) GetGemeenteOk() (*int32, bool)`

GetGemeenteOk returns a tuple with the Gemeente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemeente

`func (o *Aansluiting) SetGemeente(v int32)`

SetGemeente sets Gemeente field to given value.

### HasGemeente

`func (o *Aansluiting) HasGemeente() bool`

HasGemeente returns a boolean if a field has been set.

### SetGemeenteNil

`func (o *Aansluiting) SetGemeenteNil(b bool)`

 SetGemeenteNil sets the value for Gemeente to be an explicit nil

### UnsetGemeente
`func (o *Aansluiting) UnsetGemeente()`

UnsetGemeente ensures that no value is present for Gemeente, not even an explicit nil
### GetStraat

`func (o *Aansluiting) GetStraat() string`

GetStraat returns the Straat field if non-nil, zero value otherwise.

### GetStraatOk

`func (o *Aansluiting) GetStraatOk() (*string, bool)`

GetStraatOk returns a tuple with the Straat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStraat

`func (o *Aansluiting) SetStraat(v string)`

SetStraat sets Straat field to given value.

### HasStraat

`func (o *Aansluiting) HasStraat() bool`

HasStraat returns a boolean if a field has been set.

### GetHuisnummer

`func (o *Aansluiting) GetHuisnummer() string`

GetHuisnummer returns the Huisnummer field if non-nil, zero value otherwise.

### GetHuisnummerOk

`func (o *Aansluiting) GetHuisnummerOk() (*string, bool)`

GetHuisnummerOk returns a tuple with the Huisnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisnummer

`func (o *Aansluiting) SetHuisnummer(v string)`

SetHuisnummer sets Huisnummer field to given value.

### HasHuisnummer

`func (o *Aansluiting) HasHuisnummer() bool`

HasHuisnummer returns a boolean if a field has been set.

### GetHuisletter

`func (o *Aansluiting) GetHuisletter() string`

GetHuisletter returns the Huisletter field if non-nil, zero value otherwise.

### GetHuisletterOk

`func (o *Aansluiting) GetHuisletterOk() (*string, bool)`

GetHuisletterOk returns a tuple with the Huisletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisletter

`func (o *Aansluiting) SetHuisletter(v string)`

SetHuisletter sets Huisletter field to given value.

### HasHuisletter

`func (o *Aansluiting) HasHuisletter() bool`

HasHuisletter returns a boolean if a field has been set.

### GetToevoeging

`func (o *Aansluiting) GetToevoeging() string`

GetToevoeging returns the Toevoeging field if non-nil, zero value otherwise.

### GetToevoegingOk

`func (o *Aansluiting) GetToevoegingOk() (*string, bool)`

GetToevoegingOk returns a tuple with the Toevoeging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToevoeging

`func (o *Aansluiting) SetToevoeging(v string)`

SetToevoeging sets Toevoeging field to given value.

### HasToevoeging

`func (o *Aansluiting) HasToevoeging() bool`

HasToevoeging returns a boolean if a field has been set.

### GetPostcode

`func (o *Aansluiting) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *Aansluiting) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *Aansluiting) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *Aansluiting) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetPlaats

`func (o *Aansluiting) GetPlaats() string`

GetPlaats returns the Plaats field if non-nil, zero value otherwise.

### GetPlaatsOk

`func (o *Aansluiting) GetPlaatsOk() (*string, bool)`

GetPlaatsOk returns a tuple with the Plaats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaats

`func (o *Aansluiting) SetPlaats(v string)`

SetPlaats sets Plaats field to given value.

### HasPlaats

`func (o *Aansluiting) HasPlaats() bool`

HasPlaats returns a boolean if a field has been set.

### GetCreated

`func (o *Aansluiting) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Aansluiting) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Aansluiting) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Aansluiting) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Aansluiting) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Aansluiting) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetToestand

`func (o *Aansluiting) GetToestand() MeldingStatus`

GetToestand returns the Toestand field if non-nil, zero value otherwise.

### GetToestandOk

`func (o *Aansluiting) GetToestandOk() (*MeldingStatus, bool)`

GetToestandOk returns a tuple with the Toestand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToestand

`func (o *Aansluiting) SetToestand(v MeldingStatus)`

SetToestand sets Toestand field to given value.


### SetToestandNil

`func (o *Aansluiting) SetToestandNil(b bool)`

 SetToestandNil sets the value for Toestand to be an explicit nil

### UnsetToestand
`func (o *Aansluiting) UnsetToestand()`

UnsetToestand ensures that no value is present for Toestand, not even an explicit nil
### GetBagLookupId

`func (o *Aansluiting) GetBagLookupId() string`

GetBagLookupId returns the BagLookupId field if non-nil, zero value otherwise.

### GetBagLookupIdOk

`func (o *Aansluiting) GetBagLookupIdOk() (*string, bool)`

GetBagLookupIdOk returns a tuple with the BagLookupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagLookupId

`func (o *Aansluiting) SetBagLookupId(v string)`

SetBagLookupId sets BagLookupId field to given value.

### HasBagLookupId

`func (o *Aansluiting) HasBagLookupId() bool`

HasBagLookupId returns a boolean if a field has been set.

### GetBagId

`func (o *Aansluiting) GetBagId() string`

GetBagId returns the BagId field if non-nil, zero value otherwise.

### GetBagIdOk

`func (o *Aansluiting) GetBagIdOk() (*string, bool)`

GetBagIdOk returns a tuple with the BagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagId

`func (o *Aansluiting) SetBagId(v string)`

SetBagId sets BagId field to given value.


### GetBagNietBeschikbaar

`func (o *Aansluiting) GetBagNietBeschikbaar() bool`

GetBagNietBeschikbaar returns the BagNietBeschikbaar field if non-nil, zero value otherwise.

### GetBagNietBeschikbaarOk

`func (o *Aansluiting) GetBagNietBeschikbaarOk() (*bool, bool)`

GetBagNietBeschikbaarOk returns a tuple with the BagNietBeschikbaar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagNietBeschikbaar

`func (o *Aansluiting) SetBagNietBeschikbaar(v bool)`

SetBagNietBeschikbaar sets BagNietBeschikbaar field to given value.

### HasBagNietBeschikbaar

`func (o *Aansluiting) HasBagNietBeschikbaar() bool`

HasBagNietBeschikbaar returns a boolean if a field has been set.

### SetBagNietBeschikbaarNil

`func (o *Aansluiting) SetBagNietBeschikbaarNil(b bool)`

 SetBagNietBeschikbaarNil sets the value for BagNietBeschikbaar to be an explicit nil

### UnsetBagNietBeschikbaar
`func (o *Aansluiting) UnsetBagNietBeschikbaar()`

UnsetBagNietBeschikbaar ensures that no value is present for BagNietBeschikbaar, not even an explicit nil
### GetBagComment

`func (o *Aansluiting) GetBagComment() string`

GetBagComment returns the BagComment field if non-nil, zero value otherwise.

### GetBagCommentOk

`func (o *Aansluiting) GetBagCommentOk() (*string, bool)`

GetBagCommentOk returns a tuple with the BagComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagComment

`func (o *Aansluiting) SetBagComment(v string)`

SetBagComment sets BagComment field to given value.

### HasBagComment

`func (o *Aansluiting) HasBagComment() bool`

HasBagComment returns a boolean if a field has been set.

### GetRijksdriehoekstelX

`func (o *Aansluiting) GetRijksdriehoekstelX() float64`

GetRijksdriehoekstelX returns the RijksdriehoekstelX field if non-nil, zero value otherwise.

### GetRijksdriehoekstelXOk

`func (o *Aansluiting) GetRijksdriehoekstelXOk() (*float64, bool)`

GetRijksdriehoekstelXOk returns a tuple with the RijksdriehoekstelX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRijksdriehoekstelX

`func (o *Aansluiting) SetRijksdriehoekstelX(v float64)`

SetRijksdriehoekstelX sets RijksdriehoekstelX field to given value.

### HasRijksdriehoekstelX

`func (o *Aansluiting) HasRijksdriehoekstelX() bool`

HasRijksdriehoekstelX returns a boolean if a field has been set.

### SetRijksdriehoekstelXNil

`func (o *Aansluiting) SetRijksdriehoekstelXNil(b bool)`

 SetRijksdriehoekstelXNil sets the value for RijksdriehoekstelX to be an explicit nil

### UnsetRijksdriehoekstelX
`func (o *Aansluiting) UnsetRijksdriehoekstelX()`

UnsetRijksdriehoekstelX ensures that no value is present for RijksdriehoekstelX, not even an explicit nil
### GetRijksdriehoekstelY

`func (o *Aansluiting) GetRijksdriehoekstelY() float64`

GetRijksdriehoekstelY returns the RijksdriehoekstelY field if non-nil, zero value otherwise.

### GetRijksdriehoekstelYOk

`func (o *Aansluiting) GetRijksdriehoekstelYOk() (*float64, bool)`

GetRijksdriehoekstelYOk returns a tuple with the RijksdriehoekstelY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRijksdriehoekstelY

`func (o *Aansluiting) SetRijksdriehoekstelY(v float64)`

SetRijksdriehoekstelY sets RijksdriehoekstelY field to given value.

### HasRijksdriehoekstelY

`func (o *Aansluiting) HasRijksdriehoekstelY() bool`

HasRijksdriehoekstelY returns a boolean if a field has been set.

### SetRijksdriehoekstelYNil

`func (o *Aansluiting) SetRijksdriehoekstelYNil(b bool)`

 SetRijksdriehoekstelYNil sets the value for RijksdriehoekstelY to be an explicit nil

### UnsetRijksdriehoekstelY
`func (o *Aansluiting) UnsetRijksdriehoekstelY()`

UnsetRijksdriehoekstelY ensures that no value is present for RijksdriehoekstelY, not even an explicit nil
### GetGebruiksfunctie

`func (o *Aansluiting) GetGebruiksfunctie() GebruiksfunctieEnum`

GetGebruiksfunctie returns the Gebruiksfunctie field if non-nil, zero value otherwise.

### GetGebruiksfunctieOk

`func (o *Aansluiting) GetGebruiksfunctieOk() (*GebruiksfunctieEnum, bool)`

GetGebruiksfunctieOk returns a tuple with the Gebruiksfunctie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGebruiksfunctie

`func (o *Aansluiting) SetGebruiksfunctie(v GebruiksfunctieEnum)`

SetGebruiksfunctie sets Gebruiksfunctie field to given value.

### HasGebruiksfunctie

`func (o *Aansluiting) HasGebruiksfunctie() bool`

HasGebruiksfunctie returns a boolean if a field has been set.

### SetGebruiksfunctieNil

`func (o *Aansluiting) SetGebruiksfunctieNil(b bool)`

 SetGebruiksfunctieNil sets the value for Gebruiksfunctie to be an explicit nil

### UnsetGebruiksfunctie
`func (o *Aansluiting) UnsetGebruiksfunctie()`

UnsetGebruiksfunctie ensures that no value is present for Gebruiksfunctie, not even an explicit nil
### GetBijzonderheden

`func (o *Aansluiting) GetBijzonderheden() string`

GetBijzonderheden returns the Bijzonderheden field if non-nil, zero value otherwise.

### GetBijzonderhedenOk

`func (o *Aansluiting) GetBijzonderhedenOk() (*string, bool)`

GetBijzonderhedenOk returns a tuple with the Bijzonderheden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBijzonderheden

`func (o *Aansluiting) SetBijzonderheden(v string)`

SetBijzonderheden sets Bijzonderheden field to given value.

### HasBijzonderheden

`func (o *Aansluiting) HasBijzonderheden() bool`

HasBijzonderheden returns a boolean if a field has been set.

### SetBijzonderhedenNil

`func (o *Aansluiting) SetBijzonderhedenNil(b bool)`

 SetBijzonderhedenNil sets the value for Bijzonderheden to be an explicit nil

### UnsetBijzonderheden
`func (o *Aansluiting) UnsetBijzonderheden()`

UnsetBijzonderheden ensures that no value is present for Bijzonderheden, not even an explicit nil
### GetBron

`func (o *Aansluiting) GetBron() BronEnum`

GetBron returns the Bron field if non-nil, zero value otherwise.

### GetBronOk

`func (o *Aansluiting) GetBronOk() (*BronEnum, bool)`

GetBronOk returns a tuple with the Bron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBron

`func (o *Aansluiting) SetBron(v BronEnum)`

SetBron sets Bron field to given value.


### SetBronNil

`func (o *Aansluiting) SetBronNil(b bool)`

 SetBronNil sets the value for Bron to be an explicit nil

### UnsetBron
`func (o *Aansluiting) UnsetBron()`

UnsetBron ensures that no value is present for Bron, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


