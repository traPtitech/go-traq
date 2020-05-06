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

// UnreadChannel 未読チャンネル情報
type UnreadChannel struct {
	// チャンネルUUID
	ChannelId string `json:"channelId"`
	// 未読メッセージ数
	Count int32 `json:"count"`
	// 自分宛てメッセージが含まれているかどうか
	Noticeable bool `json:"noticeable"`
	// チャンネルの最古の未読メッセージの日時
	Since time.Time `json:"since"`
	// チャンネルの最新の未読メッセージの日時
	UpdatedAt time.Time `json:"updatedAt"`
}
