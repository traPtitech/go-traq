# QallParticipantRequestUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | ユーザーID | [optional] 
**CanPublish** | Pointer to **bool** | 発言権限 | [optional] 

## Methods

### NewQallParticipantRequestUsersInner

`func NewQallParticipantRequestUsersInner() *QallParticipantRequestUsersInner`

NewQallParticipantRequestUsersInner instantiates a new QallParticipantRequestUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQallParticipantRequestUsersInnerWithDefaults

`func NewQallParticipantRequestUsersInnerWithDefaults() *QallParticipantRequestUsersInner`

NewQallParticipantRequestUsersInnerWithDefaults instantiates a new QallParticipantRequestUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *QallParticipantRequestUsersInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *QallParticipantRequestUsersInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *QallParticipantRequestUsersInner) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *QallParticipantRequestUsersInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCanPublish

`func (o *QallParticipantRequestUsersInner) GetCanPublish() bool`

GetCanPublish returns the CanPublish field if non-nil, zero value otherwise.

### GetCanPublishOk

`func (o *QallParticipantRequestUsersInner) GetCanPublishOk() (*bool, bool)`

GetCanPublishOk returns a tuple with the CanPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPublish

`func (o *QallParticipantRequestUsersInner) SetCanPublish(v bool)`

SetCanPublish sets CanPublish field to given value.

### HasCanPublish

`func (o *QallParticipantRequestUsersInner) HasCanPublish() bool`

HasCanPublish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


