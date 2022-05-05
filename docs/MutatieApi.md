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

> MutatieZaak ZakenMutatieAangepastBrokerCreate(ctx, id).MutatieZaakGoedkeurenRequest(mutatieZaakGoedkeurenRequest).Execute()



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
    mutatieZaakGoedkeurenRequest := *openapiclient.NewMutatieZaakGoedkeurenRequest() // MutatieZaakGoedkeurenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutatieApi.ZakenMutatieAangepastBrokerCreate(context.Background(), id).MutatieZaakGoedkeurenRequest(mutatieZaakGoedkeurenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutatieApi.ZakenMutatieAangepastBrokerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMutatieAangepastBrokerCreate`: MutatieZaak
    fmt.Fprintf(os.Stdout, "Response from `MutatieApi.ZakenMutatieAangepastBrokerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMutatieAangepastBrokerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakGoedkeurenRequest** | [**MutatieZaakGoedkeurenRequest**](MutatieZaakGoedkeurenRequest.md) |  | 

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

> MutatieZaak ZakenMutatieAccorderenCreate(ctx, id).MutatieZaakGoedkeurenRequest(mutatieZaakGoedkeurenRequest).Execute()



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
    mutatieZaakGoedkeurenRequest := *openapiclient.NewMutatieZaakGoedkeurenRequest() // MutatieZaakGoedkeurenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutatieApi.ZakenMutatieAccorderenCreate(context.Background(), id).MutatieZaakGoedkeurenRequest(mutatieZaakGoedkeurenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutatieApi.ZakenMutatieAccorderenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMutatieAccorderenCreate`: MutatieZaak
    fmt.Fprintf(os.Stdout, "Response from `MutatieApi.ZakenMutatieAccorderenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMutatieAccorderenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakGoedkeurenRequest** | [**MutatieZaakGoedkeurenRequest**](MutatieZaakGoedkeurenRequest.md) |  | 

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

> MutatieZaak ZakenMutatieAfwijzenCreate(ctx, id).MutatieZaakGoedkeurenRequest(mutatieZaakGoedkeurenRequest).Execute()



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
    mutatieZaakGoedkeurenRequest := *openapiclient.NewMutatieZaakGoedkeurenRequest() // MutatieZaakGoedkeurenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutatieApi.ZakenMutatieAfwijzenCreate(context.Background(), id).MutatieZaakGoedkeurenRequest(mutatieZaakGoedkeurenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutatieApi.ZakenMutatieAfwijzenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMutatieAfwijzenCreate`: MutatieZaak
    fmt.Fprintf(os.Stdout, "Response from `MutatieApi.ZakenMutatieAfwijzenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMutatieAfwijzenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakGoedkeurenRequest** | [**MutatieZaakGoedkeurenRequest**](MutatieZaakGoedkeurenRequest.md) |  | 

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

> MutatieZaak ZakenMutatieAnnulerenCreate(ctx, id).MutatieZaakGoedkeurenRequest(mutatieZaakGoedkeurenRequest).Execute()



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
    mutatieZaakGoedkeurenRequest := *openapiclient.NewMutatieZaakGoedkeurenRequest() // MutatieZaakGoedkeurenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutatieApi.ZakenMutatieAnnulerenCreate(context.Background(), id).MutatieZaakGoedkeurenRequest(mutatieZaakGoedkeurenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutatieApi.ZakenMutatieAnnulerenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMutatieAnnulerenCreate`: MutatieZaak
    fmt.Fprintf(os.Stdout, "Response from `MutatieApi.ZakenMutatieAnnulerenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMutatieAnnulerenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakGoedkeurenRequest** | [**MutatieZaakGoedkeurenRequest**](MutatieZaakGoedkeurenRequest.md) |  | 

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

> MutatieZaak ZakenMutatieGoedkeurenCreate(ctx, id).MutatieZaakGoedkeurenRequest(mutatieZaakGoedkeurenRequest).Execute()



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
    mutatieZaakGoedkeurenRequest := *openapiclient.NewMutatieZaakGoedkeurenRequest() // MutatieZaakGoedkeurenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutatieApi.ZakenMutatieGoedkeurenCreate(context.Background(), id).MutatieZaakGoedkeurenRequest(mutatieZaakGoedkeurenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutatieApi.ZakenMutatieGoedkeurenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMutatieGoedkeurenCreate`: MutatieZaak
    fmt.Fprintf(os.Stdout, "Response from `MutatieApi.ZakenMutatieGoedkeurenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMutatieGoedkeurenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakGoedkeurenRequest** | [**MutatieZaakGoedkeurenRequest**](MutatieZaakGoedkeurenRequest.md) |  | 

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

> MutatieZaak ZakenMutatieIndienenCreate(ctx).MutatieZaakRequest(mutatieZaakRequest).Execute()



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
    mutatieZaakRequest := *openapiclient.NewMutatieZaakRequest(int32(123), "OmschrijvingVanWijziging_example") // MutatieZaakRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutatieApi.ZakenMutatieIndienenCreate(context.Background()).MutatieZaakRequest(mutatieZaakRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutatieApi.ZakenMutatieIndienenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMutatieIndienenCreate`: MutatieZaak
    fmt.Fprintf(os.Stdout, "Response from `MutatieApi.ZakenMutatieIndienenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZakenMutatieIndienenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mutatieZaakRequest** | [**MutatieZaakRequest**](MutatieZaakRequest.md) |  | 

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

> MutatieZaakAangepastInDeBroker ZakenMutatieInplannenCreate(ctx, id).MutatieZaakAangepastInDeBrokerRequest(mutatieZaakAangepastInDeBrokerRequest).Execute()



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
    mutatieZaakAangepastInDeBrokerRequest := *openapiclient.NewMutatieZaakAangepastInDeBrokerRequest() // MutatieZaakAangepastInDeBrokerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutatieApi.ZakenMutatieInplannenCreate(context.Background(), id).MutatieZaakAangepastInDeBrokerRequest(mutatieZaakAangepastInDeBrokerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutatieApi.ZakenMutatieInplannenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMutatieInplannenCreate`: MutatieZaakAangepastInDeBroker
    fmt.Fprintf(os.Stdout, "Response from `MutatieApi.ZakenMutatieInplannenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMutatieInplannenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakAangepastInDeBrokerRequest** | [**MutatieZaakAangepastInDeBrokerRequest**](MutatieZaakAangepastInDeBrokerRequest.md) |  | 

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

> PaginatedMutatieZaakList ZakenMutatieList(ctx).Gesloten(gesloten).Limit(limit).Offset(offset).Status(status).Type_(type_).Execute()



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
    status := []string{"Status_example"} // []string | mutatie = Mutatie, ingediend = Ingediend [ATSP], afgewezen = Afgewezen [Risicobeheer], goedgekeurd = Goedgekeurd [Risicobeheer], aangepast_broker = Aangepast in de Broker [ATSP], gms_toegevoegd = Toevoegingen doorgevoerd in het GMS [GMS], ingepland = Live test ingepland [ATSP], aanpassingen_geverifieerd = Aanpassingen goedgekeurd [AMS-Servicedesk], gms_verwijderd = Criteria verwijderd in het GMS [AMS-Servicedesk], geannuleerd = Geannuleerd (optional)
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutatieApi.ZakenMutatieList(context.Background()).Gesloten(gesloten).Limit(limit).Offset(offset).Status(status).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutatieApi.ZakenMutatieList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMutatieList`: PaginatedMutatieZaakList
    fmt.Fprintf(os.Stdout, "Response from `MutatieApi.ZakenMutatieList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZakenMutatieListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gesloten** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **status** | **[]string** | mutatie &#x3D; Mutatie, ingediend &#x3D; Ingediend [ATSP], afgewezen &#x3D; Afgewezen [Risicobeheer], goedgekeurd &#x3D; Goedgekeurd [Risicobeheer], aangepast_broker &#x3D; Aangepast in de Broker [ATSP], gms_toegevoegd &#x3D; Toevoegingen doorgevoerd in het GMS [GMS], ingepland &#x3D; Live test ingepland [ATSP], aanpassingen_geverifieerd &#x3D; Aanpassingen goedgekeurd [AMS-Servicedesk], gms_verwijderd &#x3D; Criteria verwijderd in het GMS [AMS-Servicedesk], geannuleerd &#x3D; Geannuleerd | 
 **type_** | **string** |  | 

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

> MutatieZaak ZakenMutatieRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.MutatieApi.ZakenMutatieRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutatieApi.ZakenMutatieRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMutatieRetrieve`: MutatieZaak
    fmt.Fprintf(os.Stdout, "Response from `MutatieApi.ZakenMutatieRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMutatieRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> MutatieZaakAangepastInDeBroker ZakenMutatieTechnischGereedCreate(ctx, id).MutatieZaakAangepastInDeBrokerRequest(mutatieZaakAangepastInDeBrokerRequest).Execute()



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
    mutatieZaakAangepastInDeBrokerRequest := *openapiclient.NewMutatieZaakAangepastInDeBrokerRequest() // MutatieZaakAangepastInDeBrokerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutatieApi.ZakenMutatieTechnischGereedCreate(context.Background(), id).MutatieZaakAangepastInDeBrokerRequest(mutatieZaakAangepastInDeBrokerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutatieApi.ZakenMutatieTechnischGereedCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenMutatieTechnischGereedCreate`: MutatieZaakAangepastInDeBroker
    fmt.Fprintf(os.Stdout, "Response from `MutatieApi.ZakenMutatieTechnischGereedCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenMutatieTechnischGereedCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutatieZaakAangepastInDeBrokerRequest** | [**MutatieZaakAangepastInDeBrokerRequest**](MutatieZaakAangepastInDeBrokerRequest.md) |  | 

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

