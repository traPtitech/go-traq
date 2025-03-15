# QallParticipantResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipantId** | Pointer to **string** | 対象参加者ID | [optional] 
**Status** | Pointer to **string** | success もしくは error | [optional] 
**ErrorMessage** | Pointer to **string** | エラーがある場合の詳細 | [optional] 

## Methods

### NewQallParticipantResponseResultsInner

`func NewQallParticipantResponseResultsInner() *QallParticipantResponseResultsInner`

NewQallParticipantResponseResultsInner instantiates a new QallParticipantResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQallParticipantResponseResultsInnerWithDefaults

`func NewQallParticipantResponseResultsInnerWithDefaults() *QallParticipantResponseResultsInner`

NewQallParticipantResponseResultsInnerWithDefaults instantiates a new QallParticipantResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipantId

`func (o *QallParticipantResponseResultsInner) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *QallParticipantResponseResultsInner) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *QallParticipantResponseResultsInner) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *QallParticipantResponseResultsInner) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### GetStatus

`func (o *QallParticipantResponseResultsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QallParticipantResponseResultsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QallParticipantResponseResultsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QallParticipantResponseResultsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *QallParticipantResponseResultsInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *QallParticipantResponseResultsInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *QallParticipantResponseResultsInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *QallParticipantResponseResultsInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


