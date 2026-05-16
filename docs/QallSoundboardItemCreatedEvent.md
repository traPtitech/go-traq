# QallSoundboardItemCreatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SoundId** | **string** | 作成されたサウンドボードアイテムのId | 
**Name** | **string** | 作成されたサウンドボードアイテムの名前 | 
**CreatorId** | **string** | 作成者のId | 

## Methods

### NewQallSoundboardItemCreatedEvent

`func NewQallSoundboardItemCreatedEvent(soundId string, name string, creatorId string, ) *QallSoundboardItemCreatedEvent`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


