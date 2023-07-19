# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | WebhookUUID | 
**BotUserId** | **string** | WebhookユーザーUUID | 
**DisplayName** | **string** | Webhookユーザー表示名 | 
**Description** | **string** | 説明 | 
**Secure** | **bool** | セキュアWebhookかどうか | 
**ChannelId** | **string** | デフォルトの投稿先チャンネルUUID | 
**OwnerId** | **string** | オーナーUUID | 
**CreatedAt** | **time.Time** | 作成日時 | 
**UpdatedAt** | **time.Time** | 更新日時 | 

## Methods

### NewWebhook

`func NewWebhook(id string, botUserId string, displayName string, description string, secure bool, channelId string, ownerId string, createdAt time.Time, updatedAt time.Time, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.


### GetBotUserId

`func (o *Webhook) GetBotUserId() string`

GetBotUserId returns the BotUserId field if non-nil, zero value otherwise.

### GetBotUserIdOk

`func (o *Webhook) GetBotUserIdOk() (*string, bool)`

GetBotUserIdOk returns a tuple with the BotUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUserId

`func (o *Webhook) SetBotUserId(v string)`

SetBotUserId sets BotUserId field to given value.


### GetDisplayName

`func (o *Webhook) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Webhook) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Webhook) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *Webhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Webhook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Webhook) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSecure

`func (o *Webhook) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *Webhook) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *Webhook) SetSecure(v bool)`

SetSecure sets Secure field to given value.


### GetChannelId

`func (o *Webhook) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *Webhook) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *Webhook) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetOwnerId

`func (o *Webhook) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Webhook) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Webhook) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetCreatedAt

`func (o *Webhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Webhook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Webhook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Webhook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


