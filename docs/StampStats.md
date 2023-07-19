# StampStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | スタンプ使用総数(同じユーザによって同じメッセージに貼られたものは複数カウントしない) | 
**TotalCount** | **int64** | スタンプ使用総数(全てカウント) | 

## Methods

### NewStampStats

`func NewStampStats(count int64, totalCount int64, ) *StampStats`

NewStampStats instantiates a new StampStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStampStatsWithDefaults

`func NewStampStatsWithDefaults() *StampStats`

NewStampStatsWithDefaults instantiates a new StampStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *StampStats) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StampStats) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StampStats) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTotalCount

`func (o *StampStats) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *StampStats) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *StampStats) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


