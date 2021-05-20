// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seeds

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/domain"
)

// runOptions will insert all default values for the options
// into the database when installing Verbis.
func (s *Seeder) runOptions() error {
	const op = "Seeder.runOptions"

	optionsSeed := domain.OptionsDBMap{
		// Site
		"site_title":       api.App.Title,
		"site_description": api.App.Description,
		"site_logo":        api.App.Logo,
		"site_url":         api.App.URL,
		// General
		"general_locale": "en_GB",
		// Theme
		// TODO: This theme needs to be a constant
		"active_theme": "Verbis",
		"homepage":     0,
		// Contact
		"contact_email":     "",
		"contact_telephone": "",
		"contact_address":   "",
		// Social
		"social_facebook":   "",
		"social_twitter":    "",
		"social_instagram":  "",
		"social_linkedin":   "",
		"social_youtube":    "",
		"social_pinterest":  "",
		"footer_text":       "",
		"footer_disclosure": "",
		// Code Injection
		"codeinjection_head": "",
		"codeinjection_foot": "",
		// Meta
		"meta_title":                "",
		"meta_description":          "",
		"meta_facebook_title":       "",
		"meta_facebook_description": "",
		"meta_facebook_image_id":    0,
		"meta_twitter_title":        "",
		"meta_twitter_description":  "",
		"meta_twitter_image_id":     0,
		// SEO
		"seo_private":           false,
		"seo_sitemap_serve":     true,
		"seo_sitemap_redirects": true,
		"seo_sitemap_excluded":  []string{},
		"seo_enforce_slash":     false,
		"seo_robots_serve":      true,
		"seo_robots":            "User-agent: *\nAllow: /",
		// Breadcrumbs
		"breadcrumbs_enable":        true,
		"breadcrumbs_title":         "",
		"breadcrumbs_separator":     "|",
		"breadcrumbs_homepage_text": "Home",
		"breadcrumbs_hide_homepage": true,
		// Media
		"media_compression":         80,
		"media_convert_webp":        true,
		"media_serve_webp":          true,
		"media_upload_max_size":     100000,
		"media_upload_max_width":    0,
		"media_upload_max_height":   0,
		"media_organise_year_month": true,
		"media_images_sizes": map[string]domain.MediaSizeOptions{
			"thumbnail": {
				Name:   "Thumbnail Size",
				Width:  550,
				Height: 300,
				Crop:   true,
			},
			"medium": {
				Name:   "Medium Size",
				Width:  992,
				Height: 0,
				Crop:   false,
			},
			"large": {
				Name:   "Large Size",
				Width:  1280,
				Height: 0,
				Crop:   false,
			},
			"hd": {
				Name:   "HD Size",
				Width:  1920,
				Height: 0,
				Crop:   false,
			},
		},
		// Cache
		"cache_server_templates":     false,
		"cache_server_field_layouts": false,
		// Frontend Caching
		"cache_frontend":            true,
		"cache_frontend_request":    "max-age",
		"cache_frontend_seconds":    31536000,
		"cache_frontend_extensions": []string{"jpg", "jpeg", "gif", "png", "ico", "cur", "webp", "jxr", "svg", "css", "js", "htc", "ttf", "tt", "otf", "eot", "woff", "woff2", "webm"},
		// Gzip
		"gzip":                     true,
		"gzip_compression":         "default-compression",
		"gzip_use_paths":           false,
		"gzip_excluded_extensions": []string{},
		"gzip_excluded_paths":      []string{},
		// Minify
		"minify_html": false,
		"minify_js":   false,
		"minify_css":  false,
		"minify_svg":  false,
		"minify_json": false,
		"minify_xml":  false,
	}

	err := s.models.Options.Insert(optionsSeed)
	if err != nil {
		return err
	}

	return nil
}
