# Admin API

The Verbis API allows you to create and manage your content easily using wide variety of different endpoints. Everything
you are able to do in the admin interface you are able to achieve with the admin API, which makes it incredibly flexible
and dynamic.

All the Verbis endpoints follow REST principles, if you've interacted with a RESTful API already, many of the concepts
will be familiar to you. Requests and responses are JSON encoded throughout with different query parameters available to
filter and search through content easily and quickly.

## Structure

### Base URL:

The following URL is used for all API endpoints.

`/api/{version}/`

### Version

After the `/api` segment of the base URL, a version is prefixed with a `v`, dependent on what version of Verbis is
currently being run. For example: `/api/v1/posts`

### Endpoints

As Verbis follows RESTful principles, the following methods are accepted for each endpoint.

- `GET`: for browsing entities, `/api/{version}/posts` or reading an entity by a particular
  key, `/api/{version}/posts/1`
- `POST`: for adding entities, `/api/{version}/posts`
- `PUT`: for updating an entity usually by ID, `/api/{version}/posts/1`
- `DELETE`: for deleting an entity, usually by ID `/api/{version}/posts/1`

## JSON Responses

All endpoints respond with a JSON encoded response, which contains the following:

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained posts",
	"meta": {},
	"data": []
}
```

- `status`: contains an integer of the resulting http status code of the call.
- `error`: contains a boolean to signify if there was an error calling the endpoint.
- `message`: contains a brief message about the status of the call, and a useful description if there was an error.
- `meta`: includes Pagination if the route is a `GET` (browse) endpoint. It also includes `request_time`
  , `response_time` and `latency` times.
- `data`: contains the main body of the call, usually this is an array, but sometimes it can be an object depending on
  the endpoint.

## Errors

If the resulting status code of an endpoint is a validation error, the API will respond with some useful information
about what keys are required as an object in the `data` key under `errors`.

```json
{
	"status": 400,
	"error": true,
	"message": "Validation failed",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {
		"errors": [
			{
				"key": "slug",
				"type": "required",
				"message": "Post Slug is required."
			},
			{
				"key": "title",
				"type": "required",
				"message": "Post Title is required."
			}
		]
	}
}
```

The error object contains the following:

- `key`: the key that contains an error.
- `type`: the type of validation error, for example `required` means it's missing from the body and `email` would mean
  that an invalid email address has been provided.
- `message`: contains a brief message to illustrate what is required for a successful `POST`, this can be especially
  useful for display information back to the user.

## Pagination

By default, Verbis limit's all `GET` requests to 15 by default. The `Pagination` object appears under the `meta` key in
the response for browse endpoints.

```json
"meta": {
"pagination": {
"page": 1,
"pages": 1,
"limit": 15,
"total": 11,
"next": false,
"prev": false
}
}
```

The pagination object contains the following:

- `page`: the current page.
- `pages`: how many pages there are in total.
- `limit`: how many items are displayed at once.
- `total`: how many items there are in total.
- `next`: is either a boolean set to `false` if there is no next page, or an integer of page number of if there is.
- `prev`: is either a boolean set to `false` if there is no previous page, or an integer of page number of if there is.

## Endpoints

Below is a table listing the current endpoints for each resource that's available in Verbis. A brief description is
included detailing what the endpoint can do.

| Resource            | Methods                              | Description                                                    |
| ------------------- | ------------------------------------ | -------------------------------------------------------------- |
| Theme               | Read                                 | Retrieves the theme's configuration file.                      |
| Templates           | Browse                               | Retrieves all page templates for the current theme.            |
| Layouts             | Browse                               | Retrieves all page layouts for the current theme.              |
| Posts               | Browse, Read, Add, Edit, Delete      | Allows for the modification and reading of posts.              |
| Categories          | Browse, Read, Add, Edit, Delete      | Allows for the modification and reading of categories.         |
| Media               | Browse, Read, Upload, Edit Delete    | Allows media to be uploaded and read.                          |
| Users               | Browse, Add, Edit, Delete            | Allows for the modification and reading of users.              |
| Options             | Browse, Add, Edit                    | Allows to add or edit an option.                               |
| Fields              | Browse                               | Retrieves page layouts based on query parameters.              |

# TODO

// Set page to 1 if the user has passed "?limit=all"

### Auth

### Site

### Theme

The `/theme` endpoint is used to retrieve the theme's `config.yml` file within the theme directory. This can be
particularly useful for establishing what resources the current theme has, and general information about the currently
activated theme including a title, description and theme version.

**Example Response:**

👉 `GET` to `/api/{version}/theme`

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained theme config",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
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

### Templates

The `/templates` endpoint is used to retrieve the theme's all page templates for the currently activated theme that
reside in the templates' folder set in the `config.yml`.

- `key`: represents the page template file name.
- `name`: is a friendly name for the page template.

**Example Response:**

👉 `GET` to `/api/{version}/templates`

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

### Layouts

The `/layoutds` endpoint is used to retrieve the theme's all page layouts for the currently activated theme that reside
in the layouts' folder set in the `config.yml`.

- `key`: represents the page layout file name.
- `name`: is a friendly name for the page layout.

**Example Response:**

👉 `GET` to `/api/{version}/layouts`

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

### Posts

Posts are the main entity of Verbis, and it contains vital data to use for theme development and filtering through content via the API.

#### The Post object.

**Example Response:**

👉 `GET` to `/api/{version}/posts`

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained posts",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": [
		{
			"post": {
				"id": 1,
				"uuid": "648f289b-6997-46df-a44f-d16eee44d0c8",
				"slug": "/posts/post-title",
				"title": "Post Title",
				"status": "published",
				"resource": "posts",
				"page_template": "default",
				"layout": "main",
				"codeinjection_head": "",
				"codeinjection_foot": "",
				"published_at": "2021-01-11T10:42:16Z",
				"created_at": "2021-01-11T10:42:26Z",
				"updated_at": "2021-01-11T10:42:26Z",
				"options": {
					"meta": {
						"twitter": {},
						"facebook": {}
					},
					"seo": {
						"public": false,
						"exclude_sitemap": false,
						"canonical": null
					}
				}
			},
			"author": {
				"id": 1,
				"uuid": "d83f45c1-1a92-4cff-88c7-545c2017cb7b",
				"first_name": "Verbis",
				"last_name": "CMS",
				"email": "hello@verbiscms.com",
				"facebook": null,
				"twitter": null,
				"linked_in": null,
				"instagram": null,
				"biography": null,
				"profile_picture_id": null,
				"role": {
					"id": 6,
					"name": "Owner",
					"description": "The user is a special user with all of the permissions as an Administrator however they cannot be deleted"
				},
				"email_verified_at": null,
				"created_at": "2021-01-11T10:41:43Z",
				"updated_at": "2021-01-11T10:41:43Z"
			},
			"category": {
				"id": 1,
				"slug": "tech",
				"name": "Tech",
				"description": "Technology category for posts.",
				"resource": "posts",
				"parent_id": null,
				"updated_at": "2021-01-11T10:42:12Z",
				"created_at": "2021-01-11T10:42:12Z"
			},
			"layout": [
				{
					"uuid": "6a4d7442-1020-490f-a3e2-436f9135bc72",
					"title": "Text Group",
					"fields": [
						{
							"uuid": "39ca0ea0-c911-4eaa-b6e0-67dfd99e5735",
							"label": "Normal",
							"name": "text",
							"type": "text",
							"instructions": "Add a text field",
							"required": false,
							"conditional_logic": null,
							"wrapper": {
								"width": 100
							},
							"options": {
								"append": "",
								"default_value": "",
								"maxlength": "20",
								"placeholder": "Placeholder",
								"prepend": ""
							}
						}
					]
				}
			],
			"fields": [
				{
					"uuid": "39ca0ea0-c911-4eaa-b6e0-67dfd99e5735",
					"type": "text",
					"name": "text",
					"key": "",
					"value": "My text field"
				}
			]
		}
	]
}
```

#### The Post object.

### Fields

### Categories

### Media

### Options

### Users

### Roles

### Cache