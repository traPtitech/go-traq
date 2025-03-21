# QallRoomStateChangedEventStateParticipantsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to **string** | ユーザーID_RandomUUID | [optional] 
**Name** | Pointer to **string** | 表示名 | [optional] 
**JoinedAt** | Pointer to **time.Time** | 参加した時刻 | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**CanPublish** | Pointer to **bool** | 発言権限 | [optional] 

## Methods

### NewQallRoomStateChangedEventStateParticipantsInner

`func NewQallRoomStateChangedEventStateParticipantsInner() *QallRoomStateChangedEventStateParticipantsInner`

NewQallRoomStateChangedEventStateParticipantsInner instantiates a new QallRoomStateChangedEventStateParticipantsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQallRoomStateChangedEventStateParticipantsInnerWithDefaults

`func NewQallRoomStateChangedEventStateParticipantsInnerWithDefaults() *QallRoomStateChangedEventStateParticipantsInner`

NewQallRoomStateChangedEventStateParticipantsInnerWithDefaults instantiates a new QallRoomStateChangedEventStateParticipantsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *QallRoomStateChangedEventStateParticipantsInner) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *QallRoomStateChangedEventStateParticipantsInner) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *QallRoomStateChangedEventStateParticipantsInner) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *QallRoomStateChangedEventStateParticipantsInner) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *QallRoomStateChangedEventStateParticipantsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QallRoomStateChangedEventStateParticipantsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QallRoomStateChangedEventStateParticipantsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QallRoomStateChangedEventStateParticipantsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetJoinedAt

`func (o *QallRoomStateChangedEventStateParticipantsInner) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *QallRoomStateChangedEventStateParticipantsInner) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *QallRoomStateChangedEventStateParticipantsInner) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *QallRoomStateChangedEventStateParticipantsInner) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetAttributes

`func (o *QallRoomStateChangedEventStateParticipantsInner) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *QallRoomStateChangedEventStateParticipantsInner) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *QallRoomStateChangedEventStateParticipantsInner) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *QallRoomStateChangedEventStateParticipantsInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCanPublish

`func (o *QallRoomStateChangedEventStateParticipantsInner) GetCanPublish() bool`

GetCanPublish returns the CanPublish field if non-nil, zero value otherwise.

### GetCanPublishOk

`func (o *QallRoomStateChangedEventStateParticipantsInner) GetCanPublishOk() (*bool, bool)`

GetCanPublishOk returns a tuple with the CanPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPublish

`func (o *QallRoomStateChangedEventStateParticipantsInner) SetCanPublish(v bool)`

SetCanPublish sets CanPublish field to given value.

### HasCanPublish

`func (o *QallRoomStateChangedEventStateParticipantsInner) HasCanPublish() bool`

HasCanPublish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


