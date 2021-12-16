# \AfsluitingApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZakenAfsluitingAfwijzenCreate**](AfsluitingApi.md#ZakenAfsluitingAfwijzenCreate) | **Post** /api/zaken/afsluiting/{id}/afwijzen | 
[**ZakenAfsluitingAnnulerenCreate**](AfsluitingApi.md#ZakenAfsluitingAnnulerenCreate) | **Post** /api/zaken/afsluiting/{id}/annuleren | 
[**ZakenAfsluitingBevestigenCreate**](AfsluitingApi.md#ZakenAfsluitingBevestigenCreate) | **Post** /api/zaken/afsluiting/{id}/bevestigen | 
[**ZakenAfsluitingGoedkeurenCreate**](AfsluitingApi.md#ZakenAfsluitingGoedkeurenCreate) | **Post** /api/zaken/afsluiting/{id}/goedkeuren | 
[**ZakenAfsluitingIndienenCreate**](AfsluitingApi.md#ZakenAfsluitingIndienenCreate) | **Post** /api/zaken/afsluiting/indienen | 
[**ZakenAfsluitingList**](AfsluitingApi.md#ZakenAfsluitingList) | **Get** /api/zaken/afsluiting/ | 
[**ZakenAfsluitingRetrieve**](AfsluitingApi.md#ZakenAfsluitingRetrieve) | **Get** /api/zaken/afsluiting/{id} | 
[**ZakenAfsluitingUitvoerenCreate**](AfsluitingApi.md#ZakenAfsluitingUitvoerenCreate) | **Post** /api/zaken/afsluiting/{id}/uitvoeren | 



## ZakenAfsluitingAfwijzenCreate

> AfsluitingZaak ZakenAfsluitingAfwijzenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenAfsluitingAfwijzenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenAfsluitingAfwijzenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakAfwijzenRequest** | [**optional.Interface of AfsluitingZaakAfwijzenRequest**](AfsluitingZaakAfwijzenRequest.md)|  | 

### Return type

[**AfsluitingZaak**](AfsluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenAfsluitingAnnulerenCreate

> AfsluitingZaak ZakenAfsluitingAnnulerenCreate(ctx, id, afsluitingZaakAnnulerenRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**afsluitingZaakAnnulerenRequest** | [**AfsluitingZaakAnnulerenRequest**](AfsluitingZaakAnnulerenRequest.md)|  | 

### Return type

[**AfsluitingZaak**](AfsluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenAfsluitingBevestigenCreate

> AfsluitingZaak ZakenAfsluitingBevestigenCreate(ctx, id, afsluitingZaakBevestigenRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**afsluitingZaakBevestigenRequest** | [**AfsluitingZaakBevestigenRequest**](AfsluitingZaakBevestigenRequest.md)|  | 

### Return type

[**AfsluitingZaak**](AfsluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenAfsluitingGoedkeurenCreate

> AfsluitingZaak ZakenAfsluitingGoedkeurenCreate(ctx, id, afsluitingZaakGoedkeurenRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**afsluitingZaakGoedkeurenRequest** | [**AfsluitingZaakGoedkeurenRequest**](AfsluitingZaakGoedkeurenRequest.md)|  | 

### Return type

[**AfsluitingZaak**](AfsluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenAfsluitingIndienenCreate

> AfsluitingZaak ZakenAfsluitingIndienenCreate(ctx, afsluitingZaakRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afsluitingZaakRequest** | [**AfsluitingZaakRequest**](AfsluitingZaakRequest.md)|  | 

### Return type

[**AfsluitingZaak**](AfsluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenAfsluitingList

> PaginatedAfsluitingZaakList ZakenAfsluitingList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ZakenAfsluitingListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenAfsluitingListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gesloten** | **optional.Bool**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **status** | **optional.String**| nieuw &#x3D; Afsluiting, ingediend &#x3D; Ingediend [ATSP], afgewezen &#x3D; Afgewezen [Risicobeheer], goedgekeurd &#x3D; Goedgekeurd [Risicobeheer], bevestigd &#x3D; Bevestigd [ATSP], uitgevoerd &#x3D; Uitgevoerd [AMS-Servicedesk], geannuleerd &#x3D; Geannuleerd | 

### Return type

[**PaginatedAfsluitingZaakList**](PaginatedAfsluitingZaakList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenAfsluitingRetrieve

> AfsluitingZaak ZakenAfsluitingRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**AfsluitingZaak**](AfsluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenAfsluitingUitvoerenCreate

> AfsluitingZaak ZakenAfsluitingUitvoerenCreate(ctx, id, afsluitingZaakUitvoerenRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**afsluitingZaakUitvoerenRequest** | [**AfsluitingZaakUitvoerenRequest**](AfsluitingZaakUitvoerenRequest.md)|  | 

### Return type

[**AfsluitingZaak**](AfsluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

