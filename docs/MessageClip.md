# MessageClip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | **string** | クリップされているフォルダのID | 
**ClippedAt** | **time.Time** | クリップされた日時 | 

## Methods

### NewMessageClip

`func NewMessageClip(folderId string, clippedAt time.Time, ) *MessageClip`

NewMessageClip instantiates a new MessageClip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageClipWithDefaults

`func NewMessageClipWithDefaults() *MessageClip`

NewMessageClipWithDefaults instantiates a new MessageClip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *MessageClip) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *MessageClip) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *MessageClip) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### GetClippedAt

`func (o *MessageClip) GetClippedAt() time.Time`

GetClippedAt returns the ClippedAt field if non-nil, zero value otherwise.

### GetClippedAtOk

`func (o *MessageClip) GetClippedAtOk() (*time.Time, bool)`

GetClippedAtOk returns a tuple with the ClippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClippedAt

`func (o *MessageClip) SetClippedAt(v time.Time)`

SetClippedAt sets ClippedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


