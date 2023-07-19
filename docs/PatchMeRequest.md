# PatchMeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | 新しい表示名 | [optional] 
**TwitterId** | Pointer to **string** | TwitterID | [optional] 
**Bio** | Pointer to **string** | 自己紹介(biography) | [optional] 
**HomeChannel** | Pointer to **string** | ホームチャンネルのUUID &#x60;00000000-0000-0000-0000-000000000000&#x60;を指定すると、ホームチャンネルが&#x60;null&#x60;に設定されます | [optional] 

## Methods

### NewPatchMeRequest

`func NewPatchMeRequest() *PatchMeRequest`

NewPatchMeRequest instantiates a new PatchMeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchMeRequestWithDefaults

`func NewPatchMeRequestWithDefaults() *PatchMeRequest`

NewPatchMeRequestWithDefaults instantiates a new PatchMeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PatchMeRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchMeRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchMeRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PatchMeRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTwitterId

`func (o *PatchMeRequest) GetTwitterId() string`

GetTwitterId returns the TwitterId field if non-nil, zero value otherwise.

### GetTwitterIdOk

`func (o *PatchMeRequest) GetTwitterIdOk() (*string, bool)`

GetTwitterIdOk returns a tuple with the TwitterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterId

`func (o *PatchMeRequest) SetTwitterId(v string)`

SetTwitterId sets TwitterId field to given value.

### HasTwitterId

`func (o *PatchMeRequest) HasTwitterId() bool`

HasTwitterId returns a boolean if a field has been set.

### GetBio

`func (o *PatchMeRequest) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *PatchMeRequest) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *PatchMeRequest) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *PatchMeRequest) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetHomeChannel

`func (o *PatchMeRequest) GetHomeChannel() string`

GetHomeChannel returns the HomeChannel field if non-nil, zero value otherwise.

### GetHomeChannelOk

`func (o *PatchMeRequest) GetHomeChannelOk() (*string, bool)`

GetHomeChannelOk returns a tuple with the HomeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeChannel

`func (o *PatchMeRequest) SetHomeChannel(v string)`

SetHomeChannel sets HomeChannel field to given value.

### HasHomeChannel

`func (o *PatchMeRequest) HasHomeChannel() bool`

HasHomeChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


