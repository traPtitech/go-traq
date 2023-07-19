# PatchStampRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | スタンプ名 | [optional] 
**CreatorId** | Pointer to **string** | 作成者UUID | [optional] 

## Methods

### NewPatchStampRequest

`func NewPatchStampRequest() *PatchStampRequest`

NewPatchStampRequest instantiates a new PatchStampRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchStampRequestWithDefaults

`func NewPatchStampRequestWithDefaults() *PatchStampRequest`

NewPatchStampRequestWithDefaults instantiates a new PatchStampRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchStampRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchStampRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchStampRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchStampRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatorId

`func (o *PatchStampRequest) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *PatchStampRequest) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *PatchStampRequest) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *PatchStampRequest) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


