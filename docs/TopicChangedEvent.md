# TopicChangedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | 変更者UUID | 
**Before** | **string** | 変更前トピック | 
**After** | **string** | 変更後トピック | 

## Methods

### NewTopicChangedEvent

`func NewTopicChangedEvent(userId string, before string, after string, ) *TopicChangedEvent`

NewTopicChangedEvent instantiates a new TopicChangedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicChangedEventWithDefaults

`func NewTopicChangedEventWithDefaults() *TopicChangedEvent`

NewTopicChangedEventWithDefaults instantiates a new TopicChangedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *TopicChangedEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TopicChangedEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TopicChangedEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetBefore

`func (o *TopicChangedEvent) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *TopicChangedEvent) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *TopicChangedEvent) SetBefore(v string)`

SetBefore sets Before field to given value.


### GetAfter

`func (o *TopicChangedEvent) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *TopicChangedEvent) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *TopicChangedEvent) SetAfter(v string)`

SetAfter sets After field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


