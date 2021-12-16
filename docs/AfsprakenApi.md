# \AfsprakenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AfsprakenBeschikbaarheidDagenRetrieve**](AfsprakenApi.md#AfsprakenBeschikbaarheidDagenRetrieve) | **Get** /api/afspraken/beschikbaarheid/{regio_code}/dagen | 
[**AfsprakenBeschikbaarheidTijdslotsRetrieve**](AfsprakenApi.md#AfsprakenBeschikbaarheidTijdslotsRetrieve) | **Get** /api/afspraken/beschikbaarheid/{regio_code}/tijdslots/{date} | 
[**AfsprakenCreateCreate**](AfsprakenApi.md#AfsprakenCreateCreate) | **Post** /api/afspraken/create/{regio_code} | 
[**AfsprakenList**](AfsprakenApi.md#AfsprakenList) | **Get** /api/afspraken/ | 
[**AfsprakenPartialUpdate**](AfsprakenApi.md#AfsprakenPartialUpdate) | **Patch** /api/afspraken/{id} | 
[**AfsprakenRetrieve**](AfsprakenApi.md#AfsprakenRetrieve) | **Get** /api/afspraken/{id} | 
[**AfsprakenUpdate**](AfsprakenApi.md#AfsprakenUpdate) | **Put** /api/afspraken/{id} | 
[**AfsprakenVerplaatsenCreate**](AfsprakenApi.md#AfsprakenVerplaatsenCreate) | **Post** /api/afspraken/{id}/verplaatsen | 



## AfsprakenBeschikbaarheidDagenRetrieve

> Bookabledays AfsprakenBeschikbaarheidDagenRetrieve(ctx, regioCode)



Haal de dagen op wanneer de AMS servicedesk beschikbaar is

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regioCode** | **string**|  | 

### Return type

[**Bookabledays**](Bookabledays.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenBeschikbaarheidTijdslotsRetrieve

> BookableTimes AfsprakenBeschikbaarheidTijdslotsRetrieve(ctx, date, regioCode)



Haal de beschikbare tijdslots op voor een dag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**date** | **string**|  | 
**regioCode** | **string**|  | 

### Return type

[**BookableTimes**](BookableTimes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenCreateCreate

> Afspraak AfsprakenCreateCreate(ctx, regioCode, afspraakRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regioCode** | **string**|  | 
**afspraakRequest** | [**AfspraakRequest**](AfspraakRequest.md)|  | 

### Return type

[**Afspraak**](Afspraak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenList

> PaginatedAfspraakList AfsprakenList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AfsprakenListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AfsprakenListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eindTijd** | **optional.Time**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **startTijd** | **optional.Time**|  | 

### Return type

[**PaginatedAfspraakList**](PaginatedAfspraakList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenPartialUpdate

> Afspraak AfsprakenPartialUpdate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AfsprakenPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AfsprakenPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAfspraakRequest** | [**optional.Interface of PatchedAfspraakRequest**](PatchedAfspraakRequest.md)|  | 

### Return type

[**Afspraak**](Afspraak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenRetrieve

> Afspraak AfsprakenRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Afspraak**](Afspraak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenUpdate

> Afspraak AfsprakenUpdate(ctx, id, afspraakRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**afspraakRequest** | [**AfspraakRequest**](AfspraakRequest.md)|  | 

### Return type

[**Afspraak**](Afspraak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenVerplaatsenCreate

> AfspraakVerplaatsen AfsprakenVerplaatsenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***AfsprakenVerplaatsenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AfsprakenVerplaatsenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afspraakVerplaatsenRequest** | [**optional.Interface of AfspraakVerplaatsenRequest**](AfspraakVerplaatsenRequest.md)|  | 

### Return type

[**AfspraakVerplaatsen**](AfspraakVerplaatsen.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

