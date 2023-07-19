# Bot

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

## Methods

### NewBot

`func NewBot(id string, botUserId string, description string, developerId string, subscribeEvents []string, mode BotMode, state BotState, createdAt time.Time, updatedAt time.Time, ) *Bot`

NewBot instantiates a new Bot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotWithDefaults

`func NewBotWithDefaults() *Bot`

NewBotWithDefaults instantiates a new Bot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bot) SetId(v string)`

SetId sets Id field to given value.


### GetBotUserId

`func (o *Bot) GetBotUserId() string`

GetBotUserId returns the BotUserId field if non-nil, zero value otherwise.

### GetBotUserIdOk

`func (o *Bot) GetBotUserIdOk() (*string, bool)`

GetBotUserIdOk returns a tuple with the BotUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUserId

`func (o *Bot) SetBotUserId(v string)`

SetBotUserId sets BotUserId field to given value.


### GetDescription

`func (o *Bot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Bot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Bot) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDeveloperId

`func (o *Bot) GetDeveloperId() string`

GetDeveloperId returns the DeveloperId field if non-nil, zero value otherwise.

### GetDeveloperIdOk

`func (o *Bot) GetDeveloperIdOk() (*string, bool)`

GetDeveloperIdOk returns a tuple with the DeveloperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperId

`func (o *Bot) SetDeveloperId(v string)`

SetDeveloperId sets DeveloperId field to given value.


### GetSubscribeEvents

`func (o *Bot) GetSubscribeEvents() []string`

GetSubscribeEvents returns the SubscribeEvents field if non-nil, zero value otherwise.

### GetSubscribeEventsOk

`func (o *Bot) GetSubscribeEventsOk() (*[]string, bool)`

GetSubscribeEventsOk returns a tuple with the SubscribeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeEvents

`func (o *Bot) SetSubscribeEvents(v []string)`

SetSubscribeEvents sets SubscribeEvents field to given value.


### GetMode

`func (o *Bot) GetMode() BotMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Bot) GetModeOk() (*BotMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Bot) SetMode(v BotMode)`

SetMode sets Mode field to given value.


### GetState

`func (o *Bot) GetState() BotState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Bot) GetStateOk() (*BotState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Bot) SetState(v BotState)`

SetState sets State field to given value.


### GetCreatedAt

`func (o *Bot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Bot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Bot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Bot) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Bot) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Bot) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


