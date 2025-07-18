/*
traQ v3

traQ v3 API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package traq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the FileInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileInfo{}

// FileInfo ファイル情報
type FileInfo struct {
	// ファイルUUID
	Id string `json:"id"`
	// ファイル名
	Name string `json:"name"`
	// MIMEタイプ
	Mime string `json:"mime"`
	// ファイルサイズ
	Size int64 `json:"size"`
	// MD5ハッシュ
	Md5 string `json:"md5"`
	// アニメーション画像かどうか
	IsAnimatedImage bool `json:"isAnimatedImage"`
	// アップロード日時
	CreatedAt  time.Time       `json:"createdAt"`
	Thumbnails []ThumbnailInfo `json:"thumbnails"`
	// Deprecated
	Thumbnail NullableFileInfoThumbnail `json:"thumbnail"`
	// 属しているチャンネルUUID
	ChannelId NullableString `json:"channelId"`
	// アップロード者UUID
	UploaderId NullableString `json:"uploaderId"`
}

type _FileInfo FileInfo

// NewFileInfo instantiates a new FileInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileInfo(id string, name string, mime string, size int64, md5 string, isAnimatedImage bool, createdAt time.Time, thumbnails []ThumbnailInfo, thumbnail NullableFileInfoThumbnail, channelId NullableString, uploaderId NullableString) *FileInfo {
	this := FileInfo{}
	this.Id = id
	this.Name = name
	this.Mime = mime
	this.Size = size
	this.Md5 = md5
	this.IsAnimatedImage = isAnimatedImage
	this.CreatedAt = createdAt
	this.Thumbnails = thumbnails
	this.Thumbnail = thumbnail
	this.ChannelId = channelId
	this.UploaderId = uploaderId
	return &this
}

// NewFileInfoWithDefaults instantiates a new FileInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileInfoWithDefaults() *FileInfo {
	this := FileInfo{}
	return &this
}

// GetId returns the Id field value
func (o *FileInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileInfo) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FileInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FileInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FileInfo) SetName(v string) {
	o.Name = v
}

// GetMime returns the Mime field value
func (o *FileInfo) GetMime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mime
}

// GetMimeOk returns a tuple with the Mime field value
// and a boolean to check if the value has been set.
func (o *FileInfo) GetMimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mime, true
}

// SetMime sets field value
func (o *FileInfo) SetMime(v string) {
	o.Mime = v
}

// GetSize returns the Size field value
func (o *FileInfo) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *FileInfo) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *FileInfo) SetSize(v int64) {
	o.Size = v
}

// GetMd5 returns the Md5 field value
func (o *FileInfo) GetMd5() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value
// and a boolean to check if the value has been set.
func (o *FileInfo) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Md5, true
}

// SetMd5 sets field value
func (o *FileInfo) SetMd5(v string) {
	o.Md5 = v
}

// GetIsAnimatedImage returns the IsAnimatedImage field value
func (o *FileInfo) GetIsAnimatedImage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAnimatedImage
}

// GetIsAnimatedImageOk returns a tuple with the IsAnimatedImage field value
// and a boolean to check if the value has been set.
func (o *FileInfo) GetIsAnimatedImageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAnimatedImage, true
}

// SetIsAnimatedImage sets field value
func (o *FileInfo) SetIsAnimatedImage(v bool) {
	o.IsAnimatedImage = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *FileInfo) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FileInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FileInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetThumbnails returns the Thumbnails field value
func (o *FileInfo) GetThumbnails() []ThumbnailInfo {
	if o == nil {
		var ret []ThumbnailInfo
		return ret
	}

	return o.Thumbnails
}

// GetThumbnailsOk returns a tuple with the Thumbnails field value
// and a boolean to check if the value has been set.
func (o *FileInfo) GetThumbnailsOk() ([]ThumbnailInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thumbnails, true
}

// SetThumbnails sets field value
func (o *FileInfo) SetThumbnails(v []ThumbnailInfo) {
	o.Thumbnails = v
}

// GetThumbnail returns the Thumbnail field value
// If the value is explicit nil, the zero value for FileInfoThumbnail will be returned
// Deprecated
func (o *FileInfo) GetThumbnail() FileInfoThumbnail {
	if o == nil || o.Thumbnail.Get() == nil {
		var ret FileInfoThumbnail
		return ret
	}

	return *o.Thumbnail.Get()
}

// GetThumbnailOk returns a tuple with the Thumbnail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *FileInfo) GetThumbnailOk() (*FileInfoThumbnail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thumbnail.Get(), o.Thumbnail.IsSet()
}

// SetThumbnail sets field value
// Deprecated
func (o *FileInfo) SetThumbnail(v FileInfoThumbnail) {
	o.Thumbnail.Set(&v)
}

// GetChannelId returns the ChannelId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FileInfo) GetChannelId() string {
	if o == nil || o.ChannelId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ChannelId.Get()
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileInfo) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelId.Get(), o.ChannelId.IsSet()
}

// SetChannelId sets field value
func (o *FileInfo) SetChannelId(v string) {
	o.ChannelId.Set(&v)
}

// GetUploaderId returns the UploaderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FileInfo) GetUploaderId() string {
	if o == nil || o.UploaderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.UploaderId.Get()
}

// GetUploaderIdOk returns a tuple with the UploaderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileInfo) GetUploaderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploaderId.Get(), o.UploaderId.IsSet()
}

// SetUploaderId sets field value
func (o *FileInfo) SetUploaderId(v string) {
	o.UploaderId.Set(&v)
}

func (o FileInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["mime"] = o.Mime
	toSerialize["size"] = o.Size
	toSerialize["md5"] = o.Md5
	toSerialize["isAnimatedImage"] = o.IsAnimatedImage
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["thumbnails"] = o.Thumbnails
	toSerialize["thumbnail"] = o.Thumbnail.Get()
	toSerialize["channelId"] = o.ChannelId.Get()
	toSerialize["uploaderId"] = o.UploaderId.Get()
	return toSerialize, nil
}

func (o *FileInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"mime",
		"size",
		"md5",
		"isAnimatedImage",
		"createdAt",
		"thumbnails",
		"thumbnail",
		"channelId",
		"uploaderId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFileInfo := _FileInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileInfo)

	if err != nil {
		return err
	}

	*o = FileInfo(varFileInfo)

	return err
}

type NullableFileInfo struct {
	value *FileInfo
	isSet bool
}

func (v NullableFileInfo) Get() *FileInfo {
	return v.value
}

func (v *NullableFileInfo) Set(val *FileInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFileInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFileInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileInfo(val *FileInfo) *NullableFileInfo {
	return &NullableFileInfo{value: val, isSet: true}
}

func (v NullableFileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
