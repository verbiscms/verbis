// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// checkbox casts the domain.FieldValue to a boolean.
// Returns errors.INVALID if the domain.FieldValue could not be cast to an bool.
func (v *Value) checkbox(value domain.FieldValue) (interface{}, error) {
	const op = "FieldResolver.checkbox"

	check, err := cast.ToBoolE(value.String())
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast checkbox field to an bool", Operation: op, Err: err}
	}

	return check, nil
}
