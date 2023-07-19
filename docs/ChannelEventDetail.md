# ChannelEventDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | 作成者UUID | 
**Before** | **string** | 変更前親チャンネルUUID | 
**After** | **string** | 変更後親チャンネルUUID | 
**On** | **[]string** | オンにされたユーザーのUUID配列 | 
**Off** | **[]string** | オフにされたユーザーのUUID配列 | 
**MessageId** | **string** | メッセージUUID | 
**Visibility** | **bool** | 変更後可視状態 | 
**Force** | **bool** | 変更後強制通知状態 | 
**ChannelId** | **string** | チャンネルUUID | 

## Methods

### NewChannelEventDetail

`func NewChannelEventDetail(userId string, before string, after string, on []string, off []string, messageId string, visibility bool, force bool, channelId string, ) *ChannelEventDetail`

NewChannelEventDetail instantiates a new ChannelEventDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelEventDetailWithDefaults

`func NewChannelEventDetailWithDefaults() *ChannelEventDetail`

NewChannelEventDetailWithDefaults instantiates a new ChannelEventDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ChannelEventDetail) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ChannelEventDetail) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ChannelEventDetail) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetBefore

`func (o *ChannelEventDetail) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ChannelEventDetail) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ChannelEventDetail) SetBefore(v string)`

SetBefore sets Before field to given value.


### GetAfter

`func (o *ChannelEventDetail) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ChannelEventDetail) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ChannelEventDetail) SetAfter(v string)`

SetAfter sets After field to given value.


### GetOn

`func (o *ChannelEventDetail) GetOn() []string`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *ChannelEventDetail) GetOnOk() (*[]string, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *ChannelEventDetail) SetOn(v []string)`

SetOn sets On field to given value.


### GetOff

`func (o *ChannelEventDetail) GetOff() []string`

GetOff returns the Off field if non-nil, zero value otherwise.

### GetOffOk

`func (o *ChannelEventDetail) GetOffOk() (*[]string, bool)`

GetOffOk returns a tuple with the Off field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOff

`func (o *ChannelEventDetail) SetOff(v []string)`

SetOff sets Off field to given value.


### GetMessageId

`func (o *ChannelEventDetail) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ChannelEventDetail) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ChannelEventDetail) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetVisibility

`func (o *ChannelEventDetail) GetVisibility() bool`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ChannelEventDetail) GetVisibilityOk() (*bool, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ChannelEventDetail) SetVisibility(v bool)`

SetVisibility sets Visibility field to given value.


### GetForce

`func (o *ChannelEventDetail) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *ChannelEventDetail) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *ChannelEventDetail) SetForce(v bool)`

SetForce sets Force field to given value.


### GetChannelId

`func (o *ChannelEventDetail) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ChannelEventDetail) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ChannelEventDetail) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


