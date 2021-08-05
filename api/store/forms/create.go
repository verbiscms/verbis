// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
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
	f.ID = int(id)

	for _, v := range f.Fields {
		err := s.fields.Insert(f.ID, v)
		if err != nil {
			return domain.Form{}, err
		}
	}

	submissions, err := s.submissions.Find(f.ID)
	if err == nil {
		f.Submissions = submissions
	}

	return f, nil
}
