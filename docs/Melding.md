# Melding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Criterium** | **int32** | The naam field of the Criterium | 
**Tijd** | [**time.Time**](time.Time.md) |  | 
**Aansluitnummer** | **string** |  | [readonly] 
**Sector** | Pointer to **int32** | Reference to the number of the sector | [optional] 
**Status** | [**MeldingStatus**](MeldingStatus.md) | 0 &#x3D; Rust, 1 &#x3D; Alarm | [optional] 
**Toestand** | [**MeldingToestand**](MeldingToestand.md) | 0 &#x3D; Actief, 1 &#x3D; test, 2 &#x3D; Passief | 
**Data** | **string** |  | [optional] 
**Actueel** | **bool** |  | [optional] 
**GmsId** | Pointer to **string** |  | [optional] 
**Oorzaak** | Pointer to **string** |  | [optional] 
**MeldingId** | Pointer to **string** |  | [optional] 
**Prio** | [**PrioEnum**](PrioEnum.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


