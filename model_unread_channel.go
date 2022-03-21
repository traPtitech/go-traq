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

// UnreadChannel 未読チャンネル情報
type UnreadChannel struct {
	// チャンネルUUID
	ChannelId string `json:"channelId"`
	// 未読メッセージ数
	Count int32 `json:"count"`
	// 自分宛てメッセージが含まれているかどうか
	Noticeable bool `json:"noticeable"`
	// チャンネルの最古の未読メッセージの日時
	Since time.Time `json:"since"`
	// チャンネルの最新の未読メッセージの日時
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewUnreadChannel instantiates a new UnreadChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnreadChannel(channelId string, count int32, noticeable bool, since time.Time, updatedAt time.Time) *UnreadChannel {
	this := UnreadChannel{}
	this.ChannelId = channelId
	this.Count = count
	this.Noticeable = noticeable
	this.Since = since
	this.UpdatedAt = updatedAt
	return &this
}

// NewUnreadChannelWithDefaults instantiates a new UnreadChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnreadChannelWithDefaults() *UnreadChannel {
	this := UnreadChannel{}
	return &this
}

// GetChannelId returns the ChannelId field value
func (o *UnreadChannel) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *UnreadChannel) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *UnreadChannel) SetChannelId(v string) {
	o.ChannelId = v
}

// GetCount returns the Count field value
func (o *UnreadChannel) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *UnreadChannel) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *UnreadChannel) SetCount(v int32) {
	o.Count = v
}

// GetNoticeable returns the Noticeable field value
func (o *UnreadChannel) GetNoticeable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Noticeable
}

// GetNoticeableOk returns a tuple with the Noticeable field value
// and a boolean to check if the value has been set.
func (o *UnreadChannel) GetNoticeableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Noticeable, true
}

// SetNoticeable sets field value
func (o *UnreadChannel) SetNoticeable(v bool) {
	o.Noticeable = v
}

// GetSince returns the Since field value
func (o *UnreadChannel) GetSince() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Since
}

// GetSinceOk returns a tuple with the Since field value
// and a boolean to check if the value has been set.
func (o *UnreadChannel) GetSinceOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Since, true
}

// SetSince sets field value
func (o *UnreadChannel) SetSince(v time.Time) {
	o.Since = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UnreadChannel) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UnreadChannel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UnreadChannel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o UnreadChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["channelId"] = o.ChannelId
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["noticeable"] = o.Noticeable
	}
	if true {
		toSerialize["since"] = o.Since
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableUnreadChannel struct {
	value *UnreadChannel
	isSet bool
}

func (v NullableUnreadChannel) Get() *UnreadChannel {
	return v.value
}

func (v *NullableUnreadChannel) Set(val *UnreadChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableUnreadChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableUnreadChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnreadChannel(val *UnreadChannel) *NullableUnreadChannel {
	return &NullableUnreadChannel{value: val, isSet: true}
}

func (v NullableUnreadChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnreadChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
