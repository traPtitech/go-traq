/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// Version バージョン・サーバーフラグ情報
type Version struct {
	// traQ(サーバー)リビジョン
	Revision string `json:"revision"`
	// traQ(サーバー)バージョン
	Version string       `json:"version"`
	Flags   VersionFlags `json:"flags"`
}