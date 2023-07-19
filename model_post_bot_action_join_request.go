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

// checks if the PostBotActionJoinRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostBotActionJoinRequest{}

// PostBotActionJoinRequest BOTチャンネル参加リクエスト
type PostBotActionJoinRequest struct {
	// チャンネルUUID
	ChannelId string `json:"channelId"`
}

// NewPostBotActionJoinRequest instantiates a new PostBotActionJoinRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostBotActionJoinRequest(channelId string) *PostBotActionJoinRequest {
	this := PostBotActionJoinRequest{}
	this.ChannelId = channelId
	return &this
}

// NewPostBotActionJoinRequestWithDefaults instantiates a new PostBotActionJoinRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostBotActionJoinRequestWithDefaults() *PostBotActionJoinRequest {
	this := PostBotActionJoinRequest{}
	return &this
}

// GetChannelId returns the ChannelId field value
func (o *PostBotActionJoinRequest) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *PostBotActionJoinRequest) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *PostBotActionJoinRequest) SetChannelId(v string) {
	o.ChannelId = v
}

func (o PostBotActionJoinRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostBotActionJoinRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channelId"] = o.ChannelId
	return toSerialize, nil
}

type NullablePostBotActionJoinRequest struct {
	value *PostBotActionJoinRequest
	isSet bool
}

func (v NullablePostBotActionJoinRequest) Get() *PostBotActionJoinRequest {
	return v.value
}

func (v *NullablePostBotActionJoinRequest) Set(val *PostBotActionJoinRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostBotActionJoinRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostBotActionJoinRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostBotActionJoinRequest(val *PostBotActionJoinRequest) *NullablePostBotActionJoinRequest {
	return &NullablePostBotActionJoinRequest{value: val, isSet: true}
}

func (v NullablePostBotActionJoinRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostBotActionJoinRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
