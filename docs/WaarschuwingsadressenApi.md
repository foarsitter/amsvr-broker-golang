# \WaarschuwingsadressenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenSectorenWaarschuwingsadressenCreate**](WaarschuwingsadressenApi.md#AansluitingenSectorenWaarschuwingsadressenCreate) | **Post** /api/aansluitingen/sectoren/{id}/waarschuwingsadressen | 
[**AansluitingenSectorenWaarschuwingsadressenList**](WaarschuwingsadressenApi.md#AansluitingenSectorenWaarschuwingsadressenList) | **Get** /api/aansluitingen/sectoren/{id}/waarschuwingsadressen | 
[**AansluitingenWaarschuwingsadressenDestroy**](WaarschuwingsadressenApi.md#AansluitingenWaarschuwingsadressenDestroy) | **Delete** /api/aansluitingen/waarschuwingsadressen/{id} | 
[**AansluitingenWaarschuwingsadressenPartialUpdate**](WaarschuwingsadressenApi.md#AansluitingenWaarschuwingsadressenPartialUpdate) | **Patch** /api/aansluitingen/waarschuwingsadressen/{id} | 
[**AansluitingenWaarschuwingsadressenRetrieve**](WaarschuwingsadressenApi.md#AansluitingenWaarschuwingsadressenRetrieve) | **Get** /api/aansluitingen/waarschuwingsadressen/{id} | 
[**AansluitingenWaarschuwingsadressenUpdate**](WaarschuwingsadressenApi.md#AansluitingenWaarschuwingsadressenUpdate) | **Put** /api/aansluitingen/waarschuwingsadressen/{id} | 



## AansluitingenSectorenWaarschuwingsadressenCreate

> Waarschuwingsadres AansluitingenSectorenWaarschuwingsadressenCreate(ctx, id).WaarschuwingsadresRequest(waarschuwingsadresRequest).Execute()





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
    waarschuwingsadresRequest := *openapiclient.NewWaarschuwingsadresRequest(int32(123)) // WaarschuwingsadresRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WaarschuwingsadressenApi.AansluitingenSectorenWaarschuwingsadressenCreate(context.Background(), id).WaarschuwingsadresRequest(waarschuwingsadresRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WaarschuwingsadressenApi.AansluitingenSectorenWaarschuwingsadressenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenWaarschuwingsadressenCreate`: Waarschuwingsadres
    fmt.Fprintf(os.Stdout, "Response from `WaarschuwingsadressenApi.AansluitingenSectorenWaarschuwingsadressenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenWaarschuwingsadressenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **waarschuwingsadresRequest** | [**WaarschuwingsadresRequest**](WaarschuwingsadresRequest.md) |  | 

### Return type

[**Waarschuwingsadres**](Waarschuwingsadres.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenWaarschuwingsadressenList

> PaginatedWaarschuwingsadresList AansluitingenSectorenWaarschuwingsadressenList(ctx, id).Limit(limit).Offset(offset).Execute()





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
    id := "id_example" // string | A unique integer value identifying this Sector
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WaarschuwingsadressenApi.AansluitingenSectorenWaarschuwingsadressenList(context.Background(), id).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WaarschuwingsadressenApi.AansluitingenSectorenWaarschuwingsadressenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenWaarschuwingsadressenList`: PaginatedWaarschuwingsadresList
    fmt.Fprintf(os.Stdout, "Response from `WaarschuwingsadressenApi.AansluitingenSectorenWaarschuwingsadressenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A unique integer value identifying this Sector | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenWaarschuwingsadressenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedWaarschuwingsadresList**](PaginatedWaarschuwingsadresList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenWaarschuwingsadressenDestroy

> AansluitingenWaarschuwingsadressenDestroy(ctx, id).Execute()



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
    resp, r, err := apiClient.WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenWaarschuwingsadressenDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenWaarschuwingsadressenPartialUpdate

> Waarschuwingsadres AansluitingenWaarschuwingsadressenPartialUpdate(ctx, id).PatchedWaarschuwingsadresRequest(patchedWaarschuwingsadresRequest).Execute()



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
    patchedWaarschuwingsadresRequest := *openapiclient.NewPatchedWaarschuwingsadresRequest() // PatchedWaarschuwingsadresRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenPartialUpdate(context.Background(), id).PatchedWaarschuwingsadresRequest(patchedWaarschuwingsadresRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenWaarschuwingsadressenPartialUpdate`: Waarschuwingsadres
    fmt.Fprintf(os.Stdout, "Response from `WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenWaarschuwingsadressenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWaarschuwingsadresRequest** | [**PatchedWaarschuwingsadresRequest**](PatchedWaarschuwingsadresRequest.md) |  | 

### Return type

[**Waarschuwingsadres**](Waarschuwingsadres.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenWaarschuwingsadressenRetrieve

> Waarschuwingsadres AansluitingenWaarschuwingsadressenRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenWaarschuwingsadressenRetrieve`: Waarschuwingsadres
    fmt.Fprintf(os.Stdout, "Response from `WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenWaarschuwingsadressenRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Waarschuwingsadres**](Waarschuwingsadres.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenWaarschuwingsadressenUpdate

> Waarschuwingsadres AansluitingenWaarschuwingsadressenUpdate(ctx, id).WaarschuwingsadresRequest(waarschuwingsadresRequest).Execute()



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
    waarschuwingsadresRequest := *openapiclient.NewWaarschuwingsadresRequest(int32(123)) // WaarschuwingsadresRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenUpdate(context.Background(), id).WaarschuwingsadresRequest(waarschuwingsadresRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenWaarschuwingsadressenUpdate`: Waarschuwingsadres
    fmt.Fprintf(os.Stdout, "Response from `WaarschuwingsadressenApi.AansluitingenWaarschuwingsadressenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenWaarschuwingsadressenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **waarschuwingsadresRequest** | [**WaarschuwingsadresRequest**](WaarschuwingsadresRequest.md) |  | 

### Return type

[**Waarschuwingsadres**](Waarschuwingsadres.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

