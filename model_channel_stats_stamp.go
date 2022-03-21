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

// ChannelStatsStamp チャンネル上の特定スタンプ統計情報
type ChannelStatsStamp struct {
	// スタンプID
	Id string `json:"id"`
	// スタンプ数(同一メッセージ上のものは複数カウントしない)
	Count int64 `json:"count"`
	// スタンプ数(同一メッセージ上のものも複数カウントする)
	Total int64 `json:"total"`
}

// NewChannelStatsStamp instantiates a new ChannelStatsStamp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelStatsStamp(id string, count int64, total int64) *ChannelStatsStamp {
	this := ChannelStatsStamp{}
	this.Id = id
	this.Count = count
	this.Total = total
	return &this
}

// NewChannelStatsStampWithDefaults instantiates a new ChannelStatsStamp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelStatsStampWithDefaults() *ChannelStatsStamp {
	this := ChannelStatsStamp{}
	return &this
}

// GetId returns the Id field value
func (o *ChannelStatsStamp) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChannelStatsStamp) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChannelStatsStamp) SetId(v string) {
	o.Id = v
}

// GetCount returns the Count field value
func (o *ChannelStatsStamp) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ChannelStatsStamp) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ChannelStatsStamp) SetCount(v int64) {
	o.Count = v
}

// GetTotal returns the Total field value
func (o *ChannelStatsStamp) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ChannelStatsStamp) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ChannelStatsStamp) SetTotal(v int64) {
	o.Total = v
}

func (o ChannelStatsStamp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableChannelStatsStamp struct {
	value *ChannelStatsStamp
	isSet bool
}

func (v NullableChannelStatsStamp) Get() *ChannelStatsStamp {
	return v.value
}

func (v *NullableChannelStatsStamp) Set(val *ChannelStatsStamp) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelStatsStamp) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelStatsStamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelStatsStamp(val *ChannelStatsStamp) *NullableChannelStatsStamp {
	return &NullableChannelStatsStamp{value: val, isSet: true}
}

func (v NullableChannelStatsStamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelStatsStamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
