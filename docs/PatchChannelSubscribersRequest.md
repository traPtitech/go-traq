# PatchChannelSubscribersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**On** | Pointer to **[]string** | 通知をオンにするユーザーのUUID配列 | [optional] 
**Off** | Pointer to **[]string** | 通知をオフにするユーザーのUUID配列 | [optional] 

## Methods

### NewPatchChannelSubscribersRequest

`func NewPatchChannelSubscribersRequest() *PatchChannelSubscribersRequest`

NewPatchChannelSubscribersRequest instantiates a new PatchChannelSubscribersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchChannelSubscribersRequestWithDefaults

`func NewPatchChannelSubscribersRequestWithDefaults() *PatchChannelSubscribersRequest`

NewPatchChannelSubscribersRequestWithDefaults instantiates a new PatchChannelSubscribersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOn

`func (o *PatchChannelSubscribersRequest) GetOn() []string`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *PatchChannelSubscribersRequest) GetOnOk() (*[]string, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *PatchChannelSubscribersRequest) SetOn(v []string)`

SetOn sets On field to given value.

### HasOn

`func (o *PatchChannelSubscribersRequest) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetOff

`func (o *PatchChannelSubscribersRequest) GetOff() []string`

GetOff returns the Off field if non-nil, zero value otherwise.

### GetOffOk

`func (o *PatchChannelSubscribersRequest) GetOffOk() (*[]string, bool)`

GetOffOk returns a tuple with the Off field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOff

`func (o *PatchChannelSubscribersRequest) SetOff(v []string)`

SetOff sets Off field to given value.

### HasOff

`func (o *PatchChannelSubscribersRequest) HasOff() bool`

HasOff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


