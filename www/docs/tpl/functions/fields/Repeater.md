# Repeater

The repeater field provides a great solution for content that is repeated such as logo's, slides, team members,
cards and so forth. 

The field type acts as a parent to the child sub-fields or `rows` and has various utility functions attached to it
to gain information about the repeater. Any kind of field can be used within a repeater making the choices and
possibilities endless!

___

## repeater

Get a slice of repeater blocks of a specific field.

### Accepts:

`string` The field name.

### Returns:

`[]map[string[interface{}` The repeater field values, or `nil` if there are no repeaters with the given name.

### Options:

Various options can be used for a repeater shown below:

| Option                  | Description                                                                  |
| ----------------------- | ---------------------------------------------------------------------------- |
| Collapsed               | If the repeater is hidden by default and not visible to the end user.        |
| Minimum Rows            | The limit of a minimum amount of rows of data to be entered.                 |
| Minimum Rows            | The limit of a maximum amount of rows of data to be entered.                 |
| Button Label            | The text shown in the 'Add Row' button.                                      |


### Examples:

**Loop through a repeater block**

This example demonstrates how to loop over a repeater block in the template.

First the value of the repeater field is assigned to a variable and `HasRows` is used to establish if there are rows within the repeater field.
A loop is used to range over the repeater fields.

To obtain the value of the repeater field you can use the `SubField` method.
If the field does not exist, it will return nil, which is why it's a smart idea to use `with` to check if the field has any value before trying
to output its contents.

```gotemplate
{{ $items := repeater "blocks" }}
{{ if $items.HasRows }}
    {{ range $row := $items }}
        {{ with $row.SubField "title" }}
            <h1>{{ . }}</h1>
        {{ end }}
        {{ with $row.SubField "content" }}
            {{ safeHTML . }}
        {{ end }}
    {{ end }}
{{ end }}
```

**Nested Repeaters**

This example demonstrates how to loop over nested repeater fields. This can be confusing stuff, however Verbis aims
to simplify the process of using repeaters recursively. 

Similar to the first example, a range loop is used to loop over the parent Repeater, however to obtain the nested repeater,
we just need to use the `repeater` function along with the row's `SubField` method. This will convert all the nested data
to data that you are able to loop over.. 


```gotemplate
{{ $items := repeater "blocks" }}
{{ if $items.HasRows }}
    {{ range $row := $items }}
        {{ $nestedItems := repeater $row.SubField "nested" }}

        {{ if $nestedItems.HasRows }}
            {{ range $nestedRow := $items }}
                {{ with $nestedRow.SubField "title" }}
                    <h1>{{ . }}</h1>
                {{ end }}
            {{ end }}
        {{ end }}

    {{ end }}
{{ end }}
```


## Repeater Methods:

The functions below are attached to the `Repeater` type.

___

## HasRows

Determines if the Repeater has any rows.

### Returns:

`bool` If the repeater has rows.

### Examples:

```gotemplate
{{ $items := repeater "blocks" }}
{{ if $items.HasRows }}
    <!-- Loop? -->
{{ end }}
```

___

## Length

Returns the amount of rows within the repeater.

### Returns:

`int` The amount of repeater rows.

### Examples:

```gotemplate
{{ $items := repeater "blocks" }}
{{ $items.Length }}
```

___

## Row Methods:

The functions below are attached to the `Row` type.
___

## SubField

Returns a Sub Field by key or nil if it wasn't found.

### Accepts:

`string` The field name.

### Returns:

`interface{}` The Sub Field value.

### Examples:

```gotemplate
{{ $items := repeater "blocks" }}
{{ if $items.HasRows }}
    {{ range $row := $items }}
        {{ $row.SubField "title" }}
    {{ end }}
{{ end }}
```

___

## HasField

Determines if a row as a specific Sub Field.

### Returns:

`bool` If the row has a Sub Field.

### Examples:

```gotemplate
{{ $items := repeater "blocks" }}
{{ if $items.HasRows }}
    {{ range $row := $items }}
        {{ if $row.HasField "title" }}
            <!-- The row has a field named "title" -->
        {{ end }}
    {{ end }}
{{ end }}
```

___

## First

Returns the first element in the repeater, or nil if
the length of the repeater is zero.

### Returns:

`interface{}` The first element or Sub Field in the repeater. 

### Examples:

```gotemplate
{{ $items := repeater "blocks" }}
{{ if $items.HasRows }}
    {{ range $row := $items }}
        {{ $row.First }}
    {{ end }}
{{ end }}
```

___

## Last

Returns the last element in the repeater, or nil if
the length of the repeater is zero.

### Returns:

`interface{}` The last element or Sub Field in the repeater.

### Examples:

```gotemplate
{{ $items := repeater "blocks" }}
{{ if $items.HasRows }}
    {{ range $row := $items }}
        {{ $row.Last }}
    {{ end }}
{{ end }}
```
