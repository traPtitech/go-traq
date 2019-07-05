# \MessageApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsChannelIDMessagesGet**](MessageApi.md#ChannelsChannelIDMessagesGet) | **Get** /channels/{channelID}/messages | 
[**ChannelsChannelIDMessagesPost**](MessageApi.md#ChannelsChannelIDMessagesPost) | **Post** /channels/{channelID}/messages | 
[**MessagesMessageIDDelete**](MessageApi.md#MessagesMessageIDDelete) | **Delete** /messages/{messageID} | 
[**MessagesMessageIDGet**](MessageApi.md#MessagesMessageIDGet) | **Get** /messages/{messageID} | 
[**MessagesMessageIDPut**](MessageApi.md#MessagesMessageIDPut) | **Put** /messages/{messageID} | 
[**MessagesMessageIDReportPost**](MessageApi.md#MessagesMessageIDReportPost) | **Post** /messages/{messageID}/report | 
[**MessagesReportsGet**](MessageApi.md#MessagesReportsGet) | **Get** /messages/reports | 
[**UsersUserIDMessagesGet**](MessageApi.md#UsersUserIDMessagesGet) | **Get** /users/{userID}/messages | 
[**UsersUserIDMessagesPost**](MessageApi.md#UsersUserIDMessagesPost) | **Post** /users/{userID}/messages | 



## ChannelsChannelIDMessagesGet

> []Message ChannelsChannelIDMessagesGet(ctx, channelID, optional)


チャンネルに存在するメッセージを取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 
 **optional** | ***ChannelsChannelIDMessagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChannelsChannelIDMessagesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| 取得する件数 1-200 | [default to 0]
 **offset** | **optional.Int32**| 取得するオフセット | [default to 0]
 **since** | **optional.Time**| 取得する時間範囲の開始日時 | [default to Thu Jan 01 09:00:00 JST 1]
 **until** | **optional.Time**| 取得する時間範囲の終了日時 | [default to Sat Jan 01 09:00:39 JST 10000]
 **inclusive** | **optional.Bool**| 範囲の端を含めるかどうか | [default to false]
 **order** | **optional.String**| 昇順か降順か | [default to desc]

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


## ChannelsChannelIDMessagesPost

> Message ChannelsChannelIDMessagesPost(ctx, channelID, optional)


チャンネルにメッセージを投稿します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 
 **optional** | ***ChannelsChannelIDMessagesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChannelsChannelIDMessagesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject20** | [**optional.Interface of InlineObject20**](InlineObject20.md)|  | 

### Return type

[**Message**](Message.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesMessageIDDelete

> MessagesMessageIDDelete(ctx, messageID)


指定したメッセージを削除します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageID** | [**string**](.md)| 操作の対象となるメッセージID | 

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


## MessagesMessageIDGet

> Message MessagesMessageIDGet(ctx, messageID)


指定したメッセージを取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageID** | [**string**](.md)| 操作の対象となるメッセージID | 

### Return type

[**Message**](Message.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesMessageIDPut

> MessagesMessageIDPut(ctx, messageID, optional)


指定したメッセージを編集します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageID** | [**string**](.md)| 操作の対象となるメッセージID | 
 **optional** | ***MessagesMessageIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MessagesMessageIDPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject24** | [**optional.Interface of InlineObject24**](InlineObject24.md)|  | 

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


## MessagesMessageIDReportPost

> MessagesMessageIDReportPost(ctx, messageID, optional)


指定したメッセージを通報します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageID** | [**string**](.md)| 操作の対象となるメッセージID | 
 **optional** | ***MessagesMessageIDReportPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MessagesMessageIDReportPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject25** | [**optional.Interface of InlineObject25**](InlineObject25.md)|  | 

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


## MessagesReportsGet

> []InlineResponse2009 MessagesReportsGet(ctx, optional)


メッセージ通報を最大50件取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesReportsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MessagesReportsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **p** | **optional.Int32**| ページ番号(ゼロオリジン) | 

### Return type

[**[]InlineResponse2009**](inline_response_200_9.md)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIDMessagesGet

> []Message UsersUserIDMessagesGet(ctx, userID, optional)


DMチャンネルに存在するメッセージを取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 
 **optional** | ***UsersUserIDMessagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersUserIDMessagesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| 取得する件数 1-200 | [default to 0]
 **offset** | **optional.Int32**| 取得するオフセット | [default to 0]
 **since** | **optional.Time**| 取得する時間範囲の開始日時 | [default to Thu Jan 01 09:00:00 JST 1]
 **until** | **optional.Time**| 取得する時間範囲の終了日時 | [default to Sat Jan 01 09:00:39 JST 10000]
 **inclusive** | **optional.Bool**| 範囲の端を含めるかどうか | [default to false]
 **order** | **optional.String**| 昇順か降順か | [default to desc]

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


## UsersUserIDMessagesPost

> Message UsersUserIDMessagesPost(ctx, userID, optional)


DMチャンネルにメッセージを投稿します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 
 **optional** | ***UsersUserIDMessagesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersUserIDMessagesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject11** | [**optional.Interface of InlineObject11**](InlineObject11.md)|  | 

### Return type

[**Message**](Message.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

