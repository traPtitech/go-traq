# ChannelStatsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ユーザーID | 
**MessageCount** | **int64** | メッセージ数 | 

## Methods

### NewChannelStatsUser

`func NewChannelStatsUser(id string, messageCount int64, ) *ChannelStatsUser`

NewChannelStatsUser instantiates a new ChannelStatsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelStatsUserWithDefaults

`func NewChannelStatsUserWithDefaults() *ChannelStatsUser`

NewChannelStatsUserWithDefaults instantiates a new ChannelStatsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelStatsUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelStatsUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelStatsUser) SetId(v string)`

SetId sets Id field to given value.


### GetMessageCount

`func (o *ChannelStatsUser) GetMessageCount() int64`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *ChannelStatsUser) GetMessageCountOk() (*int64, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *ChannelStatsUser) SetMessageCount(v int64)`

SetMessageCount sets MessageCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


