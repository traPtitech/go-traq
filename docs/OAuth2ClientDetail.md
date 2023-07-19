# OAuth2ClientDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | クライアントUUID | 
**DeveloperId** | **string** | クライアント開発者UUID | 
**Description** | **string** | 説明 | 
**Name** | **string** | クライアント名 | 
**Scopes** | [**[]OAuth2Scope**](OAuth2Scope.md) | 要求スコープの配列 | 
**CallbackUrl** | **string** | コールバックURL | 
**Secret** | **string** | クライアントシークレット | 

## Methods

### NewOAuth2ClientDetail

`func NewOAuth2ClientDetail(id string, developerId string, description string, name string, scopes []OAuth2Scope, callbackUrl string, secret string, ) *OAuth2ClientDetail`

NewOAuth2ClientDetail instantiates a new OAuth2ClientDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientDetailWithDefaults

`func NewOAuth2ClientDetailWithDefaults() *OAuth2ClientDetail`

NewOAuth2ClientDetailWithDefaults instantiates a new OAuth2ClientDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OAuth2ClientDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2ClientDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2ClientDetail) SetId(v string)`

SetId sets Id field to given value.


### GetDeveloperId

`func (o *OAuth2ClientDetail) GetDeveloperId() string`

GetDeveloperId returns the DeveloperId field if non-nil, zero value otherwise.

### GetDeveloperIdOk

`func (o *OAuth2ClientDetail) GetDeveloperIdOk() (*string, bool)`

GetDeveloperIdOk returns a tuple with the DeveloperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperId

`func (o *OAuth2ClientDetail) SetDeveloperId(v string)`

SetDeveloperId sets DeveloperId field to given value.


### GetDescription

`func (o *OAuth2ClientDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OAuth2ClientDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OAuth2ClientDetail) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *OAuth2ClientDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuth2ClientDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuth2ClientDetail) SetName(v string)`

SetName sets Name field to given value.


### GetScopes

`func (o *OAuth2ClientDetail) GetScopes() []OAuth2Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuth2ClientDetail) GetScopesOk() (*[]OAuth2Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuth2ClientDetail) SetScopes(v []OAuth2Scope)`

SetScopes sets Scopes field to given value.


### GetCallbackUrl

`func (o *OAuth2ClientDetail) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *OAuth2ClientDetail) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *OAuth2ClientDetail) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetSecret

`func (o *OAuth2ClientDetail) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *OAuth2ClientDetail) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *OAuth2ClientDetail) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


