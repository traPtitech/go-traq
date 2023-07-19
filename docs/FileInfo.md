# FileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ファイルUUID | 
**Name** | **string** | ファイル名 | 
**Mime** | **string** | MIMEタイプ | 
**Size** | **int64** | ファイルサイズ | 
**Md5** | **string** | MD5ハッシュ | 
**IsAnimatedImage** | **bool** | アニメーション画像かどうか | 
**CreatedAt** | **time.Time** | アップロード日時 | 
**Thumbnails** | [**[]ThumbnailInfo**](ThumbnailInfo.md) |  | 
**Thumbnail** | [**NullableFileInfoThumbnail**](FileInfoThumbnail.md) |  | 
**ChannelId** | **NullableString** | 属しているチャンネルUUID | 
**UploaderId** | **NullableString** | アップロード者UUID | 

## Methods

### NewFileInfo

`func NewFileInfo(id string, name string, mime string, size int64, md5 string, isAnimatedImage bool, createdAt time.Time, thumbnails []ThumbnailInfo, thumbnail NullableFileInfoThumbnail, channelId NullableString, uploaderId NullableString, ) *FileInfo`

NewFileInfo instantiates a new FileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileInfoWithDefaults

`func NewFileInfoWithDefaults() *FileInfo`

NewFileInfoWithDefaults instantiates a new FileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FileInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileInfo) SetName(v string)`

SetName sets Name field to given value.


### GetMime

`func (o *FileInfo) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *FileInfo) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *FileInfo) SetMime(v string)`

SetMime sets Mime field to given value.


### GetSize

`func (o *FileInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileInfo) SetSize(v int64)`

SetSize sets Size field to given value.


### GetMd5

`func (o *FileInfo) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *FileInfo) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *FileInfo) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### GetIsAnimatedImage

`func (o *FileInfo) GetIsAnimatedImage() bool`

GetIsAnimatedImage returns the IsAnimatedImage field if non-nil, zero value otherwise.

### GetIsAnimatedImageOk

`func (o *FileInfo) GetIsAnimatedImageOk() (*bool, bool)`

GetIsAnimatedImageOk returns a tuple with the IsAnimatedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnimatedImage

`func (o *FileInfo) SetIsAnimatedImage(v bool)`

SetIsAnimatedImage sets IsAnimatedImage field to given value.


### GetCreatedAt

`func (o *FileInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetThumbnails

`func (o *FileInfo) GetThumbnails() []ThumbnailInfo`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *FileInfo) GetThumbnailsOk() (*[]ThumbnailInfo, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *FileInfo) SetThumbnails(v []ThumbnailInfo)`

SetThumbnails sets Thumbnails field to given value.


### GetThumbnail

`func (o *FileInfo) GetThumbnail() FileInfoThumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *FileInfo) GetThumbnailOk() (*FileInfoThumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *FileInfo) SetThumbnail(v FileInfoThumbnail)`

SetThumbnail sets Thumbnail field to given value.


### SetThumbnailNil

`func (o *FileInfo) SetThumbnailNil(b bool)`

 SetThumbnailNil sets the value for Thumbnail to be an explicit nil

### UnsetThumbnail
`func (o *FileInfo) UnsetThumbnail()`

UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
### GetChannelId

`func (o *FileInfo) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *FileInfo) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *FileInfo) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### SetChannelIdNil

`func (o *FileInfo) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *FileInfo) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetUploaderId

`func (o *FileInfo) GetUploaderId() string`

GetUploaderId returns the UploaderId field if non-nil, zero value otherwise.

### GetUploaderIdOk

`func (o *FileInfo) GetUploaderIdOk() (*string, bool)`

GetUploaderIdOk returns a tuple with the UploaderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderId

`func (o *FileInfo) SetUploaderId(v string)`

SetUploaderId sets UploaderId field to given value.


### SetUploaderIdNil

`func (o *FileInfo) SetUploaderIdNil(b bool)`

 SetUploaderIdNil sets the value for UploaderId to be an explicit nil

### UnsetUploaderId
`func (o *FileInfo) UnsetUploaderId()`

UnsetUploaderId ensures that no value is present for UploaderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


