# SoundboardPlayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SoundId** | **string** | サウンドID (DB登録済み) | 
**RoomName** | **string** | 再生させたいルームのUUID | 

## Methods

### NewSoundboardPlayRequest

`func NewSoundboardPlayRequest(soundId string, roomName string, ) *SoundboardPlayRequest`

NewSoundboardPlayRequest instantiates a new SoundboardPlayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoundboardPlayRequestWithDefaults

`func NewSoundboardPlayRequestWithDefaults() *SoundboardPlayRequest`

NewSoundboardPlayRequestWithDefaults instantiates a new SoundboardPlayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoundId

`func (o *SoundboardPlayRequest) GetSoundId() string`

GetSoundId returns the SoundId field if non-nil, zero value otherwise.

### GetSoundIdOk

`func (o *SoundboardPlayRequest) GetSoundIdOk() (*string, bool)`

GetSoundIdOk returns a tuple with the SoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundId

`func (o *SoundboardPlayRequest) SetSoundId(v string)`

SetSoundId sets SoundId field to given value.


### GetRoomName

`func (o *SoundboardPlayRequest) GetRoomName() string`

GetRoomName returns the RoomName field if non-nil, zero value otherwise.

### GetRoomNameOk

`func (o *SoundboardPlayRequest) GetRoomNameOk() (*string, bool)`

GetRoomNameOk returns a tuple with the RoomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomName

`func (o *SoundboardPlayRequest) SetRoomName(v string)`

SetRoomName sets RoomName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


