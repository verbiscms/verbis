# Reflect

GoLang has primitive types such as `bool`, `int` and string as it is a statically typed language, it also
has an open type system meaning custom structs can be created such as `Posts` and the `Media` types.

Sometimes its useful to know the underlying value of a particular type within Verbis. These functions
assist with determining and comparing types. 

___

## kindIs

Returns true if the target and source types match.

### Accepts: 

`target string, src interface{}` The target string and interface to compare.

### Returns:

`bool` If the types match.

### Examples:

**Determine if a type is an integer**

Returns `true`

```gotemplate
{{ kindIs "int" 123 }}
```

___

## kindOf

Returns the type of the given `interface` as a string.

### Accepts: 

`src interface{}` The source to obtain the value of.

### Returns:

`string` The type.

### Examples:

**Obtain the kind of an integer**

Returns `int`

```gotemplate
{{ kindOf 123 }}
```

___

## typeOf

Returns the underlying type of the given value.

### Accepts: 

`src interface{}` The source to obtain the type of.

### Returns:

`string` The type.

### Examples:

**Obtain the type of a `.Post`**

This example evaluates the type of the global `.Post` variable.
Returns `domain.Post`

```gotemplate
{{ typeOf .Post }}
```

___

## typeIs

Similar to `kindIs` but its used for types, instead of primitives.

### Accepts: 

`target string, src interface{}` The target string and interface to compare.

### Returns:

`bool` If the types match.

### Examples:

**Compare the `.Post` variable**

This example compares the type of the global `.Post` variable.
Returns `true`

```gotemplate
{{ typeOf "domain.Post" .Post }}
```

___

## typeIsLike

Similar to `typeIs` but dereferences pointers.

### Accepts: 

`target string, src interface{}` The target string and interface to compare.

### Returns:

`bool` If the types match.

### Examples:

**Compare the `.Post` variable**

This example compares the type of the global `.Post` variable.
Returns `true`

```gotemplate
{{ typeOf "domain.Post" .Post }}
```

___