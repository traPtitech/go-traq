# \WebhookApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhooksGet**](WebhookApi.md#WebhooksGet) | **Get** /webhooks | 
[**WebhooksPost**](WebhookApi.md#WebhooksPost) | **Post** /webhooks | 
[**WebhooksWebhookIDDelete**](WebhookApi.md#WebhooksWebhookIDDelete) | **Delete** /webhooks/{webhookID} | 
[**WebhooksWebhookIDGet**](WebhookApi.md#WebhooksWebhookIDGet) | **Get** /webhooks/{webhookID} | 
[**WebhooksWebhookIDGithubPost**](WebhookApi.md#WebhooksWebhookIDGithubPost) | **Post** /webhooks/{webhookID}/github | 
[**WebhooksWebhookIDIconGet**](WebhookApi.md#WebhooksWebhookIDIconGet) | **Get** /webhooks/{webhookID}/icon | 
[**WebhooksWebhookIDIconPut**](WebhookApi.md#WebhooksWebhookIDIconPut) | **Put** /webhooks/{webhookID}/icon | 
[**WebhooksWebhookIDPatch**](WebhookApi.md#WebhooksWebhookIDPatch) | **Patch** /webhooks/{webhookID} | 
[**WebhooksWebhookIDPost**](WebhookApi.md#WebhooksWebhookIDPost) | **Post** /webhooks/{webhookID} | 



## WebhooksGet

> []Webhook WebhooksGet(ctx, )


自分が作成したwebhookの一覧を取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksPost

> Webhook WebhooksPost(ctx, optional)


webhookを作成します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhooksPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject27** | [**optional.Interface of InlineObject27**](InlineObject27.md)|  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksWebhookIDDelete

> WebhooksWebhookIDDelete(ctx, webhookID)


webhookを削除します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | [**string**](.md)| 操作の対象となるWebhookのID | 

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


## WebhooksWebhookIDGet

> Webhook WebhooksWebhookIDGet(ctx, webhookID)


webhookの詳細を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | [**string**](.md)| 操作の対象となるWebhookのID | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksWebhookIDGithubPost

> WebhooksWebhookIDGithubPost(ctx, webhookID, optional)


Github-Compatibleなwebhookを送信します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | [**string**](.md)| 操作の対象となるWebhookのID | 
 **optional** | ***WebhooksWebhookIDGithubPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhooksWebhookIDGithubPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.Map[string]interface{}**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksWebhookIDIconGet

> *os.File WebhooksWebhookIDIconGet(ctx, webhookID)


指定したWebhookのアイコンを取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | [**string**](.md)| 操作の対象となるWebhookのID | 

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


## WebhooksWebhookIDIconPut

> WebhooksWebhookIDIconPut(ctx, webhookID, optional)


指定したwebhookのアイコンを変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | [**string**](.md)| 操作の対象となるWebhookのID | 
 **optional** | ***WebhooksWebhookIDIconPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhooksWebhookIDIconPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.*os.File**| webhookユーザーのアイコン(1MBまでのpng, jpeg, gif) | 

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


## WebhooksWebhookIDPatch

> WebhooksWebhookIDPatch(ctx, webhookID, optional)


webhookを修正します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | [**string**](.md)| 操作の対象となるWebhookのID | 
 **optional** | ***WebhooksWebhookIDPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhooksWebhookIDPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject28** | [**optional.Interface of InlineObject28**](InlineObject28.md)|  | 

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


## WebhooksWebhookIDPost

> WebhooksWebhookIDPost(ctx, webhookID, optional)


webhookを送信します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | [**string**](.md)| 操作の対象となるWebhookのID | 
 **optional** | ***WebhooksWebhookIDPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhooksWebhookIDPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTRAQChannelId** | [**optional.Interface of string**](.md)| 投稿先のチャンネルID(変更する場合) | 
 **xTRAQSignature** | **optional.String**| リクエストボディシグネチャ。Secretが設定されている場合は必須 | 
 **body** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

