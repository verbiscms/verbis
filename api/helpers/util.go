// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package helpers

import (
	"bytes"
	"strings"
	"unicode"
)

// Check if a string exists in a slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Check if a string exists in a slice
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Split everything before deliminator
func StringsSplitLeft(str, delim string) string {
	return strings.Split(str, delim)[0]
}

// Split everything after deliminator
func StringsSplitRight(str, delim string) string {
	return strings.Join(strings.Split(str, delim)[1:], delim)
}

// Between Gets substring between two strings.
func StringsBetween(value, a, b string) string {
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

// Add space between uppercase letters, for example
// HelloWorld will convert to Hello World.
func StringsAddSpace(s string) string {
	buf := &bytes.Buffer{}
	for i, rune := range s {
		if unicode.IsUpper(rune) && i > 0 {
			buf.WriteRune(' ')
		}
		buf.WriteRune(rune)
	}
	return buf.String()
}
