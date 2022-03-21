/*
traQ v3

traQ v3 API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package traq

import (
	"encoding/json"
)

// FileInfoThumbnail サムネイル情報 サムネイルが存在しない場合はnullになります Deprecated: thumbnailsを参照してください
type FileInfoThumbnail struct {
	// MIMEタイプ
	// Deprecated
	Mime string `json:"mime"`
	// サムネイル幅
	// Deprecated
	Width *int32 `json:"width,omitempty"`
	// サムネイル高さ
	// Deprecated
	Height *int32 `json:"height,omitempty"`
}

// NewFileInfoThumbnail instantiates a new FileInfoThumbnail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileInfoThumbnail(mime string) *FileInfoThumbnail {
	this := FileInfoThumbnail{}
	this.Mime = mime
	return &this
}

// NewFileInfoThumbnailWithDefaults instantiates a new FileInfoThumbnail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileInfoThumbnailWithDefaults() *FileInfoThumbnail {
	this := FileInfoThumbnail{}
	return &this
}

// GetMime returns the Mime field value
// Deprecated
func (o *FileInfoThumbnail) GetMime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mime
}

// GetMimeOk returns a tuple with the Mime field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *FileInfoThumbnail) GetMimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mime, true
}

// SetMime sets field value
// Deprecated
func (o *FileInfoThumbnail) SetMime(v string) {
	o.Mime = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
// Deprecated
func (o *FileInfoThumbnail) GetWidth() int32 {
	if o == nil || o.Width == nil {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FileInfoThumbnail) GetWidthOk() (*int32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *FileInfoThumbnail) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
// Deprecated
func (o *FileInfoThumbnail) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
// Deprecated
func (o *FileInfoThumbnail) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FileInfoThumbnail) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *FileInfoThumbnail) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
// Deprecated
func (o *FileInfoThumbnail) SetHeight(v int32) {
	o.Height = &v
}

func (o FileInfoThumbnail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mime"] = o.Mime
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	return json.Marshal(toSerialize)
}

type NullableFileInfoThumbnail struct {
	value *FileInfoThumbnail
	isSet bool
}

func (v NullableFileInfoThumbnail) Get() *FileInfoThumbnail {
	return v.value
}

func (v *NullableFileInfoThumbnail) Set(val *FileInfoThumbnail) {
	v.value = val
	v.isSet = true
}

func (v NullableFileInfoThumbnail) IsSet() bool {
	return v.isSet
}

func (v *NullableFileInfoThumbnail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileInfoThumbnail(val *FileInfoThumbnail) *NullableFileInfoThumbnail {
	return &NullableFileInfoThumbnail{value: val, isSet: true}
}

func (v NullableFileInfoThumbnail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileInfoThumbnail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
