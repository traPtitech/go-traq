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
	"time"
)

// checks if the MessageClip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageClip{}

// MessageClip メッセージクリップ
type MessageClip struct {
	// クリップされているフォルダのID
	FolderId string `json:"folderId"`
	// クリップされた日時
	ClippedAt time.Time `json:"clippedAt"`
}

type _MessageClip MessageClip

// NewMessageClip instantiates a new MessageClip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageClip(folderId string, clippedAt time.Time) *MessageClip {
	this := MessageClip{}
	this.FolderId = folderId
	this.ClippedAt = clippedAt
	return &this
}

// NewMessageClipWithDefaults instantiates a new MessageClip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageClipWithDefaults() *MessageClip {
	this := MessageClip{}
	return &this
}

// GetFolderId returns the FolderId field value
func (o *MessageClip) GetFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *MessageClip) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *MessageClip) SetFolderId(v string) {
	o.FolderId = v
}

// GetClippedAt returns the ClippedAt field value
func (o *MessageClip) GetClippedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ClippedAt
}

// GetClippedAtOk returns a tuple with the ClippedAt field value
// and a boolean to check if the value has been set.
func (o *MessageClip) GetClippedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClippedAt, true
}

// SetClippedAt sets field value
func (o *MessageClip) SetClippedAt(v time.Time) {
	o.ClippedAt = v
}

func (o MessageClip) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageClip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folderId"] = o.FolderId
	toSerialize["clippedAt"] = o.ClippedAt
	return toSerialize, nil
}

func (o *MessageClip) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"folderId",
		"clippedAt",
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

	varMessageClip := _MessageClip{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageClip)

	if err != nil {
		return err
	}

	*o = MessageClip(varMessageClip)

	return err
}

type NullableMessageClip struct {
	value *MessageClip
	isSet bool
}

func (v NullableMessageClip) Get() *MessageClip {
	return v.value
}

func (v *NullableMessageClip) Set(val *MessageClip) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageClip) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageClip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageClip(val *MessageClip) *NullableMessageClip {
	return &NullableMessageClip{value: val, isSet: true}
}

func (v NullableMessageClip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageClip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
