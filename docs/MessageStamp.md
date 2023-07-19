# MessageStamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | ユーザーUUID | 
**StampId** | **string** | スタンプUUID | 
**Count** | **int32** | スタンプ数 | 
**CreatedAt** | **time.Time** | スタンプが最初に押された日時 | 
**UpdatedAt** | **time.Time** | スタンプが最後に押された日時 | 

## Methods

### NewMessageStamp

`func NewMessageStamp(userId string, stampId string, count int32, createdAt time.Time, updatedAt time.Time, ) *MessageStamp`

NewMessageStamp instantiates a new MessageStamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageStampWithDefaults

`func NewMessageStampWithDefaults() *MessageStamp`

NewMessageStampWithDefaults instantiates a new MessageStamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *MessageStamp) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MessageStamp) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MessageStamp) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetStampId

`func (o *MessageStamp) GetStampId() string`

GetStampId returns the StampId field if non-nil, zero value otherwise.

### GetStampIdOk

`func (o *MessageStamp) GetStampIdOk() (*string, bool)`

GetStampIdOk returns a tuple with the StampId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampId

`func (o *MessageStamp) SetStampId(v string)`

SetStampId sets StampId field to given value.


### GetCount

`func (o *MessageStamp) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MessageStamp) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MessageStamp) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCreatedAt

`func (o *MessageStamp) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageStamp) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageStamp) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MessageStamp) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MessageStamp) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MessageStamp) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


