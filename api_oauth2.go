/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// Oauth2ApiService Oauth2Api service
type Oauth2ApiService service

// Oauth2ApiCreateClientOpts Optional parameters for the method 'CreateClient'
type Oauth2ApiCreateClientOpts struct {
	PostClientRequest optional.Interface
}

/*
CreateClient OAuth2クライアントを作成
OAuth2クライアントを作成します。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *Oauth2ApiCreateClientOpts - Optional Parameters:
 * @param "PostClientRequest" (optional.Interface of PostClientRequest) -
@return OAuth2ClientDetail
*/
func (a *Oauth2ApiService) CreateClient(ctx _context.Context, localVarOptionals *Oauth2ApiCreateClientOpts) (OAuth2ClientDetail, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OAuth2ClientDetail
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clients"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.PostClientRequest.IsSet() {
		localVarOptionalPostClientRequest, localVarOptionalPostClientRequestok := localVarOptionals.PostClientRequest.Value().(PostClientRequest)
		if !localVarOptionalPostClientRequestok {
			return localVarReturnValue, nil, reportError("postClientRequest should be PostClientRequest")
		}
		localVarPostBody = &localVarOptionalPostClientRequest
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
DeleteClient OAuth2クライアントを削除
指定したOAuth2クライアントを削除します。 対象のクライアントの管理権限が必要です。正常に削除された場合、このクライアントに対する認可は全て取り消されます。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clientId OAuth2クライアントUUID
*/
func (a *Oauth2ApiService) DeleteClient(ctx _context.Context, clientId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clients/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", _neturl.QueryEscape(parameterToString(clientId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// Oauth2ApiEditClientOpts Optional parameters for the method 'EditClient'
type Oauth2ApiEditClientOpts struct {
	PatchClientRequest optional.Interface
}

/*
EditClient OAuth2クライアント情報を変更
指定したOAuth2クライアントの情報を変更します。 対象のクライアントの管理権限が必要です。 クライアント開発者UUIDを変更した場合は、変更先ユーザーにクライアント管理権限が移譲され、自分自身は権限を失います。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clientId OAuth2クライアントUUID
 * @param optional nil or *Oauth2ApiEditClientOpts - Optional Parameters:
 * @param "PatchClientRequest" (optional.Interface of PatchClientRequest) -
*/
func (a *Oauth2ApiService) EditClient(ctx _context.Context, clientId string, localVarOptionals *Oauth2ApiEditClientOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clients/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", _neturl.QueryEscape(parameterToString(clientId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.PatchClientRequest.IsSet() {
		localVarOptionalPatchClientRequest, localVarOptionalPatchClientRequestok := localVarOptionals.PatchClientRequest.Value().(PatchClientRequest)
		if !localVarOptionalPatchClientRequestok {
			return nil, reportError("patchClientRequest should be PatchClientRequest")
		}
		localVarPostBody = &localVarOptionalPatchClientRequest
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// Oauth2ApiGetClientOpts Optional parameters for the method 'GetClient'
type Oauth2ApiGetClientOpts struct {
	Detail optional.Bool
}

/*
GetClient OAuth2クライアント情報を取得
指定したOAuth2クライアントの情報を取得します。 詳細情報の取得には対象のクライアントの管理権限が必要です。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clientId OAuth2クライアントUUID
 * @param optional nil or *Oauth2ApiGetClientOpts - Optional Parameters:
 * @param "Detail" (optional.Bool) -  詳細情報を含めるかどうか
@return OneOf: OAuth2Client, OAuth2ClientDetail
*/
func (a *Oauth2ApiService) GetClient(ctx _context.Context, clientId string, localVarOptionals *Oauth2ApiGetClientOpts) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clients/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", _neturl.QueryEscape(parameterToString(clientId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Detail.IsSet() {
		localVarQueryParams.Add("detail", parameterToString(localVarOptionals.Detail.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// Oauth2ApiGetClientsOpts Optional parameters for the method 'GetClients'
type Oauth2ApiGetClientsOpts struct {
	All optional.Bool
}

/*
GetClients OAuth2クライアントのリストを取得
自身が開発者のOAuth2クライアントのリストを取得します。 &#x60;all&#x60;が&#x60;true&#x60;の場合、全開発者の全クライアントのリストを返します。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *Oauth2ApiGetClientsOpts - Optional Parameters:
 * @param "All" (optional.Bool) -  全てのクライアントを取得するかどうか
@return []OAuth2Client
*/
func (a *Oauth2ApiService) GetClients(ctx _context.Context, localVarOptionals *Oauth2ApiGetClientsOpts) ([]OAuth2Client, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []OAuth2Client
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clients"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.All.IsSet() {
		localVarQueryParams.Add("all", parameterToString(localVarOptionals.All.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetMyTokens 有効トークンのリストを取得
有効な自分に発行されたOAuth2トークンのリストを取得します。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return []ActiveOAuth2Token
*/
func (a *Oauth2ApiService) GetMyTokens(ctx _context.Context) ([]ActiveOAuth2Token, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ActiveOAuth2Token
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/me/tokens"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// Oauth2ApiGetOAuth2AuthorizeOpts Optional parameters for the method 'GetOAuth2Authorize'
type Oauth2ApiGetOAuth2AuthorizeOpts struct {
	ResponseType        optional.Interface
	RedirectUri         optional.String
	Scope               optional.String
	State               optional.String
	CodeChallenge       optional.String
	CodeChallengeMethod optional.String
	Nonce               optional.String
	Prompt              optional.Interface
}

/*
GetOAuth2Authorize OAuth2 認可エンドポイント
OAuth2 認可エンドポイント
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clientId
 * @param optional nil or *Oauth2ApiGetOAuth2AuthorizeOpts - Optional Parameters:
 * @param "ResponseType" (optional.Interface of OAuth2ResponseType) -
 * @param "RedirectUri" (optional.String) -
 * @param "Scope" (optional.String) -
 * @param "State" (optional.String) -
 * @param "CodeChallenge" (optional.String) -
 * @param "CodeChallengeMethod" (optional.String) -
 * @param "Nonce" (optional.String) -
 * @param "Prompt" (optional.Interface of OAuth2Prompt) -
*/
func (a *Oauth2ApiService) GetOAuth2Authorize(ctx _context.Context, clientId string, localVarOptionals *Oauth2ApiGetOAuth2AuthorizeOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/authorize"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.ResponseType.IsSet() {
		localVarQueryParams.Add("response_type", parameterToString(localVarOptionals.ResponseType.Value(), ""))
	}
	localVarQueryParams.Add("client_id", parameterToString(clientId, ""))
	if localVarOptionals != nil && localVarOptionals.RedirectUri.IsSet() {
		localVarQueryParams.Add("redirect_uri", parameterToString(localVarOptionals.RedirectUri.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Scope.IsSet() {
		localVarQueryParams.Add("scope", parameterToString(localVarOptionals.Scope.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarQueryParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeChallenge.IsSet() {
		localVarQueryParams.Add("code_challenge", parameterToString(localVarOptionals.CodeChallenge.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeChallengeMethod.IsSet() {
		localVarQueryParams.Add("code_challenge_method", parameterToString(localVarOptionals.CodeChallengeMethod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nonce.IsSet() {
		localVarQueryParams.Add("nonce", parameterToString(localVarOptionals.Nonce.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prompt.IsSet() {
		localVarQueryParams.Add("prompt", parameterToString(localVarOptionals.Prompt.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// Oauth2ApiPostOAuth2AuthorizeOpts Optional parameters for the method 'PostOAuth2Authorize'
type Oauth2ApiPostOAuth2AuthorizeOpts struct {
	ResponseType        optional.Interface
	RedirectUri         optional.String
	Scope               optional.String
	State               optional.String
	CodeChallenge       optional.String
	CodeChallengeMethod optional.String
	Nonce               optional.String
	Prompt              optional.Interface
}

/*
PostOAuth2Authorize OAuth2 認可エンドポイント
OAuth2 認可エンドポイント
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clientId
 * @param optional nil or *Oauth2ApiPostOAuth2AuthorizeOpts - Optional Parameters:
 * @param "ResponseType" (optional.Interface of OAuth2ResponseType) -
 * @param "RedirectUri" (optional.String) -
 * @param "Scope" (optional.String) -
 * @param "State" (optional.String) -
 * @param "CodeChallenge" (optional.String) -
 * @param "CodeChallengeMethod" (optional.String) -
 * @param "Nonce" (optional.String) -
 * @param "Prompt" (optional.Interface of OAuth2Prompt) -
*/
func (a *Oauth2ApiService) PostOAuth2Authorize(ctx _context.Context, clientId string, localVarOptionals *Oauth2ApiPostOAuth2AuthorizeOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/authorize"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ResponseType.IsSet() {
		localVarFormParams.Add("response_type", parameterToString(localVarOptionals.ResponseType.Value(), ""))
	}
	localVarFormParams.Add("client_id", parameterToString(clientId, ""))
	if localVarOptionals != nil && localVarOptionals.RedirectUri.IsSet() {
		localVarFormParams.Add("redirect_uri", parameterToString(localVarOptionals.RedirectUri.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Scope.IsSet() {
		localVarFormParams.Add("scope", parameterToString(localVarOptionals.Scope.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarFormParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeChallenge.IsSet() {
		localVarFormParams.Add("code_challenge", parameterToString(localVarOptionals.CodeChallenge.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeChallengeMethod.IsSet() {
		localVarFormParams.Add("code_challenge_method", parameterToString(localVarOptionals.CodeChallengeMethod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nonce.IsSet() {
		localVarFormParams.Add("nonce", parameterToString(localVarOptionals.Nonce.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prompt.IsSet() {
		localVarFormParams.Add("prompt", parameterToString(localVarOptionals.Prompt.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
PostOAuth2AuthorizeDecide OAuth2 認可承諾API
OAuth2 認可承諾
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param submit 承諾する場合は\\\"approve\\\"
*/
func (a *Oauth2ApiService) PostOAuth2AuthorizeDecide(ctx _context.Context, submit string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/authorize/decide"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("submit", parameterToString(submit, ""))
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// Oauth2ApiPostOAuth2TokenOpts Optional parameters for the method 'PostOAuth2Token'
type Oauth2ApiPostOAuth2TokenOpts struct {
	Code         optional.String
	RedirectUri  optional.String
	ClientId     optional.String
	CodeVerifier optional.String
	Username     optional.String
	Password     optional.String
	Scope        optional.String
	RefreshToken optional.String
	ClientSecret optional.String
}

/*
PostOAuth2Token OAuth2 トークンエンドポイント
OAuth2 トークンエンドポイント
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param grantType
 * @param optional nil or *Oauth2ApiPostOAuth2TokenOpts - Optional Parameters:
 * @param "Code" (optional.String) -
 * @param "RedirectUri" (optional.String) -
 * @param "ClientId" (optional.String) -
 * @param "CodeVerifier" (optional.String) -
 * @param "Username" (optional.String) -
 * @param "Password" (optional.String) -
 * @param "Scope" (optional.String) -
 * @param "RefreshToken" (optional.String) -
 * @param "ClientSecret" (optional.String) -
@return OAuth2Token
*/
func (a *Oauth2ApiService) PostOAuth2Token(ctx _context.Context, grantType string, localVarOptionals *Oauth2ApiPostOAuth2TokenOpts) (OAuth2Token, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OAuth2Token
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/token"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("grant_type", parameterToString(grantType, ""))
	if localVarOptionals != nil && localVarOptionals.Code.IsSet() {
		localVarFormParams.Add("code", parameterToString(localVarOptionals.Code.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RedirectUri.IsSet() {
		localVarFormParams.Add("redirect_uri", parameterToString(localVarOptionals.RedirectUri.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientId.IsSet() {
		localVarFormParams.Add("client_id", parameterToString(localVarOptionals.ClientId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeVerifier.IsSet() {
		localVarFormParams.Add("code_verifier", parameterToString(localVarOptionals.CodeVerifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Username.IsSet() {
		localVarFormParams.Add("username", parameterToString(localVarOptionals.Username.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Password.IsSet() {
		localVarFormParams.Add("password", parameterToString(localVarOptionals.Password.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Scope.IsSet() {
		localVarFormParams.Add("scope", parameterToString(localVarOptionals.Scope.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RefreshToken.IsSet() {
		localVarFormParams.Add("refresh_token", parameterToString(localVarOptionals.RefreshToken.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientSecret.IsSet() {
		localVarFormParams.Add("client_secret", parameterToString(localVarOptionals.ClientSecret.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
RevokeMyToken トークンの認可を取り消す
自分の指定したトークンの認可を取り消します。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tokenId OAuth2トークンUUID
*/
func (a *Oauth2ApiService) RevokeMyToken(ctx _context.Context, tokenId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/me/tokens/{tokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tokenId"+"}", _neturl.QueryEscape(parameterToString(tokenId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
RevokeOAuth2Token OAuth2 トークン無効化エンドポイント
OAuth2 トークン無効化エンドポイント
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param token 無効化するOAuth2トークンまたはOAuth2リフレッシュトークン
*/
func (a *Oauth2ApiService) RevokeOAuth2Token(ctx _context.Context, token string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/revoke"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("token", parameterToString(token, ""))
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
