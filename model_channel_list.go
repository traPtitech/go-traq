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

// ChannelList GET /channelsレスポンス
type ChannelList struct {
	// パブリックチャンネルの配列
	Public []Channel `json:"public"`
	// ダイレクトメッセージチャンネルの配列
	Dm []DMChannel `json:"dm,omitempty"`
}

// NewChannelList instantiates a new ChannelList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelList(public []Channel) *ChannelList {
	this := ChannelList{}
	this.Public = public
	return &this
}

// NewChannelListWithDefaults instantiates a new ChannelList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelListWithDefaults() *ChannelList {
	this := ChannelList{}
	return &this
}

// GetPublic returns the Public field value
func (o *ChannelList) GetPublic() []Channel {
	if o == nil {
		var ret []Channel
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *ChannelList) GetPublicOk() ([]Channel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Public, true
}

// SetPublic sets field value
func (o *ChannelList) SetPublic(v []Channel) {
	o.Public = v
}

// GetDm returns the Dm field value if set, zero value otherwise.
func (o *ChannelList) GetDm() []DMChannel {
	if o == nil || o.Dm == nil {
		var ret []DMChannel
		return ret
	}
	return o.Dm
}

// GetDmOk returns a tuple with the Dm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelList) GetDmOk() ([]DMChannel, bool) {
	if o == nil || o.Dm == nil {
		return nil, false
	}
	return o.Dm, true
}

// HasDm returns a boolean if a field has been set.
func (o *ChannelList) HasDm() bool {
	if o != nil && o.Dm != nil {
		return true
	}

	return false
}

// SetDm gets a reference to the given []DMChannel and assigns it to the Dm field.
func (o *ChannelList) SetDm(v []DMChannel) {
	o.Dm = v
}

func (o ChannelList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["public"] = o.Public
	}
	if o.Dm != nil {
		toSerialize["dm"] = o.Dm
	}
	return json.Marshal(toSerialize)
}

type NullableChannelList struct {
	value *ChannelList
	isSet bool
}

func (v NullableChannelList) Get() *ChannelList {
	return v.value
}

func (v *NullableChannelList) Set(val *ChannelList) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelList) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelList(val *ChannelList) *NullableChannelList {
	return &NullableChannelList{value: val, isSet: true}
}

func (v NullableChannelList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
