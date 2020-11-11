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

## getRepeater

Get a slice of repeater blocks of a specific field.

### Accepts: 

`string` The field name.

### Returns:

`[]map[string[interface{}` The repeater field values, or `nil` if there are no repeaters with the given name.

### Examples:

**Loop through a repeater block**

This example demonstrates how to loop over a repeater block in the template. 

First the value of the repeater field is assigned to a variable and `if` is used to establish if there are rows within the repeater field.
A loop is used to range over the repeater fields.

To obtain the value of the repeater field, pass index, the content (dot `.`) and the name of the field.
If the field does not exist, it will gracefully return an empty string.

```gotemplate
{{ $repeaterFields := getRepeater "repeater" }}
{{ if $repeaterFields }}
    {{ range $index, $field := $repeaterFields }}
        {{ index . "text" }}
        {{ index . "richtext" }}
    {{ end }}
{{ end }}
```
___

## getFlexible

Get the map of flexible content of a specific field.

### Accepts: 

`string` The field name.

### Returns:

`map[string[interface{}` The flexible content field values, or `nil` if there are no flexible content fields with the given name.

### Examples:

**Loop through flexible content**

This example demonstrates how to loop over a flexible content block in the template.

First the value of the flexible content field is assigned to a variable and `if` is used to establish if there are rows within the flexible content field.
A loop is used to range over the blocks.

The field type (layout) is stored in `$field.type`. The `if eq` function can be used to compare the layouts and output content based on the layout.

See `getSubField` for more information on how to output the layouts fields.

```gotemplate
{{ $flexibleContent := getFlexible "flexible" }}
{{ if $flexibleContent }}
    {{ range $layout := $flexibleContent }}

        {{ if eq $layout.type "text_block" }}
            <h2>{{ getSubField "heading" $field }}</h2>
            <p>{{ getSubField "text" $field }}</p>
        {{ end }}

        {{ if eq $layout.type "card_section" }}
            <h3>{{ getSubField "title" $field }}</h3>
            <p>{{ getSubField "content" $field }}</h3>
        {{ end }}

    {{ end }}
{{ end }}
```

**Using partials with flexible content**

Sometimes you may want to use different template files for different layouts. The example below shows how.

#### flexible-content.cms
In your master template, the same principle applies. Simply loop over the flexible content and compare the layouts using `$layout.type`

It's important to pass the context of the loop using the dot `.`. If the post data is required in the partial, pass the dollar symbol `$`.
This will enable you to still access the post data in the child template.

See `partial` for more details on including child templates.

```gotemplate
{{ $flexibleContent := getFlexible "flexible" }}
{{ if $flexibleContent }}
    {{ range $layout := $flexibleContent }}
        {{ if eq $layout.type "text_block" }}
            {{ partial "blocks/text-block.cms" . }}
        {{ end }}
    {{ end }}
{{ end }}
```

#### blocks/test.cms
In the child template, it is important to assign the layout to a variable using `index . {index-number}`. 
The index number correlates to where you passed the context (dot `.`) in the calling partial function. If you passed the global context before the loop context 
`{{ partial "blocks/text-block.cms" $ . }}`, the index would be 1.
From there, it's a simple case of using the `getSubField` function.

```gotemplate
{{ $layout := index . 0 }}
<h1>{{ getSubField "text" $layout }}</h1>
```
___

## getSubField

Get the value of a sub field within a flexible content layout.

### Accepts: 

`string, string` The field & layout name.

### Returns:

`interface{}` The field value or an empty string if the function did not find the field.

### Examples:

**Get the field value of a layout**

Get the value of a field named `"content"` in the layout being ranged.

```gotemplate
{{ getSubField "content" $layout }}
```
___

