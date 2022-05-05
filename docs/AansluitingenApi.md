# \AansluitingenApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AansluitingenContactpersonenCreate**](AansluitingenApi.md#AansluitingenContactpersonenCreate) | **Post** /api/aansluitingen/{aansluitnummer}/contactpersonen | 
[**AansluitingenContactpersonenList**](AansluitingenApi.md#AansluitingenContactpersonenList) | **Get** /api/aansluitingen/{aansluitnummer}/contactpersonen | 
[**AansluitingenCreate**](AansluitingenApi.md#AansluitingenCreate) | **Post** /api/aansluitingen | 
[**AansluitingenDestroy**](AansluitingenApi.md#AansluitingenDestroy) | **Delete** /api/aansluitingen/{aansluitnummer} | 
[**AansluitingenList**](AansluitingenApi.md#AansluitingenList) | **Get** /api/aansluitingen | 
[**AansluitingenLopendeZaakRetrieve**](AansluitingenApi.md#AansluitingenLopendeZaakRetrieve) | **Get** /api/aansluitingen/{aansluitnummer}/lopende-zaak | 
[**AansluitingenMeldingenCreate**](AansluitingenApi.md#AansluitingenMeldingenCreate) | **Post** /api/aansluitingen/{aansluitnummer}/meldingen | 
[**AansluitingenMeldingenList**](AansluitingenApi.md#AansluitingenMeldingenList) | **Get** /api/aansluitingen/{aansluitnummer}/meldingen | 
[**AansluitingenMigratieCreate**](AansluitingenApi.md#AansluitingenMigratieCreate) | **Post** /api/aansluitingen/{aansluitnummer}/migratie | 
[**AansluitingenRetrieve**](AansluitingenApi.md#AansluitingenRetrieve) | **Get** /api/aansluitingen/{aansluitnummer} | 
[**AansluitingenSearchRetrieve**](AansluitingenApi.md#AansluitingenSearchRetrieve) | **Get** /api/aansluitingen/search | 
[**AansluitingenSectorenRetrieve**](AansluitingenApi.md#AansluitingenSectorenRetrieve) | **Get** /api/aansluitingen/{aansluitnummer}/sectoren/{nummer} | 
[**AansluitingenSectorenToestandPartialUpdate**](AansluitingenApi.md#AansluitingenSectorenToestandPartialUpdate) | **Patch** /api/aansluitingen/{aansluitnummer}/sectoren/toestand | 
[**AansluitingenSectorenToestandPartialUpdate2**](AansluitingenApi.md#AansluitingenSectorenToestandPartialUpdate2) | **Patch** /api/aansluitingen/{aansluitnummer}/sectoren/{nummer}/toestand | 
[**AansluitingenToestandPartialUpdate**](AansluitingenApi.md#AansluitingenToestandPartialUpdate) | **Patch** /api/aansluitingen/{aansluitnummer}/toestand | 
[**AansluitingenTransitiesPartialUpdate**](AansluitingenApi.md#AansluitingenTransitiesPartialUpdate) | **Patch** /api/aansluitingen/{aansluitnummer}/transities | 
[**AansluitingenTransitiesRetrieve**](AansluitingenApi.md#AansluitingenTransitiesRetrieve) | **Get** /api/aansluitingen/{aansluitnummer}/transities | 
[**AansluitingenUpdate**](AansluitingenApi.md#AansluitingenUpdate) | **Put** /api/aansluitingen/{aansluitnummer} | 
[**AansluitingenValidatieRetrieve**](AansluitingenApi.md#AansluitingenValidatieRetrieve) | **Get** /api/aansluitingen/{aansluitnummer}/validatie | 



## AansluitingenContactpersonenCreate

> ContactPersoon AansluitingenContactpersonenCreate(ctx, aansluitnummer).ContactPersoonRequest(contactPersoonRequest).Execute()



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
    contactPersoonRequest := *openapiclient.NewContactPersoonRequest("Achternaam_example") // ContactPersoonRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenContactpersonenCreate(context.Background(), aansluitnummer).ContactPersoonRequest(contactPersoonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenContactpersonenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenContactpersonenCreate`: ContactPersoon
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenContactpersonenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenContactpersonenCreateRequest struct via the builder pattern


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


## AansluitingenContactpersonenList

> PaginatedContactPersoonList AansluitingenContactpersonenList(ctx, aansluitnummer).Limit(limit).Offset(offset).Execute()



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
    resp, r, err := apiClient.AansluitingenApi.AansluitingenContactpersonenList(context.Background(), aansluitnummer).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenContactpersonenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenContactpersonenList`: PaginatedContactPersoonList
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenContactpersonenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenContactpersonenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

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


## AansluitingenCreate

> Aansluiting AansluitingenCreate(ctx).AansluitingRequest(aansluitingRequest).Execute()





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
    aansluitingRequest := *openapiclient.NewAansluitingRequest("Naam_example", "Aansluitnummer_example", int32(123)) // AansluitingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenCreate(context.Background()).AansluitingRequest(aansluitingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenCreate`: Aansluiting
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aansluitingRequest** | [**AansluitingRequest**](AansluitingRequest.md) |  | 

### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenDestroy

> AansluitingenDestroy(ctx, aansluitnummer).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenDestroy(context.Background(), aansluitnummer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenDestroyRequest struct via the builder pattern


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


## AansluitingenList

> PaginatedAansluitingList AansluitingenList(ctx).Aansluitnummer(aansluitnummer).Atsp(atsp).Bijzonderheden(bijzonderheden).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Gebruiksfunctie(gebruiksfunctie).Gemeente(gemeente).IsTestAansluiting(isTestAansluiting).Limit(limit).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).Naam(naam).Offset(offset).Plaats(plaats).Postcode(postcode).Straat(straat).TestExpired(testExpired).Execute()



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
    aansluitnummer := "aansluitnummer_example" // string | Aansluitnummer (optional)
    atsp := int32(56) // int32 |  (optional)
    bijzonderheden := true // bool | Heeft bijzonderheden (optional)
    createdAfter := time.Now() // time.Time |  (optional)
    createdBefore := time.Now() // time.Time |  (optional)
    gebruiksfunctie := "gebruiksfunctie_example" // string |  (optional)
    gemeente := int32(56) // int32 |  (optional)
    isTestAansluiting := true // bool |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    modifiedAfter := time.Now() // time.Time |  (optional)
    modifiedBefore := time.Now() // time.Time |  (optional)
    naam := "naam_example" // string | Naam (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    plaats := "plaats_example" // string | Plaats (optional)
    postcode := "postcode_example" // string | Postcode (optional)
    straat := "straat_example" // string | Straat (optional)
    testExpired := true // bool | Testperiode verlopen (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenList(context.Background()).Aansluitnummer(aansluitnummer).Atsp(atsp).Bijzonderheden(bijzonderheden).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Gebruiksfunctie(gebruiksfunctie).Gemeente(gemeente).IsTestAansluiting(isTestAansluiting).Limit(limit).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).Naam(naam).Offset(offset).Plaats(plaats).Postcode(postcode).Straat(straat).TestExpired(testExpired).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenList`: PaginatedAansluitingList
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aansluitnummer** | **string** | Aansluitnummer | 
 **atsp** | **int32** |  | 
 **bijzonderheden** | **bool** | Heeft bijzonderheden | 
 **createdAfter** | **time.Time** |  | 
 **createdBefore** | **time.Time** |  | 
 **gebruiksfunctie** | **string** |  | 
 **gemeente** | **int32** |  | 
 **isTestAansluiting** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **modifiedAfter** | **time.Time** |  | 
 **modifiedBefore** | **time.Time** |  | 
 **naam** | **string** | Naam | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **plaats** | **string** | Plaats | 
 **postcode** | **string** | Postcode | 
 **straat** | **string** | Straat | 
 **testExpired** | **bool** | Testperiode verlopen | 

### Return type

[**PaginatedAansluitingList**](PaginatedAansluitingList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenLopendeZaakRetrieve

> Zaak AansluitingenLopendeZaakRetrieve(ctx, aansluitnummer).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenLopendeZaakRetrieve(context.Background(), aansluitnummer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenLopendeZaakRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenLopendeZaakRetrieve`: Zaak
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenLopendeZaakRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenLopendeZaakRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Zaak**](Zaak.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenMeldingenCreate

> Melding AansluitingenMeldingenCreate(ctx, aansluitnummer).MeldingRequest(meldingRequest).Execute()



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
    aansluitnummer := "aansluitnummer_example" // string | 
    meldingRequest := *openapiclient.NewMeldingRequest(int32(123), time.Now(), "TODO", int32(123), int32(123)) // MeldingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenMeldingenCreate(context.Background(), aansluitnummer).MeldingRequest(meldingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenMeldingenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenMeldingenCreate`: Melding
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenMeldingenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenMeldingenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **meldingRequest** | [**MeldingRequest**](MeldingRequest.md) |  | 

### Return type

[**Melding**](Melding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenMeldingenList

> PaginatedMeldingList AansluitingenMeldingenList(ctx, aansluitnummer).Limit(limit).Offset(offset).Execute()



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
    resp, r, err := apiClient.AansluitingenApi.AansluitingenMeldingenList(context.Background(), aansluitnummer).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenMeldingenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenMeldingenList`: PaginatedMeldingList
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenMeldingenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenMeldingenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

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


## AansluitingenMigratieCreate

> Aansluiting AansluitingenMigratieCreate(ctx, aansluitnummer).MigratieRequest(migratieRequest).Execute()





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
    migratieRequest := *openapiclient.NewMigratieRequest("NieuweAansluitnummer_example", int32(123)) // MigratieRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenMigratieCreate(context.Background(), aansluitnummer).MigratieRequest(migratieRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenMigratieCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenMigratieCreate`: Aansluiting
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenMigratieCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenMigratieCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **migratieRequest** | [**MigratieRequest**](MigratieRequest.md) |  | 

### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenRetrieve

> Aansluiting AansluitingenRetrieve(ctx, aansluitnummer).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenRetrieve(context.Background(), aansluitnummer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenRetrieve`: Aansluiting
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSearchRetrieve

> Search AansluitingenSearchRetrieve(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenSearchRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenSearchRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSearchRetrieve`: Search
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenSearchRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSearchRetrieveRequest struct via the builder pattern


### Return type

[**Search**](Search.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenSectorenRetrieve

> Sector AansluitingenSectorenRetrieve(ctx, aansluitnummer, nummer).Execute()



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
    nummer := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenSectorenRetrieve(context.Background(), aansluitnummer, nummer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenSectorenRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenRetrieve`: Sector
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenSectorenRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 
**nummer** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenRetrieveRequest struct via the builder pattern


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


## AansluitingenSectorenToestandPartialUpdate

> SectorToestand AansluitingenSectorenToestandPartialUpdate(ctx, aansluitnummer).PatchedSectorToestandRequest(patchedSectorToestandRequest).Execute()



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
    patchedSectorToestandRequest := *openapiclient.NewPatchedSectorToestandRequest() // PatchedSectorToestandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenSectorenToestandPartialUpdate(context.Background(), aansluitnummer).PatchedSectorToestandRequest(patchedSectorToestandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenSectorenToestandPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenToestandPartialUpdate`: SectorToestand
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenSectorenToestandPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenToestandPartialUpdateRequest struct via the builder pattern


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


## AansluitingenSectorenToestandPartialUpdate2

> SectorToestand AansluitingenSectorenToestandPartialUpdate2(ctx, aansluitnummer, nummer).PatchedSectorToestandRequest(patchedSectorToestandRequest).Execute()



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
    nummer := int32(56) // int32 | 
    patchedSectorToestandRequest := *openapiclient.NewPatchedSectorToestandRequest() // PatchedSectorToestandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenSectorenToestandPartialUpdate2(context.Background(), aansluitnummer, nummer).PatchedSectorToestandRequest(patchedSectorToestandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenSectorenToestandPartialUpdate2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenSectorenToestandPartialUpdate2`: SectorToestand
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenSectorenToestandPartialUpdate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 
**nummer** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenSectorenToestandPartialUpdate2Request struct via the builder pattern


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


## AansluitingenToestandPartialUpdate

> AansluitingToestand AansluitingenToestandPartialUpdate(ctx, aansluitnummer).PatchedAansluitingToestandRequest(patchedAansluitingToestandRequest).Execute()



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
    patchedAansluitingToestandRequest := *openapiclient.NewPatchedAansluitingToestandRequest() // PatchedAansluitingToestandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenToestandPartialUpdate(context.Background(), aansluitnummer).PatchedAansluitingToestandRequest(patchedAansluitingToestandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenToestandPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenToestandPartialUpdate`: AansluitingToestand
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenToestandPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenToestandPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAansluitingToestandRequest** | [**PatchedAansluitingToestandRequest**](PatchedAansluitingToestandRequest.md) |  | 

### Return type

[**AansluitingToestand**](AansluitingToestand.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenTransitiesPartialUpdate

> Aansluiting AansluitingenTransitiesPartialUpdate(ctx, aansluitnummer).PatchedTransitionRequest(patchedTransitionRequest).Execute()



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
    patchedTransitionRequest := *openapiclient.NewPatchedTransitionRequest() // PatchedTransitionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenTransitiesPartialUpdate(context.Background(), aansluitnummer).PatchedTransitionRequest(patchedTransitionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenTransitiesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenTransitiesPartialUpdate`: Aansluiting
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenTransitiesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenTransitiesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTransitionRequest** | [**PatchedTransitionRequest**](PatchedTransitionRequest.md) |  | 

### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenTransitiesRetrieve

> Transition AansluitingenTransitiesRetrieve(ctx, aansluitnummer).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenTransitiesRetrieve(context.Background(), aansluitnummer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenTransitiesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenTransitiesRetrieve`: Transition
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenTransitiesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenTransitiesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transition**](Transition.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenUpdate

> Aansluiting AansluitingenUpdate(ctx, aansluitnummer).AansluitingRequest(aansluitingRequest).Execute()



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
    aansluitingRequest := *openapiclient.NewAansluitingRequest("Naam_example", "Aansluitnummer_example", int32(123)) // AansluitingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenUpdate(context.Background(), aansluitnummer).AansluitingRequest(aansluitingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenUpdate`: Aansluiting
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aansluitingRequest** | [**AansluitingRequest**](AansluitingRequest.md) |  | 

### Return type

[**Aansluiting**](Aansluiting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AansluitingenValidatieRetrieve

> ValidationObject AansluitingenValidatieRetrieve(ctx, aansluitnummer).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AansluitingenApi.AansluitingenValidatieRetrieve(context.Background(), aansluitnummer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AansluitingenApi.AansluitingenValidatieRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AansluitingenValidatieRetrieve`: ValidationObject
    fmt.Fprintf(os.Stdout, "Response from `AansluitingenApi.AansluitingenValidatieRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aansluitnummer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAansluitingenValidatieRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ValidationObject**](ValidationObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

