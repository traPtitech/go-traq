# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | メッセージUUID | 
**UserId** | **string** | 投稿者UUID | 
**ChannelId** | **string** | チャンネルUUID | 
**Content** | **string** | メッセージ本文 | 
**CreatedAt** | **time.Time** | 投稿日時 | 
**UpdatedAt** | **time.Time** | 編集日時 | 
**Pinned** | **bool** | ピン留めされているかどうか | 
**Stamps** | [**[]MessageStamp**](MessageStamp.md) | 押されているスタンプの配列 | 
**ThreadId** | **NullableString** | スレッドUUID | 

## Methods

### NewMessage

`func NewMessage(id string, userId string, channelId string, content string, createdAt time.Time, updatedAt time.Time, pinned bool, stamps []MessageStamp, threadId NullableString, ) *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Message) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Message) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Message) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *Message) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Message) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Message) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetChannelId

`func (o *Message) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *Message) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *Message) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetContent

`func (o *Message) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Message) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Message) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreatedAt

`func (o *Message) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Message) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Message) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Message) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Message) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Message) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPinned

`func (o *Message) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *Message) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *Message) SetPinned(v bool)`

SetPinned sets Pinned field to given value.


### GetStamps

`func (o *Message) GetStamps() []MessageStamp`

GetStamps returns the Stamps field if non-nil, zero value otherwise.

### GetStampsOk

`func (o *Message) GetStampsOk() (*[]MessageStamp, bool)`

GetStampsOk returns a tuple with the Stamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStamps

`func (o *Message) SetStamps(v []MessageStamp)`

SetStamps sets Stamps field to given value.


### GetThreadId

`func (o *Message) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *Message) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *Message) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### SetThreadIdNil

`func (o *Message) SetThreadIdNil(b bool)`

 SetThreadIdNil sets the value for ThreadId to be an explicit nil

### UnsetThreadId
`func (o *Message) UnsetThreadId()`

UnsetThreadId ensures that no value is present for ThreadId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


