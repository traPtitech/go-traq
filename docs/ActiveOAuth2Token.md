# ActiveOAuth2Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | トークンUUID | 
**ClientId** | **string** | OAuth2クライアントUUID | 
**Scopes** | [**[]OAuth2Scope**](OAuth2Scope.md) | スコープ | 
**IssuedAt** | **time.Time** | 発行日時 | 

## Methods

### NewActiveOAuth2Token

`func NewActiveOAuth2Token(id string, clientId string, scopes []OAuth2Scope, issuedAt time.Time, ) *ActiveOAuth2Token`

NewActiveOAuth2Token instantiates a new ActiveOAuth2Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveOAuth2TokenWithDefaults

`func NewActiveOAuth2TokenWithDefaults() *ActiveOAuth2Token`

NewActiveOAuth2TokenWithDefaults instantiates a new ActiveOAuth2Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActiveOAuth2Token) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveOAuth2Token) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveOAuth2Token) SetId(v string)`

SetId sets Id field to given value.


### GetClientId

`func (o *ActiveOAuth2Token) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ActiveOAuth2Token) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ActiveOAuth2Token) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetScopes

`func (o *ActiveOAuth2Token) GetScopes() []OAuth2Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ActiveOAuth2Token) GetScopesOk() (*[]OAuth2Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ActiveOAuth2Token) SetScopes(v []OAuth2Scope)`

SetScopes sets Scopes field to given value.


### GetIssuedAt

`func (o *ActiveOAuth2Token) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *ActiveOAuth2Token) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *ActiveOAuth2Token) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


