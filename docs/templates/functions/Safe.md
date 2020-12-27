# Safe

By default, GoLang statements `{{ }}` are automatically stripped to prevent
XSS attacks. If you do not want your data to be escaped, you may use the 
following `safe` functions. 

___

## safeHTML

Unescapes HTML content with the given input. This function is especially useful
for escaping the Verbis editor content.

### Accepts: 

`interface{}` The string to be escaped.

### Returns:

`template.HTML, error` The escaped HTML or an error if the input could not 
be cast to a string.

### Examples:

**Escape a rich text**

This function will escape rich text from a field, with the following content using pipelines.

`© 2020 VerbisCMS.  <a href='https://verbiscms.com'>All rights reserved</a>.`
 
If safeHTML was not used, it would result in escaped text:
 
`&copy; 2020 VerbisCMS.  &lt;a href='https://verbiscms.com'&gt;All rights reserved&lt;/a&gt;.`

```gotemplate
{{ "richtext" | getField | safeHTML }}
```

___

## safeHTMLAttr

Unescapes HTML attribute content with the given input.

### Accepts: 

`interface{}` The string to be escaped.

### Returns:

`template.HTML, error` The escaped HTML attribute, or an error if the input could not 
be cast to a string.

### Examples:

**Escape an attribute**

This function will escape rich text from a field, with the following content using pipelines.

`© 2020 VerbisCMS.  <a href='https://verbiscms.com'>All rights reserved</a>.`
 
If safeHTML was not used, it would result in escaped text:
 
`&copy; 2020 VerbisCMS.  &lt;a href='https://verbiscms.com'&gt;All rights reserved&lt;/a&gt;.`

```gotemplate
{{ "richtext" | getField | safeHTML }}