# \GroupAPI

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserGroupAdmin**](GroupAPI.md#AddUserGroupAdmin) | **Post** /groups/{groupId}/admins | グループ管理者を追加
[**AddUserGroupMember**](GroupAPI.md#AddUserGroupMember) | **Post** /groups/{groupId}/members | グループメンバーを追加
[**ChangeUserGroupIcon**](GroupAPI.md#ChangeUserGroupIcon) | **Put** /groups/{groupId}/icon | ユーザーグループのアイコンを変更
[**CreateUserGroup**](GroupAPI.md#CreateUserGroup) | **Post** /groups | ユーザーグループを作成
[**DeleteUserGroup**](GroupAPI.md#DeleteUserGroup) | **Delete** /groups/{groupId} | ユーザーグループを削除
[**EditUserGroup**](GroupAPI.md#EditUserGroup) | **Patch** /groups/{groupId} | ユーザーグループを編集
[**EditUserGroupMember**](GroupAPI.md#EditUserGroupMember) | **Patch** /groups/{groupId}/members/{userId} | グループメンバーを編集
[**GetUserGroup**](GroupAPI.md#GetUserGroup) | **Get** /groups/{groupId} | ユーザーグループを取得
[**GetUserGroupAdmins**](GroupAPI.md#GetUserGroupAdmins) | **Get** /groups/{groupId}/admins | グループ管理者を取得
[**GetUserGroupMembers**](GroupAPI.md#GetUserGroupMembers) | **Get** /groups/{groupId}/members | グループメンバーを取得
[**GetUserGroups**](GroupAPI.md#GetUserGroups) | **Get** /groups | ユーザーグループのリストを取得
[**RemoveUserGroupAdmin**](GroupAPI.md#RemoveUserGroupAdmin) | **Delete** /groups/{groupId}/admins/{userId} | グループ管理者を削除
[**RemoveUserGroupMember**](GroupAPI.md#RemoveUserGroupMember) | **Delete** /groups/{groupId}/members/{userId} | グループメンバーを削除
[**RemoveUserGroupMembers**](GroupAPI.md#RemoveUserGroupMembers) | **Delete** /groups/{groupId}/members | グループメンバーを一括削除



## AddUserGroupAdmin

> AddUserGroupAdmin(ctx, groupId).PostUserGroupAdminRequest(postUserGroupAdminRequest).Execute()

グループ管理者を追加



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID
	postUserGroupAdminRequest := *traq.NewPostUserGroupAdminRequest("Id_example") // PostUserGroupAdminRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.AddUserGroupAdmin(context.Background(), groupId).PostUserGroupAdminRequest(postUserGroupAdminRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.AddUserGroupAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserGroupAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postUserGroupAdminRequest** | [**PostUserGroupAdminRequest**](PostUserGroupAdminRequest.md) |  | 

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


## AddUserGroupMember

> AddUserGroupMember(ctx, groupId).AddUserGroupMemberRequest(addUserGroupMemberRequest).Execute()

グループメンバーを追加



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID
	addUserGroupMemberRequest := traq.addUserGroupMember_request{UserGroupMember: traq.NewUserGroupMember("Id_example", "Role_example")} // AddUserGroupMemberRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.AddUserGroupMember(context.Background(), groupId).AddUserGroupMemberRequest(addUserGroupMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.AddUserGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addUserGroupMemberRequest** | [**AddUserGroupMemberRequest**](AddUserGroupMemberRequest.md) |  | 

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


## ChangeUserGroupIcon

> ChangeUserGroupIcon(ctx, groupId).File(file).Execute()

ユーザーグループのアイコンを変更



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID
	file := os.NewFile(1234, "some_file") // *os.File | アイコン画像(2MBまでのpng, jpeg, gif)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.ChangeUserGroupIcon(context.Background(), groupId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ChangeUserGroupIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserGroupIconRequest struct via the builder pattern


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


## CreateUserGroup

> UserGroup CreateUserGroup(ctx).PostUserGroupRequest(postUserGroupRequest).Execute()

ユーザーグループを作成



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
	postUserGroupRequest := *traq.NewPostUserGroupRequest("Name_example", "Description_example", "Type_example") // PostUserGroupRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.CreateUserGroup(context.Background()).PostUserGroupRequest(postUserGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.CreateUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserGroup`: UserGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.CreateUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postUserGroupRequest** | [**PostUserGroupRequest**](PostUserGroupRequest.md) |  | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserGroup

> DeleteUserGroup(ctx, groupId).Execute()

ユーザーグループを削除



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.DeleteUserGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeleteUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupRequest struct via the builder pattern


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


## EditUserGroup

> EditUserGroup(ctx, groupId).PatchUserGroupRequest(patchUserGroupRequest).Execute()

ユーザーグループを編集



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID
	patchUserGroupRequest := *traq.NewPatchUserGroupRequest() // PatchUserGroupRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.EditUserGroup(context.Background(), groupId).PatchUserGroupRequest(patchUserGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.EditUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchUserGroupRequest** | [**PatchUserGroupRequest**](PatchUserGroupRequest.md) |  | 

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


## EditUserGroupMember

> EditUserGroupMember(ctx, groupId, userId).PatchGroupMemberRequest(patchGroupMemberRequest).Execute()

グループメンバーを編集



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID
	patchGroupMemberRequest := *traq.NewPatchGroupMemberRequest("Role_example") // PatchGroupMemberRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.EditUserGroupMember(context.Background(), groupId, userId).PatchGroupMemberRequest(patchGroupMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.EditUserGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditUserGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchGroupMemberRequest** | [**PatchGroupMemberRequest**](PatchGroupMemberRequest.md) |  | 

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


## GetUserGroup

> UserGroup GetUserGroup(ctx, groupId).Execute()

ユーザーグループを取得



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetUserGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroup`: UserGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroupAdmins

> []string GetUserGroupAdmins(ctx, groupId).Execute()

グループ管理者を取得



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetUserGroupAdmins(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetUserGroupAdmins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroupAdmins`: []string
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetUserGroupAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroupMembers

> []UserGroupMember GetUserGroupMembers(ctx, groupId).Execute()

グループメンバーを取得



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetUserGroupMembers(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetUserGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroupMembers`: []UserGroupMember
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetUserGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserGroupMember**](UserGroupMember.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroups

> []UserGroup GetUserGroups(ctx).Execute()

ユーザーグループのリストを取得



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
	resp, r, err := apiClient.GroupAPI.GetUserGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroups`: []UserGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetUserGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupsRequest struct via the builder pattern


### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserGroupAdmin

> RemoveUserGroupAdmin(ctx, groupId, userId).Execute()

グループ管理者を削除



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.RemoveUserGroupAdmin(context.Background(), groupId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.RemoveUserGroupAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserGroupAdminRequest struct via the builder pattern


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


## RemoveUserGroupMember

> RemoveUserGroupMember(ctx, groupId, userId).Execute()

グループメンバーを削除



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.RemoveUserGroupMember(context.Background(), groupId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.RemoveUserGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 
**userId** | **string** | ユーザーUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserGroupMemberRequest struct via the builder pattern


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


## RemoveUserGroupMembers

> RemoveUserGroupMembers(ctx, groupId).Execute()

グループメンバーを一括削除



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ユーザーグループUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.RemoveUserGroupMembers(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.RemoveUserGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ユーザーグループUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserGroupMembersRequest struct via the builder pattern


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

