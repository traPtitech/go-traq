# PostChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | チャンネル名 | 
**Parent** | **NullableString** | 親チャンネルのUUID ルートに作成する場合はnullを指定 | 

## Methods

### NewPostChannelRequest

`func NewPostChannelRequest(name string, parent NullableString, ) *PostChannelRequest`

NewPostChannelRequest instantiates a new PostChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostChannelRequestWithDefaults

`func NewPostChannelRequestWithDefaults() *PostChannelRequest`

NewPostChannelRequestWithDefaults instantiates a new PostChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostChannelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *PostChannelRequest) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PostChannelRequest) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PostChannelRequest) SetParent(v string)`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *PostChannelRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PostChannelRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


