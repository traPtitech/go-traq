# \AuthorizationApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2AuthorizeDecidePost**](AuthorizationApi.md#Oauth2AuthorizeDecidePost) | **Post** /oauth2/authorize/decide | 
[**Oauth2AuthorizeGet**](AuthorizationApi.md#Oauth2AuthorizeGet) | **Get** /oauth2/authorize | 
[**Oauth2AuthorizePost**](AuthorizationApi.md#Oauth2AuthorizePost) | **Post** /oauth2/authorize | 
[**Oauth2TokenPost**](AuthorizationApi.md#Oauth2TokenPost) | **Post** /oauth2/token | 



## Oauth2AuthorizeDecidePost

> Oauth2AuthorizeDecidePost(ctx, submit)


OAuth2 認可承諾

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submit** | **string**| 承諾する場合は\\\&quot;approve\\\&quot; | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2AuthorizeGet

> Oauth2AuthorizeGet(ctx, optional)


OAuth2 認可エンドポイント

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Oauth2AuthorizeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Oauth2AuthorizeGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseType** | **optional.String**|  | 
 **clientId** | **optional.String**|  | 
 **redirectUri** | **optional.String**|  | 
 **scope** | **optional.String**|  | 
 **state** | **optional.String**|  | 
 **codeChallenge** | **optional.String**|  | 
 **codeChallengeMethod** | **optional.String**|  | 
 **nonce** | **optional.String**|  | 
 **prompt** | **optional.String**|  | 

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


## Oauth2AuthorizePost

> Oauth2AuthorizePost(ctx, optional)


OAuth2 認可エンドポイント

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Oauth2AuthorizePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Oauth2AuthorizePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseType** | **optional.String**|  | 
 **clientId** | **optional.String**|  | 
 **redirectUri** | **optional.String**|  | 
 **scope** | **optional.String**|  | 
 **state** | **optional.String**|  | 
 **codeChallenge** | **optional.String**|  | 
 **codeChallengeMethod** | **optional.String**|  | 
 **nonce** | **optional.String**|  | 
 **prompt** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2TokenPost

> InlineResponse200 Oauth2TokenPost(ctx, grantType, optional)


OAuth2 トークンエンドポイント

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantType** | **string**|  | 
 **optional** | ***Oauth2TokenPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Oauth2TokenPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **optional.String**|  | 
 **redirectUri** | **optional.String**|  | 
 **clientId** | **optional.String**|  | 
 **codeVerifier** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **password** | **optional.String**|  | 
 **scope** | **optional.String**|  | 
 **refreshToken** | **optional.String**|  | 
 **clientSecret** | **optional.String**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

