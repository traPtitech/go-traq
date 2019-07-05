# \SessionsApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMeSessionsDelete**](SessionsApi.md#UsersMeSessionsDelete) | **Delete** /users/me/sessions | 
[**UsersMeSessionsGet**](SessionsApi.md#UsersMeSessionsGet) | **Get** /users/me/sessions | 
[**UsersMeSessionsReferenceIDDelete**](SessionsApi.md#UsersMeSessionsReferenceIDDelete) | **Delete** /users/me/sessions/{referenceID} | 



## UsersMeSessionsDelete

> UsersMeSessionsDelete(ctx, )


自分のログインセッションを全てログアウトします。

### Required Parameters

This endpoint does not need any parameter.

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


## UsersMeSessionsGet

> []InlineResponse2002 UsersMeSessionsGet(ctx, )


自分のログインセッションリストを取得します。

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2002**](inline_response_200_2.md)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeSessionsReferenceIDDelete

> UsersMeSessionsReferenceIDDelete(ctx, referenceID)


対象のセッションをログアウトします。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceID** | **string**| 操作の対象となるセッションの参照ID | 

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

