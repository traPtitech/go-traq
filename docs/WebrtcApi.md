# \WebrtcApi

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebRTCState**](WebrtcApi.md#GetWebRTCState) | **Get** /webrtc/state | WebRTC状態を取得
[**PostWebRTCAuthenticate**](WebrtcApi.md#PostWebRTCAuthenticate) | **Post** /webrtc/authenticate | Skyway用認証API



## GetWebRTCState

> []WebRTCUserState GetWebRTCState(ctx).Execute()

WebRTC状態を取得



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebrtcApi.GetWebRTCState(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebrtcApi.GetWebRTCState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebRTCState`: []WebRTCUserState
    fmt.Fprintf(os.Stdout, "Response from `WebrtcApi.GetWebRTCState`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebRTCStateRequest struct via the builder pattern


### Return type

[**[]WebRTCUserState**](WebRTCUserState.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebRTCAuthenticate

> WebRTCAuthenticateResult PostWebRTCAuthenticate(ctx).PostWebRTCAuthenticateRequest(postWebRTCAuthenticateRequest).Execute()

Skyway用認証API



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
    postWebRTCAuthenticateRequest := *openapiclient.NewPostWebRTCAuthenticateRequest("PeerId_example") // PostWebRTCAuthenticateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebrtcApi.PostWebRTCAuthenticate(context.Background()).PostWebRTCAuthenticateRequest(postWebRTCAuthenticateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebrtcApi.PostWebRTCAuthenticate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWebRTCAuthenticate`: WebRTCAuthenticateResult
    fmt.Fprintf(os.Stdout, "Response from `WebrtcApi.PostWebRTCAuthenticate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWebRTCAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postWebRTCAuthenticateRequest** | [**PostWebRTCAuthenticateRequest**](PostWebRTCAuthenticateRequest.md) |  | 

### Return type

[**WebRTCAuthenticateResult**](WebRTCAuthenticateResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

