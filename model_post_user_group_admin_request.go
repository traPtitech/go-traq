/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// PostUserGroupAdminRequest グループ管理者追加リクエスト
type PostUserGroupAdminRequest struct {
	// 追加するユーザーのUUID
	Id string `json:"id"`
}