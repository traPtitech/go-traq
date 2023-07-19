# UnreadChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** | チャンネルUUID | 
**Count** | **int32** | 未読メッセージ数 | 
**Noticeable** | **bool** | 自分宛てメッセージが含まれているかどうか | 
**Since** | **time.Time** | チャンネルの最古の未読メッセージの日時 | 
**UpdatedAt** | **time.Time** | チャンネルの最新の未読メッセージの日時 | 
**OldestMessageId** | **string** | そのチャンネルの未読の中で最も古いメッセージのid | 

## Methods

### NewUnreadChannel

`func NewUnreadChannel(channelId string, count int32, noticeable bool, since time.Time, updatedAt time.Time, oldestMessageId string, ) *UnreadChannel`

NewUnreadChannel instantiates a new UnreadChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnreadChannelWithDefaults

`func NewUnreadChannelWithDefaults() *UnreadChannel`

NewUnreadChannelWithDefaults instantiates a new UnreadChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *UnreadChannel) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UnreadChannel) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UnreadChannel) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetCount

`func (o *UnreadChannel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnreadChannel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnreadChannel) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNoticeable

`func (o *UnreadChannel) GetNoticeable() bool`

GetNoticeable returns the Noticeable field if non-nil, zero value otherwise.

### GetNoticeableOk

`func (o *UnreadChannel) GetNoticeableOk() (*bool, bool)`

GetNoticeableOk returns a tuple with the Noticeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeable

`func (o *UnreadChannel) SetNoticeable(v bool)`

SetNoticeable sets Noticeable field to given value.


### GetSince

`func (o *UnreadChannel) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *UnreadChannel) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *UnreadChannel) SetSince(v time.Time)`

SetSince sets Since field to given value.


### GetUpdatedAt

`func (o *UnreadChannel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UnreadChannel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UnreadChannel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOldestMessageId

`func (o *UnreadChannel) GetOldestMessageId() string`

GetOldestMessageId returns the OldestMessageId field if non-nil, zero value otherwise.

### GetOldestMessageIdOk

`func (o *UnreadChannel) GetOldestMessageIdOk() (*string, bool)`

GetOldestMessageIdOk returns a tuple with the OldestMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestMessageId

`func (o *UnreadChannel) SetOldestMessageId(v string)`

SetOldestMessageId sets OldestMessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


