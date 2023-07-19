# BotEventLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BotId** | **string** | BOT UUID | 
**RequestId** | **string** | リクエストUUID | 
**Event** | **string** | イベントタイプ | 
**Result** | Pointer to [**BotEventResult**](BotEventResult.md) |  | [optional] 
**Code** | **int32** | ステータスコード | 
**Datetime** | **time.Time** | イベント日時 | 

## Methods

### NewBotEventLog

`func NewBotEventLog(botId string, requestId string, event string, code int32, datetime time.Time, ) *BotEventLog`

NewBotEventLog instantiates a new BotEventLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotEventLogWithDefaults

`func NewBotEventLogWithDefaults() *BotEventLog`

NewBotEventLogWithDefaults instantiates a new BotEventLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBotId

`func (o *BotEventLog) GetBotId() string`

GetBotId returns the BotId field if non-nil, zero value otherwise.

### GetBotIdOk

`func (o *BotEventLog) GetBotIdOk() (*string, bool)`

GetBotIdOk returns a tuple with the BotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotId

`func (o *BotEventLog) SetBotId(v string)`

SetBotId sets BotId field to given value.


### GetRequestId

`func (o *BotEventLog) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BotEventLog) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BotEventLog) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetEvent

`func (o *BotEventLog) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *BotEventLog) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *BotEventLog) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetResult

`func (o *BotEventLog) GetResult() BotEventResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BotEventLog) GetResultOk() (*BotEventResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BotEventLog) SetResult(v BotEventResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *BotEventLog) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCode

`func (o *BotEventLog) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BotEventLog) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BotEventLog) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDatetime

`func (o *BotEventLog) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *BotEventLog) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *BotEventLog) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


