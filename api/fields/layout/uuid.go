// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package layout

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/google/uuid"
)

// walkerByUUID
//
// Uses recursion to locate a field by UUID by comparing the given UUID
// and the field's UUID, the sub fields (repeaters) and the sub
// fields of flexible content.
// Returns a domain.Field and true if it was found.
// Returns false if it wasn't.
func walkerByUUID(uniq uuid.UUID, field domain.Field) (domain.Field, bool) {
	// Account for normal field
	if field.UUID == uniq {
		return field, true
	}

	// Account for repeaters
	if field.SubFields != nil {
		for _, subField := range field.SubFields {
			if f, found := walkerByUUID(uniq, subField); found {
				return f, true
			}
		}
	}

	// Account for flexible content
	if len(field.Layouts) != 0 {
		for _, layout := range field.Layouts {
			for _, subField := range layout.SubFields {
				if f, found := walkerByUUID(uniq, subField); found {
					return f, true
				}
			}
		}
	}

	// Field not found
	return domain.Field{}, false
}
