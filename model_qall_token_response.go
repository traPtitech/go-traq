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

// checks if the QallTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QallTokenResponse{}

// QallTokenResponse struct for QallTokenResponse
type QallTokenResponse struct {
	// LiveKit用のJWTトークン
	Token string `json:"token"`
}

// NewQallTokenResponse instantiates a new QallTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQallTokenResponse(token string) *QallTokenResponse {
	this := QallTokenResponse{}
	this.Token = token
	return &this
}

// NewQallTokenResponseWithDefaults instantiates a new QallTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQallTokenResponseWithDefaults() *QallTokenResponse {
	this := QallTokenResponse{}
	return &this
}

// GetToken returns the Token field value
func (o *QallTokenResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *QallTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *QallTokenResponse) SetToken(v string) {
	o.Token = v
}

func (o QallTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QallTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableQallTokenResponse struct {
	value *QallTokenResponse
	isSet bool
}

func (v NullableQallTokenResponse) Get() *QallTokenResponse {
	return v.value
}

func (v *NullableQallTokenResponse) Set(val *QallTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQallTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQallTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQallTokenResponse(val *QallTokenResponse) *NullableQallTokenResponse {
	return &NullableQallTokenResponse{value: val, isSet: true}
}

func (v NullableQallTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQallTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
