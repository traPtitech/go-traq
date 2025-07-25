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

// checks if the PatchUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchUserRequest{}

// PatchUserRequest ユーザー情報編集リクエスト
type PatchUserRequest struct {
	// 新しい表示名
	DisplayName *string `json:"displayName,omitempty"`
	// TwitterID
	TwitterId *string           `json:"twitterId,omitempty" validate:"regexp=^[a-zA-Z0-9_]{1,15}$"`
	State     *UserAccountState `json:"state,omitempty"`
	// ユーザーロール
	Role *string `json:"role,omitempty"`
}

// NewPatchUserRequest instantiates a new PatchUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchUserRequest() *PatchUserRequest {
	this := PatchUserRequest{}
	return &this
}

// NewPatchUserRequestWithDefaults instantiates a new PatchUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchUserRequestWithDefaults() *PatchUserRequest {
	this := PatchUserRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PatchUserRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUserRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PatchUserRequest) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PatchUserRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetTwitterId returns the TwitterId field value if set, zero value otherwise.
func (o *PatchUserRequest) GetTwitterId() string {
	if o == nil || IsNil(o.TwitterId) {
		var ret string
		return ret
	}
	return *o.TwitterId
}

// GetTwitterIdOk returns a tuple with the TwitterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUserRequest) GetTwitterIdOk() (*string, bool) {
	if o == nil || IsNil(o.TwitterId) {
		return nil, false
	}
	return o.TwitterId, true
}

// HasTwitterId returns a boolean if a field has been set.
func (o *PatchUserRequest) HasTwitterId() bool {
	if o != nil && !IsNil(o.TwitterId) {
		return true
	}

	return false
}

// SetTwitterId gets a reference to the given string and assigns it to the TwitterId field.
func (o *PatchUserRequest) SetTwitterId(v string) {
	o.TwitterId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PatchUserRequest) GetState() UserAccountState {
	if o == nil || IsNil(o.State) {
		var ret UserAccountState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUserRequest) GetStateOk() (*UserAccountState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PatchUserRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given UserAccountState and assigns it to the State field.
func (o *PatchUserRequest) SetState(v UserAccountState) {
	o.State = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PatchUserRequest) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchUserRequest) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PatchUserRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *PatchUserRequest) SetRole(v string) {
	o.Role = &v
}

func (o PatchUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.TwitterId) {
		toSerialize["twitterId"] = o.TwitterId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullablePatchUserRequest struct {
	value *PatchUserRequest
	isSet bool
}

func (v NullablePatchUserRequest) Get() *PatchUserRequest {
	return v.value
}

func (v *NullablePatchUserRequest) Set(val *PatchUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchUserRequest(val *PatchUserRequest) *NullablePatchUserRequest {
	return &NullablePatchUserRequest{value: val, isSet: true}
}

func (v NullablePatchUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
