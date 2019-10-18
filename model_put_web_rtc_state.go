/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PutWebRtcState struct for PutWebRtcState
type PutWebRtcState struct {
	// ユーザーの接続先チャンネルUUID(切断はnullを指定)
	ChannelId *string `json:"channelId,omitempty"`
	// ユーザーの状態の配列(切断は空配列を指定)
	State []string `json:"state,omitempty"`
}
