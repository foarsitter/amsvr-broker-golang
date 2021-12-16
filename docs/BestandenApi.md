# \BestandenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenBestandenCreate**](BestandenApi.md#AansluitingenBestandenCreate) | **Post** /api/aansluitingen/{aansluitnummer}/bestanden | 
[**AansluitingenBestandenDestroy**](BestandenApi.md#AansluitingenBestandenDestroy) | **Delete** /api/aansluitingen/bestanden/{id} | 
[**AansluitingenBestandenList**](BestandenApi.md#AansluitingenBestandenList) | **Get** /api/aansluitingen/{aansluitnummer}/bestanden | 
[**AansluitingenBestandenPartialUpdate**](BestandenApi.md#AansluitingenBestandenPartialUpdate) | **Patch** /api/aansluitingen/bestanden/{id} | 
[**AansluitingenBestandenRetrieve**](BestandenApi.md#AansluitingenBestandenRetrieve) | **Get** /api/aansluitingen/bestanden/{id} | 
[**AansluitingenBestandenUpdate**](BestandenApi.md#AansluitingenBestandenUpdate) | **Put** /api/aansluitingen/bestanden/{id} | 



## AansluitingenBestandenCreate

> AansluitingBestand AansluitingenBestandenCreate(ctx, aansluitnummer, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
 **optional** | ***AansluitingenBestandenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenBestandenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aansluitingBestandRequest** | [**optional.Interface of AansluitingBestandRequest**](AansluitingBestandRequest.md)|  | 

### Return type

[**AansluitingBestand**](AansluitingBestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenBestandenDestroy

> AansluitingenBestandenDestroy(ctx, id)



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


## AansluitingenBestandenList

> PaginatedAansluitingBestandList AansluitingenBestandenList(ctx, aansluitnummer, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
 **optional** | ***AansluitingenBestandenListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenBestandenListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**PaginatedAansluitingBestandList**](PaginatedAansluitingBestandList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenBestandenPartialUpdate

> AansluitingBestand AansluitingenBestandenPartialUpdate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AansluitingenBestandenPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenBestandenPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAansluitingBestandRequest** | [**optional.Interface of PatchedAansluitingBestandRequest**](PatchedAansluitingBestandRequest.md)|  | 

### Return type

[**AansluitingBestand**](AansluitingBestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenBestandenRetrieve

> AansluitingBestand AansluitingenBestandenRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**AansluitingBestand**](AansluitingBestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenBestandenUpdate

> AansluitingBestand AansluitingenBestandenUpdate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AansluitingenBestandenUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenBestandenUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aansluitingBestandRequest** | [**optional.Interface of AansluitingBestandRequest**](AansluitingBestandRequest.md)|  | 

### Return type

[**AansluitingBestand**](AansluitingBestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

