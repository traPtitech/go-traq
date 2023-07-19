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

// checks if the PostClipFolderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostClipFolderRequest{}

// PostClipFolderRequest クリップフォルダ作成リクエスト
type PostClipFolderRequest struct {
	// フォルダ名
	Name string `json:"name"`
	// 説明
	Description string `json:"description"`
}

// NewPostClipFolderRequest instantiates a new PostClipFolderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostClipFolderRequest(name string, description string) *PostClipFolderRequest {
	this := PostClipFolderRequest{}
	this.Name = name
	this.Description = description
	return &this
}

// NewPostClipFolderRequestWithDefaults instantiates a new PostClipFolderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostClipFolderRequestWithDefaults() *PostClipFolderRequest {
	this := PostClipFolderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PostClipFolderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostClipFolderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostClipFolderRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *PostClipFolderRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PostClipFolderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PostClipFolderRequest) SetDescription(v string) {
	o.Description = v
}

func (o PostClipFolderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostClipFolderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullablePostClipFolderRequest struct {
	value *PostClipFolderRequest
	isSet bool
}

func (v NullablePostClipFolderRequest) Get() *PostClipFolderRequest {
	return v.value
}

func (v *NullablePostClipFolderRequest) Set(val *PostClipFolderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostClipFolderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostClipFolderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostClipFolderRequest(val *PostClipFolderRequest) *NullablePostClipFolderRequest {
	return &NullablePostClipFolderRequest{value: val, isSet: true}
}

func (v NullablePostClipFolderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostClipFolderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
