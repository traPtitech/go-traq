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

// checks if the OAuth2Client type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2Client{}

// OAuth2Client OAuth2クライアント情報
type OAuth2Client struct {
	// クライアントUUID
	Id string `json:"id"`
	// クライアント名
	Name string `json:"name"`
	// 説明
	Description string `json:"description"`
	// クライアント開発者UUID
	DeveloperId string `json:"developerId"`
	// 要求スコープの配列
	Scopes []OAuth2Scope `json:"scopes"`
	// confidential client なら true, public client なら false
	Confidential bool `json:"confidential"`
}

// NewOAuth2Client instantiates a new OAuth2Client object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Client(id string, name string, description string, developerId string, scopes []OAuth2Scope, confidential bool) *OAuth2Client {
	this := OAuth2Client{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.DeveloperId = developerId
	this.Scopes = scopes
	this.Confidential = confidential
	return &this
}

// NewOAuth2ClientWithDefaults instantiates a new OAuth2Client object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientWithDefaults() *OAuth2Client {
	this := OAuth2Client{}
	return &this
}

// GetId returns the Id field value
func (o *OAuth2Client) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OAuth2Client) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OAuth2Client) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OAuth2Client) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *OAuth2Client) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OAuth2Client) SetDescription(v string) {
	o.Description = v
}

// GetDeveloperId returns the DeveloperId field value
func (o *OAuth2Client) GetDeveloperId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeveloperId
}

// GetDeveloperIdOk returns a tuple with the DeveloperId field value
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetDeveloperIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeveloperId, true
}

// SetDeveloperId sets field value
func (o *OAuth2Client) SetDeveloperId(v string) {
	o.DeveloperId = v
}

// GetScopes returns the Scopes field value
func (o *OAuth2Client) GetScopes() []OAuth2Scope {
	if o == nil {
		var ret []OAuth2Scope
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetScopesOk() ([]OAuth2Scope, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *OAuth2Client) SetScopes(v []OAuth2Scope) {
	o.Scopes = v
}

// GetConfidential returns the Confidential field value
func (o *OAuth2Client) GetConfidential() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Confidential
}

// GetConfidentialOk returns a tuple with the Confidential field value
// and a boolean to check if the value has been set.
func (o *OAuth2Client) GetConfidentialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidential, true
}

// SetConfidential sets field value
func (o *OAuth2Client) SetConfidential(v bool) {
	o.Confidential = v
}

func (o OAuth2Client) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2Client) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["developerId"] = o.DeveloperId
	toSerialize["scopes"] = o.Scopes
	toSerialize["confidential"] = o.Confidential
	return toSerialize, nil
}

type NullableOAuth2Client struct {
	value *OAuth2Client
	isSet bool
}

func (v NullableOAuth2Client) Get() *OAuth2Client {
	return v.value
}

func (v *NullableOAuth2Client) Set(val *OAuth2Client) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Client) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Client) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Client(val *OAuth2Client) *NullableOAuth2Client {
	return &NullableOAuth2Client{value: val, isSet: true}
}

func (v NullableOAuth2Client) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Client) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
