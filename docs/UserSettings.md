# UserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ユーザーUUID | 
**NotifyCitation** | **bool** | メッセージ引用通知の設定情報 | 

## Methods

### NewUserSettings

`func NewUserSettings(id string, notifyCitation bool, ) *UserSettings`

NewUserSettings instantiates a new UserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsWithDefaults

`func NewUserSettingsWithDefaults() *UserSettings`

NewUserSettingsWithDefaults instantiates a new UserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSettings) SetId(v string)`

SetId sets Id field to given value.


### GetNotifyCitation

`func (o *UserSettings) GetNotifyCitation() bool`

GetNotifyCitation returns the NotifyCitation field if non-nil, zero value otherwise.

### GetNotifyCitationOk

`func (o *UserSettings) GetNotifyCitationOk() (*bool, bool)`

GetNotifyCitationOk returns a tuple with the NotifyCitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCitation

`func (o *UserSettings) SetNotifyCitation(v bool)`

SetNotifyCitation sets NotifyCitation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


