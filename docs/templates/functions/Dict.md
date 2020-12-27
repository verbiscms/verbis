# Dict

As there is no built in support from GoLang's template system for `map[string]interface{}`, the dict
helper function allows you to create a dictionary for a list of key value pairs. 
This function is particularly useful for passing data to the `getPosts` function as query arguments,
or passing multiple arguments to a partial file.

___

## dict

Creates a dictionary of key value pairs, the key must always be a string.

### Accepts: 

`...interface{}` A list of key, value pairs.

### Returns:

`map[string]interface{}, error` A new map or an error if a key was not of type `string` or there
was an odd number of values.

### Examples:

**Create a dictionary**

Create a new map with key value pairs and assign to a variable. To access values inside
the map, we simply use the dot `(.)` followed by the key.

```gotemplate
{{ $map := dict "colour" "green" "height" 20 }}
{{ $map.green }}
```

**Pass multiple values to a partial**

Create a new map with key value pairs and assign to a variable

**Partial call**

Here we use the partial function with a `dict` to pass properties to render an SVG.

```gotemplate
{{ partial (dict "colour" "green" "width" 50 "height" 50) }}
```

**Partial definition**

We can access the values of the `dict` by calling them by value to render the SVG with
a specific colour, width and height.

```gotemplate
<svg width="100" height="100">
  <circle cx="{{ .width }}" cy="{{ .height }}" r="40" stroke-width="4" fill="{{ .colour }}" />
</svg>
```
___