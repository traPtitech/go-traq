/*
 * traQ API
 *
 * traQ v2 API
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// Session struct for Session
type Session struct {
	Id            string    `json:"id,omitempty"`
	LastIP        string    `json:"lastIP,omitempty"`
	LastUserAgent string    `json:"lastUserAgent,omitempty"`
	LastAccess    time.Time `json:"lastAccess,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
}
