# QallRoomStateChangedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomId** | **string** | 変更されたルームのId | 
**State** | [**QallRoomStateChangedEventState**](QallRoomStateChangedEventState.md) |  | 

## Methods

### NewQallRoomStateChangedEvent

`func NewQallRoomStateChangedEvent(roomId string, state QallRoomStateChangedEventState, ) *QallRoomStateChangedEvent`

NewQallRoomStateChangedEvent instantiates a new QallRoomStateChangedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQallRoomStateChangedEventWithDefaults

`func NewQallRoomStateChangedEventWithDefaults() *QallRoomStateChangedEvent`

NewQallRoomStateChangedEventWithDefaults instantiates a new QallRoomStateChangedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomId

`func (o *QallRoomStateChangedEvent) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *QallRoomStateChangedEvent) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *QallRoomStateChangedEvent) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.


### GetState

`func (o *QallRoomStateChangedEvent) GetState() QallRoomStateChangedEventState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QallRoomStateChangedEvent) GetStateOk() (*QallRoomStateChangedEventState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QallRoomStateChangedEvent) SetState(v QallRoomStateChangedEventState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


