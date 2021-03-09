// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// validate
//
// Returns nil if the validation passed on update/create.
// Returns errors.CONFLICT if the name already exists in the store.
func (s *Store) validate(r domain.Redirect) error {
	const op = "RedirectStore.Validate"

	exists := s.ExistsByFrom(r.From)
	if exists {
		return &errors.Error{Code: errors.CONFLICT, Message: "Validation failed, the redirect from path already exists: " + r.From, Operation: op, Err: ErrRedirectExists}
	}

	return nil
}
