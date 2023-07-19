# \OgpApi

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOgpCache**](OgpApi.md#DeleteOgpCache) | **Delete** /ogp/cache | OGP情報のキャッシュを削除
[**GetOgp**](OgpApi.md#GetOgp) | **Get** /ogp | OGP情報を取得



## DeleteOgpCache

> DeleteOgpCache(ctx).Url(url).Execute()

OGP情報のキャッシュを削除



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    url := "url_example" // string | OGPのキャッシュを削除したいURL

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OgpApi.DeleteOgpCache(context.Background()).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OgpApi.DeleteOgpCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOgpCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | OGPのキャッシュを削除したいURL | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOgp

> Ogp GetOgp(ctx).Url(url).Execute()

OGP情報を取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    url := "url_example" // string | OGPを取得したいURL

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OgpApi.GetOgp(context.Background()).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OgpApi.GetOgp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOgp`: Ogp
    fmt.Fprintf(os.Stdout, "Response from `OgpApi.GetOgp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOgpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | OGPを取得したいURL | 

### Return type

[**Ogp**](Ogp.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

