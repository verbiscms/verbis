# Attributes

The attributes functions are helpers for theme development and layout creation. As pages in Verbis are
dynamic, sometimes it's challenging targeting specific pages or resources with CSS or Javascript, which 
is where these functions come in.

___

## body

The body function returns useful classes specific to the page or resource being rendered. They can be 
used for targeting particular pages with CSS or Javascript. 

***Note:*** the function does not include `body`, just a string of classes, so you are able to add custom classes if you wish.
All special characters are removed and forward slashes and spaces are replaced with a dash `-`.

The following is outputted when the `{{ body }}` function is called.

| Description                  | Example                              |
| ---------------------------- | ------------------------------------ |
| Resource type                | `page` or `posts`                    |
| Page ID                      | `page-id-4`                          |
| Title                        | `Page Title` becomes `page-title`    |
| Page Template                | `page-template-news-archive`         |
| Page Layout                  | `page-layout-main`                   |
| Logged In                    | `logged-in` if user is authenticated |

### Returns:

`string` The body classes.

### Examples:

**Output body classes for a layout**

The resulting call would result in something similar to below.

`page page-id-4 page-title page-template-news-archive page-layout-main logged-in`

```gotemplate
<body class="{{ body }} custom-class">
<!-- Content -->
</body>
```
___

## lang

The lang function returns the current language set in the admin interface under `General`.
This can be useful for setting the `lang` attribute on the HTML tag as shown below.

### Returns:

`string` The language set.

### Examples:

**Output language for the html tag**

```gotemplate
<html lang="{{ lang }}">
<!-- Content -->
</html>
```

___

## homepage

The homepage function returns true if the current post is the homepage.

### Returns:

`bool` If the post is the homepage.

### Examples:

**Output content for the homepage only**

```gotemplate
{{ if homepage }} 
    <h2>Hello</h2> 
{{ end }}
```