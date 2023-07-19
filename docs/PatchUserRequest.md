# PatchUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | 新しい表示名 | [optional] 
**TwitterId** | Pointer to **string** | TwitterID | [optional] 
**State** | Pointer to [**UserAccountState**](UserAccountState.md) |  | [optional] 
**Role** | Pointer to **string** | ユーザーロール | [optional] 

## Methods

### NewPatchUserRequest

`func NewPatchUserRequest() *PatchUserRequest`

NewPatchUserRequest instantiates a new PatchUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchUserRequestWithDefaults

`func NewPatchUserRequestWithDefaults() *PatchUserRequest`

NewPatchUserRequestWithDefaults instantiates a new PatchUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PatchUserRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchUserRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchUserRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PatchUserRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTwitterId

`func (o *PatchUserRequest) GetTwitterId() string`

GetTwitterId returns the TwitterId field if non-nil, zero value otherwise.

### GetTwitterIdOk

`func (o *PatchUserRequest) GetTwitterIdOk() (*string, bool)`

GetTwitterIdOk returns a tuple with the TwitterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterId

`func (o *PatchUserRequest) SetTwitterId(v string)`

SetTwitterId sets TwitterId field to given value.

### HasTwitterId

`func (o *PatchUserRequest) HasTwitterId() bool`

HasTwitterId returns a boolean if a field has been set.

### GetState

`func (o *PatchUserRequest) GetState() UserAccountState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatchUserRequest) GetStateOk() (*UserAccountState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatchUserRequest) SetState(v UserAccountState)`

SetState sets State field to given value.

### HasState

`func (o *PatchUserRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRole

`func (o *PatchUserRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchUserRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchUserRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchUserRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


