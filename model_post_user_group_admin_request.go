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

// checks if the PostUserGroupAdminRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostUserGroupAdminRequest{}

// PostUserGroupAdminRequest グループ管理者追加リクエスト
type PostUserGroupAdminRequest struct {
	// 追加するユーザーのUUID
	Id string `json:"id"`
}

type _PostUserGroupAdminRequest PostUserGroupAdminRequest

// NewPostUserGroupAdminRequest instantiates a new PostUserGroupAdminRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostUserGroupAdminRequest(id string) *PostUserGroupAdminRequest {
	this := PostUserGroupAdminRequest{}
	this.Id = id
	return &this
}

// NewPostUserGroupAdminRequestWithDefaults instantiates a new PostUserGroupAdminRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostUserGroupAdminRequestWithDefaults() *PostUserGroupAdminRequest {
	this := PostUserGroupAdminRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PostUserGroupAdminRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostUserGroupAdminRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostUserGroupAdminRequest) SetId(v string) {
	o.Id = v
}

func (o PostUserGroupAdminRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostUserGroupAdminRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *PostUserGroupAdminRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPostUserGroupAdminRequest := _PostUserGroupAdminRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostUserGroupAdminRequest)

	if err != nil {
		return err
	}

	*o = PostUserGroupAdminRequest(varPostUserGroupAdminRequest)

	return err
}

type NullablePostUserGroupAdminRequest struct {
	value *PostUserGroupAdminRequest
	isSet bool
}

func (v NullablePostUserGroupAdminRequest) Get() *PostUserGroupAdminRequest {
	return v.value
}

func (v *NullablePostUserGroupAdminRequest) Set(val *PostUserGroupAdminRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostUserGroupAdminRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostUserGroupAdminRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostUserGroupAdminRequest(val *PostUserGroupAdminRequest) *NullablePostUserGroupAdminRequest {
	return &NullablePostUserGroupAdminRequest{value: val, isSet: true}
}

func (v NullablePostUserGroupAdminRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostUserGroupAdminRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
