// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package submissions

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Find
//
// Returns a form submission belonging to a form.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there was none found with the form ID available.
func (s *Store) Find(formID int) (domain.FormSubmission, error) {
	const op = "SubmissionStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("form_id", "=", formID)

	var submission domain.FormSubmission
	err := s.DB().Get(&submission, q.Build())
	if err == sql.ErrNoRows {
		return domain.FormSubmission{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No form submission exists with the form ID: %d", formID), Operation: op, Err: err}
	} else if err != nil {
		return domain.FormSubmission{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return submission, nil
}
