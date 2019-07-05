# \StampApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MessagesMessageIDStampsGet**](StampApi.md#MessagesMessageIDStampsGet) | **Get** /messages/{messageID}/stamps | 
[**MessagesMessageIDStampsStampIDDelete**](StampApi.md#MessagesMessageIDStampsStampIDDelete) | **Delete** /messages/{messageID}/stamps/{stampID} | 
[**MessagesMessageIDStampsStampIDPost**](StampApi.md#MessagesMessageIDStampsStampIDPost) | **Post** /messages/{messageID}/stamps/{stampID} | 
[**StampsGet**](StampApi.md#StampsGet) | **Get** /stamps | 
[**StampsPost**](StampApi.md#StampsPost) | **Post** /stamps | 
[**StampsStampIDDelete**](StampApi.md#StampsStampIDDelete) | **Delete** /stamps/{stampID} | 
[**StampsStampIDGet**](StampApi.md#StampsStampIDGet) | **Get** /stamps/{stampID} | 
[**StampsStampIDPatch**](StampApi.md#StampsStampIDPatch) | **Patch** /stamps/{stampID} | 
[**UsersMeStampHistoryGet**](StampApi.md#UsersMeStampHistoryGet) | **Get** /users/me/stamp-history | 



## MessagesMessageIDStampsGet

> []MessageStamp MessagesMessageIDStampsGet(ctx, messageID)


指定したメッセージに押されているスタンプを全て取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageID** | [**string**](.md)| 操作の対象となるメッセージID | 

### Return type

[**[]MessageStamp**](MessageStamp.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesMessageIDStampsStampIDDelete

> MessagesMessageIDStampsStampIDDelete(ctx, messageID, stampID)


指定したメッセージから指定したスタンプを外します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageID** | [**string**](.md)| 操作の対象となるメッセージID | 
**stampID** | [**string**](.md)| 操作の対象となるスタンプID | 

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


## MessagesMessageIDStampsStampIDPost

> MessagesMessageIDStampsStampIDPost(ctx, messageID, stampID)


指定したメッセージに指定したスタンプを押します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageID** | [**string**](.md)| 操作の対象となるメッセージID | 
**stampID** | [**string**](.md)| 操作の対象となるスタンプID | 

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


## StampsGet

> []Stamp StampsGet(ctx, )


全スタンプのリストを取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Stamp**](Stamp.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsPost

> Stamp StampsPost(ctx, name, file)


スタンプを新規作成します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| スタンプ名(半角英数字と-+_のみを含む32文字以内の文字列) | 
**file** | ***os.File*****os.File**| 1MBまでのpng, jpeg, gif, svg | 

### Return type

[**Stamp**](Stamp.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsStampIDDelete

> StampsStampIDDelete(ctx, stampID)


スタンプを削除します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stampID** | [**string**](.md)| 操作の対象となるスタンプID | 

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


## StampsStampIDGet

> Stamp StampsStampIDGet(ctx, stampID)


スタンプの情報を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stampID** | [**string**](.md)| 操作の対象となるスタンプID | 

### Return type

[**Stamp**](Stamp.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsStampIDPatch

> StampsStampIDPatch(ctx, stampID, optional)


スタンプを修正します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stampID** | [**string**](.md)| 操作の対象となるスタンプID | 
 **optional** | ***StampsStampIDPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StampsStampIDPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| スタンプ名(半角英数字と-+_のみを含む32文字以内の文字列) | 
 **file** | **optional.Interface of *os.File****optional.*os.File**| 1MBまでのpng, jpeg, gif, svg | 

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


## UsersMeStampHistoryGet

> []InlineResponse2001 UsersMeStampHistoryGet(ctx, )


自分のスタンプ履歴を最大50件取得します。結果は降順で返されます。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

