// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"encoding/json"
)

type (
	// OptionsDBMap defines the map of key value pair options
	// that are stored in the database, used for marshalling
	// and unmarshalling into the Options struct.
	OptionsDBMap map[string]interface{}
	// OptionDB represents a singular entity of an option
	// that's stored in the database.
	OptionDB struct {
		ID    int             `db:"id" json:"id"`
		Name  string          `db:"option_name" json:"option_name" binding:"required"`
		Value json.RawMessage `db:"option_value" json:"option_value" binding:"required"`
	}
	// OptionsDB represents the slice of OptionDB's.
	OptionsDB []OptionDB
	// Options defines the main system options defined in the
	// store this is used throughout the application for
	// user defined choices.
	Options struct {
		// Site
		SiteTitle       string `json:"site_title" binding:"required"`
		SiteDescription string `json:"site_description" binding:"required"`
		SiteLogo        string `json:"site_logo" binding:"required"`
		SiteURL         string `json:"site_url" binding:"required,url"`
		// Theme
		ActiveTheme string `json:"active_theme" binding:"required"`
		Homepage    int    `json:"homepage"`
		// General
		GeneralLocale string `json:"general_locale" binding:"required"`
		// Contact
		ContactEmail     string `json:"contact_email" binding:"omitempty,email"`
		ContactTelephone string `json:"contact_telephone"`
		ContactAddress   string `json:"contact_address"`
		// Social
		SocialFacebook  string `json:"social_facebook" binding:"omitempty,url"`
		SocialTwitter   string `json:"social_twitter" binding:"omitempty,url"`
		SocialInstagram string `json:"social_instagram" binding:"omitempty,url"`
		SocialLinkedIn  string `json:"social_linkedin" binding:"omitempty,url"`
		SocialYoutube   string `json:"social_youtube" binding:"omitempty,url"`
		SocialPinterest string `json:"social_pinterest" binding:"omitempty,url"`
		// Code Injection
		CodeInjectionHead string `json:"codeinjection_head" binding:"omitempty"`
		CodeInjectionFoot string `json:"codeinjection_foot" binding:"omitempty"`
		// Meta
		MetaTitle               string `json:"meta_title" binding:"omitempty"`
		MetaDescription         string `json:"meta_description" binding:"omitempty"`
		MetaFacebookTitle       string `json:"meta_facebook_title" binding:"omitempty"`
		MetaFacebookDescription string `json:"meta_facebook_description" binding:"omitempty"`
		MetaFacebookImageID     int    `json:"meta_facebook_image_id" binding:"numeric"`
		MetaTwitterTitle        string `json:"meta_twitter_title" binding:"omitempty"`
		MetaTwitterDescription  string `json:"meta_twitter_description" binding:"omitempty"`
		MetaTwitterImageID      int    `json:"meta_twitter_image_id" binding:"omitempty,numeric"`
		// SEO
		SeoPrivate          bool     `json:"seo_private"`
		SeoSitemapServe     bool     `json:"seo_sitemap_serve"`
		SeoSitemapRedirects bool     `json:"seo_sitemap_redirects"`
		SeoSitemapExcluded  []string `json:"seo_sitemap_excluded"`
		SeoEnforceSlash     bool     `json:"seo_enforce_slash"`
		SeoRobotsServe      bool     `json:"seo_robots_serve"`
		SeoRobots           string   `json:"seo_robots"`
		// Breadcrumbs
		BreadcrumbsEnable       bool   `json:"breadcrumbs_enable"`
		BreadcrumbsTitle        string `json:"breadcrumbs_title"`
		BreadcrumbsSeparator    string `json:"breadcrumbs_separator"`
		BreadcrumbsHomepageText string `json:"breadcrumbs_homepage_text"`
		BreadcrumbsHideHomePage bool   `json:"breadcrumbs_hide_homepage"`
		// Media
		MediaCompression     int        `json:"media_compression" binding:"required"`
		MediaConvertWebP     bool       `json:"media_convert_webp"`
		MediaServeWebP       bool       `json:"media_serve_webp"`
		MediaUploadMaxSize   int64      `json:"media_upload_max_size" binding:"numeric"`
		MediaUploadMaxWidth  int64      `json:"media_upload_max_width" binding:"numeric"`
		MediaUploadMaxHeight int64      `json:"media_upload_max_height" binding:"numeric"`
		MediaOrganiseDate    bool       `json:"media_organise_year_month"`
		MediaSizes           MediaSizes `json:"media_images_sizes"`
		// Server Cache
		CacheServerTemplates bool `json:"cache_server_templates"`
		CacheServerFields    bool `json:"cache_server_field_layouts"`
		// Frontend Caching
		CacheFrontend          bool     `json:"cache_frontend"`
		CacheFrontendRequest   string   `json:"cache_frontend_request"`
		CacheFrontendSeconds   int64    `json:"cache_frontend_seconds"`
		CacheFrontendExtension []string `json:"cache_frontend_extensions"`
		// Gzip
		Gzip                   bool     `json:"gzip"`
		GzipCompression        string   `json:"gzip_compression"`
		GzipUsePaths           bool     `json:"gzip_use_paths"`
		GzipExcludedExtensions []string `json:"gzip_excluded_extensions"`
		GzipExcludedPaths      []string `json:"gzip_excluded_paths"`
		// Minify
		MinifyHTML bool `json:"minify_html"`
		MinifyJS   bool `json:"minify_js"`
		MinifyCSS  bool `json:"minify_css"`
		MinifySVG  bool `json:"minify_svg"`
		MinifyJSON bool `json:"minify_json"`
		MinifyXML  bool `json:"minify_xml"`
		// Forms
		FormSendEmailAddresses []string `json:"form_send_email_addresses"`
		FormFromEmailAddress   string   `json:"form_from_email_addresses"`
		FormIncludeLogo        bool     `json:"form_email_include_logo"`
		FormEmailDisclosure    string   `json:"form_email_disclosure"`
		// Storage
		StorageProvider     StorageProvider `json:"storage_provider"`
		StorageBucket       string          `json:"storage_bucket"`
		StorageUploadRemote bool            `json:"storage_upload_remote"`
		StorageLocalBackup  bool            `json:"storage_local_backup"`
		StorageRemoteBackup bool            `json:"storage_remote_backup"`
		// Proxies
		Proxies []Proxy `json:"proxies"`
	}
)
