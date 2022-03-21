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

// Channel チャンネル
type Channel struct {
	// チャンネルUUID
	Id string `json:"id"`
	// 親チャンネルUUID
	ParentId NullableString `json:"parentId"`
	// チャンネルがアーカイブされているかどうか
	Archived bool `json:"archived"`
	// 強制通知チャンネルかどうか
	Force bool `json:"force"`
	// チャンネルトピック
	Topic string `json:"topic"`
	// チャンネル名
	Name string `json:"name"`
	// 子チャンネルのUUID配列
	Children []string `json:"children"`
}

// NewChannel instantiates a new Channel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannel(id string, parentId NullableString, archived bool, force bool, topic string, name string, children []string) *Channel {
	this := Channel{}
	this.Id = id
	this.ParentId = parentId
	this.Archived = archived
	this.Force = force
	this.Topic = topic
	this.Name = name
	this.Children = children
	return &this
}

// NewChannelWithDefaults instantiates a new Channel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelWithDefaults() *Channel {
	this := Channel{}
	return &this
}

// GetId returns the Id field value
func (o *Channel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Channel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Channel) SetId(v string) {
	o.Id = v
}

// GetParentId returns the ParentId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Channel) GetParentId() string {
	if o == nil || o.ParentId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Channel) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// SetParentId sets field value
func (o *Channel) SetParentId(v string) {
	o.ParentId.Set(&v)
}

// GetArchived returns the Archived field value
func (o *Channel) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *Channel) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *Channel) SetArchived(v bool) {
	o.Archived = v
}

// GetForce returns the Force field value
func (o *Channel) GetForce() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Force
}

// GetForceOk returns a tuple with the Force field value
// and a boolean to check if the value has been set.
func (o *Channel) GetForceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Force, true
}

// SetForce sets field value
func (o *Channel) SetForce(v bool) {
	o.Force = v
}

// GetTopic returns the Topic field value
func (o *Channel) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *Channel) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value
func (o *Channel) SetTopic(v string) {
	o.Topic = v
}

// GetName returns the Name field value
func (o *Channel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Channel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Channel) SetName(v string) {
	o.Name = v
}

// GetChildren returns the Children field value
func (o *Channel) GetChildren() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *Channel) GetChildrenOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *Channel) SetChildren(v []string) {
	o.Children = v
}

func (o Channel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if true {
		toSerialize["archived"] = o.Archived
	}
	if true {
		toSerialize["force"] = o.Force
	}
	if true {
		toSerialize["topic"] = o.Topic
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["children"] = o.Children
	}
	return json.Marshal(toSerialize)
}

type NullableChannel struct {
	value *Channel
	isSet bool
}

func (v NullableChannel) Get() *Channel {
	return v.value
}

func (v *NullableChannel) Set(val *Channel) {
	v.value = val
	v.isSet = true
}

func (v NullableChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannel(val *Channel) *NullableChannel {
	return &NullableChannel{value: val, isSet: true}
}

func (v NullableChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
