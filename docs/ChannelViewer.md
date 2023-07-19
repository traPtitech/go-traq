# ChannelViewer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | ユーザーUUID | 
**State** | [**ChannelViewState**](ChannelViewState.md) |  | 
**UpdatedAt** | **time.Time** | 更新日時 | 

## Methods

### NewChannelViewer

`func NewChannelViewer(userId string, state ChannelViewState, updatedAt time.Time, ) *ChannelViewer`

NewChannelViewer instantiates a new ChannelViewer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelViewerWithDefaults

`func NewChannelViewerWithDefaults() *ChannelViewer`

NewChannelViewerWithDefaults instantiates a new ChannelViewer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ChannelViewer) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ChannelViewer) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ChannelViewer) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetState

`func (o *ChannelViewer) GetState() ChannelViewState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ChannelViewer) GetStateOk() (*ChannelViewState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ChannelViewer) SetState(v ChannelViewState)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *ChannelViewer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChannelViewer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChannelViewer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


