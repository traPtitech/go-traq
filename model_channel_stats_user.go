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

// checks if the ChannelStatsUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelStatsUser{}

// ChannelStatsUser チャンネル上の特定ユーザー統計情報
type ChannelStatsUser struct {
	// ユーザーID
	Id string `json:"id"`
	// メッセージ数
	MessageCount int64 `json:"messageCount"`
}

// NewChannelStatsUser instantiates a new ChannelStatsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelStatsUser(id string, messageCount int64) *ChannelStatsUser {
	this := ChannelStatsUser{}
	this.Id = id
	this.MessageCount = messageCount
	return &this
}

// NewChannelStatsUserWithDefaults instantiates a new ChannelStatsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelStatsUserWithDefaults() *ChannelStatsUser {
	this := ChannelStatsUser{}
	return &this
}

// GetId returns the Id field value
func (o *ChannelStatsUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChannelStatsUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChannelStatsUser) SetId(v string) {
	o.Id = v
}

// GetMessageCount returns the MessageCount field value
func (o *ChannelStatsUser) GetMessageCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value
// and a boolean to check if the value has been set.
func (o *ChannelStatsUser) GetMessageCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageCount, true
}

// SetMessageCount sets field value
func (o *ChannelStatsUser) SetMessageCount(v int64) {
	o.MessageCount = v
}

func (o ChannelStatsUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelStatsUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["messageCount"] = o.MessageCount
	return toSerialize, nil
}

type NullableChannelStatsUser struct {
	value *ChannelStatsUser
	isSet bool
}

func (v NullableChannelStatsUser) Get() *ChannelStatsUser {
	return v.value
}

func (v *NullableChannelStatsUser) Set(val *ChannelStatsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelStatsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelStatsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelStatsUser(val *ChannelStatsUser) *NullableChannelStatsUser {
	return &NullableChannelStatsUser{value: val, isSet: true}
}

func (v NullableChannelStatsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelStatsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
