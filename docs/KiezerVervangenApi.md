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

> KiezerVervangenZaak ZakenKiezerVervangenAnnulerenCreate(ctx, id).KiezerVervangenZaakRequest(kiezerVervangenZaakRequest).Execute()



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
    kiezerVervangenZaakRequest := *openapiclient.NewKiezerVervangenZaakRequest("Description_example") // KiezerVervangenZaakRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KiezerVervangenApi.ZakenKiezerVervangenAnnulerenCreate(context.Background(), id).KiezerVervangenZaakRequest(kiezerVervangenZaakRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KiezerVervangenApi.ZakenKiezerVervangenAnnulerenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenKiezerVervangenAnnulerenCreate`: KiezerVervangenZaak
    fmt.Fprintf(os.Stdout, "Response from `KiezerVervangenApi.ZakenKiezerVervangenAnnulerenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenKiezerVervangenAnnulerenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kiezerVervangenZaakRequest** | [**KiezerVervangenZaakRequest**](KiezerVervangenZaakRequest.md) |  | 

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

> KiezerVervangenZaak ZakenKiezerVervangenGoedkeurenCreate(ctx, id).KiezerVervangenZaakRequest(kiezerVervangenZaakRequest).Execute()



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
    kiezerVervangenZaakRequest := *openapiclient.NewKiezerVervangenZaakRequest("Description_example") // KiezerVervangenZaakRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KiezerVervangenApi.ZakenKiezerVervangenGoedkeurenCreate(context.Background(), id).KiezerVervangenZaakRequest(kiezerVervangenZaakRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KiezerVervangenApi.ZakenKiezerVervangenGoedkeurenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenKiezerVervangenGoedkeurenCreate`: KiezerVervangenZaak
    fmt.Fprintf(os.Stdout, "Response from `KiezerVervangenApi.ZakenKiezerVervangenGoedkeurenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenKiezerVervangenGoedkeurenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kiezerVervangenZaakRequest** | [**KiezerVervangenZaakRequest**](KiezerVervangenZaakRequest.md) |  | 

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

> KiezerVervangenZaakIndienen ZakenKiezerVervangenIndienenCreate(ctx, aansluitnummer).KiezerVervangenZaakIndienenRequest(kiezerVervangenZaakIndienenRequest).Execute()



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
    kiezerVervangenZaakIndienenRequest := *openapiclient.NewKiezerVervangenZaakIndienenRequest(int32(123)) // KiezerVervangenZaakIndienenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KiezerVervangenApi.ZakenKiezerVervangenIndienenCreate(context.Background(), aansluitnummer).KiezerVervangenZaakIndienenRequest(kiezerVervangenZaakIndienenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KiezerVervangenApi.ZakenKiezerVervangenIndienenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenKiezerVervangenIndienenCreate`: KiezerVervangenZaakIndienen
    fmt.Fprintf(os.Stdout, "Response from `KiezerVervangenApi.ZakenKiezerVervangenIndienenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenKiezerVervangenIndienenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kiezerVervangenZaakIndienenRequest** | [**KiezerVervangenZaakIndienenRequest**](KiezerVervangenZaakIndienenRequest.md) |  | 

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

> KiezerVervangenZaakInplannen ZakenKiezerVervangenInplannenCreate(ctx, id).KiezerVervangenZaakInplannenRequest(kiezerVervangenZaakInplannenRequest).Execute()



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
    kiezerVervangenZaakInplannenRequest := *openapiclient.NewKiezerVervangenZaakInplannenRequest() // KiezerVervangenZaakInplannenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KiezerVervangenApi.ZakenKiezerVervangenInplannenCreate(context.Background(), id).KiezerVervangenZaakInplannenRequest(kiezerVervangenZaakInplannenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KiezerVervangenApi.ZakenKiezerVervangenInplannenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenKiezerVervangenInplannenCreate`: KiezerVervangenZaakInplannen
    fmt.Fprintf(os.Stdout, "Response from `KiezerVervangenApi.ZakenKiezerVervangenInplannenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenKiezerVervangenInplannenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kiezerVervangenZaakInplannenRequest** | [**KiezerVervangenZaakInplannenRequest**](KiezerVervangenZaakInplannenRequest.md) |  | 

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

> PaginatedKiezerVervangenZaakList ZakenKiezerVervangenList(ctx).Gesloten(gesloten).Limit(limit).Offset(offset).Status(status).Execute()



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
    status := []string{"Status_example"} // []string | nieuw = Nieuw, ingediend = Ingediend [ATSP], goedgekeurd = Goedgekeurd [AMS-Servicedesk], ingepland = Live test ingepland [ATSP], doorgevoerd = Doorgevoerd [GMS], geannuleerd = Geannuleerd (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KiezerVervangenApi.ZakenKiezerVervangenList(context.Background()).Gesloten(gesloten).Limit(limit).Offset(offset).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KiezerVervangenApi.ZakenKiezerVervangenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenKiezerVervangenList`: PaginatedKiezerVervangenZaakList
    fmt.Fprintf(os.Stdout, "Response from `KiezerVervangenApi.ZakenKiezerVervangenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZakenKiezerVervangenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gesloten** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **status** | **[]string** | nieuw &#x3D; Nieuw, ingediend &#x3D; Ingediend [ATSP], goedgekeurd &#x3D; Goedgekeurd [AMS-Servicedesk], ingepland &#x3D; Live test ingepland [ATSP], doorgevoerd &#x3D; Doorgevoerd [GMS], geannuleerd &#x3D; Geannuleerd | 

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

> KiezerVervangenZaak ZakenKiezerVervangenRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.KiezerVervangenApi.ZakenKiezerVervangenRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KiezerVervangenApi.ZakenKiezerVervangenRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenKiezerVervangenRetrieve`: KiezerVervangenZaak
    fmt.Fprintf(os.Stdout, "Response from `KiezerVervangenApi.ZakenKiezerVervangenRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenKiezerVervangenRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

