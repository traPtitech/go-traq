# UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | グループUUID | 
**Name** | **string** | グループ名 | 
**Description** | **string** | グループ説明 | 
**Type** | **string** | グループタイプ | 
**Icon** | **string** | グループアイコンUUID | 
**Members** | [**[]UserGroupMember**](UserGroupMember.md) | グループメンバーの配列 | 
**CreatedAt** | **time.Time** | 作成日時 | 
**UpdatedAt** | **time.Time** | 更新日時 | 
**Admins** | **[]string** | グループ管理者のUUIDの配列 | 

## Methods

### NewUserGroup

`func NewUserGroup(id string, name string, description string, type_ string, icon string, members []UserGroupMember, createdAt time.Time, updatedAt time.Time, admins []string, ) *UserGroup`

NewUserGroup instantiates a new UserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupWithDefaults

`func NewUserGroupWithDefaults() *UserGroup`

NewUserGroupWithDefaults instantiates a new UserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *UserGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGroup) SetType(v string)`

SetType sets Type field to given value.


### GetIcon

`func (o *UserGroup) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UserGroup) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UserGroup) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetMembers

`func (o *UserGroup) GetMembers() []UserGroupMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *UserGroup) GetMembersOk() (*[]UserGroupMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *UserGroup) SetMembers(v []UserGroupMember)`

SetMembers sets Members field to given value.


### GetCreatedAt

`func (o *UserGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAdmins

`func (o *UserGroup) GetAdmins() []string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *UserGroup) GetAdminsOk() (*[]string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *UserGroup) SetAdmins(v []string)`

SetAdmins sets Admins field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


