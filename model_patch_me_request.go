/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// PatchMeRequest 自分のユーザー情報変更リクエスト
type PatchMeRequest struct {
	// 新しい表示名
	DisplayName string `json:"displayName,omitempty"`
	// TwitterID
	TwitterId string `json:"twitterId,omitempty"`
	// 自己紹介(biography)
	Bio string `json:"bio,omitempty"`
	// ホームチャンネルのUUID `00000000-0000-0000-0000-000000000000`を指定すると、ホームチャンネルが`null`に設定されます
	HomeChannel string `json:"homeChannel,omitempty"`
}
