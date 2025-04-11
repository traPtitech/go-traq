# \ChannelApi

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](ChannelApi.md#CreateChannel) | **Post** /channels | チャンネルを作成
[**EditChannel**](ChannelApi.md#EditChannel) | **Patch** /channels/{channelId} | チャンネル情報を変更
[**EditChannelSubscribers**](ChannelApi.md#EditChannelSubscribers) | **Patch** /channels/{channelId}/subscribers | チャンネルの通知購読者を編集
[**EditChannelTopic**](ChannelApi.md#EditChannelTopic) | **Put** /channels/{channelId}/topic | チャンネルトピックを編集
[**GetChannel**](ChannelApi.md#GetChannel) | **Get** /channels/{channelId} | チャンネル情報を取得
[**GetChannelBots**](ChannelApi.md#GetChannelBots) | **Get** /channels/{channelId}/bots | チャンネル参加中のBOTのリストを取得
[**GetChannelEvents**](ChannelApi.md#GetChannelEvents) | **Get** /channels/{channelId}/events | チャンネルイベントのリストを取得
[**GetChannelPath**](ChannelApi.md#GetChannelPath) | **Get** /channels/{channelId}/path | 指定したチャンネルパスを取得
[**GetChannelPins**](ChannelApi.md#GetChannelPins) | **Get** /channels/{channelId}/pins | チャンネルピンのリストを取得
[**GetChannelStats**](ChannelApi.md#GetChannelStats) | **Get** /channels/{channelId}/stats | チャンネル統計情報を取得
[**GetChannelSubscribers**](ChannelApi.md#GetChannelSubscribers) | **Get** /channels/{channelId}/subscribers | チャンネルの通知購読者のリストを取得
[**GetChannelTopic**](ChannelApi.md#GetChannelTopic) | **Get** /channels/{channelId}/topic | チャンネルトピックを取得
[**GetChannelViewers**](ChannelApi.md#GetChannelViewers) | **Get** /channels/{channelId}/viewers | チャンネル閲覧者リストを取得
[**GetChannels**](ChannelApi.md#GetChannels) | **Get** /channels | チャンネルリストを取得
[**GetMessages**](ChannelApi.md#GetMessages) | **Get** /channels/{channelId}/messages | チャンネルメッセージのリストを取得
[**GetUserDMChannel**](ChannelApi.md#GetUserDMChannel) | **Get** /users/{userId}/dm-channel | DMチャンネル情報を取得
[**PostMessage**](ChannelApi.md#PostMessage) | **Post** /channels/{channelId}/messages | チャンネルにメッセージを投稿
[**SetChannelSubscribers**](ChannelApi.md#SetChannelSubscribers) | **Put** /channels/{channelId}/subscribers | チャンネルの通知購読者を設定



## CreateChannel

> Channel CreateChannel(ctx).PostChannelRequest(postChannelRequest).Execute()

チャンネルを作成



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
    postChannelRequest := *traq.NewPostChannelRequest("Name_example", "Parent_example") // PostChannelRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelApi.CreateChannel(context.Background()).PostChannelRequest(postChannelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.CreateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChannel`: Channel
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.CreateChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postChannelRequest** | [**PostChannelRequest**](PostChannelRequest.md) |  | 

### Return type

[**Channel**](Channel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditChannel

> EditChannel(ctx, channelId).PatchChannelRequest(patchChannelRequest).Execute()

チャンネル情報を変更



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
    patchChannelRequest := *traq.NewPatchChannelRequest() // PatchChannelRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.ChannelApi.EditChannel(context.Background(), channelId).PatchChannelRequest(patchChannelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.EditChannel``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEditChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchChannelRequest** | [**PatchChannelRequest**](PatchChannelRequest.md) |  | 

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
    r, err := apiClient.ChannelApi.EditChannelSubscribers(context.Background(), channelId).PatchChannelSubscribersRequest(patchChannelSubscribersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.EditChannelSubscribers``: %v\n", err)
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


## EditChannelTopic

> EditChannelTopic(ctx, channelId).PutChannelTopicRequest(putChannelTopicRequest).Execute()

チャンネルトピックを編集



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
    putChannelTopicRequest := *traq.NewPutChannelTopicRequest("Topic_example") // PutChannelTopicRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.ChannelApi.EditChannelTopic(context.Background(), channelId).PutChannelTopicRequest(putChannelTopicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.EditChannelTopic``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEditChannelTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putChannelTopicRequest** | [**PutChannelTopicRequest**](PutChannelTopicRequest.md) |  | 

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


## GetChannel

> Channel GetChannel(ctx, channelId).Execute()

チャンネル情報を取得



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
    resp, r, err := apiClient.ChannelApi.GetChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannel`: Channel
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Channel**](Channel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelBots

> []BotUser GetChannelBots(ctx, channelId).Execute()

チャンネル参加中のBOTのリストを取得



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
    resp, r, err := apiClient.ChannelApi.GetChannelBots(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelBots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannelBots`: []BotUser
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetChannelBots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelBotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BotUser**](BotUser.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelEvents

> []ChannelEvent GetChannelEvents(ctx, channelId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Execute()

チャンネルイベントのリストを取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    traq "github.com/traPtitech/go-traq"
)

func main() {
    channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | チャンネルUUID
    limit := int32(50) // int32 | 取得する件数 (optional)
    offset := int32(150) // int32 | 取得するオフセット (optional) (default to 0)
    since := time.Now() // time.Time | 取得する時間範囲の開始日時 (optional) (default to "0000-01-01T00:00Z")
    until := time.Now() // time.Time | 取得する時間範囲の終了日時 (optional)
    inclusive := true // bool | 範囲の端を含めるかどうか (optional) (default to false)
    order := "order_example" // string | 昇順か降順か (optional) (default to "desc")

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelApi.GetChannelEvents(context.Background(), channelId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannelEvents`: []ChannelEvent
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetChannelEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | 取得する件数 | 
 **offset** | **int32** | 取得するオフセット | [default to 0]
 **since** | **time.Time** | 取得する時間範囲の開始日時 | [default to &quot;0000-01-01T00:00Z&quot;]
 **until** | **time.Time** | 取得する時間範囲の終了日時 | 
 **inclusive** | **bool** | 範囲の端を含めるかどうか | [default to false]
 **order** | **string** | 昇順か降順か | [default to &quot;desc&quot;]

### Return type

[**[]ChannelEvent**](ChannelEvent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelPath

> ChannelPath GetChannelPath(ctx, channelId).Execute()

指定したチャンネルパスを取得

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
    resp, r, err := apiClient.ChannelApi.GetChannelPath(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannelPath`: ChannelPath
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetChannelPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChannelPath**](ChannelPath.md)

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
    resp, r, err := apiClient.ChannelApi.GetChannelPins(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelPins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannelPins`: []Pin
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetChannelPins`: %v\n", resp)
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


## GetChannelStats

> ChannelStats GetChannelStats(ctx, channelId).ExcludeDeletedMessages(excludeDeletedMessages).Execute()

チャンネル統計情報を取得



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
    excludeDeletedMessages := true // bool | 削除されたメッセージを除外するかどうか(デフォルト false) (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelApi.GetChannelStats(context.Background(), channelId).ExcludeDeletedMessages(excludeDeletedMessages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannelStats`: ChannelStats
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetChannelStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeDeletedMessages** | **bool** | 削除されたメッセージを除外するかどうか(デフォルト false) | 

### Return type

[**ChannelStats**](ChannelStats.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    resp, r, err := apiClient.ChannelApi.GetChannelSubscribers(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannelSubscribers`: []string
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetChannelSubscribers`: %v\n", resp)
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


## GetChannelTopic

> ChannelTopic GetChannelTopic(ctx, channelId).Execute()

チャンネルトピックを取得



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
    resp, r, err := apiClient.ChannelApi.GetChannelTopic(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannelTopic`: ChannelTopic
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetChannelTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChannelTopic**](ChannelTopic.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelViewers

> []ChannelViewer GetChannelViewers(ctx, channelId).Execute()

チャンネル閲覧者リストを取得



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
    resp, r, err := apiClient.ChannelApi.GetChannelViewers(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelViewers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannelViewers`: []ChannelViewer
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetChannelViewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelViewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ChannelViewer**](ChannelViewer.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannels

> ChannelList GetChannels(ctx).IncludeDm(includeDm).Path(path).Execute()

チャンネルリストを取得



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
    includeDm := true // bool | ダイレクトメッセージチャンネルをレスポンスに含めるかどうか (optional) (default to false)
    path := "path_example" // string | パスが一致するチャンネルのみを取得する (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelApi.GetChannels(context.Background()).IncludeDm(includeDm).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannels`: ChannelList
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDm** | **bool** | ダイレクトメッセージチャンネルをレスポンスに含めるかどうか | [default to false]
 **path** | **string** | パスが一致するチャンネルのみを取得する | 

### Return type

[**ChannelList**](ChannelList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessages

> []Message GetMessages(ctx, channelId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Execute()

チャンネルメッセージのリストを取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    traq "github.com/traPtitech/go-traq"
)

func main() {
    channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | チャンネルUUID
    limit := int32(50) // int32 | 取得する件数 (optional)
    offset := int32(150) // int32 | 取得するオフセット (optional) (default to 0)
    since := time.Now() // time.Time | 取得する時間範囲の開始日時 (optional) (default to "0000-01-01T00:00Z")
    until := time.Now() // time.Time | 取得する時間範囲の終了日時 (optional)
    inclusive := true // bool | 範囲の端を含めるかどうか (optional) (default to false)
    order := "order_example" // string | 昇順か降順か (optional) (default to "desc")

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelApi.GetMessages(context.Background(), channelId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessages`: []Message
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | 取得する件数 | 
 **offset** | **int32** | 取得するオフセット | [default to 0]
 **since** | **time.Time** | 取得する時間範囲の開始日時 | [default to &quot;0000-01-01T00:00Z&quot;]
 **until** | **time.Time** | 取得する時間範囲の終了日時 | 
 **inclusive** | **bool** | 範囲の端を含めるかどうか | [default to false]
 **order** | **string** | 昇順か降順か | [default to &quot;desc&quot;]

### Return type

[**[]Message**](Message.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDMChannel

> DMChannel GetUserDMChannel(ctx, userId).Execute()

DMチャンネル情報を取得



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
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelApi.GetUserDMChannel(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetUserDMChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserDMChannel`: DMChannel
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.GetUserDMChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDMChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DMChannel**](DMChannel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMessage

> Message PostMessage(ctx, channelId).PostMessageRequest(postMessageRequest).Execute()

チャンネルにメッセージを投稿



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
    postMessageRequest := *traq.NewPostMessageRequest("Content_example") // PostMessageRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelApi.PostMessage(context.Background(), channelId).PostMessageRequest(postMessageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.PostMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMessage`: Message
    fmt.Fprintf(os.Stdout, "Response from `ChannelApi.PostMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | チャンネルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postMessageRequest** | [**PostMessageRequest**](PostMessageRequest.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    r, err := apiClient.ChannelApi.SetChannelSubscribers(context.Background(), channelId).PutChannelSubscribersRequest(putChannelSubscribersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.SetChannelSubscribers``: %v\n", err)
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

