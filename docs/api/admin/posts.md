# Posts

Posts are the main entity of Verbis, and it contains vital data to use for theme development and filtering through
content via the API. All endpoints are prefixed with `/posts`.

## The Post object

When you retrieve posts from the API, a Post object will be returned that holds information about the post. It contains
the following:

- `post`: contains details about the post including the `slug`, `title`, page attributes such as `page_template` and any
  options that are attached including SEO and Meta information.
- `author`: contains the post category author, including the `first_name`, `last_name` and `role`.
- `category`: contains the post category information, including the `slug`, `name` and `description`.
- `layout`: contains the page layout as an array of field groups, these are key settings for the `fields`.
- `fields`: contains an array of fields that are associated with the layout.

👉 `GET` to `/api/{version}/posts`

**Example Response:**

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
				"published_at": "2020-01-01T12:00:00Z",
				"created_at": "2020-01-01T12:00:00Z",
				"updated_at": "2020-01-01T12:00:00Z",
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
				"created_at": "2020-01-01T12:00:00Z",
				"updated_at": "2020-01-01T12:00:00Z"
			},
			"category": {
				"id": 1,
				"slug": "tech",
				"name": "Tech",
				"description": "Technology category for posts.",
				"resource": "posts",
				"parent_id": null,
				"created_at": "2020-01-01T12:00:00Z",
				"updated_at": "2020-01-01T12:00:00Z"
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
								"maxlength": "",
								"placeholder": "",
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

## Retrieving Posts

The Posts endpoint allows you to filter through posts with query parameters shown below.

| Query Parameter     | Description                                                                 |
| ------------------- | --------------------------------------------------------------------------- |
| Resource            | The posts resource, `posts` or `pages` if there is no resource attached.    | 
| Status              | The status of the post, `published`, `draft` or `private`.                  |

#### Examples

**Retrieve all posts with the resource of `news`**

🔗 `GET` to `/api/{version}/posts?resource=news`

**Retrieve all posts with a status of `published`**

🔗 `GET` to `/api/{version}/posts?status=published`

**Retrieve all posts with a status of `published` and with a limit of `10`**

🔗 `GET` to `/api/{version}/posts?status=published&limit=10`

## Retrieve a specific Post

To retrieve a specific post, an ID parameter is passed after `/posts` URL, the following URL will retrieve the post with
an ID of 1.

👉 `GET` to `/api/{version}/posts/1`

## Creating a post

To create a post a `slug` and `title` is required to avoid any collisions and detect if the slug is already being used
by an existing post. If no author ID is passed, the owner will automatically be assigned. You can optionally pass a
category ID.

Required fields:
- `slug`
- `title`

Below is a minimal example of creating a post.

👉 `POST` to `/api/{version}/posts`

```json
{
	"slug": "/new-post-title",
	"title": "My awesome new post",
	"author": 1,
	"category": 1
}
```

## Updating a Post

To update a post, an ID parameter is passed after the `/posts` URL, the body is exactly the same as the `POST` route. If
no post is found with the given ID, a response of 400 will be returned.

👉 `PUT` to `/api/{version}/posts/1`

```json
{
	"title": "My awesome new post with a changed title",
	"author": 1,
	"category": 1
}
```

## Deleting a Post

To delete a post, an ID parameter is passed after the `/posts` URL. If no post is found with the given ID a response of
400 will be returned.

👉 `DELETE` to `/api/{version}/posts/1`