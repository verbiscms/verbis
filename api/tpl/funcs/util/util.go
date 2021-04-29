// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

// Len
//
// Returns the length of a variable according to its type.
// If the length of the type passed could not be
// retrieved, it will return `0`.
//
// Example: {{ len "hello" }}
// Returns: `5`
func (ns *Namespace) Len(i interface{}) int64 {
	if i == nil {
		return 0
	}

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

// Explode
//
// Breaks a string into array with a delimiter (separator).
//
// Example: {{ explode "," "hello there !" }}
// Returns: `[hello there !]`
func (ns *Namespace) Explode(delim, text interface{}) []string {
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

// Implode
//
// Returns a string from the elements of an array using a
// glue string to join them together.
//
// Example: {{ slice 1 2 3 | explode "," }}
// Returns: `[1 2 3]`
func (ns *Namespace) Implode(glue, slice interface{}) string {
	str, err := cast.ToStringE(glue)
	if err != nil || slice == nil {
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

// Seq
//
// Creates a sequence of integers.
//
// Example: {{ seq 5 }}
// Returns: `[1 2 3 4 5]`
func (ns *Namespace) Seq(size interface{}) ([]int64, error) {
	const op = "Templates.Seq"

	s, err := cast.ToInt64E(size)
	if err != nil {
		return nil, &errors.Error{Code: errors.TEMPLATE, Message: "Error casting interface to integer", Operation: op, Err: err}
	}

	if s <= 0 {
		return nil, &errors.Error{Code: errors.TEMPLATE, Message: "Sequence cannot be negative", Operation: op, Err: fmt.Errorf("sequence must be > 0")}
	}

	seq := make([]int64, 0)
	for i := int64(0); i < s; i++ {
		seq = append(seq, i)
	}

	return seq, nil
}
