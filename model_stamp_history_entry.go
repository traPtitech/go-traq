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

// checks if the StampHistoryEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StampHistoryEntry{}

// StampHistoryEntry スタンプ履歴の1項目
type StampHistoryEntry struct {
	// スタンプUUID
	StampId string `json:"stampId"`
	// 使用日時
	Datetime time.Time `json:"datetime"`
}

// NewStampHistoryEntry instantiates a new StampHistoryEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStampHistoryEntry(stampId string, datetime time.Time) *StampHistoryEntry {
	this := StampHistoryEntry{}
	this.StampId = stampId
	this.Datetime = datetime
	return &this
}

// NewStampHistoryEntryWithDefaults instantiates a new StampHistoryEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStampHistoryEntryWithDefaults() *StampHistoryEntry {
	this := StampHistoryEntry{}
	return &this
}

// GetStampId returns the StampId field value
func (o *StampHistoryEntry) GetStampId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StampId
}

// GetStampIdOk returns a tuple with the StampId field value
// and a boolean to check if the value has been set.
func (o *StampHistoryEntry) GetStampIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StampId, true
}

// SetStampId sets field value
func (o *StampHistoryEntry) SetStampId(v string) {
	o.StampId = v
}

// GetDatetime returns the Datetime field value
func (o *StampHistoryEntry) GetDatetime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *StampHistoryEntry) GetDatetimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *StampHistoryEntry) SetDatetime(v time.Time) {
	o.Datetime = v
}

func (o StampHistoryEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StampHistoryEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stampId"] = o.StampId
	toSerialize["datetime"] = o.Datetime
	return toSerialize, nil
}

type NullableStampHistoryEntry struct {
	value *StampHistoryEntry
	isSet bool
}

func (v NullableStampHistoryEntry) Get() *StampHistoryEntry {
	return v.value
}

func (v *NullableStampHistoryEntry) Set(val *StampHistoryEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableStampHistoryEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableStampHistoryEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStampHistoryEntry(val *StampHistoryEntry) *NullableStampHistoryEntry {
	return &NullableStampHistoryEntry{value: val, isSet: true}
}

func (v NullableStampHistoryEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStampHistoryEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
