package domain

import (
	"encoding/json"
)

type OptionsDB map[string]interface{}

type OptionDB struct {
	ID				int					`db:"id" json:"id"`
	Name			string 				`db:"option_name" json:"option_name" binding:"required"`
	Value			json.RawMessage		`db:"option_value" json:"option_value" binding:"required"`
}

type Options struct {
	// Site
	SiteTitle string `json:"site_title"`
	SiteDescription string `json:"site_description"`
	SiteLogo string `json:"site_logo"`
	SiteUrl string `json:"site_url"`
	// Media
	MediaCompression int `json:"media_compression"`
	MediaConvertWebP bool `json:"media_convert_webp"`
	MediaServeWebP bool `json:"media_serve_webp"`
	MediaUploadMaxSize int `json:"media_upload_max_size"`
	MediaUploadMaxWidth int `json:"media_upload_max_width"`
	MediaUploadMaxHeight int `json:"media_upload_max_height"`
	MediaOrganiseDate bool `json:"media_organise_year_month"`
	MediaSizes map[string]interface{} `json:"media_images_sizes"`
	// Cache
	CacheGlobal bool `json:"cache_global"`
	CacheLayout bool `json:"cache_layout"`
	CacheFields bool `json:"cache_fields"`
	CacheSite bool `json:"cache_site"`
	CacheTemplates bool `json:"cache_templates"`
	CacheResources bool `json:"cache_resources"`
	// Gzip
	GzipCompress bool `json:"gzip_compression"`
	GzipFiles []string `json:"gzip_files"`
}