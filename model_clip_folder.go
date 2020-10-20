/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

import (
	"time"
)

// ClipFolder クリップフォルダ情報
type ClipFolder struct {
	// フォルダUUID
	Id string `json:"id"`
	// フォルダ名
	Name string `json:"name"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// フォルダ所有者UUID
	OwnerId string `json:"ownerId"`
	// 説明
	Description string `json:"description"`
}
