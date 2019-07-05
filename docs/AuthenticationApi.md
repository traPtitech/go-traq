# \AuthenticationApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginPost**](AuthenticationApi.md#LoginPost) | **Post** /login | 
[**LogoutPost**](AuthenticationApi.md#LogoutPost) | **Post** /logout | 



## LoginPost

> LoginPost(ctx, optional)


ログインを行います。リダイレクトパラメーターが存在する場合はログイン後にリダイレクトします

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LoginPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoginPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirect** | **optional.String**| リダイレクト先 | 
 **inlineObject** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

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


## LogoutPost

> LogoutPost(ctx, optional)


ログアウトを行います。リダイレクトパラメーターが存在する場合はログアウト後にリダイレクトします

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogoutPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LogoutPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirect** | **optional.String**| リダイレクト先 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

