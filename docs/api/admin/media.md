# Media

These routes allow you to browse, read, upload, update and delete media items through the Verbis API. All media
endpoints are prefixed with `/media`.

Using these routes you can access all media items in the library and uploads different content for use with your
Verbis installation. 

## The Media object

When you retrieve media items from the API, an array of Media object's will be returned that holds information about the
media item. It contains the following:

- `url`: the public url of the media item.
- `title`, `alt` and `description`: contains useful meta information about the media item.
- `file_size`: the size (in bytes) of the item.
- `file_name`: the public file name of the item.
- `sizes`: contains an object of sizes that are specified within the `Media Options` of the admin interface, along with
  `size_name`.
- `type`: the MIME type of the file.

👉 `GET` to `/api/{version}/media`

**Example Response:**

```json

{
	"status": 200,
	"error": false,
	"message": "Successfully obtained media",
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
			"uuid": "a9322f12-541a-11eb-ae93-0242ac130002",
			"url": "/uploads/2021/01/verbis-logo.jpg",
			"title": null,
			"alt": null,
			"description": null,
			"file_size": 182674,
			"file_name": "verbis-logo.jpg",
			"sizes": {
				"hd": {
					"uuid": "a932314c-541a-11eb-ae93-0242ac130002",
					"url": "/uploads/2021/01/verbis-logo-1920x0.jpg",
					"name": "verbis-logo-1920x0.jpg",
					"size_name": "HD Size",
					"file_size": 212,
					"width": 1920,
					"height": 0,
					"crop": false
				},
				"large": {
					"uuid": "a932348a-541a-11eb-ae93-0242ac130002",
					"url": "/uploads/2021/01/verbis-logo-1280x0.jpg",
					"name": "verbis-logo-1280x0.jpg",
					"size_name": "Large Size",
					"file_size": 114,
					"width": 1280,
					"height": 0,
					"crop": false
				},
				"medium": {
					"uuid": "a93235a2-541a-11eb-ae93-0242ac130002",
					"url": "/uploads/2021/01/verbis-logo-992x0.jpg",
					"name": "verbis-logo-992x0.jpg",
					"size_name": "Medium Size",
					"file_size": 76,
					"width": 992,
					"height": 0,
					"crop": false
				},
				"thumbnail": {
					"uuid": "a932366a-541a-11eb-ae93-0242ac130002",
					"url": "/uploads/2021/01/verbis-logo-550x300.jpg",
					"name": "verbis-logo-550x300.jpg",
					"size_name": "Thumbnail Size",
					"file_size": 28,
					"width": 550,
					"height": 300,
					"crop": true
				}
			},
			"type": "image/jpeg",
			"user_id": 1,
			"created_at": "2020-01-01T12:00:00Z",
			"updated_at": "2020-01-01T12:00:00Z"
		}
	]
}
```

## Retrieve a specific Media Item

To retrieve a specific media item, an ID parameter is passed after `/media` URL, the following URL will retrieve a media
item with an ID of 1.

👉 `GET` to `/api/{version}/media/1`

## Uploading a Media Item

Media items are only permitted to be uploaded one at a time, and encoded as multipart form data, the media item should
reside under the `file` key. The API will validate MIME types, file sizes, maximum image widths and heights. If
validation failed, a 400 response will be sent. If the Admin options allow for it, images will also be converted
to `webp` file formats.

⚠️ This endpoint should be encoded as `multipart/form-data`

Required fields:

- `file`

👉 `POST` to `/api/{version}/media`

## Updating a Media Item

To update a media item, an ID parameter is passed after the `/media` URL. Whilst you can't update any of the paths or
the core media item properties, you are able to edit meta data as shown below.

👉 `PUT` to `/api/{version}/media/1`

```json
{
	"title": "Verbis",
	"alt": "Verbis Logo",
	"description": "A Verbis CMS"
}
```

## Deleting a Media Item

To delete a media item from the library, an ID parameter is passed after the `/media` URL. If no media item is found
with the given ID a response of 400 will be returned.

👉 `DELETE` to `/api/{version}/media/1`