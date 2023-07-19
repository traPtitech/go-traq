# MyUserDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ユーザーUUID | 
**Bio** | **string** | 自己紹介(biography) | 
**Groups** | **[]string** | 所属グループのUUIDの配列 | 
**Tags** | [**[]UserTag**](UserTag.md) | タグリスト | 
**UpdatedAt** | **time.Time** | 更新日時 | 
**LastOnline** | **NullableTime** | 最終オンライン日時 | 
**TwitterId** | **string** | Twitter ID | 
**Name** | **string** | ユーザー名 | 
**DisplayName** | **string** | ユーザー表示名 | 
**IconFileId** | **string** | アイコンファイルUUID | 
**Bot** | **bool** | BOTかどうか | 
**State** | [**UserAccountState**](UserAccountState.md) |  | 
**Permissions** | [**[]UserPermission**](UserPermission.md) | 所有している権限の配列 | 
**HomeChannel** | **NullableString** | ホームチャンネル | 

## Methods

### NewMyUserDetail

`func NewMyUserDetail(id string, bio string, groups []string, tags []UserTag, updatedAt time.Time, lastOnline NullableTime, twitterId string, name string, displayName string, iconFileId string, bot bool, state UserAccountState, permissions []UserPermission, homeChannel NullableString, ) *MyUserDetail`

NewMyUserDetail instantiates a new MyUserDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyUserDetailWithDefaults

`func NewMyUserDetailWithDefaults() *MyUserDetail`

NewMyUserDetailWithDefaults instantiates a new MyUserDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MyUserDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyUserDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyUserDetail) SetId(v string)`

SetId sets Id field to given value.


### GetBio

`func (o *MyUserDetail) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *MyUserDetail) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *MyUserDetail) SetBio(v string)`

SetBio sets Bio field to given value.


### GetGroups

`func (o *MyUserDetail) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MyUserDetail) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *MyUserDetail) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetTags

`func (o *MyUserDetail) GetTags() []UserTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MyUserDetail) GetTagsOk() (*[]UserTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MyUserDetail) SetTags(v []UserTag)`

SetTags sets Tags field to given value.


### GetUpdatedAt

`func (o *MyUserDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MyUserDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MyUserDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLastOnline

`func (o *MyUserDetail) GetLastOnline() time.Time`

GetLastOnline returns the LastOnline field if non-nil, zero value otherwise.

### GetLastOnlineOk

`func (o *MyUserDetail) GetLastOnlineOk() (*time.Time, bool)`

GetLastOnlineOk returns a tuple with the LastOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnline

`func (o *MyUserDetail) SetLastOnline(v time.Time)`

SetLastOnline sets LastOnline field to given value.


### SetLastOnlineNil

`func (o *MyUserDetail) SetLastOnlineNil(b bool)`

 SetLastOnlineNil sets the value for LastOnline to be an explicit nil

### UnsetLastOnline
`func (o *MyUserDetail) UnsetLastOnline()`

UnsetLastOnline ensures that no value is present for LastOnline, not even an explicit nil
### GetTwitterId

`func (o *MyUserDetail) GetTwitterId() string`

GetTwitterId returns the TwitterId field if non-nil, zero value otherwise.

### GetTwitterIdOk

`func (o *MyUserDetail) GetTwitterIdOk() (*string, bool)`

GetTwitterIdOk returns a tuple with the TwitterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterId

`func (o *MyUserDetail) SetTwitterId(v string)`

SetTwitterId sets TwitterId field to given value.


### GetName

`func (o *MyUserDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MyUserDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MyUserDetail) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *MyUserDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MyUserDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MyUserDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIconFileId

`func (o *MyUserDetail) GetIconFileId() string`

GetIconFileId returns the IconFileId field if non-nil, zero value otherwise.

### GetIconFileIdOk

`func (o *MyUserDetail) GetIconFileIdOk() (*string, bool)`

GetIconFileIdOk returns a tuple with the IconFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFileId

`func (o *MyUserDetail) SetIconFileId(v string)`

SetIconFileId sets IconFileId field to given value.


### GetBot

`func (o *MyUserDetail) GetBot() bool`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *MyUserDetail) GetBotOk() (*bool, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *MyUserDetail) SetBot(v bool)`

SetBot sets Bot field to given value.


### GetState

`func (o *MyUserDetail) GetState() UserAccountState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MyUserDetail) GetStateOk() (*UserAccountState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MyUserDetail) SetState(v UserAccountState)`

SetState sets State field to given value.


### GetPermissions

`func (o *MyUserDetail) GetPermissions() []UserPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MyUserDetail) GetPermissionsOk() (*[]UserPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MyUserDetail) SetPermissions(v []UserPermission)`

SetPermissions sets Permissions field to given value.


### GetHomeChannel

`func (o *MyUserDetail) GetHomeChannel() string`

GetHomeChannel returns the HomeChannel field if non-nil, zero value otherwise.

### GetHomeChannelOk

`func (o *MyUserDetail) GetHomeChannelOk() (*string, bool)`

GetHomeChannelOk returns a tuple with the HomeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeChannel

`func (o *MyUserDetail) SetHomeChannel(v string)`

SetHomeChannel sets HomeChannel field to given value.


### SetHomeChannelNil

`func (o *MyUserDetail) SetHomeChannelNil(b bool)`

 SetHomeChannelNil sets the value for HomeChannel to be an explicit nil

### UnsetHomeChannel
`func (o *MyUserDetail) UnsetHomeChannel()`

UnsetHomeChannel ensures that no value is present for HomeChannel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


