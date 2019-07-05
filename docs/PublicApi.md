# \PublicApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicIconUsernameGet**](PublicApi.md#PublicIconUsernameGet) | **Get** /public/icon/{username} | 



## PublicIconUsernameGet

> *os.File PublicIconUsernameGet(ctx, username)


指定したユーザーのアイコン画像を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**| 画像を取得するユーザーのユーザー名 | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, image/png, image/gif

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

