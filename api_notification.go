/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// NotificationApiService NotificationApi service
type NotificationApiService service

// NotificationApiEditChannelSubscribersOpts Optional parameters for the method 'EditChannelSubscribers'
type NotificationApiEditChannelSubscribersOpts struct {
	PatchChannelSubscribersRequest optional.Interface
}

/*
EditChannelSubscribers チャンネルの通知購読者を編集
指定したチャンネルの通知購読者を編集します。 リクエストに含めなかったユーザーの通知購読状態は変更しません。 また、存在しないユーザーを指定した場合は無視されます。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId チャンネルUUID
 * @param optional nil or *NotificationApiEditChannelSubscribersOpts - Optional Parameters:
 * @param "PatchChannelSubscribersRequest" (optional.Interface of PatchChannelSubscribersRequest) -
*/
func (a *NotificationApiService) EditChannelSubscribers(ctx _context.Context, channelId string, localVarOptionals *NotificationApiEditChannelSubscribersOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/subscribers"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", _neturl.QueryEscape(parameterToString(channelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if localVarOptionals != nil && localVarOptionals.PatchChannelSubscribersRequest.IsSet() {
		localVarOptionalPatchChannelSubscribersRequest, localVarOptionalPatchChannelSubscribersRequestok := localVarOptionals.PatchChannelSubscribersRequest.Value().(PatchChannelSubscribersRequest)
		if !localVarOptionalPatchChannelSubscribersRequestok {
			return nil, reportError("patchChannelSubscribersRequest should be PatchChannelSubscribersRequest")
		}
		localVarPostBody = &localVarOptionalPatchChannelSubscribersRequest
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
GetChannelSubscribers チャンネルの通知購読者のリストを取得
指定したチャンネルを通知購読しているユーザーのUUIDのリストを取得します。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId チャンネルUUID
@return []string
*/
func (a *NotificationApiService) GetChannelSubscribers(ctx _context.Context, channelId string) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/subscribers"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", _neturl.QueryEscape(parameterToString(channelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
		if localVarHTTPResponse.StatusCode == 200 {
			var v []string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

/*
GetMyChannelSubscriptions 自分のチャンネル購読状態を取得
自身のチャンネル購読状態を取得します。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return []UserSubscribeState
*/
func (a *NotificationApiService) GetMyChannelSubscriptions(ctx _context.Context) ([]UserSubscribeState, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []UserSubscribeState
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/me/subscriptions"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
		if localVarHTTPResponse.StatusCode == 200 {
			var v []UserSubscribeState
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

/*
GetMyUnreadChannels 未読チャンネルを取得
自分が現在未読のチャンネルの未読情報を取得します。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return []UnreadChannel
*/
func (a *NotificationApiService) GetMyUnreadChannels(ctx _context.Context) ([]UnreadChannel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []UnreadChannel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/me/unread"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
		if localVarHTTPResponse.StatusCode == 200 {
			var v []UnreadChannel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

/*
ReadChannel チャンネルを既読にする
自分が未読のチャンネルを既読にします。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId チャンネルUUID
*/
func (a *NotificationApiService) ReadChannel(ctx _context.Context, channelId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/me/unread/{channelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", _neturl.QueryEscape(parameterToString(channelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// NotificationApiRegisterFCMDeviceOpts Optional parameters for the method 'RegisterFCMDevice'
type NotificationApiRegisterFCMDeviceOpts struct {
	PostMyFcmDeviceRequest optional.Interface
}

/*
RegisterFCMDevice FCMデバイスを登録
自身のFCMデバイスを登録します。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *NotificationApiRegisterFCMDeviceOpts - Optional Parameters:
 * @param "PostMyFcmDeviceRequest" (optional.Interface of PostMyFcmDeviceRequest) -
*/
func (a *NotificationApiService) RegisterFCMDevice(ctx _context.Context, localVarOptionals *NotificationApiRegisterFCMDeviceOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/me/fcm-device"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if localVarOptionals != nil && localVarOptionals.PostMyFcmDeviceRequest.IsSet() {
		localVarOptionalPostMyFcmDeviceRequest, localVarOptionalPostMyFcmDeviceRequestok := localVarOptionals.PostMyFcmDeviceRequest.Value().(PostMyFcmDeviceRequest)
		if !localVarOptionalPostMyFcmDeviceRequestok {
			return nil, reportError("postMyFcmDeviceRequest should be PostMyFcmDeviceRequest")
		}
		localVarPostBody = &localVarOptionalPostMyFcmDeviceRequest
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// NotificationApiSetChannelSubscribeLevelOpts Optional parameters for the method 'SetChannelSubscribeLevel'
type NotificationApiSetChannelSubscribeLevelOpts struct {
	PutChannelSubscribeLevelRequest optional.Interface
}

/*
SetChannelSubscribeLevel チャンネル購読レベルを設定
自身の指定したチャンネルの購読レベルを設定します。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId チャンネルUUID
 * @param optional nil or *NotificationApiSetChannelSubscribeLevelOpts - Optional Parameters:
 * @param "PutChannelSubscribeLevelRequest" (optional.Interface of PutChannelSubscribeLevelRequest) -
*/
func (a *NotificationApiService) SetChannelSubscribeLevel(ctx _context.Context, channelId string, localVarOptionals *NotificationApiSetChannelSubscribeLevelOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/me/subscriptions/{channelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", _neturl.QueryEscape(parameterToString(channelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if localVarOptionals != nil && localVarOptionals.PutChannelSubscribeLevelRequest.IsSet() {
		localVarOptionalPutChannelSubscribeLevelRequest, localVarOptionalPutChannelSubscribeLevelRequestok := localVarOptionals.PutChannelSubscribeLevelRequest.Value().(PutChannelSubscribeLevelRequest)
		if !localVarOptionalPutChannelSubscribeLevelRequestok {
			return nil, reportError("putChannelSubscribeLevelRequest should be PutChannelSubscribeLevelRequest")
		}
		localVarPostBody = &localVarOptionalPutChannelSubscribeLevelRequest
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// NotificationApiSetChannelSubscribersOpts Optional parameters for the method 'SetChannelSubscribers'
type NotificationApiSetChannelSubscribersOpts struct {
	PutChannelSubscribersRequest optional.Interface
}

/*
SetChannelSubscribers チャンネルの通知購読者を設定
指定したチャンネルの通知購読者を設定します。 リクエストに含めなかったユーザーの通知購読状態はオフになります。 また、存在しないユーザーを指定した場合は無視されます。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId チャンネルUUID
 * @param optional nil or *NotificationApiSetChannelSubscribersOpts - Optional Parameters:
 * @param "PutChannelSubscribersRequest" (optional.Interface of PutChannelSubscribersRequest) -
*/
func (a *NotificationApiService) SetChannelSubscribers(ctx _context.Context, channelId string, localVarOptionals *NotificationApiSetChannelSubscribersOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/subscribers"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", _neturl.QueryEscape(parameterToString(channelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if localVarOptionals != nil && localVarOptionals.PutChannelSubscribersRequest.IsSet() {
		localVarOptionalPutChannelSubscribersRequest, localVarOptionalPutChannelSubscribersRequestok := localVarOptionals.PutChannelSubscribersRequest.Value().(PutChannelSubscribersRequest)
		if !localVarOptionalPutChannelSubscribersRequestok {
			return nil, reportError("putChannelSubscribersRequest should be PutChannelSubscribersRequest")
		}
		localVarPostBody = &localVarOptionalPutChannelSubscribersRequest
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
Ws WebSocket通知ストリームに接続します
# WebSocketプロトコル ## 送信 &#x60;コマンド:引数1:引数2:...&#x60;のような形式のTextMessageをサーバーに送信することで、このWebSocketセッションに対する設定が実行できる。 ### &#x60;viewstate&#x60;コマンド このWebSocketセッションが見ているチャンネル(イベントを受け取るチャンネル)を設定する。 現時点では1つのセッションに対して1つのチャンネルしか設定できない。  &#x60;viewstate:{チャンネルID}:{閲覧状態}&#x60; + チャンネルID: 対象のチャンネルID + 閲覧状態: &#x60;none&#x60;, &#x60;monitoring&#x60;, &#x60;editing&#x60;  最初の&#x60;viewstate&#x60;コマンドを送る前、または&#x60;viewstate:null&#x60;, &#x60;viewstate:&#x60;を送信した後は、このセッションはどこのチャンネルも見ていないことになる。  ### &#x60;rtcstate&#x60;コマンド 自分のWebRTC状態を変更する。 他のコネクションが既に状態を保持している場合、変更することができません。  &#x60;rtcstate:{チャンネルID}:({状態}:{セッションID})*&#x60;  コネクションが切断された場合、自分のWebRTC状態はリセットされます。  ### &#x60;timeline_streaming&#x60;コマンド 全てのパブリックチャンネルの&#x60;MESSAGE_CREATED&#x60;イベントを受け取るかどうかを設定する。 初期状態は&#x60;off&#x60;です。  &#x60;timeline_streaming:(on|off|true|false)&#x60;  ## 受信 TextMessageとして各種イベントが&#x60;type&#x60;と&#x60;body&#x60;を持つJSONとして非同期に送られます。  例:  &#x60;&#x60;&#x60;json {\&quot;type\&quot;:\&quot;USER_ONLINE\&quot;,\&quot;body\&quot;:{\&quot;id\&quot;:\&quot;7dd8e07f-7f5d-4331-9176-b56a4299768b\&quot;}} &#x60;&#x60;&#x60;  ## イベント一覧  ### &#x60;USER_JOINED&#x60; ユーザーが新規登録された。  対象: 全員  + &#x60;id&#x60;: 登録されたユーザーのId  ### &#x60;USER_UPDATED&#x60; ユーザーの情報が更新された。  対象: 全員  + &#x60;id&#x60;: 情報が更新されたユーザーのId  ### &#x60;USER_TAGS_UPDATED&#x60; ユーザーのタグが更新された。  対象: 全員  + &#x60;id&#x60;: タグが更新されたユーザーのId  ### &#x60;USER_ICON_UPDATED&#x60; ユーザーのアイコンが更新された。  対象: 全員  + &#x60;id&#x60;: アイコンが更新されたユーザーのId  ### &#x60;USER_WEBRTC_STATE_CHANGED&#x60; ユーザーのWebRTCの状態が変化した  対象: 全員  + &#x60;user_id&#x60;: 変更があったユーザーのId + &#x60;channel_id&#x60;: ユーザーの変更後の接続チャンネルのId + &#x60;sessions&#x60;: ユーザーの変更後の状態(配列)   + &#x60;state&#x60;: 状態   + &#x60;sessionId&#x60;: セッションID  ### &#x60;USER_ONLINE&#x60; ユーザーがオンラインになった。  対象: 全員  + &#x60;id&#x60;: オンラインになったユーザーのId  ### &#x60;USER_OFFLINE&#x60; ユーザーがオフラインになった。  対象: 全員  + &#x60;id&#x60;: オフラインになったユーザーのId  ### &#x60;USER_GROUP_CREATED&#x60; ユーザーグループが作成された  対象: 全員  + &#x60;id&#x60;: 作成されたユーザーグループのId  ### &#x60;USER_GROUP_UPDATED&#x60; ユーザーグループが更新された  対象: 全員  + &#x60;id&#x60;: 作成されたユーザーグループのId  ### &#x60;USER_GROUP_DELETED&#x60; ユーザーグループが削除された  対象: 全員  + &#x60;id&#x60;: 削除されたユーザーグループのId  ### &#x60;CHANNEL_CREATED&#x60; チャンネルが新規作成された。  対象: 全員  + &#x60;id&#x60;: 作成されたチャンネルのId  ### &#x60;CHANNEL_UPDATED&#x60; チャンネルの情報が変更された。  対象: 全員  + &#x60;id&#x60;: 変更があったチャンネルのId  ### &#x60;CHANNEL_DELETED&#x60; チャンネルが削除された。  対象: 全員  + &#x60;id&#x60;: 削除されたチャンネルのId  ### &#x60;CHANNEL_STARED&#x60; 自分がチャンネルをスターした。  対象: 自分  + &#x60;id&#x60;: スターしたチャンネルのId  ### &#x60;CHANNEL_UNSTARED&#x60; 自分がチャンネルのスターを解除した。  対象: 自分  + &#x60;id&#x60;: スターしたチャンネルのId  ### &#x60;MESSAGE_CREATED&#x60; メッセージが投稿された。  対象: 投稿チャンネルにハートビートを送信しているユーザー・投稿チャンネルに通知をつけているユーザー・メンションを受けたユーザー  + &#x60;id&#x60;: 投稿されたメッセージのId  ### &#x60;MESSAGE_UPDATED&#x60; メッセージが更新された。  対象: 投稿チャンネルにハートビートを送信しているユーザー  + &#x60;id&#x60;: 更新されたメッセージのId  ### &#x60;MESSAGE_DELETED&#x60; メッセージが削除された。  対象: 投稿チャンネルにハートビートを送信しているユーザー  + &#x60;id&#x60;: 削除されたメッセージのId  ### &#x60;MESSAGE_STAMPED&#x60; メッセージにスタンプが押された。  対象: 投稿チャンネルにハートビートを送信しているユーザー  + &#x60;message_id&#x60;: メッセージId + &#x60;user_id&#x60;: スタンプを押したユーザーのId + &#x60;stamp_id&#x60;: スタンプのId + &#x60;count&#x60;: そのユーザーが押した数 + &#x60;created_at&#x60;: そのユーザーがそのスタンプをそのメッセージに最初に押した日時  ### &#x60;MESSAGE_UNSTAMPED&#x60; メッセージからスタンプが外された。  対象: 投稿チャンネルにハートビートを送信しているユーザー  + &#x60;message_id&#x60;: メッセージId + &#x60;user_id&#x60;: スタンプを押したユーザーのId + &#x60;stamp_id&#x60;: スタンプのId  ### &#x60;MESSAGE_PINNED&#x60; メッセージがピン留めされた。  対象: 投稿チャンネルにハートビートを送信しているユーザー  + &#x60;message_id&#x60;: ピンされたメッセージのID + &#x60;channel_id&#x60;: ピンされたメッセージのチャンネルID  ### &#x60;MESSAGE_UNPINNED&#x60; ピン留めされたメッセージのピンが外された。  対象: 投稿チャンネルにハートビートを送信しているユーザー  + &#x60;message_id&#x60;: ピンが外されたメッセージのID + &#x60;channel_id&#x60;: ピンが外されたメッセージのチャンネルID  ### &#x60;MESSAGE_READ&#x60; 自分があるチャンネルのメッセージを読んだ。  対象: 自分  + &#x60;id&#x60;: 読んだチャンネルId  ### &#x60;STAMP_CREATED&#x60; スタンプが新しく追加された。  対象: 全員  + &#x60;id&#x60;: 作成されたスタンプのId  ### &#x60;STAMP_UPDATED&#x60; スタンプが修正された。  対象: 全員  + &#x60;id&#x60;: 修正されたスタンプのId  ### &#x60;STAMP_DELETED&#x60; スタンプが削除された。  対象: 全員  + &#x60;id&#x60;: 削除されたスタンプのId  ### &#x60;STAMP_PALETTE_CREATED&#x60; スタンプパレットが新しく追加された。  対象: 自分  + &#x60;id&#x60;: 作成されたスタンプパレットのId  ### &#x60;STAMP_PALETTE_UPDATED&#x60; スタンプパレットが修正された。  対象: 自分  + &#x60;id&#x60;: 修正されたスタンプパレットのId  ### &#x60;STAMP_PALETTE_DELETED&#x60; スタンプパレットが削除された。  対象: 自分  + &#x60;id&#x60;: 削除されたスタンプパレットのId  ### &#x60;CLIP_FOLDER_CREATED&#x60; クリップフォルダーが作成された。  対象：自分  + &#x60;id&#x60;: 作成されたクリップフォルダーのId  ### &#x60;CLIP_FOLDER_UPDATED&#x60; クリップフォルダーが修正された。  対象: 自分  + &#x60;id&#x60;: 更新されたクリップフォルダーのId  ### &#x60;CLIP_FOLDER_DELETED&#x60; クリップフォルダーが削除された。  対象: 自分  + &#x60;id&#x60;: 削除されたクリップフォルダーのId  ### &#x60;CLIP_FOLDER_MESSAGE_DELETED&#x60; クリップフォルダーからメッセージが除外された。  対象: 自分  + &#x60;folder_id&#x60;: メッセージが除外されたクリップフォルダーのId + &#x60;message_id&#x60;: クリップフォルダーから除外されたメッセージのId  ### &#x60;CLIP_FOLDER_MESSAGE_ADDED&#x60; クリップフォルダーにメッセージが追加された。  対象: 自分  + &#x60;folder_id&#x60;: メッセージが追加されたクリップフォルダーのId + &#x60;message_id&#x60;: クリップフォルダーに追加されたメッセージのId
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
*/
func (a *NotificationApiService) Ws(ctx _context.Context) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ws"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
