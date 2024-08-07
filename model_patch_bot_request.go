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

// checks if the PatchBotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchBotRequest{}

// PatchBotRequest BOT情報変更リクエスト
type PatchBotRequest struct {
	// BOTユーザー表示名
	DisplayName *string `json:"displayName,omitempty"`
	// BOTの説明
	Description *string `json:"description,omitempty"`
	// 特権
	Privileged *bool    `json:"privileged,omitempty"`
	Mode       *BotMode `json:"mode,omitempty"`
	// BOTサーバーエンドポイント
	Endpoint *string `json:"endpoint,omitempty"`
	// 移譲先の開発者UUID
	DeveloperId *string `json:"developerId,omitempty"`
	// 購読するイベント
	SubscribeEvents []string `json:"subscribeEvents,omitempty"`
	// 自己紹介(biography)
	Bio *string `json:"bio,omitempty"`
}

// NewPatchBotRequest instantiates a new PatchBotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchBotRequest() *PatchBotRequest {
	this := PatchBotRequest{}
	return &this
}

// NewPatchBotRequestWithDefaults instantiates a new PatchBotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchBotRequestWithDefaults() *PatchBotRequest {
	this := PatchBotRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PatchBotRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBotRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PatchBotRequest) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PatchBotRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchBotRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBotRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchBotRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchBotRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *PatchBotRequest) GetPrivileged() bool {
	if o == nil || IsNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBotRequest) GetPrivilegedOk() (*bool, bool) {
	if o == nil || IsNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *PatchBotRequest) HasPrivileged() bool {
	if o != nil && !IsNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *PatchBotRequest) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PatchBotRequest) GetMode() BotMode {
	if o == nil || IsNil(o.Mode) {
		var ret BotMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBotRequest) GetModeOk() (*BotMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PatchBotRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given BotMode and assigns it to the Mode field.
func (o *PatchBotRequest) SetMode(v BotMode) {
	o.Mode = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PatchBotRequest) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBotRequest) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PatchBotRequest) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *PatchBotRequest) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetDeveloperId returns the DeveloperId field value if set, zero value otherwise.
func (o *PatchBotRequest) GetDeveloperId() string {
	if o == nil || IsNil(o.DeveloperId) {
		var ret string
		return ret
	}
	return *o.DeveloperId
}

// GetDeveloperIdOk returns a tuple with the DeveloperId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBotRequest) GetDeveloperIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeveloperId) {
		return nil, false
	}
	return o.DeveloperId, true
}

// HasDeveloperId returns a boolean if a field has been set.
func (o *PatchBotRequest) HasDeveloperId() bool {
	if o != nil && !IsNil(o.DeveloperId) {
		return true
	}

	return false
}

// SetDeveloperId gets a reference to the given string and assigns it to the DeveloperId field.
func (o *PatchBotRequest) SetDeveloperId(v string) {
	o.DeveloperId = &v
}

// GetSubscribeEvents returns the SubscribeEvents field value if set, zero value otherwise.
func (o *PatchBotRequest) GetSubscribeEvents() []string {
	if o == nil || IsNil(o.SubscribeEvents) {
		var ret []string
		return ret
	}
	return o.SubscribeEvents
}

// GetSubscribeEventsOk returns a tuple with the SubscribeEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBotRequest) GetSubscribeEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubscribeEvents) {
		return nil, false
	}
	return o.SubscribeEvents, true
}

// HasSubscribeEvents returns a boolean if a field has been set.
func (o *PatchBotRequest) HasSubscribeEvents() bool {
	if o != nil && !IsNil(o.SubscribeEvents) {
		return true
	}

	return false
}

// SetSubscribeEvents gets a reference to the given []string and assigns it to the SubscribeEvents field.
func (o *PatchBotRequest) SetSubscribeEvents(v []string) {
	o.SubscribeEvents = v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *PatchBotRequest) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBotRequest) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *PatchBotRequest) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *PatchBotRequest) SetBio(v string) {
	o.Bio = &v
}

func (o PatchBotRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchBotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.DeveloperId) {
		toSerialize["developerId"] = o.DeveloperId
	}
	if !IsNil(o.SubscribeEvents) {
		toSerialize["subscribeEvents"] = o.SubscribeEvents
	}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	return toSerialize, nil
}

type NullablePatchBotRequest struct {
	value *PatchBotRequest
	isSet bool
}

func (v NullablePatchBotRequest) Get() *PatchBotRequest {
	return v.value
}

func (v *NullablePatchBotRequest) Set(val *PatchBotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBotRequest(val *PatchBotRequest) *NullablePatchBotRequest {
	return &NullablePatchBotRequest{value: val, isSet: true}
}

func (v NullablePatchBotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
