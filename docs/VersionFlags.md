# VersionFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalLogin** | **[]string** | 有効な外部ログインプロバイダ | 
**SignUpAllowed** | **bool** | ユーザーが自身で新規登録(POST /api/v3/users)可能か | 

## Methods

### NewVersionFlags

`func NewVersionFlags(externalLogin []string, signUpAllowed bool, ) *VersionFlags`

NewVersionFlags instantiates a new VersionFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionFlagsWithDefaults

`func NewVersionFlagsWithDefaults() *VersionFlags`

NewVersionFlagsWithDefaults instantiates a new VersionFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalLogin

`func (o *VersionFlags) GetExternalLogin() []string`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *VersionFlags) GetExternalLoginOk() (*[]string, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *VersionFlags) SetExternalLogin(v []string)`

SetExternalLogin sets ExternalLogin field to given value.


### GetSignUpAllowed

`func (o *VersionFlags) GetSignUpAllowed() bool`

GetSignUpAllowed returns the SignUpAllowed field if non-nil, zero value otherwise.

### GetSignUpAllowedOk

`func (o *VersionFlags) GetSignUpAllowedOk() (*bool, bool)`

GetSignUpAllowedOk returns a tuple with the SignUpAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpAllowed

`func (o *VersionFlags) SetSignUpAllowed(v bool)`

SetSignUpAllowed sets SignUpAllowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


