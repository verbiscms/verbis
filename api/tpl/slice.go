package tpl

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"reflect"
)

// slice
//
// Creates a slice (array) of passed arguments.
//
// Example: Returns `!`
// {{ $mySlice := slice "hello" "world" "!" }}
// {{ index $mySlice 2 }}
func (t *TemplateManager) slice(i ...interface{}) []interface{} {
	return i
}

// append
//
// Adds and element to the end of the slice.
//
// Example: Returns `1234`
// {{ $mySlice := slice 1 2 3 }}
// {{ append $mySlice 4 }}
func (t *TemplateManager) append(slice interface{}, i interface{}) ([]interface{}, error) {
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

// prepend
//
// Adds and element to the beginning of the slice.
//
// Example: Returns 4123
// {{ $mySlice := slice 1 2 3 }}
// {{ prepend $mySlice 4 }}
func (t *TemplateManager) prepend(slice interface{}, i interface{}) ([]interface{}, error) {
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

// first
//
// Retrieves the first element of the slice.
//
// Example: Returns `1`
// {{ $mySlice := slice 1 2 3 }}
// {{ first $mySlice }}
func (t *TemplateManager) first(slice interface{}) (interface{}, error) {
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

// last
//
// Retrieves the last element of the slice.
//
// Example: Returns `3`
// {{ $mySlice := slice 1 2 3 }}
// {{ last $mySlice }}
func (t *TemplateManager) last(slice interface{}) (interface{}, error) {
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

// reverse
//
// Reverses the slice.
//
// Example: Returns `321`
// {{ $mySlice := slice 1 2 3 }}
// {{ reverse $mySlice }}
func (t *TemplateManager) reverse(slice interface{}) ([]interface{}, error) {
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
