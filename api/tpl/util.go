package tpl

import (
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

// len
//
// Returns the length of a variable according to its type.
// If the length of the type passed could not be
// retrieved, it will return `0`.
//
// Example: {{ len "hello" }}  Returns `5`
func (t *TemplateManager) len(i interface{}) int64 {
	typ := reflect.TypeOf(i).Kind()
	switch typ {
	default:
		return 0
	case reflect.Slice, reflect.Array, reflect.String, reflect.Map:
		return int64(reflect.ValueOf(i).Len())
	case reflect.Ptr:
		test := reflect.Indirect(reflect.ValueOf(i))
		return int64(test.Len())
	}
}

// explode
//
// Breaks a string into array with a delimiter (separator).
//
// Example: {{ explode "," "hello there !" }} Returns `[hello there !]`
func (t *TemplateManager) explode(delim interface{}, text interface{}) []string {
	d, err := cast.ToStringE(delim)
	if err != nil {
		return nil
	}

	tt, err := cast.ToStringE(text)
	if err != nil {
		return nil
	}

	if len(d) > len(tt) {
		return strings.Split(d, tt)
	} else {
		return strings.Split(tt, d)
	}
}

// implode
//
// Returns a string from the elements of an array using a
// glue string to join them together.
//
// Example: {{ slice 1 2 3 | explode "," }} Returns `[1 2 3]`
func (t *TemplateManager) implode(glue interface{}, slice interface{}) string {
	str, err := cast.ToStringE(glue)
	if err != nil {
		return ""
	}

	typ := reflect.TypeOf(slice).Kind()

	switch typ {
	case reflect.Slice, reflect.Array:
		val := reflect.ValueOf(slice)

		ret := make([]string, val.Len())
		for i := 0; i < val.Len(); i++ {
			s, err := cast.ToStringE(val.Index(i).Interface())
			if err == nil {
				ret[i] = s
			}
		}

		return strings.Join(ret, str)
	default:
		return ""
	}
}
