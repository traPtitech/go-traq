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

// ActiveOAuth2Token 有効なOAuth2トークン情報
type ActiveOAuth2Token struct {
	// トークンUUID
	Id string `json:"id"`
	// OAuth2クライアントUUID
	ClientId string `json:"clientId"`
	// スコープ
	Scopes []OAuth2Scope `json:"scopes"`
	// 発行日時
	IssuedAt time.Time `json:"issuedAt"`
}
