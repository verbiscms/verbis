// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
)

// Find
//
// Returns a form by searching with the given ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the form was not found by the given ID.
func (s *Store) Find(id int) (domain.Form, error) {
	const op = "FormStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("id", "=", id).
		Limit(1)

	var form domain.Form
	err := s.DB().Get(&form, q.Build())
	if err == sql.ErrNoRows {
		return domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No form exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	fields, err := s.fields.Find(form.Id)
	if err == nil {
		form.Fields = fields
	}

	submission, err := s.submissions.Find(form.Id)
	if err == nil {
		form.Submissions = submission
	}

	return form, nil
}

// FindByUUID
//
// Returns a form by searching with the given UUID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the form was not found by the given slug.
func (s *Store) FindByUUID(uniq uuid.UUID) (domain.Form, error) {
	const op = "FormStore.FindByUUID"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("uuid", "=", uniq.String()).
		Limit(1)

	var form domain.Form
	err := s.DB().Get(&form, q.Build())
	if err == sql.ErrNoRows {
		return domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: "No form exists with the UUID: " + uniq.String(), Operation: op, Err: err}
	} else if err != nil {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	fields, err := s.fields.Find(form.Id)
	if err != nil {
		return domain.Form{}, err
	}
	form.Fields = fields

	submission, err := s.submissions.Find(form.Id)
	if err == nil {
		form.Submissions = submission
	}

	return form, nil
}
