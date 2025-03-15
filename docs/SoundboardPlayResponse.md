# SoundboardPlayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngressId** | **string** | 作成された Ingress のID | 
**Url** | Pointer to **string** | 作成された Ingress のストリームURL等 | [optional] 
**StreamKey** | Pointer to **string** | RTMP配信の場合のstream key | [optional] 

## Methods

### NewSoundboardPlayResponse

`func NewSoundboardPlayResponse(ingressId string, ) *SoundboardPlayResponse`

NewSoundboardPlayResponse instantiates a new SoundboardPlayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoundboardPlayResponseWithDefaults

`func NewSoundboardPlayResponseWithDefaults() *SoundboardPlayResponse`

NewSoundboardPlayResponseWithDefaults instantiates a new SoundboardPlayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngressId

`func (o *SoundboardPlayResponse) GetIngressId() string`

GetIngressId returns the IngressId field if non-nil, zero value otherwise.

### GetIngressIdOk

`func (o *SoundboardPlayResponse) GetIngressIdOk() (*string, bool)`

GetIngressIdOk returns a tuple with the IngressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressId

`func (o *SoundboardPlayResponse) SetIngressId(v string)`

SetIngressId sets IngressId field to given value.


### GetUrl

`func (o *SoundboardPlayResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SoundboardPlayResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SoundboardPlayResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SoundboardPlayResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStreamKey

`func (o *SoundboardPlayResponse) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *SoundboardPlayResponse) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *SoundboardPlayResponse) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.

### HasStreamKey

`func (o *SoundboardPlayResponse) HasStreamKey() bool`

HasStreamKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


