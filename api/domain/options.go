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
	SiteTitle 				string 					`json:"site_title" binding:"required"`
	SiteDescription 		string 					`json:"site_description" binding:"required"`
	SiteLogo 				string 					`json:"site_logo" binding:"required"`
	SiteUrl 				string 					`json:"site_url" binding:"required,url"`
	SitePublic 				bool 					`json:"site_public"`
	// Media
	MediaCompression		int 					`json:"media_compression" binding:"required"`
	MediaConvertWebP 		bool 					`json:"media_convert_webp"`
	MediaServeWebP 			bool 					`json:"media_serve_webp"`
	MediaUploadMaxSize 		int 					`json:"media_upload_max_size" binding:"numeric"`
	MediaUploadMaxWidth 	int 					`json:"media_upload_max_width" binding:"numeric"`
	MediaUploadMaxHeight 	int 					`json:"media_upload_max_height" binding:"numeric"`
	MediaOrganiseDate 		bool 					`json:"media_organise_year_month"`
	MediaSizes 				MediaSizes 				`json:"media_images_sizes"`
	// Contact
	ContactEmail			string 					`json:"contact_email" binding:"omitempty,email"`
	ContactTelephone		string 					`json:"contact_telephone"`
	ContactAddress			string 					`json:"contact_address"`
	// Social
	SocialFacebook			string 					`json:"social_facebook" binding:"omitempty,url"`
	SocialTwitter			string 					`json:"social_twitter" binding:"omitempty,url"`
	SocialInstagram			string 					`json:"social_instagram" binding:"omitempty,url"`
	SocialLinkedIn			string 					`json:"social_linkedin" binding:"omitempty,url"`
	SocialYoutube			string 					`json:"social_youtube" binding:"omitempty,url"`
	SocialPinterest			string 					`json:"social_pinterest" binding:"omitempty,url"`
	// Cache
	CacheGlobal 			bool 					`json:"cache_global"`
	CacheLayout 			bool 					`json:"cache_layout"`
	CacheFields 			bool 					`json:"cache_fields"`
	CacheSite 				bool 					`json:"cache_site"`
	CacheTemplates 			bool 					`json:"cache_templates"`
	CacheResources			bool 					`json:"cache_resources"`
	// Frontend Caching
	CacheFrontend			bool					`json:"cache_frontend"`
	CacheFrontendRequest	string					`json:"cache_frontend_request"`
	CacheFrontendSeconds	int64					`json:"cache_frontend_seconds"`
	CacheFrontendExtension	[]string				`json:"cache_frontend_extensions"`
	// Gzip
	Gzip 		 			bool 					`json:"gzip"`
	GzipCompression 		string 					`json:"gzip_compression"`
	GzipUsePaths		 	bool 					`json:"gzip_use_paths"`
	GzipExcludedExtensions 	[]string 				`json:"gzip_excluded_extensions"`
	GzipExcludedPaths 		[]string 				`json:"gzip_excluded_paths"`
	// Minify
	MinifyHTML 				bool 					`json:"minify_html"`
	MinifyJS				bool 					`json:"minify_js"`
	MinifyCSS				bool 					`json:"minify_css"`
	MinifySVG 				bool 					`json:"minify_svg"`
	MinifyJSON				bool 					`json:"minify_json"`
	MinifyXML				bool 					`json:"minify_xml"`
}