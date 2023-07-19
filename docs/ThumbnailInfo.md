# ThumbnailInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ThumbnailType**](ThumbnailType.md) |  | [default to THUMBNAILTYPE_IMAGE]
**Mime** | **string** | MIMEタイプ | 
**Width** | Pointer to **int32** | サムネイル幅 | [optional] 
**Height** | Pointer to **int32** | サムネイル高さ | [optional] 

## Methods

### NewThumbnailInfo

`func NewThumbnailInfo(type_ ThumbnailType, mime string, ) *ThumbnailInfo`

NewThumbnailInfo instantiates a new ThumbnailInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbnailInfoWithDefaults

`func NewThumbnailInfoWithDefaults() *ThumbnailInfo`

NewThumbnailInfoWithDefaults instantiates a new ThumbnailInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThumbnailInfo) GetType() ThumbnailType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThumbnailInfo) GetTypeOk() (*ThumbnailType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThumbnailInfo) SetType(v ThumbnailType)`

SetType sets Type field to given value.


### GetMime

`func (o *ThumbnailInfo) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *ThumbnailInfo) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *ThumbnailInfo) SetMime(v string)`

SetMime sets Mime field to given value.


### GetWidth

`func (o *ThumbnailInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ThumbnailInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ThumbnailInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ThumbnailInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ThumbnailInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ThumbnailInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ThumbnailInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ThumbnailInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


