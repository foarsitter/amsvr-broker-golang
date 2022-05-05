# \BrandweerApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrandweerGemeentesList**](BrandweerApi.md#BrandweerGemeentesList) | **Get** /api/brandweer/gemeentes | 
[**BrandweerGemeentesRetrieve**](BrandweerApi.md#BrandweerGemeentesRetrieve) | **Get** /api/brandweer/gemeentes/{id} | 
[**BrandweerMeldkamersCriteriaList**](BrandweerApi.md#BrandweerMeldkamersCriteriaList) | **Get** /api/brandweer/meldkamers/{id}/criteria | 
[**BrandweerMeldkamersList**](BrandweerApi.md#BrandweerMeldkamersList) | **Get** /api/brandweer/meldkamers | 
[**BrandweerMeldkamersRetrieve**](BrandweerApi.md#BrandweerMeldkamersRetrieve) | **Get** /api/brandweer/meldkamers/{id} | 
[**BrandweerRegiosList**](BrandweerApi.md#BrandweerRegiosList) | **Get** /api/brandweer/regios | 
[**BrandweerRegiosRetrieve**](BrandweerApi.md#BrandweerRegiosRetrieve) | **Get** /api/brandweer/regios/{id} | 



## BrandweerGemeentesList

> PaginatedGemeenteList BrandweerGemeentesList(ctx).Afkorting(afkorting).Limit(limit).Naam(naam).Offset(offset).Plaatsen(plaatsen).ProvincieAfkorting(provincieAfkorting).Regio(regio).Execute()



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
    afkorting := "afkorting_example" // string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    naam := "naam_example" // string |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    plaatsen := "plaatsen_example" // string |  (optional)
    provincieAfkorting := "provincieAfkorting_example" // string |  (optional)
    regio := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandweerApi.BrandweerGemeentesList(context.Background()).Afkorting(afkorting).Limit(limit).Naam(naam).Offset(offset).Plaatsen(plaatsen).ProvincieAfkorting(provincieAfkorting).Regio(regio).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandweerApi.BrandweerGemeentesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrandweerGemeentesList`: PaginatedGemeenteList
    fmt.Fprintf(os.Stdout, "Response from `BrandweerApi.BrandweerGemeentesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrandweerGemeentesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afkorting** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **naam** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **plaatsen** | **string** |  | 
 **provincieAfkorting** | **string** |  | 
 **regio** | **int32** |  | 

### Return type

[**PaginatedGemeenteList**](PaginatedGemeenteList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerGemeentesRetrieve

> Gemeente BrandweerGemeentesRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.BrandweerApi.BrandweerGemeentesRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandweerApi.BrandweerGemeentesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrandweerGemeentesRetrieve`: Gemeente
    fmt.Fprintf(os.Stdout, "Response from `BrandweerApi.BrandweerGemeentesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrandweerGemeentesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Gemeente**](Gemeente.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerMeldkamersCriteriaList

> PaginatedCriteriumList BrandweerMeldkamersCriteriaList(ctx, id).Limit(limit).Offset(offset).Execute()



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandweerApi.BrandweerMeldkamersCriteriaList(context.Background(), id).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandweerApi.BrandweerMeldkamersCriteriaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrandweerMeldkamersCriteriaList`: PaginatedCriteriumList
    fmt.Fprintf(os.Stdout, "Response from `BrandweerApi.BrandweerMeldkamersCriteriaList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrandweerMeldkamersCriteriaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedCriteriumList**](PaginatedCriteriumList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerMeldkamersList

> PaginatedMeldkamerList BrandweerMeldkamersList(ctx).Limit(limit).Offset(offset).Execute()



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandweerApi.BrandweerMeldkamersList(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandweerApi.BrandweerMeldkamersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrandweerMeldkamersList`: PaginatedMeldkamerList
    fmt.Fprintf(os.Stdout, "Response from `BrandweerApi.BrandweerMeldkamersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrandweerMeldkamersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedMeldkamerList**](PaginatedMeldkamerList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerMeldkamersRetrieve

> Meldkamer BrandweerMeldkamersRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.BrandweerApi.BrandweerMeldkamersRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandweerApi.BrandweerMeldkamersRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrandweerMeldkamersRetrieve`: Meldkamer
    fmt.Fprintf(os.Stdout, "Response from `BrandweerApi.BrandweerMeldkamersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrandweerMeldkamersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Meldkamer**](Meldkamer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerRegiosList

> PaginatedRegioList BrandweerRegiosList(ctx).Limit(limit).Meldkamer(meldkamer).Naam(naam).Offset(offset).Execute()



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    meldkamer := int32(56) // int32 |  (optional)
    naam := "naam_example" // string |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandweerApi.BrandweerRegiosList(context.Background()).Limit(limit).Meldkamer(meldkamer).Naam(naam).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandweerApi.BrandweerRegiosList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrandweerRegiosList`: PaginatedRegioList
    fmt.Fprintf(os.Stdout, "Response from `BrandweerApi.BrandweerRegiosList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrandweerRegiosListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **meldkamer** | **int32** |  | 
 **naam** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedRegioList**](PaginatedRegioList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandweerRegiosRetrieve

> Regio BrandweerRegiosRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.BrandweerApi.BrandweerRegiosRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandweerApi.BrandweerRegiosRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrandweerRegiosRetrieve`: Regio
    fmt.Fprintf(os.Stdout, "Response from `BrandweerApi.BrandweerRegiosRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrandweerRegiosRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Regio**](Regio.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

