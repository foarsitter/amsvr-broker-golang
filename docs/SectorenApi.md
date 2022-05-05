# \SectorenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenSectorenCreate**](SectorenApi.md#AansluitingenSectorenCreate) | **Post** /api/aansluitingen/{aansluitnummer}/sectoren | 
[**AansluitingenSectorenDestroy**](SectorenApi.md#AansluitingenSectorenDestroy) | **Delete** /api/aansluitingen/sectoren/{id} | 
[**AansluitingenSectorenList**](SectorenApi.md#AansluitingenSectorenList) | **Get** /api/aansluitingen/{aansluitnummer}/sectoren | 
[**AansluitingenSectorenRetrieveById**](SectorenApi.md#AansluitingenSectorenRetrieveById) | **Get** /api/aansluitingen/sectoren/{id} | 
[**AansluitingenSectorenToestandPartialUpdateById**](SectorenApi.md#AansluitingenSectorenToestandPartialUpdateById) | **Patch** /api/aansluitingen/sectoren/{id}/toestand | 
[**AansluitingenSectorenUpdate**](SectorenApi.md#AansluitingenSectorenUpdate) | **Put** /api/aansluitingen/sectoren/{id} | 
[**SectorList**](SectorenApi.md#SectorList) | **Get** /api/aansluitingen/sectoren | 



## AansluitingenSectorenCreate

> Sector AansluitingenSectorenCreate(ctx, aansluitnummer).SectorRequest(sectorRequest).Execute()



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
    sectorRequest := *openapiclient.NewSectorRequest(int32(123), "Naam_example") // SectorRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SectorenApi.AansluitingenSectorenCreate(context.Background(), aansluitnummer).SectorRequest(sectorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectorenApi.AansluitingenSectorenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenCreate`: Sector
    fmt.Fprintf(os.Stdout, "Response from `SectorenApi.AansluitingenSectorenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sectorRequest** | [**SectorRequest**](SectorRequest.md) |  | 

### Return type

[**Sector**](Sector.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenDestroy

> AansluitingenSectorenDestroy(ctx, id).Execute()



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
    resp, r, err := apiClient.SectorenApi.AansluitingenSectorenDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectorenApi.AansluitingenSectorenDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAansluitingenSectorenDestroyRequest struct via the builder pattern


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


## AansluitingenSectorenList

> PaginatedSectorList AansluitingenSectorenList(ctx, aansluitnummer).Limit(limit).Offset(offset).Execute()



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
    resp, r, err := apiClient.SectorenApi.AansluitingenSectorenList(context.Background(), aansluitnummer).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectorenApi.AansluitingenSectorenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenList`: PaginatedSectorList
    fmt.Fprintf(os.Stdout, "Response from `SectorenApi.AansluitingenSectorenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedSectorList**](PaginatedSectorList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenRetrieveById

> Sector AansluitingenSectorenRetrieveById(ctx, id).Execute()



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
    resp, r, err := apiClient.SectorenApi.AansluitingenSectorenRetrieveById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectorenApi.AansluitingenSectorenRetrieveById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenRetrieveById`: Sector
    fmt.Fprintf(os.Stdout, "Response from `SectorenApi.AansluitingenSectorenRetrieveById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenRetrieveByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sector**](Sector.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenToestandPartialUpdateById

> SectorToestand AansluitingenSectorenToestandPartialUpdateById(ctx, id).PatchedSectorToestandRequest(patchedSectorToestandRequest).Execute()



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
    patchedSectorToestandRequest := *openapiclient.NewPatchedSectorToestandRequest() // PatchedSectorToestandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SectorenApi.AansluitingenSectorenToestandPartialUpdateById(context.Background(), id).PatchedSectorToestandRequest(patchedSectorToestandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectorenApi.AansluitingenSectorenToestandPartialUpdateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenToestandPartialUpdateById`: SectorToestand
    fmt.Fprintf(os.Stdout, "Response from `SectorenApi.AansluitingenSectorenToestandPartialUpdateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenToestandPartialUpdateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedSectorToestandRequest** | [**PatchedSectorToestandRequest**](PatchedSectorToestandRequest.md) |  | 

### Return type

[**SectorToestand**](SectorToestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenUpdate

> SectorWrite AansluitingenSectorenUpdate(ctx, id).SectorWriteRequest(sectorWriteRequest).Execute()



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
    sectorWriteRequest := *openapiclient.NewSectorWriteRequest(int32(123), "Naam_example", []int32{int32(123)}) // SectorWriteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SectorenApi.AansluitingenSectorenUpdate(context.Background(), id).SectorWriteRequest(sectorWriteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectorenApi.AansluitingenSectorenUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenUpdate`: SectorWrite
    fmt.Fprintf(os.Stdout, "Response from `SectorenApi.AansluitingenSectorenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sectorWriteRequest** | [**SectorWriteRequest**](SectorWriteRequest.md) |  | 

### Return type

[**SectorWrite**](SectorWrite.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SectorList

> PaginatedSectorList SectorList(ctx).BagNietBeschikbaar(bagNietBeschikbaar).Huisnummer(huisnummer).IsBouwdeelNull(isBouwdeelNull).Limit(limit).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).Naam(naam).Nummer(nummer).Offset(offset).Postcode(postcode).Execute()



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
    bagNietBeschikbaar := true // bool |  (optional)
    huisnummer := "huisnummer_example" // string |  (optional)
    isBouwdeelNull := true // bool |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    modifiedAfter := time.Now() // string |  (optional)
    modifiedBefore := time.Now() // string |  (optional)
    naam := "naam_example" // string | Naam (optional)
    nummer := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    postcode := "postcode_example" // string | Postcode (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SectorenApi.SectorList(context.Background()).BagNietBeschikbaar(bagNietBeschikbaar).Huisnummer(huisnummer).IsBouwdeelNull(isBouwdeelNull).Limit(limit).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).Naam(naam).Nummer(nummer).Offset(offset).Postcode(postcode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectorenApi.SectorList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SectorList`: PaginatedSectorList
    fmt.Fprintf(os.Stdout, "Response from `SectorenApi.SectorList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSectorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bagNietBeschikbaar** | **bool** |  | 
 **huisnummer** | **string** |  | 
 **isBouwdeelNull** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **modifiedAfter** | **string** |  | 
 **modifiedBefore** | **string** |  | 
 **naam** | **string** | Naam | 
 **nummer** | **int32** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **postcode** | **string** | Postcode | 

### Return type

[**PaginatedSectorList**](PaginatedSectorList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

