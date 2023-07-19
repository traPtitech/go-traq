# SubscribersChangedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | 変更者UUID | 
**On** | **[]string** | オンにされたユーザーのUUID配列 | 
**Off** | **[]string** | オフにされたユーザーのUUID配列 | 

## Methods

### NewSubscribersChangedEvent

`func NewSubscribersChangedEvent(userId string, on []string, off []string, ) *SubscribersChangedEvent`

NewSubscribersChangedEvent instantiates a new SubscribersChangedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribersChangedEventWithDefaults

`func NewSubscribersChangedEventWithDefaults() *SubscribersChangedEvent`

NewSubscribersChangedEventWithDefaults instantiates a new SubscribersChangedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SubscribersChangedEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SubscribersChangedEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SubscribersChangedEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetOn

`func (o *SubscribersChangedEvent) GetOn() []string`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *SubscribersChangedEvent) GetOnOk() (*[]string, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *SubscribersChangedEvent) SetOn(v []string)`

SetOn sets On field to given value.


### GetOff

`func (o *SubscribersChangedEvent) GetOff() []string`

GetOff returns the Off field if non-nil, zero value otherwise.

### GetOffOk

`func (o *SubscribersChangedEvent) GetOffOk() (*[]string, bool)`

GetOffOk returns a tuple with the Off field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOff

`func (o *SubscribersChangedEvent) SetOff(v []string)`

SetOff sets Off field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


