# \BotApi

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateBot**](BotApi.md#ActivateBot) | **Post** /bots/{botId}/actions/activate | BOTをアクティベート
[**ChangeBotIcon**](BotApi.md#ChangeBotIcon) | **Put** /bots/{botId}/icon | BOTのアイコン画像を変更
[**ConnectBotWS**](BotApi.md#ConnectBotWS) | **Get** /bots/ws | WebSocket Mode BOT用通知ストリームに接続します
[**CreateBot**](BotApi.md#CreateBot) | **Post** /bots | BOTを作成
[**DeleteBot**](BotApi.md#DeleteBot) | **Delete** /bots/{botId} | BOTを削除
[**EditBot**](BotApi.md#EditBot) | **Patch** /bots/{botId} | BOT情報を変更
[**GetBot**](BotApi.md#GetBot) | **Get** /bots/{botId} | BOT情報を取得
[**GetBotIcon**](BotApi.md#GetBotIcon) | **Get** /bots/{botId}/icon | BOTのアイコン画像を取得
[**GetBotLogs**](BotApi.md#GetBotLogs) | **Get** /bots/{botId}/logs | BOTのイベントログを取得
[**GetBots**](BotApi.md#GetBots) | **Get** /bots | BOTリストを取得
[**GetChannelBots**](BotApi.md#GetChannelBots) | **Get** /channels/{channelId}/bots | チャンネル参加中のBOTのリストを取得
[**InactivateBot**](BotApi.md#InactivateBot) | **Post** /bots/{botId}/actions/inactivate | BOTをインアクティベート
[**LetBotJoinChannel**](BotApi.md#LetBotJoinChannel) | **Post** /bots/{botId}/actions/join | BOTをチャンネルに参加させる
[**LetBotLeaveChannel**](BotApi.md#LetBotLeaveChannel) | **Post** /bots/{botId}/actions/leave | BOTをチャンネルから退出させる
[**ReissueBot**](BotApi.md#ReissueBot) | **Post** /bots/{botId}/actions/reissue | BOTのトークンを再発行



## ActivateBot

> ActivateBot(ctx, botId).Execute()

BOTをアクティベート



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BotApi.ActivateBot(context.Background(), botId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.ActivateBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateBotRequest struct via the builder pattern


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


## ChangeBotIcon

> ChangeBotIcon(ctx, botId).File(file).Execute()

BOTのアイコン画像を変更



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
    file := os.NewFile(1234, "some_file") // *os.File | アイコン画像(1MBまでのpng, jpeg, gif)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BotApi.ChangeBotIcon(context.Background(), botId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.ChangeBotIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeBotIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | アイコン画像(1MBまでのpng, jpeg, gif) | 

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


## ConnectBotWS

> ConnectBotWS(ctx).Execute()

WebSocket Mode BOT用通知ストリームに接続します



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BotApi.ConnectBotWS(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.ConnectBotWS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConnectBotWSRequest struct via the builder pattern


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


## CreateBot

> BotDetail CreateBot(ctx).PostBotRequest(postBotRequest).Execute()

BOTを作成



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    postBotRequest := *openapiclient.NewPostBotRequest("Name_example", "DisplayName_example", "Description_example", openapiclient.BotMode("HTTP")) // PostBotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.CreateBot(context.Background()).PostBotRequest(postBotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.CreateBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBot`: BotDetail
    fmt.Fprintf(os.Stdout, "Response from `BotApi.CreateBot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postBotRequest** | [**PostBotRequest**](PostBotRequest.md) |  | 

### Return type

[**BotDetail**](BotDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBot

> DeleteBot(ctx, botId).Execute()

BOTを削除



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BotApi.DeleteBot(context.Background(), botId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.DeleteBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBotRequest struct via the builder pattern


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


## EditBot

> EditBot(ctx, botId).PatchBotRequest(patchBotRequest).Execute()

BOT情報を変更



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
    patchBotRequest := *openapiclient.NewPatchBotRequest() // PatchBotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BotApi.EditBot(context.Background(), botId).PatchBotRequest(patchBotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.EditBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchBotRequest** | [**PatchBotRequest**](PatchBotRequest.md) |  | 

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


## GetBot

> GetBot200Response GetBot(ctx, botId).Detail(detail).Execute()

BOT情報を取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
    detail := true // bool | 詳細情報を含めるかどうか (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.GetBot(context.Background(), botId).Detail(detail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.GetBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBot`: GetBot200Response
    fmt.Fprintf(os.Stdout, "Response from `BotApi.GetBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detail** | **bool** | 詳細情報を含めるかどうか | [default to false]

### Return type

[**GetBot200Response**](GetBot200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBotIcon

> *os.File GetBotIcon(ctx, botId).Execute()

BOTのアイコン画像を取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.GetBotIcon(context.Background(), botId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.GetBotIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBotIcon`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `BotApi.GetBotIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotIconRequest struct via the builder pattern


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


## GetBotLogs

> []BotEventLog GetBotLogs(ctx, botId).Limit(limit).Offset(offset).Execute()

BOTのイベントログを取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
    limit := int32(50) // int32 | 取得する件数 (optional)
    offset := int32(150) // int32 | 取得するオフセット (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.GetBotLogs(context.Background(), botId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.GetBotLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBotLogs`: []BotEventLog
    fmt.Fprintf(os.Stdout, "Response from `BotApi.GetBotLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | 取得する件数 | 
 **offset** | **int32** | 取得するオフセット | [default to 0]

### Return type

[**[]BotEventLog**](BotEventLog.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBots

> []Bot GetBots(ctx).All(all).Execute()

BOTリストを取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    all := true // bool | 全てのBOTを取得するかどうか (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.GetBots(context.Background()).All(all).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.GetBots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBots`: []Bot
    fmt.Fprintf(os.Stdout, "Response from `BotApi.GetBots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** | 全てのBOTを取得するかどうか | [default to false]

### Return type

[**[]Bot**](Bot.md)

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | チャンネルUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.GetChannelBots(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.GetChannelBots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannelBots`: []BotUser
    fmt.Fprintf(os.Stdout, "Response from `BotApi.GetChannelBots`: %v\n", resp)
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


## InactivateBot

> InactivateBot(ctx, botId).Execute()

BOTをインアクティベート



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BotApi.InactivateBot(context.Background(), botId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.InactivateBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInactivateBotRequest struct via the builder pattern


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


## LetBotJoinChannel

> LetBotJoinChannel(ctx, botId).PostBotActionJoinRequest(postBotActionJoinRequest).Execute()

BOTをチャンネルに参加させる



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
    postBotActionJoinRequest := *openapiclient.NewPostBotActionJoinRequest("ChannelId_example") // PostBotActionJoinRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BotApi.LetBotJoinChannel(context.Background(), botId).PostBotActionJoinRequest(postBotActionJoinRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.LetBotJoinChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLetBotJoinChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postBotActionJoinRequest** | [**PostBotActionJoinRequest**](PostBotActionJoinRequest.md) |  | 

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


## LetBotLeaveChannel

> LetBotLeaveChannel(ctx, botId).PostBotActionLeaveRequest(postBotActionLeaveRequest).Execute()

BOTをチャンネルから退出させる



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
    postBotActionLeaveRequest := *openapiclient.NewPostBotActionLeaveRequest("ChannelId_example") // PostBotActionLeaveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BotApi.LetBotLeaveChannel(context.Background(), botId).PostBotActionLeaveRequest(postBotActionLeaveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.LetBotLeaveChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLetBotLeaveChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postBotActionLeaveRequest** | [**PostBotActionLeaveRequest**](PostBotActionLeaveRequest.md) |  | 

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


## ReissueBot

> BotTokens ReissueBot(ctx, botId).Execute()

BOTのトークンを再発行



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.ReissueBot(context.Background(), botId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.ReissueBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReissueBot`: BotTokens
    fmt.Fprintf(os.Stdout, "Response from `BotApi.ReissueBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | BOTUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReissueBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotTokens**](BotTokens.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

