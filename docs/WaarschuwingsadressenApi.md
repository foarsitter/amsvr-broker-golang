# \WaarschuwingsadressenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenSectorenWaarschuwingsadressenCreate**](WaarschuwingsadressenApi.md#AansluitingenSectorenWaarschuwingsadressenCreate) | **Post** /api/aansluitingen/sectoren/{id}/waarschuwingsadressen | 
[**AansluitingenSectorenWaarschuwingsadressenList**](WaarschuwingsadressenApi.md#AansluitingenSectorenWaarschuwingsadressenList) | **Get** /api/aansluitingen/sectoren/{id}/waarschuwingsadressen | 
[**AansluitingenWaarschuwingsadressenDestroy**](WaarschuwingsadressenApi.md#AansluitingenWaarschuwingsadressenDestroy) | **Delete** /api/aansluitingen/waarschuwingsadressen/{id} | 
[**AansluitingenWaarschuwingsadressenPartialUpdate**](WaarschuwingsadressenApi.md#AansluitingenWaarschuwingsadressenPartialUpdate) | **Patch** /api/aansluitingen/waarschuwingsadressen/{id} | 
[**AansluitingenWaarschuwingsadressenRetrieve**](WaarschuwingsadressenApi.md#AansluitingenWaarschuwingsadressenRetrieve) | **Get** /api/aansluitingen/waarschuwingsadressen/{id} | 
[**AansluitingenWaarschuwingsadressenUpdate**](WaarschuwingsadressenApi.md#AansluitingenWaarschuwingsadressenUpdate) | **Put** /api/aansluitingen/waarschuwingsadressen/{id} | 



## AansluitingenSectorenWaarschuwingsadressenCreate

> Waarschuwingsadres AansluitingenSectorenWaarschuwingsadressenCreate(ctx, id, waarschuwingsadresRequest)



Waarschuwingsadres toevoegen

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**waarschuwingsadresRequest** | [**WaarschuwingsadresRequest**](WaarschuwingsadresRequest.md)|  | 

### Return type

[**Waarschuwingsadres**](Waarschuwingsadres.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenWaarschuwingsadressenList

> PaginatedWaarschuwingsadresList AansluitingenSectorenWaarschuwingsadressenList(ctx, id, optional)



Overzicht van waarschuwingsadressen van de opgegeven sector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| A unique integer value identifying this Sector | 
 **optional** | ***AansluitingenSectorenWaarschuwingsadressenListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenSectorenWaarschuwingsadressenListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**PaginatedWaarschuwingsadresList**](PaginatedWaarschuwingsadresList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenWaarschuwingsadressenDestroy

> AansluitingenWaarschuwingsadressenDestroy(ctx, id)



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


## AansluitingenWaarschuwingsadressenPartialUpdate

> Waarschuwingsadres AansluitingenWaarschuwingsadressenPartialUpdate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AansluitingenWaarschuwingsadressenPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenWaarschuwingsadressenPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWaarschuwingsadresRequest** | [**optional.Interface of PatchedWaarschuwingsadresRequest**](PatchedWaarschuwingsadresRequest.md)|  | 

### Return type

[**Waarschuwingsadres**](Waarschuwingsadres.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenWaarschuwingsadressenRetrieve

> Waarschuwingsadres AansluitingenWaarschuwingsadressenRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Waarschuwingsadres**](Waarschuwingsadres.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenWaarschuwingsadressenUpdate

> Waarschuwingsadres AansluitingenWaarschuwingsadressenUpdate(ctx, id, waarschuwingsadresRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**waarschuwingsadresRequest** | [**WaarschuwingsadresRequest**](WaarschuwingsadresRequest.md)|  | 

### Return type

[**Waarschuwingsadres**](Waarschuwingsadres.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

