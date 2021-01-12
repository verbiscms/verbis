# Fields

The `/fields` endpoint is used to obtain field layouts for a post. As field layouts can change dependent on the
different post properties, field layout's can differ and should be updated to reflect the users choice. This endpoint
accepts a wide variety of different query parameters to narrow down which field layouts should be displayed to the user.

| Query Parameter     | Description                                                 |
| ------------------- | ----------------------------------------------------------- |
| `resource`          | The post resource.                                          |
| `page_template`     | The page template.                                          |
| `layout`            | The page layout.                                            |
| `user_id`           | The user ID associated with the post.                       |
| `category_id`       | The category ID associated with the post.                   |

👉 `GET` to `/api/{version}/fields`

**Example Response:**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully obtained fields",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": [
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
	]
}
```

## Examples

**Retrieve field layouts that are assigned to a resource of `news`**

🔗 `GET` to `/api/{version}/fields?resource=news`

**Retrieve field layouts that are assigned to a page_template of `archive`**

🔗 `GET` to `/api/{version}/fields?page_template=archive`