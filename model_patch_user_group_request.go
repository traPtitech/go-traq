/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// PatchUserGroupRequest ユーザーグループ編集リクエスト
type PatchUserGroupRequest struct {
	// グループ名
	Name string `json:"name,omitempty"`
	// グループ説明
	Description string `json:"description,omitempty"`
	// グループタイプ
	Type string `json:"type,omitempty"`
}
