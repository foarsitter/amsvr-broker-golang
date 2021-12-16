# \BrandweerApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrandweerGemeentesList**](BrandweerApi.md#BrandweerGemeentesList) | **Get** /api/brandweer/gemeentes | 
[**BrandweerGemeentesRetrieve**](BrandweerApi.md#BrandweerGemeentesRetrieve) | **Get** /api/brandweer/gemeentes/{id} | 
[**BrandweerMeldkamersCriteriaList**](BrandweerApi.md#BrandweerMeldkamersCriteriaList) | **Get** /api/brandweer/meldkamers/{id}/criteria | 
[**BrandweerMeldkamersList**](BrandweerApi.md#BrandweerMeldkamersList) | **Get** /api/brandweer/meldkamers | 
[**BrandweerMeldkamersRetrieve**](BrandweerApi.md#BrandweerMeldkamersRetrieve) | **Get** /api/brandweer/meldkamers/{id} | 
[**BrandweerRegiosList**](BrandweerApi.md#BrandweerRegiosList) | **Get** /api/brandweer/regios | 
[**BrandweerRegiosRetrieve**](BrandweerApi.md#BrandweerRegiosRetrieve) | **Get** /api/brandweer/regios/{id} | 



## BrandweerGemeentesList

> PaginatedGemeenteList BrandweerGemeentesList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BrandweerGemeentesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BrandweerGemeentesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afkorting** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **naam** | **optional.String**|  | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **plaatsen** | **optional.String**|  | 
 **provincieAfkorting** | **optional.String**|  | 
 **regio** | **optional.Int32**|  | 

### Return type

[**PaginatedGemeenteList**](PaginatedGemeenteList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerGemeentesRetrieve

> Gemeente BrandweerGemeentesRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Gemeente**](Gemeente.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerMeldkamersCriteriaList

> PaginatedCriteriumList BrandweerMeldkamersCriteriaList(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***BrandweerMeldkamersCriteriaListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BrandweerMeldkamersCriteriaListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**PaginatedCriteriumList**](PaginatedCriteriumList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerMeldkamersList

> PaginatedMeldkamerList BrandweerMeldkamersList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BrandweerMeldkamersListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BrandweerMeldkamersListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**PaginatedMeldkamerList**](PaginatedMeldkamerList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerMeldkamersRetrieve

> Meldkamer BrandweerMeldkamersRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Meldkamer**](Meldkamer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerRegiosList

> PaginatedRegioList BrandweerRegiosList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BrandweerRegiosListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BrandweerRegiosListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **meldkamer** | **optional.Int32**|  | 
 **naam** | **optional.String**|  | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**PaginatedRegioList**](PaginatedRegioList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerRegiosRetrieve

> Regio BrandweerRegiosRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Regio**](Regio.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

