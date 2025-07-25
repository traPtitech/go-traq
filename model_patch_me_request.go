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

// checks if the PatchMeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchMeRequest{}

// PatchMeRequest 自分のユーザー情報変更リクエスト
type PatchMeRequest struct {
	// 新しい表示名
	DisplayName *string `json:"displayName,omitempty"`
	// TwitterID
	TwitterId *string `json:"twitterId,omitempty" validate:"regexp=^[a-zA-Z0-9_]{1,15}$"`
	// 自己紹介(biography)
	Bio *string `json:"bio,omitempty"`
	// ホームチャンネルのUUID `00000000-0000-0000-0000-000000000000`を指定すると、ホームチャンネルが`null`に設定されます
	HomeChannel *string `json:"homeChannel,omitempty"`
}

// NewPatchMeRequest instantiates a new PatchMeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchMeRequest() *PatchMeRequest {
	this := PatchMeRequest{}
	return &this
}

// NewPatchMeRequestWithDefaults instantiates a new PatchMeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchMeRequestWithDefaults() *PatchMeRequest {
	this := PatchMeRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PatchMeRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMeRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PatchMeRequest) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PatchMeRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetTwitterId returns the TwitterId field value if set, zero value otherwise.
func (o *PatchMeRequest) GetTwitterId() string {
	if o == nil || IsNil(o.TwitterId) {
		var ret string
		return ret
	}
	return *o.TwitterId
}

// GetTwitterIdOk returns a tuple with the TwitterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMeRequest) GetTwitterIdOk() (*string, bool) {
	if o == nil || IsNil(o.TwitterId) {
		return nil, false
	}
	return o.TwitterId, true
}

// HasTwitterId returns a boolean if a field has been set.
func (o *PatchMeRequest) HasTwitterId() bool {
	if o != nil && !IsNil(o.TwitterId) {
		return true
	}

	return false
}

// SetTwitterId gets a reference to the given string and assigns it to the TwitterId field.
func (o *PatchMeRequest) SetTwitterId(v string) {
	o.TwitterId = &v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *PatchMeRequest) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMeRequest) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *PatchMeRequest) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *PatchMeRequest) SetBio(v string) {
	o.Bio = &v
}

// GetHomeChannel returns the HomeChannel field value if set, zero value otherwise.
func (o *PatchMeRequest) GetHomeChannel() string {
	if o == nil || IsNil(o.HomeChannel) {
		var ret string
		return ret
	}
	return *o.HomeChannel
}

// GetHomeChannelOk returns a tuple with the HomeChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMeRequest) GetHomeChannelOk() (*string, bool) {
	if o == nil || IsNil(o.HomeChannel) {
		return nil, false
	}
	return o.HomeChannel, true
}

// HasHomeChannel returns a boolean if a field has been set.
func (o *PatchMeRequest) HasHomeChannel() bool {
	if o != nil && !IsNil(o.HomeChannel) {
		return true
	}

	return false
}

// SetHomeChannel gets a reference to the given string and assigns it to the HomeChannel field.
func (o *PatchMeRequest) SetHomeChannel(v string) {
	o.HomeChannel = &v
}

func (o PatchMeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchMeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.TwitterId) {
		toSerialize["twitterId"] = o.TwitterId
	}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.HomeChannel) {
		toSerialize["homeChannel"] = o.HomeChannel
	}
	return toSerialize, nil
}

type NullablePatchMeRequest struct {
	value *PatchMeRequest
	isSet bool
}

func (v NullablePatchMeRequest) Get() *PatchMeRequest {
	return v.value
}

func (v *NullablePatchMeRequest) Set(val *PatchMeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchMeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchMeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchMeRequest(val *PatchMeRequest) *NullablePatchMeRequest {
	return &NullablePatchMeRequest{value: val, isSet: true}
}

func (v NullablePatchMeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchMeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
