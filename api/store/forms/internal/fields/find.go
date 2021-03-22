// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Find
//
// Returns a slice of form fields belonging to a form.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no form fields available.
func (s *Store) Find(formID int) (domain.FormFields, error) {
	const op = "FieldStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("form_id", "=", formID)

	var fields domain.FormFields
	err := s.DB().Select(&fields, q.Build())
	if err == sql.ErrNoRows {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No form fields exists with the form ID: %d", formID), Operation: op, Err: err}
	} else if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return fields, nil
}
