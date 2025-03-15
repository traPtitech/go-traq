# SoundboardItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SoundId** | **string** | サーバが発行したサウンドID | 
**SoundName** | **string** | ユーザが指定した表示用のサウンド名 | 
**StampId** | **string** | 任意のスタンプID等、サウンドに紐づく拡張情報 | 
**CreatorId** | **string** | 作成者のユーザID | 

## Methods

### NewSoundboardItem

`func NewSoundboardItem(soundId string, soundName string, stampId string, creatorId string, ) *SoundboardItem`

NewSoundboardItem instantiates a new SoundboardItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoundboardItemWithDefaults

`func NewSoundboardItemWithDefaults() *SoundboardItem`

NewSoundboardItemWithDefaults instantiates a new SoundboardItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoundId

`func (o *SoundboardItem) GetSoundId() string`

GetSoundId returns the SoundId field if non-nil, zero value otherwise.

### GetSoundIdOk

`func (o *SoundboardItem) GetSoundIdOk() (*string, bool)`

GetSoundIdOk returns a tuple with the SoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundId

`func (o *SoundboardItem) SetSoundId(v string)`

SetSoundId sets SoundId field to given value.


### GetSoundName

`func (o *SoundboardItem) GetSoundName() string`

GetSoundName returns the SoundName field if non-nil, zero value otherwise.

### GetSoundNameOk

`func (o *SoundboardItem) GetSoundNameOk() (*string, bool)`

GetSoundNameOk returns a tuple with the SoundName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundName

`func (o *SoundboardItem) SetSoundName(v string)`

SetSoundName sets SoundName field to given value.


### GetStampId

`func (o *SoundboardItem) GetStampId() string`

GetStampId returns the StampId field if non-nil, zero value otherwise.

### GetStampIdOk

`func (o *SoundboardItem) GetStampIdOk() (*string, bool)`

GetStampIdOk returns a tuple with the StampId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampId

`func (o *SoundboardItem) SetStampId(v string)`

SetStampId sets StampId field to given value.


### GetCreatorId

`func (o *SoundboardItem) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *SoundboardItem) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *SoundboardItem) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


