## Cache

The `/cache` endpoint is used to clear system cache in production. No body is needed for this endpoint. The following is
cleared when you post a request to `/cache`:

- Templates
- Assets
- Uploads
- Field layouts

👉 `POST` to `/api/{version}/cache`

**Example Response:**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully cleared server cache",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {}
}
```