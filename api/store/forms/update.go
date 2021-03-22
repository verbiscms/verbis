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

// Update
//
// Returns an updated form.
// Returns errors.CONFLICT if the validation failed.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not obtain the newly created ID.
func (s *Store) Update(f domain.Form) (domain.Form, error) {
	const op = "FormStore.Create"

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("uuid", "?").
		Column("name", f.Name).
		Column("email_send", f.EmailSend).
		Column("email_message", f.EmailMessage).
		Column("email_subject", f.EmailSubject).
		Column("store_db", f.StoreDB).
		Column("updated_at", "NOW()").
		Where("id", "=", f.Id)

	_, err := s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating form with the name: " + f.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	for _, v := range f.Fields {
		err := s.fields.Insert(f.Id, v)
		if err != nil {
			return domain.Form{}, err
		}
	}

	submissions, err := s.submissions.Find(f.Id)
	if err == nil {
		f.Submissions = submissions
	}

	return f, nil
}
