# ChannelEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | イベントタイプ | 
**Datetime** | **time.Time** | イベント日時 | 
**Detail** | [**ChannelEventDetail**](ChannelEventDetail.md) |  | 

## Methods

### NewChannelEvent

`func NewChannelEvent(type_ string, datetime time.Time, detail ChannelEventDetail, ) *ChannelEvent`

NewChannelEvent instantiates a new ChannelEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelEventWithDefaults

`func NewChannelEventWithDefaults() *ChannelEvent`

NewChannelEventWithDefaults instantiates a new ChannelEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChannelEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelEvent) SetType(v string)`

SetType sets Type field to given value.


### GetDatetime

`func (o *ChannelEvent) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *ChannelEvent) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *ChannelEvent) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetDetail

`func (o *ChannelEvent) GetDetail() ChannelEventDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ChannelEvent) GetDetailOk() (*ChannelEventDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ChannelEvent) SetDetail(v ChannelEventDetail)`

SetDetail sets Detail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


