# \MutatieApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZakenMutatieAangepastBrokerCreate**](MutatieApi.md#ZakenMutatieAangepastBrokerCreate) | **Post** /api/zaken/mutatie/{id}/aangepast_broker | 
[**ZakenMutatieAccorderenCreate**](MutatieApi.md#ZakenMutatieAccorderenCreate) | **Post** /api/zaken/mutatie/{id}/accorderen | 
[**ZakenMutatieAfwijzenCreate**](MutatieApi.md#ZakenMutatieAfwijzenCreate) | **Post** /api/zaken/mutatie/{id}/afwijzen | 
[**ZakenMutatieAnnulerenCreate**](MutatieApi.md#ZakenMutatieAnnulerenCreate) | **Post** /api/zaken/mutatie/{id}/annuleren | 
[**ZakenMutatieGoedkeurenCreate**](MutatieApi.md#ZakenMutatieGoedkeurenCreate) | **Post** /api/zaken/mutatie/{id}/goedkeuren | 
[**ZakenMutatieIndienenCreate**](MutatieApi.md#ZakenMutatieIndienenCreate) | **Post** /api/zaken/mutatie/indienen | 
[**ZakenMutatieInplannenCreate**](MutatieApi.md#ZakenMutatieInplannenCreate) | **Post** /api/zaken/mutatie/{id}/inplannen | 
[**ZakenMutatieList**](MutatieApi.md#ZakenMutatieList) | **Get** /api/zaken/mutatie/ | 
[**ZakenMutatieRetrieve**](MutatieApi.md#ZakenMutatieRetrieve) | **Get** /api/zaken/mutatie/{id} | 
[**ZakenMutatieTechnischGereedCreate**](MutatieApi.md#ZakenMutatieTechnischGereedCreate) | **Post** /api/zaken/mutatie/{id}/technisch-gereed | 



## ZakenMutatieAangepastBrokerCreate

> MutatieZaak ZakenMutatieAangepastBrokerCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenMutatieAangepastBrokerCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenMutatieAangepastBrokerCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakGoedkeurenRequest** | [**optional.Interface of MutatieZaakGoedkeurenRequest**](MutatieZaakGoedkeurenRequest.md)|  | 

### Return type

[**MutatieZaak**](MutatieZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMutatieAccorderenCreate

> MutatieZaak ZakenMutatieAccorderenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenMutatieAccorderenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenMutatieAccorderenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakGoedkeurenRequest** | [**optional.Interface of MutatieZaakGoedkeurenRequest**](MutatieZaakGoedkeurenRequest.md)|  | 

### Return type

[**MutatieZaak**](MutatieZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMutatieAfwijzenCreate

> MutatieZaak ZakenMutatieAfwijzenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenMutatieAfwijzenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenMutatieAfwijzenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakGoedkeurenRequest** | [**optional.Interface of MutatieZaakGoedkeurenRequest**](MutatieZaakGoedkeurenRequest.md)|  | 

### Return type

[**MutatieZaak**](MutatieZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMutatieAnnulerenCreate

> MutatieZaak ZakenMutatieAnnulerenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenMutatieAnnulerenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenMutatieAnnulerenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakGoedkeurenRequest** | [**optional.Interface of MutatieZaakGoedkeurenRequest**](MutatieZaakGoedkeurenRequest.md)|  | 

### Return type

[**MutatieZaak**](MutatieZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMutatieGoedkeurenCreate

> MutatieZaak ZakenMutatieGoedkeurenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenMutatieGoedkeurenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenMutatieGoedkeurenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakGoedkeurenRequest** | [**optional.Interface of MutatieZaakGoedkeurenRequest**](MutatieZaakGoedkeurenRequest.md)|  | 

### Return type

[**MutatieZaak**](MutatieZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMutatieIndienenCreate

> MutatieZaak ZakenMutatieIndienenCreate(ctx, mutatieZaakRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mutatieZaakRequest** | [**MutatieZaakRequest**](MutatieZaakRequest.md)|  | 

### Return type

[**MutatieZaak**](MutatieZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMutatieInplannenCreate

> MutatieZaakAangepastInDeBroker ZakenMutatieInplannenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenMutatieInplannenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenMutatieInplannenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakAangepastInDeBrokerRequest** | [**optional.Interface of MutatieZaakAangepastInDeBrokerRequest**](MutatieZaakAangepastInDeBrokerRequest.md)|  | 

### Return type

[**MutatieZaakAangepastInDeBroker**](MutatieZaakAangepastInDeBroker.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMutatieList

> PaginatedMutatieZaakList ZakenMutatieList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ZakenMutatieListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenMutatieListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gesloten** | **optional.Bool**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **status** | **optional.String**| mutatie &#x3D; Mutatie, ingediend &#x3D; Ingediend [ATSP], afgewezen &#x3D; Afgewezen [Risicobeheer], goedgekeurd &#x3D; Goedgekeurd [Risicobeheer], aangepast_broker &#x3D; Aangepast in de Broker [ATSP], gms_toegevoegd &#x3D; Toevoegingen doorgevoerd in het GMS [GMS], ingepland &#x3D; Live test ingepland [ATSP], aanpassingen_geverifieerd &#x3D; Live test goedgekeurd [AMS-Servicedesk], gms_verwijderd &#x3D; Criteria verwijderd in het GMS [AMS-Servicedesk], geannuleerd &#x3D; Geannuleerd | 
 **type_** | **optional.String**|  | 

### Return type

[**PaginatedMutatieZaakList**](PaginatedMutatieZaakList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMutatieRetrieve

> MutatieZaak ZakenMutatieRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**MutatieZaak**](MutatieZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMutatieTechnischGereedCreate

> MutatieZaakAangepastInDeBroker ZakenMutatieTechnischGereedCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenMutatieTechnischGereedCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenMutatieTechnischGereedCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakAangepastInDeBrokerRequest** | [**optional.Interface of MutatieZaakAangepastInDeBrokerRequest**](MutatieZaakAangepastInDeBrokerRequest.md)|  | 

### Return type

[**MutatieZaakAangepastInDeBroker**](MutatieZaakAangepastInDeBroker.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

