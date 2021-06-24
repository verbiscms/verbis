# Date & Time

These functions aid the formatting of date and time variables. Golang's date formatting is a little different 
to other languages in that it doesn't use any special codes (e.g `%Y` for Year). Instead it uses a standard 
time, which is:

`Mon Jan 2 15:04:05 MST 2006  (MST is GMT-0700)`

You can read more about GoLang time and date formatting [here](https://yourbasic.org/golang/format-parse-string-time-date-example/).

These functions can be useful for formatting post's `CreatedAt` and `UpdatedAt` fields to a user
friendly format.

___

## now

Returns the current time and date for use with other date functions.

### Returns:

`Time` The current date and time.

### Examples:

**Obtain the current date and time**

Results in something similar to:
`2009-11-10 23:00:00 +0000 UTC m=+0.000000001`

```gotemplate
{{ now }}
```

___

## date

Formats the given date with a string. 

### Accepts: 

`format string, date interface{}` The format and date.

### Returns:

`string` The formatted date time.

### Examples:

**Format to day, month and year with the current time**

Results in something similar to:
`01/12/2020`

```gotemplate
{{ date "02/01/2006" now }}
```

**Format to RFC822 with the current time**

Results in something similar to:
`01 Dec 20 15:00 BST`

```gotemplate
{{ date "02 Jan 06 15:04 MST" now }}
```
___

## dateInZone

Is the same function as `date` but accepts a timezone.

### Accepts: 

`format string, date interface{}, zone string` The format, date and zone.

### Returns:

`string` The formatted date time.

### Examples:

**Format to day, month and year with the current time (01/12/2020) and London (GMT) time**

Results in something similar to:
`01/12/2020`

```gotemplate
{{ dateInZone "02/01/2006" now "Europe/London" }}
```

___

## ago

Calculates the duration from the current `time.Now` in seconds resolution.

### Accepts:  

`time interface{}` The time to be calculated.

### Returns:

`string` The calculated duration.

### Examples:

**Calculate the time since the Post was updated**

Results in something similar to:
`4h22m1s`

```gotemplate
{{ ago .UpdatedAt }}
```

___

## duration

Formats a given number of seconds as a `time.Duration`.

### Accepts: 

`seconds interface{}` The duration to be formatted.

### Returns:

`string` The formatted duration.

### Examples:

**Format seconds**

Results in `1m25s`

```gotemplate
{{ duration 85 }}
```

___

## htmlDate

Formats a given date for use with HTML date pickers.

### Accepts: 

`date interface{}` The date to be formatted.

### Returns:

`string` The formatted date.

### Examples:

**Format the current date for a form date picker**

Results in something similar to `2020-05-01`

```gotemplate
<input type="date" value="{{ htmlDate now }}">
```

___

## htmlDateInZone

Is the same function as `htmlZone` but accepts a timezone.

### Accepts: 

`date interface{}, zone string` The date to be formatted and zone.

### Returns:

`string` The formatted date.

### Examples:

**Format the current date in a timezone for a form date picker**

Results in something similar to `2020-05-01`

```gotemplate
<input type="date" value="{{ htmlDateInZone now "Europe/London" }}">
```