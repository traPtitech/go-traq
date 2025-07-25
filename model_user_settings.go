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

// checks if the UserSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSettings{}

// UserSettings ユーザー設定の情報
type UserSettings struct {
	// ユーザーUUID
	Id string `json:"id"`
	// メッセージ引用通知の設定情報
	NotifyCitation bool `json:"notifyCitation"`
}

type _UserSettings UserSettings

// NewUserSettings instantiates a new UserSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSettings(id string, notifyCitation bool) *UserSettings {
	this := UserSettings{}
	this.Id = id
	this.NotifyCitation = notifyCitation
	return &this
}

// NewUserSettingsWithDefaults instantiates a new UserSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingsWithDefaults() *UserSettings {
	this := UserSettings{}
	return &this
}

// GetId returns the Id field value
func (o *UserSettings) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSettings) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSettings) SetId(v string) {
	o.Id = v
}

// GetNotifyCitation returns the NotifyCitation field value
func (o *UserSettings) GetNotifyCitation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NotifyCitation
}

// GetNotifyCitationOk returns a tuple with the NotifyCitation field value
// and a boolean to check if the value has been set.
func (o *UserSettings) GetNotifyCitationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyCitation, true
}

// SetNotifyCitation sets field value
func (o *UserSettings) SetNotifyCitation(v bool) {
	o.NotifyCitation = v
}

func (o UserSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["notifyCitation"] = o.NotifyCitation
	return toSerialize, nil
}

func (o *UserSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUserSettings := _UserSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSettings)

	if err != nil {
		return err
	}

	*o = UserSettings(varUserSettings)

	return err
}

type NullableUserSettings struct {
	value *UserSettings
	isSet bool
}

func (v NullableUserSettings) Get() *UserSettings {
	return v.value
}

func (v *NullableUserSettings) Set(val *UserSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSettings(val *UserSettings) *NullableUserSettings {
	return &NullableUserSettings{value: val, isSet: true}
}

func (v NullableUserSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
