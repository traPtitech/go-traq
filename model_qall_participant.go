/*
traQ v3

traQ v3 API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package traq

import (
	"encoding/json"
	"time"
)

// checks if the QallParticipant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QallParticipant{}

// QallParticipant ルーム内の参加者一覧
type QallParticipant struct {
	// ユーザーID_RandomUUID
	Identity *string `json:"identity,omitempty"`
	// 表示名
	Name *string `json:"name,omitempty"`
	// 参加した時刻
	JoinedAt *time.Time `json:"joinedAt,omitempty"`
	// ユーザーに関連付けられたカスタム属性
	Attributes *map[string]string `json:"attributes,omitempty"`
	// 発言権限
	CanPublish *bool `json:"canPublish,omitempty"`
}

// NewQallParticipant instantiates a new QallParticipant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQallParticipant() *QallParticipant {
	this := QallParticipant{}
	return &this
}

// NewQallParticipantWithDefaults instantiates a new QallParticipant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQallParticipantWithDefaults() *QallParticipant {
	this := QallParticipant{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *QallParticipant) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QallParticipant) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *QallParticipant) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *QallParticipant) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QallParticipant) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QallParticipant) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QallParticipant) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QallParticipant) SetName(v string) {
	o.Name = &v
}

// GetJoinedAt returns the JoinedAt field value if set, zero value otherwise.
func (o *QallParticipant) GetJoinedAt() time.Time {
	if o == nil || IsNil(o.JoinedAt) {
		var ret time.Time
		return ret
	}
	return *o.JoinedAt
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QallParticipant) GetJoinedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.JoinedAt) {
		return nil, false
	}
	return o.JoinedAt, true
}

// HasJoinedAt returns a boolean if a field has been set.
func (o *QallParticipant) HasJoinedAt() bool {
	if o != nil && !IsNil(o.JoinedAt) {
		return true
	}

	return false
}

// SetJoinedAt gets a reference to the given time.Time and assigns it to the JoinedAt field.
func (o *QallParticipant) SetJoinedAt(v time.Time) {
	o.JoinedAt = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *QallParticipant) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QallParticipant) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *QallParticipant) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *QallParticipant) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetCanPublish returns the CanPublish field value if set, zero value otherwise.
func (o *QallParticipant) GetCanPublish() bool {
	if o == nil || IsNil(o.CanPublish) {
		var ret bool
		return ret
	}
	return *o.CanPublish
}

// GetCanPublishOk returns a tuple with the CanPublish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QallParticipant) GetCanPublishOk() (*bool, bool) {
	if o == nil || IsNil(o.CanPublish) {
		return nil, false
	}
	return o.CanPublish, true
}

// HasCanPublish returns a boolean if a field has been set.
func (o *QallParticipant) HasCanPublish() bool {
	if o != nil && !IsNil(o.CanPublish) {
		return true
	}

	return false
}

// SetCanPublish gets a reference to the given bool and assigns it to the CanPublish field.
func (o *QallParticipant) SetCanPublish(v bool) {
	o.CanPublish = &v
}

func (o QallParticipant) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QallParticipant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.JoinedAt) {
		toSerialize["joinedAt"] = o.JoinedAt
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.CanPublish) {
		toSerialize["canPublish"] = o.CanPublish
	}
	return toSerialize, nil
}

type NullableQallParticipant struct {
	value *QallParticipant
	isSet bool
}

func (v NullableQallParticipant) Get() *QallParticipant {
	return v.value
}

func (v *NullableQallParticipant) Set(val *QallParticipant) {
	v.value = val
	v.isSet = true
}

func (v NullableQallParticipant) IsSet() bool {
	return v.isSet
}

func (v *NullableQallParticipant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQallParticipant(val *QallParticipant) *NullableQallParticipant {
	return &NullableQallParticipant{value: val, isSet: true}
}

func (v NullableQallParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQallParticipant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
