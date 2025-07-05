# \AuthenticationAPI

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyExternalAccounts**](AuthenticationAPI.md#GetMyExternalAccounts) | **Get** /users/me/ex-accounts | 外部ログインアカウント一覧を取得
[**GetMySessions**](AuthenticationAPI.md#GetMySessions) | **Get** /users/me/sessions | 自分のログインセッションリストを取得
[**LinkExternalAccount**](AuthenticationAPI.md#LinkExternalAccount) | **Post** /users/me/ex-accounts/link | 外部ログインアカウントを紐付ける
[**Login**](AuthenticationAPI.md#Login) | **Post** /login | ログイン
[**Logout**](AuthenticationAPI.md#Logout) | **Post** /logout | ログアウト
[**RevokeMySession**](AuthenticationAPI.md#RevokeMySession) | **Delete** /users/me/sessions/{sessionId} | セッションを無効化
[**UnlinkExternalAccount**](AuthenticationAPI.md#UnlinkExternalAccount) | **Post** /users/me/ex-accounts/unlink | 外部ログインアカウントの紐付けを解除



## GetMyExternalAccounts

> []ExternalProviderUser GetMyExternalAccounts(ctx).Execute()

外部ログインアカウント一覧を取得



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
	resp, r, err := apiClient.AuthenticationAPI.GetMyExternalAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetMyExternalAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyExternalAccounts`: []ExternalProviderUser
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetMyExternalAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyExternalAccountsRequest struct via the builder pattern


### Return type

[**[]ExternalProviderUser**](ExternalProviderUser.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMySessions

> []LoginSession GetMySessions(ctx).Execute()

自分のログインセッションリストを取得



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
	resp, r, err := apiClient.AuthenticationAPI.GetMySessions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetMySessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMySessions`: []LoginSession
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetMySessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMySessionsRequest struct via the builder pattern


### Return type

[**[]LoginSession**](LoginSession.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkExternalAccount

> LinkExternalAccount(ctx).PostLinkExternalAccount(postLinkExternalAccount).Execute()

外部ログインアカウントを紐付ける



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
	postLinkExternalAccount := *traq.NewPostLinkExternalAccount("ProviderName_example") // PostLinkExternalAccount |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.LinkExternalAccount(context.Background()).PostLinkExternalAccount(postLinkExternalAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.LinkExternalAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkExternalAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postLinkExternalAccount** | [**PostLinkExternalAccount**](PostLinkExternalAccount.md) |  | 

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


## Login

> Login(ctx).Redirect(redirect).PostLoginRequest(postLoginRequest).Execute()

ログイン



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
	redirect := "redirect_example" // string | リダイレクト先 (optional)
	postLoginRequest := *traq.NewPostLoginRequest("Name_example", "Password_example") // PostLoginRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.Login(context.Background()).Redirect(redirect).PostLoginRequest(postLoginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirect** | **string** | リダイレクト先 | 
 **postLoginRequest** | [**PostLoginRequest**](PostLoginRequest.md) |  | 

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


## Logout

> Logout(ctx).Redirect(redirect).All(all).Execute()

ログアウト



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
	redirect := "redirect_example" // string | リダイレクト先 (optional)
	all := true // bool | 全てのセッションでログアウトするかどうか (optional) (default to false)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.Logout(context.Background()).Redirect(redirect).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirect** | **string** | リダイレクト先 | 
 **all** | **bool** | 全てのセッションでログアウトするかどうか | [default to false]

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


## RevokeMySession

> RevokeMySession(ctx, sessionId).Execute()

セッションを無効化



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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | セッションUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.RevokeMySession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.RevokeMySession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | セッションUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeMySessionRequest struct via the builder pattern


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


## UnlinkExternalAccount

> UnlinkExternalAccount(ctx).PostUnlinkExternalAccount(postUnlinkExternalAccount).Execute()

外部ログインアカウントの紐付けを解除



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
	postUnlinkExternalAccount := *traq.NewPostUnlinkExternalAccount("ProviderName_example") // PostUnlinkExternalAccount |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.UnlinkExternalAccount(context.Background()).PostUnlinkExternalAccount(postUnlinkExternalAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UnlinkExternalAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkExternalAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postUnlinkExternalAccount** | [**PostUnlinkExternalAccount**](PostUnlinkExternalAccount.md) |  | 

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

