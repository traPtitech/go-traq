# Pin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | ピン留めしたユーザーUUID | 
**PinnedAt** | **time.Time** | ピン留めされた日時 | 
**Message** | [**Message**](Message.md) |  | 

## Methods

### NewPin

`func NewPin(userId string, pinnedAt time.Time, message Message, ) *Pin`

NewPin instantiates a new Pin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinWithDefaults

`func NewPinWithDefaults() *Pin`

NewPinWithDefaults instantiates a new Pin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *Pin) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Pin) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Pin) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPinnedAt

`func (o *Pin) GetPinnedAt() time.Time`

GetPinnedAt returns the PinnedAt field if non-nil, zero value otherwise.

### GetPinnedAtOk

`func (o *Pin) GetPinnedAtOk() (*time.Time, bool)`

GetPinnedAtOk returns a tuple with the PinnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedAt

`func (o *Pin) SetPinnedAt(v time.Time)`

SetPinnedAt sets PinnedAt field to given value.


### GetMessage

`func (o *Pin) GetMessage() Message`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Pin) GetMessageOk() (*Message, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Pin) SetMessage(v Message)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


