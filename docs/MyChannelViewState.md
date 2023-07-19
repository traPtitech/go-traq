# MyChannelViewState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | WSセッションの識別子 | 
**ChannelId** | **string** | チャンネルUUID | 
**State** | [**ChannelViewState**](ChannelViewState.md) |  | 

## Methods

### NewMyChannelViewState

`func NewMyChannelViewState(key string, channelId string, state ChannelViewState, ) *MyChannelViewState`

NewMyChannelViewState instantiates a new MyChannelViewState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyChannelViewStateWithDefaults

`func NewMyChannelViewStateWithDefaults() *MyChannelViewState`

NewMyChannelViewStateWithDefaults instantiates a new MyChannelViewState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MyChannelViewState) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MyChannelViewState) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MyChannelViewState) SetKey(v string)`

SetKey sets Key field to given value.


### GetChannelId

`func (o *MyChannelViewState) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MyChannelViewState) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MyChannelViewState) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetState

`func (o *MyChannelViewState) GetState() ChannelViewState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MyChannelViewState) GetStateOk() (*ChannelViewState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MyChannelViewState) SetState(v ChannelViewState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


