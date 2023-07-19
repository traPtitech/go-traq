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

// checks if the OgpMedia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OgpMedia{}

// OgpMedia OGPに含まれる画像の情報
type OgpMedia struct {
	Url       string         `json:"url"`
	SecureUrl NullableString `json:"secureUrl"`
	Type      NullableString `json:"type"`
	Width     NullableInt32  `json:"width"`
	Height    NullableInt32  `json:"height"`
}

// NewOgpMedia instantiates a new OgpMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOgpMedia(url string, secureUrl NullableString, type_ NullableString, width NullableInt32, height NullableInt32) *OgpMedia {
	this := OgpMedia{}
	this.Url = url
	this.SecureUrl = secureUrl
	this.Type = type_
	this.Width = width
	this.Height = height
	return &this
}

// NewOgpMediaWithDefaults instantiates a new OgpMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOgpMediaWithDefaults() *OgpMedia {
	this := OgpMedia{}
	return &this
}

// GetUrl returns the Url field value
func (o *OgpMedia) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OgpMedia) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *OgpMedia) SetUrl(v string) {
	o.Url = v
}

// GetSecureUrl returns the SecureUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OgpMedia) GetSecureUrl() string {
	if o == nil || o.SecureUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.SecureUrl.Get()
}

// GetSecureUrlOk returns a tuple with the SecureUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OgpMedia) GetSecureUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecureUrl.Get(), o.SecureUrl.IsSet()
}

// SetSecureUrl sets field value
func (o *OgpMedia) SetSecureUrl(v string) {
	o.SecureUrl.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OgpMedia) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OgpMedia) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *OgpMedia) SetType(v string) {
	o.Type.Set(&v)
}

// GetWidth returns the Width field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *OgpMedia) GetWidth() int32 {
	if o == nil || o.Width.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OgpMedia) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// SetWidth sets field value
func (o *OgpMedia) SetWidth(v int32) {
	o.Width.Set(&v)
}

// GetHeight returns the Height field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *OgpMedia) GetHeight() int32 {
	if o == nil || o.Height.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OgpMedia) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// SetHeight sets field value
func (o *OgpMedia) SetHeight(v int32) {
	o.Height.Set(&v)
}

func (o OgpMedia) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OgpMedia) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["secureUrl"] = o.SecureUrl.Get()
	toSerialize["type"] = o.Type.Get()
	toSerialize["width"] = o.Width.Get()
	toSerialize["height"] = o.Height.Get()
	return toSerialize, nil
}

type NullableOgpMedia struct {
	value *OgpMedia
	isSet bool
}

func (v NullableOgpMedia) Get() *OgpMedia {
	return v.value
}

func (v *NullableOgpMedia) Set(val *OgpMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableOgpMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableOgpMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOgpMedia(val *OgpMedia) *NullableOgpMedia {
	return &NullableOgpMedia{value: val, isSet: true}
}

func (v NullableOgpMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOgpMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
