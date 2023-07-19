# UserTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | **string** | タグUUID | 
**Tag** | **string** | タグ文字列 | 
**IsLocked** | **bool** | タグがロックされているか | 
**CreatedAt** | **time.Time** | タグ付与日時 | 
**UpdatedAt** | **time.Time** | タグ更新日時 | 

## Methods

### NewUserTag

`func NewUserTag(tagId string, tag string, isLocked bool, createdAt time.Time, updatedAt time.Time, ) *UserTag`

NewUserTag instantiates a new UserTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTagWithDefaults

`func NewUserTagWithDefaults() *UserTag`

NewUserTagWithDefaults instantiates a new UserTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *UserTag) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *UserTag) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *UserTag) SetTagId(v string)`

SetTagId sets TagId field to given value.


### GetTag

`func (o *UserTag) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *UserTag) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *UserTag) SetTag(v string)`

SetTag sets Tag field to given value.


### GetIsLocked

`func (o *UserTag) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UserTag) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UserTag) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetCreatedAt

`func (o *UserTag) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserTag) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserTag) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserTag) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserTag) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserTag) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


