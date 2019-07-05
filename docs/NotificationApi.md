# \NotificationApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsChannelIDNotificationGet**](NotificationApi.md#ChannelsChannelIDNotificationGet) | **Get** /channels/{channelID}/notification | 
[**ChannelsChannelIDNotificationPut**](NotificationApi.md#ChannelsChannelIDNotificationPut) | **Put** /channels/{channelID}/notification | 
[**NotificationDevicePost**](NotificationApi.md#NotificationDevicePost) | **Post** /notification/device | 
[**NotificationGet**](NotificationApi.md#NotificationGet) | **Get** /notification | 
[**UsersMeNotificationGet**](NotificationApi.md#UsersMeNotificationGet) | **Get** /users/me/notification | 
[**UsersUserIDNotificationGet**](NotificationApi.md#UsersUserIDNotificationGet) | **Get** /users/{userID}/notification | 



## ChannelsChannelIDNotificationGet

> []string ChannelsChannelIDNotificationGet(ctx, channelID)


通知を点けているユーザーのIDの配列を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 

### Return type

**[]string**

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelIDNotificationPut

> ChannelsChannelIDNotificationPut(ctx, channelID, optional)


チャンネルの通知状況を変更します。 リクエストに含めなかったユーザーIDのユーザーの通知状況は変更しません。 また、存在しないユーザーのIDを指定した場合は無視されます。 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 
 **optional** | ***ChannelsChannelIDNotificationPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChannelsChannelIDNotificationPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject21** | [**optional.Interface of InlineObject21**](InlineObject21.md)|  | 

### Return type

 (empty response body)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationDevicePost

> NotificationDevicePost(ctx, optional)


FCMデバイスを登録します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationDevicePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NotificationDevicePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject23** | [**optional.Interface of InlineObject23**](InlineObject23.md)|  | 

### Return type

 (empty response body)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationGet

> NotificationGet(ctx, )


通知ストリーム(Server Sent Events)に接続します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeNotificationGet

> []string UsersMeNotificationGet(ctx, )


自分が通知を入れているチャンネルのリストを取得します

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIDNotificationGet

> []string UsersUserIDNotificationGet(ctx, userID)


ユーザーが通知を入れているチャンネルのリストを取得します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 

### Return type

**[]string**

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

