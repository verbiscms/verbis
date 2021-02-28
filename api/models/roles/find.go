// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Find
//
//
func (s *Store) Find(id int64) (domain.Role, error) {
	const op = "Redirects.Find"

	q := s.Builder.Select("*").From(TableName).WhereRaw("`id` = ?").Limit(1)

	var r domain.Role
	err := s.Get(&r, q.Build(), id)
	if err == sql.ErrNoRows {
		return domain.Role{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No redirect exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.Role{}, &errors.Error{Code: errors.INTERNAL, Message: "Error executing sql query", Operation: op, Err: err}
	}

	return r, nil
}
