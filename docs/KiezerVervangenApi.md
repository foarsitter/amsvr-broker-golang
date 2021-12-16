# \KiezerVervangenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZakenKiezerVervangenAnnulerenCreate**](KiezerVervangenApi.md#ZakenKiezerVervangenAnnulerenCreate) | **Post** /api/zaken/kiezer-vervangen/{id}/annuleren | 
[**ZakenKiezerVervangenGoedkeurenCreate**](KiezerVervangenApi.md#ZakenKiezerVervangenGoedkeurenCreate) | **Post** /api/zaken/kiezer-vervangen/{id}/goedkeuren | 
[**ZakenKiezerVervangenIndienenCreate**](KiezerVervangenApi.md#ZakenKiezerVervangenIndienenCreate) | **Post** /api/zaken/kiezer-vervangen/{aansluitnummer}/indienen | 
[**ZakenKiezerVervangenInplannenCreate**](KiezerVervangenApi.md#ZakenKiezerVervangenInplannenCreate) | **Post** /api/zaken/kiezer-vervangen/{id}/inplannen | 
[**ZakenKiezerVervangenList**](KiezerVervangenApi.md#ZakenKiezerVervangenList) | **Get** /api/zaken/kiezer-vervangen/ | 
[**ZakenKiezerVervangenRetrieve**](KiezerVervangenApi.md#ZakenKiezerVervangenRetrieve) | **Get** /api/zaken/kiezer-vervangen/{id} | 



## ZakenKiezerVervangenAnnulerenCreate

> KiezerVervangenZaak ZakenKiezerVervangenAnnulerenCreate(ctx, id, kiezerVervangenZaakRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**kiezerVervangenZaakRequest** | [**KiezerVervangenZaakRequest**](KiezerVervangenZaakRequest.md)|  | 

### Return type

[**KiezerVervangenZaak**](KiezerVervangenZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenKiezerVervangenGoedkeurenCreate

> KiezerVervangenZaak ZakenKiezerVervangenGoedkeurenCreate(ctx, id, kiezerVervangenZaakRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**kiezerVervangenZaakRequest** | [**KiezerVervangenZaakRequest**](KiezerVervangenZaakRequest.md)|  | 

### Return type

[**KiezerVervangenZaak**](KiezerVervangenZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenKiezerVervangenIndienenCreate

> KiezerVervangenZaakIndienen ZakenKiezerVervangenIndienenCreate(ctx, aansluitnummer, kiezerVervangenZaakIndienenRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
**kiezerVervangenZaakIndienenRequest** | [**KiezerVervangenZaakIndienenRequest**](KiezerVervangenZaakIndienenRequest.md)|  | 

### Return type

[**KiezerVervangenZaakIndienen**](KiezerVervangenZaakIndienen.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenKiezerVervangenInplannenCreate

> KiezerVervangenZaakInplannen ZakenKiezerVervangenInplannenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenKiezerVervangenInplannenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenKiezerVervangenInplannenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kiezerVervangenZaakInplannenRequest** | [**optional.Interface of KiezerVervangenZaakInplannenRequest**](KiezerVervangenZaakInplannenRequest.md)|  | 

### Return type

[**KiezerVervangenZaakInplannen**](KiezerVervangenZaakInplannen.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenKiezerVervangenList

> PaginatedKiezerVervangenZaakList ZakenKiezerVervangenList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ZakenKiezerVervangenListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenKiezerVervangenListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gesloten** | **optional.Bool**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **status** | **optional.String**| nieuw &#x3D; Nieuw, ingediend &#x3D; Ingediend [ATSP], ingepland &#x3D; Live test ingepland [ATSP], goedgekeurd &#x3D; Goedgekeurd [AMS-Servicedesk], doorgevoerd &#x3D; Doorgevoerd [GMS], geannuleerd &#x3D; Geannuleerd | 

### Return type

[**PaginatedKiezerVervangenZaakList**](PaginatedKiezerVervangenZaakList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenKiezerVervangenRetrieve

> KiezerVervangenZaak ZakenKiezerVervangenRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**KiezerVervangenZaak**](KiezerVervangenZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

