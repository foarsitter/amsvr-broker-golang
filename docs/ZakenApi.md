# \ZakenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZakenAfsprakenCreate**](ZakenApi.md#ZakenAfsprakenCreate) | **Post** /api/zaken/{id}/afspraken | 
[**ZakenAfsprakenList**](ZakenApi.md#ZakenAfsprakenList) | **Get** /api/zaken/{id}/afspraken | 



## ZakenAfsprakenCreate

> AfspraakCreate ZakenAfsprakenCreate(ctx, id).AfspraakCreateRequest(afspraakCreateRequest).Execute()





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
    id := "id_example" // string | 
    afspraakCreateRequest := *openapiclient.NewAfspraakCreateRequest(time.Now()) // AfspraakCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZakenApi.ZakenAfsprakenCreate(context.Background(), id).AfspraakCreateRequest(afspraakCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZakenApi.ZakenAfsprakenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenAfsprakenCreate`: AfspraakCreate
    fmt.Fprintf(os.Stdout, "Response from `ZakenApi.ZakenAfsprakenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenAfsprakenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afspraakCreateRequest** | [**AfspraakCreateRequest**](AfspraakCreateRequest.md) |  | 

### Return type

[**AfspraakCreate**](AfspraakCreate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZakenAfsprakenList

> PaginatedAfspraakCreateList ZakenAfsprakenList(ctx, id).Limit(limit).Offset(offset).Execute()





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
    id := "id_example" // string | 
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZakenApi.ZakenAfsprakenList(context.Background(), id).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZakenApi.ZakenAfsprakenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZakenAfsprakenList`: PaginatedAfspraakCreateList
    fmt.Fprintf(os.Stdout, "Response from `ZakenApi.ZakenAfsprakenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZakenAfsprakenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedAfspraakCreateList**](PaginatedAfspraakCreateList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

