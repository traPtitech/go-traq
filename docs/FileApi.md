# \FileApi

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFile**](FileApi.md#DeleteFile) | **Delete** /files/{fileId} | ファイルを削除
[**GetFile**](FileApi.md#GetFile) | **Get** /files/{fileId} | ファイルをダウンロード
[**GetFileMeta**](FileApi.md#GetFileMeta) | **Get** /files/{fileId}/meta | ファイルメタを取得
[**GetFiles**](FileApi.md#GetFiles) | **Get** /files | ファイルメタのリストを取得
[**GetThumbnailImage**](FileApi.md#GetThumbnailImage) | **Get** /files/{fileId}/thumbnail | サムネイル画像を取得
[**PostFile**](FileApi.md#PostFile) | **Post** /files | ファイルをアップロード



## DeleteFile

> DeleteFile(ctx, fileId).Execute()

ファイルを削除



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
    fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ファイルUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FileApi.DeleteFile(context.Background(), fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.DeleteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ファイルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


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


## GetFile

> *os.File GetFile(ctx, fileId).Dl(dl).Execute()

ファイルをダウンロード



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
    fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ファイルUUID
    dl := int32(56) // int32 | 1を指定するとレスポンスにContent-Dispositionヘッダーが付与されます (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileApi.GetFile(context.Background(), fileId).Dl(dl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.GetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FileApi.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ファイルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dl** | **int32** | 1を指定するとレスポンスにContent-Dispositionヘッダーが付与されます | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileMeta

> FileInfo GetFileMeta(ctx, fileId).Execute()

ファイルメタを取得



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
    fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ファイルUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileApi.GetFileMeta(context.Background(), fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.GetFileMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileMeta`: FileInfo
    fmt.Fprintf(os.Stdout, "Response from `FileApi.GetFileMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ファイルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileInfo**](FileInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiles

> []FileInfo GetFiles(ctx).ChannelId(channelId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Mine(mine).Execute()

ファイルメタのリストを取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | アップロード先チャンネルUUID (optional)
    limit := int32(50) // int32 | 取得する件数 (optional)
    offset := int32(150) // int32 | 取得するオフセット (optional) (default to 0)
    since := time.Now() // time.Time | 取得する時間範囲の開始日時 (optional) (default to "0000-01-01T00:00Z")
    until := time.Now() // time.Time | 取得する時間範囲の終了日時 (optional)
    inclusive := true // bool | 範囲の端を含めるかどうか (optional) (default to false)
    order := "order_example" // string | 昇順か降順か (optional) (default to "desc")
    mine := true // bool | アップロード者が自分のファイルのみを取得するか (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileApi.GetFiles(context.Background()).ChannelId(channelId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Mine(mine).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.GetFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFiles`: []FileInfo
    fmt.Fprintf(os.Stdout, "Response from `FileApi.GetFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | アップロード先チャンネルUUID | 
 **limit** | **int32** | 取得する件数 | 
 **offset** | **int32** | 取得するオフセット | [default to 0]
 **since** | **time.Time** | 取得する時間範囲の開始日時 | [default to &quot;0000-01-01T00:00Z&quot;]
 **until** | **time.Time** | 取得する時間範囲の終了日時 | 
 **inclusive** | **bool** | 範囲の端を含めるかどうか | [default to false]
 **order** | **string** | 昇順か降順か | [default to &quot;desc&quot;]
 **mine** | **bool** | アップロード者が自分のファイルのみを取得するか | [default to false]

### Return type

[**[]FileInfo**](FileInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThumbnailImage

> *os.File GetThumbnailImage(ctx, fileId).Type_(type_).Execute()

サムネイル画像を取得



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
    fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ファイルUUID
    type_ := openapiclient.ThumbnailType("image") // ThumbnailType | 取得するサムネイルのタイプ (optional) (default to "image")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileApi.GetThumbnailImage(context.Background(), fileId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.GetThumbnailImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThumbnailImage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FileApi.GetThumbnailImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ファイルUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**ThumbnailType**](ThumbnailType.md) | 取得するサムネイルのタイプ | [default to &quot;image&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFile

> FileInfo PostFile(ctx).File(file).ChannelId(channelId).Execute()

ファイルをアップロード



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
    file := os.NewFile(1234, "some_file") // *os.File | ファイル本体
    channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | アップロード先チャンネルUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileApi.PostFile(context.Background()).File(file).ChannelId(channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.PostFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFile`: FileInfo
    fmt.Fprintf(os.Stdout, "Response from `FileApi.PostFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | ファイル本体 | 
 **channelId** | **string** | アップロード先チャンネルUUID | 

### Return type

[**FileInfo**](FileInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

