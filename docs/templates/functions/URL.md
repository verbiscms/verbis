# URL




___

## query

Returns a query parameter by the given key.

### Accepts: 

`str interface{}` The query parameter key

### Returns:

`string` The query parameter value, or an empty string if the given key could not be cast
to a string, or if the value was not found.

### Examples:

**Obtain a query parameter**

Given the url of `https://verbiscms.com/test=123`, executing the function below
will return `123`. 

```gotemplate
{{ query "test" }}
```


