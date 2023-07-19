# LoginSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | セッションUUID | 
**IssuedAt** | **time.Time** | 発行日時 | 

## Methods

### NewLoginSession

`func NewLoginSession(id string, issuedAt time.Time, ) *LoginSession`

NewLoginSession instantiates a new LoginSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginSessionWithDefaults

`func NewLoginSessionWithDefaults() *LoginSession`

NewLoginSessionWithDefaults instantiates a new LoginSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoginSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoginSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoginSession) SetId(v string)`

SetId sets Id field to given value.


### GetIssuedAt

`func (o *LoginSession) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *LoginSession) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *LoginSession) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


