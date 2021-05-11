

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

