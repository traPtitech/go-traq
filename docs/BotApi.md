# \BotApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotsBotIDChannelsGet**](BotApi.md#BotsBotIDChannelsGet) | **Get** /bots/{botID}/channels | 
[**BotsBotIDDelete**](BotApi.md#BotsBotIDDelete) | **Delete** /bots/{botID} | 
[**BotsBotIDDetailGet**](BotApi.md#BotsBotIDDetailGet) | **Get** /bots/{botID}/detail | 
[**BotsBotIDEventsPut**](BotApi.md#BotsBotIDEventsPut) | **Put** /bots/{botID}/events | 
[**BotsBotIDGet**](BotApi.md#BotsBotIDGet) | **Get** /bots/{botID} | 
[**BotsBotIDIconGet**](BotApi.md#BotsBotIDIconGet) | **Get** /bots/{botID}/icon | 
[**BotsBotIDIconPut**](BotApi.md#BotsBotIDIconPut) | **Put** /bots/{botID}/icon | 
[**BotsBotIDPatch**](BotApi.md#BotsBotIDPatch) | **Patch** /bots/{botID} | 
[**BotsBotIDReissuePost**](BotApi.md#BotsBotIDReissuePost) | **Post** /bots/{botID}/reissue | 
[**BotsBotIDStatePut**](BotApi.md#BotsBotIDStatePut) | **Put** /bots/{botID}/state | 
[**BotsGet**](BotApi.md#BotsGet) | **Get** /bots | 
[**BotsPost**](BotApi.md#BotsPost) | **Post** /bots | 
[**ChannelsChannelIDBotsBotIDDelete**](BotApi.md#ChannelsChannelIDBotsBotIDDelete) | **Delete** /channels/{channelID}/bots/{botID} | 
[**ChannelsChannelIDBotsGet**](BotApi.md#ChannelsChannelIDBotsGet) | **Get** /channels/{channelID}/bots | 
[**ChannelsChannelIDBotsPost**](BotApi.md#ChannelsChannelIDBotsPost) | **Post** /channels/{channelID}/bots | 



## BotsBotIDChannelsGet

> []string BotsBotIDChannelsGet(ctx, botID)


Botが参加しているチャンネルのUUIDの配列を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botID** | [**string**](.md)| 操作の対象となるBotのID | 

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


## BotsBotIDDelete

> BotsBotIDDelete(ctx, botID)


Botを削除します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botID** | [**string**](.md)| 操作の対象となるBotのID | 

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


## BotsBotIDDetailGet

> BotDetail BotsBotIDDetailGet(ctx, botID)


Botの詳細を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botID** | [**string**](.md)| 操作の対象となるBotのID | 

### Return type

[**BotDetail**](BotDetail.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotsBotIDEventsPut

> BotsBotIDEventsPut(ctx, botID, optional)


Botの購読イベントを変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botID** | [**string**](.md)| 操作の対象となるBotのID | 
 **optional** | ***BotsBotIDEventsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BotsBotIDEventsPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject36** | [**optional.Interface of InlineObject36**](InlineObject36.md)|  | 

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


## BotsBotIDGet

> Bot BotsBotIDGet(ctx, botID)


Botを取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botID** | [**string**](.md)| 操作の対象となるBotのID | 

### Return type

[**Bot**](Bot.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotsBotIDIconGet

> *os.File BotsBotIDIconGet(ctx, botID)


指定したBotのアイコンを取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botID** | [**string**](.md)| 操作の対象となるBotのID | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, image/png, image/gif

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotsBotIDIconPut

> BotsBotIDIconPut(ctx, botID, optional)


指定したBotのアイコンを変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botID** | [**string**](.md)| 操作の対象となるBotのID | 
 **optional** | ***BotsBotIDIconPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BotsBotIDIconPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.*os.File**| Botユーザーのアイコン(1MBまでのpng, jpeg, gif) | 

### Return type

 (empty response body)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotsBotIDPatch

> BotsBotIDPatch(ctx, botID, optional)


Bot情報を変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botID** | [**string**](.md)| 操作の対象となるBotのID | 
 **optional** | ***BotsBotIDPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BotsBotIDPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject35** | [**optional.Interface of InlineObject35**](InlineObject35.md)|  | 

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


## BotsBotIDReissuePost

> InlineResponse20011 BotsBotIDReissuePost(ctx, botID)


Botの各種トークンを再発行します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botID** | [**string**](.md)| 操作の対象となるBotのID | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotsBotIDStatePut

> BotsBotIDStatePut(ctx, botID, optional)


Botの状態を変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botID** | [**string**](.md)| 操作の対象となるBotのID | 
 **optional** | ***BotsBotIDStatePutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BotsBotIDStatePutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject37** | [**optional.Interface of InlineObject37**](InlineObject37.md)|  | 

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


## BotsGet

> []Bot BotsGet(ctx, )


自分が作成したBotの一覧を取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Bot**](Bot.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotsPost

> BotDetail BotsPost(ctx, optional)


Botを作成します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BotsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BotsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject34** | [**optional.Interface of InlineObject34**](InlineObject34.md)|  | 

### Return type

[**BotDetail**](BotDetail.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelIDBotsBotIDDelete

> ChannelsChannelIDBotsBotIDDelete(ctx, channelID, botID)


Botをチャンネルから退出させます。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 
**botID** | [**string**](.md)| 操作の対象となるBotのID | 

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


## ChannelsChannelIDBotsGet

> []InlineResponse2007 ChannelsChannelIDBotsGet(ctx, channelID)


チャンネルに参加しているBot一覧を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 

### Return type

[**[]InlineResponse2007**](inline_response_200_7.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelIDBotsPost

> InlineResponse2008 ChannelsChannelIDBotsPost(ctx, channelID, optional)


チャンネルにBotを参加させます。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 
 **optional** | ***ChannelsChannelIDBotsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChannelsChannelIDBotsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject22** | [**optional.Interface of InlineObject22**](InlineObject22.md)|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

