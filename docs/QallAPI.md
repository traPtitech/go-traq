# \QallAPI

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeParticipantRole**](QallAPI.md#ChangeParticipantRole) | **Patch** /qall/rooms/{roomId}/participants | ルームでの発言権限を変更
[**GetLiveKitToken**](QallAPI.md#GetLiveKitToken) | **Get** /qall/token | LiveKitトークンを取得
[**GetQallEndpoints**](QallAPI.md#GetQallEndpoints) | **Get** /qall/endpoints | LiveKitエンドポイントを取得
[**GetRoomMetadata**](QallAPI.md#GetRoomMetadata) | **Get** /qall/rooms/{roomId}/metadata | ルームのメタデータを取得
[**GetRooms**](QallAPI.md#GetRooms) | **Get** /qall/rooms | ルームと参加者の一覧を取得
[**GetSoundboardList**](QallAPI.md#GetSoundboardList) | **Get** /qall/soundboard | サウンドボード用の音声一覧を取得
[**LiveKitWebhook**](QallAPI.md#LiveKitWebhook) | **Post** /qall/webhook | LiveKit Webhook受信
[**PostSoundboard**](QallAPI.md#PostSoundboard) | **Post** /qall/soundboard | サウンドボード用の短い音声ファイルをアップロード
[**PostSoundboardPlay**](QallAPI.md#PostSoundboardPlay) | **Post** /qall/soundboard/play | アップロード済み音声を LiveKit ルームで再生
[**UpdateRoomMetadata**](QallAPI.md#UpdateRoomMetadata) | **Patch** /qall/rooms/{roomId}/metadata | ルームのメタデータを更新



## ChangeParticipantRole

> QallParticipantResponse ChangeParticipantRole(ctx, roomId).QallParticipantRequest(qallParticipantRequest).Execute()

ルームでの発言権限を変更



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
	roomId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ルームUUID
	qallParticipantRequest := []traq.QallParticipantRequest{*traq.NewQallParticipantRequest([]traq.QallParticipantRequestUsersInner{*traq.NewQallParticipantRequestUsersInner()})} // []QallParticipantRequest | 発言権限を変更する参加者の情報

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.QallAPI.ChangeParticipantRole(context.Background(), roomId).QallParticipantRequest(qallParticipantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QallAPI.ChangeParticipantRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeParticipantRole`: QallParticipantResponse
	fmt.Fprintf(os.Stdout, "Response from `QallAPI.ChangeParticipantRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string** | ルームUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeParticipantRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qallParticipantRequest** | [**[]QallParticipantRequest**](QallParticipantRequest.md) | 発言権限を変更する参加者の情報 | 

### Return type

[**QallParticipantResponse**](QallParticipantResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveKitToken

> QallTokenResponse GetLiveKitToken(ctx).RoomId(roomId).IsWebinar(isWebinar).Execute()

LiveKitトークンを取得



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
	roomId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ルームUUID (optional)
	isWebinar := true // bool | ウェビナールームかどうか(デフォルト false) (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.QallAPI.GetLiveKitToken(context.Background()).RoomId(roomId).IsWebinar(isWebinar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QallAPI.GetLiveKitToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveKitToken`: QallTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `QallAPI.GetLiveKitToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveKitTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomId** | **string** | ルームUUID | 
 **isWebinar** | **bool** | ウェビナールームかどうか(デフォルト false) | 

### Return type

[**QallTokenResponse**](QallTokenResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQallEndpoints

> QallEndpointResponse GetQallEndpoints(ctx).Execute()

LiveKitエンドポイントを取得



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
	resp, r, err := apiClient.QallAPI.GetQallEndpoints(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QallAPI.GetQallEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQallEndpoints`: QallEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `QallAPI.GetQallEndpoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQallEndpointsRequest struct via the builder pattern


### Return type

[**QallEndpointResponse**](QallEndpointResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomMetadata

> QallMetadataResponse GetRoomMetadata(ctx, roomId).Execute()

ルームのメタデータを取得



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
	roomId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ルームUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.QallAPI.GetRoomMetadata(context.Background(), roomId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QallAPI.GetRoomMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomMetadata`: QallMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `QallAPI.GetRoomMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string** | ルームUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QallMetadataResponse**](QallMetadataResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRooms

> []QallRoomWithParticipants GetRooms(ctx).Execute()

ルームと参加者の一覧を取得



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
	resp, r, err := apiClient.QallAPI.GetRooms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QallAPI.GetRooms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRooms`: []QallRoomWithParticipants
	fmt.Fprintf(os.Stdout, "Response from `QallAPI.GetRooms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomsRequest struct via the builder pattern


### Return type

[**[]QallRoomWithParticipants**](QallRoomWithParticipants.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoundboardList

> []SoundboardItem GetSoundboardList(ctx).Execute()

サウンドボード用の音声一覧を取得



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
	resp, r, err := apiClient.QallAPI.GetSoundboardList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QallAPI.GetSoundboardList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSoundboardList`: []SoundboardItem
	fmt.Fprintf(os.Stdout, "Response from `QallAPI.GetSoundboardList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoundboardListRequest struct via the builder pattern


### Return type

[**[]SoundboardItem**](SoundboardItem.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveKitWebhook

> LiveKitWebhook(ctx).Body(body).Execute()

LiveKit Webhook受信



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.QallAPI.LiveKitWebhook(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QallAPI.LiveKitWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLiveKitWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/webhook+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSoundboard

> SoundboardUploadResponse PostSoundboard(ctx).Audio(audio).SoundName(soundName).StampId(stampId).Execute()

サウンドボード用の短い音声ファイルをアップロード



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
	audio := os.NewFile(1234, "some_file") // *os.File | アップロードする音声ファイル(20秒以内)
	soundName := "soundName_example" // string | ユーザが自由につけるサウンド名
	stampId := "stampId_example" // string | アイコンスタンプID (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.QallAPI.PostSoundboard(context.Background()).Audio(audio).SoundName(soundName).StampId(stampId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QallAPI.PostSoundboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSoundboard`: SoundboardUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `QallAPI.PostSoundboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSoundboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audio** | ***os.File** | アップロードする音声ファイル(20秒以内) | 
 **soundName** | **string** | ユーザが自由につけるサウンド名 | 
 **stampId** | **string** | アイコンスタンプID | 

### Return type

[**SoundboardUploadResponse**](SoundboardUploadResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSoundboardPlay

> SoundboardPlayResponse PostSoundboardPlay(ctx).SoundboardPlayRequest(soundboardPlayRequest).Execute()

アップロード済み音声を LiveKit ルームで再生



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
	soundboardPlayRequest := *traq.NewSoundboardPlayRequest("SoundId_example", "RoomName_example") // SoundboardPlayRequest | 

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.QallAPI.PostSoundboardPlay(context.Background()).SoundboardPlayRequest(soundboardPlayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QallAPI.PostSoundboardPlay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSoundboardPlay`: SoundboardPlayResponse
	fmt.Fprintf(os.Stdout, "Response from `QallAPI.PostSoundboardPlay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSoundboardPlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **soundboardPlayRequest** | [**SoundboardPlayRequest**](SoundboardPlayRequest.md) |  | 

### Return type

[**SoundboardPlayResponse**](SoundboardPlayResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomMetadata

> UpdateRoomMetadata(ctx, roomId).QallMetadataRequest(qallMetadataRequest).Execute()

ルームのメタデータを更新



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
	roomId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ルームUUID
	qallMetadataRequest := *traq.NewQallMetadataRequest() // QallMetadataRequest | ルームのメタデータ

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.QallAPI.UpdateRoomMetadata(context.Background(), roomId).QallMetadataRequest(qallMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QallAPI.UpdateRoomMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string** | ルームUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoomMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qallMetadataRequest** | [**QallMetadataRequest**](QallMetadataRequest.md) | ルームのメタデータ | 

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

