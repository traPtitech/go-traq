/*
 * traQ API
 *
 * traQ v2 API
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// BotTokens struct for BotTokens
type BotTokens struct {
	// Verification Token
	VerificationToken string `json:"verificationToken,omitempty"`
	// BOTアクセストークン
	AccessToken string `json:"accessToken,omitempty"`
	// BOTインストールコード
	BotCode string `json:"botCode,omitempty"`
}