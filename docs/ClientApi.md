# \ClientApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientsClientIDDelete**](ClientApi.md#ClientsClientIDDelete) | **Delete** /clients/{clientID} | 
[**ClientsClientIDDetailGet**](ClientApi.md#ClientsClientIDDetailGet) | **Get** /clients/{clientID}/detail | 
[**ClientsClientIDGet**](ClientApi.md#ClientsClientIDGet) | **Get** /clients/{clientID} | 
[**ClientsClientIDPatch**](ClientApi.md#ClientsClientIDPatch) | **Patch** /clients/{clientID} | 
[**ClientsGet**](ClientApi.md#ClientsGet) | **Get** /clients | 
[**ClientsPost**](ClientApi.md#ClientsPost) | **Post** /clients | 
[**UsersMeTokensGet**](ClientApi.md#UsersMeTokensGet) | **Get** /users/me/tokens | 
[**UsersMeTokensOauth2TokenIDDelete**](ClientApi.md#UsersMeTokensOauth2TokenIDDelete) | **Delete** /users/me/tokens/{oauth2TokenID} | 



## ClientsClientIDDelete

> ClientsClientIDDelete(ctx, clientID)


指定したクライアントIDのクライアントを削除します。 正常に削除された場合、このクライアントを通じての認可は全て取り消されます。 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientID** | [**string**](.md)| 操作の対象となるclientのID | 

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


## ClientsClientIDDetailGet

> OwnedClientInfo ClientsClientIDDetailGet(ctx, clientID)


クライアントの詳細を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientID** | [**string**](.md)| 操作の対象となるclientのID | 

### Return type

[**OwnedClientInfo**](OwnedClientInfo.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientsClientIDGet

> ClientInfo ClientsClientIDGet(ctx, clientID)


指定したクライアントIDのクライアントの情報を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientID** | [**string**](.md)| 操作の対象となるclientのID | 

### Return type

[**ClientInfo**](ClientInfo.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientsClientIDPatch

> ClientsClientIDPatch(ctx, clientID, optional)


指定したクライアントIDのクライアントの情報を変更します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientID** | [**string**](.md)| 操作の対象となるclientのID | 
 **optional** | ***ClientsClientIDPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClientsClientIDPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject33** | [**optional.Interface of InlineObject33**](InlineObject33.md)|  | 

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


## ClientsGet

> []OwnedClientInfo ClientsGet(ctx, )


自分が登録しているクライアントの一覧を取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]OwnedClientInfo**](OwnedClientInfo.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientsPost

> OwnedClientInfo ClientsPost(ctx, optional)


クライアントを登録します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClientsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject32** | [**optional.Interface of InlineObject32**](InlineObject32.md)|  | 

### Return type

[**OwnedClientInfo**](OwnedClientInfo.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeTokensGet

> []AllowedClientInfo UsersMeTokensGet(ctx, )


自分が許可しているクライアントの一覧とトークン情報を取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AllowedClientInfo**](AllowedClientInfo.md)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeTokensOauth2TokenIDDelete

> UsersMeTokensOauth2TokenIDDelete(ctx, oauth2TokenID)


指定したトークンの認可を取り消します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oauth2TokenID** | [**string**](.md)| 操作の対象となるTokenのID | 

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

