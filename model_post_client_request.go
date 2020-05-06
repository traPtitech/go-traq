/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PostClientRequest OAuth2クライアント作成リクエスト
type PostClientRequest struct {
	// クライアント名
	Name string `json:"name"`
	// コールバックURL
	CallbackUrl string `json:"callbackUrl"`
	// 要求スコープの配列
	Scopes []OAuth2Scope `json:"scopes"`
	// 説明
	Description string `json:"description"`
}
