# Auth

The auth functions used within templates are used to establish if the user is logged in to the backend of the
Verbis website. These functions relies on the `verbis-cookie` with a value of the current user's token. 
If the user clears their cookies, usage may vary until the user logs out again.

These functions can be handy to display hidden content or media to logged in users.

___

## isAuth

Checks if the current user is logged in to the backend.

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

## isAdmin

Checks if the current user is logged in to the backend and if the user
has administrator or owner privileges.

### Returns:

`bool` The current user state.

### Examples:

**See if the current user is an adminstrator or owner**

```
{{ if isAdmin }}
    // Output for logged admin users only.
{{ end }}
```