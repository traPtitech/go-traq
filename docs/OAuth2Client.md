# OAuth2Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | クライアントUUID | 
**Name** | **string** | クライアント名 | 
**Description** | **string** | 説明 | 
**DeveloperId** | **string** | クライアント開発者UUID | 
**Scopes** | [**[]OAuth2Scope**](OAuth2Scope.md) | 要求スコープの配列 | 

## Methods

### NewOAuth2Client

`func NewOAuth2Client(id string, name string, description string, developerId string, scopes []OAuth2Scope, ) *OAuth2Client`

NewOAuth2Client instantiates a new OAuth2Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientWithDefaults

`func NewOAuth2ClientWithDefaults() *OAuth2Client`

NewOAuth2ClientWithDefaults instantiates a new OAuth2Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OAuth2Client) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2Client) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2Client) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OAuth2Client) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuth2Client) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuth2Client) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OAuth2Client) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OAuth2Client) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OAuth2Client) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDeveloperId

`func (o *OAuth2Client) GetDeveloperId() string`

GetDeveloperId returns the DeveloperId field if non-nil, zero value otherwise.

### GetDeveloperIdOk

`func (o *OAuth2Client) GetDeveloperIdOk() (*string, bool)`

GetDeveloperIdOk returns a tuple with the DeveloperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperId

`func (o *OAuth2Client) SetDeveloperId(v string)`

SetDeveloperId sets DeveloperId field to given value.


### GetScopes

`func (o *OAuth2Client) GetScopes() []OAuth2Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuth2Client) GetScopesOk() (*[]OAuth2Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuth2Client) SetScopes(v []OAuth2Scope)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


