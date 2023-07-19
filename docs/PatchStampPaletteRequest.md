# PatchStampPaletteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | パレット名 | [optional] 
**Description** | Pointer to **string** | 説明 | [optional] 
**Stamps** | Pointer to **[]string** | パレット内のスタンプUUIDの配列 | [optional] 

## Methods

### NewPatchStampPaletteRequest

`func NewPatchStampPaletteRequest() *PatchStampPaletteRequest`

NewPatchStampPaletteRequest instantiates a new PatchStampPaletteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchStampPaletteRequestWithDefaults

`func NewPatchStampPaletteRequestWithDefaults() *PatchStampPaletteRequest`

NewPatchStampPaletteRequestWithDefaults instantiates a new PatchStampPaletteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchStampPaletteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchStampPaletteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchStampPaletteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchStampPaletteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchStampPaletteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchStampPaletteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchStampPaletteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchStampPaletteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStamps

`func (o *PatchStampPaletteRequest) GetStamps() []string`

GetStamps returns the Stamps field if non-nil, zero value otherwise.

### GetStampsOk

`func (o *PatchStampPaletteRequest) GetStampsOk() (*[]string, bool)`

GetStampsOk returns a tuple with the Stamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStamps

`func (o *PatchStampPaletteRequest) SetStamps(v []string)`

SetStamps sets Stamps field to given value.

### HasStamps

`func (o *PatchStampPaletteRequest) HasStamps() bool`

HasStamps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


