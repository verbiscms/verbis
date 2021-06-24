# Regex

Regular expressions (shortened as regex or regexp) also referred to as rational expression) is a sequence of characters that 
define a search pattern. Usually such patterns are used by string-searching algorithms for "find" or "find and replace" operations on strings, or for input validation. 
Verbis has some useful functions to find and replace on strings shown below.

___

## regexMatch

Returns true if the input string contains and matches of the regular expression pattern.

### Accepts: 

`regex string, str string` The regular expression pattern and string to check.

### Returns:

`bool` If any matches have been found.

### Examples:

**Determine if a string starts with Verbis**

Returns `true`

```gotemplate
{{ regexMatch "^Verbis" "Verbis CMS" }}
```

___

## regexFindAll

Returns a slice of all matches of the regular expressions with the given input string.
The last parameter `i` determines the number of substrings to return, where `-1` returns all matches.

### Accepts: 

`regex string, str string, amount int` The regular expression pattern, string to check and 
 amount of matches to return.

### Returns:

`[]string` Slice of string matches.

### Examples:

**Replace all integers with expression**

Returns `[1 3 5 7]`

```gotemplate
{{ regexFindAll "[1,3,5,7]" "123456789" -1 }}
```

___

## regexFind

Return the first (left most) match of the regular expression in the input string.

### Accepts: 

`regex string, str string` The regular expression pattern and string to check.

### Returns:

`string` The first regular expression match.

### Examples:

**Find the first part of a string**

Returns `verbisc`

```gotemplate
{{ regexFind "verbis.?" "verbiscms" }}
```

___

## regexReplaceAll
  
Returns a copy of the input string, replacing matches of the Regexp with the replacement string. 
Within the string replacement, $ signs are interpreted as in Expand, so for instance $1 
represents the first submatch.

### Accepts: 

`regex string, str string` The regular expression pattern and string to replace.

### Returns:

`string` The replaced string.

### Examples:

**Replace a whole string with pattern matches**

Returns `-W-xxW-`

```gotemplate
{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W" }}
```

___

## regexReplaceAllLiteral
  
Returns a copy of the input string, replacing matches of the Regexp with the replacement string replacement.
The replacement string is substituted directly, without using Expand.

### Accepts: 

`regex string, str string` The regular expression pattern and string to replace.

### Returns:

`string` The replaced string.

### Examples:

**Replace a whole string with pattern matches**

Returns `-${1}-${1}-`

```gotemplate
{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}" }}
```

___

## regexSplit
  
Slices the input string into substrings separated by the expression and returns a slice of the
substrings between expression matches. The last parameter `i` determines the number of 
substrings to return, where `-1` returns all matches.

### Accepts: 

`regex string, str string, amount int` The regular expression pattern, string to replace and amount
of matches to return.

### Returns:

`string` The replaced string.

### Examples:

**Split part of a string using regular expression**

Returns `[ver  s]`

```gotemplate
{{ regexSplit "b+" "verbis" -1 }}
```

___

## regexQuoteMeta
  
QuoteMeta returns a string that escapes all regular expression metacharacters
inside the argument text; the returned string is a regular expression matching
the literal text.

### Accepts: 

`str string` The string to be replaced 

### Returns:

`string` The replaced string.

### Examples:

**Split part of a string using regular expression**

Returns `verbis`

```gotemplate
{{ regexQuoteMeta "verbis+?" }}
```