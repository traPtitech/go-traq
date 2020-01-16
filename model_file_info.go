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

// FileInfo struct for FileInfo
type FileInfo struct {
	// ファイルUUID
	FileId string `json:"fileId,omitempty"`
	// ファイル名
	Name string `json:"name,omitempty"`
	// MIMEタイプ
	Mime string `json:"mime,omitempty"`
	// ファイルサイズ
	Size int32 `json:"size,omitempty"`
	// MD5ハッシュ
	Md5 string `json:"md5,omitempty"`
	// サムネイルがあるかどうか
	HasThumb bool `json:"hasThumb,omitempty"`
	// サムネイル幅
	ThumbWidth int32 `json:"thumbWidth,omitempty"`
	// サムネイル高さ
	ThumbHeight int32 `json:"thumbHeight,omitempty"`
	// アップロード日時
	Datetime time.Time `json:"datetime,omitempty"`
}
