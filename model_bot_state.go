/*
traQ v3

traQ v3 API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package traq

import (
	"encoding/json"
	"fmt"
)

// BotState BOT状態 0: 停止 1: 有効 2: 一時停止
type BotState int32

// List of BotState
const (
	BOTSTATE_deactivated BotState = 0
	BOTSTATE_active      BotState = 1
	BOTSTATE_suspended   BotState = 2
)

// All allowed values of BotState enum
var AllowedBotStateEnumValues = []BotState{
	0,
	1,
	2,
}

func (v *BotState) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BotState(value)
	for _, existing := range AllowedBotStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BotState", value)
}

// NewBotStateFromValue returns a pointer to a valid BotState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBotStateFromValue(v int32) (*BotState, error) {
	ev := BotState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BotState: valid values are %v", v, AllowedBotStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BotState) IsValid() bool {
	for _, existing := range AllowedBotStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BotState value
func (v BotState) Ptr() *BotState {
	return &v
}

type NullableBotState struct {
	value *BotState
	isSet bool
}

func (v NullableBotState) Get() *BotState {
	return v.value
}

func (v *NullableBotState) Set(val *BotState) {
	v.value = val
	v.isSet = true
}

func (v NullableBotState) IsSet() bool {
	return v.isSet
}

func (v *NullableBotState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBotState(val *BotState) *NullableBotState {
	return &NullableBotState{value: val, isSet: true}
}

func (v NullableBotState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBotState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
