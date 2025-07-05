# \UserAPI

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserTag**](UserAPI.md#AddUserTag) | **Post** /users/{userId}/tags | ユーザーにタグを追加
[**ChangeUserIcon**](UserAPI.md#ChangeUserIcon) | **Put** /users/{userId}/icon | ユーザーのアイコン画像を変更します
[**ChangeUserPassword**](UserAPI.md#ChangeUserPassword) | **Put** /users/{userId}/password | ユーザーのパスワードを変更
[**CreateUser**](UserAPI.md#CreateUser) | **Post** /users | ユーザーを登録
[**EditUser**](UserAPI.md#EditUser) | **Patch** /users/{userId} | ユーザー情報を変更
[**EditUserTag**](UserAPI.md#EditUserTag) | **Patch** /users/{userId}/tags/{tagId} | ユーザーのタグを編集
[**GetDirectMessages**](UserAPI.md#GetDirectMessages) | **Get** /users/{userId}/messages | ダイレクトメッセージのリストを取得
[**GetUser**](UserAPI.md#GetUser) | **Get** /users/{userId} | ユーザー詳細情報を取得
[**GetUserDMChannel**](UserAPI.md#GetUserDMChannel) | **Get** /users/{userId}/dm-channel | DMチャンネル情報を取得
[**GetUserIcon**](UserAPI.md#GetUserIcon) | **Get** /users/{userId}/icon | ユーザーのアイコン画像を取得
[**GetUserStats**](UserAPI.md#GetUserStats) | **Get** /users/{userId}/stats | ユーザー統計情報を取得
[**GetUserTags**](UserAPI.md#GetUserTags) | **Get** /users/{userId}/tags | ユーザーのタグリストを取得
[**GetUsers**](UserAPI.md#GetUsers) | **Get** /users | ユーザーのリストを取得
[**PostDirectMessage**](UserAPI.md#PostDirectMessage) | **Post** /users/{userId}/messages | ダイレクトメッセージを送信
[**RemoveUserTag**](UserAPI.md#RemoveUserTag) | **Delete** /users/{userId}/tags/{tagId} | ユーザーからタグを削除します



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
	resp, r, err := apiClient.UserAPI.AddUserTag(context.Background(), userId).PostUserTagRequest(postUserTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AddUserTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserTag`: UserTag
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AddUserTag`: %v\n", resp)
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


## ChangeUserIcon

> ChangeUserIcon(ctx, userId).File(file).Execute()

ユーザーのアイコン画像を変更します



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
	file := os.NewFile(1234, "some_file") // *os.File | アイコン画像(2MBまでのpng, jpeg, gif)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.ChangeUserIcon(context.Background(), userId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ChangeUserIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserIconRequest struct via the builder pattern


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


## ChangeUserPassword

> ChangeUserPassword(ctx, userId).PutUserPasswordRequest(putUserPasswordRequest).Execute()

ユーザーのパスワードを変更



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
	putUserPasswordRequest := *traq.NewPutUserPasswordRequest("NewPassword_example") // PutUserPasswordRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.ChangeUserPassword(context.Background(), userId).PutUserPasswordRequest(putUserPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ChangeUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putUserPasswordRequest** | [**PutUserPasswordRequest**](PutUserPasswordRequest.md) |  | 

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


## CreateUser

> UserDetail CreateUser(ctx).PostUserRequest(postUserRequest).Execute()

ユーザーを登録



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
	postUserRequest := *traq.NewPostUserRequest("Name_example") // PostUserRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateUser(context.Background()).PostUserRequest(postUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: UserDetail
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postUserRequest** | [**PostUserRequest**](PostUserRequest.md) |  | 

### Return type

[**UserDetail**](UserDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditUser

> EditUser(ctx, userId).PatchUserRequest(patchUserRequest).Execute()

ユーザー情報を変更



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
	patchUserRequest := *traq.NewPatchUserRequest() // PatchUserRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.EditUser(context.Background(), userId).PatchUserRequest(patchUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.EditUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchUserRequest** | [**PatchUserRequest**](PatchUserRequest.md) |  | 

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
	r, err := apiClient.UserAPI.EditUserTag(context.Background(), userId, tagId).PatchUserTagRequest(patchUserTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.EditUserTag``: %v\n", err)
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
	resp, r, err := apiClient.UserAPI.GetDirectMessages(context.Background(), userId).Limit(limit).Offset(offset).Since(since).Until(until).Inclusive(inclusive).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetDirectMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectMessages`: []Message
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetDirectMessages`: %v\n", resp)
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


## GetUser

> UserDetail GetUser(ctx, userId).Execute()

ユーザー詳細情報を取得



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
	resp, r, err := apiClient.UserAPI.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: UserDetail
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserDetail**](UserDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDMChannel

> DMChannel GetUserDMChannel(ctx, userId).Execute()

DMチャンネル情報を取得



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserDMChannel(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserDMChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserDMChannel`: DMChannel
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserDMChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDMChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DMChannel**](DMChannel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserIcon

> *os.File GetUserIcon(ctx, userId).Execute()

ユーザーのアイコン画像を取得



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
	resp, r, err := apiClient.UserAPI.GetUserIcon(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserIcon`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserIconRequest struct via the builder pattern


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


## GetUserStats

> UserStats GetUserStats(ctx, userId).Execute()

ユーザー統計情報を取得



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
	resp, r, err := apiClient.UserAPI.GetUserStats(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserStats`: UserStats
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserStats**](UserStats.md)

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
	resp, r, err := apiClient.UserAPI.GetUserTags(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTags`: []UserTag
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserTags`: %v\n", resp)
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


## GetUsers

> []User GetUsers(ctx).IncludeSuspended(includeSuspended).Name(name).Execute()

ユーザーのリストを取得



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
	includeSuspended := true // bool | アカウントがアクティブでないユーザーを含め、全てのユーザーを取得するかどうか (optional) (default to false)
	name := "name_example" // string | 名前が一致するアカウントのみを取得する (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUsers(context.Background()).IncludeSuspended(includeSuspended).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: []User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeSuspended** | **bool** | アカウントがアクティブでないユーザーを含め、全てのユーザーを取得するかどうか | [default to false]
 **name** | **string** | 名前が一致するアカウントのみを取得する | 

### Return type

[**[]User**](User.md)

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
	resp, r, err := apiClient.UserAPI.PostDirectMessage(context.Background(), userId).PostMessageRequest(postMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PostDirectMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDirectMessage`: Message
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PostDirectMessage`: %v\n", resp)
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
	r, err := apiClient.UserAPI.RemoveUserTag(context.Background(), userId, tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RemoveUserTag``: %v\n", err)
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

