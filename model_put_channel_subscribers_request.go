/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// PutChannelSubscribersRequest 通知をオンにするユーザーのUUID配列
type PutChannelSubscribersRequest struct {
	// 通知をオンにするユーザーのUUID配列
	On []string `json:"on"`
}