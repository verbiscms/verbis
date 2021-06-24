# Fields

Fields are the building blocks of a Verbis theme, it's important to grasp the foundations on how they work to,
accelerate your theme building. The fields value will be automatically rendered to the fields value, meaning
if the field is a post, media item, or  user, their respective types will be returned allowing you to filter
through rich data, quickly.

___

## field

Get the value of a specific field specified in the layout.

### Accepts: 

`string, agrs (optional)` The field name & optional post ID or Post.

### Returns:

`interface{}` The field value or nil if the function did not find the field.

### Examples:

**Get the value from the current post**

Obtain the value of a field named `text`.

```gotemplate
{{ field "text" }}
```

**Get the value from a specific post**

Obtain the value of a field named `text` from the post with the ID of 10.

```gotemplate
{{ field "text" 10 }}
```

**Check if a field exists:**

As `field` returns nil if the field is not set, you can use it to see if the value exists.
See also `hasField`.

```gotemplate
{{ if field "text" }}
     <!-- The field named 'text' exists. -->
{{ end }}
```

Or you can use with to re-assign the context to the magic `dot`.

```gotemplate
{{ with field "text" }}
    {{ . }}
{{ end }}
```

**Assign to a variable:**

You can also assign the contents of a field to a variable to be used later on in the template.

```gotemplate
{{ $text := field "text" }}
{{ with $text }}
    {{ . }}
{{ end }}
```
___

## fields

Get all fields associated with a post. This function is especially useful for debugging.

### Accepts: 

`integer (optional)` Optional post ID.

### Returns:

`map[string[interface{}` The field values, or nil if there are no fields set.

### Examples:

**Output all fields from the current post**

This example demonstrates how to output all the fields in the current post.

```gotemplate
{{ fields }}
```

**Output all fields from a specific post**

This example demonstrates how to output all the fields with the post ID of 10.

```gotemplate
{{ fields 10 }}
```
___


## fieldObject

Gets the raw object of a specific field before the value is rendered. This cab be especially useful
for debugging.

### Accepts:

`string, args (optional)` The field name & optional post ID.

### Returns:

`PostField` The field object.

### Examples:

**Debug a field**

This example demonstrates how to output all the fields with the post ID of 10.

```gotemplate
{{ dump (field "text") }}
```
___
