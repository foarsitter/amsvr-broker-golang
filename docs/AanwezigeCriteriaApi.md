# \AanwezigeCriteriaApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenAanwezigeCriteriaDestroy**](AanwezigeCriteriaApi.md#AansluitingenAanwezigeCriteriaDestroy) | **Delete** /api/aansluitingen/aanwezige-criteria/{id} | 
[**AansluitingenAanwezigeCriteriaPartialUpdate**](AanwezigeCriteriaApi.md#AansluitingenAanwezigeCriteriaPartialUpdate) | **Patch** /api/aansluitingen/aanwezige-criteria/{id} | 
[**AansluitingenAanwezigeCriteriaRetrieve**](AanwezigeCriteriaApi.md#AansluitingenAanwezigeCriteriaRetrieve) | **Get** /api/aansluitingen/aanwezige-criteria/{id} | 
[**AansluitingenAanwezigeCriteriaUpdate**](AanwezigeCriteriaApi.md#AansluitingenAanwezigeCriteriaUpdate) | **Put** /api/aansluitingen/aanwezige-criteria/{id} | 
[**AansluitingenSectorenAanwezigeCriteriaCreate**](AanwezigeCriteriaApi.md#AansluitingenSectorenAanwezigeCriteriaCreate) | **Post** /api/aansluitingen/sectoren/{id}/aanwezige-criteria | 
[**AansluitingenSectorenAanwezigeCriteriaList**](AanwezigeCriteriaApi.md#AansluitingenSectorenAanwezigeCriteriaList) | **Get** /api/aansluitingen/sectoren/{id}/aanwezige-criteria | 



## AansluitingenAanwezigeCriteriaDestroy

> AansluitingenAanwezigeCriteriaDestroy(ctx, id).Execute()



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
    resp, r, err := apiClient.AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAansluitingenAanwezigeCriteriaDestroyRequest struct via the builder pattern


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


## AansluitingenAanwezigeCriteriaPartialUpdate

> AanwezigCriterium AansluitingenAanwezigeCriteriaPartialUpdate(ctx, id).PatchedAanwezigCriteriumRequest(patchedAanwezigCriteriumRequest).Execute()



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
    patchedAanwezigCriteriumRequest := *openapiclient.NewPatchedAanwezigCriteriumRequest() // PatchedAanwezigCriteriumRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaPartialUpdate(context.Background(), id).PatchedAanwezigCriteriumRequest(patchedAanwezigCriteriumRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenAanwezigeCriteriaPartialUpdate`: AanwezigCriterium
    fmt.Fprintf(os.Stdout, "Response from `AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenAanwezigeCriteriaPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAanwezigCriteriumRequest** | [**PatchedAanwezigCriteriumRequest**](PatchedAanwezigCriteriumRequest.md) |  | 

### Return type

[**AanwezigCriterium**](AanwezigCriterium.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenAanwezigeCriteriaRetrieve

> AanwezigCriterium AansluitingenAanwezigeCriteriaRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenAanwezigeCriteriaRetrieve`: AanwezigCriterium
    fmt.Fprintf(os.Stdout, "Response from `AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenAanwezigeCriteriaRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AanwezigCriterium**](AanwezigCriterium.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenAanwezigeCriteriaUpdate

> AanwezigCriterium AansluitingenAanwezigeCriteriaUpdate(ctx, id).AanwezigCriteriumRequest(aanwezigCriteriumRequest).Execute()



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
    aanwezigCriteriumRequest := *openapiclient.NewAanwezigCriteriumRequest(int32(123)) // AanwezigCriteriumRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaUpdate(context.Background(), id).AanwezigCriteriumRequest(aanwezigCriteriumRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenAanwezigeCriteriaUpdate`: AanwezigCriterium
    fmt.Fprintf(os.Stdout, "Response from `AanwezigeCriteriaApi.AansluitingenAanwezigeCriteriaUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenAanwezigeCriteriaUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aanwezigCriteriumRequest** | [**AanwezigCriteriumRequest**](AanwezigCriteriumRequest.md) |  | 

### Return type

[**AanwezigCriterium**](AanwezigCriterium.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenAanwezigeCriteriaCreate

> AanwezigCriterium AansluitingenSectorenAanwezigeCriteriaCreate(ctx, id).AanwezigCriteriumRequest(aanwezigCriteriumRequest).Execute()



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
    aanwezigCriteriumRequest := *openapiclient.NewAanwezigCriteriumRequest(int32(123)) // AanwezigCriteriumRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AanwezigeCriteriaApi.AansluitingenSectorenAanwezigeCriteriaCreate(context.Background(), id).AanwezigCriteriumRequest(aanwezigCriteriumRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AanwezigeCriteriaApi.AansluitingenSectorenAanwezigeCriteriaCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenAanwezigeCriteriaCreate`: AanwezigCriterium
    fmt.Fprintf(os.Stdout, "Response from `AanwezigeCriteriaApi.AansluitingenSectorenAanwezigeCriteriaCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenAanwezigeCriteriaCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aanwezigCriteriumRequest** | [**AanwezigCriteriumRequest**](AanwezigCriteriumRequest.md) |  | 

### Return type

[**AanwezigCriterium**](AanwezigCriterium.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenAanwezigeCriteriaList

> PaginatedAanwezigCriteriumList AansluitingenSectorenAanwezigeCriteriaList(ctx, id).Limit(limit).Offset(offset).Execute()



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
    resp, r, err := apiClient.AanwezigeCriteriaApi.AansluitingenSectorenAanwezigeCriteriaList(context.Background(), id).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AanwezigeCriteriaApi.AansluitingenSectorenAanwezigeCriteriaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenAanwezigeCriteriaList`: PaginatedAanwezigCriteriumList
    fmt.Fprintf(os.Stdout, "Response from `AanwezigeCriteriaApi.AansluitingenSectorenAanwezigeCriteriaList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenAanwezigeCriteriaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedAanwezigCriteriumList**](PaginatedAanwezigCriteriumList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

