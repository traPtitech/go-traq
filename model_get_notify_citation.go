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

// checks if the GetNotifyCitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNotifyCitation{}

// GetNotifyCitation メッセージ引用通知の設定情報
type GetNotifyCitation struct {
	NotifyCitation bool `json:"notifyCitation"`
}

type _GetNotifyCitation GetNotifyCitation

// NewGetNotifyCitation instantiates a new GetNotifyCitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNotifyCitation(notifyCitation bool) *GetNotifyCitation {
	this := GetNotifyCitation{}
	this.NotifyCitation = notifyCitation
	return &this
}

// NewGetNotifyCitationWithDefaults instantiates a new GetNotifyCitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNotifyCitationWithDefaults() *GetNotifyCitation {
	this := GetNotifyCitation{}
	return &this
}

// GetNotifyCitation returns the NotifyCitation field value
func (o *GetNotifyCitation) GetNotifyCitation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NotifyCitation
}

// GetNotifyCitationOk returns a tuple with the NotifyCitation field value
// and a boolean to check if the value has been set.
func (o *GetNotifyCitation) GetNotifyCitationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyCitation, true
}

// SetNotifyCitation sets field value
func (o *GetNotifyCitation) SetNotifyCitation(v bool) {
	o.NotifyCitation = v
}

func (o GetNotifyCitation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNotifyCitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifyCitation"] = o.NotifyCitation
	return toSerialize, nil
}

func (o *GetNotifyCitation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notifyCitation",
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

	varGetNotifyCitation := _GetNotifyCitation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetNotifyCitation)

	if err != nil {
		return err
	}

	*o = GetNotifyCitation(varGetNotifyCitation)

	return err
}

type NullableGetNotifyCitation struct {
	value *GetNotifyCitation
	isSet bool
}

func (v NullableGetNotifyCitation) Get() *GetNotifyCitation {
	return v.value
}

func (v *NullableGetNotifyCitation) Set(val *GetNotifyCitation) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNotifyCitation) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNotifyCitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNotifyCitation(val *GetNotifyCitation) *NullableGetNotifyCitation {
	return &NullableGetNotifyCitation{value: val, isSet: true}
}

func (v NullableGetNotifyCitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNotifyCitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
