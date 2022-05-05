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

> MigratieZaak ZakenMigratieAnnulerenCreate(ctx, id).MigratieZaakXRequest(migratieZaakXRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 
    migratieZaakXRequest := *openapiclient.NewMigratieZaakXRequest("Description_example") // MigratieZaakXRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigratieApi.ZakenMigratieAnnulerenCreate(context.Background(), id).MigratieZaakXRequest(migratieZaakXRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigratieApi.ZakenMigratieAnnulerenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMigratieAnnulerenCreate`: MigratieZaak
    fmt.Fprintf(os.Stdout, "Response from `MigratieApi.ZakenMigratieAnnulerenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMigratieAnnulerenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **migratieZaakXRequest** | [**MigratieZaakXRequest**](MigratieZaakXRequest.md) |  | 

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

> MigratieZaak ZakenMigratieGoedkeurenCreate(ctx, id).MigratieZaakXRequest(migratieZaakXRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 
    migratieZaakXRequest := *openapiclient.NewMigratieZaakXRequest("Description_example") // MigratieZaakXRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigratieApi.ZakenMigratieGoedkeurenCreate(context.Background(), id).MigratieZaakXRequest(migratieZaakXRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigratieApi.ZakenMigratieGoedkeurenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMigratieGoedkeurenCreate`: MigratieZaak
    fmt.Fprintf(os.Stdout, "Response from `MigratieApi.ZakenMigratieGoedkeurenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMigratieGoedkeurenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **migratieZaakXRequest** | [**MigratieZaakXRequest**](MigratieZaakXRequest.md) |  | 

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

> Aansluiting ZakenMigratieIndienenCreate(ctx, aansluitnummer).MigratieRequest(migratieRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aansluitnummer := "aansluitnummer_example" // string | 
    migratieRequest := *openapiclient.NewMigratieRequest("NieuweAansluitnummer_example", int32(123)) // MigratieRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigratieApi.ZakenMigratieIndienenCreate(context.Background(), aansluitnummer).MigratieRequest(migratieRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigratieApi.ZakenMigratieIndienenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMigratieIndienenCreate`: Aansluiting
    fmt.Fprintf(os.Stdout, "Response from `MigratieApi.ZakenMigratieIndienenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMigratieIndienenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **migratieRequest** | [**MigratieRequest**](MigratieRequest.md) |  | 

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

> MigratieZaakLiveTestInplannen ZakenMigratieInplannenCreate(ctx, id).MigratieZaakLiveTestInplannenRequest(migratieZaakLiveTestInplannenRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 
    migratieZaakLiveTestInplannenRequest := *openapiclient.NewMigratieZaakLiveTestInplannenRequest() // MigratieZaakLiveTestInplannenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigratieApi.ZakenMigratieInplannenCreate(context.Background(), id).MigratieZaakLiveTestInplannenRequest(migratieZaakLiveTestInplannenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigratieApi.ZakenMigratieInplannenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMigratieInplannenCreate`: MigratieZaakLiveTestInplannen
    fmt.Fprintf(os.Stdout, "Response from `MigratieApi.ZakenMigratieInplannenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMigratieInplannenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **migratieZaakLiveTestInplannenRequest** | [**MigratieZaakLiveTestInplannenRequest**](MigratieZaakLiveTestInplannenRequest.md) |  | 

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

> PaginatedMigratieZaakList ZakenMigratieList(ctx).Gesloten(gesloten).Limit(limit).Offset(offset).Status(status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gesloten := true // bool |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    status := []string{"Status_example"} // []string | nieuw = Nieuw, ingediend = Ingediend [ATSP], gms_doorgevoerd = Doorgevoerd [GMS], ingepland = Live test ingepland [ATSP], goedgekeurd = Goedgekeurd [AMS-Servicedesk], geannuleerd = Geannuleerd, gms_verwijderd = Latende aansluiting verwijderd [GMS] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigratieApi.ZakenMigratieList(context.Background()).Gesloten(gesloten).Limit(limit).Offset(offset).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigratieApi.ZakenMigratieList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMigratieList`: PaginatedMigratieZaakList
    fmt.Fprintf(os.Stdout, "Response from `MigratieApi.ZakenMigratieList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZakenMigratieListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gesloten** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **status** | **[]string** | nieuw &#x3D; Nieuw, ingediend &#x3D; Ingediend [ATSP], gms_doorgevoerd &#x3D; Doorgevoerd [GMS], ingepland &#x3D; Live test ingepland [ATSP], goedgekeurd &#x3D; Goedgekeurd [AMS-Servicedesk], geannuleerd &#x3D; Geannuleerd, gms_verwijderd &#x3D; Latende aansluiting verwijderd [GMS] | 

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

> MigratieZaak ZakenMigratieRetrieve(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigratieApi.ZakenMigratieRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigratieApi.ZakenMigratieRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMigratieRetrieve`: MigratieZaak
    fmt.Fprintf(os.Stdout, "Response from `MigratieApi.ZakenMigratieRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMigratieRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

