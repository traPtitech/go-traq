# \Oauth2API

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClient**](Oauth2API.md#CreateClient) | **Post** /clients | OAuth2クライアントを作成
[**DeleteClient**](Oauth2API.md#DeleteClient) | **Delete** /clients/{clientId} | OAuth2クライアントを削除
[**EditClient**](Oauth2API.md#EditClient) | **Patch** /clients/{clientId} | OAuth2クライアント情報を変更
[**GetClient**](Oauth2API.md#GetClient) | **Get** /clients/{clientId} | OAuth2クライアント情報を取得
[**GetClients**](Oauth2API.md#GetClients) | **Get** /clients | OAuth2クライアントのリストを取得
[**GetMyTokens**](Oauth2API.md#GetMyTokens) | **Get** /users/me/tokens | 有効トークンのリストを取得
[**GetOAuth2Authorize**](Oauth2API.md#GetOAuth2Authorize) | **Get** /oauth2/authorize | OAuth2 認可エンドポイント
[**PostOAuth2Authorize**](Oauth2API.md#PostOAuth2Authorize) | **Post** /oauth2/authorize | OAuth2 認可エンドポイント
[**PostOAuth2AuthorizeDecide**](Oauth2API.md#PostOAuth2AuthorizeDecide) | **Post** /oauth2/authorize/decide | OAuth2 認可承諾API
[**PostOAuth2Token**](Oauth2API.md#PostOAuth2Token) | **Post** /oauth2/token | OAuth2 トークンエンドポイント
[**RevokeClientTokens**](Oauth2API.md#RevokeClientTokens) | **Delete** /clients/{clientId}/tokens | OAuthクライアントのトークンを削除
[**RevokeMyToken**](Oauth2API.md#RevokeMyToken) | **Delete** /users/me/tokens/{tokenId} | トークンの認可を取り消す
[**RevokeOAuth2Token**](Oauth2API.md#RevokeOAuth2Token) | **Post** /oauth2/revoke | OAuth2 トークン無効化エンドポイント



## CreateClient

> OAuth2ClientDetail CreateClient(ctx).PostClientRequest(postClientRequest).Execute()

OAuth2クライアントを作成



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
	postClientRequest := *traq.NewPostClientRequest("Name_example", "CallbackUrl_example", []traq.OAuth2Scope{traq.OAuth2Scope("openid")}, "Description_example") // PostClientRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.Oauth2API.CreateClient(context.Background()).PostClientRequest(postClientRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.CreateClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClient`: OAuth2ClientDetail
	fmt.Fprintf(os.Stdout, "Response from `Oauth2API.CreateClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postClientRequest** | [**PostClientRequest**](PostClientRequest.md) |  | 

### Return type

[**OAuth2ClientDetail**](OAuth2ClientDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClient

> DeleteClient(ctx, clientId).Execute()

OAuth2クライアントを削除



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
	clientId := "clientId_example" // string | OAuth2クライアントUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.DeleteClient(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.DeleteClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | OAuth2クライアントUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRequest struct via the builder pattern


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


## EditClient

> EditClient(ctx, clientId).PatchClientRequest(patchClientRequest).Execute()

OAuth2クライアント情報を変更



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
	clientId := "clientId_example" // string | OAuth2クライアントUUID
	patchClientRequest := *traq.NewPatchClientRequest() // PatchClientRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.EditClient(context.Background(), clientId).PatchClientRequest(patchClientRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.EditClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | OAuth2クライアントUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchClientRequest** | [**PatchClientRequest**](PatchClientRequest.md) |  | 

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


## GetClient

> GetClient200Response GetClient(ctx, clientId).Detail(detail).Execute()

OAuth2クライアント情報を取得



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
	clientId := "clientId_example" // string | OAuth2クライアントUUID
	detail := true // bool | 詳細情報を含めるかどうか (optional) (default to false)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.Oauth2API.GetClient(context.Background(), clientId).Detail(detail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.GetClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClient`: GetClient200Response
	fmt.Fprintf(os.Stdout, "Response from `Oauth2API.GetClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | OAuth2クライアントUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detail** | **bool** | 詳細情報を含めるかどうか | [default to false]

### Return type

[**GetClient200Response**](GetClient200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClients

> []OAuth2Client GetClients(ctx).All(all).Execute()

OAuth2クライアントのリストを取得



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
	all := true // bool | 全てのクライアントを取得するかどうか (optional) (default to false)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.Oauth2API.GetClients(context.Background()).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.GetClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClients`: []OAuth2Client
	fmt.Fprintf(os.Stdout, "Response from `Oauth2API.GetClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** | 全てのクライアントを取得するかどうか | [default to false]

### Return type

[**[]OAuth2Client**](OAuth2Client.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyTokens

> []ActiveOAuth2Token GetMyTokens(ctx).Execute()

有効トークンのリストを取得



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
	resp, r, err := apiClient.Oauth2API.GetMyTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.GetMyTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyTokens`: []ActiveOAuth2Token
	fmt.Fprintf(os.Stdout, "Response from `Oauth2API.GetMyTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyTokensRequest struct via the builder pattern


### Return type

[**[]ActiveOAuth2Token**](ActiveOAuth2Token.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuth2Authorize

> GetOAuth2Authorize(ctx).ClientId(clientId).ResponseType(responseType).RedirectUri(redirectUri).Scope(scope).State(state).CodeChallenge(codeChallenge).CodeChallengeMethod(codeChallengeMethod).Nonce(nonce).Prompt(prompt).Execute()

OAuth2 認可エンドポイント



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
	clientId := "clientId_example" // string | 
	responseType := traq.OAuth2ResponseType("code") // OAuth2ResponseType |  (optional)
	redirectUri := "redirectUri_example" // string |  (optional)
	scope := "scope_example" // string |  (optional)
	state := "state_example" // string |  (optional)
	codeChallenge := "codeChallenge_example" // string |  (optional)
	codeChallengeMethod := "codeChallengeMethod_example" // string |  (optional)
	nonce := "nonce_example" // string |  (optional)
	prompt := traq.OAuth2Prompt("none") // OAuth2Prompt |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.GetOAuth2Authorize(context.Background()).ClientId(clientId).ResponseType(responseType).RedirectUri(redirectUri).Scope(scope).State(state).CodeChallenge(codeChallenge).CodeChallengeMethod(codeChallengeMethod).Nonce(nonce).Prompt(prompt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.GetOAuth2Authorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuth2AuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **responseType** | [**OAuth2ResponseType**](OAuth2ResponseType.md) |  | 
 **redirectUri** | **string** |  | 
 **scope** | **string** |  | 
 **state** | **string** |  | 
 **codeChallenge** | **string** |  | 
 **codeChallengeMethod** | **string** |  | 
 **nonce** | **string** |  | 
 **prompt** | [**OAuth2Prompt**](OAuth2Prompt.md) |  | 

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


## PostOAuth2Authorize

> PostOAuth2Authorize(ctx).ClientId(clientId).ResponseType(responseType).RedirectUri(redirectUri).Scope(scope).State(state).CodeChallenge(codeChallenge).CodeChallengeMethod(codeChallengeMethod).Nonce(nonce).Prompt(prompt).Execute()

OAuth2 認可エンドポイント



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
	clientId := "clientId_example" // string | 
	responseType := traq.OAuth2ResponseType("code") // OAuth2ResponseType |  (optional)
	redirectUri := "redirectUri_example" // string |  (optional)
	scope := "scope_example" // string |  (optional)
	state := "state_example" // string |  (optional)
	codeChallenge := "codeChallenge_example" // string |  (optional)
	codeChallengeMethod := "codeChallengeMethod_example" // string |  (optional)
	nonce := "nonce_example" // string |  (optional)
	prompt := traq.OAuth2Prompt("none") // OAuth2Prompt |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.PostOAuth2Authorize(context.Background()).ClientId(clientId).ResponseType(responseType).RedirectUri(redirectUri).Scope(scope).State(state).CodeChallenge(codeChallenge).CodeChallengeMethod(codeChallengeMethod).Nonce(nonce).Prompt(prompt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.PostOAuth2Authorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOAuth2AuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **responseType** | [**OAuth2ResponseType**](OAuth2ResponseType.md) |  | 
 **redirectUri** | **string** |  | 
 **scope** | **string** |  | 
 **state** | **string** |  | 
 **codeChallenge** | **string** |  | 
 **codeChallengeMethod** | **string** |  | 
 **nonce** | **string** |  | 
 **prompt** | [**OAuth2Prompt**](OAuth2Prompt.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOAuth2AuthorizeDecide

> PostOAuth2AuthorizeDecide(ctx).Submit(submit).Execute()

OAuth2 認可承諾API



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
	submit := "submit_example" // string | 承諾する場合は\\\"approve\\\"

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.PostOAuth2AuthorizeDecide(context.Background()).Submit(submit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.PostOAuth2AuthorizeDecide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOAuth2AuthorizeDecideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submit** | **string** | 承諾する場合は\\\&quot;approve\\\&quot; | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOAuth2Token

> OAuth2Token PostOAuth2Token(ctx).GrantType(grantType).Code(code).RedirectUri(redirectUri).ClientId(clientId).CodeVerifier(codeVerifier).Username(username).Password(password).Scope(scope).RefreshToken(refreshToken).ClientSecret(clientSecret).Execute()

OAuth2 トークンエンドポイント



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
	grantType := "grantType_example" // string | 
	code := "code_example" // string |  (optional)
	redirectUri := "redirectUri_example" // string |  (optional)
	clientId := "clientId_example" // string |  (optional)
	codeVerifier := "codeVerifier_example" // string |  (optional)
	username := "username_example" // string |  (optional)
	password := "password_example" // string |  (optional)
	scope := "scope_example" // string |  (optional)
	refreshToken := "refreshToken_example" // string |  (optional)
	clientSecret := "clientSecret_example" // string |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.Oauth2API.PostOAuth2Token(context.Background()).GrantType(grantType).Code(code).RedirectUri(redirectUri).ClientId(clientId).CodeVerifier(codeVerifier).Username(username).Password(password).Scope(scope).RefreshToken(refreshToken).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.PostOAuth2Token``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOAuth2Token`: OAuth2Token
	fmt.Fprintf(os.Stdout, "Response from `Oauth2API.PostOAuth2Token`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOAuth2TokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **code** | **string** |  | 
 **redirectUri** | **string** |  | 
 **clientId** | **string** |  | 
 **codeVerifier** | **string** |  | 
 **username** | **string** |  | 
 **password** | **string** |  | 
 **scope** | **string** |  | 
 **refreshToken** | **string** |  | 
 **clientSecret** | **string** |  | 

### Return type

[**OAuth2Token**](OAuth2Token.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeClientTokens

> RevokeClientTokens(ctx, clientId).Execute()

OAuthクライアントのトークンを削除



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
	clientId := "clientId_example" // string | OAuth2クライアントUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.RevokeClientTokens(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.RevokeClientTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | OAuth2クライアントUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeClientTokensRequest struct via the builder pattern


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


## RevokeMyToken

> RevokeMyToken(ctx, tokenId).Execute()

トークンの認可を取り消す



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
	tokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | OAuth2トークンUUID

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.RevokeMyToken(context.Background(), tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.RevokeMyToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | OAuth2トークンUUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeMyTokenRequest struct via the builder pattern


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


## RevokeOAuth2Token

> RevokeOAuth2Token(ctx).Token(token).Execute()

OAuth2 トークン無効化エンドポイント



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
	token := "token_example" // string | 無効化するOAuth2トークンまたはOAuth2リフレッシュトークン

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	r, err := apiClient.Oauth2API.RevokeOAuth2Token(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2API.RevokeOAuth2Token``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOAuth2TokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | 無効化するOAuth2トークンまたはOAuth2リフレッシュトークン | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

