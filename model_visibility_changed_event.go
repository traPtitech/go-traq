/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// VisibilityChangedEvent チャンネル可視状態変更イベント
type VisibilityChangedEvent struct {
	// 変更者UUID
	UserId string `json:"userId"`
	// 変更後可視状態
	Visibility bool `json:"visibility"`
}
