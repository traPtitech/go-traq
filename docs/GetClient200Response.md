# GetClient200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | クライアントUUID | 
**Name** | **string** | クライアント名 | 
**Description** | **string** | 説明 | 
**DeveloperId** | **string** | クライアント開発者UUID | 
**Scopes** | [**[]OAuth2Scope**](OAuth2Scope.md) | 要求スコープの配列 | 
**Confidential** | **bool** | confidential client なら true, public client なら false | 
**CallbackUrl** | **string** | コールバックURL | 
**Secret** | **string** | クライアントシークレット | 

## Methods

### NewGetClient200Response

`func NewGetClient200Response(id string, name string, description string, developerId string, scopes []OAuth2Scope, confidential bool, callbackUrl string, secret string, ) *GetClient200Response`

NewGetClient200Response instantiates a new GetClient200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClient200ResponseWithDefaults

`func NewGetClient200ResponseWithDefaults() *GetClient200Response`

NewGetClient200ResponseWithDefaults instantiates a new GetClient200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClient200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClient200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClient200Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetClient200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClient200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClient200Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetClient200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetClient200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetClient200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDeveloperId

`func (o *GetClient200Response) GetDeveloperId() string`

GetDeveloperId returns the DeveloperId field if non-nil, zero value otherwise.

### GetDeveloperIdOk

`func (o *GetClient200Response) GetDeveloperIdOk() (*string, bool)`

GetDeveloperIdOk returns a tuple with the DeveloperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperId

`func (o *GetClient200Response) SetDeveloperId(v string)`

SetDeveloperId sets DeveloperId field to given value.


### GetScopes

`func (o *GetClient200Response) GetScopes() []OAuth2Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetClient200Response) GetScopesOk() (*[]OAuth2Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetClient200Response) SetScopes(v []OAuth2Scope)`

SetScopes sets Scopes field to given value.


### GetConfidential

`func (o *GetClient200Response) GetConfidential() bool`

GetConfidential returns the Confidential field if non-nil, zero value otherwise.

### GetConfidentialOk

`func (o *GetClient200Response) GetConfidentialOk() (*bool, bool)`

GetConfidentialOk returns a tuple with the Confidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidential

`func (o *GetClient200Response) SetConfidential(v bool)`

SetConfidential sets Confidential field to given value.


### GetCallbackUrl

`func (o *GetClient200Response) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *GetClient200Response) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *GetClient200Response) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetSecret

`func (o *GetClient200Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GetClient200Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GetClient200Response) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


