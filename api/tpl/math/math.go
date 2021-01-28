package math

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
	"math"
)

// add
//
// Returns a range of numbers that have been added in `int64`
//
// Example: {{ add 2 2 }} Returns 4
func (ns *Namespace) add(i ...interface{}) int64 {
	var num int64 = 0
	for _, a := range i {
		num += cast.ToInt64(a)
	}
	return num
}

// subtract
//
// Returns subtracted numbers in `int64`
//
// Example: {{ subtract 100 10 }} Results in 90
func (ns *Namespace) subtract(a, b interface{}) int64 {
	return cast.ToInt64(a) - cast.ToInt64(b)
}

// divide
//
// Returns divided numbers in `int64`
//
// Example: {{ divide 16 4 }} Results in 4
func (ns *Namespace) divide(a, b interface{}) (int64, error) {
	const op = "Templates.divide"

	aa := cast.ToInt64(a)
	bb := cast.ToInt64(b)

	if bb == 0 {
		return 0, &errors.Error{Code: errors.TEMPLATE, Message: "Cannot divide by zero", Operation: op, Err: fmt.Errorf("integer divide by zero")}
	}

	return aa / bb, nil
}

// multiply
//
// Returns a range of numbers that have been multiplied in `int64`
//
// Example: {{ multiply 4 4 }} Results in 16
func (ns *Namespace) multiply(a interface{}, i ...interface{}) int64 {
	val := cast.ToInt64(a)
	for _, b := range i {
		val = val * cast.ToInt64(b)
	}
	return val
}

// modulus
//
// Returns remainder of two numbers in `int64`
//
// Example: {{ mod 10 9 }} Results in 1.
func (ns *Namespace) modulus(a, b interface{}) (int64, error) {
	const op = "Templates.divide"

	aa := cast.ToInt64(a)
	bb := cast.ToInt64(b)

	if bb == 0 {
		return 0, &errors.Error{Code: errors.TEMPLATE, Message: "Cannot divide by zero", Operation: op, Err: fmt.Errorf("integer divide by zero")}
	}

	return aa % bb, nil
}

// round
//
// Rounds up to the nearest integer rounding halfway
// from zero. Returns `float 64`.
//
// Example: {{ round 10.2 }} Results in 10
func (ns *Namespace) round(i interface{}) float64 {
	return math.Round(cast.ToFloat64(i))
}

// ceil
//
// Rounds up to the nearest float value, returns `float 64`
//
// Example: {{ ceil 9.32 }} Results in 10
func (ns *Namespace) ceil(i interface{}) float64 {
	return math.Ceil(cast.ToFloat64(i))
}

// floor
//
// Rounds down to the nearest float value , returns `float 64`
//
// Example: {{ floor 9.62 }} Results in 9
func (ns *Namespace) floor(i interface{}) float64 {
	return math.Floor(cast.ToFloat64(i))
}

// min
//
// Finds the smallest numeric value in a slice of numbers, returns `int64`
//
// Example: {{ min 20 1 100 }} Results in 1
func (ns *Namespace) min(a interface{}, i ...interface{}) int64 {
	val := cast.ToInt64(a)
	for _, a := range i {
		b := cast.ToInt64(a)
		if cast.ToInt64(a) < val {
			val = b
		}
	}
	return val
}

// max
//
// Finds the largest numeric value in a slice of numbers, returns `int64`
//
// Example: {{ max 20 1 100 }} Results in 100
func (ns *Namespace) max(a interface{}, i ...interface{}) int64 {
	val := cast.ToInt64(a)
	for _, a := range i {
		b := cast.ToInt64(a)
		if cast.ToInt64(a) > val {
			val = b
		}
	}
	return val
}
