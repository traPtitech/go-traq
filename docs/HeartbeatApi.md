# \HeartbeatApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HeartbeatGet**](HeartbeatApi.md#HeartbeatGet) | **Get** /heartbeat | 
[**HeartbeatPost**](HeartbeatApi.md#HeartbeatPost) | **Post** /heartbeat | 



## HeartbeatGet

> HeartbeatRes HeartbeatGet(ctx, channelId)


チャンネルを現在見ている人・編集している人を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | [**string**](.md)| 現在いるチャンネルId | 

### Return type

[**HeartbeatRes**](HeartbeatRes.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeartbeatPost

> HeartbeatPost(ctx, optional)


どのチャンネルを見ているか・編集しているかを送信します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HeartbeatPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeartbeatPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject14** | [**optional.Interface of InlineObject14**](InlineObject14.md)|  | 

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

