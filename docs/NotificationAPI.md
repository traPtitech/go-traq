# \NotificationAPI

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditChannelSubscribers**](NotificationAPI.md#EditChannelSubscribers) | **Patch** /channels/{channelId}/subscribers | チャンネルの通知購読者を編集
[**GetChannelSubscribers**](NotificationAPI.md#GetChannelSubscribers) | **Get** /channels/{channelId}/subscribers | チャンネルの通知購読者のリストを取得
[**GetMyChannelSubscriptions**](NotificationAPI.md#GetMyChannelSubscriptions) | **Get** /users/me/subscriptions | 自分のチャンネル購読状態を取得
[**GetMyUnreadChannels**](NotificationAPI.md#GetMyUnreadChannels) | **Get** /users/me/unread | 未読チャンネルを取得
[**GetMyViewStates**](NotificationAPI.md#GetMyViewStates) | **Get** /users/me/view-states | 自身のチャンネル閲覧状態一覧を取得
[**ReadChannel**](NotificationAPI.md#ReadChannel) | **Delete** /users/me/unread/{channelId} | チャンネルを既読にする
[**RegisterFCMDevice**](NotificationAPI.md#RegisterFCMDevice) | **Post** /users/me/fcm-device | FCMデバイスを登録
[**SetChannelSubscribeLevel**](NotificationAPI.md#SetChannelSubscribeLevel) | **Put** /users/me/subscriptions/{channelId} | チャンネル購読レベルを設定
[**SetChannelSubscribers**](NotificationAPI.md#SetChannelSubscribers) | **Put** /channels/{channelId}/subscribers | チャンネルの通知購読者を設定
[**Ws**](NotificationAPI.md#Ws) | **Get** /ws | WebSocket通知ストリームに接続します



## EditChannelSubscribers

> EditChannelSubscribers(ctx, channelId).PatchChannelSubscribersRequest(patchChannelSubscribersRequest).Execute()

チャンネルの通知購読者を編集



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
	patchChannelSubscribersRequest := *traq.NewPatchChannelSubscribersRequest() // PatchChannelSubscribersRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.NotificationAPI.EditChannelSubscribers(context.Background(), channelId).PatchChannelSubscribersRequest(patchChannelSubscribersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.EditChannelSubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditChannelSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchChannelSubscribersRequest** | [**PatchChannelSubscribersRequest**](PatchChannelSubscribersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelSubscribers

> []string GetChannelSubscribers(ctx, channelId).Execute()

チャンネルの通知購読者のリストを取得



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
	resp, r, err := apiClient.NotificationAPI.GetChannelSubscribers(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.GetChannelSubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelSubscribers`: []string
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.GetChannelSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyChannelSubscriptions

> []UserSubscribeState GetMyChannelSubscriptions(ctx).Execute()

自分のチャンネル購読状態を取得



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

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationAPI.GetMyChannelSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.GetMyChannelSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyChannelSubscriptions`: []UserSubscribeState
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.GetMyChannelSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyChannelSubscriptionsRequest struct via the builder pattern


### Return type

[**[]UserSubscribeState**](UserSubscribeState.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyUnreadChannels

> []UnreadChannel GetMyUnreadChannels(ctx).Execute()

未読チャンネルを取得



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

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationAPI.GetMyUnreadChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.GetMyUnreadChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyUnreadChannels`: []UnreadChannel
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.GetMyUnreadChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyUnreadChannelsRequest struct via the builder pattern


### Return type

[**[]UnreadChannel**](UnreadChannel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyViewStates

> []MyChannelViewState GetMyViewStates(ctx).Execute()

自身のチャンネル閲覧状態一覧を取得



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

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationAPI.GetMyViewStates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.GetMyViewStates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyViewStates`: []MyChannelViewState
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.GetMyViewStates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyViewStatesRequest struct via the builder pattern


### Return type

[**[]MyChannelViewState**](MyChannelViewState.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadChannel

> ReadChannel(ctx, channelId).Execute()

チャンネルを既読にする



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
	r, err := apiClient.NotificationAPI.ReadChannel(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.ReadChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadChannelRequest struct via the builder pattern


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


## RegisterFCMDevice

> RegisterFCMDevice(ctx).PostMyFCMDeviceRequest(postMyFCMDeviceRequest).Execute()

FCMデバイスを登録



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
	postMyFCMDeviceRequest := *traq.NewPostMyFCMDeviceRequest("bk3RNwTe3H0:CI2k_HHwgIpoDKCIZvvDMExUdFQ3P1") // PostMyFCMDeviceRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.NotificationAPI.RegisterFCMDevice(context.Background()).PostMyFCMDeviceRequest(postMyFCMDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.RegisterFCMDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterFCMDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postMyFCMDeviceRequest** | [**PostMyFCMDeviceRequest**](PostMyFCMDeviceRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetChannelSubscribeLevel

> SetChannelSubscribeLevel(ctx, channelId).PutChannelSubscribeLevelRequest(putChannelSubscribeLevelRequest).Execute()

チャンネル購読レベルを設定



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
	putChannelSubscribeLevelRequest := *traq.NewPutChannelSubscribeLevelRequest(traq.ChannelSubscribeLevel(0)) // PutChannelSubscribeLevelRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.NotificationAPI.SetChannelSubscribeLevel(context.Background(), channelId).PutChannelSubscribeLevelRequest(putChannelSubscribeLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.SetChannelSubscribeLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetChannelSubscribeLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putChannelSubscribeLevelRequest** | [**PutChannelSubscribeLevelRequest**](PutChannelSubscribeLevelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetChannelSubscribers

> SetChannelSubscribers(ctx, channelId).PutChannelSubscribersRequest(putChannelSubscribersRequest).Execute()

チャンネルの通知購読者を設定



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
	putChannelSubscribersRequest := *traq.NewPutChannelSubscribersRequest([]string{"On_example"}) // PutChannelSubscribersRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.NotificationAPI.SetChannelSubscribers(context.Background(), channelId).PutChannelSubscribersRequest(putChannelSubscribersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.SetChannelSubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetChannelSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putChannelSubscribersRequest** | [**PutChannelSubscribersRequest**](PutChannelSubscribersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ws

> Ws(ctx).Execute()

WebSocket通知ストリームに接続します



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

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.NotificationAPI.Ws(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.Ws``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWsRequest struct via the builder pattern


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

