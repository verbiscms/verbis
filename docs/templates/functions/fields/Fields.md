# Fields

Fields are the building blocks of a Verbis theme, it's important to grasp the foundations on how they work to,
accelerate your theme building. Usually interface is rednered in the template when obtaining a field, however
if the field is a post, media item, or  user, their respective types will be returned allowing you to filter
through rich data, quickly.

___

## getField

Get the value of a specific field specified in the layout.

### Accepts: 

`string, integer (optional)` The field name & optional post ID.

### Returns:

`interface{}` The field value or an empty string if the function did not find the field.

### Examples:

**Get the value from the current post**

Obtain the value of a field named `text`.

```gotemplate
{{ getField "text" }}
```

**Get the value from a specific post**

Obtain the value of a field named `text` from the post with the ID of 10.

```gotemplate
{{ getField "text" 10 }}
```

**Check if a field exists:**

As `getField` returns an empty string if the field is not set, you can use it to see if the value exists.
See also `hasField`.

```gotemplate
{{ if getField "text" }}
     The field named 'text' exists...
{{ end }}
```

**Assign to a variable:**

You can also assign the contents of a field to a variable to be used later on in the template.

```gotemplate
{{ $text := getField "text" }}
{{ $text }}
```
___

## hasField

Check if a field value exists.

### Accepts: 

`string` The field name.

### Returns:

`bool` True if the field exists.

### Examples:

**Check the value**

See if the field `text` exists.

```gotemplate
{{ if hasField "text" }}
     The field named 'text' exists...
{{ end }}
```
___

## getFields

Get all fields associated with a post. This function is especially useful for debugging.

### Accepts: 

`integer (optional)` Optional post ID.

### Returns:

`map[string[interface{}` The field values, or nil if there are no fields set.

### Examples:

**Output all fields from the current post**

This example demonstrates how to output all the fields in the current post.

```gotemplate
{{ getFields }}
```

**Output all fields from a specific post**

This example demonstrates how to output all the fields with the post ID of 10.

```gotemplate
{{ getFields 10 }}
```
___

