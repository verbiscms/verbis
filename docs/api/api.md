# Admin API

The Verbis API allows you to create and manage your content easily using wide variety of different endpoints. Everything
you are able to do in the admin interface you are able to achieve with the admin API, which makes it incredibly flexible
and dynamic.

All the Verbis endpoints follow REST principles, if you've interacted with a RESTful API already, many of the concepts
will be familiar to you. Requests and responses are JSON encoded throughout with different query parameters available to
filter and search through content easily and quickly.

___

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

___

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

___

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

___

## Filtering

When requesting a list of resources via the API through browse endpoints, you can apply filters to search through the array of entities.

`/posts&filter={"resource":[{"operator":"=", "value":"verbis"}]}`

In the above url:

- `resource` can be any property attached to the object.
- `operator`: is an allowed operator detailed in the table below.
- `value`: is a value whose type corresponds to the allowed type detailed below.

This call will search all posts that are equal to verbis.

### Available operators:

You can search through a list of entities in Verbis with the following operators:

| Operator      | Description                                                             |
| ------------- | ----------------------------------------------------------------------- |
| `=`           | Equal to.                                                               |
| `>`           | Greater than.                                                           |
| `>=`          | Greater than or equal to.                                               |
| `<`           | Less than.                                                              |
| `<=`          | Less than or equal to.                                                  |
| `<>`          | Not equal to.                                                           | 
| `LIKE`        | True if the operand matches a pattern.                                  |
| `IN`          | True if the operand is equal to one of a list of expressions.           |
| `NOT LIKE`    | True if the operand does not match a pattern.                           |

___

// Set page to 1 if the user has passed "?limit=all"

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
| Media               | Browse, Read, Upload, Edit, Delete   | Allows media to be uploaded and read.                          |
| Users               | Browse, Add, Edit, Delete            | Allows for the modification and reading of users.              |
| Options             | Browse, Add, Edit                    | Allows to add or edit an option.                               |
| Fields              | Browse                               | Retrieves page layouts based on query parameters.              |



## Auth

## Site

The `/site` endpoint is used to retrieve the global Site object which contains important information about the Verbis installation.

- The `title`, `description`, `logo`, `url` can all be updated in the admin interface.
- The `url` contains the current version of Verbis.

**Example Response:**

👉 `GET` to `/api/{version}/site`

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained site config",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {
		"title": "Verbis",
		"description": "A Verbis website. Publish online, build a business, work from home",
		"logo": "/verbis/images/verbis-logo.svg",
		"url": "http://127.0.0.1:8080",
		"version": "0.0.1"
	}
}
```

## Theme

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

## Templates

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

## Layouts

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

## Posts

Posts are the main entity of Verbis, and it contains vital data to use for theme development and filtering through
content via the API.

### The Post object

When you retrieve posts from the API, a Post object will be returned that holds information about the post. It contains
the following:

- `post`: contains details about the post including the `slug`, `title`, page attributes such as `page_template` and any
  options that are attached including SEO and Meta information
- `author`: contains the post category author, including the `first_name`, `last_name` and `role`.
- `category`: contains the post category information, including the `slug`, `name` and `description`.
- `layout`: contains the page layout as an array of field groups, these are key settings for the `fields`.
- `fields`: contains an array of fields that are associated with the layout.

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
				"uuid": "266c4f82-53fb-11eb-ae93-0242ac130002",
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
				"uuid": "266c51ee-53fb-11eb-ae93-0242ac130002",
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
					"uuid": "266c52d4-53fb-11eb-ae93-0242ac130002",
					"title": "Text Group",
					"fields": [
						{
							"uuid": "266c53a6-53fb-11eb-ae93-0242ac130002",
							"label": "Normal",
							"name": "text",
							"type": "text",
							"instructions": "Add a text field.",
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
					"uuid": "266c548c-53fb-11eb-ae93-0242ac130002",
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

___

### Retrieving Posts

The Posts endpoint allows you to filter through posts with query parameters and [filtering](#Filtering) shown below.

| Parameter           | Description                                                                 |
| ------------------- | --------------------------------------------------------------------------- |
| Resource            | The posts resource, `posts` or `pages` if there is no resource attached.    | 
| Status              | The status of the post, `published`, `draft` or `private`.                  |

#### Examples

**Retrieve all posts with the resource of `news`**

🔗 `GET` to `/posts?resource=news`

**Retrieve all posts with a status of `published`**

🔗 `GET` to `/posts?status=published`

**Retrieve all posts with a status of `published` and with a limit of `10`**

🔗 `GET` to `/posts?status=published&limit=10`

### Retrieve a specific Post

To retrieve a specific post, an ID parameter is passed after `/posts` URL, the following URL will retrieve the post with
an ID of 10.
___

### Creating a post

To create a post a slug and title is required to avoid any collisions and detect if the slug is already being used by an
existing post. If no author ID is passed, the owner will automatically be assigned. You can optionally pass a category
ID.

Required fields:

- `slug`
- `title`

Below is a minimal example of creating a post.

👉 `POST` to `/posts`

```json
{
	"slug": "/new-post-title",
	"title": "My awesome new post",
	"author": 1,
	"category": 1
}
```

___

### Updating a Post

To update a post, an ID parameter is passed after the `/posts` URL, the body is exactly the same as the `POST` route. If
no post is found with the given ID a response of 400 will be returned.

👉 `PUT` to `/posts/1`

```json
{
	"title": "My awesome new post with a changed title",
	"author": 1,
	"category": 1
}
```

___

### Deleting a Post

To delete a post, an ID parameter is passed after the `/posts` URL. If no post is found with the given ID a response of
400 will be returned.

👉 `DELETE` to `/posts/1`

___

## Categories

___

## Media

___

## Fields

___

## Options

___

## Users

___

## Roles

___

## Cache

___