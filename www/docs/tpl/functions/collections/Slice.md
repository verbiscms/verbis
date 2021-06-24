# Slice

As there is no built in slice (array) types GoLang's templating engine, the `Slice` functions
below help you to create and manipulate collections of data easily.
As slices are a collection of `interface{}`'s, mixed values can be used, but proceed with caution.

**Reminder:** you are able to access any value with the `index` property.
For example, `{{ index $mySlice 0 }}` will output the first value of the slice.

___

## slice

Creates a slice (array) of passed arguments.

### Accepts: 

`i ...interface{}` The values to be converted.

### Returns:

`[]interface` The created slice.

### Examples:

**Create a slice of words**

This example demonstrates how to create a slice of words and assign it to the
`$mySlice` variable. The `index` function is used to output the third word.
Returns `!`

```gotemplate
{{ $mySlice := slice "hello" "world" "!" }}
{{ index $mySlice 2 }}
```

**Create a slice of integers and range over them**

This example demonstrates how to create a slice of integers and assign it to the
`$numbers` variable. A range loop is then used to loop over the slice and output 
the value.
Returns `12345`

```gotemplate
{{ $numbers := slice 1 2 3 4 5 }}
{{ range $number := $numbers }}
    {{ . }}
{{ end }}
```

___

## append

Adds and element to the end of the slice.

### Accepts: 

`slice interface{}, append interface{}` The slice and value to be appended.

### Returns:

`[]interface, error` The slice with the appended value or an error if the
slice passed is not of type `Slice` or `Array`.

### Examples:

**Append to a slice of integers**

Returns `1234`.

```gotemplate
{{ $mySlice := slice 1 2 3 }}
{{ append $mySlice 4 }}
```

___

## prepend

Adds and element to the beginning of the slice.

### Accepts: 

`slice interface{}, prepnd interface{}` The slice and value to be prepended.

### Returns:

`[]interface, error` The slice with the prepended value or an error if the
slice passed is not of type `Slice` or `Array`.

### Examples:

**Prepend to a slice of integers**

Returns `4123`

```gotemplate
{{ $mySlice := slice 1 2 3 }}
{{ prepend $mySlice 4 }}
```

___

## first

Retrieves the first element of the slice.

### Accepts: 

`slice interface{}` The slice.

### Returns:

`inetface{}, error` The first value of the slice or an error if the
slice passed is not of type `Slice` or `Array`.

### Examples:

**Retrieve the first element of a slice of integers**

Returns `1`

```gotemplate
{{ $mySlice := slice 1 2 3 }}
{{ first $mySlice }}
```

___

## last

Retrieves the last element of the slice.

### Accepts: 

`slice interface{}` The slice.

### Returns:

`inetface{}, error` The last value of the slice or an error if the
slice passed is not of type `Slice` or `Array`.

### Examples:

**Retrieve the last element of a slice of integers**

Returns `3`

```gotemplate
{{ $mySlice := slice 1 2 3 }}
{{ last $mySlice }}
```

___

## reverse

Reverses the slice.

### Accepts: 

`slice interface{}` The slice to be reversed.

### Returns:

`interface{}, error` The reversed slice or an error if the
slice passed is not of type `Slice` or `Array`.

### Examples:

**Reverse a slice of integers**

Returns `321`

```gotemplate
{{ $mySlice := slice 1 2 3 }}
{{ reverse $mySlice }}
```