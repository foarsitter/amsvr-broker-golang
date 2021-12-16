# NieuweAansluitingZaak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | 
**Status** | [**NieuweAansluitingZaakStatusEnum**](NieuweAansluitingZaakStatusEnum.md) | nieuw &#x3D; Nieuw, ingediend &#x3D; Ingediend ter aanvraag [ATSP], afgekeurd &#x3D; Aanvraag afgekeurd [Risicobeheer], realisatie &#x3D; Aanvraag goedgekeurd [Risicobeheer], gms_doorgevoerd &#x3D; Doorgevoerd [GMS], technisch_gereed &#x3D; Live test inplannen [ATSP], goedgekeurd &#x3D; Goedgekeurd [AMS-Servicedesk], opgeleverd &#x3D; Actief gezet [Risicobeheer], geannuleerd &#x3D; Geannuleerd | [readonly] 
**StatusChanged** | [**time.Time**](time.Time.md) |  | [readonly] 
**Created** | [**time.Time**](time.Time.md) |  | [readonly] 
**Modified** | [**time.Time**](time.Time.md) |  | [readonly] 
**User** | **int32** |  | [readonly] 
**Gesloten** | **bool** |  | [readonly] 
**GeslotenDoor** | **int32** |  | [readonly] 
**MogelijkeTransities** | **[]string** |  | [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


