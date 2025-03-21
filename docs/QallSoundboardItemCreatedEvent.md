# QallSoundboardItemCreatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SoundId** | Pointer to **string** | 作成されたサウンドボードアイテムのId | [optional] 
**Name** | Pointer to **string** | 作成されたサウンドボードアイテムの名前 | [optional] 
**CreatorId** | Pointer to **string** | 作成者のId | [optional] 

## Methods

### NewQallSoundboardItemCreatedEvent

`func NewQallSoundboardItemCreatedEvent() *QallSoundboardItemCreatedEvent`

NewQallSoundboardItemCreatedEvent instantiates a new QallSoundboardItemCreatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQallSoundboardItemCreatedEventWithDefaults

`func NewQallSoundboardItemCreatedEventWithDefaults() *QallSoundboardItemCreatedEvent`

NewQallSoundboardItemCreatedEventWithDefaults instantiates a new QallSoundboardItemCreatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoundId

`func (o *QallSoundboardItemCreatedEvent) GetSoundId() string`

GetSoundId returns the SoundId field if non-nil, zero value otherwise.

### GetSoundIdOk

`func (o *QallSoundboardItemCreatedEvent) GetSoundIdOk() (*string, bool)`

GetSoundIdOk returns a tuple with the SoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundId

`func (o *QallSoundboardItemCreatedEvent) SetSoundId(v string)`

SetSoundId sets SoundId field to given value.

### HasSoundId

`func (o *QallSoundboardItemCreatedEvent) HasSoundId() bool`

HasSoundId returns a boolean if a field has been set.

### GetName

`func (o *QallSoundboardItemCreatedEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QallSoundboardItemCreatedEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QallSoundboardItemCreatedEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QallSoundboardItemCreatedEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatorId

`func (o *QallSoundboardItemCreatedEvent) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *QallSoundboardItemCreatedEvent) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *QallSoundboardItemCreatedEvent) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *QallSoundboardItemCreatedEvent) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


