# ClipFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | フォルダUUID | 
**Name** | **string** | フォルダ名 | 
**CreatedAt** | **time.Time** | 作成日時 | 
**OwnerId** | **string** | フォルダ所有者UUID | 
**Description** | **string** | 説明 | 

## Methods

### NewClipFolder

`func NewClipFolder(id string, name string, createdAt time.Time, ownerId string, description string, ) *ClipFolder`

NewClipFolder instantiates a new ClipFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClipFolderWithDefaults

`func NewClipFolderWithDefaults() *ClipFolder`

NewClipFolderWithDefaults instantiates a new ClipFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClipFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClipFolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClipFolder) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ClipFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClipFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClipFolder) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *ClipFolder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClipFolder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClipFolder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetOwnerId

`func (o *ClipFolder) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ClipFolder) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ClipFolder) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetDescription

`func (o *ClipFolder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClipFolder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClipFolder) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


