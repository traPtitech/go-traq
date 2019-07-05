# \UnreadApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMeUnreadChannelsChannelIDDelete**](UnreadApi.md#UsersMeUnreadChannelsChannelIDDelete) | **Delete** /users/me/unread/channels/{channelID} | 
[**UsersMeUnreadChannelsGet**](UnreadApi.md#UsersMeUnreadChannelsGet) | **Get** /users/me/unread/channels | 



## UsersMeUnreadChannelsChannelIDDelete

> UsersMeUnreadChannelsChannelIDDelete(ctx, channelID)


指定されたチャンネルの未読メッセージを既読にします。存在しないチャンネルIDを指定した場合は、無視されます。

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


## UsersMeUnreadChannelsGet

> []InlineResponse2005 UsersMeUnreadChannelsGet(ctx, )


未読チャンネル情報のリストを取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2005**](inline_response_200_5.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

