/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TopicChangedEvent TopicChanged
type TopicChangedEvent struct {
	// 変更者UUID
	UserId string `json:"userId,omitempty"`
	// 変更前トピック
	Before string `json:"before,omitempty"`
	// 変更後トピック
	After string `json:"after,omitempty"`
}
