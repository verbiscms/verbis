// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/services/forms"
	"github.com/google/uuid"
)

// Submit
//
//
func (s *Store) Submit(form domain.Form, sub domain.FormSubmission, values forms.FormValues) error {
	const op = "FormStore.Submit"

	formValues, err := values.JSON()
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not process the form fields for storing", Operation: op, Err: err}
	}

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("uuid", "?").
		Column("form_id", form.Id).
		Column("fields", formValues).
		Column("ip", sub.IPAddress).
		Column("agent", sub.UserAgent).
		Column("sent_at", "NOW()")

	_, err = s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error creating form submission with the name: " + form.Name, Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
