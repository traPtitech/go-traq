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
	"strings"
)

// NotificationApiService NotificationApi service
type NotificationApiService service

type NotificationApiEditChannelSubscribersRequest struct {
	ctx                            context.Context
	ApiService                     *NotificationApiService
	channelId                      string
	patchChannelSubscribersRequest *PatchChannelSubscribersRequest
}

func (r NotificationApiEditChannelSubscribersRequest) PatchChannelSubscribersRequest(patchChannelSubscribersRequest PatchChannelSubscribersRequest) NotificationApiEditChannelSubscribersRequest {
	r.patchChannelSubscribersRequest = &patchChannelSubscribersRequest
	return r
}

func (r NotificationApiEditChannelSubscribersRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditChannelSubscribersExecute(r)
}

/*
EditChannelSubscribers チャンネルの通知購読者を編集

指定したチャンネルの通知購読者を編集します。
リクエストに含めなかったユーザーの通知購読状態は変更しません。
また、存在しないユーザーを指定した場合は無視されます。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId チャンネルUUID
	@return NotificationApiEditChannelSubscribersRequest
*/
func (a *NotificationApiService) EditChannelSubscribers(ctx context.Context, channelId string) NotificationApiEditChannelSubscribersRequest {
	return NotificationApiEditChannelSubscribersRequest{
		ApiService: a,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (a *NotificationApiService) EditChannelSubscribersExecute(r NotificationApiEditChannelSubscribersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.EditChannelSubscribers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channelId}/subscribers"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

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
	localVarPostBody = r.patchChannelSubscribersRequest
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

type NotificationApiGetChannelSubscribersRequest struct {
	ctx        context.Context
	ApiService *NotificationApiService
	channelId  string
}

func (r NotificationApiGetChannelSubscribersRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetChannelSubscribersExecute(r)
}

/*
GetChannelSubscribers チャンネルの通知購読者のリストを取得

指定したチャンネルを通知購読しているユーザーのUUIDのリストを取得します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId チャンネルUUID
	@return NotificationApiGetChannelSubscribersRequest
*/
func (a *NotificationApiService) GetChannelSubscribers(ctx context.Context, channelId string) NotificationApiGetChannelSubscribersRequest {
	return NotificationApiGetChannelSubscribersRequest{
		ApiService: a,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
//
//	@return []string
func (a *NotificationApiService) GetChannelSubscribersExecute(r NotificationApiGetChannelSubscribersRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.GetChannelSubscribers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channelId}/subscribers"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

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

type NotificationApiGetMyChannelSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *NotificationApiService
}

func (r NotificationApiGetMyChannelSubscriptionsRequest) Execute() ([]UserSubscribeState, *http.Response, error) {
	return r.ApiService.GetMyChannelSubscriptionsExecute(r)
}

/*
GetMyChannelSubscriptions 自分のチャンネル購読状態を取得

自身のチャンネル購読状態を取得します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return NotificationApiGetMyChannelSubscriptionsRequest
*/
func (a *NotificationApiService) GetMyChannelSubscriptions(ctx context.Context) NotificationApiGetMyChannelSubscriptionsRequest {
	return NotificationApiGetMyChannelSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []UserSubscribeState
func (a *NotificationApiService) GetMyChannelSubscriptionsExecute(r NotificationApiGetMyChannelSubscriptionsRequest) ([]UserSubscribeState, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []UserSubscribeState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.GetMyChannelSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/subscriptions"

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

type NotificationApiGetMyUnreadChannelsRequest struct {
	ctx        context.Context
	ApiService *NotificationApiService
}

func (r NotificationApiGetMyUnreadChannelsRequest) Execute() ([]UnreadChannel, *http.Response, error) {
	return r.ApiService.GetMyUnreadChannelsExecute(r)
}

/*
GetMyUnreadChannels 未読チャンネルを取得

自分が現在未読のチャンネルの未読情報を取得します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return NotificationApiGetMyUnreadChannelsRequest
*/
func (a *NotificationApiService) GetMyUnreadChannels(ctx context.Context) NotificationApiGetMyUnreadChannelsRequest {
	return NotificationApiGetMyUnreadChannelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []UnreadChannel
func (a *NotificationApiService) GetMyUnreadChannelsExecute(r NotificationApiGetMyUnreadChannelsRequest) ([]UnreadChannel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []UnreadChannel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.GetMyUnreadChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/unread"

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

type NotificationApiGetMyViewStatesRequest struct {
	ctx        context.Context
	ApiService *NotificationApiService
}

func (r NotificationApiGetMyViewStatesRequest) Execute() ([]MyChannelViewState, *http.Response, error) {
	return r.ApiService.GetMyViewStatesExecute(r)
}

/*
GetMyViewStates 自身のチャンネル閲覧状態一覧を取得

自身のチャンネル閲覧状態一覧を取得します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return NotificationApiGetMyViewStatesRequest
*/
func (a *NotificationApiService) GetMyViewStates(ctx context.Context) NotificationApiGetMyViewStatesRequest {
	return NotificationApiGetMyViewStatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []MyChannelViewState
func (a *NotificationApiService) GetMyViewStatesExecute(r NotificationApiGetMyViewStatesRequest) ([]MyChannelViewState, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []MyChannelViewState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.GetMyViewStates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/view-states"

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

type NotificationApiReadChannelRequest struct {
	ctx        context.Context
	ApiService *NotificationApiService
	channelId  string
}

func (r NotificationApiReadChannelRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReadChannelExecute(r)
}

/*
ReadChannel チャンネルを既読にする

自分が未読のチャンネルを既読にします。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId チャンネルUUID
	@return NotificationApiReadChannelRequest
*/
func (a *NotificationApiService) ReadChannel(ctx context.Context, channelId string) NotificationApiReadChannelRequest {
	return NotificationApiReadChannelRequest{
		ApiService: a,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (a *NotificationApiService) ReadChannelExecute(r NotificationApiReadChannelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.ReadChannel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/unread/{channelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

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

type NotificationApiRegisterFCMDeviceRequest struct {
	ctx                    context.Context
	ApiService             *NotificationApiService
	postMyFCMDeviceRequest *PostMyFCMDeviceRequest
}

func (r NotificationApiRegisterFCMDeviceRequest) PostMyFCMDeviceRequest(postMyFCMDeviceRequest PostMyFCMDeviceRequest) NotificationApiRegisterFCMDeviceRequest {
	r.postMyFCMDeviceRequest = &postMyFCMDeviceRequest
	return r
}

func (r NotificationApiRegisterFCMDeviceRequest) Execute() (*http.Response, error) {
	return r.ApiService.RegisterFCMDeviceExecute(r)
}

/*
RegisterFCMDevice FCMデバイスを登録

自身のFCMデバイスを登録します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return NotificationApiRegisterFCMDeviceRequest
*/
func (a *NotificationApiService) RegisterFCMDevice(ctx context.Context) NotificationApiRegisterFCMDeviceRequest {
	return NotificationApiRegisterFCMDeviceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *NotificationApiService) RegisterFCMDeviceExecute(r NotificationApiRegisterFCMDeviceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.RegisterFCMDevice")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/fcm-device"

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
	localVarPostBody = r.postMyFCMDeviceRequest
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

type NotificationApiSetChannelSubscribeLevelRequest struct {
	ctx                             context.Context
	ApiService                      *NotificationApiService
	channelId                       string
	putChannelSubscribeLevelRequest *PutChannelSubscribeLevelRequest
}

func (r NotificationApiSetChannelSubscribeLevelRequest) PutChannelSubscribeLevelRequest(putChannelSubscribeLevelRequest PutChannelSubscribeLevelRequest) NotificationApiSetChannelSubscribeLevelRequest {
	r.putChannelSubscribeLevelRequest = &putChannelSubscribeLevelRequest
	return r
}

func (r NotificationApiSetChannelSubscribeLevelRequest) Execute() (*http.Response, error) {
	return r.ApiService.SetChannelSubscribeLevelExecute(r)
}

/*
SetChannelSubscribeLevel チャンネル購読レベルを設定

自身の指定したチャンネルの購読レベルを設定します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId チャンネルUUID
	@return NotificationApiSetChannelSubscribeLevelRequest
*/
func (a *NotificationApiService) SetChannelSubscribeLevel(ctx context.Context, channelId string) NotificationApiSetChannelSubscribeLevelRequest {
	return NotificationApiSetChannelSubscribeLevelRequest{
		ApiService: a,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (a *NotificationApiService) SetChannelSubscribeLevelExecute(r NotificationApiSetChannelSubscribeLevelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.SetChannelSubscribeLevel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/subscriptions/{channelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

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
	localVarPostBody = r.putChannelSubscribeLevelRequest
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

type NotificationApiSetChannelSubscribersRequest struct {
	ctx                          context.Context
	ApiService                   *NotificationApiService
	channelId                    string
	putChannelSubscribersRequest *PutChannelSubscribersRequest
}

func (r NotificationApiSetChannelSubscribersRequest) PutChannelSubscribersRequest(putChannelSubscribersRequest PutChannelSubscribersRequest) NotificationApiSetChannelSubscribersRequest {
	r.putChannelSubscribersRequest = &putChannelSubscribersRequest
	return r
}

func (r NotificationApiSetChannelSubscribersRequest) Execute() (*http.Response, error) {
	return r.ApiService.SetChannelSubscribersExecute(r)
}

/*
SetChannelSubscribers チャンネルの通知購読者を設定

指定したチャンネルの通知購読者を設定します。
リクエストに含めなかったユーザーの通知購読状態はオフになります。
また、存在しないユーザーを指定した場合は無視されます。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId チャンネルUUID
	@return NotificationApiSetChannelSubscribersRequest
*/
func (a *NotificationApiService) SetChannelSubscribers(ctx context.Context, channelId string) NotificationApiSetChannelSubscribersRequest {
	return NotificationApiSetChannelSubscribersRequest{
		ApiService: a,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (a *NotificationApiService) SetChannelSubscribersExecute(r NotificationApiSetChannelSubscribersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.SetChannelSubscribers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channelId}/subscribers"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

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
	localVarPostBody = r.putChannelSubscribersRequest
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

type NotificationApiWsRequest struct {
	ctx        context.Context
	ApiService *NotificationApiService
}

func (r NotificationApiWsRequest) Execute() (*http.Response, error) {
	return r.ApiService.WsExecute(r)
}

/*
Ws WebSocket通知ストリームに接続します

# WebSocketプロトコル
## 送信
`コマンド:引数1:引数2:...`のような形式のTextMessageをサーバーに送信することで、このWebSocketセッションに対する設定が実行できる。
### `viewstate`コマンド
このWebSocketセッションが見ているチャンネル(イベントを受け取るチャンネル)を設定する。
現時点では1つのセッションに対して1つのチャンネルしか設定できない。

`viewstate:{チャンネルID}:{閲覧状態}`
+ チャンネルID: 対象のチャンネルID
+ 閲覧状態: `none`, `monitoring`, `editing`

最初の`viewstate`コマンドを送る前、または`viewstate:null`, `viewstate:`を送信した後は、このセッションはどこのチャンネルも見ていないことになる。

### `rtcstate`コマンド
自分のWebRTC状態を変更する。
他のコネクションが既に状態を保持している場合、変更することができません。

`rtcstate:{チャンネルID}:({状態}:{セッションID})*`

コネクションが切断された場合、自分のWebRTC状態はリセットされます。

### `timeline_streaming`コマンド
全てのパブリックチャンネルの`MESSAGE_CREATED`イベントを受け取るかどうかを設定する。
初期状態は`off`です。

`timeline_streaming:(on|off|true|false)`

## 受信
TextMessageとして各種イベントが`type`と`body`を持つJSONとして非同期に送られます。

例:
```json
{"type":"USER_ONLINE","body":{"id":"7dd8e07f-7f5d-4331-9176-b56a4299768b"}}
```

## イベント一覧

### `USER_JOINED`
ユーザーが新規登録された。

対象: 全員

+ `id`: 登録されたユーザーのId

### `USER_UPDATED`
ユーザーの情報が更新された。

対象: 全員

+ `id`: 情報が更新されたユーザーのId

### `USER_TAGS_UPDATED`
ユーザーのタグが更新された。

対象: 全員

+ `id`: タグが更新されたユーザーのId
+ `tag_id`: 更新されたタグのId

### `USER_ICON_UPDATED`
ユーザーのアイコンが更新された。

対象: 全員

+ `id`: アイコンが更新されたユーザーのId

### `USER_WEBRTC_STATE_CHANGED`
ユーザーのWebRTCの状態が変化した

対象: 全員

+ `user_id`: 変更があったユーザーのId
+ `channel_id`: ユーザーの変更後の接続チャンネルのId
+ `sessions`: ユーザーの変更後の状態(配列)
  - `state`: 状態
  - `sessionId`: セッションID

### `USER_VIEWSTATE_CHANGED`
ユーザーのチャンネルの閲覧状態が変化した

対象: 変化したWSセッションを含めた、該当ユーザーのWSセッション全て

+ `view_states`: 変化したWSセッションを含めた、該当ユーザーの変更後の状態(配列)
  - `key`: WSセッションの識別子
  - `channel_id`: 閲覧しているチャンネルId
  - `state`: 閲覧状態

### `USER_ONLINE`
ユーザーがオンラインになった。

対象: 全員

+ `id`: オンラインになったユーザーのId

### `USER_OFFLINE`
ユーザーがオフラインになった。

対象: 全員

+ `id`: オフラインになったユーザーのId

### `USER_GROUP_CREATED`
ユーザーグループが作成された

対象: 全員

+ `id`: 作成されたユーザーグループのId

### `USER_GROUP_UPDATED`
ユーザーグループが更新された

対象: 全員

+ `id`: 作成されたユーザーグループのId

### `USER_GROUP_DELETED`
ユーザーグループが削除された

対象: 全員

+ `id`: 削除されたユーザーグループのId

### `CHANNEL_CREATED`
チャンネルが新規作成された。

対象: 該当チャンネルを閲覧可能な全員

+ `id`: 作成されたチャンネルのId
+ `dm_user_id`: (DMの場合のみ) DM相手のユーザーId

### `CHANNEL_UPDATED`
チャンネルの情報が変更された。

対象: 該当チャンネルを閲覧可能な全員

+ `id`: 変更があったチャンネルのId
+ `dm_user_id`: (DMの場合のみ) DM相手のユーザーId

### `CHANNEL_DELETED`
チャンネルが削除された。

対象: 該当チャンネルを閲覧可能な全員

+ `id`: 削除されたチャンネルのId
+ `dm_user_id`: (DMの場合のみ) DM相手のユーザーId

### `CHANNEL_STARED`
自分がチャンネルをスターした。

対象: 自分

+ `id`: スターしたチャンネルのId

### `CHANNEL_UNSTARED`
自分がチャンネルのスターを解除した。

対象: 自分

+ `id`: スターしたチャンネルのId

### `CHANNEL_VIEWERS_CHANGED`
チャンネルの閲覧者が変化した。

対象: 該当チャンネルを閲覧しているユーザー

+ `id`: 変化したチャンネルのId
+ `viewers`: 変化後の閲覧者(配列)
  - `userId`: ユーザーId
  - `state`: 閲覧状態
  - `updatedAt`: 閲覧状態の更新日時

### `CHANNEL_SUBSCRIBERS_CHANGED`
チャンネルの購読者が変化した。

対象: 該当チャンネルを閲覧しているユーザー

+ `id`: 変化したチャンネルのId

### `MESSAGE_CREATED`
メッセージが投稿された。

対象: 投稿チャンネルを閲覧しているユーザー・投稿チャンネルに通知をつけているユーザー・メンションを受けたユーザー

+ `id`: 投稿されたメッセージのId
+ `is_citing`: 投稿されたメッセージがWebSocketを接続しているユーザーの投稿を引用しているかどうか

### `MESSAGE_UPDATED`
メッセージが更新された。

対象: 投稿チャンネルを閲覧しているユーザー

+ `id`: 更新されたメッセージのId

### `MESSAGE_DELETED`
メッセージが削除された。

対象: 投稿チャンネルを閲覧しているユーザー

+ `id`: 削除されたメッセージのId

### `MESSAGE_STAMPED`
メッセージにスタンプが押された。

対象: 投稿チャンネルを閲覧しているユーザー

+ `message_id`: メッセージId
+ `user_id`: スタンプを押したユーザーのId
+ `stamp_id`: スタンプのId
+ `count`: そのユーザーが押した数
+ `created_at`: そのユーザーがそのスタンプをそのメッセージに最初に押した日時

### `MESSAGE_UNSTAMPED`
メッセージからスタンプが外された。

対象: 投稿チャンネルを閲覧しているユーザー

+ `message_id`: メッセージId
+ `user_id`: スタンプを押したユーザーのId
+ `stamp_id`: スタンプのId

### `MESSAGE_PINNED`
メッセージがピン留めされた。

対象: 投稿チャンネルを閲覧しているユーザー

+ `message_id`: ピンされたメッセージのID
+ `channel_id`: ピンされたメッセージのチャンネルID

### `MESSAGE_UNPINNED`
ピン留めされたメッセージのピンが外された。

対象: 投稿チャンネルを閲覧しているユーザー

+ `message_id`: ピンが外されたメッセージのID
+ `channel_id`: ピンが外されたメッセージのチャンネルID

### `MESSAGE_READ`
自分があるチャンネルのメッセージを読んだ。

対象: 自分

+ `id`: 読んだチャンネルId

### `STAMP_CREATED`
スタンプが新しく追加された。

対象: 全員

+ `id`: 作成されたスタンプのId

### `STAMP_UPDATED`
スタンプが修正された。

対象: 全員

+ `id`: 修正されたスタンプのId

### `STAMP_DELETED`
スタンプが削除された。

対象: 全員

+ `id`: 削除されたスタンプのId

### `STAMP_PALETTE_CREATED`
スタンプパレットが新しく追加された。

対象: 自分

+ `id`: 作成されたスタンプパレットのId

### `STAMP_PALETTE_UPDATED`
スタンプパレットが修正された。

対象: 自分

+ `id`: 修正されたスタンプパレットのId

### `STAMP_PALETTE_DELETED`
スタンプパレットが削除された。

対象: 自分

+ `id`: 削除されたスタンプパレットのId

### `CLIP_FOLDER_CREATED`
クリップフォルダーが作成された。

対象：自分

+ `id`: 作成されたクリップフォルダーのId

### `CLIP_FOLDER_UPDATED`
クリップフォルダーが修正された。

対象: 自分

+ `id`: 更新されたクリップフォルダーのId

### `CLIP_FOLDER_DELETED`
クリップフォルダーが削除された。

対象: 自分

+ `id`: 削除されたクリップフォルダーのId

### `CLIP_FOLDER_MESSAGE_DELETED`
クリップフォルダーからメッセージが除外された。

対象: 自分

+ `folder_id`: メッセージが除外されたクリップフォルダーのId
+ `message_id`: クリップフォルダーから除外されたメッセージのId

### `CLIP_FOLDER_MESSAGE_ADDED`
クリップフォルダーにメッセージが追加された。

対象: 自分

+ `folder_id`: メッセージが追加されたクリップフォルダーのId
+ `message_id`: クリップフォルダーに追加されたメッセージのId

### `QALL_ROOM_STATE_CHANGED`
ルーム状態が変更された。

対象: 全員

+ `room_id`: 変更されたルームのId
+ `state`: 変更後のルーム状態
  - `roomId`: ルームのID
  - `participants`: ルーム内の参加者(配列)
  - `identity`: ユーザーID_RandomUUID
  - `name`: 表示名
  - `joinedAt`: 参加した時刻
  - `attributes`: ユーザーに関連付けられたカスタム属性
  - `canPublish`: 発言権限
  - `isWebinar`: ウェビナールームかどうか
  - `metadata`: ルームに関連付けられたカスタム属性

### `QALL_SOUNDBOARD_ITEM_CREATED`
サウンドボードアイテムが作成された。

対象: 全員

+ `sound_id`: 作成されたサウンドのId
+ `name`: サウンド名
+ `creator_id`: 作成者のId

### `QALL_SOUNDBOARD_ITEM_DELETED`
サウンドボードアイテムが削除された。

対象: 全員

+ `sound_id`: 削除されたサウンドのId

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return NotificationApiWsRequest
*/
func (a *NotificationApiService) Ws(ctx context.Context) NotificationApiWsRequest {
	return NotificationApiWsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *NotificationApiService) WsExecute(r NotificationApiWsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.Ws")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ws"

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
