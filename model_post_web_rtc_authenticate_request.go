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

// checks if the PostWebRTCAuthenticateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostWebRTCAuthenticateRequest{}

// PostWebRTCAuthenticateRequest skyway用認証リクエスト
type PostWebRTCAuthenticateRequest struct {
	// ピアID
	PeerId string `json:"peerId"`
}

type _PostWebRTCAuthenticateRequest PostWebRTCAuthenticateRequest

// NewPostWebRTCAuthenticateRequest instantiates a new PostWebRTCAuthenticateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostWebRTCAuthenticateRequest(peerId string) *PostWebRTCAuthenticateRequest {
	this := PostWebRTCAuthenticateRequest{}
	this.PeerId = peerId
	return &this
}

// NewPostWebRTCAuthenticateRequestWithDefaults instantiates a new PostWebRTCAuthenticateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostWebRTCAuthenticateRequestWithDefaults() *PostWebRTCAuthenticateRequest {
	this := PostWebRTCAuthenticateRequest{}
	return &this
}

// GetPeerId returns the PeerId field value
func (o *PostWebRTCAuthenticateRequest) GetPeerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeerId
}

// GetPeerIdOk returns a tuple with the PeerId field value
// and a boolean to check if the value has been set.
func (o *PostWebRTCAuthenticateRequest) GetPeerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeerId, true
}

// SetPeerId sets field value
func (o *PostWebRTCAuthenticateRequest) SetPeerId(v string) {
	o.PeerId = v
}

func (o PostWebRTCAuthenticateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostWebRTCAuthenticateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["peerId"] = o.PeerId
	return toSerialize, nil
}

func (o *PostWebRTCAuthenticateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"peerId",
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

	varPostWebRTCAuthenticateRequest := _PostWebRTCAuthenticateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostWebRTCAuthenticateRequest)

	if err != nil {
		return err
	}

	*o = PostWebRTCAuthenticateRequest(varPostWebRTCAuthenticateRequest)

	return err
}

type NullablePostWebRTCAuthenticateRequest struct {
	value *PostWebRTCAuthenticateRequest
	isSet bool
}

func (v NullablePostWebRTCAuthenticateRequest) Get() *PostWebRTCAuthenticateRequest {
	return v.value
}

func (v *NullablePostWebRTCAuthenticateRequest) Set(val *PostWebRTCAuthenticateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostWebRTCAuthenticateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostWebRTCAuthenticateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostWebRTCAuthenticateRequest(val *PostWebRTCAuthenticateRequest) *NullablePostWebRTCAuthenticateRequest {
	return &NullablePostWebRTCAuthenticateRequest{value: val, isSet: true}
}

func (v NullablePostWebRTCAuthenticateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostWebRTCAuthenticateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
