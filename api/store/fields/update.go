// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// update
//
// Returns an updated post field.
// Returns errors.INTERNAL if the SQL query was invalid or no rows were effected.
func (s *Store) update(f domain.PostField) (domain.PostField, error) {
	const op = "FieldStore.Update"

	// NOTE! Finding By UUID does not work, Vue passing wrong Data (UUID).

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("value", f.OriginalValue).
		Column("field_key", f.Key).
		Where("post_id", "=", f.PostId).
		Where("field_key", "=", f.Key).
		Where("name", "=", f.Name)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: "Error updating the post field name the uuid: " + f.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return f, nil
}
