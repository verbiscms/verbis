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
		"social_facebook_url" : "",
		"social_twitter_url" : "",
		"social_instagram_url" : "",
		"social_linkedin_url" : "",
		"footer_text" : "",
		"codeinjection_head": "",
		"codeinjection_foot": "",
		"media_compression": 80,
		"media_convert_webp": true,
		"media_serve_webp": true,
		"media_upload_max_size": 100000,
		"media_upload_max_width": 0,
		"media_upload_max_height": 0,
		"media_organise_year_month": true,
		"media_images_sizes": domain.MediaSizes{
			"thumbnail": domain.MediaSize{
				Name: "Thumbnail Size",
				Width: 550,
				Height: 300,
				Crop: true,
			},
			"medium": domain.MediaSize{
				Name: "Medium Size",
				Width: 992,
				Height: 0,
				Crop: false,
			},
			"large": domain.MediaSize{
				Name: "Large Size",
				Width: 1280,
				Height: 0,
				Crop: false,
			},
			"hd": domain.MediaSize{
				Name: "HD Size",
				Width: 1920,
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
