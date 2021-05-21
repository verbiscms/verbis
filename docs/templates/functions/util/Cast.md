# Cast

These functions are an easy way to convert between different types within Verbis. They take in an `interface{}` and return
the desired type. To avoid interrupting the flow of templates, no errors are returned if the conversion failed, instead
the 0 or nil value for that type will be returned. 

All the functions below are prefixed with `to` for example, `toInt`.
___

## toBool

Casts an `interface` to a `bool` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`bool` The cast value or `false` if an error occurred.

### Examples:

**Cast a string to `bool`**

```gotemplate
{{ toBool "true" }}
```

___

## toString

Casts an `interface` to a `string` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`time.Time` The cast value or `""` (an empty string) if an error occurred.

### Examples:

**Cast a `int` to `string`**

```gotemplate
{{ toString "123" }}
```

___

## toSlice

Casts an `interface` to a `[]interface` type.
This is particularly useful for range loops, as Verbis automatically
resolves singular types. 

### Accepts:

`i interface{}` The value to be converted.

### Returns:

`[]interface{}` The cast value or `nil` if an error occurred.

### Examples:

**Cast a `int` to `slice`**

```gotemplate
{{ toSlice 1 }}
```

___

## toTime

Casts an `interface` to a `time.Time` type.

### Accepts:

`i interface{}` The value to be converted.

### Returns:

`time.Time` The cast value or `0001-01-01 00:00:00 +0000 UTC` if an error occurred.

### Examples:

**Cast a string to `time.Time`**

```gotemplate
{{ toTime "22 May 90 20:39:39 GMT" }}
```

___

## toDuration

Casts an `interface` to a `time.Duration` type.

### Accepts:

`i interface{}` The value to be converted.

### Returns:

`time.Duration` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `time.Duration`**

```gotemplate
{{ toDuration "123" }}
```

___

## toInt

Casts an `interface` to a `int` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`int` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `int`**

```gotemplate
{{ toInt "123" }}
```

___

## toInt8

Casts an `interface` to a `int8` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`int8` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `int8`**

```gotemplate
{{ toInt8 "123" }}
```

___

## toInt16

Casts an `interface` to a `int16` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`int16` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `int16`**

```gotemplate
{{ toInt16 "123" }}
```

___

## toInt32

Casts an `interface` to a `int32` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`int32` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `int32`**

```gotemplate
{{ toInt32 "123" }}
```

___

## toUInt

Casts an `interface` to a `uint` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`uint` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `uint`**

```gotemplate
{{ toUInt "123" }}
```

___

## toUInt8

Casts an `interface` to a `uint8` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`uint8` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `uint8`**

```gotemplate
{{ toUInt8 "123" }}
```

___

## toUInt16

Casts an `interface` to a `uint16` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`uint16` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `uint16`**

```gotemplate
{{ toUInt16 "123" }}
```

___

## toUInt32

Casts an `interface` to a `uint32` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`uint32` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `uint32`**

```gotemplate
{{ toUInt32 "123" }}
```

___

## toUInt64

Casts an `interface` to a `uint64` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`uint64` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `uint64`**

```gotemplate
{{ toUInt64 "123" }}
```

___

## toFloat32

Casts an `interface` to a `float32` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`float32` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `float32`**

```gotemplate
{{ toFloat32 "1.123" }}
```

___

## toFloat64

Casts an `interface` to a `float64` type.

### Accepts: 

`i interface{}` The value to be converted.

### Returns:

`float64` The cast value or `0` if an error occurred.

### Examples:

**Cast a `string` to `float64`**

```gotemplate
{{ toFloat64 "1.123" }}
```