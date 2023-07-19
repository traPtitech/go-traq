# UserDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ユーザーUUID | 
**State** | [**UserAccountState**](UserAccountState.md) |  | 
**Bot** | **bool** | BOTかどうか | 
**IconFileId** | **string** | アイコンファイルUUID | 
**DisplayName** | **string** | ユーザー表示名 | 
**Name** | **string** | ユーザー名 | 
**TwitterId** | **string** | Twitter ID | 
**LastOnline** | **NullableTime** | 最終オンライン日時 | 
**UpdatedAt** | **time.Time** | 更新日時 | 
**Tags** | [**[]UserTag**](UserTag.md) | タグリスト | 
**Groups** | **[]string** | 所属グループのUUIDの配列 | 
**Bio** | **string** | 自己紹介(biography) | 
**HomeChannel** | **NullableString** | ホームチャンネル | 

## Methods

### NewUserDetail

`func NewUserDetail(id string, state UserAccountState, bot bool, iconFileId string, displayName string, name string, twitterId string, lastOnline NullableTime, updatedAt time.Time, tags []UserTag, groups []string, bio string, homeChannel NullableString, ) *UserDetail`

NewUserDetail instantiates a new UserDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailWithDefaults

`func NewUserDetailWithDefaults() *UserDetail`

NewUserDetailWithDefaults instantiates a new UserDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDetail) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *UserDetail) GetState() UserAccountState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserDetail) GetStateOk() (*UserAccountState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserDetail) SetState(v UserAccountState)`

SetState sets State field to given value.


### GetBot

`func (o *UserDetail) GetBot() bool`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *UserDetail) GetBotOk() (*bool, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *UserDetail) SetBot(v bool)`

SetBot sets Bot field to given value.


### GetIconFileId

`func (o *UserDetail) GetIconFileId() string`

GetIconFileId returns the IconFileId field if non-nil, zero value otherwise.

### GetIconFileIdOk

`func (o *UserDetail) GetIconFileIdOk() (*string, bool)`

GetIconFileIdOk returns a tuple with the IconFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFileId

`func (o *UserDetail) SetIconFileId(v string)`

SetIconFileId sets IconFileId field to given value.


### GetDisplayName

`func (o *UserDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetName

`func (o *UserDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDetail) SetName(v string)`

SetName sets Name field to given value.


### GetTwitterId

`func (o *UserDetail) GetTwitterId() string`

GetTwitterId returns the TwitterId field if non-nil, zero value otherwise.

### GetTwitterIdOk

`func (o *UserDetail) GetTwitterIdOk() (*string, bool)`

GetTwitterIdOk returns a tuple with the TwitterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterId

`func (o *UserDetail) SetTwitterId(v string)`

SetTwitterId sets TwitterId field to given value.


### GetLastOnline

`func (o *UserDetail) GetLastOnline() time.Time`

GetLastOnline returns the LastOnline field if non-nil, zero value otherwise.

### GetLastOnlineOk

`func (o *UserDetail) GetLastOnlineOk() (*time.Time, bool)`

GetLastOnlineOk returns a tuple with the LastOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnline

`func (o *UserDetail) SetLastOnline(v time.Time)`

SetLastOnline sets LastOnline field to given value.


### SetLastOnlineNil

`func (o *UserDetail) SetLastOnlineNil(b bool)`

 SetLastOnlineNil sets the value for LastOnline to be an explicit nil

### UnsetLastOnline
`func (o *UserDetail) UnsetLastOnline()`

UnsetLastOnline ensures that no value is present for LastOnline, not even an explicit nil
### GetUpdatedAt

`func (o *UserDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTags

`func (o *UserDetail) GetTags() []UserTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UserDetail) GetTagsOk() (*[]UserTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UserDetail) SetTags(v []UserTag)`

SetTags sets Tags field to given value.


### GetGroups

`func (o *UserDetail) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserDetail) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserDetail) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetBio

`func (o *UserDetail) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UserDetail) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UserDetail) SetBio(v string)`

SetBio sets Bio field to given value.


### GetHomeChannel

`func (o *UserDetail) GetHomeChannel() string`

GetHomeChannel returns the HomeChannel field if non-nil, zero value otherwise.

### GetHomeChannelOk

`func (o *UserDetail) GetHomeChannelOk() (*string, bool)`

GetHomeChannelOk returns a tuple with the HomeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeChannel

`func (o *UserDetail) SetHomeChannel(v string)`

SetHomeChannel sets HomeChannel field to given value.


### SetHomeChannelNil

`func (o *UserDetail) SetHomeChannelNil(b bool)`

 SetHomeChannelNil sets the value for HomeChannel to be an explicit nil

### UnsetHomeChannel
`func (o *UserDetail) UnsetHomeChannel()`

UnsetHomeChannel ensures that no value is present for HomeChannel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


