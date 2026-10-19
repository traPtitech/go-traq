# PostThreadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | スレッド名 | [optional] 
**Parent** | Pointer to **string** | 親メッセージのUUID | [optional] 

## Methods

### NewPostThreadRequest

`func NewPostThreadRequest() *PostThreadRequest`

NewPostThreadRequest instantiates a new PostThreadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostThreadRequestWithDefaults

`func NewPostThreadRequestWithDefaults() *PostThreadRequest`

NewPostThreadRequestWithDefaults instantiates a new PostThreadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostThreadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostThreadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostThreadRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostThreadRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *PostThreadRequest) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PostThreadRequest) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PostThreadRequest) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PostThreadRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


