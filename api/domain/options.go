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
	OptionsSite
	OptionsMedia
	OptionsCache
	OptionsGzip
}

type OptionsSite struct {
	Title string `json:"site_title"`
	Description string `json:"site_description"`
	Logo string `json:"site_logo"`
	Url string `json:"site_url"`
}

type OptionsMedia struct {
	Compression int `json:"media_compression"`
	ConvertWebP bool `json:"media_convert_webp"`
	ServeWebP bool `json:"media_serve_webp"`
	UploadMaxSize int `json:"media_upload_max_size"`
	UploadMaxWidth int `json:"media_upload_max_width"`
	UploadMaxHeight int `json:"media_upload_max_height"`
	OrganiseDate bool `json:"media_organise_year_month"`
	Sizes MediaSize `json:"media_images_sizes"`
}

type OptionsCache struct {
	Cache bool `json:"cache_global"`
	Layout bool `json:"cache_layout"`
	Fields bool `json:"cache_fields"`
	Site bool `json:"cache_site"`
	Templates bool `json:"cache_templates"`
	Resources bool `json:"cache_resources"`
}

type OptionsGzip struct {
	Compress bool `json:"gzip_compression"`
	Files []string `json:"gzip_files"`
}