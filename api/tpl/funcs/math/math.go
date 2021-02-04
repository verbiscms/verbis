// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
	"math"
)

// Add
//
// Returns a range of numbers that have been added in `int64`
//
// Example: {{ add 2 2 }}
// Returns: `4`
func (ns *Namespace) Add(i ...interface{}) int64 {
	var num int64 = 0
	for _, a := range i {
		num += cast.ToInt64(a)
	}
	return num
}

// Subtract
//
// Returns subtracted numbers in `int64`
//
// Example: {{ subtract 100 10 }}
// Returns: `90`
func (ns *Namespace) Subtract(a, b interface{}) int64 {
	return cast.ToInt64(a) - cast.ToInt64(b)
}

// Divide
//
// Returns divided numbers in `int64`
//
// Example: {{ divide 16 4 }}
// Returns: `4`
func (ns *Namespace) Divide(a, b interface{}) (int64, error) {
	const op = "Templates.Divide"

	aa := cast.ToInt64(a)
	bb := cast.ToInt64(b)

	if bb == 0 {
		return 0, &errors.Error{Code: errors.TEMPLATE, Message: "Cannot divide by zero", Operation: op, Err: fmt.Errorf("integer divide by zero")}
	}

	return aa / bb, nil
}

// Multiply
//
// Returns a range of numbers that have been multiplied in `int64`
//
// Example: {{ multiply 4 4 }}
// Returns: `16`
func (ns *Namespace) Multiply(a interface{}, i ...interface{}) int64 {
	val := cast.ToInt64(a)
	for _, b := range i {
		val = val * cast.ToInt64(b)
	}
	return val
}

// Modulus
//
// Returns remainder of two numbers in `int64`
//
// Example: {{ mod 10 9 }}
// Returns: `1`
func (ns *Namespace) Modulus(a, b interface{}) (int64, error) {
	const op = "Templates.Modulus"

	aa := cast.ToInt64(a)
	bb := cast.ToInt64(b)

	if bb == 0 {
		return 0, &errors.Error{Code: errors.TEMPLATE, Message: "Cannot divide by zero", Operation: op, Err: fmt.Errorf("integer divide by zero")}
	}

	return aa % bb, nil
}

// Round
//
// Rounds up to the nearest integer rounding halfway
// from zero. Returns `float 64`.
//
// Example: {{ round 10.2 }}
// Returns: `10`
func (ns *Namespace) Round(i interface{}) float64 {
	return math.Round(cast.ToFloat64(i))
}

// Ceil
//
// Rounds up to the nearest float value, returns `float 64`
//
// Example: {{ ceil 9.32 }}
// Returns: `10`
func (ns *Namespace) Ceil(i interface{}) float64 {
	return math.Ceil(cast.ToFloat64(i))
}

// Floor
//
// Rounds down to the nearest float value , returns `float 64`
//
// Example: {{ floor 9.62 }}
// Returns: `9`
func (ns *Namespace) Floor(i interface{}) float64 {
	return math.Floor(cast.ToFloat64(i))
}

// Min
//
// Finds the smallest numeric value in a slice of numbers, returns `int64`
//
// Example: {{ min 20 1 100 }}
// Returns: `1`
func (ns *Namespace) Min(a interface{}, i ...interface{}) int64 {
	val := cast.ToInt64(a)
	for _, a := range i {
		b := cast.ToInt64(a)
		if cast.ToInt64(a) < val {
			val = b
		}
	}
	return val
}

// Max
//
// Finds the largest numeric value in a slice of numbers, returns `int64`
//
// Example: {{ max 20 1 100 }}
// Returns: `100`
func (ns *Namespace) Max(a interface{}, i ...interface{}) int64 {
	val := cast.ToInt64(a)
	for _, a := range i {
		b := cast.ToInt64(a)
		if cast.ToInt64(a) > val {
			val = b
		}
	}
	return val
}
