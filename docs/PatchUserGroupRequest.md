# PatchUserGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | グループ名 | [optional] 
**Description** | Pointer to **string** | グループ説明 | [optional] 
**Type** | Pointer to **string** | グループタイプ | [optional] 

## Methods

### NewPatchUserGroupRequest

`func NewPatchUserGroupRequest() *PatchUserGroupRequest`

NewPatchUserGroupRequest instantiates a new PatchUserGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchUserGroupRequestWithDefaults

`func NewPatchUserGroupRequestWithDefaults() *PatchUserGroupRequest`

NewPatchUserGroupRequestWithDefaults instantiates a new PatchUserGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchUserGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchUserGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchUserGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchUserGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchUserGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchUserGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchUserGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchUserGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PatchUserGroupRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchUserGroupRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchUserGroupRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchUserGroupRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


