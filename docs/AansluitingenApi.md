# \AansluitingenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenContactpersonenCreate**](AansluitingenApi.md#AansluitingenContactpersonenCreate) | **Post** /api/aansluitingen/{aansluitnummer}/contactpersonen | 
[**AansluitingenContactpersonenList**](AansluitingenApi.md#AansluitingenContactpersonenList) | **Get** /api/aansluitingen/{aansluitnummer}/contactpersonen | 
[**AansluitingenCreate**](AansluitingenApi.md#AansluitingenCreate) | **Post** /api/aansluitingen | 
[**AansluitingenDestroy**](AansluitingenApi.md#AansluitingenDestroy) | **Delete** /api/aansluitingen/{aansluitnummer} | 
[**AansluitingenList**](AansluitingenApi.md#AansluitingenList) | **Get** /api/aansluitingen | 
[**AansluitingenLopendeZaakRetrieve**](AansluitingenApi.md#AansluitingenLopendeZaakRetrieve) | **Get** /api/aansluitingen/{aansluitnummer}/lopende-zaak | 
[**AansluitingenMeldingenCreate**](AansluitingenApi.md#AansluitingenMeldingenCreate) | **Post** /api/aansluitingen/{aansluitnummer}/meldingen | 
[**AansluitingenMeldingenList**](AansluitingenApi.md#AansluitingenMeldingenList) | **Get** /api/aansluitingen/{aansluitnummer}/meldingen | 
[**AansluitingenMigratieCreate**](AansluitingenApi.md#AansluitingenMigratieCreate) | **Post** /api/aansluitingen/{aansluitnummer}/migratie | 
[**AansluitingenRetrieve**](AansluitingenApi.md#AansluitingenRetrieve) | **Get** /api/aansluitingen/{aansluitnummer} | 
[**AansluitingenSearchRetrieve**](AansluitingenApi.md#AansluitingenSearchRetrieve) | **Get** /api/aansluitingen/search | 
[**AansluitingenSectorenRetrieve**](AansluitingenApi.md#AansluitingenSectorenRetrieve) | **Get** /api/aansluitingen/{aansluitnummer}/sectoren/{nummer} | 
[**AansluitingenSectorenToestandPartialUpdate**](AansluitingenApi.md#AansluitingenSectorenToestandPartialUpdate) | **Patch** /api/aansluitingen/{aansluitnummer}/sectoren/{nummer}/toestand | 
[**AansluitingenToestandPartialUpdate**](AansluitingenApi.md#AansluitingenToestandPartialUpdate) | **Patch** /api/aansluitingen/{aansluitnummer}/toestand | 
[**AansluitingenTransitiesPartialUpdate**](AansluitingenApi.md#AansluitingenTransitiesPartialUpdate) | **Patch** /api/aansluitingen/{aansluitnummer}/transities | 
[**AansluitingenTransitiesRetrieve**](AansluitingenApi.md#AansluitingenTransitiesRetrieve) | **Get** /api/aansluitingen/{aansluitnummer}/transities | 
[**AansluitingenUpdate**](AansluitingenApi.md#AansluitingenUpdate) | **Put** /api/aansluitingen/{aansluitnummer} | 
[**AansluitingenValidatieRetrieve**](AansluitingenApi.md#AansluitingenValidatieRetrieve) | **Get** /api/aansluitingen/{aansluitnummer}/validatie | 



## AansluitingenContactpersonenCreate

> ContactPersoon AansluitingenContactpersonenCreate(ctx, aansluitnummer, contactPersoonRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
**contactPersoonRequest** | [**ContactPersoonRequest**](ContactPersoonRequest.md)|  | 

### Return type

[**ContactPersoon**](ContactPersoon.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenContactpersonenList

> PaginatedContactPersoonList AansluitingenContactpersonenList(ctx, aansluitnummer, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
 **optional** | ***AansluitingenContactpersonenListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenContactpersonenListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**PaginatedContactPersoonList**](PaginatedContactPersoonList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenCreate

> Aansluiting AansluitingenCreate(ctx, aansluitingRequest)



 # Gegevensvalidatie adressering Om te voorkomen dat er fouten ontstaan in de adressering van aansluitingen worden adressen gevalideerd in de BAG. Hiervoor wordt de API Locatieserver van PDOK gebruikt. Deze dienst is vrijelijk te gebruiken: https://github.com/PDOK/locatieserver/wiki/API-Locatieserver  Er zijn drie verschillende wegen om een aansluiting correct toe te voegen.  1. Lookup met bag_lookup_id: via deze methode is het mogelijk om zelf suggesties op te vragen bij de PDOK Locatieserver. Vervolgens halen wij aan onze kant de gevens uit de BAG. Meegestuurde waarden worden voor  bijvoorbeeld huisnummer worden genegeerd.  2. Suggest via Postcode en huisnummer (en optioneel straat, huisletter en toevoeging): via deze methode proberen wij te matchen met de BAG. Als er een match wordt gevonden dan worden de gevens uit BAG als leidend beschouwd.  3. bag_niet_beschikbaar = True: hierbij wordt een motivatie verwacht in het bag_comment veld. Deze is verplicht. Adresvelden worden via deze weg niet gevalideerd. Wel is het noodzakelijk om de juiste gemeente op te geven. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitingRequest** | [**AansluitingRequest**](AansluitingRequest.md)|  | 

### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenDestroy

> AansluitingenDestroy(ctx, aansluitnummer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenList

> PaginatedAansluitingList AansluitingenList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AansluitingenListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aansluitnummer** | **optional.String**| Aansluitnummer | 
 **atsp** | **optional.Int32**|  | 
 **bijzonderheden** | **optional.Bool**| Heeft bijzonderheden | 
 **createdAfter** | **optional.Time**|  | 
 **createdBefore** | **optional.Time**|  | 
 **gebruiksfunctie** | **optional.String**|  | 
 **gemeente** | **optional.Int32**|  | 
 **isTestAansluiting** | **optional.Bool**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **modifiedAfter** | **optional.Time**|  | 
 **modifiedBefore** | **optional.Time**|  | 
 **naam** | **optional.String**| Naam | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **plaats** | **optional.String**| Plaats | 
 **postcode** | **optional.String**| Postcode | 
 **straat** | **optional.String**| Straat | 
 **testExpired** | **optional.Bool**| Testperiode verlopen | 

### Return type

[**PaginatedAansluitingList**](PaginatedAansluitingList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenLopendeZaakRetrieve

> Zaak AansluitingenLopendeZaakRetrieve(ctx, aansluitnummer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 

### Return type

[**Zaak**](Zaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenMeldingenCreate

> Melding AansluitingenMeldingenCreate(ctx, aansluitnummer, meldingRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
**meldingRequest** | [**MeldingRequest**](MeldingRequest.md)|  | 

### Return type

[**Melding**](Melding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenMeldingenList

> PaginatedMeldingList AansluitingenMeldingenList(ctx, aansluitnummer, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
 **optional** | ***AansluitingenMeldingenListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenMeldingenListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**PaginatedMeldingList**](PaginatedMeldingList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenMigratieCreate

> Aansluiting AansluitingenMigratieCreate(ctx, aansluitnummer, migratieRequest)



Het is mogelijk om een bestaande aansluiting over te nemen van een andere ATSP. Ter verificatie vragen wij om straat + huisnummer (zonder toevoegingen)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
**migratieRequest** | [**MigratieRequest**](MigratieRequest.md)|  | 

### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenRetrieve

> Aansluiting AansluitingenRetrieve(ctx, aansluitnummer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 

### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSearchRetrieve

> Search AansluitingenSearchRetrieve(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Search**](Search.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenRetrieve

> Sector AansluitingenSectorenRetrieve(ctx, aansluitnummer, nummer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
**nummer** | **int32**|  | 

### Return type

[**Sector**](Sector.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenToestandPartialUpdate

> SectorToestand AansluitingenSectorenToestandPartialUpdate(ctx, aansluitnummer, nummer, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
**nummer** | **int32**|  | 
 **optional** | ***AansluitingenSectorenToestandPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenSectorenToestandPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedSectorToestandRequest** | [**optional.Interface of PatchedSectorToestandRequest**](PatchedSectorToestandRequest.md)|  | 

### Return type

[**SectorToestand**](SectorToestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenToestandPartialUpdate

> AansluitingToestand AansluitingenToestandPartialUpdate(ctx, aansluitnummer, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
 **optional** | ***AansluitingenToestandPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenToestandPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAansluitingToestandRequest** | [**optional.Interface of PatchedAansluitingToestandRequest**](PatchedAansluitingToestandRequest.md)|  | 

### Return type

[**AansluitingToestand**](AansluitingToestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenTransitiesPartialUpdate

> Aansluiting AansluitingenTransitiesPartialUpdate(ctx, aansluitnummer, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
 **optional** | ***AansluitingenTransitiesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenTransitiesPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTransitionRequest** | [**optional.Interface of PatchedTransitionRequest**](PatchedTransitionRequest.md)|  | 

### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenTransitiesRetrieve

> Transition AansluitingenTransitiesRetrieve(ctx, aansluitnummer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 

### Return type

[**Transition**](Transition.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenUpdate

> Aansluiting AansluitingenUpdate(ctx, aansluitnummer, aansluitingRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
**aansluitingRequest** | [**AansluitingRequest**](AansluitingRequest.md)|  | 

### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenValidatieRetrieve

> ValidationObject AansluitingenValidatieRetrieve(ctx, aansluitnummer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 

### Return type

[**ValidationObject**](ValidationObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

