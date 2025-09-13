# PostMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | メッセージ本文 | 
**Embed** | Pointer to **bool** | メンション・チャンネルリンクを自動埋め込みするか | [optional] [default to false]
**Nonce** | Pointer to **string** | メッセージ送信の確認に使うことができる任意の識別子(投稿でのみ使用可) | [optional] 

## Methods

### NewPostMessageRequest

`func NewPostMessageRequest(content string, ) *PostMessageRequest`

NewPostMessageRequest instantiates a new PostMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostMessageRequestWithDefaults

`func NewPostMessageRequestWithDefaults() *PostMessageRequest`

NewPostMessageRequestWithDefaults instantiates a new PostMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PostMessageRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PostMessageRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PostMessageRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetEmbed

`func (o *PostMessageRequest) GetEmbed() bool`

GetEmbed returns the Embed field if non-nil, zero value otherwise.

### GetEmbedOk

`func (o *PostMessageRequest) GetEmbedOk() (*bool, bool)`

GetEmbedOk returns a tuple with the Embed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbed

`func (o *PostMessageRequest) SetEmbed(v bool)`

SetEmbed sets Embed field to given value.

### HasEmbed

`func (o *PostMessageRequest) HasEmbed() bool`

HasEmbed returns a boolean if a field has been set.

### GetNonce

`func (o *PostMessageRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *PostMessageRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *PostMessageRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *PostMessageRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


