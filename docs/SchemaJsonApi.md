# \SchemaJsonApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchemaJsonRetrieve**](SchemaJsonApi.md#SchemaJsonRetrieve) | **Get** /schema.json | 



## SchemaJsonRetrieve

> map[string]interface{} SchemaJsonRetrieve(ctx).Lang(lang).Execute()





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
    lang := "lang_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaJsonApi.SchemaJsonRetrieve(context.Background()).Lang(lang).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaJsonApi.SchemaJsonRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaJsonRetrieve`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SchemaJsonApi.SchemaJsonRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchemaJsonRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lang** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.oai.openapi+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

