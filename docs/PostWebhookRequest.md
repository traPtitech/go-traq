# PostWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Webhookユーザーの表示名 | 
**Description** | **string** | 説明 | 
**ChannelId** | **string** | デフォルトの投稿先チャンネルUUID | 
**Secret** | **string** | Webhookシークレット | 

## Methods

### NewPostWebhookRequest

`func NewPostWebhookRequest(name string, description string, channelId string, secret string, ) *PostWebhookRequest`

NewPostWebhookRequest instantiates a new PostWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWebhookRequestWithDefaults

`func NewPostWebhookRequestWithDefaults() *PostWebhookRequest`

NewPostWebhookRequestWithDefaults instantiates a new PostWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostWebhookRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetChannelId

`func (o *PostWebhookRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PostWebhookRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PostWebhookRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetSecret

`func (o *PostWebhookRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PostWebhookRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PostWebhookRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


