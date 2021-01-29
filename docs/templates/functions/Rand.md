# Random

These functions generate random strings & numeric data but with different base character sets.
Integer, Floats, Alpha & Alphanumeric can be generated using the functions below. 

___

## randInt

Generates a random integer between two given values.

### Accepts: 

`min interface{}, max interface{}` The minimum and maximum values.

### Returns:

`int` A random integer.

### Examples:

**Generate a random number between 1 and 10**

```gotemplate
{{ randInt 1 10 }}
```

___

## randFloat

Generates a random float between two given values.

### Accepts: 

`min interface{}, max interface{}` The minimum and maximum values.

### Returns:

`float64` A random float.

### Examples:

**Generate a random float between 2.5 and 10**

```gotemplate
{{ randFloat 2.5 10 }}
```

___

## randAlpha

Generates a random alpha string using `a-zA-Z` from a given length.

### Accepts: 

`length int64` The length of the random string.

### Returns:

`string` A random alpha string.

### Examples:

**Generate a random alpha string with the length of 20**

Results in something similar to:
`EsRZeHmmFcOpsyWrNCeT`

```gotemplate
{{ randAlpha 20 }}
```

___

## randAlphaNum

Generates a random alphanumeric string using `a-zA-Z0-9` from a given length.

### Accepts: 

`length int64` The length of the random string.

### Returns:

`string` A random alphanumeric string.

### Examples:

**Generate a random alphanumeric string with the length of 20**

Results in something similar to:
`QO57mDsdThGb59ewu1lK`

```gotemplate
{{ randAlphaNum 20 }}
```
___