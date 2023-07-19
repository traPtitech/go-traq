# ForcedNotificationChangedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | 変更者UUID | 
**Force** | **bool** | 変更後強制通知状態 | 

## Methods

### NewForcedNotificationChangedEvent

`func NewForcedNotificationChangedEvent(userId string, force bool, ) *ForcedNotificationChangedEvent`

NewForcedNotificationChangedEvent instantiates a new ForcedNotificationChangedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForcedNotificationChangedEventWithDefaults

`func NewForcedNotificationChangedEventWithDefaults() *ForcedNotificationChangedEvent`

NewForcedNotificationChangedEventWithDefaults instantiates a new ForcedNotificationChangedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ForcedNotificationChangedEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ForcedNotificationChangedEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ForcedNotificationChangedEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetForce

`func (o *ForcedNotificationChangedEvent) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *ForcedNotificationChangedEvent) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *ForcedNotificationChangedEvent) SetForce(v bool)`

SetForce sets Force field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


