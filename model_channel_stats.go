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

// ChannelStats チャンネル統計情報
type ChannelStats struct {
	// チャンネルの総投稿メッセージ数(削除されたものも含む)
	TotalMessageCount int64 `json:"totalMessageCount"`
	// チャンネル上のスタンプ統計情報
	Stamps []ChannelStatsStamp `json:"stamps"`
	// チャンネル上のユーザー統計情報
	Users []ChannelStatsUser `json:"users"`
	// 統計情報日時
	Datetime time.Time `json:"datetime"`
}

// NewChannelStats instantiates a new ChannelStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelStats(totalMessageCount int64, stamps []ChannelStatsStamp, users []ChannelStatsUser, datetime time.Time) *ChannelStats {
	this := ChannelStats{}
	this.TotalMessageCount = totalMessageCount
	this.Stamps = stamps
	this.Users = users
	this.Datetime = datetime
	return &this
}

// NewChannelStatsWithDefaults instantiates a new ChannelStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelStatsWithDefaults() *ChannelStats {
	this := ChannelStats{}
	return &this
}

// GetTotalMessageCount returns the TotalMessageCount field value
func (o *ChannelStats) GetTotalMessageCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalMessageCount
}

// GetTotalMessageCountOk returns a tuple with the TotalMessageCount field value
// and a boolean to check if the value has been set.
func (o *ChannelStats) GetTotalMessageCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMessageCount, true
}

// SetTotalMessageCount sets field value
func (o *ChannelStats) SetTotalMessageCount(v int64) {
	o.TotalMessageCount = v
}

// GetStamps returns the Stamps field value
func (o *ChannelStats) GetStamps() []ChannelStatsStamp {
	if o == nil {
		var ret []ChannelStatsStamp
		return ret
	}

	return o.Stamps
}

// GetStampsOk returns a tuple with the Stamps field value
// and a boolean to check if the value has been set.
func (o *ChannelStats) GetStampsOk() ([]ChannelStatsStamp, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stamps, true
}

// SetStamps sets field value
func (o *ChannelStats) SetStamps(v []ChannelStatsStamp) {
	o.Stamps = v
}

// GetUsers returns the Users field value
func (o *ChannelStats) GetUsers() []ChannelStatsUser {
	if o == nil {
		var ret []ChannelStatsUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ChannelStats) GetUsersOk() ([]ChannelStatsUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ChannelStats) SetUsers(v []ChannelStatsUser) {
	o.Users = v
}

// GetDatetime returns the Datetime field value
func (o *ChannelStats) GetDatetime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *ChannelStats) GetDatetimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *ChannelStats) SetDatetime(v time.Time) {
	o.Datetime = v
}

func (o ChannelStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalMessageCount"] = o.TotalMessageCount
	}
	if true {
		toSerialize["stamps"] = o.Stamps
	}
	if true {
		toSerialize["users"] = o.Users
	}
	if true {
		toSerialize["datetime"] = o.Datetime
	}
	return json.Marshal(toSerialize)
}

type NullableChannelStats struct {
	value *ChannelStats
	isSet bool
}

func (v NullableChannelStats) Get() *ChannelStats {
	return v.value
}

func (v *NullableChannelStats) Set(val *ChannelStats) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelStats) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelStats(val *ChannelStats) *NullableChannelStats {
	return &NullableChannelStats{value: val, isSet: true}
}

func (v NullableChannelStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
