# UserStatsStamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | スタンプID | 
**Count** | **int64** | スタンプ数(同一メッセージ上のものは複数カウントしない) | 
**Total** | **int64** | スタンプ数(同一メッセージ上のものも複数カウントする) | 

## Methods

### NewUserStatsStamp

`func NewUserStatsStamp(id string, count int64, total int64, ) *UserStatsStamp`

NewUserStatsStamp instantiates a new UserStatsStamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatsStampWithDefaults

`func NewUserStatsStampWithDefaults() *UserStatsStamp`

NewUserStatsStampWithDefaults instantiates a new UserStatsStamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserStatsStamp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserStatsStamp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserStatsStamp) SetId(v string)`

SetId sets Id field to given value.


### GetCount

`func (o *UserStatsStamp) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UserStatsStamp) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UserStatsStamp) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTotal

`func (o *UserStatsStamp) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UserStatsStamp) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UserStatsStamp) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


