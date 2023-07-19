/*
traQ v3

traQ v3 API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package traq

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// OgpApiService OgpApi service
type OgpApiService service

type OgpApiDeleteOgpCacheRequest struct {
	ctx        context.Context
	ApiService *OgpApiService
	url        *string
}

// OGPのキャッシュを削除したいURL
func (r OgpApiDeleteOgpCacheRequest) Url(url string) OgpApiDeleteOgpCacheRequest {
	r.url = &url
	return r
}

func (r OgpApiDeleteOgpCacheRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOgpCacheExecute(r)
}

/*
DeleteOgpCache OGP情報のキャッシュを削除

指定されたURLのOGP情報のキャッシュを削除します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OgpApiDeleteOgpCacheRequest
*/
func (a *OgpApiService) DeleteOgpCache(ctx context.Context) OgpApiDeleteOgpCacheRequest {
	return OgpApiDeleteOgpCacheRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OgpApiService) DeleteOgpCacheExecute(r OgpApiDeleteOgpCacheRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OgpApiService.DeleteOgpCache")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ogp/cache"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.url == nil {
		return nil, reportError("url is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "url", r.url, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OgpApiGetOgpRequest struct {
	ctx        context.Context
	ApiService *OgpApiService
	url        *string
}

// OGPを取得したいURL
func (r OgpApiGetOgpRequest) Url(url string) OgpApiGetOgpRequest {
	r.url = &url
	return r
}

func (r OgpApiGetOgpRequest) Execute() (*Ogp, *http.Response, error) {
	return r.ApiService.GetOgpExecute(r)
}

/*
GetOgp OGP情報を取得

指定されたURLのOGP情報を取得します。
指定されたURLに対するOGP情報が見つからなかった場合、typeがemptyに設定された空のOGP情報を返します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OgpApiGetOgpRequest
*/
func (a *OgpApiService) GetOgp(ctx context.Context) OgpApiGetOgpRequest {
	return OgpApiGetOgpRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Ogp
func (a *OgpApiService) GetOgpExecute(r OgpApiGetOgpRequest) (*Ogp, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Ogp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OgpApiService.GetOgp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ogp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.url == nil {
		return localVarReturnValue, nil, reportError("url is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "url", r.url, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
