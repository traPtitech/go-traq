# \UserApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGet**](UserApi.md#UsersGet) | **Get** /users | 
[**UsersMeGet**](UserApi.md#UsersMeGet) | **Get** /users/me | 
[**UsersMeIconGet**](UserApi.md#UsersMeIconGet) | **Get** /users/me/icon | 
[**UsersMeIconPut**](UserApi.md#UsersMeIconPut) | **Put** /users/me/icon | 
[**UsersMePasswordPut**](UserApi.md#UsersMePasswordPut) | **Put** /users/me/password | 
[**UsersMePatch**](UserApi.md#UsersMePatch) | **Patch** /users/me | 
[**UsersMeQrCodeGet**](UserApi.md#UsersMeQrCodeGet) | **Get** /users/me/qr-code | 
[**UsersPost**](UserApi.md#UsersPost) | **Post** /users | 
[**UsersUserIDGet**](UserApi.md#UsersUserIDGet) | **Get** /users/{userID} | 
[**UsersUserIDIconGet**](UserApi.md#UsersUserIDIconGet) | **Get** /users/{userID}/icon | 



## UsersGet

> []User UsersGet(ctx, )


全ユーザーのリストを取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]User**](User.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeGet

> Me UsersMeGet(ctx, )


自分のユーザー情報を取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Me**](Me.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeIconGet

> *os.File UsersMeIconGet(ctx, )


自分のアイコン画像を取得します。

### Required Parameters

This endpoint does not need any parameter.

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


## UsersMeIconPut

> UsersMeIconPut(ctx, optional)


自分のアイコンを更新します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersMeIconPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMeIconPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.*os.File**| 1MBまでのpng, jpeg, gif | 

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


## UsersMePasswordPut

> UsersMePasswordPut(ctx, optional)


自分のパスワードを変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersMePasswordPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMePasswordPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject6** | [**optional.Interface of InlineObject6**](InlineObject6.md)|  | 

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


## UsersMePatch

> UsersMePatch(ctx, optional)


自分のユーザー情報を変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersMePatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMePatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject5** | [**optional.Interface of InlineObject5**](InlineObject5.md)|  | 

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


## UsersMeQrCodeGet

> *os.File UsersMeQrCodeGet(ctx, )


自分のQRコードを取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[***os.File**](*os.File.md)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPost

> UsersPost(ctx, optional)


ユーザー登録します

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject4** | [**optional.Interface of InlineObject4**](InlineObject4.md)|  | 

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


## UsersUserIDGet

> UserDetail UsersUserIDGet(ctx, userID)


ユーザーの詳細を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 

### Return type

[**UserDetail**](UserDetail.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIDIconGet

> *os.File UsersUserIDIconGet(ctx, userID)


ユーザーのアイコン画像を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | [**string**](.md)| 操作の対象となるユーザーID | 

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

