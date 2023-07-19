# PostUserGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | グループ名 | 
**Description** | **string** | 説明 | 
**Type** | **string** | グループタイプ | 

## Methods

### NewPostUserGroupRequest

`func NewPostUserGroupRequest(name string, description string, type_ string, ) *PostUserGroupRequest`

NewPostUserGroupRequest instantiates a new PostUserGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostUserGroupRequestWithDefaults

`func NewPostUserGroupRequestWithDefaults() *PostUserGroupRequest`

NewPostUserGroupRequestWithDefaults instantiates a new PostUserGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostUserGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostUserGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostUserGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostUserGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostUserGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostUserGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *PostUserGroupRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostUserGroupRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostUserGroupRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


