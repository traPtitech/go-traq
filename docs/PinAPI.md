# \PinAPI

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePin**](PinAPI.md#CreatePin) | **Post** /messages/{messageId}/pin | ピン留めする
[**GetChannelPins**](PinAPI.md#GetChannelPins) | **Get** /channels/{channelId}/pins | チャンネルピンのリストを取得
[**GetPin**](PinAPI.md#GetPin) | **Get** /messages/{messageId}/pin | ピン留めを取得
[**RemovePin**](PinAPI.md#RemovePin) | **Delete** /messages/{messageId}/pin | ピン留めを外す



## CreatePin

> MessagePin CreatePin(ctx, messageId).Execute()

ピン留めする



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	traq "github.com/traPtitech/go-traq"
)

func main() {
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | メッセージUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.PinAPI.CreatePin(context.Background(), messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PinAPI.CreatePin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePin`: MessagePin
	fmt.Fprintf(os.Stdout, "Response from `PinAPI.CreatePin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | メッセージUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessagePin**](MessagePin.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelPins

> []Pin GetChannelPins(ctx, channelId).Execute()

チャンネルピンのリストを取得



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	traq "github.com/traPtitech/go-traq"
)

func main() {
	channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | チャンネルUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.PinAPI.GetChannelPins(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PinAPI.GetChannelPins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelPins`: []Pin
	fmt.Fprintf(os.Stdout, "Response from `PinAPI.GetChannelPins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelPinsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Pin**](Pin.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPin

> MessagePin GetPin(ctx, messageId).Execute()

ピン留めを取得



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	traq "github.com/traPtitech/go-traq"
)

func main() {
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | メッセージUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.PinAPI.GetPin(context.Background(), messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PinAPI.GetPin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPin`: MessagePin
	fmt.Fprintf(os.Stdout, "Response from `PinAPI.GetPin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | メッセージUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessagePin**](MessagePin.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePin

> RemovePin(ctx, messageId).Execute()

ピン留めを外す



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	traq "github.com/traPtitech/go-traq"
)

func main() {
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | メッセージUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.PinAPI.RemovePin(context.Background(), messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PinAPI.RemovePin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | メッセージUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

