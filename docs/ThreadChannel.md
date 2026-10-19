# ThreadChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | チャンネルUUID | 
**ParentChannelId** | **string** | 親チャンネルUUID | 
**ParentMessageId** | **string** | 親メッセージUUID | 
**Archived** | **bool** | スレッドがアーカイブされているかどうか | 
**Name** | **string** | スレッド名 | 

## Methods

### NewThreadChannel

`func NewThreadChannel(id string, parentChannelId string, parentMessageId string, archived bool, name string, ) *ThreadChannel`

NewThreadChannel instantiates a new ThreadChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadChannelWithDefaults

`func NewThreadChannelWithDefaults() *ThreadChannel`

NewThreadChannelWithDefaults instantiates a new ThreadChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThreadChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreadChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreadChannel) SetId(v string)`

SetId sets Id field to given value.


### GetParentChannelId

`func (o *ThreadChannel) GetParentChannelId() string`

GetParentChannelId returns the ParentChannelId field if non-nil, zero value otherwise.

### GetParentChannelIdOk

`func (o *ThreadChannel) GetParentChannelIdOk() (*string, bool)`

GetParentChannelIdOk returns a tuple with the ParentChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChannelId

`func (o *ThreadChannel) SetParentChannelId(v string)`

SetParentChannelId sets ParentChannelId field to given value.


### GetParentMessageId

`func (o *ThreadChannel) GetParentMessageId() string`

GetParentMessageId returns the ParentMessageId field if non-nil, zero value otherwise.

### GetParentMessageIdOk

`func (o *ThreadChannel) GetParentMessageIdOk() (*string, bool)`

GetParentMessageIdOk returns a tuple with the ParentMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentMessageId

`func (o *ThreadChannel) SetParentMessageId(v string)`

SetParentMessageId sets ParentMessageId field to given value.


### GetArchived

`func (o *ThreadChannel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ThreadChannel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ThreadChannel) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetName

`func (o *ThreadChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreadChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreadChannel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


