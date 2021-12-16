# \MeldingenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeldingenList**](MeldingenApi.md#MeldingenList) | **Get** /api/meldingen | 



## MeldingenList

> PaginatedMeldingList MeldingenList(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeldingenListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeldingenListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actueel** | **optional.Bool**|  | 
 **createdAfter** | **optional.Time**|  | 
 **createdBefore** | **optional.Time**|  | 
 **criterium** | **optional.Int32**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **modifiedAfter** | **optional.Time**|  | 
 **modifiedBefore** | **optional.Time**|  | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 
 **toestand** | **optional.Int32**| 0 &#x3D; Actief, 1 &#x3D; test, 2 &#x3D; Passief | 

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

