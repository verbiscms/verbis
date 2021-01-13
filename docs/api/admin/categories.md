# Categories

These routes allow you to browse, read, create, update and delete categories through the Verbis API. All endpoints are
prefixed with `/categories`. These routes can be particularly powerful for updating and creating filtered posts

## The Category object

When you retrieve categories from the API, an array of Category objects will be returned that holds information about
the category.

👉 `GET` to `/api/{version}/categories`

**Example Response:**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained categories",
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
	"data": [
		{
			"id": 1,
			"uuid": "0eb7db58-5415-11eb-ae93-0242ac130002",
			"slug": "tech",
			"name": "Tech",
			"description": "Technology category for posts.",
			"resource": "posts",
			"parent_id": null,
			"archive_id": null,
			"created_at": "2020-01-01T12:00:00Z",
			"updated_at": "2020-01-01T12:00:00Z"
		},
		{
			"id": 2,
			"uuid": "0eb7de14-5415-11eb-ae93-0242ac130002",
			"slug": "sports",
			"name": "Sports",
			"description": "Sports category for posts.",
			"resource": "posts",
			"parent_id": null,
			"archive_id": null,
			"created_at": "2020-01-01T12:00:00Z",
			"updated_at": "2020-01-01T12:00:00Z"
		}
	]
}
```

## Retrieve a specific Category

To retrieve a specific category, an ID parameter is passed after `/categories` URL, the following URL will retrieve a
category with an ID of 1.

👉 `GET` to `/api/{version}/categories/1`

## Creating a category

To create a category a `slug`, `name` and `resource` are required, as a category has to be attached to a resource.

Required fields:

- `slug`
- `name`
- `resource`

Below is a minimal example of creating a category.

👉 `POST` to `/api/{version}/categories`

```json
{
	"slug": "/tech",
	"name": "Tech",
	"resource": "posts"
}
```

## Updating a Category

To update a category, an ID parameter is passed after the `/categories` URL, the body is exactly the same as the `POST`
route. If no category is found with the given ID, a response of 400 will be returned.

👉 `PUT` to `/api/{version}/categories/1`

```json
{
	"name": "Updated category name"
}
```

## Deleting a Category

To delete a category, an ID parameter is passed after the `/category` URL. If no category is found with the given ID a
response of 400 will be returned.

👉 `DELETE` to `/api/{version}/categories/1`