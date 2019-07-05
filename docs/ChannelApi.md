# \ChannelApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsChannelIDChildrenPost**](ChannelApi.md#ChannelsChannelIDChildrenPost) | **Post** /channels/{channelID}/children | 
[**ChannelsChannelIDDelete**](ChannelApi.md#ChannelsChannelIDDelete) | **Delete** /channels/{channelID} | 
[**ChannelsChannelIDGet**](ChannelApi.md#ChannelsChannelIDGet) | **Get** /channels/{channelID} | 
[**ChannelsChannelIDParentPut**](ChannelApi.md#ChannelsChannelIDParentPut) | **Put** /channels/{channelID}/parent | 
[**ChannelsChannelIDPatch**](ChannelApi.md#ChannelsChannelIDPatch) | **Patch** /channels/{channelID} | 
[**ChannelsChannelIDTopicGet**](ChannelApi.md#ChannelsChannelIDTopicGet) | **Get** /channels/{channelID}/topic | 
[**ChannelsChannelIDTopicPut**](ChannelApi.md#ChannelsChannelIDTopicPut) | **Put** /channels/{channelID}/topic | 
[**ChannelsGet**](ChannelApi.md#ChannelsGet) | **Get** /channels | 
[**ChannelsPost**](ChannelApi.md#ChannelsPost) | **Post** /channels | 



## ChannelsChannelIDChildrenPost

> Channel ChannelsChannelIDChildrenPost(ctx, channelID, optional)


子チャンネルを作成します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 
 **optional** | ***ChannelsChannelIDChildrenPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChannelsChannelIDChildrenPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject18** | [**optional.Interface of InlineObject18**](InlineObject18.md)|  | 

### Return type

[**Channel**](Channel.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelIDDelete

> ChannelsChannelIDDelete(ctx, channelID)


チャンネルを削除します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelIDGet

> Channel ChannelsChannelIDGet(ctx, channelID)


チャンネルの情報を返します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 

### Return type

[**Channel**](Channel.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelIDParentPut

> ChannelsChannelIDParentPut(ctx, channelID, optional)


チャンネルの親チャンネルを変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 
 **optional** | ***ChannelsChannelIDParentPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChannelsChannelIDParentPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject17** | [**optional.Interface of InlineObject17**](InlineObject17.md)|  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelIDPatch

> ChannelsChannelIDPatch(ctx, channelID, optional)


チャンネルの情報を変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 
 **optional** | ***ChannelsChannelIDPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChannelsChannelIDPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject16** | [**optional.Interface of InlineObject16**](InlineObject16.md)|  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelIDTopicGet

> InlineResponse2006 ChannelsChannelIDTopicGet(ctx, channelID)


チャンネルの説明を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelIDTopicPut

> ChannelsChannelIDTopicPut(ctx, channelID, optional)


チャンネルの説明を変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 
 **optional** | ***ChannelsChannelIDTopicPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChannelsChannelIDTopicPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject19** | [**optional.Interface of InlineObject19**](InlineObject19.md)|  | 

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


## ChannelsGet

> []Channel ChannelsGet(ctx, )


(すべての)チャンネルのリストを取得します。 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Channel**](Channel.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsPost

> Channel ChannelsPost(ctx, optional)


チャンネルを作成します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChannelsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject15** | [**optional.Interface of InlineObject15**](InlineObject15.md)|  | 

### Return type

[**Channel**](Channel.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

