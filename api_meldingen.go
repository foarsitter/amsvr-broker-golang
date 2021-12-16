/*
 * AMSVR Broker API
 *
 * # Introductie     Voor u heeft u de online documentatie van de AMSVR Broker API. Omdat het een Nederlands domein      betreft worden er Engelse en Nederlandse terminologie door elkaar heen gebruikt.  Liefhebbers van de Swagger UI kunnen [hier terecht](/swagger-ui). Daarnaast is de API eveneens beschikbaar in de op endpoint niveau zoals bijvoorbeeld [hier](/api/aansluitingen).  ##        
 *
 * API version: release-0.6.0.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package broker

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// MeldingenApiService MeldingenApi service
type MeldingenApiService service

// MeldingenListOpts Optional parameters for the method 'MeldingenList'
type MeldingenListOpts struct {
    Actueel optional.Bool
    CreatedAfter optional.Time
    CreatedBefore optional.Time
    Criterium optional.Int32
    Limit optional.Int32
    ModifiedAfter optional.Time
    ModifiedBefore optional.Time
    Offset optional.Int32
    Toestand optional.Int32
}

/*
MeldingenList Method for MeldingenList
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *MeldingenListOpts - Optional Parameters:
 * @param "Actueel" (optional.Bool) - 
 * @param "CreatedAfter" (optional.Time) - 
 * @param "CreatedBefore" (optional.Time) - 
 * @param "Criterium" (optional.Int32) - 
 * @param "Limit" (optional.Int32) -  Number of results to return per page.
 * @param "ModifiedAfter" (optional.Time) - 
 * @param "ModifiedBefore" (optional.Time) - 
 * @param "Offset" (optional.Int32) -  The initial index from which to return the results.
 * @param "Toestand" (optional.Int32) -  0 = Actief, 1 = test, 2 = Passief
@return PaginatedMeldingList
*/
func (a *MeldingenApiService) MeldingenList(ctx _context.Context, localVarOptionals *MeldingenListOpts) (PaginatedMeldingList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedMeldingList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/meldingen"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Actueel.IsSet() {
		localVarQueryParams.Add("actueel", parameterToString(localVarOptionals.Actueel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreatedAfter.IsSet() {
		localVarQueryParams.Add("created_after", parameterToString(localVarOptionals.CreatedAfter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreatedBefore.IsSet() {
		localVarQueryParams.Add("created_before", parameterToString(localVarOptionals.CreatedBefore.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Criterium.IsSet() {
		localVarQueryParams.Add("criterium", parameterToString(localVarOptionals.Criterium.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ModifiedAfter.IsSet() {
		localVarQueryParams.Add("modified_after", parameterToString(localVarOptionals.ModifiedAfter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ModifiedBefore.IsSet() {
		localVarQueryParams.Add("modified_before", parameterToString(localVarOptionals.ModifiedBefore.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Toestand.IsSet() {
		localVarQueryParams.Add("toestand", parameterToString(localVarOptionals.Toestand.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}