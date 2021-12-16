# AfsluitingZaak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | 
**Status** | [**AfsluitingZaakStatusEnum**](AfsluitingZaakStatusEnum.md) | nieuw &#x3D; Afsluiting, ingediend &#x3D; Ingediend [ATSP], afgewezen &#x3D; Afgewezen [Risicobeheer], goedgekeurd &#x3D; Goedgekeurd [Risicobeheer], bevestigd &#x3D; Bevestigd [ATSP], uitgevoerd &#x3D; Uitgevoerd [AMS-Servicedesk], geannuleerd &#x3D; Geannuleerd | [readonly] 
**StatusChanged** | [**time.Time**](time.Time.md) |  | [readonly] 
**Created** | [**time.Time**](time.Time.md) |  | [readonly] 
**Modified** | [**time.Time**](time.Time.md) |  | [readonly] 
**User** | **int32** |  | [readonly] 
**Gesloten** | **bool** |  | [readonly] 
**GeslotenDoor** | **int32** |  | [readonly] 
**Reden** | **string** |  | 
**RedenVanAfwijzing** | Pointer to **string** |  | [optional] 
**GeplandeDatum** | **string** | Voor wanneer staat de afsluiting gepland? | 
**DefinitieveDatum** | **string** | Wanneer wordt de aansluiting definitief verwijderd? | [readonly] 
**MogelijkeTransities** | **[]string** |  | [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


