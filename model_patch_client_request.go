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

// checks if the PatchClientRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchClientRequest{}

// PatchClientRequest OAuth2クライアント情報変更リクエスト
type PatchClientRequest struct {
	// クライアント名
	Name *string `json:"name,omitempty"`
	// 説明
	Description *string `json:"description,omitempty"`
	// コールバックURL
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// クライアント開発者UUID
	DeveloperId *string `json:"developerId,omitempty"`
	// confidential client なら true, public client なら false
	Confidential *bool `json:"confidential,omitempty"`
}

// NewPatchClientRequest instantiates a new PatchClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchClientRequest() *PatchClientRequest {
	this := PatchClientRequest{}
	return &this
}

// NewPatchClientRequestWithDefaults instantiates a new PatchClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchClientRequestWithDefaults() *PatchClientRequest {
	this := PatchClientRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchClientRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchClientRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchClientRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchClientRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchClientRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchClientRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchClientRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchClientRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *PatchClientRequest) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchClientRequest) GetCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *PatchClientRequest) HasCallbackUrl() bool {
	if o != nil && !IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *PatchClientRequest) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetDeveloperId returns the DeveloperId field value if set, zero value otherwise.
func (o *PatchClientRequest) GetDeveloperId() string {
	if o == nil || IsNil(o.DeveloperId) {
		var ret string
		return ret
	}
	return *o.DeveloperId
}

// GetDeveloperIdOk returns a tuple with the DeveloperId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchClientRequest) GetDeveloperIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeveloperId) {
		return nil, false
	}
	return o.DeveloperId, true
}

// HasDeveloperId returns a boolean if a field has been set.
func (o *PatchClientRequest) HasDeveloperId() bool {
	if o != nil && !IsNil(o.DeveloperId) {
		return true
	}

	return false
}

// SetDeveloperId gets a reference to the given string and assigns it to the DeveloperId field.
func (o *PatchClientRequest) SetDeveloperId(v string) {
	o.DeveloperId = &v
}

// GetConfidential returns the Confidential field value if set, zero value otherwise.
func (o *PatchClientRequest) GetConfidential() bool {
	if o == nil || IsNil(o.Confidential) {
		var ret bool
		return ret
	}
	return *o.Confidential
}

// GetConfidentialOk returns a tuple with the Confidential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchClientRequest) GetConfidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Confidential) {
		return nil, false
	}
	return o.Confidential, true
}

// HasConfidential returns a boolean if a field has been set.
func (o *PatchClientRequest) HasConfidential() bool {
	if o != nil && !IsNil(o.Confidential) {
		return true
	}

	return false
}

// SetConfidential gets a reference to the given bool and assigns it to the Confidential field.
func (o *PatchClientRequest) SetConfidential(v bool) {
	o.Confidential = &v
}

func (o PatchClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchClientRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CallbackUrl) {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if !IsNil(o.DeveloperId) {
		toSerialize["developerId"] = o.DeveloperId
	}
	if !IsNil(o.Confidential) {
		toSerialize["confidential"] = o.Confidential
	}
	return toSerialize, nil
}

type NullablePatchClientRequest struct {
	value *PatchClientRequest
	isSet bool
}

func (v NullablePatchClientRequest) Get() *PatchClientRequest {
	return v.value
}

func (v *NullablePatchClientRequest) Set(val *PatchClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchClientRequest(val *PatchClientRequest) *NullablePatchClientRequest {
	return &NullablePatchClientRequest{value: val, isSet: true}
}

func (v NullablePatchClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
