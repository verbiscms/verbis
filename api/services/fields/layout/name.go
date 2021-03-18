// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package layout

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

// walkerByName
//
// Uses recursion to locate a field by name by comparing the given name
// and the field's name, the sub fields (repeaters) and the sub
// fields of flexible content.
// Returns a domain.Field and true if it was found.
// Returns false if it wasn't.
func walkerByName(name string, field domain.Field) (domain.Field, bool) {
	// Account for normal field
	if field.Name == name {
		return field, true
	}

	// Account for repeaters
	if field.SubFields != nil {
		for _, subField := range field.SubFields {
			if f, found := walkerByName(name, subField); found {
				return f, true
			}
		}
	}

	// Account for flexible content
	if len(field.Layouts) != 0 {
		for _, layout := range field.Layouts {
			for _, subField := range layout.SubFields {
				if f, found := walkerByName(name, subField); found {
					return f, true
				}
			}
		}
	}

	// Field not found
	return domain.Field{}, false
}
