# OS

These functions return information about the environment and operating system.
**They should be used with caution** Sensitive information can be exposed to the frontend if
not used properly such as database credentials and other secret keys.

___

## env

Reads an environment variable from the system or `.env` file of the project
by a key.

### Accepts: 

`key string` The environment variable key.

### Returns:

`string` The value.

### Examples:

**Detecting if Verbis is in debug mode**

This could be used for displaying different debug data or content if Verbis is
in debug mode.

```gotemplate
{{ env "APP_DEBUG" }}
```

___

## expandEnv

Substitutes environment variables in a string using the `$` dollar symbol.

### Accepts: 

`str string` The string to be expanded.

### Returns:

`string` The value.

### Examples:

**Output the application name in a header**

```gotemplate
<header>
    <p>{{ expandEnv "Welcome to $APP_NAME" }}</p>
</header>
```

___