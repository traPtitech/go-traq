# BotDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | BOT UUID | 
**UpdatedAt** | **time.Time** | 更新日時 | 
**CreatedAt** | **time.Time** | 作成日時 | 
**Mode** | [**BotMode**](BotMode.md) |  | 
**State** | [**BotState**](BotState.md) |  | 
**SubscribeEvents** | **[]string** | BOTが購読しているイベントの配列 | 
**DeveloperId** | **string** | BOT開発者UUID | 
**Description** | **string** | 説明 | 
**BotUserId** | **string** | BOTユーザーUUID | 
**Tokens** | [**BotTokens**](BotTokens.md) |  | 
**Endpoint** | **string** | BOTサーバーエンドポイント | 
**Privileged** | **bool** | 特権BOTかどうか | 
**Channels** | **[]string** | BOTが参加しているチャンネルのUUID配列 | 

## Methods

### NewBotDetail

`func NewBotDetail(id string, updatedAt time.Time, createdAt time.Time, mode BotMode, state BotState, subscribeEvents []string, developerId string, description string, botUserId string, tokens BotTokens, endpoint string, privileged bool, channels []string, ) *BotDetail`

NewBotDetail instantiates a new BotDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotDetailWithDefaults

`func NewBotDetailWithDefaults() *BotDetail`

NewBotDetailWithDefaults instantiates a new BotDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BotDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BotDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BotDetail) SetId(v string)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *BotDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BotDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BotDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *BotDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BotDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BotDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMode

`func (o *BotDetail) GetMode() BotMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BotDetail) GetModeOk() (*BotMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BotDetail) SetMode(v BotMode)`

SetMode sets Mode field to given value.


### GetState

`func (o *BotDetail) GetState() BotState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BotDetail) GetStateOk() (*BotState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BotDetail) SetState(v BotState)`

SetState sets State field to given value.


### GetSubscribeEvents

`func (o *BotDetail) GetSubscribeEvents() []string`

GetSubscribeEvents returns the SubscribeEvents field if non-nil, zero value otherwise.

### GetSubscribeEventsOk

`func (o *BotDetail) GetSubscribeEventsOk() (*[]string, bool)`

GetSubscribeEventsOk returns a tuple with the SubscribeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeEvents

`func (o *BotDetail) SetSubscribeEvents(v []string)`

SetSubscribeEvents sets SubscribeEvents field to given value.


### GetDeveloperId

`func (o *BotDetail) GetDeveloperId() string`

GetDeveloperId returns the DeveloperId field if non-nil, zero value otherwise.

### GetDeveloperIdOk

`func (o *BotDetail) GetDeveloperIdOk() (*string, bool)`

GetDeveloperIdOk returns a tuple with the DeveloperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperId

`func (o *BotDetail) SetDeveloperId(v string)`

SetDeveloperId sets DeveloperId field to given value.


### GetDescription

`func (o *BotDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BotDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BotDetail) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBotUserId

`func (o *BotDetail) GetBotUserId() string`

GetBotUserId returns the BotUserId field if non-nil, zero value otherwise.

### GetBotUserIdOk

`func (o *BotDetail) GetBotUserIdOk() (*string, bool)`

GetBotUserIdOk returns a tuple with the BotUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUserId

`func (o *BotDetail) SetBotUserId(v string)`

SetBotUserId sets BotUserId field to given value.


### GetTokens

`func (o *BotDetail) GetTokens() BotTokens`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *BotDetail) GetTokensOk() (*BotTokens, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *BotDetail) SetTokens(v BotTokens)`

SetTokens sets Tokens field to given value.


### GetEndpoint

`func (o *BotDetail) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *BotDetail) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *BotDetail) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetPrivileged

`func (o *BotDetail) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *BotDetail) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *BotDetail) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.


### GetChannels

`func (o *BotDetail) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *BotDetail) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *BotDetail) SetChannels(v []string)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


