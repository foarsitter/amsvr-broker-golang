# \ContactpersonenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenContactpersonenDestroy**](ContactpersonenApi.md#AansluitingenContactpersonenDestroy) | **Delete** /api/aansluitingen/contactpersonen/{id} | 
[**AansluitingenContactpersonenPartialUpdate**](ContactpersonenApi.md#AansluitingenContactpersonenPartialUpdate) | **Patch** /api/aansluitingen/contactpersonen/{id} | 
[**AansluitingenContactpersonenRetrieve**](ContactpersonenApi.md#AansluitingenContactpersonenRetrieve) | **Get** /api/aansluitingen/contactpersonen/{id} | 
[**AansluitingenContactpersonenUpdate**](ContactpersonenApi.md#AansluitingenContactpersonenUpdate) | **Put** /api/aansluitingen/contactpersonen/{id} | 
[**ContactpersoonList**](ContactpersonenApi.md#ContactpersoonList) | **Get** /api/aansluitingen/contactpersonen | 



## AansluitingenContactpersonenDestroy

> AansluitingenContactpersonenDestroy(ctx, id).Execute()



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
    resp, r, err := apiClient.ContactpersonenApi.AansluitingenContactpersonenDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactpersonenApi.AansluitingenContactpersonenDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAansluitingenContactpersonenDestroyRequest struct via the builder pattern


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


## AansluitingenContactpersonenPartialUpdate

> ContactPersoon AansluitingenContactpersonenPartialUpdate(ctx, id).PatchedContactPersoonRequest(patchedContactPersoonRequest).Execute()



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
    patchedContactPersoonRequest := *openapiclient.NewPatchedContactPersoonRequest() // PatchedContactPersoonRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactpersonenApi.AansluitingenContactpersonenPartialUpdate(context.Background(), id).PatchedContactPersoonRequest(patchedContactPersoonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactpersonenApi.AansluitingenContactpersonenPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenContactpersonenPartialUpdate`: ContactPersoon
    fmt.Fprintf(os.Stdout, "Response from `ContactpersonenApi.AansluitingenContactpersonenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenContactpersonenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedContactPersoonRequest** | [**PatchedContactPersoonRequest**](PatchedContactPersoonRequest.md) |  | 

### Return type

[**ContactPersoon**](ContactPersoon.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenContactpersonenRetrieve

> ContactPersoon AansluitingenContactpersonenRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.ContactpersonenApi.AansluitingenContactpersonenRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactpersonenApi.AansluitingenContactpersonenRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenContactpersonenRetrieve`: ContactPersoon
    fmt.Fprintf(os.Stdout, "Response from `ContactpersonenApi.AansluitingenContactpersonenRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenContactpersonenRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContactPersoon**](ContactPersoon.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenContactpersonenUpdate

> ContactPersoon AansluitingenContactpersonenUpdate(ctx, id).ContactPersoonRequest(contactPersoonRequest).Execute()



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
    contactPersoonRequest := *openapiclient.NewContactPersoonRequest("Achternaam_example") // ContactPersoonRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactpersonenApi.AansluitingenContactpersonenUpdate(context.Background(), id).ContactPersoonRequest(contactPersoonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactpersonenApi.AansluitingenContactpersonenUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenContactpersonenUpdate`: ContactPersoon
    fmt.Fprintf(os.Stdout, "Response from `ContactpersonenApi.AansluitingenContactpersonenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenContactpersonenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactPersoonRequest** | [**ContactPersoonRequest**](ContactPersoonRequest.md) |  | 

### Return type

[**ContactPersoon**](ContactPersoon.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactpersoonList

> PaginatedContactPersoonList ContactpersoonList(ctx).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Limit(limit).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).Offset(offset).Voornaam(voornaam).Execute()



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
    createdAfter := time.Now() // time.Time |  (optional)
    createdBefore := time.Now() // time.Time |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    modifiedAfter := time.Now() // time.Time |  (optional)
    modifiedBefore := time.Now() // time.Time |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    voornaam := "voornaam_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactpersonenApi.ContactpersoonList(context.Background()).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Limit(limit).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).Offset(offset).Voornaam(voornaam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactpersonenApi.ContactpersoonList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactpersoonList`: PaginatedContactPersoonList
    fmt.Fprintf(os.Stdout, "Response from `ContactpersonenApi.ContactpersoonList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactpersoonListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAfter** | **time.Time** |  | 
 **createdBefore** | **time.Time** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **modifiedAfter** | **time.Time** |  | 
 **modifiedBefore** | **time.Time** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **voornaam** | **string** |  | 

### Return type

[**PaginatedContactPersoonList**](PaginatedContactPersoonList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

