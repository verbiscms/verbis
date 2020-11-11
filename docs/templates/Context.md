# Context, Scope & Dot

It may be somewhat difficult moving from other languages such as Javascript, Python, PHP or others
to the Golang Template eco system and constraints that come with it.

The context & scope is key to understanding how to output data on the frontend of your Verbis templates.
Below we go into more detail about scope, context amongst other condurums faced when developing a GoLang
template.

## Context

The dot `{{ . }}` always refers to the current context. The context could be any number of things, such as
being inside a `{{ range }}` loop or rendering a partial.

At the very top level, the template that's rendered within Verbis, the dot `{{ . }}` refers to all the 
data made available to it such as `Site` and `Post` (see `variables` for more info).
Inside a loop however, the context changes.

## Global

Wait a minute ✋, how can I access my global `Site` and `Post` data if I'm inside a loop?

Luckily GoLang templates offer the `{{ $ }}` dollar sign 💵 to help you with this. The dollar
contains data set to the starting value of the template. Which means you can still access 
valuable data even if you are inside a loop.

For example, below shows a range loop over `$navItems` and we output of the particular nav item
using the dot, but we have lost context in terms of the global scope, as the dot `{{ . }}` changes
value. To access the global context we can use `{{ $.Site }}`.

```gotemplate
<ul>
{{ range $navItems }}
  <li>
    <a href="">{{ . }}</a>
    {{ $.Site.Title }}
  </li>
{{ end }}
</ul>
```

## Clever tricks 💡

### Defining variables

We can assign the global context to a variable at the very top of a template, to preserve the scope
of the global template data. This preservers the dollar `{{ $ }}` and avoids anyone redefining our 
magical global context.

```gotemplate
{{ $global := $ }}
<h1>{{ $global.Site.Title }}</h1>
```

### With

With, `with` we can redefine what the context and the dot `{{ . }}` means. This makes your templating
 more readable by shortening your variable names by using `shifting context` as opposed to
`unshifting context`. Take a look at below, you will see the difference!

**Using `with`**

```gotemplate
{{ with .Post }}
 	{{/* Here the dot is now .Post */}} 
	<h4>{{ .Title }}</h4>
    <a href="{{ .Slug }}">{{ .Title }}</a>
{{ end }}
```

**Using global scope**

```gotemplate
{{ with .Post.Title }}
	<h1>{{ .Post.Title }}</h1>
    <a href="{{ .Post.Slug }}">{{ .Post.Title }}</a>
{{ end }}
```







