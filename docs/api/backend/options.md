## Options

Options define the core parameters of Verbis, and how it should accept and serve content. These routes all you to read,
create and update options through the Verbis API. All endpoints are prefixed with `/options`.

👉 `GET` to `/api/{version}/options`

**Example Response:**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained options",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {
		"cache_fields": true,
		"cache_frontend": true,
		"cache_frontend_extensions": [
			"jpg",
			"jpeg",
			"gif",
			"png",
			"ico",
			"cur",
			"webp",
			"jxr",
			"svg",
			"css",
			"js",
			"htc",
			"ttf",
			"tt",
			"otf",
			"eot",
			"woff",
			"woff2"
		],
		"cache_frontend_request": "max-age",
		"cache_frontend_seconds": 31536000,
		"cache_global": true,
		"cache_layout": true,
		"cache_resources": true,
		"cache_site": true,
		"cache_templates": true,
		"codeinjection_foot": "",
		"codeinjection_head": "",
		"contact_address": "",
		"contact_email": "",
		"contact_telephone": "",
		"footer_disclosure": "",
		"footer_text": "",
		"general_locale": "en_GB",
		"gzip": true,
		"gzip_compression": "default-compression",
		"gzip_excluded_extensions": [],
		"gzip_excluded_paths": [],
		"gzip_use_paths": false,
		"media_compression": 80,
		"media_convert_webp": true,
		"media_images_sizes": {
			"hd": {
				"crop": false,
				"height": 0,
				"name": "HD Size",
				"width": 1920
			},
			"large": {
				"crop": false,
				"height": 0,
				"name": "Large Size",
				"width": 1280
			},
			"medium": {
				"crop": false,
				"height": 0,
				"name": "Medium Size",
				"width": 992
			},
			"thumbnail": {
				"crop": true,
				"height": 300,
				"name": "Thumbnail Size",
				"width": 550
			}
		},
		"media_organise_year_month": true,
		"media_serve_webp": true,
		"media_upload_max_height": 0,
		"media_upload_max_size": 100000,
		"media_upload_max_width": 0,
		"meta_description": "",
		"meta_facebook_description": "",
		"meta_facebook_image_id": 0,
		"meta_facebook_title": "",
		"meta_title": "",
		"meta_twitter_description": "",
		"meta_twitter_image_id": 0,
		"meta_twitter_title": "",
		"minify_css": false,
		"minify_html": false,
		"minify_js": false,
		"minify_json": false,
		"minify_svg": false,
		"minify_xml": false,
		"seo_public": false,
		"seo_redirects": null,
		"seo_robots": "\"User-Agent: *\\nAllow: /\"",
		"seo_robots_serve": true,
		"seo_sitemap_excluded": [],
		"seo_sitemap_redirects": true,
		"seo_sitemap_serve": true,
		"site_description": "A Verbis website. Publish online, build a business, work from home",
		"site_logo": "/verbis/images/verbis-logo.svg",
		"site_title": "Verbis",
		"site_url": "http://127.0.0.1:8080",
		"social_facebook": "",
		"social_instagram": "",
		"social_linkedin": "",
		"social_pinterest": "",
		"social_twitter": "",
		"social_youtube": ""
	}
}
```

## Create or Update an Option

WIP: Experimental

