package templates

import (
	"github.com/spf13/cast"
	"math"
)

// add
//
// Returns a range of numbers that have been added in `int64`
//
// Example: {{ add 2 2 }} Returns 4
func (t *TemplateFunctions) add(i ...interface{}) int64 {
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
func (t *TemplateFunctions) subtract(a, b interface{}) int64 {
	return cast.ToInt64(a) - cast.ToInt64(b)
}

// divide
//
// Returns divided numbers in `int64`
//
// Example: {{ divide 16 4 }} Results in 4
func (t *TemplateFunctions) divide(a, b interface{}) int64 {
	return cast.ToInt64(a) / cast.ToInt64(b)
}

// multiply
//
// Returns a range of numbers that have been multiplied in `int64`
//
// Example: {{ add 4 4 }} Results in 16
func (t *TemplateFunctions) multiply(a interface{}, i ...interface{}) int64 {
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
func (t *TemplateFunctions) modulus(a, b interface{}) int64 {
	return cast.ToInt64(a) % cast.ToInt64(b)
}

// round
//
// Rounds up to the nearest integer rounding halfway
// from zero. Returns `float 64`.
//
// Example: {{ round 10.2 }} Results in 10
func (t *TemplateFunctions) round(i interface{}) float64 {
	return math.Round(cast.ToFloat64(i))
}

// ceil
//
// Rounds up to the nearest float value, returns `float 64`
//
// Example: {{ ceil 9.32 }} Results in 10
func (t *TemplateFunctions) ceil(i interface{}) float64 {
	return math.Ceil(cast.ToFloat64(i))
}

// floor
//
// Rounds down to the nearest float value , returns `float 64`
//
// Example: {{ floor 9.62 }} Results in 9
func (t *TemplateFunctions) floor(i interface{}) float64 {
	return math.Floor(cast.ToFloat64(i))
}

// min
//
// Finds the smallest numeric value in a slice of numbers, returns `int64`
//
// Example: {{ min 20 1 100 }} Results in 1
func (t *TemplateFunctions) min(a interface{}, i ...interface{}) int64 {
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
func (t *TemplateFunctions) max(a interface{}, i ...interface{}) int64 {
	val := cast.ToInt64(a)
	for _, a := range i {
		b := cast.ToInt64(a)
		if cast.ToInt64(a) > val {
			val = b
		}
	}
	return val
}
