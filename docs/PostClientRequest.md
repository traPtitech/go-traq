# PostClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | クライアント名 | 
**CallbackUrl** | **string** | コールバックURL | 
**Scopes** | [**[]OAuth2Scope**](OAuth2Scope.md) | 要求スコープの配列 | 
**Description** | **string** | 説明 | 

## Methods

### NewPostClientRequest

`func NewPostClientRequest(name string, callbackUrl string, scopes []OAuth2Scope, description string, ) *PostClientRequest`

NewPostClientRequest instantiates a new PostClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostClientRequestWithDefaults

`func NewPostClientRequestWithDefaults() *PostClientRequest`

NewPostClientRequestWithDefaults instantiates a new PostClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostClientRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostClientRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostClientRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCallbackUrl

`func (o *PostClientRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *PostClientRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *PostClientRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetScopes

`func (o *PostClientRequest) GetScopes() []OAuth2Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostClientRequest) GetScopesOk() (*[]OAuth2Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostClientRequest) SetScopes(v []OAuth2Scope)`

SetScopes sets Scopes field to given value.


### GetDescription

`func (o *PostClientRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostClientRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostClientRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


