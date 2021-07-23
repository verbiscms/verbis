// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// create
//
// Returns a new post field upon creation.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) create(f domain.PostField) (domain.PostField, error) {
	const op = "FieldStore.Create"

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("uuid", "?").
		Column("post_id", f.PostId).
		Column("type", f.Type).
		Column("name", f.Name).
		Column("value", f.OriginalValue).
		Column("field_key", f.Key)

	_, err := s.DB().Exec(q.Build(), uuid.New().String())
	if err == sql.ErrNoRows {
		return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating field with the name: " + f.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	//id, err := result.LastInsertId()
	//if err != nil {
	//	return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created field ID", Operation: op, Err: err}
	//}
	//f.ID = int(id)

	return f, nil
}
