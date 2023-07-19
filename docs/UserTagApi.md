# \UserTagApi

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMyUserTag**](UserTagApi.md#AddMyUserTag) | **Post** /users/me/tags | 自分にタグを追加
[**AddUserTag**](UserTagApi.md#AddUserTag) | **Post** /users/{userId}/tags | ユーザーにタグを追加
[**EditMyUserTag**](UserTagApi.md#EditMyUserTag) | **Patch** /users/me/tags/{tagId} | 自分のタグを編集
[**EditUserTag**](UserTagApi.md#EditUserTag) | **Patch** /users/{userId}/tags/{tagId} | ユーザーのタグを編集
[**GetMyUserTags**](UserTagApi.md#GetMyUserTags) | **Get** /users/me/tags | 自分のタグリストを取得
[**GetTag**](UserTagApi.md#GetTag) | **Get** /tags/{tagId} | タグ情報を取得
[**GetUserTags**](UserTagApi.md#GetUserTags) | **Get** /users/{userId}/tags | ユーザーのタグリストを取得
[**RemoveMyUserTag**](UserTagApi.md#RemoveMyUserTag) | **Delete** /users/me/tags/{tagId} | 自分からタグを削除します
[**RemoveUserTag**](UserTagApi.md#RemoveUserTag) | **Delete** /users/{userId}/tags/{tagId} | ユーザーからタグを削除します



## AddMyUserTag

> UserTag AddMyUserTag(ctx).PostUserTagRequest(postUserTagRequest).Execute()

自分にタグを追加



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
    postUserTagRequest := *openapiclient.NewPostUserTagRequest("Tag_example") // PostUserTagRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTagApi.AddMyUserTag(context.Background()).PostUserTagRequest(postUserTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTagApi.AddMyUserTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMyUserTag`: UserTag
    fmt.Fprintf(os.Stdout, "Response from `UserTagApi.AddMyUserTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMyUserTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postUserTagRequest** | [**PostUserTagRequest**](PostUserTagRequest.md) |  | 

### Return type

[**UserTag**](UserTag.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserTag

> UserTag AddUserTag(ctx, userId).PostUserTagRequest(postUserTagRequest).Execute()

ユーザーにタグを追加



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
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID
    postUserTagRequest := *openapiclient.NewPostUserTagRequest("Tag_example") // PostUserTagRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTagApi.AddUserTag(context.Background(), userId).PostUserTagRequest(postUserTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTagApi.AddUserTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserTag`: UserTag
    fmt.Fprintf(os.Stdout, "Response from `UserTagApi.AddUserTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postUserTagRequest** | [**PostUserTagRequest**](PostUserTagRequest.md) |  | 

### Return type

[**UserTag**](UserTag.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditMyUserTag

> EditMyUserTag(ctx, tagId).PatchUserTagRequest(patchUserTagRequest).Execute()

自分のタグを編集



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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | タグUUID
    patchUserTagRequest := *openapiclient.NewPatchUserTagRequest(false) // PatchUserTagRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserTagApi.EditMyUserTag(context.Background(), tagId).PatchUserTagRequest(patchUserTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTagApi.EditMyUserTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | タグUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMyUserTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchUserTagRequest** | [**PatchUserTagRequest**](PatchUserTagRequest.md) |  | 

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


## EditUserTag

> EditUserTag(ctx, userId, tagId).PatchUserTagRequest(patchUserTagRequest).Execute()

ユーザーのタグを編集



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
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | タグUUID
    patchUserTagRequest := *openapiclient.NewPatchUserTagRequest(false) // PatchUserTagRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserTagApi.EditUserTag(context.Background(), userId, tagId).PatchUserTagRequest(patchUserTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTagApi.EditUserTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 
**tagId** | **string** | タグUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditUserTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchUserTagRequest** | [**PatchUserTagRequest**](PatchUserTagRequest.md) |  | 

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


## GetMyUserTags

> []UserTag GetMyUserTags(ctx).Execute()

自分のタグリストを取得



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
    resp, r, err := apiClient.UserTagApi.GetMyUserTags(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTagApi.GetMyUserTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMyUserTags`: []UserTag
    fmt.Fprintf(os.Stdout, "Response from `UserTagApi.GetMyUserTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyUserTagsRequest struct via the builder pattern


### Return type

[**[]UserTag**](UserTag.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTag

> Tag GetTag(ctx, tagId).Execute()

タグ情報を取得



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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | タグUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTagApi.GetTag(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTagApi.GetTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `UserTagApi.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | タグUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tag**](Tag.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTags

> []UserTag GetUserTags(ctx, userId).Execute()

ユーザーのタグリストを取得



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
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTagApi.GetUserTags(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTagApi.GetUserTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserTags`: []UserTag
    fmt.Fprintf(os.Stdout, "Response from `UserTagApi.GetUserTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserTag**](UserTag.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMyUserTag

> RemoveMyUserTag(ctx, tagId).Execute()

自分からタグを削除します



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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | タグUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserTagApi.RemoveMyUserTag(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTagApi.RemoveMyUserTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | タグUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMyUserTagRequest struct via the builder pattern


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


## RemoveUserTag

> RemoveUserTag(ctx, userId, tagId).Execute()

ユーザーからタグを削除します



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
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | タグUUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserTagApi.RemoveUserTag(context.Background(), userId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTagApi.RemoveUserTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 
**tagId** | **string** | タグUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserTagRequest struct via the builder pattern


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

