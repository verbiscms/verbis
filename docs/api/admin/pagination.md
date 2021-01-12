# Pagination

By default, Verbis limit's all `GET` requests to **15** by default. The `Pagination` object appears under the `meta` key in
the response for browse endpoints.

## The Pagination object

```json
{
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
}
```

The pagination object contains the following:

- `page`: the current page.
- `pages`: how many pages there are in total.
- `limit`: how many items are displayed at once.
- `total`: how many items there are in total.
- `next`: is either a boolean set to `false` if there is no next page, or an integer of the next page number of if there is.
- `prev`: is either a boolean set to `false` if there is no previous page, or an integer of the previous page number of if there is.

## Pagination parameters for requests

The Verbis API allows you to navigate through paginated content easily with the query parameters described below:

| Query Parameter     | Default      | Description                                                                     |
| ------------------- | ------------ | ------------------------------------------------------------------------------- |
| `page`              | `1`          | The page number of the query.                                                   |
| `limit`             | `15`         | Allows you to limit the number of records returned `all` for every record.      |
| `order_by`          | `asc`        | Order by a particular column.                                                   |
| `order_direction`   | `id`         | Establish the order direction of the query.                                     |

## Examples

**Set a custom limit and page**

This example demonstrates how to obtain posts in batches of ten, and display the second page.

`/posts?limit=10&page=2`

**Retrieve all rows**

This example demonstrates how to obtain all posts with no paging.

⚠️ **Be careful** when setting the `limit` to `all` it may slow down requests depending on how big your Verbis installation is.

`/posts?limit=all`

**Change order by and direction**

This example demonstrates how to order posts by title in descending order.

`/posts?order_by=title&order_direction=desc`