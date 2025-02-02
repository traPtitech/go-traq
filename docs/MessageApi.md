# \MessageApi

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMessageStamp**](MessageApi.md#AddMessageStamp) | **Post** /messages/{messageId}/stamps/{stampId} | スタンプを押す
[**CreatePin**](MessageApi.md#CreatePin) | **Post** /messages/{messageId}/pin | ピン留めする
[**DeleteMessage**](MessageApi.md#DeleteMessage) | **Delete** /messages/{messageId} | メッセージを削除
[**EditMessage**](MessageApi.md#EditMessage) | **Put** /messages/{messageId} | メッセージを編集
[**GetDirectMessages**](MessageApi.md#GetDirectMessages) | **Get** /users/{userId}/messages | ダイレクトメッセージのリストを取得
[**GetMessage**](MessageApi.md#GetMessage) | **Get** /messages/{messageId} | メッセージを取得
[**GetMessageClips**](MessageApi.md#GetMessageClips) | **Get** /messages/{messageId}/clips | 自分のクリップを取得
[**GetMessageStamps**](MessageApi.md#GetMessageStamps) | **Get** /messages/{messageId}/stamps | メッセージのスタンプリストを取得
[**GetMessages**](MessageApi.md#GetMessages) | **Get** /channels/{channelId}/messages | チャンネルメッセージのリストを取得
[**GetPin**](MessageApi.md#GetPin) | **Get** /messages/{messageId}/pin | ピン留めを取得
[**PostDirectMessage**](MessageApi.md#PostDirectMessage) | **Post** /users/{userId}/messages | ダイレクトメッセージを送信
[**PostMessage**](MessageApi.md#PostMessage) | **Post** /channels/{channelId}/messages | チャンネルにメッセージを投稿
[**RemoveMessageStamp**](MessageApi.md#RemoveMessageStamp) | **Delete** /messages/{messageId}/stamps/{stampId} | スタンプを消す
[**RemovePin**](MessageApi.md#RemovePin) | **Delete** /messages/{messageId}/pin | ピン留めを外す
[**SearchMessages**](MessageApi.md#SearchMessages) | **Get** /messages | メッセージを検索



## AddMessageStamp

> AddMessageStamp(ctx, messageId, stampId).PostMessageStampRequest(postMessageStampRequest).Execute()

スタンプを押す



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
    stampId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプUUID
    postMessageStampRequest := *traq.NewPostMessageStampRequest(int32(123)) // PostMessageStampRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.MessageApi.AddMessageStamp(context.Background(), messageId, stampId).PostMessageStampRequest(postMessageStampRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.AddMessageStamp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | メッセージUUID | 
**stampId** | **string** | スタンプUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMessageStampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postMessageStampRequest** | [**PostMessageStampRequest**](PostMessageStampRequest.md) |  | 

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
    resp, r, err := apiClient.MessageApi.CreatePin(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.CreatePin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePin`: MessagePin
    fmt.Fprintf(os.Stdout, "Response from `MessageApi.CreatePin`: %v\n", resp)
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


## DeleteMessage

> DeleteMessage(ctx, messageId).Execute()

メッセージを削除



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
    r, err := apiClient.MessageApi.DeleteMessage(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.DeleteMessage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMessageRequest struct via the builder pattern


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


## EditMessage

> EditMessage(ctx, messageId).PostMessageRequest(postMessageRequest).Execute()

メッセージを編集



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
    postMessageRequest := *traq.NewPostMessageRequest("Content_example") // PostMessageRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.MessageApi.EditMessage(context.Background(), messageId).PostMessageRequest(postMessageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.EditMessage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEditMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postMessageRequest** | [**PostMessageRequest**](PostMessageRequest.md) |  | 

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


## GetDirectMessages

> []Message GetDirectMessages(ctx, userId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Execute()

ダイレクトメッセージのリストを取得



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
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID
    limit := int32(50) // int32 | 取得する件数 (optional)
    offset := int32(150) // int32 | 取得するオフセット (optional) (default to 0)
    since := time.Now() // time.Time | 取得する時間範囲の開始日時 (optional) (default to "0000-01-01T00:00Z")
    until := time.Now() // time.Time | 取得する時間範囲の終了日時 (optional)
    inclusive := true // bool | 範囲の端を含めるかどうか (optional) (default to false)
    order := "order_example" // string | 昇順か降順か (optional) (default to "desc")

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageApi.GetDirectMessages(context.Background(), userId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.GetDirectMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirectMessages`: []Message
    fmt.Fprintf(os.Stdout, "Response from `MessageApi.GetDirectMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectMessagesRequest struct via the builder pattern


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


## GetMessage

> Message GetMessage(ctx, messageId).Execute()

メッセージを取得



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
    resp, r, err := apiClient.MessageApi.GetMessage(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.GetMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessage`: Message
    fmt.Fprintf(os.Stdout, "Response from `MessageApi.GetMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | メッセージUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Message**](Message.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageClips

> []MessageClip GetMessageClips(ctx, messageId).Execute()

自分のクリップを取得



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
    resp, r, err := apiClient.MessageApi.GetMessageClips(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.GetMessageClips``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageClips`: []MessageClip
    fmt.Fprintf(os.Stdout, "Response from `MessageApi.GetMessageClips`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | メッセージUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageClipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MessageClip**](MessageClip.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageStamps

> []MessageStamp GetMessageStamps(ctx, messageId).Execute()

メッセージのスタンプリストを取得



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
    resp, r, err := apiClient.MessageApi.GetMessageStamps(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.GetMessageStamps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageStamps`: []MessageStamp
    fmt.Fprintf(os.Stdout, "Response from `MessageApi.GetMessageStamps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | メッセージUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageStampsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MessageStamp**](MessageStamp.md)

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
    resp, r, err := apiClient.MessageApi.GetMessages(context.Background(), channelId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.GetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessages`: []Message
    fmt.Fprintf(os.Stdout, "Response from `MessageApi.GetMessages`: %v\n", resp)
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
    resp, r, err := apiClient.MessageApi.GetPin(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.GetPin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPin`: MessagePin
    fmt.Fprintf(os.Stdout, "Response from `MessageApi.GetPin`: %v\n", resp)
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


## PostDirectMessage

> Message PostDirectMessage(ctx, userId).PostMessageRequest(postMessageRequest).Execute()

ダイレクトメッセージを送信



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
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID
    postMessageRequest := *traq.NewPostMessageRequest("Content_example") // PostMessageRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageApi.PostDirectMessage(context.Background(), userId).PostMessageRequest(postMessageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.PostDirectMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDirectMessage`: Message
    fmt.Fprintf(os.Stdout, "Response from `MessageApi.PostDirectMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDirectMessageRequest struct via the builder pattern


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
    resp, r, err := apiClient.MessageApi.PostMessage(context.Background(), channelId).PostMessageRequest(postMessageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.PostMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMessage`: Message
    fmt.Fprintf(os.Stdout, "Response from `MessageApi.PostMessage`: %v\n", resp)
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


## RemoveMessageStamp

> RemoveMessageStamp(ctx, messageId, stampId).Execute()

スタンプを消す



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
    stampId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプUUID

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.MessageApi.RemoveMessageStamp(context.Background(), messageId, stampId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.RemoveMessageStamp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | メッセージUUID | 
**stampId** | **string** | スタンプUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMessageStampRequest struct via the builder pattern


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
    r, err := apiClient.MessageApi.RemovePin(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.RemovePin``: %v\n", err)
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


## SearchMessages

> MessageSearchResult SearchMessages(ctx).Word(word).After(after).Before(before).In(in).To(to).From(from).Citation(citation).Bot(bot).HasURL(hasURL).HasAttachments(hasAttachments).HasImage(hasImage).HasVideo(hasVideo).HasAudio(hasAudio).Limit(limit).Offset(offset).Sort(sort).Execute()

メッセージを検索



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
    word := ""phrase match" +(foo | bar) -baz" // string | 検索ワード Simple-Query-String-Syntaxをパースして検索します  (optional)
    after := time.Now() // time.Time | 投稿日時が指定日時より後 (optional)
    before := time.Now() // time.Time | 投稿日時が指定日時より前 (optional)
    in := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | メッセージが投稿されたチャンネル (optional)
    to := []string{"Inner_example"} // []string | メンションされたユーザー (optional)
    from := []string{"Inner_example"} // []string | メッセージを投稿したユーザー (optional)
    citation := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 引用しているメッセージ (optional)
    bot := true // bool | メッセージを投稿したユーザーがBotかどうか (optional)
    hasURL := true // bool | メッセージがURLを含むか (optional)
    hasAttachments := true // bool | メッセージが添付ファイルを含むか (optional)
    hasImage := true // bool | メッセージが画像を含むか (optional)
    hasVideo := true // bool | メッセージが動画を含むか (optional)
    hasAudio := true // bool | メッセージが音声ファイルを含むか (optional)
    limit := int32(56) // int32 | 検索結果から取得するメッセージの最大件数 (optional)
    offset := int32(56) // int32 | 検索結果から取得するメッセージのオフセット (optional)
    sort := "sort_example" // string | ソート順 (作成日時が新しい `createdAt`, 作成日時が古い `-createdAt`, 更新日時が新しい `updatedAt`, 更新日時が古い `-updatedAt`) (optional) (default to "-createdAt")

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageApi.SearchMessages(context.Background()).Word(word).After(after).Before(before).In(in).To(to).From(from).Citation(citation).Bot(bot).HasURL(hasURL).HasAttachments(hasAttachments).HasImage(hasImage).HasVideo(hasVideo).HasAudio(hasAudio).Limit(limit).Offset(offset).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageApi.SearchMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMessages`: MessageSearchResult
    fmt.Fprintf(os.Stdout, "Response from `MessageApi.SearchMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **word** | **string** | 検索ワード Simple-Query-String-Syntaxをパースして検索します  | 
 **after** | **time.Time** | 投稿日時が指定日時より後 | 
 **before** | **time.Time** | 投稿日時が指定日時より前 | 
 **in** | **string** | メッセージが投稿されたチャンネル | 
 **to** | **[]string** | メンションされたユーザー | 
 **from** | **[]string** | メッセージを投稿したユーザー | 
 **citation** | **string** | 引用しているメッセージ | 
 **bot** | **bool** | メッセージを投稿したユーザーがBotかどうか | 
 **hasURL** | **bool** | メッセージがURLを含むか | 
 **hasAttachments** | **bool** | メッセージが添付ファイルを含むか | 
 **hasImage** | **bool** | メッセージが画像を含むか | 
 **hasVideo** | **bool** | メッセージが動画を含むか | 
 **hasAudio** | **bool** | メッセージが音声ファイルを含むか | 
 **limit** | **int32** | 検索結果から取得するメッセージの最大件数 | 
 **offset** | **int32** | 検索結果から取得するメッセージのオフセット | 
 **sort** | **string** | ソート順 (作成日時が新しい &#x60;createdAt&#x60;, 作成日時が古い &#x60;-createdAt&#x60;, 更新日時が新しい &#x60;updatedAt&#x60;, 更新日時が古い &#x60;-updatedAt&#x60;) | [default to &quot;-createdAt&quot;]

### Return type

[**MessageSearchResult**](MessageSearchResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

