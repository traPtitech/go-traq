# PatchChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | チャンネル名 | [optional] 
**Archived** | Pointer to **bool** | アーカイブされているかどうか | [optional] 
**Force** | Pointer to **bool** | 強制通知チャンネルかどうか | [optional] 
**Parent** | Pointer to **string** | 親チャンネルUUID | [optional] 

## Methods

### NewPatchChannelRequest

`func NewPatchChannelRequest() *PatchChannelRequest`

NewPatchChannelRequest instantiates a new PatchChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchChannelRequestWithDefaults

`func NewPatchChannelRequestWithDefaults() *PatchChannelRequest`

NewPatchChannelRequestWithDefaults instantiates a new PatchChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchChannelRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchChannelRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchived

`func (o *PatchChannelRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PatchChannelRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PatchChannelRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PatchChannelRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetForce

`func (o *PatchChannelRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *PatchChannelRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *PatchChannelRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *PatchChannelRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetParent

`func (o *PatchChannelRequest) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchChannelRequest) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchChannelRequest) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchChannelRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


