// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package submissions

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Create
//
// Returns nil if the submission was successfully created.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) Create(f domain.FormSubmission) error {
	const op = "FieldStore.Create"

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("uuid", "?").
		Column("form_id", f.FormId).
		Column("fields", "?").
		Column("ip_address", f.IPAddress).
		Column("user_agent", f.UserAgent).
		Column("sent_at", "NOW()")

	_, err := s.DB().Exec(q.Build(), uuid.New().String(), f.Fields)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error creating form submission with the form ID: %d", f.Id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
