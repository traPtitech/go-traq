# PatchClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | クライアント名 | [optional] 
**Description** | Pointer to **string** | 説明 | [optional] 
**CallbackUrl** | Pointer to **string** | コールバックURL | [optional] 
**DeveloperId** | Pointer to **string** | クライアント開発者UUID | [optional] 

## Methods

### NewPatchClientRequest

`func NewPatchClientRequest() *PatchClientRequest`

NewPatchClientRequest instantiates a new PatchClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchClientRequestWithDefaults

`func NewPatchClientRequestWithDefaults() *PatchClientRequest`

NewPatchClientRequestWithDefaults instantiates a new PatchClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchClientRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchClientRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchClientRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchClientRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchClientRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchClientRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchClientRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchClientRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *PatchClientRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *PatchClientRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *PatchClientRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *PatchClientRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetDeveloperId

`func (o *PatchClientRequest) GetDeveloperId() string`

GetDeveloperId returns the DeveloperId field if non-nil, zero value otherwise.

### GetDeveloperIdOk

`func (o *PatchClientRequest) GetDeveloperIdOk() (*string, bool)`

GetDeveloperIdOk returns a tuple with the DeveloperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperId

`func (o *PatchClientRequest) SetDeveloperId(v string)`

SetDeveloperId sets DeveloperId field to given value.

### HasDeveloperId

`func (o *PatchClientRequest) HasDeveloperId() bool`

HasDeveloperId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


