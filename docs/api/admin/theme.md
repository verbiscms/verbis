# Theme

The theme endpoints are helpers to obtain configuration files from the theme's directory and obtain the templates and
layouts that are available.

## Theme Config

The `/theme` endpoint is used to retrieve the theme's `config.yml` file within the theme directory. This can be
particularly useful for establishing what resources the current theme has, and general information about the currently
activated theme including a title, description and theme version.

👉 `GET` to `/api/{version}/theme`

**Example Response:**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained theme config",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms",
		"pagination": {
			"page": 1,
			"pages": 1,
			"limit": 15,
			"total": 1,
			"next": false,
			"prev": false
		}
	},
	"data": {
		"theme": {
			"title": "A Verbis Theme",
			"description": "Stock theme for verbis",
			"version": "0.0.1"
		},
		"resources": {
			"posts": {
				"name": "posts",
				"friendly_name": "Posts",
				"singular_name": "Post",
				"slug": "/posts",
				"icon": "feather feather-clipboard"
			}
		},
		"assets_path": "/assets",
		"file_extension": ".cms",
		"template_dir": "templates",
		"layout_dir": "layouts",
		"editor": {
			"modules": [
				"blockquote",
				"code_block",
				"code_block_highlight",
				"code_view",
				"hardbreak",
				"h1",
				"h2",
				"h3",
				"h4",
				"h5",
				"h6",
				"paragraph",
				"hr",
				"ul",
				"ol",
				"bold",
				"code",
				"italic",
				"link",
				"strike",
				"underline",
				"history",
				"search",
				"trailing_node",
				"color",
				"table"
			],
			"options": {
				"link": {
					"rel": "noopener noreferrer nofollow"
				},
				"ol": {
					"class": "list list-ordered"
				},
				"palette": [
					"#fad839",
					"#ca1f26",
					"#1b3990",
					"#333"
				],
				"ul": {
					"class": "list"
				}
			}
		}
	}
}
```

## Templates

The `/templates` endpoint is used to retrieve the theme's all page templates for the currently activated theme that
reside in the templates' folder set in the `config.yml`.

- `key`: represents the page template file name.
- `name`: is a friendly name for the page template.

👉 `GET` to `/api/{version}/templates`

**Example Response:**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained templates",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {
		"templates": [
			{
				"key": "default",
				"name": "Default"
			},
			{
				"key": "archive",
				"name": "Archive"
			},
			{
				"key": "archive-single",
				"name": "Archive Single"
			}
		]
	}
}
```

## Layouts

The `/layouts` endpoint is used to retrieve the theme's all page layouts for the currently activated theme that reside
in the layouts' folder set in the `config.yml`.

- `key`: represents the page layout file name.
- `name`: is a friendly name for the page layout.

👉 `GET` to `/api/{version}/layouts`

**Example Response:**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained layouts",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {
		"layouts": [
			{
				"key": "default",
				"name": "Default"
			},
			{
				"key": "main",
				"name": "Main"
			}
		]
	}
}
```