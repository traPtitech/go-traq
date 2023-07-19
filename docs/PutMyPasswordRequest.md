# PutMyPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | 現在のパスワード | 
**NewPassword** | **string** | 新しいパスワード | 

## Methods

### NewPutMyPasswordRequest

`func NewPutMyPasswordRequest(password string, newPassword string, ) *PutMyPasswordRequest`

NewPutMyPasswordRequest instantiates a new PutMyPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutMyPasswordRequestWithDefaults

`func NewPutMyPasswordRequestWithDefaults() *PutMyPasswordRequest`

NewPutMyPasswordRequestWithDefaults instantiates a new PutMyPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PutMyPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PutMyPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PutMyPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetNewPassword

`func (o *PutMyPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *PutMyPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *PutMyPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


