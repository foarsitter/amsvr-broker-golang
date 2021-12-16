# MeldingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criterium** | **int32** | The naam field of the Criterium | 
**Tijd** | [**time.Time**](time.Time.md) |  | 
**Sector** | Pointer to **int32** | Reference to the number of the sector | [optional] 
**Status** | [**MeldingStatus**](MeldingStatus.md) | 0 &#x3D; Rust, 1 &#x3D; Alarm | [optional] 
**Toestand** | [**MeldingToestand**](MeldingToestand.md) | 0 &#x3D; Actief, 1 &#x3D; test, 2 &#x3D; Passief | 
**Data** | **string** |  | [optional] 
**Actueel** | **bool** |  | [optional] 
**GmsId** | Pointer to **string** |  | [optional] 
**Oorzaak** | Pointer to **string** |  | [optional] 
**FepNr** | **int32** |  | 
**VolgNr** | **int32** |  | 
**MeldingId** | Pointer to **string** |  | [optional] 
**Prio** | [**PrioEnum**](PrioEnum.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


