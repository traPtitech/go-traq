/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// MessagePin ピン情報
type MessagePin struct {
	// ピン留めしたユーザーUUID
	UserId string `json:"userId"`
	// ピン留めされた日時
	PinnedAt time.Time `json:"pinnedAt"`
}
