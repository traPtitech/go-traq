# \FileApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesFileIDDelete**](FileApi.md#FilesFileIDDelete) | **Delete** /files/{fileID} | 
[**FilesFileIDGet**](FileApi.md#FilesFileIDGet) | **Get** /files/{fileID} | 
[**FilesFileIDMetaGet**](FileApi.md#FilesFileIDMetaGet) | **Get** /files/{fileID}/meta | 
[**FilesFileIDThumbnailGet**](FileApi.md#FilesFileIDThumbnailGet) | **Get** /files/{fileID}/thumbnail | 
[**FilesPost**](FileApi.md#FilesPost) | **Post** /files | 



## FilesFileIDDelete

> FilesFileIDDelete(ctx, fileID)


指定したファイルを削除します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| 操作の対象となるファイルID | 

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


## FilesFileIDGet

> *os.File FilesFileIDGet(ctx, fileID, optional)


指定したファイルの中身を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| 操作の対象となるファイルID | 
 **optional** | ***FilesFileIDGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilesFileIDGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dl** | **optional.Int32**| ファイルをダウンロードする場合に1を指定する | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesFileIDMetaGet

> FileInfo FilesFileIDMetaGet(ctx, fileID)


指定したファイルのメタデータを取得します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| 操作の対象となるファイルID | 

### Return type

[**FileInfo**](FileInfo.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesFileIDThumbnailGet

> *os.File FilesFileIDThumbnailGet(ctx, fileID)


指定したファイルのサムネイルを取得します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| 操作の対象となるファイルID | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesPost

> FileInfo FilesPost(ctx, file, optional)


ファイルをアップロードします

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**file** | ***os.File*****os.File**| ファイル本体 | 
 **optional** | ***FilesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aclReadable** | **optional.String**| ファイルにアクセスすることが可能なユーザーのUUIDの文字列表現をカンマ区切りで連結した文字列、または&#39;all&#39;(デフォルト) | [default to all]

### Return type

[**FileInfo**](FileInfo.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

