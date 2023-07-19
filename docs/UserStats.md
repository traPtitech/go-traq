# UserStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalMessageCount** | **int64** | ユーザーの総投稿メッセージ数(削除されたものも含む) | 
**Stamps** | [**[]UserStatsStamp**](UserStatsStamp.md) | ユーザーのスタンプ統計情報 | 
**Datetime** | **time.Time** | 統計情報日時 | 

## Methods

### NewUserStats

`func NewUserStats(totalMessageCount int64, stamps []UserStatsStamp, datetime time.Time, ) *UserStats`

NewUserStats instantiates a new UserStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatsWithDefaults

`func NewUserStatsWithDefaults() *UserStats`

NewUserStatsWithDefaults instantiates a new UserStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalMessageCount

`func (o *UserStats) GetTotalMessageCount() int64`

GetTotalMessageCount returns the TotalMessageCount field if non-nil, zero value otherwise.

### GetTotalMessageCountOk

`func (o *UserStats) GetTotalMessageCountOk() (*int64, bool)`

GetTotalMessageCountOk returns a tuple with the TotalMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessageCount

`func (o *UserStats) SetTotalMessageCount(v int64)`

SetTotalMessageCount sets TotalMessageCount field to given value.


### GetStamps

`func (o *UserStats) GetStamps() []UserStatsStamp`

GetStamps returns the Stamps field if non-nil, zero value otherwise.

### GetStampsOk

`func (o *UserStats) GetStampsOk() (*[]UserStatsStamp, bool)`

GetStampsOk returns a tuple with the Stamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStamps

`func (o *UserStats) SetStamps(v []UserStatsStamp)`

SetStamps sets Stamps field to given value.


### GetDatetime

`func (o *UserStats) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *UserStats) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *UserStats) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


