// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// validate
//
// Returns nil if the validation passed on update/create.
// Returns errors.CONFLICT if the name already exists in the store.
func (s *Store) validate(c domain.Category) error {
	const op = "CategoryStore.Validate"

	exists := s.ExistsByName(c.Name)
	if exists {
		return &errors.Error{Code: errors.CONFLICT, Message: "Validation failed, the category name already exists: " + c.Name, Operation: op, Err: ErrCategoryExists}
	}

	return nil
}
