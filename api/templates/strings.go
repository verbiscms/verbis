package templates

import (
	"github.com/spf13/cast"
	"strings"
)

// len
//
// Returns length of given interface{} in `int`.
//

func (t *TemplateManager) len(a interface{}) int {
	return len(cast.ToString(a))
}

// replace
//
// Returns new replaced string with all matches.
//
// Example: {{ replace "" "-" "hello verbis cms" }} Returns `hello-verbis-cms`
func (t *TemplateManager) replace(old, new, src string) string {
	return strings.Replace(old, new, src, -1)
}

// substr
//
// Returns new substring of the given string.
//
// Example: {{ substr "hello verbis" 0 5 }} Returns `hello`
func (t *TemplateManager) substr(str string, start, end interface{}) string {
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
// Example: {{ trunc "hello verbis" -5 }} Returns `verbis`
func (t *TemplateManager) trunc(str string, a interface{}) string {
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
// Example: {{ ellipsis "hello verbis cms!" 11 }} Returns `hello verbis...`
func (t *TemplateManager) ellipsis(str string, len int) string {
	marker := "..."
	if len < 4 {
		return str
	}
	return t.substr(str, 0, len) + marker
}
