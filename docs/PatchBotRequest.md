# PatchBotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | BOTユーザー表示名 | [optional] 
**Description** | Pointer to **string** | BOTの説明 | [optional] 
**Privileged** | Pointer to **bool** | 特権 | [optional] 
**Mode** | Pointer to [**BotMode**](BotMode.md) |  | [optional] 
**Endpoint** | Pointer to **string** | BOTサーバーエンドポイント | [optional] 
**DeveloperId** | Pointer to **string** | 移譲先の開発者UUID | [optional] 
**SubscribeEvents** | Pointer to **[]string** | 購読するイベント | [optional] 
**Bio** | Pointer to **string** | 自己紹介(biography) | [optional] 

## Methods

### NewPatchBotRequest

`func NewPatchBotRequest() *PatchBotRequest`

NewPatchBotRequest instantiates a new PatchBotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchBotRequestWithDefaults

`func NewPatchBotRequestWithDefaults() *PatchBotRequest`

NewPatchBotRequestWithDefaults instantiates a new PatchBotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PatchBotRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchBotRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchBotRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PatchBotRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchBotRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchBotRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchBotRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchBotRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrivileged

`func (o *PatchBotRequest) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *PatchBotRequest) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *PatchBotRequest) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *PatchBotRequest) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetMode

`func (o *PatchBotRequest) GetMode() BotMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchBotRequest) GetModeOk() (*BotMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchBotRequest) SetMode(v BotMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchBotRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetEndpoint

`func (o *PatchBotRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PatchBotRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PatchBotRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PatchBotRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetDeveloperId

`func (o *PatchBotRequest) GetDeveloperId() string`

GetDeveloperId returns the DeveloperId field if non-nil, zero value otherwise.

### GetDeveloperIdOk

`func (o *PatchBotRequest) GetDeveloperIdOk() (*string, bool)`

GetDeveloperIdOk returns a tuple with the DeveloperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperId

`func (o *PatchBotRequest) SetDeveloperId(v string)`

SetDeveloperId sets DeveloperId field to given value.

### HasDeveloperId

`func (o *PatchBotRequest) HasDeveloperId() bool`

HasDeveloperId returns a boolean if a field has been set.

### GetSubscribeEvents

`func (o *PatchBotRequest) GetSubscribeEvents() []string`

GetSubscribeEvents returns the SubscribeEvents field if non-nil, zero value otherwise.

### GetSubscribeEventsOk

`func (o *PatchBotRequest) GetSubscribeEventsOk() (*[]string, bool)`

GetSubscribeEventsOk returns a tuple with the SubscribeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeEvents

`func (o *PatchBotRequest) SetSubscribeEvents(v []string)`

SetSubscribeEvents sets SubscribeEvents field to given value.

### HasSubscribeEvents

`func (o *PatchBotRequest) HasSubscribeEvents() bool`

HasSubscribeEvents returns a boolean if a field has been set.

### GetBio

`func (o *PatchBotRequest) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *PatchBotRequest) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *PatchBotRequest) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *PatchBotRequest) HasBio() bool`

HasBio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


