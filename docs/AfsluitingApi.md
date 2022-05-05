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

> AfsluitingZaak ZakenAfsluitingAfwijzenCreate(ctx, id).AfsluitingZaakAfwijzenRequest(afsluitingZaakAfwijzenRequest).Execute()



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
    afsluitingZaakAfwijzenRequest := *openapiclient.NewAfsluitingZaakAfwijzenRequest() // AfsluitingZaakAfwijzenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsluitingApi.ZakenAfsluitingAfwijzenCreate(context.Background(), id).AfsluitingZaakAfwijzenRequest(afsluitingZaakAfwijzenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsluitingApi.ZakenAfsluitingAfwijzenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenAfsluitingAfwijzenCreate`: AfsluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `AfsluitingApi.ZakenAfsluitingAfwijzenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenAfsluitingAfwijzenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakAfwijzenRequest** | [**AfsluitingZaakAfwijzenRequest**](AfsluitingZaakAfwijzenRequest.md) |  | 

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

> AfsluitingZaak ZakenAfsluitingAnnulerenCreate(ctx, id).AfsluitingZaakAnnulerenRequest(afsluitingZaakAnnulerenRequest).Execute()



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
    afsluitingZaakAnnulerenRequest := *openapiclient.NewAfsluitingZaakAnnulerenRequest("Description_example") // AfsluitingZaakAnnulerenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsluitingApi.ZakenAfsluitingAnnulerenCreate(context.Background(), id).AfsluitingZaakAnnulerenRequest(afsluitingZaakAnnulerenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsluitingApi.ZakenAfsluitingAnnulerenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenAfsluitingAnnulerenCreate`: AfsluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `AfsluitingApi.ZakenAfsluitingAnnulerenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenAfsluitingAnnulerenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakAnnulerenRequest** | [**AfsluitingZaakAnnulerenRequest**](AfsluitingZaakAnnulerenRequest.md) |  | 

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

> AfsluitingZaak ZakenAfsluitingBevestigenCreate(ctx, id).AfsluitingZaakBevestigenRequest(afsluitingZaakBevestigenRequest).Execute()



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
    afsluitingZaakBevestigenRequest := *openapiclient.NewAfsluitingZaakBevestigenRequest("Description_example") // AfsluitingZaakBevestigenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsluitingApi.ZakenAfsluitingBevestigenCreate(context.Background(), id).AfsluitingZaakBevestigenRequest(afsluitingZaakBevestigenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsluitingApi.ZakenAfsluitingBevestigenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenAfsluitingBevestigenCreate`: AfsluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `AfsluitingApi.ZakenAfsluitingBevestigenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenAfsluitingBevestigenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakBevestigenRequest** | [**AfsluitingZaakBevestigenRequest**](AfsluitingZaakBevestigenRequest.md) |  | 

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

> AfsluitingZaak ZakenAfsluitingGoedkeurenCreate(ctx, id).AfsluitingZaakGoedkeurenRequest(afsluitingZaakGoedkeurenRequest).Execute()



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
    afsluitingZaakGoedkeurenRequest := *openapiclient.NewAfsluitingZaakGoedkeurenRequest("Description_example") // AfsluitingZaakGoedkeurenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsluitingApi.ZakenAfsluitingGoedkeurenCreate(context.Background(), id).AfsluitingZaakGoedkeurenRequest(afsluitingZaakGoedkeurenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsluitingApi.ZakenAfsluitingGoedkeurenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenAfsluitingGoedkeurenCreate`: AfsluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `AfsluitingApi.ZakenAfsluitingGoedkeurenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenAfsluitingGoedkeurenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakGoedkeurenRequest** | [**AfsluitingZaakGoedkeurenRequest**](AfsluitingZaakGoedkeurenRequest.md) |  | 

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

> AfsluitingZaak ZakenAfsluitingIndienenCreate(ctx).AfsluitingZaakRequest(afsluitingZaakRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    afsluitingZaakRequest := *openapiclient.NewAfsluitingZaakRequest(int32(123), "Reden_example", time.Now()) // AfsluitingZaakRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsluitingApi.ZakenAfsluitingIndienenCreate(context.Background()).AfsluitingZaakRequest(afsluitingZaakRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsluitingApi.ZakenAfsluitingIndienenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenAfsluitingIndienenCreate`: AfsluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `AfsluitingApi.ZakenAfsluitingIndienenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZakenAfsluitingIndienenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afsluitingZaakRequest** | [**AfsluitingZaakRequest**](AfsluitingZaakRequest.md) |  | 

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

> PaginatedAfsluitingZaakList ZakenAfsluitingList(ctx).Gesloten(gesloten).Limit(limit).Offset(offset).Status(status).Execute()



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
    status := []string{"Status_example"} // []string | nieuw = Afsluiting, ingediend = Ingediend [ATSP], afgewezen = Afgewezen [Risicobeheer], goedgekeurd = Goedgekeurd [Risicobeheer], bevestigd = Bevestigd [ATSP], uitgevoerd = Uitgevoerd [AMS-Servicedesk], verwijderd = Verwijderd [GMS], geannuleerd = Geannuleerd (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsluitingApi.ZakenAfsluitingList(context.Background()).Gesloten(gesloten).Limit(limit).Offset(offset).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsluitingApi.ZakenAfsluitingList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenAfsluitingList`: PaginatedAfsluitingZaakList
    fmt.Fprintf(os.Stdout, "Response from `AfsluitingApi.ZakenAfsluitingList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZakenAfsluitingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gesloten** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **status** | **[]string** | nieuw &#x3D; Afsluiting, ingediend &#x3D; Ingediend [ATSP], afgewezen &#x3D; Afgewezen [Risicobeheer], goedgekeurd &#x3D; Goedgekeurd [Risicobeheer], bevestigd &#x3D; Bevestigd [ATSP], uitgevoerd &#x3D; Uitgevoerd [AMS-Servicedesk], verwijderd &#x3D; Verwijderd [GMS], geannuleerd &#x3D; Geannuleerd | 

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

> AfsluitingZaak ZakenAfsluitingRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.AfsluitingApi.ZakenAfsluitingRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsluitingApi.ZakenAfsluitingRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenAfsluitingRetrieve`: AfsluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `AfsluitingApi.ZakenAfsluitingRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenAfsluitingRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> AfsluitingZaak ZakenAfsluitingUitvoerenCreate(ctx, id).AfsluitingZaakUitvoerenRequest(afsluitingZaakUitvoerenRequest).Execute()



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
    afsluitingZaakUitvoerenRequest := *openapiclient.NewAfsluitingZaakUitvoerenRequest("Description_example") // AfsluitingZaakUitvoerenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsluitingApi.ZakenAfsluitingUitvoerenCreate(context.Background(), id).AfsluitingZaakUitvoerenRequest(afsluitingZaakUitvoerenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsluitingApi.ZakenAfsluitingUitvoerenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenAfsluitingUitvoerenCreate`: AfsluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `AfsluitingApi.ZakenAfsluitingUitvoerenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenAfsluitingUitvoerenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakUitvoerenRequest** | [**AfsluitingZaakUitvoerenRequest**](AfsluitingZaakUitvoerenRequest.md) |  | 

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

