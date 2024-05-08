# OIDCTraqUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bio** | **string** | 自己紹介(biography) | 
**Groups** | **[]string** | 所属グループのUUIDの配列 | 
**Tags** | [**[]UserTag**](UserTag.md) | タグリスト | 
**LastOnline** | **NullableTime** | 最終オンライン日時 | 
**TwitterId** | **string** | Twitter ID | 
**DisplayName** | **string** | ユーザー表示名 | 
**IconFileId** | **string** | アイコンファイルUUID | 
**Bot** | **bool** | BOTかどうか | 
**State** | [**UserAccountState**](UserAccountState.md) |  | 
**Permissions** | [**[]UserPermission**](UserPermission.md) | 所有している権限の配列 | 
**HomeChannel** | **NullableString** | ホームチャンネル | 

## Methods

### NewOIDCTraqUserInfo

`func NewOIDCTraqUserInfo(bio string, groups []string, tags []UserTag, lastOnline NullableTime, twitterId string, displayName string, iconFileId string, bot bool, state UserAccountState, permissions []UserPermission, homeChannel NullableString, ) *OIDCTraqUserInfo`

NewOIDCTraqUserInfo instantiates a new OIDCTraqUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCTraqUserInfoWithDefaults

`func NewOIDCTraqUserInfoWithDefaults() *OIDCTraqUserInfo`

NewOIDCTraqUserInfoWithDefaults instantiates a new OIDCTraqUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBio

`func (o *OIDCTraqUserInfo) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *OIDCTraqUserInfo) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *OIDCTraqUserInfo) SetBio(v string)`

SetBio sets Bio field to given value.


### GetGroups

`func (o *OIDCTraqUserInfo) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *OIDCTraqUserInfo) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *OIDCTraqUserInfo) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetTags

`func (o *OIDCTraqUserInfo) GetTags() []UserTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OIDCTraqUserInfo) GetTagsOk() (*[]UserTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OIDCTraqUserInfo) SetTags(v []UserTag)`

SetTags sets Tags field to given value.


### GetLastOnline

`func (o *OIDCTraqUserInfo) GetLastOnline() time.Time`

GetLastOnline returns the LastOnline field if non-nil, zero value otherwise.

### GetLastOnlineOk

`func (o *OIDCTraqUserInfo) GetLastOnlineOk() (*time.Time, bool)`

GetLastOnlineOk returns a tuple with the LastOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnline

`func (o *OIDCTraqUserInfo) SetLastOnline(v time.Time)`

SetLastOnline sets LastOnline field to given value.


### SetLastOnlineNil

`func (o *OIDCTraqUserInfo) SetLastOnlineNil(b bool)`

 SetLastOnlineNil sets the value for LastOnline to be an explicit nil

### UnsetLastOnline
`func (o *OIDCTraqUserInfo) UnsetLastOnline()`

UnsetLastOnline ensures that no value is present for LastOnline, not even an explicit nil
### GetTwitterId

`func (o *OIDCTraqUserInfo) GetTwitterId() string`

GetTwitterId returns the TwitterId field if non-nil, zero value otherwise.

### GetTwitterIdOk

`func (o *OIDCTraqUserInfo) GetTwitterIdOk() (*string, bool)`

GetTwitterIdOk returns a tuple with the TwitterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterId

`func (o *OIDCTraqUserInfo) SetTwitterId(v string)`

SetTwitterId sets TwitterId field to given value.


### GetDisplayName

`func (o *OIDCTraqUserInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OIDCTraqUserInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OIDCTraqUserInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIconFileId

`func (o *OIDCTraqUserInfo) GetIconFileId() string`

GetIconFileId returns the IconFileId field if non-nil, zero value otherwise.

### GetIconFileIdOk

`func (o *OIDCTraqUserInfo) GetIconFileIdOk() (*string, bool)`

GetIconFileIdOk returns a tuple with the IconFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFileId

`func (o *OIDCTraqUserInfo) SetIconFileId(v string)`

SetIconFileId sets IconFileId field to given value.


### GetBot

`func (o *OIDCTraqUserInfo) GetBot() bool`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *OIDCTraqUserInfo) GetBotOk() (*bool, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *OIDCTraqUserInfo) SetBot(v bool)`

SetBot sets Bot field to given value.


### GetState

`func (o *OIDCTraqUserInfo) GetState() UserAccountState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OIDCTraqUserInfo) GetStateOk() (*UserAccountState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OIDCTraqUserInfo) SetState(v UserAccountState)`

SetState sets State field to given value.


### GetPermissions

`func (o *OIDCTraqUserInfo) GetPermissions() []UserPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OIDCTraqUserInfo) GetPermissionsOk() (*[]UserPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OIDCTraqUserInfo) SetPermissions(v []UserPermission)`

SetPermissions sets Permissions field to given value.


### GetHomeChannel

`func (o *OIDCTraqUserInfo) GetHomeChannel() string`

GetHomeChannel returns the HomeChannel field if non-nil, zero value otherwise.

### GetHomeChannelOk

`func (o *OIDCTraqUserInfo) GetHomeChannelOk() (*string, bool)`

GetHomeChannelOk returns a tuple with the HomeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeChannel

`func (o *OIDCTraqUserInfo) SetHomeChannel(v string)`

SetHomeChannel sets HomeChannel field to given value.


### SetHomeChannelNil

`func (o *OIDCTraqUserInfo) SetHomeChannelNil(b bool)`

 SetHomeChannelNil sets the value for HomeChannel to be an explicit nil

### UnsetHomeChannel
`func (o *OIDCTraqUserInfo) UnsetHomeChannel()`

UnsetHomeChannel ensures that no value is present for HomeChannel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


