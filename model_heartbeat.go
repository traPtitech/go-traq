/*
 * traQ API
 *
 * traQ v2 API
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Heartbeat struct for Heartbeat
type Heartbeat struct {
	Status HeartbeatStatus `json:"status"`
	// 現在いるチャンネルId
	ChannelId string `json:"channelId"`
}
