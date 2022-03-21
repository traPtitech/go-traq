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

// PostMessageStampRequest スタンプを押すリクエスト
type PostMessageStampRequest struct {
	// 押す数
	Count int32 `json:"count"`
}

// NewPostMessageStampRequest instantiates a new PostMessageStampRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostMessageStampRequest(count int32) *PostMessageStampRequest {
	this := PostMessageStampRequest{}
	this.Count = count
	return &this
}

// NewPostMessageStampRequestWithDefaults instantiates a new PostMessageStampRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostMessageStampRequestWithDefaults() *PostMessageStampRequest {
	this := PostMessageStampRequest{}
	return &this
}

// GetCount returns the Count field value
func (o *PostMessageStampRequest) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PostMessageStampRequest) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PostMessageStampRequest) SetCount(v int32) {
	o.Count = v
}

func (o PostMessageStampRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullablePostMessageStampRequest struct {
	value *PostMessageStampRequest
	isSet bool
}

func (v NullablePostMessageStampRequest) Get() *PostMessageStampRequest {
	return v.value
}

func (v *NullablePostMessageStampRequest) Set(val *PostMessageStampRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostMessageStampRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostMessageStampRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostMessageStampRequest(val *PostMessageStampRequest) *NullablePostMessageStampRequest {
	return &NullablePostMessageStampRequest{value: val, isSet: true}
}

func (v NullablePostMessageStampRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostMessageStampRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
