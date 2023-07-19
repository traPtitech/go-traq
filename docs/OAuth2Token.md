# OAuth2Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**TokenType** | **string** |  | 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**IdToken** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuth2Token

`func NewOAuth2Token(accessToken string, tokenType string, ) *OAuth2Token`

NewOAuth2Token instantiates a new OAuth2Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2TokenWithDefaults

`func NewOAuth2TokenWithDefaults() *OAuth2Token`

NewOAuth2TokenWithDefaults instantiates a new OAuth2Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OAuth2Token) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuth2Token) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuth2Token) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *OAuth2Token) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *OAuth2Token) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *OAuth2Token) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *OAuth2Token) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *OAuth2Token) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *OAuth2Token) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *OAuth2Token) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRefreshToken

`func (o *OAuth2Token) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuth2Token) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuth2Token) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OAuth2Token) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetScope

`func (o *OAuth2Token) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuth2Token) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuth2Token) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuth2Token) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetIdToken

`func (o *OAuth2Token) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *OAuth2Token) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *OAuth2Token) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *OAuth2Token) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


