# ChildCreatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | 作成者UUID | 
**ChannelId** | **string** | チャンネルUUID | 

## Methods

### NewChildCreatedEvent

`func NewChildCreatedEvent(userId string, channelId string, ) *ChildCreatedEvent`

NewChildCreatedEvent instantiates a new ChildCreatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildCreatedEventWithDefaults

`func NewChildCreatedEventWithDefaults() *ChildCreatedEvent`

NewChildCreatedEventWithDefaults instantiates a new ChildCreatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ChildCreatedEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ChildCreatedEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ChildCreatedEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetChannelId

`func (o *ChildCreatedEvent) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ChildCreatedEvent) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ChildCreatedEvent) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


