# Verbis Header/Footer & Meta

Using the `{{ verbisHead }}` and `{{ verbisFoot }}` functions are crucial helpers to output rich
SEO data to increase the theme's visibility on common search engines.  

## verbisHead

Obtain crucial meta information & Verbis specific data. This function should be
outputted in the `<head></head>`.

### Returns:

`template.HTML` The following is returned to the rendered `verbisHead`.

- Global Code injection header.
- Post code injection header.
- `norobots` if the site is not public or the post is set to private.
- Meta description (Post description overrides global).
- Twitter & Facebook cards (Post social cards overrides global).
- Canonical link.

### Examples:

**Output the verbis header**

```gotemplate
<head>
    {{ verbisHead }}
</head>
```

## verbisFoot

Obtain code injection both globally and post specific. This function shoudl be
outputted just before the closing `<body>`.

### Returns:

`template.HTML` The following is returned to the rendered `verbisFoot`.

- Global code injection footer.
- Post code injection footer.


### Examples:

**Output the verbis footer**

```gotemplate
{{ verbisFoot }}
</body>
```

## metaTitle

Retrieve the meta title for use in the `<title></title>` element.

### Returns:

`string` The post meta title or the global meta title if the post specific meta 
could not be found.

- Global code injection footer.
- Post code injection footer.

**Output the verbis footer**

```gotemplate
<title>Verbis - {{ metaTitle }}</title>
```