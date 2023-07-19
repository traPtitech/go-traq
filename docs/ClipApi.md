# \ClipApi

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClipMessage**](ClipApi.md#ClipMessage) | **Post** /clip-folders/{folderId}/messages | メッセージをクリップフォルダに追加
[**CreateClipFolder**](ClipApi.md#CreateClipFolder) | **Post** /clip-folders | クリップフォルダを作成
[**DeleteClipFolder**](ClipApi.md#DeleteClipFolder) | **Delete** /clip-folders/{folderId} | クリップフォルダを削除
[**EditClipFolder**](ClipApi.md#EditClipFolder) | **Patch** /clip-folders/{folderId} | クリップフォルダ情報を編集
[**GetClipFolder**](ClipApi.md#GetClipFolder) | **Get** /clip-folders/{folderId} | クリップフォルダ情報を取得
[**GetClipFolders**](ClipApi.md#GetClipFolders) | **Get** /clip-folders | クリップフォルダのリストを取得
[**GetClips**](ClipApi.md#GetClips) | **Get** /clip-folders/{folderId}/messages | フォルダ内のクリップのリストを取得
[**GetMessageClips**](ClipApi.md#GetMessageClips) | **Get** /messages/{messageId}/clips | 自分のクリップを取得
[**UnclipMessage**](ClipApi.md#UnclipMessage) | **Delete** /clip-folders/{folderId}/messages/{messageId} | メッセージをクリップフォルダから除外



## ClipMessage

> ClippedMessage ClipMessage(ctx, folderId).PostClipFolderMessageRequest(postClipFolderMessageRequest).Execute()

メッセージをクリップフォルダに追加



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
    folderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | クリップフォルダUUID
    postClipFolderMessageRequest := *traq.NewPostClipFolderMessageRequest("MessageId_example") // PostClipFolderMessageRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ClipApi.ClipMessage(context.Background(), folderId).PostClipFolderMessageRequest(postClipFolderMessageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.ClipMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClipMessage`: ClippedMessage
    fmt.Fprintf(os.Stdout, "Response from `ClipApi.ClipMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | クリップフォルダUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClipMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postClipFolderMessageRequest** | [**PostClipFolderMessageRequest**](PostClipFolderMessageRequest.md) |  | 

### Return type

[**ClippedMessage**](ClippedMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClipFolder

> ClipFolder CreateClipFolder(ctx).PostClipFolderRequest(postClipFolderRequest).Execute()

クリップフォルダを作成



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
    postClipFolderRequest := *traq.NewPostClipFolderRequest("Name_example", "Description_example") // PostClipFolderRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ClipApi.CreateClipFolder(context.Background()).PostClipFolderRequest(postClipFolderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.CreateClipFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClipFolder`: ClipFolder
    fmt.Fprintf(os.Stdout, "Response from `ClipApi.CreateClipFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClipFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postClipFolderRequest** | [**PostClipFolderRequest**](PostClipFolderRequest.md) |  | 

### Return type

[**ClipFolder**](ClipFolder.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClipFolder

> DeleteClipFolder(ctx, folderId).Execute()

クリップフォルダを削除



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
    folderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | クリップフォルダUUID

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.ClipApi.DeleteClipFolder(context.Background(), folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.DeleteClipFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | クリップフォルダUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClipFolderRequest struct via the builder pattern


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


## EditClipFolder

> EditClipFolder(ctx, folderId).PatchClipFolderRequest(patchClipFolderRequest).Execute()

クリップフォルダ情報を編集



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
    folderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | クリップフォルダUUID
    patchClipFolderRequest := *traq.NewPatchClipFolderRequest() // PatchClipFolderRequest |  (optional)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.ClipApi.EditClipFolder(context.Background(), folderId).PatchClipFolderRequest(patchClipFolderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.EditClipFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | クリップフォルダUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditClipFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchClipFolderRequest** | [**PatchClipFolderRequest**](PatchClipFolderRequest.md) |  | 

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


## GetClipFolder

> ClipFolder GetClipFolder(ctx, folderId).Execute()

クリップフォルダ情報を取得



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
    folderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | クリップフォルダUUID

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ClipApi.GetClipFolder(context.Background(), folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.GetClipFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClipFolder`: ClipFolder
    fmt.Fprintf(os.Stdout, "Response from `ClipApi.GetClipFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | クリップフォルダUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClipFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClipFolder**](ClipFolder.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClipFolders

> []ClipFolder GetClipFolders(ctx).Execute()

クリップフォルダのリストを取得



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
    resp, r, err := apiClient.ClipApi.GetClipFolders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.GetClipFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClipFolders`: []ClipFolder
    fmt.Fprintf(os.Stdout, "Response from `ClipApi.GetClipFolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClipFoldersRequest struct via the builder pattern


### Return type

[**[]ClipFolder**](ClipFolder.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClips

> []ClippedMessage GetClips(ctx, folderId).Limit(limit).Offset(offset).Order(order).Execute()

フォルダ内のクリップのリストを取得



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
    folderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | クリップフォルダUUID
    limit := int32(50) // int32 | 取得する件数 (optional)
    offset := int32(150) // int32 | 取得するオフセット (optional) (default to 0)
    order := "order_example" // string | 昇順か降順か (optional) (default to "desc")

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ClipApi.GetClips(context.Background(), folderId).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.GetClips``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClips`: []ClippedMessage
    fmt.Fprintf(os.Stdout, "Response from `ClipApi.GetClips`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | クリップフォルダUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | 取得する件数 | 
 **offset** | **int32** | 取得するオフセット | [default to 0]
 **order** | **string** | 昇順か降順か | [default to &quot;desc&quot;]

### Return type

[**[]ClippedMessage**](ClippedMessage.md)

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
    resp, r, err := apiClient.ClipApi.GetMessageClips(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.GetMessageClips``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageClips`: []MessageClip
    fmt.Fprintf(os.Stdout, "Response from `ClipApi.GetMessageClips`: %v\n", resp)
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


## UnclipMessage

> UnclipMessage(ctx, folderId, messageId).Execute()

メッセージをクリップフォルダから除外



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
    folderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | クリップフォルダUUID
    messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | メッセージUUID

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    r, err := apiClient.ClipApi.UnclipMessage(context.Background(), folderId, messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.UnclipMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | クリップフォルダUUID | 
**messageId** | **string** | メッセージUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnclipMessageRequest struct via the builder pattern


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

