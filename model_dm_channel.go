/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DmChannel ダイレクトメッセージチャンネル
type DmChannel struct {
	// チャンネルUUID
	Id string `json:"id"`
	// 送信先相手のUUID
	UserId string `json:"userId"`
}
