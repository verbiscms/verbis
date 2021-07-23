// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"strings"
)

// number casts the domain.FieldValue to int64.
// Returns errors.INVALID if the domain.FieldValue could not be cast to an int64.
func (v *Value) number(value domain.FieldValue) (interface{}, error) {
	const op = "FieldResolver.Number"

	val := value.String()

	if strings.Contains(val, ".") {
		f, err := cast.ToFloat64E(val)
		if err != nil {
			return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast field to an integer", Operation: op, Err: err}
		}
		return f, nil
	}

	n, err := cast.ToInt64E(val)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast field to an integer", Operation: op, Err: err}
	}
	return n, nil
}
