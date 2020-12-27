# Math

The math functions can handle values within the range of integer and float types allowing easy integration into templates.
The functions below mostly operate on a `int64` values which are returned from the function.

___

## add

Add two or more numbers together.

### Accepts: 

`numbers ...interface{}` Two or more values to be added.

### Returns:

`int64` The sum of the input.

### Examples:

**Adding 2 numbers together**

Returns `4`

```gotemplate
{{ add 2 2 }}
```

**Adding 3 numbers together**

Returns `10`

```gotemplate
{{ add 2 4 4 }}
```
___

## subtract

Subtract one value from another.

### Accepts: 

`a, b interface{}` The values to be subtracted 

### Returns:

`int64` The subtracted value.

### Examples:

**Subtracting two numbers**

Returns `90`

```gotemplate
{{ subtract 100 10 }}
```

___

## divide

Divide one value into another.

### Accepts: 

`a, b interface{}` The values to be divided 

### Returns:

`int64` The divided value.

### Examples:

**Dividing two numbers**

Returns `4`

```gotemplate
{{ divide 16 4 }}
```

___

## multiply

Multiply two or more numbers together.

### Accepts: 

`...interface{}` Two or more values to be multiplied.

### Returns:

`int64` The multiplied value.

### Examples:

**Multiplying 2 numbers together**

Returns `16`

```gotemplate
{{ add 4 4 }}
```

**Multiplying 3 numbers together**

Returns `8`

```gotemplate
{{ add 2 2 2 }}
```

___

## mod

Calculates the remainder after dividing one number by another.

### Accepts: 

`a, b interface{}` The values to be reduced to a range. 

### Returns:

`int64` The remainder value.

### Examples:

**Obtain the remainder from two numbers**

Returns `1`

```gotemplate
{{ mod 10 9 }}
```

___

## round

Round to the nearest integer, rounding halfway from zero.

### Accepts: 

`number interface{}` The number to be rounded.

### Returns:

`int64` The rounded value.

### Examples:

**Rounding a float value (down)**

Returns `10`

```gotemplate
{{ round 10.2 }}
```

**Rounding a float value (up) Returns `11`**

```gotemplate
{{ round 10.6 }}
```

___

## ceil

Ceil rounds a number up to the next largest integer.

### Accepts: 

`number interface{}` The number to be rounded.

### Returns:

`float64` The rounded value.

### Examples:

**Ceil a float value**

Returns `10`

```gotemplate
{{ ceil 9.32 }}
```
___

## floor

Floor rounds a number down to the next largest integer.

### Accepts: 

`number interface{}` The number to be rounded.

### Returns:

`float64` The rounded value.

### Examples:

**Ceil a float value**

Returns `9`

```gotemplate
{{ floor 9.62 }}
```

___

## min

Min returns the lowest integer in a slice of numbers.

### Accepts: 

`numbers interface{}` The slice of integers.

### Returns:

`int64` The lowest value.

### Examples:

**Obtain the smallest number from a slice**

Returns `1`

```gotemplate
{{ min 20 1 100 }}
```

___

## max

Max returns the highest integer in a slice of numbers.

### Accepts: 

`numbers interface{}` The slice of integers.

### Returns:

`int64` The highest value.

### Examples:

**Obtain the largest number from a slice**

Returns `100`

```gotemplate
{{ max 20 1 100 }}
```
