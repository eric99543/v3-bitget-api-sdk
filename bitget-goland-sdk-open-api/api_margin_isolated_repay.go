/*
Bitget Open API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// MarginIsolatedRepayApiService MarginIsolatedRepayApi service
type MarginIsolatedRepayApiService service

type ApiIsolateRepayListRequest struct {
	ctx        context.Context
	ApiService *MarginIsolatedRepayApiService
	symbol     *string
	startTime  *string
	coin       *string
	repayId    *string
	endTime    *string
	pageSize   *string
	pageId     *string
}

// symbol
func (r ApiIsolateRepayListRequest) Symbol(symbol string) ApiIsolateRepayListRequest {
	r.symbol = &symbol
	return r
}

// startTime
func (r ApiIsolateRepayListRequest) StartTime(startTime string) ApiIsolateRepayListRequest {
	r.startTime = &startTime
	return r
}

// coin
func (r ApiIsolateRepayListRequest) Coin(coin string) ApiIsolateRepayListRequest {
	r.coin = &coin
	return r
}

// repayId
func (r ApiIsolateRepayListRequest) RepayId(repayId string) ApiIsolateRepayListRequest {
	r.repayId = &repayId
	return r
}

// endTime
func (r ApiIsolateRepayListRequest) EndTime(endTime string) ApiIsolateRepayListRequest {
	r.endTime = &endTime
	return r
}

// pageSize
func (r ApiIsolateRepayListRequest) PageSize(pageSize string) ApiIsolateRepayListRequest {
	r.pageSize = &pageSize
	return r
}

// pageId
func (r ApiIsolateRepayListRequest) PageId(pageId string) ApiIsolateRepayListRequest {
	r.pageId = &pageId
	return r
}

func (r ApiIsolateRepayListRequest) Execute() (*ApiResponseResultOfMarginIsolatedRepayInfoResult, *http.Response, error) {
	return r.ApiService.IsolateRepayListExecute(r)
}

/*
IsolateRepayList list

Get liquidation List

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIsolateRepayListRequest
*/
func (a *MarginIsolatedRepayApiService) IsolateRepayList(ctx context.Context) ApiIsolateRepayListRequest {
	return ApiIsolateRepayListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiResponseResultOfMarginIsolatedRepayInfoResult
func (a *MarginIsolatedRepayApiService) IsolateRepayListExecute(r ApiIsolateRepayListRequest) (*ApiResponseResultOfMarginIsolatedRepayInfoResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiResponseResultOfMarginIsolatedRepayInfoResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarginIsolatedRepayApiService.IsolateRepayList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/margin/v1/isolated/repay/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, reportError("startTime is required and must be specified")
	}

	localVarQueryParams.Add("symbol", parameterToString(*r.symbol, ""))
	if r.coin != nil {
		localVarQueryParams.Add("coin", parameterToString(*r.coin, ""))
	}
	if r.repayId != nil {
		localVarQueryParams.Add("repayId", parameterToString(*r.repayId, ""))
	}
	localVarQueryParams.Add("startTime", parameterToString(*r.startTime, ""))
	if r.endTime != nil {
		localVarQueryParams.Add("endTime", parameterToString(*r.endTime, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.pageId != nil {
		localVarQueryParams.Add("pageId", parameterToString(*r.pageId, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ACCESS_KEY"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ACCESS-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ACCESS_PASSPHRASE"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ACCESS-PASSPHRASE"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ACCESS_SIGN"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ACCESS-SIGN"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ACCESS_TIMESTAMP"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ACCESS-TIMESTAMP"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["SECRET_KEY"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["SECRET-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiResponseResultOfVoid
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiResponseResultOfVoid
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiResponseResultOfVoid
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
