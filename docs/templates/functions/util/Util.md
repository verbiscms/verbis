# Util

The functions below are helpers for theme development and template rendering. 
___

## len

Returns the length of a variable according to its type. If the length of the type passed
could not be retrieved, it will return `0`.

The following will be calculated from the len function:

| Type                 | Returned                                          |
| -------------------- | ------------------------------------------------- |
| Slice or Array       | The number of elements in i.                      |
| Pointer to Array     | The number of elements in *i (even if i is nil).  |                    |
| Map                  | The number of key value pairs in a map.           |
| String               | The number of bytes in i.                         |


### Accepts: 

`i interface{}` The item of length to retrieve

### Returns:

`int64` The length of the item.

### Examples:

**Obtain the length of a string**

Returns `5`

```gotemplate
{{ len "hello" }}
```

**Obtain the length of a slice**

Returns `3`

```gotemplate
{{ slice 1 2 3 | len  }}
```
___

## explode

Breaks a string into array with a delimiter (separator).
This function will return nil if there was a problem casting the parameters to strings


### Accepts: 

`delimiter interface{}, text interface{}` The explode delimiter and text to explode.

### Returns:

`[]string` The slice of strings or nil.

### Examples:

**Explode a sentence by a space**

Returns `[hello there !]`

```gotemplate
{{ explode "," "hello there !" }}
```
___

## implode

Returns a string from the elements of an array using a glue string to join them together.
This function will return nil if there was a problem casting the parameters.

### Accepts: 

`glue interface{}, slice interface{}` The explode delimiter and text to explode.

### Returns:

`string` The joined string or "" (empty string)".

### Examples:

**Implode a sentence by a comma**

Returns `[1 2 3]`

```gotemplate
{{ slice 1 2 3 | explode "," }}
```
___

## seq

Creates a sequence of integers, useful for using an integer to loop over as there is no
built-in extension for a traditional `for` loop in GoLang templates.

### Accepts:

`size interface{}` The size of the array

### Returns:

`[]int64, error` A slice of integers to range over, or an error if the size paramater
could not be cast or is less than or equal to zero.

### Examples:

**Create a loop of 5**

Returns `[1 2 3 4 5]`

```gotemplate
{{ range $val := seq 5 }}
    {{ $val }}
{{ end }}
```
___