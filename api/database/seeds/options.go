package seeds

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/domain"
)

// runOptions will insert all default values for the options
// into the database when installing Verbis.
func (s *Seeder) runOptions() error {
	const op = "Seeder.runOptions"

	optionsSeed := domain.OptionsDB{
		"site_title" : api.App.Title,
		"site_description": api.App.Description,
		"site_logo": api.App.Logo,
		"site_url" : api.App.Url,
		"contact_email" : "",
		"contact_telephone" : "",
		"contact_address" : "",
		"social_facebook" : "",
		"social_twitter" : "",
		"social_instagram" : "",
		"social_linkedin" : "",
		"social_youtube" : "",
		"social_pinterest" : "",
		"footer_text" : "",
		"footer_disclosure" : "",
		"seo_public" : false,
		"meta_title" : "",
		"meta_description" : "",
		"meta_facebook_title" : "",
		"meta_facebook_description" : "",
		"meta_facebook_image_id" : "",
		"meta_twitter_title" : "",
		"meta_twitter_description" : "",
		"meta_twitter_image_id" : "",
		"codeinjection_head": "",
		"codeinjection_foot": "",
		"media_compression": 80,
		"media_convert_webp": true,
		"media_serve_webp": true,
		"media_upload_max_size": 100000,
		"media_upload_max_width": 0,
		"media_upload_max_height": 0,
		"media_organise_year_month": true,
		"media_images_sizes": map[string]domain.MediaSizeOptions{
			"thumbnail": {
				Name:   "Thumbnail Size",
				Width:  550,
				Height: 300,
				Crop: true,
			},
			"medium": {
				Name:   "Medium Size",
				Width:  992,
				Height: 0,
				Crop: false,
			},
			"large": {
				Name:   "Large Size",
				Width:  1280,
				Height: 0,
				Crop: false,
			},
			"hd": {
				Name:   "HD Size",
				Width:  1920,
				Height: 0,
				Crop: false,
			},
		},
		"cache_global": true,
		"cache_layout": true,
		"cache_fields": true,
		"cache_site": true,
		"cache_templates": true,
		"cache_resources": true,
		"gzip_compression": true,
		"gzip_files": []string{
			"text/css",
			"text/javascript",
			"text/xml",
			"text/plain",
			"text/x-component",
			"application/javascript",
			"application/json",
			"application/xml",
			"application/rss+xml",
			"font/truetype",
			"font/opentype",
			"application/vnd.ms-fontobject",
			"image/svg+xml",
		},
	}

	err := s.models.Options.UpdateCreate(optionsSeed); if err != nil {
		return err
	}

	return nil
}
