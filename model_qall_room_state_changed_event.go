/*
traQ v3

traQ v3 API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package traq

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the QallRoomStateChangedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QallRoomStateChangedEvent{}

// QallRoomStateChangedEvent Qallのルーム状態が変更された
type QallRoomStateChangedEvent struct {
	RoomStates []QallRoomStateChangedEventRoomStatesInner `json:"roomStates"`
}

type _QallRoomStateChangedEvent QallRoomStateChangedEvent

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

func (o *QallRoomStateChangedEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roomStates",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varQallRoomStateChangedEvent := _QallRoomStateChangedEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQallRoomStateChangedEvent)

	if err != nil {
		return err
	}

	*o = QallRoomStateChangedEvent(varQallRoomStateChangedEvent)

	return err
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
