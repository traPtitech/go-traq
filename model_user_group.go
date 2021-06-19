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

// UserGroup ユーザーグループ
type UserGroup struct {
	// グループUUID
	Id string `json:"id"`
	// グループ名
	Name string `json:"name"`
	// グループ説明
	Description string `json:"description"`
	// グループタイプ
	Type string `json:"type"`
	// グループアイコンUUID
	Icon string `json:"icon"`
	// グループメンバーの配列
	Members []UserGroupMember `json:"members"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
	// グループ管理者のUUIDの配列
	Admins []string `json:"admins"`
}
