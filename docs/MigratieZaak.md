# MigratieZaak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | 
**Status** | [**MigratieZaakStatusEnum**](MigratieZaakStatusEnum.md) | nieuw &#x3D; Nieuw, ingediend &#x3D; Ingediend [ATSP], ingepland &#x3D; Live test ingepland [ATSP], gms_doorgevoerd &#x3D; Doorgevoerd [GMS], goedgekeurd &#x3D; Goedgekeurd [AMS-Servicedesk], geannuleerd &#x3D; Geannuleerd, gms_verwijderd &#x3D; Latende aansluiting verwijderd [GMS] | [readonly] 
**StatusChanged** | [**time.Time**](time.Time.md) |  | [readonly] 
**Created** | [**time.Time**](time.Time.md) |  | [readonly] 
**Modified** | [**time.Time**](time.Time.md) |  | [readonly] 
**User** | **int32** |  | [readonly] 
**Gesloten** | **bool** |  | [readonly] 
**GeslotenDoor** | **int32** |  | [readonly] 
**MogelijkeTransities** | **[]string** |  | [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


