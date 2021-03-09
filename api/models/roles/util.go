// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// validate
//
// Returns nil if the validation passed on update/create.
// Returns errors.CONFLICT if the name already exists in the store.
func (s *Store) validate(r domain.Role) error {
	const op = "RoleStore.Validate"

	exists := s.Exists(r.Name)
	if exists {
		return &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Validation failed, the role ID already exists: %d", r.Id), Operation: op, Err: ErrRoleExists}
	}

	return nil
}
