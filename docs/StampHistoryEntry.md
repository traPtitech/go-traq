# StampHistoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StampId** | **string** | スタンプUUID | 
**Datetime** | **time.Time** | 使用日時 | 

## Methods

### NewStampHistoryEntry

`func NewStampHistoryEntry(stampId string, datetime time.Time, ) *StampHistoryEntry`

NewStampHistoryEntry instantiates a new StampHistoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStampHistoryEntryWithDefaults

`func NewStampHistoryEntryWithDefaults() *StampHistoryEntry`

NewStampHistoryEntryWithDefaults instantiates a new StampHistoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStampId

`func (o *StampHistoryEntry) GetStampId() string`

GetStampId returns the StampId field if non-nil, zero value otherwise.

### GetStampIdOk

`func (o *StampHistoryEntry) GetStampIdOk() (*string, bool)`

GetStampIdOk returns a tuple with the StampId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampId

`func (o *StampHistoryEntry) SetStampId(v string)`

SetStampId sets StampId field to given value.


### GetDatetime

`func (o *StampHistoryEntry) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *StampHistoryEntry) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *StampHistoryEntry) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


