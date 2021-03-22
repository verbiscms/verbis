// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"database/sql"
	//"encoding/json"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Find
//
// Returns a option by searching with the given name.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the option was not found by the given name.
func (s *Store) Find(name string) (interface{}, error) {
	const op = "OptionStore.Find"

	q := s.Builder().
		Select("option_value").
		From(s.Schema()+TableName).
		Where("option_name", "=", name).
		Limit(1)

	var value interface{}
	err := s.DB().Get(&value, q.Build())
	if err == sql.ErrNoRows {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No option exists with the name: " + name, Operation: op, Err: err}
	} else if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return value, nil
}
