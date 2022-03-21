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

// PatchStampRequest スタンプ情報変更リクエスト
type PatchStampRequest struct {
	// スタンプ名
	Name *string `json:"name,omitempty"`
	// 作成者UUID
	CreatorId *string `json:"creatorId,omitempty"`
}

// NewPatchStampRequest instantiates a new PatchStampRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchStampRequest() *PatchStampRequest {
	this := PatchStampRequest{}
	return &this
}

// NewPatchStampRequestWithDefaults instantiates a new PatchStampRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchStampRequestWithDefaults() *PatchStampRequest {
	this := PatchStampRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchStampRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStampRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchStampRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchStampRequest) SetName(v string) {
	o.Name = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *PatchStampRequest) GetCreatorId() string {
	if o == nil || o.CreatorId == nil {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStampRequest) GetCreatorIdOk() (*string, bool) {
	if o == nil || o.CreatorId == nil {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *PatchStampRequest) HasCreatorId() bool {
	if o != nil && o.CreatorId != nil {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *PatchStampRequest) SetCreatorId(v string) {
	o.CreatorId = &v
}

func (o PatchStampRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CreatorId != nil {
		toSerialize["creatorId"] = o.CreatorId
	}
	return json.Marshal(toSerialize)
}

type NullablePatchStampRequest struct {
	value *PatchStampRequest
	isSet bool
}

func (v NullablePatchStampRequest) Get() *PatchStampRequest {
	return v.value
}

func (v *NullablePatchStampRequest) Set(val *PatchStampRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchStampRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchStampRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchStampRequest(val *PatchStampRequest) *NullablePatchStampRequest {
	return &NullablePatchStampRequest{value: val, isSet: true}
}

func (v NullablePatchStampRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchStampRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
