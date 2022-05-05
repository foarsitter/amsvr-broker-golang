# \TelefoonnummersApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenContactpersonenTelefoonnummersCreate**](TelefoonnummersApi.md#AansluitingenContactpersonenTelefoonnummersCreate) | **Post** /api/aansluitingen/contactpersonen/{id}/telefoonnummers | 
[**AansluitingenContactpersonenTelefoonnummersList**](TelefoonnummersApi.md#AansluitingenContactpersonenTelefoonnummersList) | **Get** /api/aansluitingen/contactpersonen/{id}/telefoonnummers | 
[**AansluitingenTelefoonnummersDestroy**](TelefoonnummersApi.md#AansluitingenTelefoonnummersDestroy) | **Delete** /api/aansluitingen/telefoonnummers/{id} | 
[**AansluitingenTelefoonnummersPartialUpdate**](TelefoonnummersApi.md#AansluitingenTelefoonnummersPartialUpdate) | **Patch** /api/aansluitingen/telefoonnummers/{id} | 
[**AansluitingenTelefoonnummersRetrieve**](TelefoonnummersApi.md#AansluitingenTelefoonnummersRetrieve) | **Get** /api/aansluitingen/telefoonnummers/{id} | 
[**AansluitingenTelefoonnummersUpdate**](TelefoonnummersApi.md#AansluitingenTelefoonnummersUpdate) | **Put** /api/aansluitingen/telefoonnummers/{id} | 



## AansluitingenContactpersonenTelefoonnummersCreate

> Telefoonnummer AansluitingenContactpersonenTelefoonnummersCreate(ctx, id).TelefoonnummerRequest(telefoonnummerRequest).Execute()



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
    telefoonnummerRequest := *openapiclient.NewTelefoonnummerRequest("Telefoonnummer_example") // TelefoonnummerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TelefoonnummersApi.AansluitingenContactpersonenTelefoonnummersCreate(context.Background(), id).TelefoonnummerRequest(telefoonnummerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelefoonnummersApi.AansluitingenContactpersonenTelefoonnummersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenContactpersonenTelefoonnummersCreate`: Telefoonnummer
    fmt.Fprintf(os.Stdout, "Response from `TelefoonnummersApi.AansluitingenContactpersonenTelefoonnummersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenContactpersonenTelefoonnummersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **telefoonnummerRequest** | [**TelefoonnummerRequest**](TelefoonnummerRequest.md) |  | 

### Return type

[**Telefoonnummer**](Telefoonnummer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenContactpersonenTelefoonnummersList

> PaginatedTelefoonnummerList AansluitingenContactpersonenTelefoonnummersList(ctx, id).Limit(limit).Offset(offset).Execute()



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
    resp, r, err := apiClient.TelefoonnummersApi.AansluitingenContactpersonenTelefoonnummersList(context.Background(), id).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelefoonnummersApi.AansluitingenContactpersonenTelefoonnummersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenContactpersonenTelefoonnummersList`: PaginatedTelefoonnummerList
    fmt.Fprintf(os.Stdout, "Response from `TelefoonnummersApi.AansluitingenContactpersonenTelefoonnummersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenContactpersonenTelefoonnummersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedTelefoonnummerList**](PaginatedTelefoonnummerList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenTelefoonnummersDestroy

> AansluitingenTelefoonnummersDestroy(ctx, id).Execute()



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
    resp, r, err := apiClient.TelefoonnummersApi.AansluitingenTelefoonnummersDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelefoonnummersApi.AansluitingenTelefoonnummersDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAansluitingenTelefoonnummersDestroyRequest struct via the builder pattern


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


## AansluitingenTelefoonnummersPartialUpdate

> Telefoonnummer AansluitingenTelefoonnummersPartialUpdate(ctx, id).PatchedTelefoonnummerRequest(patchedTelefoonnummerRequest).Execute()



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
    patchedTelefoonnummerRequest := *openapiclient.NewPatchedTelefoonnummerRequest() // PatchedTelefoonnummerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TelefoonnummersApi.AansluitingenTelefoonnummersPartialUpdate(context.Background(), id).PatchedTelefoonnummerRequest(patchedTelefoonnummerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelefoonnummersApi.AansluitingenTelefoonnummersPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenTelefoonnummersPartialUpdate`: Telefoonnummer
    fmt.Fprintf(os.Stdout, "Response from `TelefoonnummersApi.AansluitingenTelefoonnummersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenTelefoonnummersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTelefoonnummerRequest** | [**PatchedTelefoonnummerRequest**](PatchedTelefoonnummerRequest.md) |  | 

### Return type

[**Telefoonnummer**](Telefoonnummer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenTelefoonnummersRetrieve

> Telefoonnummer AansluitingenTelefoonnummersRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.TelefoonnummersApi.AansluitingenTelefoonnummersRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelefoonnummersApi.AansluitingenTelefoonnummersRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenTelefoonnummersRetrieve`: Telefoonnummer
    fmt.Fprintf(os.Stdout, "Response from `TelefoonnummersApi.AansluitingenTelefoonnummersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenTelefoonnummersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Telefoonnummer**](Telefoonnummer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenTelefoonnummersUpdate

> Telefoonnummer AansluitingenTelefoonnummersUpdate(ctx, id).TelefoonnummerRequest(telefoonnummerRequest).Execute()



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
    telefoonnummerRequest := *openapiclient.NewTelefoonnummerRequest("Telefoonnummer_example") // TelefoonnummerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TelefoonnummersApi.AansluitingenTelefoonnummersUpdate(context.Background(), id).TelefoonnummerRequest(telefoonnummerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelefoonnummersApi.AansluitingenTelefoonnummersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenTelefoonnummersUpdate`: Telefoonnummer
    fmt.Fprintf(os.Stdout, "Response from `TelefoonnummersApi.AansluitingenTelefoonnummersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenTelefoonnummersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **telefoonnummerRequest** | [**TelefoonnummerRequest**](TelefoonnummerRequest.md) |  | 

### Return type

[**Telefoonnummer**](Telefoonnummer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

