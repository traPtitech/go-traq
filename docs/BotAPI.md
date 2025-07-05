# \BotAPI

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateBot**](BotAPI.md#ActivateBot) | **Post** /bots/{botId}/actions/activate | BOTをアクティベート
[**ChangeBotIcon**](BotAPI.md#ChangeBotIcon) | **Put** /bots/{botId}/icon | BOTのアイコン画像を変更
[**ConnectBotWS**](BotAPI.md#ConnectBotWS) | **Get** /bots/ws | WebSocket Mode BOT用通知ストリームに接続します
[**CreateBot**](BotAPI.md#CreateBot) | **Post** /bots | BOTを作成
[**DeleteBot**](BotAPI.md#DeleteBot) | **Delete** /bots/{botId} | BOTを削除
[**EditBot**](BotAPI.md#EditBot) | **Patch** /bots/{botId} | BOT情報を変更
[**GetBot**](BotAPI.md#GetBot) | **Get** /bots/{botId} | BOT情報を取得
[**GetBotIcon**](BotAPI.md#GetBotIcon) | **Get** /bots/{botId}/icon | BOTのアイコン画像を取得
[**GetBotLogs**](BotAPI.md#GetBotLogs) | **Get** /bots/{botId}/logs | BOTのイベントログを取得
[**GetBots**](BotAPI.md#GetBots) | **Get** /bots | BOTリストを取得
[**GetChannelBots**](BotAPI.md#GetChannelBots) | **Get** /channels/{channelId}/bots | チャンネル参加中のBOTのリストを取得
[**InactivateBot**](BotAPI.md#InactivateBot) | **Post** /bots/{botId}/actions/inactivate | BOTをインアクティベート
[**LetBotJoinChannel**](BotAPI.md#LetBotJoinChannel) | **Post** /bots/{botId}/actions/join | BOTをチャンネルに参加させる
[**LetBotLeaveChannel**](BotAPI.md#LetBotLeaveChannel) | **Post** /bots/{botId}/actions/leave | BOTをチャンネルから退出させる
[**ReissueBot**](BotAPI.md#ReissueBot) | **Post** /bots/{botId}/actions/reissue | BOTのトークンを再発行



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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.ActivateBot(context.Background(), botId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.ActivateBot``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
	file := os.NewFile(1234, "some_file") // *os.File | アイコン画像(2MBまでのpng, jpeg, gif)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.ChangeBotIcon(context.Background(), botId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.ChangeBotIcon``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.ConnectBotWS(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.ConnectBotWS``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	postBotRequest := *traq.NewPostBotRequest("Name_example", "DisplayName_example", "Description_example", traq.BotMode("HTTP")) // PostBotRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.BotAPI.CreateBot(context.Background()).PostBotRequest(postBotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.CreateBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBot`: BotDetail
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.CreateBot`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.DeleteBot(context.Background(), botId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.DeleteBot``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
	patchBotRequest := *traq.NewPatchBotRequest() // PatchBotRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.EditBot(context.Background(), botId).PatchBotRequest(patchBotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.EditBot``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
	detail := true // bool | 詳細情報を含めるかどうか (optional) (default to false)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.BotAPI.GetBot(context.Background(), botId).Detail(detail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.GetBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBot`: GetBot200Response
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.GetBot`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.BotAPI.GetBotIcon(context.Background(), botId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.GetBotIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBotIcon`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.GetBotIcon`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
	limit := int32(50) // int32 | 取得する件数 (optional)
	offset := int32(150) // int32 | 取得するオフセット (optional) (default to 0)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.BotAPI.GetBotLogs(context.Background(), botId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.GetBotLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBotLogs`: []BotEventLog
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.GetBotLogs`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	all := true // bool | 全てのBOTを取得するかどうか (optional) (default to false)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.BotAPI.GetBots(context.Background()).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.GetBots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBots`: []Bot
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.GetBots`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | チャンネルUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.BotAPI.GetChannelBots(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.GetChannelBots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelBots`: []BotUser
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.GetChannelBots`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.InactivateBot(context.Background(), botId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.InactivateBot``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
	postBotActionJoinRequest := *traq.NewPostBotActionJoinRequest("ChannelId_example") // PostBotActionJoinRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.LetBotJoinChannel(context.Background(), botId).PostBotActionJoinRequest(postBotActionJoinRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.LetBotJoinChannel``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID
	postBotActionLeaveRequest := *traq.NewPostBotActionLeaveRequest("ChannelId_example") // PostBotActionLeaveRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.LetBotLeaveChannel(context.Background(), botId).PostBotActionLeaveRequest(postBotActionLeaveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.LetBotLeaveChannel``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BOTUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.BotAPI.ReissueBot(context.Background(), botId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.ReissueBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReissueBot`: BotTokens
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.ReissueBot`: %v\n", resp)
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

