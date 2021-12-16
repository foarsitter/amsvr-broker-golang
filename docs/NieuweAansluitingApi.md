# \NieuweAansluitingApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZakenNieuweAansluitingAfkeurenCreate**](NieuweAansluitingApi.md#ZakenNieuweAansluitingAfkeurenCreate) | **Post** /api/zaken/nieuwe-aansluiting/{id}/afkeuren | 
[**ZakenNieuweAansluitingAnnulerenCreate**](NieuweAansluitingApi.md#ZakenNieuweAansluitingAnnulerenCreate) | **Post** /api/zaken/nieuwe-aansluiting/{id}/annuleren | 
[**ZakenNieuweAansluitingGoedkeurenCreate**](NieuweAansluitingApi.md#ZakenNieuweAansluitingGoedkeurenCreate) | **Post** /api/zaken/nieuwe-aansluiting/{id}/goedkeuren | 
[**ZakenNieuweAansluitingIndienenCreate**](NieuweAansluitingApi.md#ZakenNieuweAansluitingIndienenCreate) | **Post** /api/zaken/nieuwe-aansluiting/{id}/indienen | 
[**ZakenNieuweAansluitingInplannenCreate**](NieuweAansluitingApi.md#ZakenNieuweAansluitingInplannenCreate) | **Post** /api/zaken/nieuwe-aansluiting/{id}/inplannen | 
[**ZakenNieuweAansluitingList**](NieuweAansluitingApi.md#ZakenNieuweAansluitingList) | **Get** /api/zaken/nieuwe-aansluiting/ | 
[**ZakenNieuweAansluitingOpleverenCreate**](NieuweAansluitingApi.md#ZakenNieuweAansluitingOpleverenCreate) | **Post** /api/zaken/nieuwe-aansluiting/{id}/opleveren | 
[**ZakenNieuweAansluitingRealisatieAfrondenCreate**](NieuweAansluitingApi.md#ZakenNieuweAansluitingRealisatieAfrondenCreate) | **Post** /api/zaken/nieuwe-aansluiting/{id}/realisatie_afronden | 
[**ZakenNieuweAansluitingRealisatieGoedkeurenCreate**](NieuweAansluitingApi.md#ZakenNieuweAansluitingRealisatieGoedkeurenCreate) | **Post** /api/zaken/nieuwe-aansluiting/{id}/realisatie_goedkeuren | 
[**ZakenNieuweAansluitingRetrieve**](NieuweAansluitingApi.md#ZakenNieuweAansluitingRetrieve) | **Get** /api/zaken/nieuwe-aansluiting/{id} | 



## ZakenNieuweAansluitingAfkeurenCreate

> NieuweAansluitingZaak ZakenNieuweAansluitingAfkeurenCreate(ctx, id, afsluitingZaakTransitieRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**afsluitingZaakTransitieRequest** | [**AfsluitingZaakTransitieRequest**](AfsluitingZaakTransitieRequest.md)|  | 

### Return type

[**NieuweAansluitingZaak**](NieuweAansluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenNieuweAansluitingAnnulerenCreate

> NieuweAansluitingZaak ZakenNieuweAansluitingAnnulerenCreate(ctx, id, afsluitingZaakTransitieRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**afsluitingZaakTransitieRequest** | [**AfsluitingZaakTransitieRequest**](AfsluitingZaakTransitieRequest.md)|  | 

### Return type

[**NieuweAansluitingZaak**](NieuweAansluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenNieuweAansluitingGoedkeurenCreate

> NieuweAansluitingZaak ZakenNieuweAansluitingGoedkeurenCreate(ctx, id, nieuweAfsluitingZaakGoedkeurenRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**nieuweAfsluitingZaakGoedkeurenRequest** | [**NieuweAfsluitingZaakGoedkeurenRequest**](NieuweAfsluitingZaakGoedkeurenRequest.md)|  | 

### Return type

[**NieuweAansluitingZaak**](NieuweAansluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenNieuweAansluitingIndienenCreate

> NieuweAansluitingZaak ZakenNieuweAansluitingIndienenCreate(ctx, id, afsluitingZaakTransitieRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**afsluitingZaakTransitieRequest** | [**AfsluitingZaakTransitieRequest**](AfsluitingZaakTransitieRequest.md)|  | 

### Return type

[**NieuweAansluitingZaak**](NieuweAansluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenNieuweAansluitingInplannenCreate

> NieuweAansluitingZaakLiveTest ZakenNieuweAansluitingInplannenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenNieuweAansluitingInplannenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenNieuweAansluitingInplannenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nieuweAansluitingZaakLiveTestRequest** | [**optional.Interface of NieuweAansluitingZaakLiveTestRequest**](NieuweAansluitingZaakLiveTestRequest.md)|  | 

### Return type

[**NieuweAansluitingZaakLiveTest**](NieuweAansluitingZaakLiveTest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenNieuweAansluitingList

> PaginatedNieuweAansluitingZaakList ZakenNieuweAansluitingList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ZakenNieuweAansluitingListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenNieuweAansluitingListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gesloten** | **optional.Bool**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **status** | **optional.String**| nieuw &#x3D; Nieuw, ingediend &#x3D; Ingediend ter aanvraag [ATSP], afgekeurd &#x3D; Aanvraag afgekeurd [Risicobeheer], realisatie &#x3D; Aanvraag goedgekeurd [Risicobeheer], gms_doorgevoerd &#x3D; Doorgevoerd [GMS], technisch_gereed &#x3D; Live test inplannen [ATSP], goedgekeurd &#x3D; Goedgekeurd [AMS-Servicedesk], opgeleverd &#x3D; Actief gezet [Risicobeheer], geannuleerd &#x3D; Geannuleerd | 

### Return type

[**PaginatedNieuweAansluitingZaakList**](PaginatedNieuweAansluitingZaakList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenNieuweAansluitingOpleverenCreate

> NieuweAansluitingZaak ZakenNieuweAansluitingOpleverenCreate(ctx, id, afsluitingZaakTransitieRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**afsluitingZaakTransitieRequest** | [**AfsluitingZaakTransitieRequest**](AfsluitingZaakTransitieRequest.md)|  | 

### Return type

[**NieuweAansluitingZaak**](NieuweAansluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenNieuweAansluitingRealisatieAfrondenCreate

> NieuweAansluitingZaakLiveTest ZakenNieuweAansluitingRealisatieAfrondenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenNieuweAansluitingRealisatieAfrondenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenNieuweAansluitingRealisatieAfrondenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nieuweAansluitingZaakLiveTestRequest** | [**optional.Interface of NieuweAansluitingZaakLiveTestRequest**](NieuweAansluitingZaakLiveTestRequest.md)|  | 

### Return type

[**NieuweAansluitingZaakLiveTest**](NieuweAansluitingZaakLiveTest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenNieuweAansluitingRealisatieGoedkeurenCreate

> NieuweAansluitingZaak ZakenNieuweAansluitingRealisatieGoedkeurenCreate(ctx, id, afsluitingZaakTransitieRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**afsluitingZaakTransitieRequest** | [**AfsluitingZaakTransitieRequest**](AfsluitingZaakTransitieRequest.md)|  | 

### Return type

[**NieuweAansluitingZaak**](NieuweAansluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenNieuweAansluitingRetrieve

> NieuweAansluitingZaak ZakenNieuweAansluitingRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**NieuweAansluitingZaak**](NieuweAansluitingZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

