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

// PostUserRequest ユーザー登録リクエスト
type PostUserRequest struct {
	// ユーザー名
	Name string `json:"name"`
	// パスワード
	Password *string `json:"password,omitempty"`
}

// NewPostUserRequest instantiates a new PostUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostUserRequest(name string) *PostUserRequest {
	this := PostUserRequest{}
	this.Name = name
	return &this
}

// NewPostUserRequestWithDefaults instantiates a new PostUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostUserRequestWithDefaults() *PostUserRequest {
	this := PostUserRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PostUserRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostUserRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostUserRequest) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PostUserRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PostUserRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PostUserRequest) SetPassword(v string) {
	o.Password = &v
}

func (o PostUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullablePostUserRequest struct {
	value *PostUserRequest
	isSet bool
}

func (v NullablePostUserRequest) Get() *PostUserRequest {
	return v.value
}

func (v *NullablePostUserRequest) Set(val *PostUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostUserRequest(val *PostUserRequest) *NullablePostUserRequest {
	return &NullablePostUserRequest{value: val, isSet: true}
}

func (v NullablePostUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
