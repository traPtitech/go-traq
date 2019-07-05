# \ClipApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMeClipsClipIDDelete**](ClipApi.md#UsersMeClipsClipIDDelete) | **Delete** /users/me/clips/{clipID} | 
[**UsersMeClipsClipIDFolderGet**](ClipApi.md#UsersMeClipsClipIDFolderGet) | **Get** /users/me/clips/{clipID}/folder | 
[**UsersMeClipsClipIDFolderPut**](ClipApi.md#UsersMeClipsClipIDFolderPut) | **Put** /users/me/clips/{clipID}/folder | 
[**UsersMeClipsClipIDGet**](ClipApi.md#UsersMeClipsClipIDGet) | **Get** /users/me/clips/{clipID} | 
[**UsersMeClipsFoldersFolderIDDelete**](ClipApi.md#UsersMeClipsFoldersFolderIDDelete) | **Delete** /users/me/clips/folders/{folderID} | 
[**UsersMeClipsFoldersFolderIDGet**](ClipApi.md#UsersMeClipsFoldersFolderIDGet) | **Get** /users/me/clips/folders/{folderID} | 
[**UsersMeClipsFoldersFolderIDPatch**](ClipApi.md#UsersMeClipsFoldersFolderIDPatch) | **Patch** /users/me/clips/folders/{folderID} | 
[**UsersMeClipsFoldersGet**](ClipApi.md#UsersMeClipsFoldersGet) | **Get** /users/me/clips/folders | 
[**UsersMeClipsFoldersPost**](ClipApi.md#UsersMeClipsFoldersPost) | **Post** /users/me/clips/folders | 
[**UsersMeClipsGet**](ClipApi.md#UsersMeClipsGet) | **Get** /users/me/clips | 
[**UsersMeClipsPost**](ClipApi.md#UsersMeClipsPost) | **Post** /users/me/clips | 



## UsersMeClipsClipIDDelete

> UsersMeClipsClipIDDelete(ctx, clipID)


指定したクリップを削除します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clipID** | **string**| 捜査の対象となるクリップのID | 

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


## UsersMeClipsClipIDFolderGet

> ClipsFolder UsersMeClipsClipIDFolderGet(ctx, clipID)


指定したクリップのフォルダ情報を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clipID** | **string**| 捜査の対象となるクリップのID | 

### Return type

[**ClipsFolder**](ClipsFolder.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeClipsClipIDFolderPut

> UsersMeClipsClipIDFolderPut(ctx, clipID, optional)


指定したクリップのフォルダを変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clipID** | **string**| 捜査の対象となるクリップのID | 
 **optional** | ***UsersMeClipsClipIDFolderPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMeClipsClipIDFolderPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject8** | [**optional.Interface of InlineObject8**](InlineObject8.md)|  | 

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


## UsersMeClipsClipIDGet

> Message UsersMeClipsClipIDGet(ctx, clipID)


指定したクリップのメッセージを取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clipID** | **string**| 捜査の対象となるクリップのID | 

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


## UsersMeClipsFoldersFolderIDDelete

> UsersMeClipsFoldersFolderIDDelete(ctx, folderID)


クリップフォルダを削除します。フォルダ内のクリップは全て削除されます。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderID** | [**string**](.md)| 操作の対象となるクリップフォルダのID | 

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


## UsersMeClipsFoldersFolderIDGet

> []InlineResponse2004 UsersMeClipsFoldersFolderIDGet(ctx, folderID)


フォルダ内のクリップ一覧を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderID** | [**string**](.md)| 操作の対象となるクリップフォルダのID | 

### Return type

[**[]InlineResponse2004**](inline_response_200_4.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeClipsFoldersFolderIDPatch

> UsersMeClipsFoldersFolderIDPatch(ctx, folderID, optional)


クリップフォルダ名を変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderID** | [**string**](.md)| 操作の対象となるクリップフォルダのID | 
 **optional** | ***UsersMeClipsFoldersFolderIDPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMeClipsFoldersFolderIDPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject10** | [**optional.Interface of InlineObject10**](InlineObject10.md)|  | 

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


## UsersMeClipsFoldersGet

> []ClipsFolder UsersMeClipsFoldersGet(ctx, )


クリップフォルダ一覧を取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ClipsFolder**](ClipsFolder.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeClipsFoldersPost

> ClipsFolder UsersMeClipsFoldersPost(ctx, optional)


クリップフォルダを作成します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersMeClipsFoldersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMeClipsFoldersPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject9** | [**optional.Interface of InlineObject9**](InlineObject9.md)|  | 

### Return type

[**ClipsFolder**](ClipsFolder.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeClipsGet

> []InlineResponse2003 UsersMeClipsGet(ctx, )


全てのクリップを取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2003**](inline_response_200_3.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeClipsPost

> InlineResponse201 UsersMeClipsPost(ctx, optional)


新しくメッセージをクリップします。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersMeClipsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMeClipsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject7** | [**optional.Interface of InlineObject7**](InlineObject7.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

