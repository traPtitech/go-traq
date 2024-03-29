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

// checks if the PostBotActionLeaveRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostBotActionLeaveRequest{}

// PostBotActionLeaveRequest BOTチャンネル退出リクエスト
type PostBotActionLeaveRequest struct {
	// チャンネルUUID
	ChannelId string `json:"channelId"`
}

// NewPostBotActionLeaveRequest instantiates a new PostBotActionLeaveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostBotActionLeaveRequest(channelId string) *PostBotActionLeaveRequest {
	this := PostBotActionLeaveRequest{}
	this.ChannelId = channelId
	return &this
}

// NewPostBotActionLeaveRequestWithDefaults instantiates a new PostBotActionLeaveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostBotActionLeaveRequestWithDefaults() *PostBotActionLeaveRequest {
	this := PostBotActionLeaveRequest{}
	return &this
}

// GetChannelId returns the ChannelId field value
func (o *PostBotActionLeaveRequest) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *PostBotActionLeaveRequest) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *PostBotActionLeaveRequest) SetChannelId(v string) {
	o.ChannelId = v
}

func (o PostBotActionLeaveRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostBotActionLeaveRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channelId"] = o.ChannelId
	return toSerialize, nil
}

type NullablePostBotActionLeaveRequest struct {
	value *PostBotActionLeaveRequest
	isSet bool
}

func (v NullablePostBotActionLeaveRequest) Get() *PostBotActionLeaveRequest {
	return v.value
}

func (v *NullablePostBotActionLeaveRequest) Set(val *PostBotActionLeaveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostBotActionLeaveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostBotActionLeaveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostBotActionLeaveRequest(val *PostBotActionLeaveRequest) *NullablePostBotActionLeaveRequest {
	return &NullablePostBotActionLeaveRequest{value: val, isSet: true}
}

func (v NullablePostBotActionLeaveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostBotActionLeaveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
