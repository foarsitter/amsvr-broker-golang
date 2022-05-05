# \AfsprakenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AfsprakenBeschikbaarheidDagenRetrieve**](AfsprakenApi.md#AfsprakenBeschikbaarheidDagenRetrieve) | **Get** /api/afspraken/beschikbaarheid/{regio_code}/dagen | 
[**AfsprakenBeschikbaarheidTijdslotsRetrieve**](AfsprakenApi.md#AfsprakenBeschikbaarheidTijdslotsRetrieve) | **Get** /api/afspraken/beschikbaarheid/{regio_code}/tijdslots/{date} | 
[**AfsprakenCreateCreate**](AfsprakenApi.md#AfsprakenCreateCreate) | **Post** /api/afspraken/create/{regio_code} | 
[**AfsprakenList**](AfsprakenApi.md#AfsprakenList) | **Get** /api/afspraken/ | 
[**AfsprakenPartialUpdate**](AfsprakenApi.md#AfsprakenPartialUpdate) | **Patch** /api/afspraken/{id} | 
[**AfsprakenRetrieve**](AfsprakenApi.md#AfsprakenRetrieve) | **Get** /api/afspraken/{id} | 
[**AfsprakenUpdate**](AfsprakenApi.md#AfsprakenUpdate) | **Put** /api/afspraken/{id} | 
[**AfsprakenVerplaatsenCreate**](AfsprakenApi.md#AfsprakenVerplaatsenCreate) | **Post** /api/afspraken/{id}/verplaatsen | 



## AfsprakenBeschikbaarheidDagenRetrieve

> Bookabledays AfsprakenBeschikbaarheidDagenRetrieve(ctx, regioCode).Execute()





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
    regioCode := "regioCode_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsprakenApi.AfsprakenBeschikbaarheidDagenRetrieve(context.Background(), regioCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsprakenApi.AfsprakenBeschikbaarheidDagenRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfsprakenBeschikbaarheidDagenRetrieve`: Bookabledays
    fmt.Fprintf(os.Stdout, "Response from `AfsprakenApi.AfsprakenBeschikbaarheidDagenRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regioCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfsprakenBeschikbaarheidDagenRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Bookabledays**](Bookabledays.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenBeschikbaarheidTijdslotsRetrieve

> BookableTimes AfsprakenBeschikbaarheidTijdslotsRetrieve(ctx, date, regioCode).Execute()





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
    date := "date_example" // string | 
    regioCode := "regioCode_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsprakenApi.AfsprakenBeschikbaarheidTijdslotsRetrieve(context.Background(), date, regioCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsprakenApi.AfsprakenBeschikbaarheidTijdslotsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfsprakenBeschikbaarheidTijdslotsRetrieve`: BookableTimes
    fmt.Fprintf(os.Stdout, "Response from `AfsprakenApi.AfsprakenBeschikbaarheidTijdslotsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**date** | **string** |  | 
**regioCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfsprakenBeschikbaarheidTijdslotsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BookableTimes**](BookableTimes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenCreateCreate

> Afspraak AfsprakenCreateCreate(ctx, regioCode).AfspraakRequest(afspraakRequest).Execute()



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
    regioCode := "regioCode_example" // string | 
    afspraakRequest := *openapiclient.NewAfspraakRequest(time.Now(), "Titel_example") // AfspraakRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsprakenApi.AfsprakenCreateCreate(context.Background(), regioCode).AfspraakRequest(afspraakRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsprakenApi.AfsprakenCreateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfsprakenCreateCreate`: Afspraak
    fmt.Fprintf(os.Stdout, "Response from `AfsprakenApi.AfsprakenCreateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regioCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfsprakenCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afspraakRequest** | [**AfspraakRequest**](AfspraakRequest.md) |  | 

### Return type

[**Afspraak**](Afspraak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenList

> PaginatedAfspraakList AfsprakenList(ctx).EindTijd(eindTijd).Limit(limit).Offset(offset).StartTijd(startTijd).Execute()



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
    eindTijd := time.Now() // time.Time |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    startTijd := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsprakenApi.AfsprakenList(context.Background()).EindTijd(eindTijd).Limit(limit).Offset(offset).StartTijd(startTijd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsprakenApi.AfsprakenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfsprakenList`: PaginatedAfspraakList
    fmt.Fprintf(os.Stdout, "Response from `AfsprakenApi.AfsprakenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAfsprakenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eindTijd** | **time.Time** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **startTijd** | **time.Time** |  | 

### Return type

[**PaginatedAfspraakList**](PaginatedAfspraakList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenPartialUpdate

> Afspraak AfsprakenPartialUpdate(ctx, id).PatchedAfspraakRequest(patchedAfspraakRequest).Execute()



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
    patchedAfspraakRequest := *openapiclient.NewPatchedAfspraakRequest() // PatchedAfspraakRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsprakenApi.AfsprakenPartialUpdate(context.Background(), id).PatchedAfspraakRequest(patchedAfspraakRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsprakenApi.AfsprakenPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfsprakenPartialUpdate`: Afspraak
    fmt.Fprintf(os.Stdout, "Response from `AfsprakenApi.AfsprakenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfsprakenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAfspraakRequest** | [**PatchedAfspraakRequest**](PatchedAfspraakRequest.md) |  | 

### Return type

[**Afspraak**](Afspraak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenRetrieve

> Afspraak AfsprakenRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.AfsprakenApi.AfsprakenRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsprakenApi.AfsprakenRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfsprakenRetrieve`: Afspraak
    fmt.Fprintf(os.Stdout, "Response from `AfsprakenApi.AfsprakenRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfsprakenRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Afspraak**](Afspraak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenUpdate

> Afspraak AfsprakenUpdate(ctx, id).AfspraakRequest(afspraakRequest).Execute()



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
    id := int32(56) // int32 | 
    afspraakRequest := *openapiclient.NewAfspraakRequest(time.Now(), "Titel_example") // AfspraakRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsprakenApi.AfsprakenUpdate(context.Background(), id).AfspraakRequest(afspraakRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsprakenApi.AfsprakenUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfsprakenUpdate`: Afspraak
    fmt.Fprintf(os.Stdout, "Response from `AfsprakenApi.AfsprakenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfsprakenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afspraakRequest** | [**AfspraakRequest**](AfspraakRequest.md) |  | 

### Return type

[**Afspraak**](Afspraak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfsprakenVerplaatsenCreate

> AfspraakCreate AfsprakenVerplaatsenCreate(ctx, id).AfspraakCreateRequest(afspraakCreateRequest).Execute()



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
    id := int32(56) // int32 | 
    afspraakCreateRequest := *openapiclient.NewAfspraakCreateRequest(time.Now()) // AfspraakCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AfsprakenApi.AfsprakenVerplaatsenCreate(context.Background(), id).AfspraakCreateRequest(afspraakCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AfsprakenApi.AfsprakenVerplaatsenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfsprakenVerplaatsenCreate`: AfspraakCreate
    fmt.Fprintf(os.Stdout, "Response from `AfsprakenApi.AfsprakenVerplaatsenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfsprakenVerplaatsenCreateRequest struct via the builder pattern


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

