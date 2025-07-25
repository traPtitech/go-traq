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

// checks if the SoundboardUploadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoundboardUploadResponse{}

// SoundboardUploadResponse struct for SoundboardUploadResponse
type SoundboardUploadResponse struct {
	// 登録されたサウンドID (ファイル名)
	SoundId string `json:"soundId"`
}

type _SoundboardUploadResponse SoundboardUploadResponse

// NewSoundboardUploadResponse instantiates a new SoundboardUploadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoundboardUploadResponse(soundId string) *SoundboardUploadResponse {
	this := SoundboardUploadResponse{}
	this.SoundId = soundId
	return &this
}

// NewSoundboardUploadResponseWithDefaults instantiates a new SoundboardUploadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoundboardUploadResponseWithDefaults() *SoundboardUploadResponse {
	this := SoundboardUploadResponse{}
	return &this
}

// GetSoundId returns the SoundId field value
func (o *SoundboardUploadResponse) GetSoundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SoundId
}

// GetSoundIdOk returns a tuple with the SoundId field value
// and a boolean to check if the value has been set.
func (o *SoundboardUploadResponse) GetSoundIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SoundId, true
}

// SetSoundId sets field value
func (o *SoundboardUploadResponse) SetSoundId(v string) {
	o.SoundId = v
}

func (o SoundboardUploadResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoundboardUploadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["soundId"] = o.SoundId
	return toSerialize, nil
}

func (o *SoundboardUploadResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"soundId",
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

	varSoundboardUploadResponse := _SoundboardUploadResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSoundboardUploadResponse)

	if err != nil {
		return err
	}

	*o = SoundboardUploadResponse(varSoundboardUploadResponse)

	return err
}

type NullableSoundboardUploadResponse struct {
	value *SoundboardUploadResponse
	isSet bool
}

func (v NullableSoundboardUploadResponse) Get() *SoundboardUploadResponse {
	return v.value
}

func (v *NullableSoundboardUploadResponse) Set(val *SoundboardUploadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSoundboardUploadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSoundboardUploadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoundboardUploadResponse(val *SoundboardUploadResponse) *NullableSoundboardUploadResponse {
	return &NullableSoundboardUploadResponse{value: val, isSet: true}
}

func (v NullableSoundboardUploadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoundboardUploadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
