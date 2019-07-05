# \StarApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMeStarsChannelIDDelete**](StarApi.md#UsersMeStarsChannelIDDelete) | **Delete** /users/me/stars/{channelID} | 
[**UsersMeStarsChannelIDPut**](StarApi.md#UsersMeStarsChannelIDPut) | **Put** /users/me/stars/{channelID} | 
[**UsersMeStarsGet**](StarApi.md#UsersMeStarsGet) | **Get** /users/me/stars | 



## UsersMeStarsChannelIDDelete

> UsersMeStarsChannelIDDelete(ctx, channelID)


+| お気に入りチャンネルリストから削除します。 既にお気に入りチャンネルリストに無いチャンネルを指定した場合は無視されます(204)。

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


## UsersMeStarsChannelIDPut

> UsersMeStarsChannelIDPut(ctx, channelID)


お気に入りチャンネルリストにチャンネルを追加します。

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


## UsersMeStarsGet

> []string UsersMeStarsGet(ctx, )


お気に入りチャンネルリストを取得します。

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

