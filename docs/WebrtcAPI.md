# \WebrtcAPI

All URIs are relative to *https://q.trap.jp/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebRTCState**](WebrtcAPI.md#GetWebRTCState) | **Get** /webrtc/state | WebRTC状態を取得
[**PostWebRTCAuthenticate**](WebrtcAPI.md#PostWebRTCAuthenticate) | **Post** /webrtc/authenticate | Skyway用認証API



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
	traq "github.com/traPtitech/go-traq"
)

func main() {

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.WebrtcAPI.GetWebRTCState(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebrtcAPI.GetWebRTCState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebRTCState`: []WebRTCUserState
	fmt.Fprintf(os.Stdout, "Response from `WebrtcAPI.GetWebRTCState`: %v\n", resp)
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
	traq "github.com/traPtitech/go-traq"
)

func main() {
	postWebRTCAuthenticateRequest := *traq.NewPostWebRTCAuthenticateRequest("PeerId_example") // PostWebRTCAuthenticateRequest |  (optional)

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	resp, r, err := apiClient.WebrtcAPI.PostWebRTCAuthenticate(context.Background()).PostWebRTCAuthenticateRequest(postWebRTCAuthenticateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebrtcAPI.PostWebRTCAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebRTCAuthenticate`: WebRTCAuthenticateResult
	fmt.Fprintf(os.Stdout, "Response from `WebrtcAPI.PostWebRTCAuthenticate`: %v\n", resp)
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

