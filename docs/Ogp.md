# Ogp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** |  | 
**Url** | **string** |  | 
**Images** | [**[]OgpMedia**](OgpMedia.md) |  | 
**Description** | **string** |  | 
**Videos** | [**[]OgpMedia**](OgpMedia.md) |  | 

## Methods

### NewOgp

`func NewOgp(type_ string, title string, url string, images []OgpMedia, description string, videos []OgpMedia, ) *Ogp`

NewOgp instantiates a new Ogp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOgpWithDefaults

`func NewOgpWithDefaults() *Ogp`

NewOgpWithDefaults instantiates a new Ogp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Ogp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ogp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ogp) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *Ogp) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Ogp) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Ogp) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *Ogp) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Ogp) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Ogp) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetImages

`func (o *Ogp) GetImages() []OgpMedia`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Ogp) GetImagesOk() (*[]OgpMedia, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Ogp) SetImages(v []OgpMedia)`

SetImages sets Images field to given value.


### GetDescription

`func (o *Ogp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ogp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Ogp) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVideos

`func (o *Ogp) GetVideos() []OgpMedia`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *Ogp) GetVideosOk() (*[]OgpMedia, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *Ogp) SetVideos(v []OgpMedia)`

SetVideos sets Videos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


