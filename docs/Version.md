# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **string** | traQ(サーバー)リビジョン | 
**Version** | **string** | traQ(サーバー)バージョン | 
**Flags** | [**VersionFlags**](VersionFlags.md) |  | 

## Methods

### NewVersion

`func NewVersion(revision string, version string, flags VersionFlags, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *Version) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Version) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Version) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetVersion

`func (o *Version) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Version) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Version) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFlags

`func (o *Version) GetFlags() VersionFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Version) GetFlagsOk() (*VersionFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Version) SetFlags(v VersionFlags)`

SetFlags sets Flags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


