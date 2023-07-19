# \ActivityApi

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivityTimeline**](ActivityApi.md#GetActivityTimeline) | **Get** /activity/timeline | アクテビティタイムラインを取得
[**GetOnlineUsers**](ActivityApi.md#GetOnlineUsers) | **Get** /activity/onlines | オンラインユーザーリストを取得



## GetActivityTimeline

> []ActivityTimelineMessage GetActivityTimeline(ctx).Limit(limit).All(all).PerChannel(perChannel).Execute()

アクテビティタイムラインを取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    traq "github.com/traPtitech/go-traq"
)

func main() {
    limit := int32(56) // int32 | 取得する件数 (optional) (default to 50)
    all := true // bool | 全てのチャンネルのタイムラインを取得する (optional) (default to false)
    perChannel := true // bool | 同じチャンネルのメッセージは最新のもののみ取得するか (optional) (default to false)

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.GetActivityTimeline(context.Background()).Limit(limit).All(all).PerChannel(perChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GetActivityTimeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActivityTimeline`: []ActivityTimelineMessage
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GetActivityTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | 取得する件数 | [default to 50]
 **all** | **bool** | 全てのチャンネルのタイムラインを取得する | [default to false]
 **perChannel** | **bool** | 同じチャンネルのメッセージは最新のもののみ取得するか | [default to false]

### Return type

[**[]ActivityTimelineMessage**](ActivityTimelineMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnlineUsers

> []string GetOnlineUsers(ctx).Execute()

オンラインユーザーリストを取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    traq "github.com/traPtitech/go-traq"
)

func main() {

    configuration := traq.NewConfiguration()
    apiClient := traq.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.GetOnlineUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GetOnlineUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOnlineUsers`: []string
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GetOnlineUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnlineUsersRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

