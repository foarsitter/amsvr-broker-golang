# Aansluiting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Naam** | **string** |  | 
**Aansluitnummer** | **string** |  | 
**Status** | [**AansluitingStatusEnum**](AansluitingStatusEnum.md) | 0 &#x3D; Nieuw, 10 &#x3D; Ingediend ter aanvraag [ATSP], 20 &#x3D; Aanvraag afgekeurd [Risicobeheer], 30 &#x3D; Aanvraag goedgekeurd [Risicobeheer], 50 &#x3D; Live test ingepland [ATSP], 55 &#x3D; Goedgekeurd [AMS-servicedesk], 60 &#x3D; Actief gezet [Risicobeheer], 5 &#x3D; Migratie gestart [ATSP], 61 &#x3D; Migratie goedgekeurd [AMS-Servicedesk] | [readonly] 
**Atsp** | **int32** |  | 
**Gemeente** | Pointer to **int32** |  | [optional] 
**Straat** | **string** |  | [optional] 
**Huisnummer** | **string** |  | [optional] 
**Huisletter** | **string** |  | [optional] 
**Toevoeging** | **string** |  | [optional] 
**Postcode** | **string** |  | [optional] 
**Plaats** | **string** |  | [optional] 
**Created** | [**time.Time**](time.Time.md) |  | [readonly] 
**Modified** | [**time.Time**](time.Time.md) |  | [readonly] 
**Toestand** | [**MeldingStatus**](MeldingStatus.md) | 0 &#x3D; Passief, 1 &#x3D; Actief | [readonly] 
**BagLookupId** | **string** |  | [optional] 
**BagId** | **string** |  | [readonly] 
**BagNietBeschikbaar** | Pointer to **bool** |  | [optional] 
**BagComment** | **string** |  | [optional] 
**RijksdriehoekstelX** | Pointer to **string** | https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel | [optional] 
**RijksdriehoekstelY** | Pointer to **string** | https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel | [optional] 
**Gebruiksfunctie** | Pointer to [**GebruiksfunctieEnum**](GebruiksfunctieEnum.md) |  | [optional] 
**Bijzonderheden** | Pointer to **string** |  | [optional] 
**Bron** | [**BronEnum**](BronEnum.md) |  | [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


