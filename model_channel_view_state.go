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

// ChannelViewState 閲覧状態
type ChannelViewState string

// List of ChannelViewState
const (
	CHANNELVIEWSTATE_NONE       ChannelViewState = "none"
	CHANNELVIEWSTATE_MONITORING ChannelViewState = "monitoring"
	CHANNELVIEWSTATE_EDITING    ChannelViewState = "editing"
)

// All allowed values of ChannelViewState enum
var AllowedChannelViewStateEnumValues = []ChannelViewState{
	"none",
	"monitoring",
	"editing",
}

func (v *ChannelViewState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelViewState(value)
	for _, existing := range AllowedChannelViewStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChannelViewState", value)
}

// NewChannelViewStateFromValue returns a pointer to a valid ChannelViewState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelViewStateFromValue(v string) (*ChannelViewState, error) {
	ev := ChannelViewState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelViewState: valid values are %v", v, AllowedChannelViewStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelViewState) IsValid() bool {
	for _, existing := range AllowedChannelViewStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChannelViewState value
func (v ChannelViewState) Ptr() *ChannelViewState {
	return &v
}

type NullableChannelViewState struct {
	value *ChannelViewState
	isSet bool
}

func (v NullableChannelViewState) Get() *ChannelViewState {
	return v.value
}

func (v *NullableChannelViewState) Set(val *ChannelViewState) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelViewState) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelViewState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelViewState(val *ChannelViewState) *NullableChannelViewState {
	return &NullableChannelViewState{value: val, isSet: true}
}

func (v NullableChannelViewState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelViewState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
