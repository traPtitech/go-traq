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

// checks if the StampWithThumbnail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StampWithThumbnail{}

// StampWithThumbnail スタンプ情報とサムネイルの有無
type StampWithThumbnail struct {
	// スタンプUUID
	Id string `json:"id"`
	// スタンプ名
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9_-]{1,32}$"`
	// 作成者UUID
	CreatorId string `json:"creatorId"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
	// ファイルUUID
	FileId string `json:"fileId"`
	// Unicode絵文字か
	IsUnicode bool `json:"isUnicode"`
	// サムネイルの有無
	HasThumbnail bool `json:"hasThumbnail"`
}

type _StampWithThumbnail StampWithThumbnail

// NewStampWithThumbnail instantiates a new StampWithThumbnail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStampWithThumbnail(id string, name string, creatorId string, createdAt time.Time, updatedAt time.Time, fileId string, isUnicode bool, hasThumbnail bool) *StampWithThumbnail {
	this := StampWithThumbnail{}
	this.Id = id
	this.Name = name
	this.CreatorId = creatorId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.FileId = fileId
	this.IsUnicode = isUnicode
	this.HasThumbnail = hasThumbnail
	return &this
}

// NewStampWithThumbnailWithDefaults instantiates a new StampWithThumbnail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStampWithThumbnailWithDefaults() *StampWithThumbnail {
	this := StampWithThumbnail{}
	return &this
}

// GetId returns the Id field value
func (o *StampWithThumbnail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StampWithThumbnail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StampWithThumbnail) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *StampWithThumbnail) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StampWithThumbnail) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StampWithThumbnail) SetName(v string) {
	o.Name = v
}

// GetCreatorId returns the CreatorId field value
func (o *StampWithThumbnail) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *StampWithThumbnail) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *StampWithThumbnail) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *StampWithThumbnail) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *StampWithThumbnail) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *StampWithThumbnail) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *StampWithThumbnail) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *StampWithThumbnail) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *StampWithThumbnail) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetFileId returns the FileId field value
func (o *StampWithThumbnail) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *StampWithThumbnail) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *StampWithThumbnail) SetFileId(v string) {
	o.FileId = v
}

// GetIsUnicode returns the IsUnicode field value
func (o *StampWithThumbnail) GetIsUnicode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUnicode
}

// GetIsUnicodeOk returns a tuple with the IsUnicode field value
// and a boolean to check if the value has been set.
func (o *StampWithThumbnail) GetIsUnicodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUnicode, true
}

// SetIsUnicode sets field value
func (o *StampWithThumbnail) SetIsUnicode(v bool) {
	o.IsUnicode = v
}

// GetHasThumbnail returns the HasThumbnail field value
func (o *StampWithThumbnail) GetHasThumbnail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasThumbnail
}

// GetHasThumbnailOk returns a tuple with the HasThumbnail field value
// and a boolean to check if the value has been set.
func (o *StampWithThumbnail) GetHasThumbnailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasThumbnail, true
}

// SetHasThumbnail sets field value
func (o *StampWithThumbnail) SetHasThumbnail(v bool) {
	o.HasThumbnail = v
}

func (o StampWithThumbnail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StampWithThumbnail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["creatorId"] = o.CreatorId
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["fileId"] = o.FileId
	toSerialize["isUnicode"] = o.IsUnicode
	toSerialize["hasThumbnail"] = o.HasThumbnail
	return toSerialize, nil
}

func (o *StampWithThumbnail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"creatorId",
		"createdAt",
		"updatedAt",
		"fileId",
		"isUnicode",
		"hasThumbnail",
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

	varStampWithThumbnail := _StampWithThumbnail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStampWithThumbnail)

	if err != nil {
		return err
	}

	*o = StampWithThumbnail(varStampWithThumbnail)

	return err
}

type NullableStampWithThumbnail struct {
	value *StampWithThumbnail
	isSet bool
}

func (v NullableStampWithThumbnail) Get() *StampWithThumbnail {
	return v.value
}

func (v *NullableStampWithThumbnail) Set(val *StampWithThumbnail) {
	v.value = val
	v.isSet = true
}

func (v NullableStampWithThumbnail) IsSet() bool {
	return v.isSet
}

func (v *NullableStampWithThumbnail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStampWithThumbnail(val *StampWithThumbnail) *NullableStampWithThumbnail {
	return &NullableStampWithThumbnail{value: val, isSet: true}
}

func (v NullableStampWithThumbnail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStampWithThumbnail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
