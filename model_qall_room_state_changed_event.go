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

// checks if the QallRoomStateChangedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QallRoomStateChangedEvent{}

// QallRoomStateChangedEvent Qallのルーム状態が変更された
type QallRoomStateChangedEvent struct {
	RoomStates []QallRoomStateChangedEventRoomStatesInner `json:"roomStates"`
}

// NewQallRoomStateChangedEvent instantiates a new QallRoomStateChangedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQallRoomStateChangedEvent(roomStates []QallRoomStateChangedEventRoomStatesInner) *QallRoomStateChangedEvent {
	this := QallRoomStateChangedEvent{}
	this.RoomStates = roomStates
	return &this
}

// NewQallRoomStateChangedEventWithDefaults instantiates a new QallRoomStateChangedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQallRoomStateChangedEventWithDefaults() *QallRoomStateChangedEvent {
	this := QallRoomStateChangedEvent{}
	return &this
}

// GetRoomStates returns the RoomStates field value
func (o *QallRoomStateChangedEvent) GetRoomStates() []QallRoomStateChangedEventRoomStatesInner {
	if o == nil {
		var ret []QallRoomStateChangedEventRoomStatesInner
		return ret
	}

	return o.RoomStates
}

// GetRoomStatesOk returns a tuple with the RoomStates field value
// and a boolean to check if the value has been set.
func (o *QallRoomStateChangedEvent) GetRoomStatesOk() ([]QallRoomStateChangedEventRoomStatesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomStates, true
}

// SetRoomStates sets field value
func (o *QallRoomStateChangedEvent) SetRoomStates(v []QallRoomStateChangedEventRoomStatesInner) {
	o.RoomStates = v
}

func (o QallRoomStateChangedEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QallRoomStateChangedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roomStates"] = o.RoomStates
	return toSerialize, nil
}

type NullableQallRoomStateChangedEvent struct {
	value *QallRoomStateChangedEvent
	isSet bool
}

func (v NullableQallRoomStateChangedEvent) Get() *QallRoomStateChangedEvent {
	return v.value
}

func (v *NullableQallRoomStateChangedEvent) Set(val *QallRoomStateChangedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableQallRoomStateChangedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableQallRoomStateChangedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQallRoomStateChangedEvent(val *QallRoomStateChangedEvent) *NullableQallRoomStateChangedEvent {
	return &NullableQallRoomStateChangedEvent{value: val, isSet: true}
}

func (v NullableQallRoomStateChangedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQallRoomStateChangedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
