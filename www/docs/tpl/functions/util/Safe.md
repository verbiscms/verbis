# Safe

By default, GoLang statements `{{ }}` are automatically stripped to prevent
XSS attacks. If you do not want your data to be escaped, you may use the 
following `safe` functions. 

If unsafe content reached a CSS value or URL the special value of `ZgotmplZ` will
be outtputted, indicating you need to unescape the content.

___

## safeHTML

Unescapes HTML content with the given input. This function is especially useful
for escaping the Verbis editor content.

### Accepts: 

`str interface{}` The string to be unescaped.

### Returns:

`template.HTML, error` The unescaped HTML or an error if the input could not 
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
HTMLAttr encapsulates an HTML attribute from a trusted source, for example, ` dir="ltr"`.

### Accepts: 

`str interface{}` The string to be unescaped.

### Returns:

`template.HTMLAttr, error` The unescaped HTML attribute, or an error if the input could not 
be cast to a string.

___

## safeCSS

Declare the given string as a `safe` CSS string.

CSS encapsulates known safe content that matches any of:
1. The CSS3 stylesheet production, such as `p { color: purple }`.
2. The CSS3 rule production, such as `a[href=~"https:"].foo#bar`.
3. CSS3 declaration productions, such as `color: red; margin: 2px`.
4. The CSS3 value production, such as `rgba(0, 0, 255, 127)`.

### Accepts: 

`str interface{}` The string to be unescaped.

### Returns:

`template.CSS, error` The unescaped CSS, or an error if the input could not 
be cast to a string.

### Examples:

**Style a paragraph from a field**

This function will render safe CSS from a field, to a paragraph element.

`<p style="color: purple;">Hello Verbis</p>`
 
If safeCSS was not used, it would result in escaped text:
 
`<p style="ZgotmplZ">Hello Verbis</p>`

```gotemplate
<p style="{{ "style" | getField | safeCSS }}">Hello Verbis</p>
```

___

## safeJS

Declare the given string as a `safe` JS string.
safeJS encapsulates a known safe EcmaScript5 Expression, for example `(x + y * z())`.

### Accepts: 

`str interface{}` The string to be unescaped.

### Returns:

`template.JS, error` The unescaped JS, or an error if the input could not 
be cast to a string.

___

## safeJSStr

Declare the given string as a `safe` JS string.
JSStr encapsulates a sequence of characters meant to be embedded between quotes in a JavaScript expression.
The string must match a series of StringCharacters:
    StringCharacter :: SourceCharacter but not `\` or LineTerminator | EscapeSequence
Note that LineContinuations are not allowed.
JSStr("foo\\nbar") is fine, but JSStr("foo\\\nbar") is not.

### Accepts: 

`str interface{}` The string to be unescaped.

### Returns:

`template.JSStr, error` The unescaped JS string, or an error if the input could not 
be cast to a string.

___

## safeUrl

Declares the provided string as a `safe URL` or URL substring.

### Accepts: 

`str interface{}` The string to be unescaped.

### Returns:

`template.URL, error` The unescaped URL, or an error if the input could not 
be cast to a string.

### Examples:

**Outputting a safe URL from a field**

This function will render a safe URL from a field, to a paragraph element.

`<p style="color: purple;">Hello Verbis</p>`
 
If safeURL was not used, it would result in escaped text:
 
`<a href="ZgotmplZ">Hello Verbis</a>`

```gotemplate
<a href="{{ "url" | getField | safeURL }}">Click me!</a>
```
