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

// checks if the PostStarRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostStarRequest{}

// PostStarRequest スター追加リクエスト
type PostStarRequest struct {
	// チャンネルUUID
	ChannelId string `json:"channelId"`
}

// NewPostStarRequest instantiates a new PostStarRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostStarRequest(channelId string) *PostStarRequest {
	this := PostStarRequest{}
	this.ChannelId = channelId
	return &this
}

// NewPostStarRequestWithDefaults instantiates a new PostStarRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostStarRequestWithDefaults() *PostStarRequest {
	this := PostStarRequest{}
	return &this
}

// GetChannelId returns the ChannelId field value
func (o *PostStarRequest) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *PostStarRequest) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *PostStarRequest) SetChannelId(v string) {
	o.ChannelId = v
}

func (o PostStarRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostStarRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channelId"] = o.ChannelId
	return toSerialize, nil
}

type NullablePostStarRequest struct {
	value *PostStarRequest
	isSet bool
}

func (v NullablePostStarRequest) Get() *PostStarRequest {
	return v.value
}

func (v *NullablePostStarRequest) Set(val *PostStarRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostStarRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostStarRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostStarRequest(val *PostStarRequest) *NullablePostStarRequest {
	return &NullablePostStarRequest{value: val, isSet: true}
}

func (v NullablePostStarRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostStarRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
