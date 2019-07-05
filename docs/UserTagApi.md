# \UserTagApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagsTagIDGet**](UserTagApi.md#TagsTagIDGet) | **Get** /tags/{tagID} | 
[**UsersUserIDTagsGet**](UserTagApi.md#UsersUserIDTagsGet) | **Get** /users/{userID}/tags | 
[**UsersUserIDTagsPost**](UserTagApi.md#UsersUserIDTagsPost) | **Post** /users/{userID}/tags | 
[**UsersUserIDTagsTagIDDelete**](UserTagApi.md#UsersUserIDTagsTagIDDelete) | **Delete** /users/{userID}/tags/{tagID} | 
[**UsersUserIDTagsTagIDPatch**](UserTagApi.md#UsersUserIDTagsTagIDPatch) | **Patch** /users/{userID}/tags/{tagID} | 



## TagsTagIDGet

> InlineResponse20010 TagsTagIDGet(ctx, tagID)


指定されたタグの情報を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagID** | [**string**](.md)| 操作の対象となるタグID | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIDTagsGet

> []Tag UsersUserIDTagsGet(ctx, userID)


ユーザーのタグのリストを取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIDTagsPost

> UsersUserIDTagsPost(ctx, userID, optional)


ユーザーにタグを追加します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 
 **optional** | ***UsersUserIDTagsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersUserIDTagsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject12** | [**optional.Interface of InlineObject12**](InlineObject12.md)|  | 

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


## UsersUserIDTagsTagIDDelete

> UsersUserIDTagsTagIDDelete(ctx, userID, tagID)


ユーザーから指定したタグを削除します。既に存在しないタグを削除しようとした場合は無視されます(204)。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 
**tagID** | [**string**](.md)| 操作の対象となるタグID | 

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


## UsersUserIDTagsTagIDPatch

> UsersUserIDTagsTagIDPatch(ctx, userID, tagID, optional)


タグのロック、アンロックを変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 
**tagID** | [**string**](.md)| 操作の対象となるタグID | 
 **optional** | ***UsersUserIDTagsTagIDPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersUserIDTagsTagIDPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject13** | [**optional.Interface of InlineObject13**](InlineObject13.md)|  | 

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

