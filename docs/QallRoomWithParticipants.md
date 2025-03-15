# QallRoomWithParticipants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomId** | **string** | ルームのID | 
**Participants** | [**[]QallParticipant**](QallParticipant.md) |  | 
**IsWebinar** | Pointer to **bool** | ウェビナールームかどうか | [optional] 
**Metadata** | Pointer to **string** | ルームに関連付けられたカスタム属性 | [optional] 

## Methods

### NewQallRoomWithParticipants

`func NewQallRoomWithParticipants(roomId string, participants []QallParticipant, ) *QallRoomWithParticipants`

NewQallRoomWithParticipants instantiates a new QallRoomWithParticipants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQallRoomWithParticipantsWithDefaults

`func NewQallRoomWithParticipantsWithDefaults() *QallRoomWithParticipants`

NewQallRoomWithParticipantsWithDefaults instantiates a new QallRoomWithParticipants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomId

`func (o *QallRoomWithParticipants) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *QallRoomWithParticipants) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *QallRoomWithParticipants) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.


### GetParticipants

`func (o *QallRoomWithParticipants) GetParticipants() []QallParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *QallRoomWithParticipants) GetParticipantsOk() (*[]QallParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *QallRoomWithParticipants) SetParticipants(v []QallParticipant)`

SetParticipants sets Participants field to given value.


### GetIsWebinar

`func (o *QallRoomWithParticipants) GetIsWebinar() bool`

GetIsWebinar returns the IsWebinar field if non-nil, zero value otherwise.

### GetIsWebinarOk

`func (o *QallRoomWithParticipants) GetIsWebinarOk() (*bool, bool)`

GetIsWebinarOk returns a tuple with the IsWebinar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebinar

`func (o *QallRoomWithParticipants) SetIsWebinar(v bool)`

SetIsWebinar sets IsWebinar field to given value.

### HasIsWebinar

`func (o *QallRoomWithParticipants) HasIsWebinar() bool`

HasIsWebinar returns a boolean if a field has been set.

### GetMetadata

`func (o *QallRoomWithParticipants) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QallRoomWithParticipants) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QallRoomWithParticipants) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QallRoomWithParticipants) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


