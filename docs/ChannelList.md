# ChannelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | [**[]Channel**](Channel.md) | パブリックチャンネルの配列 | 
**Dm** | Pointer to [**[]DMChannel**](DMChannel.md) | ダイレクトメッセージチャンネルの配列 | [optional] 

## Methods

### NewChannelList

`func NewChannelList(public []Channel, ) *ChannelList`

NewChannelList instantiates a new ChannelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelListWithDefaults

`func NewChannelListWithDefaults() *ChannelList`

NewChannelListWithDefaults instantiates a new ChannelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *ChannelList) GetPublic() []Channel`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ChannelList) GetPublicOk() (*[]Channel, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ChannelList) SetPublic(v []Channel)`

SetPublic sets Public field to given value.


### GetDm

`func (o *ChannelList) GetDm() []DMChannel`

GetDm returns the Dm field if non-nil, zero value otherwise.

### GetDmOk

`func (o *ChannelList) GetDmOk() (*[]DMChannel, bool)`

GetDmOk returns a tuple with the Dm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDm

`func (o *ChannelList) SetDm(v []DMChannel)`

SetDm sets Dm field to given value.

### HasDm

`func (o *ChannelList) HasDm() bool`

HasDm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


