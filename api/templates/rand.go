package templates

import (
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/spf13/cast"
	"math/rand"
	"time"
)

// randInt
//
// Returns a random integer between min and max values.
func (t *TemplateFunctions) randInt(a, b interface{}) int {
	min := cast.ToInt(a)
	max := cast.ToInt(b)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

// randFloat
//
// Returns a random float between min and max values.
func (t *TemplateFunctions) randFloat(a, b interface{}) float64 {
	min := cast.ToFloat64(a)
	max := cast.ToFloat64(b)
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64() * (max - min)
}

// randAlpha
//
// Returns a random alpha string by the given length.
func (t *TemplateFunctions) randAlpha(len int) string {
	return encryption.RandomString(len, false)
}

// randAlphaNum
//
// Returns a random alpha numeric string by the given length.
func (t *TemplateFunctions) randAlphaNum(len int) string {
	return encryption.RandomString(len, true)
}