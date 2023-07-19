# PostBotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | BOTユーザーID 自動的に接頭辞\&quot;BOT_\&quot;が付与されます | 
**DisplayName** | **string** | BOTユーザー表示名 | 
**Description** | **string** | BOTの説明 | 
**Mode** | [**BotMode**](BotMode.md) |  | 
**Endpoint** | Pointer to **string** | BOTサーバーエンドポイント BOT動作モードがHTTPの場合必須です | [optional] 

## Methods

### NewPostBotRequest

`func NewPostBotRequest(name string, displayName string, description string, mode BotMode, ) *PostBotRequest`

NewPostBotRequest instantiates a new PostBotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostBotRequestWithDefaults

`func NewPostBotRequestWithDefaults() *PostBotRequest`

NewPostBotRequestWithDefaults instantiates a new PostBotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostBotRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostBotRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostBotRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *PostBotRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PostBotRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PostBotRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *PostBotRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostBotRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostBotRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMode

`func (o *PostBotRequest) GetMode() BotMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PostBotRequest) GetModeOk() (*BotMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PostBotRequest) SetMode(v BotMode)`

SetMode sets Mode field to given value.


### GetEndpoint

`func (o *PostBotRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PostBotRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PostBotRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PostBotRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


