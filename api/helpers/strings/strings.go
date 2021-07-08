// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"bytes"
	"math/rand"
	"strings"
	"unicode"
)

// InSlice checks if a string exists in a slice,
func InSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Between gets substring between two strings.
func Between(value, a, b string) string {
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}

	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}

	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}

	return value[posFirstAdjusted:posLast]
}

// AddSpace between uppercase letters, for example
// HelloWorld will convert to Hello World.
func AddSpace(s string) string {
	buf := &bytes.Buffer{}
	for i, rune := range s {
		if unicode.IsUpper(rune) && i > 0 {
			buf.WriteRune(' ')
		}
		buf.WriteRune(rune)
	}
	return buf.String()
}

// Random generates a random string of n length. Contains
// numeric characters if set to true.
func Random(n int64, numeric bool) string {
	var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if !numeric {
		characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}

// TrimSlashes removes the prefix and suffix of a forward
// slash "/" from the given string.
func TrimSlashes(s string) string {
	return strings.TrimSuffix(strings.TrimPrefix(s, "/"), "/")
}
