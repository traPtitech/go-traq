# OIDCUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | **string** | ユーザーUUID | 
**Name** | **string** | ユーザー名 | 
**PreferredUsername** | **string** | ユーザー名 | 
**Picture** | **string** | アイコン画像URL | 
**UpdatedAt** | Pointer to **int64** | 更新日時 | [optional] 
**Traq** | Pointer to [**OIDCTraqUserInfo**](OIDCTraqUserInfo.md) |  | [optional] 

## Methods

### NewOIDCUserInfo

`func NewOIDCUserInfo(sub string, name string, preferredUsername string, picture string, ) *OIDCUserInfo`

NewOIDCUserInfo instantiates a new OIDCUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCUserInfoWithDefaults

`func NewOIDCUserInfoWithDefaults() *OIDCUserInfo`

NewOIDCUserInfoWithDefaults instantiates a new OIDCUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *OIDCUserInfo) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *OIDCUserInfo) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *OIDCUserInfo) SetSub(v string)`

SetSub sets Sub field to given value.


### GetName

`func (o *OIDCUserInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OIDCUserInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OIDCUserInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPreferredUsername

`func (o *OIDCUserInfo) GetPreferredUsername() string`

GetPreferredUsername returns the PreferredUsername field if non-nil, zero value otherwise.

### GetPreferredUsernameOk

`func (o *OIDCUserInfo) GetPreferredUsernameOk() (*string, bool)`

GetPreferredUsernameOk returns a tuple with the PreferredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUsername

`func (o *OIDCUserInfo) SetPreferredUsername(v string)`

SetPreferredUsername sets PreferredUsername field to given value.


### GetPicture

`func (o *OIDCUserInfo) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *OIDCUserInfo) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *OIDCUserInfo) SetPicture(v string)`

SetPicture sets Picture field to given value.


### GetUpdatedAt

`func (o *OIDCUserInfo) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OIDCUserInfo) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OIDCUserInfo) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OIDCUserInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTraq

`func (o *OIDCUserInfo) GetTraq() OIDCTraqUserInfo`

GetTraq returns the Traq field if non-nil, zero value otherwise.

### GetTraqOk

`func (o *OIDCUserInfo) GetTraqOk() (*OIDCTraqUserInfo, bool)`

GetTraqOk returns a tuple with the Traq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraq

`func (o *OIDCUserInfo) SetTraq(v OIDCTraqUserInfo)`

SetTraq sets Traq field to given value.

### HasTraq

`func (o *OIDCUserInfo) HasTraq() bool`

HasTraq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


