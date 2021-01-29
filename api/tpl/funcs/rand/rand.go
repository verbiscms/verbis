package rand

import (
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/spf13/cast"
	"math/rand"
	"time"
)

// Int
//
// Returns a random integer between min and max values.
//
// Example: {{ randInt 1 10 }}
func (ns *Namespace) Int(a, b interface{}) int {
	min, err := cast.ToIntE(a)
	if err != nil || a == nil {
		return 0
	}

	max, err := cast.ToIntE(b)
	if err != nil || b == nil {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Float
//
// Returns a random float between min and max values.
//
// Example: {{ randFloat 2.5 10 }}
func (ns *Namespace) Float(a, b interface{}) float64 {
	min, err := cast.ToFloat64E(a)
	if err != nil || a == nil {
		return 0.0
	}

	max, err := cast.ToFloat64E(b)
	if err != nil || b == nil {
		return 0.0
	}

	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

// Alpha
//
// Returns a random alpha string by the given length.
//
// Example: {{ randAlpha 20 }}
func (ns *Namespace) Alpha(len int64) string {
	return encryption.RandomString(len, false)
}

// AlphaNum
//
// Returns a random alpha numeric string by the given length.
//
// Example: {{ randAlphaNum 20 }}
func (ns *Namespace) AlphaNum(len int64) string {
	return encryption.RandomString(len, true)
}
