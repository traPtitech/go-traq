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

// checks if the QallEndpointResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QallEndpointResponse{}

// QallEndpointResponse struct for QallEndpointResponse
type QallEndpointResponse struct {
	// LiveKitのエンドポイント
	Endpoint string `json:"endpoint"`
}

// NewQallEndpointResponse instantiates a new QallEndpointResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQallEndpointResponse(endpoint string) *QallEndpointResponse {
	this := QallEndpointResponse{}
	this.Endpoint = endpoint
	return &this
}

// NewQallEndpointResponseWithDefaults instantiates a new QallEndpointResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQallEndpointResponseWithDefaults() *QallEndpointResponse {
	this := QallEndpointResponse{}
	return &this
}

// GetEndpoint returns the Endpoint field value
func (o *QallEndpointResponse) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *QallEndpointResponse) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *QallEndpointResponse) SetEndpoint(v string) {
	o.Endpoint = v
}

func (o QallEndpointResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QallEndpointResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endpoint"] = o.Endpoint
	return toSerialize, nil
}

type NullableQallEndpointResponse struct {
	value *QallEndpointResponse
	isSet bool
}

func (v NullableQallEndpointResponse) Get() *QallEndpointResponse {
	return v.value
}

func (v *NullableQallEndpointResponse) Set(val *QallEndpointResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQallEndpointResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQallEndpointResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQallEndpointResponse(val *QallEndpointResponse) *NullableQallEndpointResponse {
	return &NullableQallEndpointResponse{value: val, isSet: true}
}

func (v NullableQallEndpointResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQallEndpointResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
