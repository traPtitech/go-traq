# VisibilityChangedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | 変更者UUID | 
**Visibility** | **bool** | 変更後可視状態 | 

## Methods

### NewVisibilityChangedEvent

`func NewVisibilityChangedEvent(userId string, visibility bool, ) *VisibilityChangedEvent`

NewVisibilityChangedEvent instantiates a new VisibilityChangedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisibilityChangedEventWithDefaults

`func NewVisibilityChangedEventWithDefaults() *VisibilityChangedEvent`

NewVisibilityChangedEventWithDefaults instantiates a new VisibilityChangedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *VisibilityChangedEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VisibilityChangedEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VisibilityChangedEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVisibility

`func (o *VisibilityChangedEvent) GetVisibility() bool`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *VisibilityChangedEvent) GetVisibilityOk() (*bool, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *VisibilityChangedEvent) SetVisibility(v bool)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


