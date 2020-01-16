/*
 * traQ API
 *
 * traQ v2 API
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PatchUserGroup struct for PatchUserGroup
type PatchUserGroup struct {
	// グループ名
	Name string `json:"name,omitempty"`
	// 説明
	Description string `json:"description,omitempty"`
	// 管理ユーザー
	AdminUserId string `json:"adminUserId,omitempty"`
}
