# Filtering

When requesting a list of resources via the API through browse endpoints, you can apply filters to search through the
array of entities.

`/posts&filter={"resource":[{"operator":"=", "value":"verbis"}]}`

In the above url:

- `resource` can be any property attached to the object.
- `operator`: is an allowed operator detailed in the table below.
- `value`: is a value whose type corresponds to the allowed type detailed below.

This call will search all posts that are equal to verbis. You can use `AND`'s for each parameter by appending another
object to the filter array.

## Available operators:

You can search through a list of entities in Verbis with the following operators:

| Operator      | Description                                                             |
| ------------- | ----------------------------------------------------------------------- |
| `=`           | Equal to.                                                               |
| `>`           | Greater than.                                                           |
| `>=`          | Greater than or equal to.                                               |
| `<`           | Less than.                                                              |
| `<=`          | Less than or equal to.                                                  |
| `<>`          | Not equal to.                                                           | 
| `LIKE`        | True if the operand matches a pattern.                                  |
| `IN`          | True if the operand is equal to one of a list of expressions.           |
| `NOT LIKE`    | True if the operand does not match a pattern.                           |

## Examples

**Filter through posts by title**

This example demonstrates how to search through posts with a title that iS `LIKE` `verbis`.

`/posts&filter={"title":[{"operator":"LIKE", "value":"verbis"}]}`

**Filter through posts by title and page template**

This example demonstrates how to search through posts with a title that is `LIKE` `verbis` OR if the post has a page
template `LIKE` `archive`.

`/posts&filter={"title":[{"operator":"LIKE", "value":"verbis"}], "page_template":[{"operator":"LIKE", "value":"archive"}]}`

**Filter through posts with and conditional**

This example demonstrates how to search through posts with a slug that is `LIKE` `verbis` AND `LIKE` `cms`.

`/posts&filter={"slug":[{"operator":"LIKE", "value":"verbis"},{"operator":"LIKE", "value":"cms"}]}`