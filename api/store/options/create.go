// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"database/sql"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/errors"
)

// Create
//
// Returns a nil upon creation.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(name string, value interface{}) error {
	const op = "OptionStore.Create"

	v, err := s.marshal(value)
	if err != nil {
		return err
	}

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("option_name", name).
		Column("option_value", "?")

	_, err = s.DB().Exec(q.Build(), v)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error creating option with the name: " + name, Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
