# Class

The auth functions used within templates are used to establish if the user is logged in to the backend of the
Verbis website. These functions rely on the `verbis-cookie` with a value of the current user's token. 
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


https://wordpress.stackexchange.com/questions/210097/language-attributes-for-two-languages