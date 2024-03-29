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

// UserAccountState ユーザーアカウント状態 0: 停止 1: 有効 2: 一時停止
type UserAccountState int32

// List of UserAccountState
const (
	USERACCOUNTSTATE_deactivated UserAccountState = 0
	USERACCOUNTSTATE_active      UserAccountState = 1
	USERACCOUNTSTATE_suspended   UserAccountState = 2
)

// All allowed values of UserAccountState enum
var AllowedUserAccountStateEnumValues = []UserAccountState{
	0,
	1,
	2,
}

func (v *UserAccountState) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserAccountState(value)
	for _, existing := range AllowedUserAccountStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserAccountState", value)
}

// NewUserAccountStateFromValue returns a pointer to a valid UserAccountState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserAccountStateFromValue(v int32) (*UserAccountState, error) {
	ev := UserAccountState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserAccountState: valid values are %v", v, AllowedUserAccountStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserAccountState) IsValid() bool {
	for _, existing := range AllowedUserAccountStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserAccountState value
func (v UserAccountState) Ptr() *UserAccountState {
	return &v
}

type NullableUserAccountState struct {
	value *UserAccountState
	isSet bool
}

func (v NullableUserAccountState) Get() *UserAccountState {
	return v.value
}

func (v *NullableUserAccountState) Set(val *UserAccountState) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccountState) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccountState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccountState(val *UserAccountState) *NullableUserAccountState {
	return &NullableUserAccountState{value: val, isSet: true}
}

func (v NullableUserAccountState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccountState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
