package strings

import (
	"github.com/spf13/cast"
	"strings"
)

// replace
//
// Returns new replaced string with all matches.
//
// Returns `hello-verbis-cms`
// Example: {{ replace "" "-" "hello verbis cms" }}
func (ns *Namespace) replace(old, new, src string) string {
	return strings.Replace(old, new, src, -1)
}

// substr
//
// Returns new substring of the given string.
//
// Returns `hello`
// Example: {{ substr "hello verbis" 0 5 }}
func (ns *Namespace) substr(str string, start, end interface{}) string {
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
//
// Returns `verbis`
// Example: {{ trunc "hello verbis" -5 }}
func (ns *Namespace) trunc(str string, a interface{}) string {
	i := cast.ToInt(a)
	if i < 0 && len(str)+i > 0 {
		return str[len(str)+i:]
	}
	if i >= 0 && len(str) > i {
		return str[:i]
	}
	return str
}

// ellipsis
//
// Returns a ellipsis (...) string from the given length.
//
// Returns `hello verbis...`
// Example: {{ ellipsis "hello verbis cms!" 11 }}
func (ns *Namespace) ellipsis(str string, len interface{}) string {
	i := cast.ToInt(len)
	marker := "..."
	if i < 4 {
		return str
	}
	return ns.substr(str, 0, i) + marker
}
