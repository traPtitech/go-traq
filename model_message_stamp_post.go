/*
 * traQ API
 *
 * traQ v2 API
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MessageStampPost struct for MessageStampPost
type MessageStampPost struct {
	// 押す個数(1-100)、100より大きい場合は100として受け取ります
	Count int32 `json:"count"`
}