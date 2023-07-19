# PinRemovedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | 変更者UUID | 
**MessageId** | **string** | メッセージUUID | 

## Methods

### NewPinRemovedEvent

`func NewPinRemovedEvent(userId string, messageId string, ) *PinRemovedEvent`

NewPinRemovedEvent instantiates a new PinRemovedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinRemovedEventWithDefaults

`func NewPinRemovedEventWithDefaults() *PinRemovedEvent`

NewPinRemovedEventWithDefaults instantiates a new PinRemovedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PinRemovedEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PinRemovedEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PinRemovedEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetMessageId

`func (o *PinRemovedEvent) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *PinRemovedEvent) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *PinRemovedEvent) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


