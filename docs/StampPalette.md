# StampPalette

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | スタンプパレットUUID | 
**Name** | **string** | パレット名 | 
**Stamps** | **[]string** | パレット内のスタンプのUUID配列 | 
**CreatorId** | **string** | 作成者UUID | 
**CreatedAt** | **time.Time** | パレット作成日時 | 
**UpdatedAt** | **time.Time** | パレット更新日時 | 
**Description** | **string** | パレット説明 | 

## Methods

### NewStampPalette

`func NewStampPalette(id string, name string, stamps []string, creatorId string, createdAt time.Time, updatedAt time.Time, description string, ) *StampPalette`

NewStampPalette instantiates a new StampPalette object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStampPaletteWithDefaults

`func NewStampPaletteWithDefaults() *StampPalette`

NewStampPaletteWithDefaults instantiates a new StampPalette object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StampPalette) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StampPalette) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StampPalette) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StampPalette) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StampPalette) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StampPalette) SetName(v string)`

SetName sets Name field to given value.


### GetStamps

`func (o *StampPalette) GetStamps() []string`

GetStamps returns the Stamps field if non-nil, zero value otherwise.

### GetStampsOk

`func (o *StampPalette) GetStampsOk() (*[]string, bool)`

GetStampsOk returns a tuple with the Stamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStamps

`func (o *StampPalette) SetStamps(v []string)`

SetStamps sets Stamps field to given value.


### GetCreatorId

`func (o *StampPalette) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *StampPalette) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *StampPalette) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetCreatedAt

`func (o *StampPalette) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StampPalette) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StampPalette) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *StampPalette) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StampPalette) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StampPalette) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDescription

`func (o *StampPalette) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StampPalette) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StampPalette) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


