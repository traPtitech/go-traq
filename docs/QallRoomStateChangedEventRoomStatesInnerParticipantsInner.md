# QallRoomStateChangedEventRoomStatesInnerParticipantsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | **string** | ユーザーID_RandomUUID | 
**Name** | **string** | 表示名 | 
**JoinedAt** | **time.Time** | 参加した時刻 | 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**CanPublish** | **bool** | 発言権限 | 

## Methods

### NewQallRoomStateChangedEventRoomStatesInnerParticipantsInner

`func NewQallRoomStateChangedEventRoomStatesInnerParticipantsInner(identity string, name string, joinedAt time.Time, canPublish bool, ) *QallRoomStateChangedEventRoomStatesInnerParticipantsInner`

NewQallRoomStateChangedEventRoomStatesInnerParticipantsInner instantiates a new QallRoomStateChangedEventRoomStatesInnerParticipantsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQallRoomStateChangedEventRoomStatesInnerParticipantsInnerWithDefaults

`func NewQallRoomStateChangedEventRoomStatesInnerParticipantsInnerWithDefaults() *QallRoomStateChangedEventRoomStatesInnerParticipantsInner`

NewQallRoomStateChangedEventRoomStatesInnerParticipantsInnerWithDefaults instantiates a new QallRoomStateChangedEventRoomStatesInnerParticipantsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) SetIdentity(v string)`

SetIdentity sets Identity field to given value.


### GetName

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) SetName(v string)`

SetName sets Name field to given value.


### GetJoinedAt

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.


### GetAttributes

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCanPublish

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) GetCanPublish() bool`

GetCanPublish returns the CanPublish field if non-nil, zero value otherwise.

### GetCanPublishOk

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) GetCanPublishOk() (*bool, bool)`

GetCanPublishOk returns a tuple with the CanPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPublish

`func (o *QallRoomStateChangedEventRoomStatesInnerParticipantsInner) SetCanPublish(v bool)`

SetCanPublish sets CanPublish field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


