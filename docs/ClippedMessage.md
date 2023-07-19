# ClippedMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | [**Message**](Message.md) |  | 
**ClippedAt** | **time.Time** | クリップした日時 | 

## Methods

### NewClippedMessage

`func NewClippedMessage(message Message, clippedAt time.Time, ) *ClippedMessage`

NewClippedMessage instantiates a new ClippedMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClippedMessageWithDefaults

`func NewClippedMessageWithDefaults() *ClippedMessage`

NewClippedMessageWithDefaults instantiates a new ClippedMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ClippedMessage) GetMessage() Message`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClippedMessage) GetMessageOk() (*Message, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClippedMessage) SetMessage(v Message)`

SetMessage sets Message field to given value.


### GetClippedAt

`func (o *ClippedMessage) GetClippedAt() time.Time`

GetClippedAt returns the ClippedAt field if non-nil, zero value otherwise.

### GetClippedAtOk

`func (o *ClippedMessage) GetClippedAtOk() (*time.Time, bool)`

GetClippedAtOk returns a tuple with the ClippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClippedAt

`func (o *ClippedMessage) SetClippedAt(v time.Time)`

SetClippedAt sets ClippedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


