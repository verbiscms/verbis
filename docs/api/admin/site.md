# Site

The `/site` endpoint is used to retrieve the global Site object which contains important information about the Verbis
installation.

Apart from the `Auth` routes, this is the **only endpoint that does not require authentication**.

- The `title`, `description`, `logo`, `url` can all be updated in the admin interface.
- The `url` contains the current version of Verbis.

👉 `GET` to `/api/{version}/site`

**Example Response:**

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