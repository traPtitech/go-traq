# ChannelStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalMessageCount** | **int64** | チャンネルの総投稿メッセージ数(削除されたものも含む) | 
**Stamps** | [**[]ChannelStatsStamp**](ChannelStatsStamp.md) | チャンネル上のスタンプ統計情報 | 
**Users** | [**[]ChannelStatsUser**](ChannelStatsUser.md) | チャンネル上のユーザー統計情報 | 
**Datetime** | **time.Time** | 統計情報日時 | 

## Methods

### NewChannelStats

`func NewChannelStats(totalMessageCount int64, stamps []ChannelStatsStamp, users []ChannelStatsUser, datetime time.Time, ) *ChannelStats`

NewChannelStats instantiates a new ChannelStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelStatsWithDefaults

`func NewChannelStatsWithDefaults() *ChannelStats`

NewChannelStatsWithDefaults instantiates a new ChannelStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalMessageCount

`func (o *ChannelStats) GetTotalMessageCount() int64`

GetTotalMessageCount returns the TotalMessageCount field if non-nil, zero value otherwise.

### GetTotalMessageCountOk

`func (o *ChannelStats) GetTotalMessageCountOk() (*int64, bool)`

GetTotalMessageCountOk returns a tuple with the TotalMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessageCount

`func (o *ChannelStats) SetTotalMessageCount(v int64)`

SetTotalMessageCount sets TotalMessageCount field to given value.


### GetStamps

`func (o *ChannelStats) GetStamps() []ChannelStatsStamp`

GetStamps returns the Stamps field if non-nil, zero value otherwise.

### GetStampsOk

`func (o *ChannelStats) GetStampsOk() (*[]ChannelStatsStamp, bool)`

GetStampsOk returns a tuple with the Stamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStamps

`func (o *ChannelStats) SetStamps(v []ChannelStatsStamp)`

SetStamps sets Stamps field to given value.


### GetUsers

`func (o *ChannelStats) GetUsers() []ChannelStatsUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ChannelStats) GetUsersOk() (*[]ChannelStatsUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ChannelStats) SetUsers(v []ChannelStatsUser)`

SetUsers sets Users field to given value.


### GetDatetime

`func (o *ChannelStats) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *ChannelStats) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *ChannelStats) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


