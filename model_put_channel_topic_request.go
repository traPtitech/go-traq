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

// checks if the PutChannelTopicRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutChannelTopicRequest{}

// PutChannelTopicRequest チャンネルトピック編集リクエスト
type PutChannelTopicRequest struct {
	// トピック
	Topic string `json:"topic"`
}

type _PutChannelTopicRequest PutChannelTopicRequest

// NewPutChannelTopicRequest instantiates a new PutChannelTopicRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutChannelTopicRequest(topic string) *PutChannelTopicRequest {
	this := PutChannelTopicRequest{}
	this.Topic = topic
	return &this
}

// NewPutChannelTopicRequestWithDefaults instantiates a new PutChannelTopicRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutChannelTopicRequestWithDefaults() *PutChannelTopicRequest {
	this := PutChannelTopicRequest{}
	return &this
}

// GetTopic returns the Topic field value
func (o *PutChannelTopicRequest) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *PutChannelTopicRequest) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value
func (o *PutChannelTopicRequest) SetTopic(v string) {
	o.Topic = v
}

func (o PutChannelTopicRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutChannelTopicRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["topic"] = o.Topic
	return toSerialize, nil
}

func (o *PutChannelTopicRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"topic",
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

	varPutChannelTopicRequest := _PutChannelTopicRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPutChannelTopicRequest)

	if err != nil {
		return err
	}

	*o = PutChannelTopicRequest(varPutChannelTopicRequest)

	return err
}

type NullablePutChannelTopicRequest struct {
	value *PutChannelTopicRequest
	isSet bool
}

func (v NullablePutChannelTopicRequest) Get() *PutChannelTopicRequest {
	return v.value
}

func (v *NullablePutChannelTopicRequest) Set(val *PutChannelTopicRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutChannelTopicRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutChannelTopicRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutChannelTopicRequest(val *PutChannelTopicRequest) *NullablePutChannelTopicRequest {
	return &NullablePutChannelTopicRequest{value: val, isSet: true}
}

func (v NullablePutChannelTopicRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutChannelTopicRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
