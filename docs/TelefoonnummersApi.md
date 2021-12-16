# \TelefoonnummersApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenContactpersonenTelefoonnummersCreate**](TelefoonnummersApi.md#AansluitingenContactpersonenTelefoonnummersCreate) | **Post** /api/aansluitingen/contactpersonen/{id}/telefoonnummers | 
[**AansluitingenContactpersonenTelefoonnummersList**](TelefoonnummersApi.md#AansluitingenContactpersonenTelefoonnummersList) | **Get** /api/aansluitingen/contactpersonen/{id}/telefoonnummers | 
[**AansluitingenTelefoonnummersDestroy**](TelefoonnummersApi.md#AansluitingenTelefoonnummersDestroy) | **Delete** /api/aansluitingen/telefoonnummers/{id} | 
[**AansluitingenTelefoonnummersPartialUpdate**](TelefoonnummersApi.md#AansluitingenTelefoonnummersPartialUpdate) | **Patch** /api/aansluitingen/telefoonnummers/{id} | 
[**AansluitingenTelefoonnummersRetrieve**](TelefoonnummersApi.md#AansluitingenTelefoonnummersRetrieve) | **Get** /api/aansluitingen/telefoonnummers/{id} | 
[**AansluitingenTelefoonnummersUpdate**](TelefoonnummersApi.md#AansluitingenTelefoonnummersUpdate) | **Put** /api/aansluitingen/telefoonnummers/{id} | 



## AansluitingenContactpersonenTelefoonnummersCreate

> Telefoonnummer AansluitingenContactpersonenTelefoonnummersCreate(ctx, id, telefoonnummerRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**telefoonnummerRequest** | [**TelefoonnummerRequest**](TelefoonnummerRequest.md)|  | 

### Return type

[**Telefoonnummer**](Telefoonnummer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenContactpersonenTelefoonnummersList

> PaginatedTelefoonnummerList AansluitingenContactpersonenTelefoonnummersList(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AansluitingenContactpersonenTelefoonnummersListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenContactpersonenTelefoonnummersListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**PaginatedTelefoonnummerList**](PaginatedTelefoonnummerList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenTelefoonnummersDestroy

> AansluitingenTelefoonnummersDestroy(ctx, id)



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


## AansluitingenTelefoonnummersPartialUpdate

> Telefoonnummer AansluitingenTelefoonnummersPartialUpdate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AansluitingenTelefoonnummersPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AansluitingenTelefoonnummersPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTelefoonnummerRequest** | [**optional.Interface of PatchedTelefoonnummerRequest**](PatchedTelefoonnummerRequest.md)|  | 

### Return type

[**Telefoonnummer**](Telefoonnummer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenTelefoonnummersRetrieve

> Telefoonnummer AansluitingenTelefoonnummersRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Telefoonnummer**](Telefoonnummer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenTelefoonnummersUpdate

> Telefoonnummer AansluitingenTelefoonnummersUpdate(ctx, id, telefoonnummerRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**telefoonnummerRequest** | [**TelefoonnummerRequest**](TelefoonnummerRequest.md)|  | 

### Return type

[**Telefoonnummer**](Telefoonnummer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

