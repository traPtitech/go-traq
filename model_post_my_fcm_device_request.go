/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// PostMyFcmDeviceRequest FCMデバイス登録リクエスト
type PostMyFcmDeviceRequest struct {
	// FCMのデバイストークン
	Token string `json:"token"`
}
