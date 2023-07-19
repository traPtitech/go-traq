# PostStampPaletteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stamps** | **[]string** | パレット内のスタンプのUUID配列 | 
**Name** | **string** | パレット名 | 
**Description** | **string** | 説明 | 

## Methods

### NewPostStampPaletteRequest

`func NewPostStampPaletteRequest(stamps []string, name string, description string, ) *PostStampPaletteRequest`

NewPostStampPaletteRequest instantiates a new PostStampPaletteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostStampPaletteRequestWithDefaults

`func NewPostStampPaletteRequestWithDefaults() *PostStampPaletteRequest`

NewPostStampPaletteRequestWithDefaults instantiates a new PostStampPaletteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStamps

`func (o *PostStampPaletteRequest) GetStamps() []string`

GetStamps returns the Stamps field if non-nil, zero value otherwise.

### GetStampsOk

`func (o *PostStampPaletteRequest) GetStampsOk() (*[]string, bool)`

GetStampsOk returns a tuple with the Stamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStamps

`func (o *PostStampPaletteRequest) SetStamps(v []string)`

SetStamps sets Stamps field to given value.


### GetName

`func (o *PostStampPaletteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostStampPaletteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostStampPaletteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostStampPaletteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostStampPaletteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostStampPaletteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


