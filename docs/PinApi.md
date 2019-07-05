# \PinApi

All URIs are relative to *https://q.trap.jp/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsChannelIDPinsGet**](PinApi.md#ChannelsChannelIDPinsGet) | **Get** /channels/{channelID}/pins | 
[**PinsPinIDDelete**](PinApi.md#PinsPinIDDelete) | **Delete** /pins/{pinID} | 
[**PinsPinIDGet**](PinApi.md#PinsPinIDGet) | **Get** /pins/{pinID} | 
[**PinsPost**](PinApi.md#PinsPost) | **Post** /pins | 



## ChannelsChannelIDPinsGet

> []Pin ChannelsChannelIDPinsGet(ctx, channelID)


チャンネルのピン留め一覧を取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelID** | [**string**](.md)| 操作の対象となるチャンネルのID | 

### Return type

[**[]Pin**](Pin.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinsPinIDDelete

> PinsPinIDDelete(ctx, pinID)


ピン留めを外します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pinID** | **string**| 操作の対象となるピン留めID | 

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


## PinsPinIDGet

> Pin PinsPinIDGet(ctx, pinID)


ピン留めを取得します。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pinID** | **string**| 操作の対象となるピン留めID | 

### Return type

[**Pin**](Pin.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinsPost

> InlineResponse2011 PinsPost(ctx, optional)


チャンネルにメッセージをピン留めします。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PinsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PinsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject26** | [**optional.Interface of InlineObject26**](InlineObject26.md)|  | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[traqOAuth2](../README.md#traqOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

