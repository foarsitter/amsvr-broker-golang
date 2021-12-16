# ContactPersoon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Aansluiting** | **int32** |  | [readonly] 
**Voornaam** | **string** |  | 
**Tussenvoegsel** | Pointer to **string** |  | [optional] 
**Achternaam** | **string** |  | 
**IsTechnischContact** | **bool** |  | [optional] 
**Telefoonnummers** | [**[]Telefoonnummer**](Telefoonnummer.md) |  | [readonly] 
**Email** | Pointer to **string** | contactpersoon t.b.v. informatie voor de repressieve dienst (verplicht voor een technisch contact) | [optional] 
**Created** | [**time.Time**](time.Time.md) |  | [readonly] 
**Modified** | [**time.Time**](time.Time.md) |  | [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


