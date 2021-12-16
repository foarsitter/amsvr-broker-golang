# \AanwezigeCriteriaApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenAanwezigeCriteriaDestroy**](AanwezigeCriteriaApi.md#AansluitingenAanwezigeCriteriaDestroy) | **Delete** /api/aansluitingen/aanwezige-criteria/{id} | 
[**AansluitingenAanwezigeCriteriaPartialUpdate**](AanwezigeCriteriaApi.md#AansluitingenAanwezigeCriteriaPartialUpdate) | **Patch** /api/aansluitingen/aanwezige-criteria/{id} | 
[**AansluitingenAanwezigeCriteriaRetrieve**](AanwezigeCriteriaApi.md#AansluitingenAanwezigeCriteriaRetrieve) | **Get** /api/aansluitingen/aanwezige-criteria/{id} | 
[**AansluitingenAanwezigeCriteriaUpdate**](AanwezigeCriteriaApi.md#AansluitingenAanwezigeCriteriaUpdate) | **Put** /api/aansluitingen/aanwezige-criteria/{id} | 
[**AansluitingenSectorenAanwezigeCriteriaCreate**](AanwezigeCriteriaApi.md#AansluitingenSectorenAanwezigeCriteriaCreate) | **Post** /api/aansluitingen/sectoren/{id}/aanwezige-criteria | 
[**AansluitingenSectorenAanwezigeCriteriaList**](AanwezigeCriteriaApi.md#AansluitingenSectorenAanwezigeCriteriaList) | **Get** /api/aansluitingen/sectoren/{id}/aanwezige-criteria | 



## AansluitingenAanwezigeCriteriaDestroy

> AansluitingenAanwezigeCriteriaDestroy(ctx, id)



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


## AansluitingenAanwezigeCriteriaPartialUpdate

> AanwezigCriterium AansluitingenAanwezigeCriteriaPartialUpdate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AansluitingenAanwezigeCriteriaPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenAanwezigeCriteriaPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAanwezigCriteriumRequest** | [**optional.Interface of PatchedAanwezigCriteriumRequest**](PatchedAanwezigCriteriumRequest.md)|  | 

### Return type

[**AanwezigCriterium**](AanwezigCriterium.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenAanwezigeCriteriaRetrieve

> AanwezigCriterium AansluitingenAanwezigeCriteriaRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**AanwezigCriterium**](AanwezigCriterium.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenAanwezigeCriteriaUpdate

> AanwezigCriterium AansluitingenAanwezigeCriteriaUpdate(ctx, id, aanwezigCriteriumRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**aanwezigCriteriumRequest** | [**AanwezigCriteriumRequest**](AanwezigCriteriumRequest.md)|  | 

### Return type

[**AanwezigCriterium**](AanwezigCriterium.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenAanwezigeCriteriaCreate

> AanwezigCriterium AansluitingenSectorenAanwezigeCriteriaCreate(ctx, id, aanwezigCriteriumRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**aanwezigCriteriumRequest** | [**AanwezigCriteriumRequest**](AanwezigCriteriumRequest.md)|  | 

### Return type

[**AanwezigCriterium**](AanwezigCriterium.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenAanwezigeCriteriaList

> PaginatedAanwezigCriteriumList AansluitingenSectorenAanwezigeCriteriaList(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AansluitingenSectorenAanwezigeCriteriaListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenSectorenAanwezigeCriteriaListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**PaginatedAanwezigCriteriumList**](PaginatedAanwezigCriteriumList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

