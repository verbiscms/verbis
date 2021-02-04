// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

import (
	"fmt"
	"reflect"
)

// KindIs
//
// Returns true if the target and source types match.
//
// Returns `true`
// Example: {{ kindIs "int" 123 }}
func (ns *Namespace) KindIs(target string, src interface{}) bool {
	return target == ns.KindOf(src)
}

// KindOf
//
// Returns the type of the given `interface` as a string.
//
// Example: {{ kindOf 123 }}
// Returns: `int`
func (ns *Namespace) KindOf(src interface{}) string {
	return reflect.ValueOf(src).Kind().String()
}

// TypeOf
//
// Returns the underlying type of the given value.
//
// Example: {{ typeOf .Post }}
// Returns: `domain.Post`
func (ns *Namespace) TypeOf(src interface{}) string {
	return fmt.Sprintf("%T", src)
}

// TypeIs
//
// Similar to `kindIs` but its used for types, instead of primitives.
//
// Example: {{ typeOf "domain.Post" .Post }}
// Returns: `true`
func (ns *Namespace) TypeIs(target string, src interface{}) bool {
	return target == ns.TypeOf(src)
}

// TypeIsLike
//
// Similar to `kindIs` but its used for types, instead of primitives.
//
// Example: {{ typeOf "domain.Post" .Post }}
// Returns: true
func (ns *Namespace) TypeIsLike(target string, src interface{}) bool {
	s := ns.TypeOf(src)
	return target == s || "*"+target == s
}
