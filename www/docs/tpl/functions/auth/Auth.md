# Auth

The auth functions used within templates are used to establish if the user is logged in to the backend of the
Verbis website and check role's of the user. These functions rely on the `verbis-cookie` with a value of the current user's token. 
If the user clears their cookies, usage may vary until the user logs out and in again.

These functions can be handy to display hidden content or media to logged in users.

___

## auth

Check if the current user is logged in to the backend.

### Returns:

`bool` The current user state.

### Examples:

**See if the current user is logged in**

```gotemplate
{{ if auth }}
    // Output for logged in users only.
{{ end }}
```

___

## admin

Check if the current user is logged in to the backend and if the user
has administrator or owner privileges.

### Returns:

`bool` The current user state.

### Examples:

**See if the current user is an adminstrator or owner**

```gotemplate
{{ if admin }}
    // Output for logged admin users only.
{{ end }}
```