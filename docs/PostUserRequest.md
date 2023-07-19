# PostUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | ユーザー名 | 
**Password** | Pointer to **string** | パスワード | [optional] 

## Methods

### NewPostUserRequest

`func NewPostUserRequest(name string, ) *PostUserRequest`

NewPostUserRequest instantiates a new PostUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostUserRequestWithDefaults

`func NewPostUserRequestWithDefaults() *PostUserRequest`

NewPostUserRequestWithDefaults instantiates a new PostUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *PostUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PostUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


