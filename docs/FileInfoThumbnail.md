# FileInfoThumbnail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mime** | **string** | MIMEタイプ | 
**Width** | Pointer to **int32** | サムネイル幅 | [optional] 
**Height** | Pointer to **int32** | サムネイル高さ | [optional] 

## Methods

### NewFileInfoThumbnail

`func NewFileInfoThumbnail(mime string, ) *FileInfoThumbnail`

NewFileInfoThumbnail instantiates a new FileInfoThumbnail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileInfoThumbnailWithDefaults

`func NewFileInfoThumbnailWithDefaults() *FileInfoThumbnail`

NewFileInfoThumbnailWithDefaults instantiates a new FileInfoThumbnail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMime

`func (o *FileInfoThumbnail) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *FileInfoThumbnail) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *FileInfoThumbnail) SetMime(v string)`

SetMime sets Mime field to given value.


### GetWidth

`func (o *FileInfoThumbnail) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *FileInfoThumbnail) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *FileInfoThumbnail) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *FileInfoThumbnail) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *FileInfoThumbnail) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *FileInfoThumbnail) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *FileInfoThumbnail) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *FileInfoThumbnail) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


