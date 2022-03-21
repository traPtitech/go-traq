/*
traQ v3

traQ v3 API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package traq

import (
	"encoding/json"
	"time"
)

// ClipFolder クリップフォルダ情報
type ClipFolder struct {
	// フォルダUUID
	Id string `json:"id"`
	// フォルダ名
	Name string `json:"name"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// フォルダ所有者UUID
	OwnerId string `json:"ownerId"`
	// 説明
	Description string `json:"description"`
}

// NewClipFolder instantiates a new ClipFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClipFolder(id string, name string, createdAt time.Time, ownerId string, description string) *ClipFolder {
	this := ClipFolder{}
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.OwnerId = ownerId
	this.Description = description
	return &this
}

// NewClipFolderWithDefaults instantiates a new ClipFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClipFolderWithDefaults() *ClipFolder {
	this := ClipFolder{}
	return &this
}

// GetId returns the Id field value
func (o *ClipFolder) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClipFolder) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClipFolder) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ClipFolder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClipFolder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClipFolder) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ClipFolder) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ClipFolder) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ClipFolder) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetOwnerId returns the OwnerId field value
func (o *ClipFolder) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *ClipFolder) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *ClipFolder) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetDescription returns the Description field value
func (o *ClipFolder) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ClipFolder) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ClipFolder) SetDescription(v string) {
	o.Description = v
}

func (o ClipFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["ownerId"] = o.OwnerId
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableClipFolder struct {
	value *ClipFolder
	isSet bool
}

func (v NullableClipFolder) Get() *ClipFolder {
	return v.value
}

func (v *NullableClipFolder) Set(val *ClipFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableClipFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableClipFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClipFolder(val *ClipFolder) *NullableClipFolder {
	return &NullableClipFolder{value: val, isSet: true}
}

func (v NullableClipFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClipFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
