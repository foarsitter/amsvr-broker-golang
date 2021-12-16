# Sector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | [readonly] 
**Nummer** | **int32** |  | 
**Naam** | **string** |  | 
**Straat** | **string** |  | [optional] 
**Huisnummer** | **string** |  | [optional] 
**Huisletter** | **string** |  | [optional] 
**Toevoeging** | **string** |  | [optional] 
**Postcode** | **string** |  | [optional] 
**Plaats** | **string** |  | [optional] 
**RijksdriehoekstelX** | Pointer to **string** | https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel | [optional] 
**RijksdriehoekstelY** | Pointer to **string** | https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel | [optional] 
**Waarschuwingsadressen** | [**[]Waarschuwingsadres**](Waarschuwingsadres.md) |  | [readonly] 
**Created** | [**time.Time**](time.Time.md) |  | [readonly] 
**Modified** | [**time.Time**](time.Time.md) |  | [readonly] 
**Toestand** | Pointer to [**MeldingStatus**](MeldingStatus.md) | 0 &#x3D; Test, 1 &#x3D; Actief | [optional] 
**Gebruiksfunctie** | Pointer to [**GebruiksfunctieEnum**](GebruiksfunctieEnum.md) |  | [optional] 
**BagId** | **string** |  | [readonly] 
**BagLookupId** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

