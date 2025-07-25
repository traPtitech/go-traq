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

// checks if the PostBotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostBotRequest{}

// PostBotRequest BOT作成リクエスト
type PostBotRequest struct {
	// BOTユーザーID 自動的に接頭辞\"BOT_\"が付与されます
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9_-]{1,16}$"`
	// BOTユーザー表示名
	DisplayName string `json:"displayName"`
	// BOTの説明
	Description string  `json:"description"`
	Mode        BotMode `json:"mode"`
	// BOTサーバーエンドポイント BOT動作モードがHTTPの場合必須です
	Endpoint *string `json:"endpoint,omitempty"`
}

type _PostBotRequest PostBotRequest

// NewPostBotRequest instantiates a new PostBotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostBotRequest(name string, displayName string, description string, mode BotMode) *PostBotRequest {
	this := PostBotRequest{}
	this.Name = name
	this.DisplayName = displayName
	this.Description = description
	this.Mode = mode
	return &this
}

// NewPostBotRequestWithDefaults instantiates a new PostBotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostBotRequestWithDefaults() *PostBotRequest {
	this := PostBotRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PostBotRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostBotRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostBotRequest) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *PostBotRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PostBotRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PostBotRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDescription returns the Description field value
func (o *PostBotRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PostBotRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PostBotRequest) SetDescription(v string) {
	o.Description = v
}

// GetMode returns the Mode field value
func (o *PostBotRequest) GetMode() BotMode {
	if o == nil {
		var ret BotMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *PostBotRequest) GetModeOk() (*BotMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *PostBotRequest) SetMode(v BotMode) {
	o.Mode = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PostBotRequest) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostBotRequest) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PostBotRequest) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *PostBotRequest) SetEndpoint(v string) {
	o.Endpoint = &v
}

func (o PostBotRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostBotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName
	toSerialize["description"] = o.Description
	toSerialize["mode"] = o.Mode
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	return toSerialize, nil
}

func (o *PostBotRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"displayName",
		"description",
		"mode",
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

	varPostBotRequest := _PostBotRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostBotRequest)

	if err != nil {
		return err
	}

	*o = PostBotRequest(varPostBotRequest)

	return err
}

type NullablePostBotRequest struct {
	value *PostBotRequest
	isSet bool
}

func (v NullablePostBotRequest) Get() *PostBotRequest {
	return v.value
}

func (v *NullablePostBotRequest) Set(val *PostBotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostBotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostBotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostBotRequest(val *PostBotRequest) *NullablePostBotRequest {
	return &NullablePostBotRequest{value: val, isSet: true}
}

func (v NullablePostBotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostBotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
