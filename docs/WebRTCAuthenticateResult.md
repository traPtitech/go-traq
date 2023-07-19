# WebRTCAuthenticateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerId** | **string** | ピアID | 
**Ttl** | **int32** | TTL | 
**Timestamp** | **int64** | タイムスタンプ | 
**AuthToken** | **string** | 認証トークン | 

## Methods

### NewWebRTCAuthenticateResult

`func NewWebRTCAuthenticateResult(peerId string, ttl int32, timestamp int64, authToken string, ) *WebRTCAuthenticateResult`

NewWebRTCAuthenticateResult instantiates a new WebRTCAuthenticateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRTCAuthenticateResultWithDefaults

`func NewWebRTCAuthenticateResultWithDefaults() *WebRTCAuthenticateResult`

NewWebRTCAuthenticateResultWithDefaults instantiates a new WebRTCAuthenticateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerId

`func (o *WebRTCAuthenticateResult) GetPeerId() string`

GetPeerId returns the PeerId field if non-nil, zero value otherwise.

### GetPeerIdOk

`func (o *WebRTCAuthenticateResult) GetPeerIdOk() (*string, bool)`

GetPeerIdOk returns a tuple with the PeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerId

`func (o *WebRTCAuthenticateResult) SetPeerId(v string)`

SetPeerId sets PeerId field to given value.


### GetTtl

`func (o *WebRTCAuthenticateResult) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *WebRTCAuthenticateResult) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *WebRTCAuthenticateResult) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetTimestamp

`func (o *WebRTCAuthenticateResult) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebRTCAuthenticateResult) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebRTCAuthenticateResult) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetAuthToken

`func (o *WebRTCAuthenticateResult) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *WebRTCAuthenticateResult) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *WebRTCAuthenticateResult) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


