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

// ForcedNotificationChangedEvent チャンネル強制通知状態変更イベント
type ForcedNotificationChangedEvent struct {
	// 変更者UUID
	UserId string `json:"userId"`
	// 変更後強制通知状態
	Force bool `json:"force"`
}

// NewForcedNotificationChangedEvent instantiates a new ForcedNotificationChangedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForcedNotificationChangedEvent(userId string, force bool) *ForcedNotificationChangedEvent {
	this := ForcedNotificationChangedEvent{}
	this.UserId = userId
	this.Force = force
	return &this
}

// NewForcedNotificationChangedEventWithDefaults instantiates a new ForcedNotificationChangedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForcedNotificationChangedEventWithDefaults() *ForcedNotificationChangedEvent {
	this := ForcedNotificationChangedEvent{}
	return &this
}

// GetUserId returns the UserId field value
func (o *ForcedNotificationChangedEvent) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ForcedNotificationChangedEvent) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ForcedNotificationChangedEvent) SetUserId(v string) {
	o.UserId = v
}

// GetForce returns the Force field value
func (o *ForcedNotificationChangedEvent) GetForce() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Force
}

// GetForceOk returns a tuple with the Force field value
// and a boolean to check if the value has been set.
func (o *ForcedNotificationChangedEvent) GetForceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Force, true
}

// SetForce sets field value
func (o *ForcedNotificationChangedEvent) SetForce(v bool) {
	o.Force = v
}

func (o ForcedNotificationChangedEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableForcedNotificationChangedEvent struct {
	value *ForcedNotificationChangedEvent
	isSet bool
}

func (v NullableForcedNotificationChangedEvent) Get() *ForcedNotificationChangedEvent {
	return v.value
}

func (v *NullableForcedNotificationChangedEvent) Set(val *ForcedNotificationChangedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableForcedNotificationChangedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableForcedNotificationChangedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForcedNotificationChangedEvent(val *ForcedNotificationChangedEvent) *NullableForcedNotificationChangedEvent {
	return &NullableForcedNotificationChangedEvent{value: val, isSet: true}
}

func (v NullableForcedNotificationChangedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForcedNotificationChangedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
