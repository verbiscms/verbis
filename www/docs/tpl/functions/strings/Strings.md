# Strings

The string functions are a variety of different functions to help manipulate a whole or
part of a string.
___

## trim

Strips trailing white space from the beginning and end of a string.

### Accepts: 

`str string` The string to be trimmed.

### Returns:

`string` The trimmed string.

### Examples:

**Replace a string with dashes**

Returns `hello verbis`

```gotemplate
{{ trim "    hello verbis     " }}
```

___

## upper

Converts an entire string to uppercase characters.

### Accepts: 

`str string` The string to be converted.

### Returns:

`string` The uppercase string.

### Examples:

**Convert a string to uppercase characters**

Returns `HELLO VERBIS`

```gotemplate
{{ upper "hello verbis" }}
```

___

## lower

Converts an entire string to lowercase characters.

### Accepts: 

`str string` The string to be converted.

### Returns:

`string` The lowercase string.

### Examples:

**Convert a string to lowercase characters**

Returns `hello verbis`

```gotemplate
{{ lower "hELLo VERBIS" }}
```

___

## title

Converts an entire string to title case.

### Accepts: 

`str string` The string to be converted.

### Returns:

`string` The title cased string.

### Examples:

**Convert a string to title case**

Returns `Hello Verbis`

```gotemplate
{{ title "hello verbis" }}
```

___

## replace

Replaces a sub string with a new string with all matches.

### Accepts: 

`src string, old, string, new string` The source string, the old string to be replaced, and the new string.

### Returns:

`string` The replaced string.

### Examples:

**Replace a string with dashes**

Returns `hello-verbis-cms`

```gotemplate
{{ replace "hello verbis cms" "" "-" }}
```

**Replace part of a string**

Returns `hello world!`

```gotemplate
{{ replace "hello verbis cms" "verbis cms" "world!" }}
```

___

## substr

Extracts characters from a string, between two specified indices, and returns the new sub string.

### Accepts: 

`str string, start interface{} end interface{}` The source string, and the start and end indexes.

### Returns:

`string` The replaced string.

### Examples:

**Cut out part of a string**

Returns `hello`

```gotemplate
{{ substr "hello verbis" 0 5 }}
```

___

## trunc

Removes (truncates) a part of a string from the given index. Negatives can
be used.

### Accepts: 

`str string, index interface{}` The source string, and the index to be truncated.

### Returns:

`string` The truncated string.

### Examples:

**Cut out part of a string (positive)**

Returns `hello`

```gotemplate
{{ trunc "hello verbis" 5 }}
```

**Cut out part of a string (negative)**

Returns `verbis`

```gotemplate
{{ trunc "hello verbis" -5 }}
```

___

## ellipsis

Truncates a string with ellipsis `(...)` from a given length.

### Accepts: 

`str string, length int` The source string, and the index to be truncated with ellipsis.

### Returns:

`string` The truncated string.

### Examples:

**Replace a sentence with ellipsis**

Returns `hello verbis...`

```gotemplate
{{ ellipsis "hello verbis cms!" 11 }}
```
