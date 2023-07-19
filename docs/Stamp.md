# Stamp

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

## Methods

### NewStamp

`func NewStamp(id string, name string, creatorId string, createdAt time.Time, updatedAt time.Time, fileId string, isUnicode bool, ) *Stamp`

NewStamp instantiates a new Stamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStampWithDefaults

`func NewStampWithDefaults() *Stamp`

NewStampWithDefaults instantiates a new Stamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Stamp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Stamp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Stamp) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Stamp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stamp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stamp) SetName(v string)`

SetName sets Name field to given value.


### GetCreatorId

`func (o *Stamp) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Stamp) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Stamp) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetCreatedAt

`func (o *Stamp) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Stamp) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Stamp) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Stamp) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Stamp) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Stamp) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFileId

`func (o *Stamp) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Stamp) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Stamp) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetIsUnicode

`func (o *Stamp) GetIsUnicode() bool`

GetIsUnicode returns the IsUnicode field if non-nil, zero value otherwise.

### GetIsUnicodeOk

`func (o *Stamp) GetIsUnicodeOk() (*bool, bool)`

GetIsUnicodeOk returns a tuple with the IsUnicode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnicode

`func (o *Stamp) SetIsUnicode(v bool)`

SetIsUnicode sets IsUnicode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


