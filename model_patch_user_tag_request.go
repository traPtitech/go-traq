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

// checks if the PatchUserTagRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchUserTagRequest{}

// PatchUserTagRequest ユーザーのタグの編集リクエスト
type PatchUserTagRequest struct {
	// タグのロック状態
	IsLocked bool `json:"isLocked"`
}

// NewPatchUserTagRequest instantiates a new PatchUserTagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchUserTagRequest(isLocked bool) *PatchUserTagRequest {
	this := PatchUserTagRequest{}
	this.IsLocked = isLocked
	return &this
}

// NewPatchUserTagRequestWithDefaults instantiates a new PatchUserTagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchUserTagRequestWithDefaults() *PatchUserTagRequest {
	this := PatchUserTagRequest{}
	return &this
}

// GetIsLocked returns the IsLocked field value
func (o *PatchUserTagRequest) GetIsLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *PatchUserTagRequest) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *PatchUserTagRequest) SetIsLocked(v bool) {
	o.IsLocked = v
}

func (o PatchUserTagRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchUserTagRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isLocked"] = o.IsLocked
	return toSerialize, nil
}

type NullablePatchUserTagRequest struct {
	value *PatchUserTagRequest
	isSet bool
}

func (v NullablePatchUserTagRequest) Get() *PatchUserTagRequest {
	return v.value
}

func (v *NullablePatchUserTagRequest) Set(val *PatchUserTagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchUserTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchUserTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchUserTagRequest(val *PatchUserTagRequest) *NullablePatchUserTagRequest {
	return &NullablePatchUserTagRequest{value: val, isSet: true}
}

func (v NullablePatchUserTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchUserTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
