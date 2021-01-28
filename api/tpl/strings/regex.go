package strings

import "regexp"

// regexMatch
//
// Returns true if the input string contains and
// matches of the regular expression pattern.
//
// Returns true
// Example: {{ regexMatch "^Verbis" "Verbis CMS" }}
func (ns *Namespace) regexMatch(regex string, str string) bool {
	match, _ := regexp.MatchString(regex, str)
	return match
}

// regexReplaceAll
//
// Returns a slice of all matches of the regular
// expressions with the given input string.
//
// Returns [1 3 5 7]
// Example: {{ regexFindAll "[1,3,5,7]" "123456789" -1 }}
func (ns *Namespace) regexFindAll(regex string, str string, i int) []string {
	r := regexp.MustCompile(regex)
	return r.FindAllString(str, i)
}

// regexFind
//
// Return the first (left most) match of the
// regular expression in the input string
//
// Returns verbisc
// Example: {{ regexFind "verbis.?" "verbiscms" }}
func (ns *Namespace) regexFind(regex string, str string) string {
	r := regexp.MustCompile(regex)
	return r.FindString(str)
}

// regexReplaceAll
//
// Returns a copy of the input string, replacing matches of the Regexp with the replacement string.
// Within the string replacement, $ signs are interpreted as in Expand, so for instance $1
// represents the first submatch.
//
// Returns -W-xxW-
// Example: {{ regexReplaceAll "a(x*)b" "-ab-axxb-" "\${1}W" }}
func (ns *Namespace) regexReplaceAll(regex string, str string, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(str, repl)
}

// regexReplaceAllLiteral
//
// Returns a copy of the input string, replacing matches of the Regexp with the replacement string
// replacement. The replacement string is substituted directly, without using Expand.
//
// Returns `-${1}-${1}-`
// Example: {{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}" }}
func (ns *Namespace) regexReplaceAllLiteral(regex string, str string, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllLiteralString(str, repl)
}

// regexSplit
//
// Slices the input string into substrings separated by the expression and returns a slice of the
// substrings between expression matches. The last parameter `i` determines the number of
// substrings to return, where `-1` returns all matches.
//
// Returns `[ver  s]`
// Example: {{ regexSplit "b+" "verbis" -1 }}
func (ns *Namespace) regexSplit(regex string, str string, i int) []string {
	r := regexp.MustCompile(regex)
	return r.Split(str, i)
}

// regexQuoteMeta
//
// QuoteMeta returns a string that escapes all regular expression metacharacters
// inside the argument text; the returned string is a regular expression matching
// the literal text.
//
// Returns `verbis`
// Example: {{ regexQuoteMeta "verbis+?" }}
func (ns *Namespace) regexQuoteMeta(str string) string {
	return regexp.QuoteMeta(str)
}
