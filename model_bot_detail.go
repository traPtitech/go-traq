/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// BotDetail struct for BotDetail
type BotDetail struct {
	// BOT UUID
	BotId string `json:"botId,omitempty"`
	// BOTユーザーUUID
	BotUserId string `json:"botUserId,omitempty"`
	// BOT説明
	Description string `json:"description,omitempty"`
	// BOTが購読しているイベントの配列
	SubscribeEvents []string `json:"subscribeEvents,omitempty"`
	// BOT状態
	State int32 `json:"state,omitempty"`
	// BOT作成者UUID
	CreatorId string `json:"creatorId,omitempty"`
	// BOT作成日時
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// BOT更新日時
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// BOT認証コード
	VerificationCode string `json:"verificationCode,omitempty"`
	// BOTアクセストークン
	AccessToken string `json:"accessToken,omitempty"`
	// BOTサーバーエンドポイント
	PostUrl string `json:"postUrl,omitempty"`
	// 権限昇格されているかどうか
	Privileged bool `json:"privileged,omitempty"`
	// BOTインストールコード
	BotCode string `json:"botCode,omitempty"`
}
