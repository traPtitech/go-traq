# QallRoomStateChangedEventRoomStatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomId** | **string** | ルームのID | 
**Participants** | [**[]QallRoomStateChangedEventRoomStatesInnerParticipantsInner**](QallRoomStateChangedEventRoomStatesInnerParticipantsInner.md) |  | 
**IsWebinar** | **bool** | ウェビナールームかどうか | 
**Metadata** | Pointer to **string** | ルームに関連付けられたカスタム属性 | [optional] 

## Methods

### NewQallRoomStateChangedEventRoomStatesInner

`func NewQallRoomStateChangedEventRoomStatesInner(roomId string, participants []QallRoomStateChangedEventRoomStatesInnerParticipantsInner, isWebinar bool, ) *QallRoomStateChangedEventRoomStatesInner`

NewQallRoomStateChangedEventRoomStatesInner instantiates a new QallRoomStateChangedEventRoomStatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQallRoomStateChangedEventRoomStatesInnerWithDefaults

`func NewQallRoomStateChangedEventRoomStatesInnerWithDefaults() *QallRoomStateChangedEventRoomStatesInner`

NewQallRoomStateChangedEventRoomStatesInnerWithDefaults instantiates a new QallRoomStateChangedEventRoomStatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomId

`func (o *QallRoomStateChangedEventRoomStatesInner) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *QallRoomStateChangedEventRoomStatesInner) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *QallRoomStateChangedEventRoomStatesInner) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.


### GetParticipants

`func (o *QallRoomStateChangedEventRoomStatesInner) GetParticipants() []QallRoomStateChangedEventRoomStatesInnerParticipantsInner`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *QallRoomStateChangedEventRoomStatesInner) GetParticipantsOk() (*[]QallRoomStateChangedEventRoomStatesInnerParticipantsInner, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *QallRoomStateChangedEventRoomStatesInner) SetParticipants(v []QallRoomStateChangedEventRoomStatesInnerParticipantsInner)`

SetParticipants sets Participants field to given value.


### GetIsWebinar

`func (o *QallRoomStateChangedEventRoomStatesInner) GetIsWebinar() bool`

GetIsWebinar returns the IsWebinar field if non-nil, zero value otherwise.

### GetIsWebinarOk

`func (o *QallRoomStateChangedEventRoomStatesInner) GetIsWebinarOk() (*bool, bool)`

GetIsWebinarOk returns a tuple with the IsWebinar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebinar

`func (o *QallRoomStateChangedEventRoomStatesInner) SetIsWebinar(v bool)`

SetIsWebinar sets IsWebinar field to given value.


### GetMetadata

`func (o *QallRoomStateChangedEventRoomStatesInner) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QallRoomStateChangedEventRoomStatesInner) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QallRoomStateChangedEventRoomStatesInner) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QallRoomStateChangedEventRoomStatesInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


