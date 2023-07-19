# MessageSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalHits** | **int64** | 検索にヒットしたメッセージ件数 | 
**Hits** | [**[]Message**](Message.md) | 検索にヒットしたメッセージの配列 | 

## Methods

### NewMessageSearchResult

`func NewMessageSearchResult(totalHits int64, hits []Message, ) *MessageSearchResult`

NewMessageSearchResult instantiates a new MessageSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageSearchResultWithDefaults

`func NewMessageSearchResultWithDefaults() *MessageSearchResult`

NewMessageSearchResultWithDefaults instantiates a new MessageSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalHits

`func (o *MessageSearchResult) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *MessageSearchResult) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *MessageSearchResult) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.


### GetHits

`func (o *MessageSearchResult) GetHits() []Message`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *MessageSearchResult) GetHitsOk() (*[]Message, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *MessageSearchResult) SetHits(v []Message)`

SetHits sets Hits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


