# \SectorenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenSectorenCreate**](SectorenApi.md#AansluitingenSectorenCreate) | **Post** /api/aansluitingen/{aansluitnummer}/sectoren | 
[**AansluitingenSectorenDestroy**](SectorenApi.md#AansluitingenSectorenDestroy) | **Delete** /api/aansluitingen/sectoren/{id} | 
[**AansluitingenSectorenList**](SectorenApi.md#AansluitingenSectorenList) | **Get** /api/aansluitingen/{aansluitnummer}/sectoren | 
[**AansluitingenSectorenRetrieveById**](SectorenApi.md#AansluitingenSectorenRetrieveById) | **Get** /api/aansluitingen/sectoren/{id} | 
[**AansluitingenSectorenToestandPartialUpdateById**](SectorenApi.md#AansluitingenSectorenToestandPartialUpdateById) | **Patch** /api/aansluitingen/sectoren/{id}/toestand | 
[**AansluitingenSectorenUpdate**](SectorenApi.md#AansluitingenSectorenUpdate) | **Put** /api/aansluitingen/sectoren/{id} | 
[**SectorList**](SectorenApi.md#SectorList) | **Get** /api/aansluitingen/sectoren | 



## AansluitingenSectorenCreate

> Sector AansluitingenSectorenCreate(ctx, aansluitnummer, sectorRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
**sectorRequest** | [**SectorRequest**](SectorRequest.md)|  | 

### Return type

[**Sector**](Sector.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenDestroy

> AansluitingenSectorenDestroy(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

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


## AansluitingenSectorenList

> PaginatedSectorList AansluitingenSectorenList(ctx, aansluitnummer, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
 **optional** | ***AansluitingenSectorenListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenSectorenListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**PaginatedSectorList**](PaginatedSectorList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenRetrieveById

> Sector AansluitingenSectorenRetrieveById(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

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


## AansluitingenSectorenToestandPartialUpdateById

> SectorToestand AansluitingenSectorenToestandPartialUpdateById(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AansluitingenSectorenToestandPartialUpdateByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenSectorenToestandPartialUpdateByIdOpts struct


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


## AansluitingenSectorenUpdate

> Sector AansluitingenSectorenUpdate(ctx, id, sectorRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**sectorRequest** | [**SectorRequest**](SectorRequest.md)|  | 

### Return type

[**Sector**](Sector.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SectorList

> PaginatedSectorList SectorList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SectorListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SectorListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bagNietBeschikbaar** | **optional.Bool**|  | 
 **huisnummer** | **optional.String**|  | 
 **isBouwdeelNull** | **optional.Bool**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **modifiedAfter** | **optional.String**|  | 
 **modifiedBefore** | **optional.String**|  | 
 **naam** | **optional.String**| Naam | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **postcode** | **optional.String**| Postcode | 

### Return type

[**PaginatedSectorList**](PaginatedSectorList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

