# UserSubscribeState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** | チャンネルUUID | 
**Level** | [**ChannelSubscribeLevel**](ChannelSubscribeLevel.md) |  | 

## Methods

### NewUserSubscribeState

`func NewUserSubscribeState(channelId string, level ChannelSubscribeLevel, ) *UserSubscribeState`

NewUserSubscribeState instantiates a new UserSubscribeState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSubscribeStateWithDefaults

`func NewUserSubscribeStateWithDefaults() *UserSubscribeState`

NewUserSubscribeStateWithDefaults instantiates a new UserSubscribeState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *UserSubscribeState) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UserSubscribeState) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UserSubscribeState) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetLevel

`func (o *UserSubscribeState) GetLevel() ChannelSubscribeLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *UserSubscribeState) GetLevelOk() (*ChannelSubscribeLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *UserSubscribeState) SetLevel(v ChannelSubscribeLevel)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


