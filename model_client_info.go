/*
 * traQ API
 *
 * traQ v2 API
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ClientInfo struct for ClientInfo
type ClientInfo struct {
	// クライアントID
	ClientId string `json:"clientId,omitempty"`
	// クライアント名
	Name string `json:"name,omitempty"`
	// クライアント説明
	Description string `json:"description,omitempty"`
	// クライアント作成者UUID
	CreatorId string `json:"creatorId,omitempty"`
}
