// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

// tags defines the the array of values that are held
// in the `tags` field type.
type tags []string

// tags
//
// Uses the Array() function on the domain.FieldValue type to split
// the value by a comma delimiter, and loops over the values to
// build up a tags array to be sent back.
func (v *Value) tags(value domain.FieldValue) (interface{}, error) {
	var t tags
	for _, v := range value.Array() {
		if v != "" {
			t = append(t, v)
		}
	}
	return t, nil
}
