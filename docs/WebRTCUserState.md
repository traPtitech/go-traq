# WebRTCUserState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | ユーザーUUID | 
**ChannelId** | **string** | チャンネルUUID | 
**Sessions** | [**[]Session**](Session.md) | セッションの配列 | 

## Methods

### NewWebRTCUserState

`func NewWebRTCUserState(userId string, channelId string, sessions []Session, ) *WebRTCUserState`

NewWebRTCUserState instantiates a new WebRTCUserState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRTCUserStateWithDefaults

`func NewWebRTCUserStateWithDefaults() *WebRTCUserState`

NewWebRTCUserStateWithDefaults instantiates a new WebRTCUserState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *WebRTCUserState) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebRTCUserState) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebRTCUserState) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetChannelId

`func (o *WebRTCUserState) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *WebRTCUserState) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *WebRTCUserState) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetSessions

`func (o *WebRTCUserState) GetSessions() []Session`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *WebRTCUserState) GetSessionsOk() (*[]Session, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *WebRTCUserState) SetSessions(v []Session)`

SetSessions sets Sessions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


