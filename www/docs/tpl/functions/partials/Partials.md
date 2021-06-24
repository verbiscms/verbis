# Partials

Using the partial method is an easy way to include child components or layouts into templates by passing
the partial name (relative to the theme) and any data to be included in the rendered partial.

### Variable Scoping

If the argument list to a partial call is only one, you are able to access the data with the dot `{{ . }}` instead
of using the index function `{{ index . 0 }}`. Using the `dict` function along with partials are a powerful way to
include data in a child template.

Partials are unaware of global scope, meaning you need to pass the global context `$` if you wish to access
any of the global Verbis data such as `Post` or `Media`.

## partial

Render's a child component with given data.

### Returns:

`template.HTML, error` The rendered html or an error if the file was not found or could not be executed.

File extensions can be **anything** and not limited to the file extension set out in the theme options.

### Examples:

**Render a partial svg using a dict**

This example demonstrates how to include a circle svg with a dictionary of properties including 
`radius` and`fill`. As we are only passing one argument to the partial function (the dict) we
are able to access it using dot notation, for example: `.fill`.

```gotemplate
{{ partial "partials/circle.svg" (dict "radius" 50 "fill" "red") }}
```

The template which is stored (relative to the theme) in `partials/circle.svg`.

```gotemplate
<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
  <circle cx="50" cy="50" r="{{ .radius }}" fill="{{ .fill }}"/>
</svg>
```

**Render a partial with global data**

This example demonstrates how to pass the global Verbis data to a partial file.
As the dollar `$` is passed to the `user.cms` file, we are able to access the Post's authors first name
and website.

```gotemplate
{{ partial "partials/user.cms" $ }}
```

The template which is stored (relative to the theme) in `partials/user.cms`.

```gotemplate
<div class="user">
    <div class="user-name">{{ .Post.Author.FirstName }}</div>
    <div class="user-website">{{ .Post.Author.Website }}</div>
</div>
```

**Render a partial without a dictionary**

If for some reason you didn't' want to pass a dictionary, you can pass a slice of variables to the partial
to render. However, the `{{ index }}` function needs to be used along with the context (the dot - `.`) and
the index relative to the argument list.

```gotemplate
{{ partial "partials/card.cms" "Card title" "Card body" }}
```

The template which is stored (relative to the theme) in `partials/card.cms`.

```gotemplate
<div class="card">
    <div class="card-title">{{ index . 0 }}</div>
    <div class="card-body">{{ index . 1 }}</div>
</div>
```

___
