# UserOfflineEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | オフラインになったユーザーのId | 
**LastOnline** | **time.Time** | 最後の接続の切断をサーバーが検出した日時 | 

## Methods

### NewUserOfflineEvent

`func NewUserOfflineEvent(id string, lastOnline time.Time, ) *UserOfflineEvent`

NewUserOfflineEvent instantiates a new UserOfflineEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOfflineEventWithDefaults

`func NewUserOfflineEventWithDefaults() *UserOfflineEvent`

NewUserOfflineEventWithDefaults instantiates a new UserOfflineEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserOfflineEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserOfflineEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserOfflineEvent) SetId(v string)`

SetId sets Id field to given value.


### GetLastOnline

`func (o *UserOfflineEvent) GetLastOnline() time.Time`

GetLastOnline returns the LastOnline field if non-nil, zero value otherwise.

### GetLastOnlineOk

`func (o *UserOfflineEvent) GetLastOnlineOk() (*time.Time, bool)`

GetLastOnlineOk returns a tuple with the LastOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnline

`func (o *UserOfflineEvent) SetLastOnline(v time.Time)`

SetLastOnline sets LastOnline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


