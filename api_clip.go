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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// ClipApiService ClipApi service
type ClipApiService service

type ClipApiApiClipMessageRequest struct {
	ctx                          context.Context
	ApiService                   *ClipApiService
	folderId                     string
	postClipFolderMessageRequest *PostClipFolderMessageRequest
}

func (r ClipApiApiClipMessageRequest) PostClipFolderMessageRequest(postClipFolderMessageRequest PostClipFolderMessageRequest) ClipApiApiClipMessageRequest {
	r.postClipFolderMessageRequest = &postClipFolderMessageRequest
	return r
}

func (r ClipApiApiClipMessageRequest) Execute() (*ClippedMessage, *http.Response, error) {
	return r.ApiService.ClipMessageExecute(r)
}

/*
ClipMessage メッセージをクリップフォルダに追加

指定したメッセージを指定したクリップフォルダに追加します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId クリップフォルダUUID
 @return ClipApiApiClipMessageRequest
*/
func (a *ClipApiService) ClipMessage(ctx context.Context, folderId string) ClipApiApiClipMessageRequest {
	return ClipApiApiClipMessageRequest{
		ApiService: a,
		ctx:        ctx,
		folderId:   folderId,
	}
}

// Execute executes the request
//  @return ClippedMessage
func (a *ClipApiService) ClipMessageExecute(r ClipApiApiClipMessageRequest) (*ClippedMessage, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClippedMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClipApiService.ClipMessage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clip-folders/{folderId}/messages"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterToString(r.folderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.postClipFolderMessageRequest
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

type ClipApiApiCreateClipFolderRequest struct {
	ctx                   context.Context
	ApiService            *ClipApiService
	postClipFolderRequest *PostClipFolderRequest
}

func (r ClipApiApiCreateClipFolderRequest) PostClipFolderRequest(postClipFolderRequest PostClipFolderRequest) ClipApiApiCreateClipFolderRequest {
	r.postClipFolderRequest = &postClipFolderRequest
	return r
}

func (r ClipApiApiCreateClipFolderRequest) Execute() (*ClipFolder, *http.Response, error) {
	return r.ApiService.CreateClipFolderExecute(r)
}

/*
CreateClipFolder クリップフォルダを作成

クリップフォルダを作成します。
既にあるフォルダと同名のフォルダを作成することは可能です。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClipApiApiCreateClipFolderRequest
*/
func (a *ClipApiService) CreateClipFolder(ctx context.Context) ClipApiApiCreateClipFolderRequest {
	return ClipApiApiCreateClipFolderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return ClipFolder
func (a *ClipApiService) CreateClipFolderExecute(r ClipApiApiCreateClipFolderRequest) (*ClipFolder, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClipFolder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClipApiService.CreateClipFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clip-folders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.postClipFolderRequest
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

type ClipApiApiDeleteClipFolderRequest struct {
	ctx        context.Context
	ApiService *ClipApiService
	folderId   string
}

func (r ClipApiApiDeleteClipFolderRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClipFolderExecute(r)
}

/*
DeleteClipFolder クリップフォルダを削除

指定したクリップフォルダを削除します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId クリップフォルダUUID
 @return ClipApiApiDeleteClipFolderRequest
*/
func (a *ClipApiService) DeleteClipFolder(ctx context.Context, folderId string) ClipApiApiDeleteClipFolderRequest {
	return ClipApiApiDeleteClipFolderRequest{
		ApiService: a,
		ctx:        ctx,
		folderId:   folderId,
	}
}

// Execute executes the request
func (a *ClipApiService) DeleteClipFolderExecute(r ClipApiApiDeleteClipFolderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClipApiService.DeleteClipFolder")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clip-folders/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterToString(r.folderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ClipApiApiEditClipFolderRequest struct {
	ctx                    context.Context
	ApiService             *ClipApiService
	folderId               string
	patchClipFolderRequest *PatchClipFolderRequest
}

func (r ClipApiApiEditClipFolderRequest) PatchClipFolderRequest(patchClipFolderRequest PatchClipFolderRequest) ClipApiApiEditClipFolderRequest {
	r.patchClipFolderRequest = &patchClipFolderRequest
	return r
}

func (r ClipApiApiEditClipFolderRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditClipFolderExecute(r)
}

/*
EditClipFolder クリップフォルダ情報を編集

指定したクリップフォルダの情報を編集します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId クリップフォルダUUID
 @return ClipApiApiEditClipFolderRequest
*/
func (a *ClipApiService) EditClipFolder(ctx context.Context, folderId string) ClipApiApiEditClipFolderRequest {
	return ClipApiApiEditClipFolderRequest{
		ApiService: a,
		ctx:        ctx,
		folderId:   folderId,
	}
}

// Execute executes the request
func (a *ClipApiService) EditClipFolderExecute(r ClipApiApiEditClipFolderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClipApiService.EditClipFolder")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clip-folders/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterToString(r.folderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.patchClipFolderRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ClipApiApiGetClipFolderRequest struct {
	ctx        context.Context
	ApiService *ClipApiService
	folderId   string
}

func (r ClipApiApiGetClipFolderRequest) Execute() (*ClipFolder, *http.Response, error) {
	return r.ApiService.GetClipFolderExecute(r)
}

/*
GetClipFolder クリップフォルダ情報を取得

指定したクリップフォルダの情報を取得します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId クリップフォルダUUID
 @return ClipApiApiGetClipFolderRequest
*/
func (a *ClipApiService) GetClipFolder(ctx context.Context, folderId string) ClipApiApiGetClipFolderRequest {
	return ClipApiApiGetClipFolderRequest{
		ApiService: a,
		ctx:        ctx,
		folderId:   folderId,
	}
}

// Execute executes the request
//  @return ClipFolder
func (a *ClipApiService) GetClipFolderExecute(r ClipApiApiGetClipFolderRequest) (*ClipFolder, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClipFolder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClipApiService.GetClipFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clip-folders/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterToString(r.folderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ClipApiApiGetClipFoldersRequest struct {
	ctx        context.Context
	ApiService *ClipApiService
}

func (r ClipApiApiGetClipFoldersRequest) Execute() ([]ClipFolder, *http.Response, error) {
	return r.ApiService.GetClipFoldersExecute(r)
}

/*
GetClipFolders クリップフォルダのリストを取得

自身が所有するクリップフォルダのリストを取得します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClipApiApiGetClipFoldersRequest
*/
func (a *ClipApiService) GetClipFolders(ctx context.Context) ClipApiApiGetClipFoldersRequest {
	return ClipApiApiGetClipFoldersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []ClipFolder
func (a *ClipApiService) GetClipFoldersExecute(r ClipApiApiGetClipFoldersRequest) ([]ClipFolder, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ClipFolder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClipApiService.GetClipFolders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clip-folders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ClipApiApiGetClipsRequest struct {
	ctx        context.Context
	ApiService *ClipApiService
	folderId   string
	limit      *int32
	offset     *int32
	order      *string
}

// 取得する件数
func (r ClipApiApiGetClipsRequest) Limit(limit int32) ClipApiApiGetClipsRequest {
	r.limit = &limit
	return r
}

// 取得するオフセット
func (r ClipApiApiGetClipsRequest) Offset(offset int32) ClipApiApiGetClipsRequest {
	r.offset = &offset
	return r
}

// 昇順か降順か
func (r ClipApiApiGetClipsRequest) Order(order string) ClipApiApiGetClipsRequest {
	r.order = &order
	return r
}

func (r ClipApiApiGetClipsRequest) Execute() ([]ClippedMessage, *http.Response, error) {
	return r.ApiService.GetClipsExecute(r)
}

/*
GetClips フォルダ内のクリップのリストを取得

指定したフォルダ内のクリップのリストを取得します。
`order`を指定しない場合、クリップした日時の新しい順で返されます。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId クリップフォルダUUID
 @return ClipApiApiGetClipsRequest
*/
func (a *ClipApiService) GetClips(ctx context.Context, folderId string) ClipApiApiGetClipsRequest {
	return ClipApiApiGetClipsRequest{
		ApiService: a,
		ctx:        ctx,
		folderId:   folderId,
	}
}

// Execute executes the request
//  @return []ClippedMessage
func (a *ClipApiService) GetClipsExecute(r ClipApiApiGetClipsRequest) ([]ClippedMessage, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ClippedMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClipApiService.GetClips")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clip-folders/{folderId}/messages"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterToString(r.folderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
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

type ClipApiApiGetMessageClipsRequest struct {
	ctx        context.Context
	ApiService *ClipApiService
	messageId  string
}

func (r ClipApiApiGetMessageClipsRequest) Execute() ([]MessageClip, *http.Response, error) {
	return r.ApiService.GetMessageClipsExecute(r)
}

/*
GetMessageClips 自分のクリップを取得

対象のメッセージの自分のクリップの一覧を返します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param messageId メッセージUUID
 @return ClipApiApiGetMessageClipsRequest
*/
func (a *ClipApiService) GetMessageClips(ctx context.Context, messageId string) ClipApiApiGetMessageClipsRequest {
	return ClipApiApiGetMessageClipsRequest{
		ApiService: a,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
//  @return []MessageClip
func (a *ClipApiService) GetMessageClipsExecute(r ClipApiApiGetMessageClipsRequest) ([]MessageClip, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []MessageClip
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClipApiService.GetMessageClips")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages/{messageId}/clips"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", url.PathEscape(parameterToString(r.messageId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ClipApiApiUnclipMessageRequest struct {
	ctx        context.Context
	ApiService *ClipApiService
	folderId   string
	messageId  string
}

func (r ClipApiApiUnclipMessageRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnclipMessageExecute(r)
}

/*
UnclipMessage メッセージをクリップフォルダから除外

指定したフォルダから指定したメッセージのクリップを除外します。
既に外されているメッセージを指定した場合は204を返します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId クリップフォルダUUID
 @param messageId メッセージUUID
 @return ClipApiApiUnclipMessageRequest
*/
func (a *ClipApiService) UnclipMessage(ctx context.Context, folderId string, messageId string) ClipApiApiUnclipMessageRequest {
	return ClipApiApiUnclipMessageRequest{
		ApiService: a,
		ctx:        ctx,
		folderId:   folderId,
		messageId:  messageId,
	}
}

// Execute executes the request
func (a *ClipApiService) UnclipMessageExecute(r ClipApiApiUnclipMessageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClipApiService.UnclipMessage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clip-folders/{folderId}/messages/{messageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterToString(r.folderId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", url.PathEscape(parameterToString(r.messageId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
