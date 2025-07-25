/*
traQ v3

traQ v3 API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package traq

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AuthenticationAPIService AuthenticationAPI service
type AuthenticationAPIService service

type AuthenticationAPIGetMyExternalAccountsRequest struct {
	ctx        context.Context
	ApiService *AuthenticationAPIService
}

func (r AuthenticationAPIGetMyExternalAccountsRequest) Execute() ([]ExternalProviderUser, *http.Response, error) {
	return r.ApiService.GetMyExternalAccountsExecute(r)
}

/*
GetMyExternalAccounts 外部ログインアカウント一覧を取得

自分に紐付けられている外部ログインアカウント一覧を取得します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AuthenticationAPIGetMyExternalAccountsRequest
*/
func (a *AuthenticationAPIService) GetMyExternalAccounts(ctx context.Context) AuthenticationAPIGetMyExternalAccountsRequest {
	return AuthenticationAPIGetMyExternalAccountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ExternalProviderUser
func (a *AuthenticationAPIService) GetMyExternalAccountsExecute(r AuthenticationAPIGetMyExternalAccountsRequest) ([]ExternalProviderUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ExternalProviderUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIService.GetMyExternalAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/ex-accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type AuthenticationAPIGetMySessionsRequest struct {
	ctx        context.Context
	ApiService *AuthenticationAPIService
}

func (r AuthenticationAPIGetMySessionsRequest) Execute() ([]LoginSession, *http.Response, error) {
	return r.ApiService.GetMySessionsExecute(r)
}

/*
GetMySessions 自分のログインセッションリストを取得

自分のログインセッションのリストを取得します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AuthenticationAPIGetMySessionsRequest
*/
func (a *AuthenticationAPIService) GetMySessions(ctx context.Context) AuthenticationAPIGetMySessionsRequest {
	return AuthenticationAPIGetMySessionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []LoginSession
func (a *AuthenticationAPIService) GetMySessionsExecute(r AuthenticationAPIGetMySessionsRequest) ([]LoginSession, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []LoginSession
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIService.GetMySessions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/sessions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type AuthenticationAPILinkExternalAccountRequest struct {
	ctx                     context.Context
	ApiService              *AuthenticationAPIService
	postLinkExternalAccount *PostLinkExternalAccount
}

func (r AuthenticationAPILinkExternalAccountRequest) PostLinkExternalAccount(postLinkExternalAccount PostLinkExternalAccount) AuthenticationAPILinkExternalAccountRequest {
	r.postLinkExternalAccount = &postLinkExternalAccount
	return r
}

func (r AuthenticationAPILinkExternalAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.LinkExternalAccountExecute(r)
}

/*
LinkExternalAccount 外部ログインアカウントを紐付ける

自分に外部ログインアカウントを紐付けます。
指定した`providerName`がサーバー側で有効である必要があります。
リクエストが受理された場合、外部サービスの認証画面にリダイレクトされ、認証される必要があります。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AuthenticationAPILinkExternalAccountRequest
*/
func (a *AuthenticationAPIService) LinkExternalAccount(ctx context.Context) AuthenticationAPILinkExternalAccountRequest {
	return AuthenticationAPILinkExternalAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AuthenticationAPIService) LinkExternalAccountExecute(r AuthenticationAPILinkExternalAccountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIService.LinkExternalAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/ex-accounts/link"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.postLinkExternalAccount
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type AuthenticationAPILoginRequest struct {
	ctx              context.Context
	ApiService       *AuthenticationAPIService
	redirect         *string
	postLoginRequest *PostLoginRequest
}

// リダイレクト先
func (r AuthenticationAPILoginRequest) Redirect(redirect string) AuthenticationAPILoginRequest {
	r.redirect = &redirect
	return r
}

func (r AuthenticationAPILoginRequest) PostLoginRequest(postLoginRequest PostLoginRequest) AuthenticationAPILoginRequest {
	r.postLoginRequest = &postLoginRequest
	return r
}

func (r AuthenticationAPILoginRequest) Execute() (*http.Response, error) {
	return r.ApiService.LoginExecute(r)
}

/*
Login ログイン

ログインします。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AuthenticationAPILoginRequest
*/
func (a *AuthenticationAPIService) Login(ctx context.Context) AuthenticationAPILoginRequest {
	return AuthenticationAPILoginRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AuthenticationAPIService) LoginExecute(r AuthenticationAPILoginRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIService.Login")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.redirect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "redirect", r.redirect, "form", "")
	}
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
	localVarPostBody = r.postLoginRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type AuthenticationAPILogoutRequest struct {
	ctx        context.Context
	ApiService *AuthenticationAPIService
	redirect   *string
	all        *bool
}

// リダイレクト先
func (r AuthenticationAPILogoutRequest) Redirect(redirect string) AuthenticationAPILogoutRequest {
	r.redirect = &redirect
	return r
}

// 全てのセッションでログアウトするかどうか
func (r AuthenticationAPILogoutRequest) All(all bool) AuthenticationAPILogoutRequest {
	r.all = &all
	return r
}

func (r AuthenticationAPILogoutRequest) Execute() (*http.Response, error) {
	return r.ApiService.LogoutExecute(r)
}

/*
Logout ログアウト

ログアウトします。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AuthenticationAPILogoutRequest
*/
func (a *AuthenticationAPIService) Logout(ctx context.Context) AuthenticationAPILogoutRequest {
	return AuthenticationAPILogoutRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AuthenticationAPIService) LogoutExecute(r AuthenticationAPILogoutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIService.Logout")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logout"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.redirect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "redirect", r.redirect, "form", "")
	}
	if r.all != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "all", r.all, "form", "")
	} else {
		var defaultValue bool = false
		r.all = &defaultValue
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type AuthenticationAPIRevokeMySessionRequest struct {
	ctx        context.Context
	ApiService *AuthenticationAPIService
	sessionId  string
}

func (r AuthenticationAPIRevokeMySessionRequest) Execute() (*http.Response, error) {
	return r.ApiService.RevokeMySessionExecute(r)
}

/*
RevokeMySession セッションを無効化

指定した自分のセッションを無効化(ログアウト)します。
既に存在しない・無効化されているセッションを指定した場合も`204`を返します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sessionId セッションUUID
	@return AuthenticationAPIRevokeMySessionRequest
*/
func (a *AuthenticationAPIService) RevokeMySession(ctx context.Context, sessionId string) AuthenticationAPIRevokeMySessionRequest {
	return AuthenticationAPIRevokeMySessionRequest{
		ApiService: a,
		ctx:        ctx,
		sessionId:  sessionId,
	}
}

// Execute executes the request
func (a *AuthenticationAPIService) RevokeMySessionExecute(r AuthenticationAPIRevokeMySessionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIService.RevokeMySession")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/sessions/{sessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterValueToString(r.sessionId, "sessionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type AuthenticationAPIUnlinkExternalAccountRequest struct {
	ctx                       context.Context
	ApiService                *AuthenticationAPIService
	postUnlinkExternalAccount *PostUnlinkExternalAccount
}

func (r AuthenticationAPIUnlinkExternalAccountRequest) PostUnlinkExternalAccount(postUnlinkExternalAccount PostUnlinkExternalAccount) AuthenticationAPIUnlinkExternalAccountRequest {
	r.postUnlinkExternalAccount = &postUnlinkExternalAccount
	return r
}

func (r AuthenticationAPIUnlinkExternalAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnlinkExternalAccountExecute(r)
}

/*
UnlinkExternalAccount 外部ログインアカウントの紐付けを解除

自分に紐付けられている外部ログインアカウントの紐付けを解除します。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AuthenticationAPIUnlinkExternalAccountRequest
*/
func (a *AuthenticationAPIService) UnlinkExternalAccount(ctx context.Context) AuthenticationAPIUnlinkExternalAccountRequest {
	return AuthenticationAPIUnlinkExternalAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AuthenticationAPIService) UnlinkExternalAccountExecute(r AuthenticationAPIUnlinkExternalAccountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIService.UnlinkExternalAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/ex-accounts/unlink"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.postUnlinkExternalAccount
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
