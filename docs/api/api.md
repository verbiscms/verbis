# Admin API

The Verbis API allows you to create and manage your content easily using wide variety of different endpoints. 
Everything you are able to do in the admin interface you are able to achieve with the admin API, which makes it incredibly flexible and
dynamic. 

All the Verbis endpoints follow REST principles, if you've interacted with a RESTful API already, many of the concepts will be familiar to you.
Requests and responses are JSON encoded throughout with different query parameters available to filter and search through content easily and
quickly. 

## Structure

### Base URL:
The following URL is used for all API endpoints.

`/api/{version}/`

### Version

After the `/api` segment of the base URL, a version is prefixed with a `v`, dependent on what version of Verbis is currently being run.
For example: `/api/v1/posts`

### Endpoints

As Verbis follows RESTful principles, the following methods are accepted for each endpoint.

- `GET`: for browsing entities, `/api/{version}/posts` or reading an entity by a particular key, `/api/{version}/posts/1`
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
- `meta`: includes Pagination if the route is a `GET` (browse) endpoint. It also includes `request_time`, `response_time` and `latency` times.
- `data`: contains the main body of the call, usually this is an array, but sometimes it can be an object depending on the endpoint.


## Errors

If the resulting status code of an endpoint is a validation error, the API will respond with some useful information about what keys are
required as an object in the `data` key under `errors`.

```json
{
    "status": 400,
    "error": true,
    "message": "Validation failed",
    "meta": {
        "request_time": "2021-01-11 09:23:07.245776 +0000 UTC",
        "response_time": "2021-01-11 09:23:07.286635 +0000 UTC",
        "latency_time": "20.857ms"
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
- `type`: the type of validation error, for example `required` means it's missing from the body and `email` would mean that an invalid email address has been provided.
- `message`: contains a brief message to illustrate what is required for a successful `POST`, this can be especially useful for display information back to the user.


## Pagination

By default, Verbis limit's all `GET` requests to 15 by default. The `Pagination` object appears under the `meta` key in the response for browse endpoints.

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

Below is a table listing the current endpoints for each resource that's available in Verbis. A brief description is included detailing what the endpoint can do.

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

Brief intro
Table of endpoints


	// Set page to 1 if the user has passed "?limit=all"

### Auth

### Site

### Theme

### Templates

### Layouts

### Posts

### Fields

### Categories

### Media

### Options

### Users

### Roles

### Cache