package templates

import (
	"github.com/spf13/cast"
	"strings"
)

// len
//
// Returns length of given interface{} in `int`.
func (t *TemplateFunctions) len(a interface{}) int {
	return len(cast.ToString(a))
}

// replace
//
// Returns new replaced string with all matches.
func (t *TemplateFunctions) replace(old, new, src string) string {
	return strings.Replace(old, new, src, -1)
}

// substr
//
// Returns new substring of the given string.
func (t *TemplateFunctions) substr(str string, start, end interface{}) string {
	st := cast.ToInt(start)
	en := cast.ToInt(end)
	if st < 0 {
		return str[:en]
	}
	if en < 0 || en > len(str) {
		return str[st:]
	}
	return str[st:en]
}

// trunc
//
// Returns a truncated string with no suffix, negatives apply.
func (t *TemplateFunctions) trunc(s string, a interface{}) string {
	i := cast.ToInt(a)
	if i < 0 && len(s)+i > 0 {
		return s[len(s)+i:]
	}
	if i >= 0 && len(s) > i {
		return s[:i]
	}
	return s
}

// ellipsis
//
// Returns a ellipsis (...) string from the given length
func (t *TemplateFunctions) ellipsis(len int, s string) string {
	marker := "..."
	if len < 4 {
		return s
	}
	return t.substr(s, 0, len) + marker
}
