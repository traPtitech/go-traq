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

// BotMode BOT動作モード  HTTP: HTTP Mode WebSocket: WebSocket Mode
type BotMode string

// List of BotMode
const (
	BOTMODE_HTTP      BotMode = "HTTP"
	BOTMODE_WebSocket BotMode = "WebSocket"
)

// All allowed values of BotMode enum
var AllowedBotModeEnumValues = []BotMode{
	"HTTP",
	"WebSocket",
}

func (v *BotMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BotMode(value)
	for _, existing := range AllowedBotModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BotMode", value)
}

// NewBotModeFromValue returns a pointer to a valid BotMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBotModeFromValue(v string) (*BotMode, error) {
	ev := BotMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BotMode: valid values are %v", v, AllowedBotModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BotMode) IsValid() bool {
	for _, existing := range AllowedBotModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BotMode value
func (v BotMode) Ptr() *BotMode {
	return &v
}

type NullableBotMode struct {
	value *BotMode
	isSet bool
}

func (v NullableBotMode) Get() *BotMode {
	return v.value
}

func (v *NullableBotMode) Set(val *BotMode) {
	v.value = val
	v.isSet = true
}

func (v NullableBotMode) IsSet() bool {
	return v.isSet
}

func (v *NullableBotMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBotMode(val *BotMode) *NullableBotMode {
	return &NullableBotMode{value: val, isSet: true}
}

func (v NullableBotMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBotMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
