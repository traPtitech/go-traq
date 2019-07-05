# \UserGroupApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsGet**](UserGroupApi.md#GroupsGet) | **Get** /groups | 
[**GroupsGroupIDDelete**](UserGroupApi.md#GroupsGroupIDDelete) | **Delete** /groups/{groupID} | 
[**GroupsGroupIDGet**](UserGroupApi.md#GroupsGroupIDGet) | **Get** /groups/{groupID} | 
[**GroupsGroupIDMembersGet**](UserGroupApi.md#GroupsGroupIDMembersGet) | **Get** /groups/{groupID}/members | 
[**GroupsGroupIDMembersPost**](UserGroupApi.md#GroupsGroupIDMembersPost) | **Post** /groups/{groupID}/members | 
[**GroupsGroupIDMembersUserIDDelete**](UserGroupApi.md#GroupsGroupIDMembersUserIDDelete) | **Delete** /groups/{groupID}/members/{userID} | 
[**GroupsGroupIDPatch**](UserGroupApi.md#GroupsGroupIDPatch) | **Patch** /groups/{groupID} | 
[**GroupsPost**](UserGroupApi.md#GroupsPost) | **Post** /groups | 
[**UsersMeGroupsGet**](UserGroupApi.md#UsersMeGroupsGet) | **Get** /users/me/groups | 
[**UsersUserIDGroupsGet**](UserGroupApi.md#UsersUserIDGroupsGet) | **Get** /users/{userID}/groups | 



## GroupsGet

> []UserGroup GroupsGet(ctx, )


全てのユーザーグループを取得します

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIDDelete

> GroupsGroupIDDelete(ctx, groupID)


ユーザーグループを削除します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | [**string**](.md)| 操作の対象となるユーザーグループID | 

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


## GroupsGroupIDGet

> UserGroup GroupsGroupIDGet(ctx, groupID)


ユーザーグループを取得します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | [**string**](.md)| 操作の対象となるユーザーグループID | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIDMembersGet

> []string GroupsGroupIDMembersGet(ctx, groupID)


ユーザーグループのメンバーのIDを取得します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | [**string**](.md)| 操作の対象となるユーザーグループID | 

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


## GroupsGroupIDMembersPost

> GroupsGroupIDMembersPost(ctx, groupID, optional)


ユーザーグループにメンバーを追加します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | [**string**](.md)| 操作の対象となるユーザーグループID | 
 **optional** | ***GroupsGroupIDMembersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGroupIDMembersPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject31** | [**optional.Interface of InlineObject31**](InlineObject31.md)|  | 

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


## GroupsGroupIDMembersUserIDDelete

> GroupsGroupIDMembersUserIDDelete(ctx, groupID, userID)


ユーザーグループからメンバーを削除します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | [**string**](.md)| 操作の対象となるユーザーグループID | 
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 

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


## GroupsGroupIDPatch

> GroupsGroupIDPatch(ctx, groupID, optional)


ユーザーグループの情報を変更します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | [**string**](.md)| 操作の対象となるユーザーグループID | 
 **optional** | ***GroupsGroupIDPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGroupIDPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject30** | [**optional.Interface of InlineObject30**](InlineObject30.md)|  | 

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


## GroupsPost

> UserGroup GroupsPost(ctx, optional)


ユーザーグループを作成します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject29** | [**optional.Interface of InlineObject29**](InlineObject29.md)|  | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeGroupsGet

> []string UsersMeGroupsGet(ctx, )


所属するユーザーグループのIDを取得します

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


## UsersUserIDGroupsGet

> []string UsersUserIDGroupsGet(ctx, userID)


所属するユーザーグループのIDを取得します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 

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

