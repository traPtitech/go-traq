# PatchWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Webhookユーザー表示名 | [optional] 
**Description** | Pointer to **string** | 説明 | [optional] 
**ChannelId** | Pointer to **string** | デフォルトの投稿先チャンネルUUID | [optional] 
**Secret** | Pointer to **string** | Webhookシークレット | [optional] 
**OwnerId** | Pointer to **string** | 移譲先のユーザーUUID | [optional] 

## Methods

### NewPatchWebhookRequest

`func NewPatchWebhookRequest() *PatchWebhookRequest`

NewPatchWebhookRequest instantiates a new PatchWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchWebhookRequestWithDefaults

`func NewPatchWebhookRequestWithDefaults() *PatchWebhookRequest`

NewPatchWebhookRequestWithDefaults instantiates a new PatchWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchWebhookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChannelId

`func (o *PatchWebhookRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PatchWebhookRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PatchWebhookRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *PatchWebhookRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetSecret

`func (o *PatchWebhookRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchWebhookRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchWebhookRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchWebhookRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetOwnerId

`func (o *PatchWebhookRequest) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *PatchWebhookRequest) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *PatchWebhookRequest) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *PatchWebhookRequest) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


