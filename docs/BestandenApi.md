# \BestandenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenBestandenCreate**](BestandenApi.md#AansluitingenBestandenCreate) | **Post** /api/aansluitingen/{aansluitnummer}/bestanden | 
[**AansluitingenBestandenDestroy**](BestandenApi.md#AansluitingenBestandenDestroy) | **Delete** /api/aansluitingen/bestanden/{id} | 
[**AansluitingenBestandenList**](BestandenApi.md#AansluitingenBestandenList) | **Get** /api/aansluitingen/{aansluitnummer}/bestanden | 
[**AansluitingenBestandenPartialUpdate**](BestandenApi.md#AansluitingenBestandenPartialUpdate) | **Patch** /api/aansluitingen/bestanden/{id} | 
[**AansluitingenBestandenRetrieve**](BestandenApi.md#AansluitingenBestandenRetrieve) | **Get** /api/aansluitingen/bestanden/{id} | 
[**AansluitingenBestandenUpdate**](BestandenApi.md#AansluitingenBestandenUpdate) | **Put** /api/aansluitingen/bestanden/{id} | 



## AansluitingenBestandenCreate

> AansluitingBestand AansluitingenBestandenCreate(ctx, aansluitnummer).AansluitingBestandRequest(aansluitingBestandRequest).Execute()



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
    aansluitingBestandRequest := *openapiclient.NewAansluitingBestandRequest() // AansluitingBestandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BestandenApi.AansluitingenBestandenCreate(context.Background(), aansluitnummer).AansluitingBestandRequest(aansluitingBestandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BestandenApi.AansluitingenBestandenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenBestandenCreate`: AansluitingBestand
    fmt.Fprintf(os.Stdout, "Response from `BestandenApi.AansluitingenBestandenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenBestandenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aansluitingBestandRequest** | [**AansluitingBestandRequest**](AansluitingBestandRequest.md) |  | 

### Return type

[**AansluitingBestand**](AansluitingBestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenBestandenDestroy

> AansluitingenBestandenDestroy(ctx, id).Execute()



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
    resp, r, err := apiClient.BestandenApi.AansluitingenBestandenDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BestandenApi.AansluitingenBestandenDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAansluitingenBestandenDestroyRequest struct via the builder pattern


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


## AansluitingenBestandenList

> PaginatedAansluitingBestandList AansluitingenBestandenList(ctx, aansluitnummer).Limit(limit).Offset(offset).Execute()



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BestandenApi.AansluitingenBestandenList(context.Background(), aansluitnummer).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BestandenApi.AansluitingenBestandenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenBestandenList`: PaginatedAansluitingBestandList
    fmt.Fprintf(os.Stdout, "Response from `BestandenApi.AansluitingenBestandenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenBestandenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedAansluitingBestandList**](PaginatedAansluitingBestandList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenBestandenPartialUpdate

> AansluitingBestand AansluitingenBestandenPartialUpdate(ctx, id).PatchedAansluitingBestandRequest(patchedAansluitingBestandRequest).Execute()



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
    patchedAansluitingBestandRequest := *openapiclient.NewPatchedAansluitingBestandRequest() // PatchedAansluitingBestandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BestandenApi.AansluitingenBestandenPartialUpdate(context.Background(), id).PatchedAansluitingBestandRequest(patchedAansluitingBestandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BestandenApi.AansluitingenBestandenPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenBestandenPartialUpdate`: AansluitingBestand
    fmt.Fprintf(os.Stdout, "Response from `BestandenApi.AansluitingenBestandenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenBestandenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAansluitingBestandRequest** | [**PatchedAansluitingBestandRequest**](PatchedAansluitingBestandRequest.md) |  | 

### Return type

[**AansluitingBestand**](AansluitingBestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenBestandenRetrieve

> AansluitingBestand AansluitingenBestandenRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.BestandenApi.AansluitingenBestandenRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BestandenApi.AansluitingenBestandenRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenBestandenRetrieve`: AansluitingBestand
    fmt.Fprintf(os.Stdout, "Response from `BestandenApi.AansluitingenBestandenRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenBestandenRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AansluitingBestand**](AansluitingBestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenBestandenUpdate

> AansluitingBestand AansluitingenBestandenUpdate(ctx, id).AansluitingBestandRequest(aansluitingBestandRequest).Execute()



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
    aansluitingBestandRequest := *openapiclient.NewAansluitingBestandRequest() // AansluitingBestandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BestandenApi.AansluitingenBestandenUpdate(context.Background(), id).AansluitingBestandRequest(aansluitingBestandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BestandenApi.AansluitingenBestandenUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenBestandenUpdate`: AansluitingBestand
    fmt.Fprintf(os.Stdout, "Response from `BestandenApi.AansluitingenBestandenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenBestandenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aansluitingBestandRequest** | [**AansluitingBestandRequest**](AansluitingBestandRequest.md) |  | 

### Return type

[**AansluitingBestand**](AansluitingBestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

