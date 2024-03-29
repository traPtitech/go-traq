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

// checks if the ChildCreatedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChildCreatedEvent{}

// ChildCreatedEvent 子チャンネル作成イベント
type ChildCreatedEvent struct {
	// 作成者UUID
	UserId string `json:"userId"`
	// チャンネルUUID
	ChannelId string `json:"channelId"`
}

// NewChildCreatedEvent instantiates a new ChildCreatedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChildCreatedEvent(userId string, channelId string) *ChildCreatedEvent {
	this := ChildCreatedEvent{}
	this.UserId = userId
	this.ChannelId = channelId
	return &this
}

// NewChildCreatedEventWithDefaults instantiates a new ChildCreatedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChildCreatedEventWithDefaults() *ChildCreatedEvent {
	this := ChildCreatedEvent{}
	return &this
}

// GetUserId returns the UserId field value
func (o *ChildCreatedEvent) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ChildCreatedEvent) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ChildCreatedEvent) SetUserId(v string) {
	o.UserId = v
}

// GetChannelId returns the ChannelId field value
func (o *ChildCreatedEvent) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *ChildCreatedEvent) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *ChildCreatedEvent) SetChannelId(v string) {
	o.ChannelId = v
}

func (o ChildCreatedEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChildCreatedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["channelId"] = o.ChannelId
	return toSerialize, nil
}

type NullableChildCreatedEvent struct {
	value *ChildCreatedEvent
	isSet bool
}

func (v NullableChildCreatedEvent) Get() *ChildCreatedEvent {
	return v.value
}

func (v *NullableChildCreatedEvent) Set(val *ChildCreatedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableChildCreatedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableChildCreatedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChildCreatedEvent(val *ChildCreatedEvent) *NullableChildCreatedEvent {
	return &NullableChildCreatedEvent{value: val, isSet: true}
}

func (v NullableChildCreatedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChildCreatedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
