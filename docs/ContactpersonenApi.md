# \ContactpersonenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenContactpersonenDestroy**](ContactpersonenApi.md#AansluitingenContactpersonenDestroy) | **Delete** /api/aansluitingen/contactpersonen/{id} | 
[**AansluitingenContactpersonenPartialUpdate**](ContactpersonenApi.md#AansluitingenContactpersonenPartialUpdate) | **Patch** /api/aansluitingen/contactpersonen/{id} | 
[**AansluitingenContactpersonenRetrieve**](ContactpersonenApi.md#AansluitingenContactpersonenRetrieve) | **Get** /api/aansluitingen/contactpersonen/{id} | 
[**AansluitingenContactpersonenUpdate**](ContactpersonenApi.md#AansluitingenContactpersonenUpdate) | **Put** /api/aansluitingen/contactpersonen/{id} | 
[**ContactpersoonList**](ContactpersonenApi.md#ContactpersoonList) | **Get** /api/aansluitingen/contactpersonen | 



## AansluitingenContactpersonenDestroy

> AansluitingenContactpersonenDestroy(ctx, id)



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


## AansluitingenContactpersonenPartialUpdate

> ContactPersoon AansluitingenContactpersonenPartialUpdate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AansluitingenContactpersonenPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenContactpersonenPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedContactPersoonRequest** | [**optional.Interface of PatchedContactPersoonRequest**](PatchedContactPersoonRequest.md)|  | 

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


## AansluitingenContactpersonenRetrieve

> ContactPersoon AansluitingenContactpersonenRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**ContactPersoon**](ContactPersoon.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenContactpersonenUpdate

> ContactPersoon AansluitingenContactpersonenUpdate(ctx, id, contactPersoonRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
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


## ContactpersoonList

> PaginatedContactPersoonList ContactpersoonList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContactpersoonListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContactpersoonListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAfter** | **optional.Time**|  | 
 **createdBefore** | **optional.Time**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **modifiedAfter** | **optional.Time**|  | 
 **modifiedBefore** | **optional.Time**|  | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **voornaam** | **optional.String**|  | 

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

