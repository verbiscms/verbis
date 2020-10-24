# Templates

## Fields

### getField

Returns the value of a specific field specified in the layout.

#### Returns:

`{{ mixed }}` The field value or `an empty string if the function did not find the field.

#### Examples:

**Get the value**

Obtain the value of a field named `text` 

```
{{ getField "text" }}
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






