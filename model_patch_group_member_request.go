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

// checks if the PatchGroupMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchGroupMemberRequest{}

// PatchGroupMemberRequest ユーザーグループメンバー編集リクエスト
type PatchGroupMemberRequest struct {
	// ユーザーの役割
	Role string `json:"role"`
}

// NewPatchGroupMemberRequest instantiates a new PatchGroupMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchGroupMemberRequest(role string) *PatchGroupMemberRequest {
	this := PatchGroupMemberRequest{}
	this.Role = role
	return &this
}

// NewPatchGroupMemberRequestWithDefaults instantiates a new PatchGroupMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchGroupMemberRequestWithDefaults() *PatchGroupMemberRequest {
	this := PatchGroupMemberRequest{}
	return &this
}

// GetRole returns the Role field value
func (o *PatchGroupMemberRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *PatchGroupMemberRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *PatchGroupMemberRequest) SetRole(v string) {
	o.Role = v
}

func (o PatchGroupMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchGroupMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

type NullablePatchGroupMemberRequest struct {
	value *PatchGroupMemberRequest
	isSet bool
}

func (v NullablePatchGroupMemberRequest) Get() *PatchGroupMemberRequest {
	return v.value
}

func (v *NullablePatchGroupMemberRequest) Set(val *PatchGroupMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchGroupMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchGroupMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchGroupMemberRequest(val *PatchGroupMemberRequest) *NullablePatchGroupMemberRequest {
	return &NullablePatchGroupMemberRequest{value: val, isSet: true}
}

func (v NullablePatchGroupMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchGroupMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
