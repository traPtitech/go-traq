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

// checks if the ChannelPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPath{}

// ChannelPath チャンネルパス
type ChannelPath struct {
	// チャンネルパス
	Path string `json:"path"`
}

// NewChannelPath instantiates a new ChannelPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPath(path string) *ChannelPath {
	this := ChannelPath{}
	this.Path = path
	return &this
}

// NewChannelPathWithDefaults instantiates a new ChannelPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPathWithDefaults() *ChannelPath {
	this := ChannelPath{}
	return &this
}

// GetPath returns the Path field value
func (o *ChannelPath) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ChannelPath) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ChannelPath) SetPath(v string) {
	o.Path = v
}

func (o ChannelPath) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

type NullableChannelPath struct {
	value *ChannelPath
	isSet bool
}

func (v NullableChannelPath) Get() *ChannelPath {
	return v.value
}

func (v *NullableChannelPath) Set(val *ChannelPath) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPath) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPath(val *ChannelPath) *NullableChannelPath {
	return &NullableChannelPath{value: val, isSet: true}
}

func (v NullableChannelPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
