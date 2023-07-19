# ExternalProviderUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | 外部サービス名 | 
**LinkedAt** | **string** | 紐付けた日時 | 
**ExternalName** | **string** | 外部アカウント名 | 

## Methods

### NewExternalProviderUser

`func NewExternalProviderUser(providerName string, linkedAt string, externalName string, ) *ExternalProviderUser`

NewExternalProviderUser instantiates a new ExternalProviderUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalProviderUserWithDefaults

`func NewExternalProviderUserWithDefaults() *ExternalProviderUser`

NewExternalProviderUserWithDefaults instantiates a new ExternalProviderUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *ExternalProviderUser) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ExternalProviderUser) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ExternalProviderUser) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetLinkedAt

`func (o *ExternalProviderUser) GetLinkedAt() string`

GetLinkedAt returns the LinkedAt field if non-nil, zero value otherwise.

### GetLinkedAtOk

`func (o *ExternalProviderUser) GetLinkedAtOk() (*string, bool)`

GetLinkedAtOk returns a tuple with the LinkedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAt

`func (o *ExternalProviderUser) SetLinkedAt(v string)`

SetLinkedAt sets LinkedAt field to given value.


### GetExternalName

`func (o *ExternalProviderUser) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *ExternalProviderUser) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *ExternalProviderUser) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


