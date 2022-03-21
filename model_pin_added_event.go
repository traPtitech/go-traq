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

// PinAddedEvent ピン追加イベント
type PinAddedEvent struct {
	// 変更者UUID
	UserId string `json:"userId"`
	// メッセージUUID
	MessageId string `json:"messageId"`
}

// NewPinAddedEvent instantiates a new PinAddedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinAddedEvent(userId string, messageId string) *PinAddedEvent {
	this := PinAddedEvent{}
	this.UserId = userId
	this.MessageId = messageId
	return &this
}

// NewPinAddedEventWithDefaults instantiates a new PinAddedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinAddedEventWithDefaults() *PinAddedEvent {
	this := PinAddedEvent{}
	return &this
}

// GetUserId returns the UserId field value
func (o *PinAddedEvent) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PinAddedEvent) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PinAddedEvent) SetUserId(v string) {
	o.UserId = v
}

// GetMessageId returns the MessageId field value
func (o *PinAddedEvent) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *PinAddedEvent) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *PinAddedEvent) SetMessageId(v string) {
	o.MessageId = v
}

func (o PinAddedEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["messageId"] = o.MessageId
	}
	return json.Marshal(toSerialize)
}

type NullablePinAddedEvent struct {
	value *PinAddedEvent
	isSet bool
}

func (v NullablePinAddedEvent) Get() *PinAddedEvent {
	return v.value
}

func (v *NullablePinAddedEvent) Set(val *PinAddedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullablePinAddedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullablePinAddedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePinAddedEvent(val *PinAddedEvent) *NullablePinAddedEvent {
	return &NullablePinAddedEvent{value: val, isSet: true}
}

func (v NullablePinAddedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePinAddedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
