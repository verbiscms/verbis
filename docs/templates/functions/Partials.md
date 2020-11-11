# Partials

Using the partial method is an easy way to include child components or layouts into templates by passing
the partial name (relative to the theme) and any data to be included in the rendered partial.

## isAuth

Check if the current user is logged in to the backend.

### Returns:

`bool` The current user state.

### Examples:

**See if the current user is logged in**

```
{{ if isAuth }}
    // Output for logged in users only.
{{ end }}
```

___
