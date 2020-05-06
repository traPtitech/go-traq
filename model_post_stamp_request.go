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
	"os"
)

// PostStampRequest スタンプ作成リクエスト
type PostStampRequest struct {
	// スタンプ名
	Name string `json:"name"`
	// スタンプ画像(1MBまでのpng, jpeg, gif)
	File *os.File `json:"file"`
}
