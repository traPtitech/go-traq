/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// PostUnlinkExternalAccount POST /users/me/ex-accounts/unlink 用リクエストボディ
type PostUnlinkExternalAccount struct {
	// 外部サービス名
	ProviderName string `json:"providerName"`
}
