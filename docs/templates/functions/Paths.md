# Paths

These functions return useful path information of the Verbis project.
**Use with caution** exposing sensitive information to the frontend of the website
can be dangerous.

___

## basePath

Returns the base path (root) of the project.

### Returns:

`string` The path.

### Example:

```gotemplate
{{ basePath }}
```

___

## adminPath

Returns the admin path of the project, where the SPA (Vue) is stored.

### Returns:

`string` The path.

### Example:

```gotemplate
{{ adminPath }}
```

___

## apiPath

Returns the API path of the project, where backend logic is stored.

### Returns:

`string` The path.

### Example:

```gotemplate
{{ apiPath }}
```

___

## themePath

Returns the currently active theme path.

### Returns:

`string` The path.

### Example:

```gotemplate
{{ themePath }}
```

___

## uploadsPath

Returns the uploads path of the project.

### Returns:

`string` The path.

### Example:

```gotemplate
{{ uploadsPath }}
```

___

## assetsPath

Returns the assets URL for the theme, for example "/assets".

### Returns:

`string` The path.

### Example:

```gotemplate
{{ uploadsPath }}
```

___

## storagePath

Returns the storage path of the project, where fields, uploads and logs are stored.

### Returns:

`string` The path.

### Example:

```gotemplate
{{ uploadsPath }}
```

___

## templatesPath

Returns the directory where page templates are stored.

### Returns:

`string` The path.

### Example:

```gotemplate
{{ templatesPath }}
```

___

## layoutsPath

Returns the directory where page layouts are stored.

### Returns:

`string` The path.

### Example:

```gotemplate
{{ layoutsPath }}
```
