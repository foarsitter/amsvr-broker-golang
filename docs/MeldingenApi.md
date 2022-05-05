# \MeldingenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeldingenList**](MeldingenApi.md#MeldingenList) | **Get** /api/meldingen | 



## MeldingenList

> PaginatedMeldingList MeldingenList(ctx).Actueel(actueel).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Criterium(criterium).Limit(limit).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).Offset(offset).Toestand(toestand).Execute()



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
    actueel := true // bool |  (optional)
    createdAfter := time.Now() // time.Time |  (optional)
    createdBefore := time.Now() // time.Time |  (optional)
    criterium := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    modifiedAfter := time.Now() // time.Time |  (optional)
    modifiedBefore := time.Now() // time.Time |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    toestand := int32(56) // int32 | 0 = Actief, 1 = Test, 2 = Passief (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MeldingenApi.MeldingenList(context.Background()).Actueel(actueel).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Criterium(criterium).Limit(limit).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).Offset(offset).Toestand(toestand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeldingenApi.MeldingenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeldingenList`: PaginatedMeldingList
    fmt.Fprintf(os.Stdout, "Response from `MeldingenApi.MeldingenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeldingenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actueel** | **bool** |  | 
 **createdAfter** | **time.Time** |  | 
 **createdBefore** | **time.Time** |  | 
 **criterium** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **modifiedAfter** | **time.Time** |  | 
 **modifiedBefore** | **time.Time** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **toestand** | **int32** | 0 &#x3D; Actief, 1 &#x3D; Test, 2 &#x3D; Passief | 

### Return type

[**PaginatedMeldingList**](PaginatedMeldingList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

