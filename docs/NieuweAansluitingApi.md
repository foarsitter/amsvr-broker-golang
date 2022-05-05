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

> NieuweAansluitingZaak ZakenNieuweAansluitingAfkeurenCreate(ctx, id).AfsluitingZaakTransitieRequest(afsluitingZaakTransitieRequest).Execute()



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
    afsluitingZaakTransitieRequest := *openapiclient.NewAfsluitingZaakTransitieRequest("Description_example") // AfsluitingZaakTransitieRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NieuweAansluitingApi.ZakenNieuweAansluitingAfkeurenCreate(context.Background(), id).AfsluitingZaakTransitieRequest(afsluitingZaakTransitieRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NieuweAansluitingApi.ZakenNieuweAansluitingAfkeurenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenNieuweAansluitingAfkeurenCreate`: NieuweAansluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `NieuweAansluitingApi.ZakenNieuweAansluitingAfkeurenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenNieuweAansluitingAfkeurenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakTransitieRequest** | [**AfsluitingZaakTransitieRequest**](AfsluitingZaakTransitieRequest.md) |  | 

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

> NieuweAansluitingZaak ZakenNieuweAansluitingAnnulerenCreate(ctx, id).AfsluitingZaakTransitieRequest(afsluitingZaakTransitieRequest).Execute()



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
    afsluitingZaakTransitieRequest := *openapiclient.NewAfsluitingZaakTransitieRequest("Description_example") // AfsluitingZaakTransitieRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NieuweAansluitingApi.ZakenNieuweAansluitingAnnulerenCreate(context.Background(), id).AfsluitingZaakTransitieRequest(afsluitingZaakTransitieRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NieuweAansluitingApi.ZakenNieuweAansluitingAnnulerenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenNieuweAansluitingAnnulerenCreate`: NieuweAansluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `NieuweAansluitingApi.ZakenNieuweAansluitingAnnulerenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenNieuweAansluitingAnnulerenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakTransitieRequest** | [**AfsluitingZaakTransitieRequest**](AfsluitingZaakTransitieRequest.md) |  | 

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

> NieuweAansluitingZaak ZakenNieuweAansluitingGoedkeurenCreate(ctx, id).NieuweAfsluitingZaakGoedkeurenRequest(nieuweAfsluitingZaakGoedkeurenRequest).Execute()



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
    nieuweAfsluitingZaakGoedkeurenRequest := *openapiclient.NewNieuweAfsluitingZaakGoedkeurenRequest("Description_example", openapiclient.GebruiksfunctieEnum("Gezondheidszorgfunctie met bedgebied")) // NieuweAfsluitingZaakGoedkeurenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NieuweAansluitingApi.ZakenNieuweAansluitingGoedkeurenCreate(context.Background(), id).NieuweAfsluitingZaakGoedkeurenRequest(nieuweAfsluitingZaakGoedkeurenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NieuweAansluitingApi.ZakenNieuweAansluitingGoedkeurenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenNieuweAansluitingGoedkeurenCreate`: NieuweAansluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `NieuweAansluitingApi.ZakenNieuweAansluitingGoedkeurenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenNieuweAansluitingGoedkeurenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nieuweAfsluitingZaakGoedkeurenRequest** | [**NieuweAfsluitingZaakGoedkeurenRequest**](NieuweAfsluitingZaakGoedkeurenRequest.md) |  | 

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

> NieuweAansluitingZaak ZakenNieuweAansluitingIndienenCreate(ctx, id).AfsluitingZaakTransitieRequest(afsluitingZaakTransitieRequest).Execute()



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
    afsluitingZaakTransitieRequest := *openapiclient.NewAfsluitingZaakTransitieRequest("Description_example") // AfsluitingZaakTransitieRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NieuweAansluitingApi.ZakenNieuweAansluitingIndienenCreate(context.Background(), id).AfsluitingZaakTransitieRequest(afsluitingZaakTransitieRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NieuweAansluitingApi.ZakenNieuweAansluitingIndienenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenNieuweAansluitingIndienenCreate`: NieuweAansluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `NieuweAansluitingApi.ZakenNieuweAansluitingIndienenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenNieuweAansluitingIndienenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakTransitieRequest** | [**AfsluitingZaakTransitieRequest**](AfsluitingZaakTransitieRequest.md) |  | 

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

