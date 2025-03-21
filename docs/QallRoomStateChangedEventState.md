# QallRoomStateChangedEventState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomId** | Pointer to **string** | ルームのID | [optional] 
**Participants** | Pointer to [**[]QallRoomStateChangedEventStateParticipantsInner**](QallRoomStateChangedEventStateParticipantsInner.md) |  | [optional] 
**IsWebinar** | Pointer to **bool** | ウェビナールームかどうか | [optional] 
**Metadata** | Pointer to **string** | ルームに関連付けられたカスタム属性 | [optional] 

## Methods

### NewQallRoomStateChangedEventState

`func NewQallRoomStateChangedEventState() *QallRoomStateChangedEventState`

NewQallRoomStateChangedEventState instantiates a new QallRoomStateChangedEventState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQallRoomStateChangedEventStateWithDefaults

`func NewQallRoomStateChangedEventStateWithDefaults() *QallRoomStateChangedEventState`

NewQallRoomStateChangedEventStateWithDefaults instantiates a new QallRoomStateChangedEventState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomId

`func (o *QallRoomStateChangedEventState) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *QallRoomStateChangedEventState) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *QallRoomStateChangedEventState) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *QallRoomStateChangedEventState) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetParticipants

`func (o *QallRoomStateChangedEventState) GetParticipants() []QallRoomStateChangedEventStateParticipantsInner`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *QallRoomStateChangedEventState) GetParticipantsOk() (*[]QallRoomStateChangedEventStateParticipantsInner, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *QallRoomStateChangedEventState) SetParticipants(v []QallRoomStateChangedEventStateParticipantsInner)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *QallRoomStateChangedEventState) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetIsWebinar

`func (o *QallRoomStateChangedEventState) GetIsWebinar() bool`

GetIsWebinar returns the IsWebinar field if non-nil, zero value otherwise.

### GetIsWebinarOk

`func (o *QallRoomStateChangedEventState) GetIsWebinarOk() (*bool, bool)`

GetIsWebinarOk returns a tuple with the IsWebinar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebinar

`func (o *QallRoomStateChangedEventState) SetIsWebinar(v bool)`

SetIsWebinar sets IsWebinar field to given value.

### HasIsWebinar

`func (o *QallRoomStateChangedEventState) HasIsWebinar() bool`

HasIsWebinar returns a boolean if a field has been set.

### GetMetadata

`func (o *QallRoomStateChangedEventState) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QallRoomStateChangedEventState) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QallRoomStateChangedEventState) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QallRoomStateChangedEventState) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


