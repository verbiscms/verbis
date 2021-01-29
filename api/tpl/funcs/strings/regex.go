package strings

import "regexp"

// Match
//
// Returns true if the input string contains and
// matches of the regular expression pattern.
//
// Example: {{ regexMatch "^Verbis" "Verbis CMS" }}
// Returns `true`
func (ns *Namespace) Match(regex string, str string) bool {
	match, _ := regexp.MatchString(regex, str)
	return match
}

// FindAll
//
// Returns a slice of all matches of the regular
// expressions with the given input string.
//
// Example: {{ regexFindAll "[1,3,5,7]" "123456789" -1 }}
// Returns: `[1 3 5 7]`
func (ns *Namespace) FindAll(regex string, str string, i int) []string {
	r := regexp.MustCompile(regex)
	return r.FindAllString(str, i)
}

// Find
//
// Return the first (left most) match of the
// regular expression in the input string
//
// Example: {{ regexFind "verbis.?" "verbiscms" }}
// Returns `verbisc`
func (ns *Namespace) Find(regex string, str string) string {
	r := regexp.MustCompile(regex)
	return r.FindString(str)
}

// ReplaceAll
//
// Returns a copy of the input string, replacing matches of the Regexp with the replacement string.
// Within the string replacement, $ signs are interpreted as in Expand, so for instance $1
// represents the first submatch.
//
// Example: {{ regexReplaceAll "a(x*)b" "-ab-axxb-" "\${1}W" }}
// Returns: `-W-xxW-`
func (ns *Namespace) ReplaceAll(regex string, str string, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(str, repl)
}

// ReplaceAllLiteral
//
// Returns a copy of the input string, replacing matches of the Regexp with the replacement string
// replacement. The replacement string is substituted directly, without using Expand.
//
// Example: {{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}" }}
// Returns: `-${1}-${1}-`
func (ns *Namespace) ReplaceAllLiteral(regex string, str string, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllLiteralString(str, repl)
}

// Split
//
// Slices the input string into substrings separated by the expression and returns a slice of the
// substrings between expression matches. The last parameter `i` determines the number of
// substrings to return, where `-1` returns all matches.
//
// Example: {{ regexSplit "b+" "verbis" -1 }}
// Returns: `[ver  s]`
func (ns *Namespace) Split(regex string, str string, i int) []string {
	r := regexp.MustCompile(regex)
	return r.Split(str, i)
}

// QuoteMeta
//
// QuoteMeta returns a string that escapes all regular expression metacharacters
// inside the argument text; the returned string is a regular expression matching
// the literal text.
//
// Example: {{ regexQuoteMeta "verbis+?" }}
// Returns: `verbis`
func (ns *Namespace) QuoteMeta(str string) string {
	return regexp.QuoteMeta(str)
}
