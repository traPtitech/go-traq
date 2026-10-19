# PatchThreadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | スレッド名 | [optional] 
**Archived** | Pointer to **bool** | アーカイブされているかどうか | [optional] 

## Methods

### NewPatchThreadRequest

`func NewPatchThreadRequest() *PatchThreadRequest`

NewPatchThreadRequest instantiates a new PatchThreadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchThreadRequestWithDefaults

`func NewPatchThreadRequestWithDefaults() *PatchThreadRequest`

NewPatchThreadRequestWithDefaults instantiates a new PatchThreadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchThreadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchThreadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchThreadRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchThreadRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchived

`func (o *PatchThreadRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PatchThreadRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PatchThreadRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PatchThreadRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


