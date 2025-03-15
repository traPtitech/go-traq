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

// checks if the QallParticipantResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QallParticipantResponse{}

// QallParticipantResponse struct for QallParticipantResponse
type QallParticipantResponse struct {
	Results []QallParticipantResponseResultsInner `json:"results,omitempty"`
}

// NewQallParticipantResponse instantiates a new QallParticipantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQallParticipantResponse() *QallParticipantResponse {
	this := QallParticipantResponse{}
	return &this
}

// NewQallParticipantResponseWithDefaults instantiates a new QallParticipantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQallParticipantResponseWithDefaults() *QallParticipantResponse {
	this := QallParticipantResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *QallParticipantResponse) GetResults() []QallParticipantResponseResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []QallParticipantResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QallParticipantResponse) GetResultsOk() ([]QallParticipantResponseResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *QallParticipantResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []QallParticipantResponseResultsInner and assigns it to the Results field.
func (o *QallParticipantResponse) SetResults(v []QallParticipantResponseResultsInner) {
	o.Results = v
}

func (o QallParticipantResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QallParticipantResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableQallParticipantResponse struct {
	value *QallParticipantResponse
	isSet bool
}

func (v NullableQallParticipantResponse) Get() *QallParticipantResponse {
	return v.value
}

func (v *NullableQallParticipantResponse) Set(val *QallParticipantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQallParticipantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQallParticipantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQallParticipantResponse(val *QallParticipantResponse) *NullableQallParticipantResponse {
	return &NullableQallParticipantResponse{value: val, isSet: true}
}

func (v NullableQallParticipantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQallParticipantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
