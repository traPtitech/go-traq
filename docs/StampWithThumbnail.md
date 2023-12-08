# StampWithThumbnail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | スタンプUUID | 
**Name** | **string** | スタンプ名 | 
**CreatorId** | **string** | 作成者UUID | 
**CreatedAt** | **time.Time** | 作成日時 | 
**UpdatedAt** | **time.Time** | 更新日時 | 
**FileId** | **string** | ファイルUUID | 
**IsUnicode** | **bool** | Unicode絵文字か | 
**HasThumbnail** | **bool** | サムネイルの有無 | 

## Methods

### NewStampWithThumbnail

`func NewStampWithThumbnail(id string, name string, creatorId string, createdAt time.Time, updatedAt time.Time, fileId string, isUnicode bool, hasThumbnail bool, ) *StampWithThumbnail`

NewStampWithThumbnail instantiates a new StampWithThumbnail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStampWithThumbnailWithDefaults

`func NewStampWithThumbnailWithDefaults() *StampWithThumbnail`

NewStampWithThumbnailWithDefaults instantiates a new StampWithThumbnail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StampWithThumbnail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StampWithThumbnail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StampWithThumbnail) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StampWithThumbnail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StampWithThumbnail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StampWithThumbnail) SetName(v string)`

SetName sets Name field to given value.


### GetCreatorId

`func (o *StampWithThumbnail) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *StampWithThumbnail) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *StampWithThumbnail) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetCreatedAt

`func (o *StampWithThumbnail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StampWithThumbnail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StampWithThumbnail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *StampWithThumbnail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StampWithThumbnail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StampWithThumbnail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFileId

`func (o *StampWithThumbnail) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *StampWithThumbnail) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *StampWithThumbnail) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetIsUnicode

`func (o *StampWithThumbnail) GetIsUnicode() bool`

GetIsUnicode returns the IsUnicode field if non-nil, zero value otherwise.

### GetIsUnicodeOk

`func (o *StampWithThumbnail) GetIsUnicodeOk() (*bool, bool)`

GetIsUnicodeOk returns a tuple with the IsUnicode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnicode

`func (o *StampWithThumbnail) SetIsUnicode(v bool)`

SetIsUnicode sets IsUnicode field to given value.


### GetHasThumbnail

`func (o *StampWithThumbnail) GetHasThumbnail() bool`

GetHasThumbnail returns the HasThumbnail field if non-nil, zero value otherwise.

### GetHasThumbnailOk

`func (o *StampWithThumbnail) GetHasThumbnailOk() (*bool, bool)`

GetHasThumbnailOk returns a tuple with the HasThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasThumbnail

`func (o *StampWithThumbnail) SetHasThumbnail(v bool)`

SetHasThumbnail sets HasThumbnail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


