package slice

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"reflect"
)

// Slice
//
// Creates a slice (array) of passed arguments.
//
// Example: {{ slice "hello" "world" "!" }}
// Returns: `[hello world !]`
func (ns *Namespace) Slice(i ...interface{}) []interface{} {
	return i
}

// Append
//
// Adds and element to the end of the slice.
//
// Example: {{ append (slice "hello" "world" "!") "verbis" }}
// Returns: `[hello world ! verbis]`
func (ns *Namespace) Append(slice interface{}, i interface{}) ([]interface{}, error) {
	const op = "Templates.append"

	typ := reflect.TypeOf(slice).Kind()

	switch typ {
	default:
		return nil, &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Cannot append on type: %s", typ), Operation: op, Err: fmt.Errorf("unable to append to slice with type: %s", typ)}
	case reflect.Slice, reflect.Array:
		val := reflect.ValueOf(slice)

		ret := make([]interface{}, val.Len())
		for i := 0; i < val.Len(); i++ {
			ret[i] = val.Index(i).Interface()
		}

		return append(ret, i), nil
	}
}

// Prepend
//
// Adds and element to the beginning of the slice.
//
// Example: {{ prepend (slice "hello" "world" "!") "verbis" }}
// Returns: `[verbis hello world !]`
func (ns *Namespace) Prepend(slice interface{}, i interface{}) ([]interface{}, error) {
	const op = "Templates.append"

	typ := reflect.TypeOf(slice).Kind()

	switch typ {
	default:
		return nil, &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Cannot prepend on type: %s", typ), Operation: op, Err: fmt.Errorf("unable to prepend to slice with type: %s", typ)}
	case reflect.Slice, reflect.Array:
		val := reflect.ValueOf(slice)

		ret := make([]interface{}, val.Len())
		for i := 0; i < val.Len(); i++ {
			ret[i] = val.Index(i).Interface()
		}

		return append([]interface{}{i}, ret...), nil
	}
}

// First
//
// Retrieves the first element of the slice.
//
// Example: {{ first (slice "hello" "world" "!") }}
// Returns: `hello`
func (ns *Namespace) First(slice interface{}) (interface{}, error) {
	const op = "Templates.first"

	typ := reflect.TypeOf(slice).Kind()

	switch typ {
	default:
		return nil, &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Cannot get the last element on type: %s", typ), Operation: op, Err: fmt.Errorf("unable to get first element of slice with type: %s", typ)}
	case reflect.Slice, reflect.Array:
		val := reflect.ValueOf(slice)

		if val.Len() == 0 {
			return nil, nil
		}

		return val.Index(0).Interface(), nil
	}
}

// Last
//
// Retrieves the last element of the slice.
//
// Example: {{ last (slice "hello" "world" "!") }}
// Returns: `!`
func (ns *Namespace) Last(slice interface{}) (interface{}, error) {
	const op = "Templates.last"

	typ := reflect.TypeOf(slice).Kind()

	switch typ {
	case reflect.Slice, reflect.Array:
		val := reflect.ValueOf(slice)

		if val.Len() == 0 {
			return nil, nil
		}

		return val.Index(val.Len() - 1).Interface(), nil
	default:
		return nil, &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Cannot get the last element on type: %s", typ), Operation: op, Err: fmt.Errorf("unable to get last element of slice with type: %s", typ)}
	}
}

// Reverse
//
// Reverses the slice.
//
// Example: {{ reverse (slice "hello" "world" "!") }}
// Returns: `[! world hello]`
func (ns *Namespace) Reverse(slice interface{}) ([]interface{}, error) {
	const op = "Templates.reverse"

	typ := reflect.TypeOf(slice).Kind()

	switch typ {
	default:
		return nil, &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Cannot get the last element on type: %s", typ), Operation: op, Err: fmt.Errorf("unable to get reverse slice of type: %s", typ)}
	case reflect.Slice, reflect.Array:
		val := reflect.ValueOf(slice)

		reversed := make([]interface{}, val.Len())
		for i := 0; i < val.Len(); i++ {
			reversed[val.Len()-i-1] = val.Index(i).Interface()
		}

		return reversed, nil
	}
}
