# Users

These routes allow you to browse, read, create, update and delete Users through the Verbis API. All endpoints are
prefixed with `/users`.

## The User object

When you retrieve categories from the API, an array of User objects will be returned that holds information about the
user.

👉 `GET` to `/api/{version}/users`

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
		}
	]
}
```

## Retrieve a specific User

To retrieve a specific user, an ID parameter is passed after `/users` URL, the following URL will retrieve a user with
an ID of 1.

👉 `GET` to `/api/{version}/users/1`

## Creating a User

To create a user multiple fields are required as shown below. A `role_id` needs to be attached to the request to assign
the newly created user a role. A response 400 will be returned if the user already exists or validation failed.

Required fields:

- `first_name`
- `last_name`
- `email`
- `role_id`
- `password`
- `confirm_password`

Below is a minimal example of creating a user.

👉 `POST` to `/api/{version}/categories`

```json
{
	"first_name": "Verbis",
	"last_name": "CMS",
	"email": "hello@verbiscms.com",
	"role_id": 2,
	"password": "mypassword",
	"confirm_password": "mypassword"
}
```

## Updating a User

To update a user, an ID parameter is passed after the `/users` URL, the body is exactly the same as the `POST`
route. If no user is found with the given ID, a response of 400 will be returned.

👉 `PUT` to `/api/{version}/users/1`

```json
{
	"first_name": "CMS"
}
```

## Deleting a User

To delete a user, an ID parameter is passed after the `/users` URL. If no user is found with the given ID a response of
400 will be returned.

👉 `DELETE` to `/api/{version}/users/1`