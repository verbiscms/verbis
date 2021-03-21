// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
)

// Create
//
// Returns a new form upon creation.
// Returns errors.CONFLICT if the the form (name) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(f domain.Form) (domain.Form, error) {
	const op = "FormStore.Create"

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("uuid", "?").
		Column("name", f.Name).
		Column("email_send", f.EmailSend).
		Column("email_message", f.EmailMessage).
		Column("email_subject", f.EmailSubject).
		Column("store_db", f.StoreDB).
		Column("updated_at", "NOW()").
		Column("created_at", "NOW()")

	result, err := s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating form with the name: " + f.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created form ID", Operation: op, Err: err}
	}
	f.Id = int(id)

	// TODO, Insert/Update into Form Fields

	return f, nil
}
