# MutatieZaak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | 
**Status** | [**MutatieZaakStatusEnum**](MutatieZaakStatusEnum.md) | mutatie &#x3D; Mutatie, ingediend &#x3D; Ingediend [ATSP], afgewezen &#x3D; Afgewezen [Risicobeheer], goedgekeurd &#x3D; Goedgekeurd [Risicobeheer], aangepast_broker &#x3D; Aangepast in de Broker [ATSP], gms_toegevoegd &#x3D; Toevoegingen doorgevoerd in het GMS [GMS], ingepland &#x3D; Live test ingepland [ATSP], aanpassingen_geverifieerd &#x3D; Live test goedgekeurd [AMS-Servicedesk], gms_verwijderd &#x3D; Criteria verwijderd in het GMS [AMS-Servicedesk], geannuleerd &#x3D; Geannuleerd | [readonly] 
**StatusChanged** | [**time.Time**](time.Time.md) |  | [readonly] 
**Created** | [**time.Time**](time.Time.md) |  | [readonly] 
**Modified** | [**time.Time**](time.Time.md) |  | [readonly] 
**User** | **int32** |  | [readonly] 
**Gesloten** | **bool** |  | [readonly] 
**GeslotenDoor** | **int32** |  | [readonly] 
**Type** | [**MutatieZaakTypeEnum**](MutatieZaakTypeEnum.md) |  | [optional] 
**Reden** | Pointer to **string** |  | [optional] 
**OmschrijvingVanWijziging** | **string** |  | 
**BouwdelenToevoegen** | **bool** |  | [optional] 
**MogelijkeTransities** | **[]string** |  | [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


