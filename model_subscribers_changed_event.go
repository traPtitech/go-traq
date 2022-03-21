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

// SubscribersChangedEvent 購読者変更イベント
type SubscribersChangedEvent struct {
	// 変更者UUID
	UserId string `json:"userId"`
	// オンにされたユーザーのUUID配列
	On []string `json:"on"`
	// オフにされたユーザーのUUID配列
	Off []string `json:"off"`
}

// NewSubscribersChangedEvent instantiates a new SubscribersChangedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribersChangedEvent(userId string, on []string, off []string) *SubscribersChangedEvent {
	this := SubscribersChangedEvent{}
	this.UserId = userId
	this.On = on
	this.Off = off
	return &this
}

// NewSubscribersChangedEventWithDefaults instantiates a new SubscribersChangedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribersChangedEventWithDefaults() *SubscribersChangedEvent {
	this := SubscribersChangedEvent{}
	return &this
}

// GetUserId returns the UserId field value
func (o *SubscribersChangedEvent) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SubscribersChangedEvent) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SubscribersChangedEvent) SetUserId(v string) {
	o.UserId = v
}

// GetOn returns the On field value
func (o *SubscribersChangedEvent) GetOn() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.On
}

// GetOnOk returns a tuple with the On field value
// and a boolean to check if the value has been set.
func (o *SubscribersChangedEvent) GetOnOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.On, true
}

// SetOn sets field value
func (o *SubscribersChangedEvent) SetOn(v []string) {
	o.On = v
}

// GetOff returns the Off field value
func (o *SubscribersChangedEvent) GetOff() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Off
}

// GetOffOk returns a tuple with the Off field value
// and a boolean to check if the value has been set.
func (o *SubscribersChangedEvent) GetOffOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Off, true
}

// SetOff sets field value
func (o *SubscribersChangedEvent) SetOff(v []string) {
	o.Off = v
}

func (o SubscribersChangedEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["on"] = o.On
	}
	if true {
		toSerialize["off"] = o.Off
	}
	return json.Marshal(toSerialize)
}

type NullableSubscribersChangedEvent struct {
	value *SubscribersChangedEvent
	isSet bool
}

func (v NullableSubscribersChangedEvent) Get() *SubscribersChangedEvent {
	return v.value
}

func (v *NullableSubscribersChangedEvent) Set(val *SubscribersChangedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribersChangedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribersChangedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribersChangedEvent(val *SubscribersChangedEvent) *NullableSubscribersChangedEvent {
	return &NullableSubscribersChangedEvent{value: val, isSet: true}
}

func (v NullableSubscribersChangedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribersChangedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
