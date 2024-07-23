# \WebhookApi

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeWebhookIcon**](WebhookApi.md#ChangeWebhookIcon) | **Put** /webhooks/{webhookId}/icon | Webhookのアイコンを変更
[**CreateWebhook**](WebhookApi.md#CreateWebhook) | **Post** /webhooks | Webhookを新規作成
[**DeleteWebhook**](WebhookApi.md#DeleteWebhook) | **Delete** /webhooks/{webhookId} | Webhookを削除
[**EditWebhook**](WebhookApi.md#EditWebhook) | **Patch** /webhooks/{webhookId} | Webhook情報を変更
[**GetWebhook**](WebhookApi.md#GetWebhook) | **Get** /webhooks/{webhookId} | Webhook情報を取得
[**GetWebhookIcon**](WebhookApi.md#GetWebhookIcon) | **Get** /webhooks/{webhookId}/icon | Webhookのアイコンを取得
[**GetWebhookMessages**](WebhookApi.md#GetWebhookMessages) | **Get** /webhooks/{webhookId}/messages | Webhookの投稿メッセージのリストを取得
[**GetWebhooks**](WebhookApi.md#GetWebhooks) | **Get** /webhooks | Webhook情報のリストを取得します
[**PostWebhook**](WebhookApi.md#PostWebhook) | **Post** /webhooks/{webhookId} | Webhookを送信



## ChangeWebhookIcon

> ChangeWebhookIcon(ctx, webhookId).File(file).Execute()

Webhookのアイコンを変更



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | WebhookUUID
    file := os.NewFile(1234, "some_file") // *os.File | アイコン画像(2MBまでのpng, jpeg, gif)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.WebhookApi.ChangeWebhookIcon(context.Background(), webhookId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.ChangeWebhookIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | WebhookUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeWebhookIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | アイコン画像(2MBまでのpng, jpeg, gif) | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> Webhook CreateWebhook(ctx).PostWebhookRequest(postWebhookRequest).Execute()

Webhookを新規作成



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
    postWebhookRequest := *traq.NewPostWebhookRequest("Name_example", "Description_example", "ChannelId_example", "Secret_example") // PostWebhookRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.CreateWebhook(context.Background()).PostWebhookRequest(postWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postWebhookRequest** | [**PostWebhookRequest**](PostWebhookRequest.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, webhookId).Execute()

Webhookを削除



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | WebhookUUID

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.WebhookApi.DeleteWebhook(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.DeleteWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | WebhookUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


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


## EditWebhook

> EditWebhook(ctx, webhookId).PatchWebhookRequest(patchWebhookRequest).Execute()

Webhook情報を変更



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | WebhookUUID
    patchWebhookRequest := *traq.NewPatchWebhookRequest() // PatchWebhookRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.WebhookApi.EditWebhook(context.Background(), webhookId).PatchWebhookRequest(patchWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.EditWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | WebhookUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchWebhookRequest** | [**PatchWebhookRequest**](PatchWebhookRequest.md) |  | 

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


## GetWebhook

> Webhook GetWebhook(ctx, webhookId).Execute()

Webhook情報を取得



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | WebhookUUID

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetWebhook(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | WebhookUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookIcon

> *os.File GetWebhookIcon(ctx, webhookId).Execute()

Webhookのアイコンを取得



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | WebhookUUID

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetWebhookIcon(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhookIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookIcon`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhookIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | WebhookUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, image/gif, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookMessages

> []Message GetWebhookMessages(ctx, webhookId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Execute()

Webhookの投稿メッセージのリストを取得



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | WebhookUUID
    limit := int32(50) // int32 | 取得する件数 (optional)
    offset := int32(150) // int32 | 取得するオフセット (optional) (default to 0)
    since := time.Now() // time.Time | 取得する時間範囲の開始日時 (optional) (default to "0000-01-01T00:00Z")
    until := time.Now() // time.Time | 取得する時間範囲の終了日時 (optional)
    inclusive := true // bool | 範囲の端を含めるかどうか (optional) (default to false)
    order := "order_example" // string | 昇順か降順か (optional) (default to "desc")

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetWebhookMessages(context.Background(), webhookId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhookMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookMessages`: []Message
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhookMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | WebhookUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookMessagesRequest struct via the builder pattern


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


## GetWebhooks

> []Webhook GetWebhooks(ctx).All(all).Execute()

Webhook情報のリストを取得します



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
    all := true // bool | 全てのWebhookを取得します。権限が必要です。 (optional) (default to false)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetWebhooks(context.Background()).All(all).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooks`: []Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** | 全てのWebhookを取得します。権限が必要です。 | [default to false]

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhook

> PostWebhook(ctx, webhookId).XTRAQSignature(xTRAQSignature).XTRAQChannelId(xTRAQChannelId).Embed(embed).Body(body).Execute()

Webhookを送信



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | WebhookUUID
    xTRAQSignature := "xTRAQSignature_example" // string | リクエストボディシグネチャ(Secretが設定されている場合は必須) (optional)
    xTRAQChannelId := "xTRAQChannelId_example" // string | 投稿先のチャンネルID(変更する場合) (optional)
    embed := int32(56) // int32 | メンション・チャンネルリンクを自動埋め込みする場合に1を指定する (optional) (default to 0)
    body := "body_example" // string |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.WebhookApi.PostWebhook(context.Background(), webhookId).XTRAQSignature(xTRAQSignature).XTRAQChannelId(xTRAQChannelId).Embed(embed).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.PostWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | WebhookUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTRAQSignature** | **string** | リクエストボディシグネチャ(Secretが設定されている場合は必須) | 
 **xTRAQChannelId** | **string** | 投稿先のチャンネルID(変更する場合) | 
 **embed** | **int32** | メンション・チャンネルリンクを自動埋め込みする場合に1を指定する | [default to 0]
 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

