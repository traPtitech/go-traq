# GetBot200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | BOT UUID | 
**BotUserId** | **string** | BOTユーザーUUID | 
**Description** | **string** | 説明 | 
**DeveloperId** | **string** | BOT開発者UUID | 
**SubscribeEvents** | **[]string** | BOTが購読しているイベントの配列 | 
**Mode** | [**BotMode**](BotMode.md) |  | 
**State** | [**BotState**](BotState.md) |  | 
**CreatedAt** | **time.Time** | 作成日時 | 
**UpdatedAt** | **time.Time** | 更新日時 | 
**Tokens** | [**BotTokens**](BotTokens.md) |  | 
**Endpoint** | **string** | BOTサーバーエンドポイント | 
**Privileged** | **bool** | 特権BOTかどうか | 
**Channels** | **[]string** | BOTが参加しているチャンネルのUUID配列 | 

## Methods

### NewGetBot200Response

`func NewGetBot200Response(id string, botUserId string, description string, developerId string, subscribeEvents []string, mode BotMode, state BotState, createdAt time.Time, updatedAt time.Time, tokens BotTokens, endpoint string, privileged bool, channels []string, ) *GetBot200Response`

NewGetBot200Response instantiates a new GetBot200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBot200ResponseWithDefaults

`func NewGetBot200ResponseWithDefaults() *GetBot200Response`

NewGetBot200ResponseWithDefaults instantiates a new GetBot200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetBot200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBot200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBot200Response) SetId(v string)`

SetId sets Id field to given value.


### GetBotUserId

`func (o *GetBot200Response) GetBotUserId() string`

GetBotUserId returns the BotUserId field if non-nil, zero value otherwise.

### GetBotUserIdOk

`func (o *GetBot200Response) GetBotUserIdOk() (*string, bool)`

GetBotUserIdOk returns a tuple with the BotUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUserId

`func (o *GetBot200Response) SetBotUserId(v string)`

SetBotUserId sets BotUserId field to given value.


### GetDescription

`func (o *GetBot200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetBot200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetBot200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDeveloperId

`func (o *GetBot200Response) GetDeveloperId() string`

GetDeveloperId returns the DeveloperId field if non-nil, zero value otherwise.

### GetDeveloperIdOk

`func (o *GetBot200Response) GetDeveloperIdOk() (*string, bool)`

GetDeveloperIdOk returns a tuple with the DeveloperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperId

`func (o *GetBot200Response) SetDeveloperId(v string)`

SetDeveloperId sets DeveloperId field to given value.


### GetSubscribeEvents

`func (o *GetBot200Response) GetSubscribeEvents() []string`

GetSubscribeEvents returns the SubscribeEvents field if non-nil, zero value otherwise.

### GetSubscribeEventsOk

`func (o *GetBot200Response) GetSubscribeEventsOk() (*[]string, bool)`

GetSubscribeEventsOk returns a tuple with the SubscribeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeEvents

`func (o *GetBot200Response) SetSubscribeEvents(v []string)`

SetSubscribeEvents sets SubscribeEvents field to given value.


### GetMode

`func (o *GetBot200Response) GetMode() BotMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetBot200Response) GetModeOk() (*BotMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetBot200Response) SetMode(v BotMode)`

SetMode sets Mode field to given value.


### GetState

`func (o *GetBot200Response) GetState() BotState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetBot200Response) GetStateOk() (*BotState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetBot200Response) SetState(v BotState)`

SetState sets State field to given value.


### GetCreatedAt

`func (o *GetBot200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetBot200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetBot200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetBot200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetBot200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetBot200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTokens

`func (o *GetBot200Response) GetTokens() BotTokens`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *GetBot200Response) GetTokensOk() (*BotTokens, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *GetBot200Response) SetTokens(v BotTokens)`

SetTokens sets Tokens field to given value.


### GetEndpoint

`func (o *GetBot200Response) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GetBot200Response) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GetBot200Response) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetPrivileged

`func (o *GetBot200Response) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *GetBot200Response) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *GetBot200Response) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.


### GetChannels

`func (o *GetBot200Response) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *GetBot200Response) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *GetBot200Response) SetChannels(v []string)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


