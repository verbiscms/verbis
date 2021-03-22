// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Map
//
// Returns options as a map.
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no options available.
func (s *Store) Map() (domain.OptionsDBMap, error) {
	const op = "OptionStore.Map"

	q := s.Builder().
		From(s.Schema() + TableName)

	var o domain.OptionsDB
	err := s.DB().Select(&o, q.Build())
	if err == sql.ErrNoRows {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No options available", Operation: op, Err: err}
	} else if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	opts := make(domain.OptionsDBMap)
	for _, v := range o {
		unValue, err := s.unmarshal(&v.Value)
		if err != nil {
			return nil, err
		}
		opts[v.Name] = unValue
	}

	return opts, nil
}
