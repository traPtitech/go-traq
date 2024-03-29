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

// checks if the PatchClipFolderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchClipFolderRequest{}

// PatchClipFolderRequest クリップフォルダ情報編集リクエスト
type PatchClipFolderRequest struct {
	// フォルダ名
	Name *string `json:"name,omitempty"`
	// 説明
	Description *string `json:"description,omitempty"`
}

// NewPatchClipFolderRequest instantiates a new PatchClipFolderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchClipFolderRequest() *PatchClipFolderRequest {
	this := PatchClipFolderRequest{}
	return &this
}

// NewPatchClipFolderRequestWithDefaults instantiates a new PatchClipFolderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchClipFolderRequestWithDefaults() *PatchClipFolderRequest {
	this := PatchClipFolderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchClipFolderRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchClipFolderRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchClipFolderRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchClipFolderRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchClipFolderRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchClipFolderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchClipFolderRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchClipFolderRequest) SetDescription(v string) {
	o.Description = &v
}

func (o PatchClipFolderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchClipFolderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullablePatchClipFolderRequest struct {
	value *PatchClipFolderRequest
	isSet bool
}

func (v NullablePatchClipFolderRequest) Get() *PatchClipFolderRequest {
	return v.value
}

func (v *NullablePatchClipFolderRequest) Set(val *PatchClipFolderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchClipFolderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchClipFolderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchClipFolderRequest(val *PatchClipFolderRequest) *NullablePatchClipFolderRequest {
	return &NullablePatchClipFolderRequest{value: val, isSet: true}
}

func (v NullablePatchClipFolderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchClipFolderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
