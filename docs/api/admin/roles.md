# Roles

The `/roles` endpoint is used to retrieve the all roles that a user can be assigned to. This is especially useful for
the `POST` endpoint to create new users, as you can display what roles thew user can be assigned to before creating a
new user.

- `id`: the ID of the role stored in the database.
- `name`: the name of the role.
- `description`: a brief description of the role and the role's permissions.

👉 `GET` to `/api/{version}/roles`

**Example Response:**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained user roles",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": [
		{
			"id": 1,
			"name": "Banned",
			"description": "The user has been banned from the system."
		},
		{
			"id": 2,
			"name": "Contributor",
			"description": "The user can create and edit their own draft posts, but they are unable to edit drafts of users or published posts."
		},
		{
			"id": 3,
			"name": "Author",
			"description": "The user can write, edit and publish their own posts."
		},
		{
			"id": 4,
			"name": "Editor",
			"description": "The user can do everything defined in the Author role but they can also edit and publish posts of others, as well as their own."
		},
		{
			"id": 5,
			"name": "Administrator",
			"description": "The user can do everything defined in the Editor role but they can also edit site settings and data. Additionally they can manage users"
		},
		{
			"id": 6,
			"name": "Owner",
			"description": "The user is a special user with all of the permissions as an Administrator however they cannot be deleted"
		}
	]
}
```


