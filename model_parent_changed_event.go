/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// ParentChangedEvent 親チャンネル変更イベント
type ParentChangedEvent struct {
	// 変更者UUID
	UserId string `json:"userId"`
	// 変更前親チャンネルUUID
	Before string `json:"before"`
	// 変更後親チャンネルUUID
	After string `json:"after"`
}
