# \UserTagAPI

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMyUserTag**](UserTagAPI.md#AddMyUserTag) | **Post** /users/me/tags | 自分にタグを追加
[**AddUserTag**](UserTagAPI.md#AddUserTag) | **Post** /users/{userId}/tags | ユーザーにタグを追加
[**EditMyUserTag**](UserTagAPI.md#EditMyUserTag) | **Patch** /users/me/tags/{tagId} | 自分のタグを編集
[**EditUserTag**](UserTagAPI.md#EditUserTag) | **Patch** /users/{userId}/tags/{tagId} | ユーザーのタグを編集
[**GetMyUserTags**](UserTagAPI.md#GetMyUserTags) | **Get** /users/me/tags | 自分のタグリストを取得
[**GetTag**](UserTagAPI.md#GetTag) | **Get** /tags/{tagId} | タグ情報を取得
[**GetUserTags**](UserTagAPI.md#GetUserTags) | **Get** /users/{userId}/tags | ユーザーのタグリストを取得
[**RemoveMyUserTag**](UserTagAPI.md#RemoveMyUserTag) | **Delete** /users/me/tags/{tagId} | 自分からタグを削除します
[**RemoveUserTag**](UserTagAPI.md#RemoveUserTag) | **Delete** /users/{userId}/tags/{tagId} | ユーザーからタグを削除します



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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	postUserTagRequest := *traq.NewPostUserTagRequest("Tag_example") // PostUserTagRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTagAPI.AddMyUserTag(context.Background()).PostUserTagRequest(postUserTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTagAPI.AddMyUserTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMyUserTag`: UserTag
	fmt.Fprintf(os.Stdout, "Response from `UserTagAPI.AddMyUserTag`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID
	postUserTagRequest := *traq.NewPostUserTagRequest("Tag_example") // PostUserTagRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTagAPI.AddUserTag(context.Background(), userId).PostUserTagRequest(postUserTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTagAPI.AddUserTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserTag`: UserTag
	fmt.Fprintf(os.Stdout, "Response from `UserTagAPI.AddUserTag`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | タグUUID
	patchUserTagRequest := *traq.NewPatchUserTagRequest(false) // PatchUserTagRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.UserTagAPI.EditMyUserTag(context.Background(), tagId).PatchUserTagRequest(patchUserTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTagAPI.EditMyUserTag``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | タグUUID
	patchUserTagRequest := *traq.NewPatchUserTagRequest(false) // PatchUserTagRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.UserTagAPI.EditUserTag(context.Background(), userId, tagId).PatchUserTagRequest(patchUserTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTagAPI.EditUserTag``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTagAPI.GetMyUserTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTagAPI.GetMyUserTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyUserTags`: []UserTag
	fmt.Fprintf(os.Stdout, "Response from `UserTagAPI.GetMyUserTags`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | タグUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTagAPI.GetTag(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTagAPI.GetTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTag`: Tag
	fmt.Fprintf(os.Stdout, "Response from `UserTagAPI.GetTag`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTagAPI.GetUserTags(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTagAPI.GetUserTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTags`: []UserTag
	fmt.Fprintf(os.Stdout, "Response from `UserTagAPI.GetUserTags`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | タグUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.UserTagAPI.RemoveMyUserTag(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTagAPI.RemoveMyUserTag``: %v\n", err)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | タグUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.UserTagAPI.RemoveUserTag(context.Background(), userId, tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTagAPI.RemoveUserTag``: %v\n", err)
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

