package strings

import (
	"github.com/spf13/cast"
	"strings"
)

// Replace
//
// Returns new replaced string with all matches.
//
// Example: {{ replace "" "-" "hello verbis cms" }}
// Returns: `hello-verbis-cms`
func (ns *Namespace) Replace(old, new, src string) string {
	return strings.Replace(old, new, src, -1)
}

// Substr
//
// Returns new substring of the given string.
//
// Example: {{ substr "hello verbis" 0 5 }}
// Returns: `hello`
func (ns *Namespace) Substr(str string, start, end interface{}) string {
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

// Trunc
//
// Returns a truncated string with no suffix, negatives apply.
//
// Example: {{ trunc "hello verbis" -5 }}
// Returns: `verbis`
func (ns *Namespace) Trunc(str string, a interface{}) string {
	i := cast.ToInt(a)
	if i < 0 && len(str)+i > 0 {
		return str[len(str)+i:]
	}
	if i >= 0 && len(str) > i {
		return str[:i]
	}
	return str
}

// Ellipsis
//
// Returns a ellipsis (...) string from the given length.
//
// Example: {{ ellipsis "hello verbis cms!" 11 }
// Returns: `hello verbis...`}
func (ns *Namespace) Ellipsis(str string, len interface{}) string {
	i := cast.ToInt(len)
	marker := "..."
	if i < 4 {
		return str
	}
	return ns.Substr(str, 0, i) + marker
}
