# \MuteApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMeMuteChannelIDDelete**](MuteApi.md#UsersMeMuteChannelIDDelete) | **Delete** /users/me/mute/{channelID} | 
[**UsersMeMuteChannelIDPost**](MuteApi.md#UsersMeMuteChannelIDPost) | **Post** /users/me/mute/{channelID} | 
[**UsersMeMuteGet**](MuteApi.md#UsersMeMuteGet) | **Get** /users/me/mute | 



## UsersMeMuteChannelIDDelete

> UsersMeMuteChannelIDDelete(ctx, channelID)


指定したチャンネルのミュートを解除します。既に解除されていた場合は204を返します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 

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


## UsersMeMuteChannelIDPost

> UsersMeMuteChannelIDPost(ctx, channelID)


指定したチャンネルをミュートします。ただし、強制通知チャンネルはミュートできません。既にミュートしていた場合は204を返します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 

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


## UsersMeMuteGet

> []string UsersMeMuteGet(ctx, )


ミュートしているチャンネルのIDの配列を返します。

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

