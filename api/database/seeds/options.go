package seeds

import (
	"cms/api"
	"cms/api/domain"
)

func (s *Seeder) runOptions() error {

	optionsSeed := domain.Options{
		"site_title" : api.AppTitle,
		"site_description": api.AppDescription,
		"site_logo": api.AppLogo,
		"site_url" : api.AppUrl,
		"media_compression": 80,
		"media_convert_webp": true,
		"media_upload_max_size": 100000,
		"media_upload_max_width": 0,
		"media_upload_max_height": 0,
		"media_organise_year_month": false,
		"media_serve_webp": true,
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
		"cache": true,
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

	s.models.Options.UpdateCreate(optionsSeed)

	return nil
}
