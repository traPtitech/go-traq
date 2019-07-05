# \ActivityApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivityLatestMessagesGet**](ActivityApi.md#ActivityLatestMessagesGet) | **Get** /activity/latest-messages | 



## ActivityLatestMessagesGet

> []Message ActivityLatestMessagesGet(ctx, optional)


各チャンネルの最新のメッセージ１件をメッセージ作成日時の降順で取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ActivityLatestMessagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ActivityLatestMessagesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| 取得する件数 1-50 | [default to 50]
 **subscribe** | **optional.Bool**| 購読チャンネルのみを取得する | [default to true]

### Return type

[**[]Message**](Message.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

