# \StampAPI

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMessageStamp**](StampAPI.md#AddMessageStamp) | **Post** /messages/{messageId}/stamps/{stampId} | スタンプを押す
[**ChangeStampImage**](StampAPI.md#ChangeStampImage) | **Put** /stamps/{stampId}/image | スタンプ画像を変更
[**CreateStamp**](StampAPI.md#CreateStamp) | **Post** /stamps | スタンプを作成
[**CreateStampPalette**](StampAPI.md#CreateStampPalette) | **Post** /stamp-palettes | スタンプパレットを作成
[**DeleteStamp**](StampAPI.md#DeleteStamp) | **Delete** /stamps/{stampId} | スタンプを削除
[**DeleteStampPalette**](StampAPI.md#DeleteStampPalette) | **Delete** /stamp-palettes/{paletteId} | スタンプパレットを削除
[**EditStamp**](StampAPI.md#EditStamp) | **Patch** /stamps/{stampId} | スタンプ情報を変更
[**EditStampPalette**](StampAPI.md#EditStampPalette) | **Patch** /stamp-palettes/{paletteId} | スタンプパレットを編集
[**GetMessageStamps**](StampAPI.md#GetMessageStamps) | **Get** /messages/{messageId}/stamps | メッセージのスタンプリストを取得
[**GetMyStampHistory**](StampAPI.md#GetMyStampHistory) | **Get** /users/me/stamp-history | スタンプ履歴を取得
[**GetStamp**](StampAPI.md#GetStamp) | **Get** /stamps/{stampId} | スタンプ情報を取得
[**GetStampImage**](StampAPI.md#GetStampImage) | **Get** /stamps/{stampId}/image | スタンプ画像を取得
[**GetStampPalette**](StampAPI.md#GetStampPalette) | **Get** /stamp-palettes/{paletteId} | スタンプパレットを取得
[**GetStampPalettes**](StampAPI.md#GetStampPalettes) | **Get** /stamp-palettes | スタンプパレットのリストを取得
[**GetStampStats**](StampAPI.md#GetStampStats) | **Get** /stamps/{stampId}/stats | スタンプ統計情報を取得
[**GetStamps**](StampAPI.md#GetStamps) | **Get** /stamps | スタンプリストを取得
[**RemoveMessageStamp**](StampAPI.md#RemoveMessageStamp) | **Delete** /messages/{messageId}/stamps/{stampId} | スタンプを消す



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
	r, err := apiClient.StampAPI.AddMessageStamp(context.Background(), messageId, stampId).PostMessageStampRequest(postMessageStampRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.AddMessageStamp``: %v\n", err)
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


## ChangeStampImage

> ChangeStampImage(ctx, stampId).File(file).Execute()

スタンプ画像を変更



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
	stampId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプUUID
	file := os.NewFile(1234, "some_file") // *os.File | スタンプ画像(1MBまでのpng, jpeg, gif)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.StampAPI.ChangeStampImage(context.Background(), stampId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.ChangeStampImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stampId** | **string** | スタンプUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeStampImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | スタンプ画像(1MBまでのpng, jpeg, gif) | 

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


## CreateStamp

> Stamp CreateStamp(ctx).Name(name).File(file).Execute()

スタンプを作成



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
	name := "name_example" // string | スタンプ名
	file := os.NewFile(1234, "some_file") // *os.File | スタンプ画像(1MBまでのpng, jpeg, gif)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.StampAPI.CreateStamp(context.Background()).Name(name).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.CreateStamp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStamp`: Stamp
	fmt.Fprintf(os.Stdout, "Response from `StampAPI.CreateStamp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | スタンプ名 | 
 **file** | ***os.File** | スタンプ画像(1MBまでのpng, jpeg, gif) | 

### Return type

[**Stamp**](Stamp.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStampPalette

> StampPalette CreateStampPalette(ctx).PostStampPaletteRequest(postStampPaletteRequest).Execute()

スタンプパレットを作成



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
	postStampPaletteRequest := *traq.NewPostStampPaletteRequest([]string{"Stamps_example"}, "Name_example", "Description_example") // PostStampPaletteRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.StampAPI.CreateStampPalette(context.Background()).PostStampPaletteRequest(postStampPaletteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.CreateStampPalette``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStampPalette`: StampPalette
	fmt.Fprintf(os.Stdout, "Response from `StampAPI.CreateStampPalette`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStampPaletteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postStampPaletteRequest** | [**PostStampPaletteRequest**](PostStampPaletteRequest.md) |  | 

### Return type

[**StampPalette**](StampPalette.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStamp

> DeleteStamp(ctx, stampId).Execute()

スタンプを削除



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
	stampId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.StampAPI.DeleteStamp(context.Background(), stampId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.DeleteStamp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stampId** | **string** | スタンプUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStampRequest struct via the builder pattern


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


## DeleteStampPalette

> DeleteStampPalette(ctx, paletteId).Execute()

スタンプパレットを削除



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
	paletteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプパレットUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.StampAPI.DeleteStampPalette(context.Background(), paletteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.DeleteStampPalette``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paletteId** | **string** | スタンプパレットUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStampPaletteRequest struct via the builder pattern


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


## EditStamp

> EditStamp(ctx, stampId).PatchStampRequest(patchStampRequest).Execute()

スタンプ情報を変更



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
	stampId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプUUID
	patchStampRequest := *traq.NewPatchStampRequest() // PatchStampRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.StampAPI.EditStamp(context.Background(), stampId).PatchStampRequest(patchStampRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.EditStamp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stampId** | **string** | スタンプUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditStampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchStampRequest** | [**PatchStampRequest**](PatchStampRequest.md) |  | 

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


## EditStampPalette

> EditStampPalette(ctx, paletteId).PatchStampPaletteRequest(patchStampPaletteRequest).Execute()

スタンプパレットを編集



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
	paletteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプパレットUUID
	patchStampPaletteRequest := *traq.NewPatchStampPaletteRequest() // PatchStampPaletteRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.StampAPI.EditStampPalette(context.Background(), paletteId).PatchStampPaletteRequest(patchStampPaletteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.EditStampPalette``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paletteId** | **string** | スタンプパレットUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditStampPaletteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchStampPaletteRequest** | [**PatchStampPaletteRequest**](PatchStampPaletteRequest.md) |  | 

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
	resp, r, err := apiClient.StampAPI.GetMessageStamps(context.Background(), messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.GetMessageStamps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageStamps`: []MessageStamp
	fmt.Fprintf(os.Stdout, "Response from `StampAPI.GetMessageStamps`: %v\n", resp)
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


## GetMyStampHistory

> []StampHistoryEntry GetMyStampHistory(ctx).Limit(limit).Execute()

スタンプ履歴を取得



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
	limit := int32(56) // int32 | 件数 (optional) (default to 100)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.StampAPI.GetMyStampHistory(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.GetMyStampHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyStampHistory`: []StampHistoryEntry
	fmt.Fprintf(os.Stdout, "Response from `StampAPI.GetMyStampHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyStampHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | 件数 | [default to 100]

### Return type

[**[]StampHistoryEntry**](StampHistoryEntry.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStamp

> Stamp GetStamp(ctx, stampId).Execute()

スタンプ情報を取得



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
	stampId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.StampAPI.GetStamp(context.Background(), stampId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.GetStamp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStamp`: Stamp
	fmt.Fprintf(os.Stdout, "Response from `StampAPI.GetStamp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stampId** | **string** | スタンプUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Stamp**](Stamp.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStampImage

> *os.File GetStampImage(ctx, stampId).Execute()

スタンプ画像を取得



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
	stampId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.StampAPI.GetStampImage(context.Background(), stampId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.GetStampImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStampImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `StampAPI.GetStampImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stampId** | **string** | スタンプUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStampImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, image/gif, image/jpeg, image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStampPalette

> StampPalette GetStampPalette(ctx, paletteId).Execute()

スタンプパレットを取得



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
	paletteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプパレットUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.StampAPI.GetStampPalette(context.Background(), paletteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.GetStampPalette``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStampPalette`: StampPalette
	fmt.Fprintf(os.Stdout, "Response from `StampAPI.GetStampPalette`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paletteId** | **string** | スタンプパレットUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStampPaletteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StampPalette**](StampPalette.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStampPalettes

> []StampPalette GetStampPalettes(ctx).Execute()

スタンプパレットのリストを取得



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
	resp, r, err := apiClient.StampAPI.GetStampPalettes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.GetStampPalettes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStampPalettes`: []StampPalette
	fmt.Fprintf(os.Stdout, "Response from `StampAPI.GetStampPalettes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStampPalettesRequest struct via the builder pattern


### Return type

[**[]StampPalette**](StampPalette.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStampStats

> StampStats GetStampStats(ctx, stampId).Execute()

スタンプ統計情報を取得



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
	stampId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | スタンプUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.StampAPI.GetStampStats(context.Background(), stampId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.GetStampStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStampStats`: StampStats
	fmt.Fprintf(os.Stdout, "Response from `StampAPI.GetStampStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stampId** | **string** | スタンプUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStampStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StampStats**](StampStats.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStamps

> []StampWithThumbnail GetStamps(ctx).IncludeUnicode(includeUnicode).Type_(type_).Execute()

スタンプリストを取得



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
	includeUnicode := true // bool | Unicode絵文字を含ませるかどうか Deprecated: typeクエリを指定しなければ全てのスタンプを取得できるため、そちらを利用してください  (optional) (default to true)
	type_ := "type__example" // string | 取得するスタンプの種類 (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.StampAPI.GetStamps(context.Background()).IncludeUnicode(includeUnicode).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.GetStamps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStamps`: []StampWithThumbnail
	fmt.Fprintf(os.Stdout, "Response from `StampAPI.GetStamps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStampsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUnicode** | **bool** | Unicode絵文字を含ませるかどうか Deprecated: typeクエリを指定しなければ全てのスタンプを取得できるため、そちらを利用してください  | [default to true]
 **type_** | **string** | 取得するスタンプの種類 | 

### Return type

[**[]StampWithThumbnail**](StampWithThumbnail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
	r, err := apiClient.StampAPI.RemoveMessageStamp(context.Background(), messageId, stampId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampAPI.RemoveMessageStamp``: %v\n", err)
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

