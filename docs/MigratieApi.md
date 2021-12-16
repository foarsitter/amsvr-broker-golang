# \MigratieApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZakenMigratieAnnulerenCreate**](MigratieApi.md#ZakenMigratieAnnulerenCreate) | **Post** /api/zaken/migratie/{id}/annuleren | 
[**ZakenMigratieGoedkeurenCreate**](MigratieApi.md#ZakenMigratieGoedkeurenCreate) | **Post** /api/zaken/migratie/{id}/goedkeuren | 
[**ZakenMigratieIndienenCreate**](MigratieApi.md#ZakenMigratieIndienenCreate) | **Post** /api/zaken/migratie/{aansluitnummer}/indienen | 
[**ZakenMigratieInplannenCreate**](MigratieApi.md#ZakenMigratieInplannenCreate) | **Post** /api/zaken/migratie/{id}/inplannen | 
[**ZakenMigratieList**](MigratieApi.md#ZakenMigratieList) | **Get** /api/zaken/migratie/ | 
[**ZakenMigratieRetrieve**](MigratieApi.md#ZakenMigratieRetrieve) | **Get** /api/zaken/migratie/{id} | 



## ZakenMigratieAnnulerenCreate

> MigratieZaak ZakenMigratieAnnulerenCreate(ctx, id, migratieZaakXRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**migratieZaakXRequest** | [**MigratieZaakXRequest**](MigratieZaakXRequest.md)|  | 

### Return type

[**MigratieZaak**](MigratieZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMigratieGoedkeurenCreate

> MigratieZaak ZakenMigratieGoedkeurenCreate(ctx, id, migratieZaakXRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**migratieZaakXRequest** | [**MigratieZaakXRequest**](MigratieZaakXRequest.md)|  | 

### Return type

[**MigratieZaak**](MigratieZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMigratieIndienenCreate

> Aansluiting ZakenMigratieIndienenCreate(ctx, aansluitnummer, migratieRequest)



Het is mogelijk om een bestaande aansluiting over te nemen van een andere ATSP. Ter verificatie vragen wij om straat + huisnummer (zonder toevoegingen)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string**|  | 
**migratieRequest** | [**MigratieRequest**](MigratieRequest.md)|  | 

### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMigratieInplannenCreate

> MigratieZaakLiveTestInplannen ZakenMigratieInplannenCreate(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***ZakenMigratieInplannenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenMigratieInplannenCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **migratieZaakLiveTestInplannenRequest** | [**optional.Interface of MigratieZaakLiveTestInplannenRequest**](MigratieZaakLiveTestInplannenRequest.md)|  | 

### Return type

[**MigratieZaakLiveTestInplannen**](MigratieZaakLiveTestInplannen.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMigratieList

> PaginatedMigratieZaakList ZakenMigratieList(ctx, optional)



Een migratie zaak wordt automatisch gestart bij het uitvoeren van een migratie

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ZakenMigratieListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZakenMigratieListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gesloten** | **optional.Bool**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **status** | **optional.String**| nieuw &#x3D; Nieuw, ingediend &#x3D; Ingediend [ATSP], ingepland &#x3D; Live test ingepland [ATSP], gms_doorgevoerd &#x3D; Doorgevoerd [GMS], goedgekeurd &#x3D; Goedgekeurd [AMS-Servicedesk], geannuleerd &#x3D; Geannuleerd, gms_verwijderd &#x3D; Latende aansluiting verwijderd [GMS] | 

### Return type

[**PaginatedMigratieZaakList**](PaginatedMigratieZaakList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenMigratieRetrieve

> MigratieZaak ZakenMigratieRetrieve(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**MigratieZaak**](MigratieZaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

