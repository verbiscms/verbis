# Templates

# Fields

Short introduction here

___

## getField

Gets the value of a specific field specified in the layout.

### Returns:

`interface{}` The field value or an empty string if the function did not find the field.

### Examples:

**Get the value from the current post**

Obtain the value of a field named `text` 

```
{{ getField "text" }}
```

**Get the value from a specific post**

Obtain the value of a field named `text` from the post with the ID of 10

```
{{ getField "text" 10 }}
```

**Check if a field exists:**

As `getField` returns an empty string if the field is not set, you can use it to see if the value exists.
See also `hasField`.

```
{{ if getField "text" }}
     The field named 'text' exists...
{{ end }}
```

**Assign to a variable:**

You can also assign the contents of a field to a variable to be used later on in the template.

```
{{ $text := getField "text" }}
{{ $text }}
```

___

### hasField

Checks if a field value exists.

#### Returns:

`boolean` True if the field exists

#### Examples:

**Check the value**

See if the field `text` exists.

```
{{ if hasField "text" }}
     The field named 'text' exists...
{{ end }}
```

___

### getFields

Get all fields associated with a post. This function is especially useful for debugging.

#### Returns:

`map[string[interface{}` The field values, or nil if there are no fields set.

#### Examples:

**Output all fields from the current post**

This example demonstrates how to output all the fields in the current post.

```
{{ getFields }}
```

**Output all fields from a specific post**

This example demonstrates how to output all the fields with the post ID of 10.

```
{{ getFields 10 }}
```

___

### getRepeater

Get a repeater block associated

#### Returns:

`[]map[string[interface{}` The field values, or nil if there are no fields set.