> NieuweAansluitingZaakLiveTest ZakenNieuweAansluitingInplannenCreate(ctx, id).NieuweAansluitingZaakLiveTestRequest(nieuweAansluitingZaakLiveTestRequest).Execute()



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
    nieuweAansluitingZaakLiveTestRequest := *openapiclient.NewNieuweAansluitingZaakLiveTestRequest() // NieuweAansluitingZaakLiveTestRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NieuweAansluitingApi.ZakenNieuweAansluitingInplannenCreate(context.Background(), id).NieuweAansluitingZaakLiveTestRequest(nieuweAansluitingZaakLiveTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NieuweAansluitingApi.ZakenNieuweAansluitingInplannenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenNieuweAansluitingInplannenCreate`: NieuweAansluitingZaakLiveTest
    fmt.Fprintf(os.Stdout, "Response from `NieuweAansluitingApi.ZakenNieuweAansluitingInplannenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenNieuweAansluitingInplannenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nieuweAansluitingZaakLiveTestRequest** | [**NieuweAansluitingZaakLiveTestRequest**](NieuweAansluitingZaakLiveTestRequest.md) |  | 

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

> PaginatedNieuweAansluitingZaakList ZakenNieuweAansluitingList(ctx).Gesloten(gesloten).Limit(limit).Offset(offset).Status(status).Execute()



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
    status := []string{"Status_example"} // []string | nieuw = Nieuw, ingediend = Ingediend ter aanvraag [ATSP], afgekeurd = Aanvraag afgekeurd [Risicobeheer], realisatie = Aanvraag goedgekeurd [Risicobeheer], gms_doorgevoerd = Doorgevoerd [GMS], technisch_gereed = Live test ingepland [ATSP], goedgekeurd = Goedgekeurd [AMS-Servicedesk], opgeleverd = Actief gezet [Risicobeheer], geannuleerd = Geannuleerd (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NieuweAansluitingApi.ZakenNieuweAansluitingList(context.Background()).Gesloten(gesloten).Limit(limit).Offset(offset).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NieuweAansluitingApi.ZakenNieuweAansluitingList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenNieuweAansluitingList`: PaginatedNieuweAansluitingZaakList
    fmt.Fprintf(os.Stdout, "Response from `NieuweAansluitingApi.ZakenNieuweAansluitingList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZakenNieuweAansluitingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gesloten** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **status** | **[]string** | nieuw &#x3D; Nieuw, ingediend &#x3D; Ingediend ter aanvraag [ATSP], afgekeurd &#x3D; Aanvraag afgekeurd [Risicobeheer], realisatie &#x3D; Aanvraag goedgekeurd [Risicobeheer], gms_doorgevoerd &#x3D; Doorgevoerd [GMS], technisch_gereed &#x3D; Live test ingepland [ATSP], goedgekeurd &#x3D; Goedgekeurd [AMS-Servicedesk], opgeleverd &#x3D; Actief gezet [Risicobeheer], geannuleerd &#x3D; Geannuleerd | 

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

> NieuweAansluitingZaak ZakenNieuweAansluitingOpleverenCreate(ctx, id).AfsluitingZaakTransitieRequest(afsluitingZaakTransitieRequest).Execute()



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
    afsluitingZaakTransitieRequest := *openapiclient.NewAfsluitingZaakTransitieRequest("Description_example") // AfsluitingZaakTransitieRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NieuweAansluitingApi.ZakenNieuweAansluitingOpleverenCreate(context.Background(), id).AfsluitingZaakTransitieRequest(afsluitingZaakTransitieRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NieuweAansluitingApi.ZakenNieuweAansluitingOpleverenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenNieuweAansluitingOpleverenCreate`: NieuweAansluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `NieuweAansluitingApi.ZakenNieuweAansluitingOpleverenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenNieuweAansluitingOpleverenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakTransitieRequest** | [**AfsluitingZaakTransitieRequest**](AfsluitingZaakTransitieRequest.md) |  | 

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

> NieuweAansluitingZaakLiveTest ZakenNieuweAansluitingRealisatieAfrondenCreate(ctx, id).NieuweAansluitingZaakLiveTestRequest(nieuweAansluitingZaakLiveTestRequest).Execute()



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
    nieuweAansluitingZaakLiveTestRequest := *openapiclient.NewNieuweAansluitingZaakLiveTestRequest() // NieuweAansluitingZaakLiveTestRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NieuweAansluitingApi.ZakenNieuweAansluitingRealisatieAfrondenCreate(context.Background(), id).NieuweAansluitingZaakLiveTestRequest(nieuweAansluitingZaakLiveTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NieuweAansluitingApi.ZakenNieuweAansluitingRealisatieAfrondenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenNieuweAansluitingRealisatieAfrondenCreate`: NieuweAansluitingZaakLiveTest
    fmt.Fprintf(os.Stdout, "Response from `NieuweAansluitingApi.ZakenNieuweAansluitingRealisatieAfrondenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenNieuweAansluitingRealisatieAfrondenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nieuweAansluitingZaakLiveTestRequest** | [**NieuweAansluitingZaakLiveTestRequest**](NieuweAansluitingZaakLiveTestRequest.md) |  | 

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

> NieuweAansluitingZaak ZakenNieuweAansluitingRealisatieGoedkeurenCreate(ctx, id).AfsluitingZaakTransitieRequest(afsluitingZaakTransitieRequest).Execute()



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
    afsluitingZaakTransitieRequest := *openapiclient.NewAfsluitingZaakTransitieRequest("Description_example") // AfsluitingZaakTransitieRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NieuweAansluitingApi.ZakenNieuweAansluitingRealisatieGoedkeurenCreate(context.Background(), id).AfsluitingZaakTransitieRequest(afsluitingZaakTransitieRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NieuweAansluitingApi.ZakenNieuweAansluitingRealisatieGoedkeurenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenNieuweAansluitingRealisatieGoedkeurenCreate`: NieuweAansluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `NieuweAansluitingApi.ZakenNieuweAansluitingRealisatieGoedkeurenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenNieuweAansluitingRealisatieGoedkeurenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afsluitingZaakTransitieRequest** | [**AfsluitingZaakTransitieRequest**](AfsluitingZaakTransitieRequest.md) |  | 

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

> NieuweAansluitingZaak ZakenNieuweAansluitingRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.NieuweAansluitingApi.ZakenNieuweAansluitingRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NieuweAansluitingApi.ZakenNieuweAansluitingRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenNieuweAansluitingRetrieve`: NieuweAansluitingZaak
    fmt.Fprintf(os.Stdout, "Response from `NieuweAansluitingApi.ZakenNieuweAansluitingRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenNieuweAansluitingRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

