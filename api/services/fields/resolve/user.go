// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// user resolves a user from the given value.
// Returns the domain.User if it was found and no error occurred.
// Returns errors.INVALID if the domain.FieldValue could not be cast to an integer.
func (v *Value) user(value domain.FieldValue) (interface{}, error) {
	const op = "FieldResolver.user"

	id, err := value.Int()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast user ID to an integer", Operation: op, Err: err}
	}

	user, err := v.deps.Store.User.Find(id)
	if err != nil {
		return nil, err
	}

	return user.HideCredentials(), nil
}
