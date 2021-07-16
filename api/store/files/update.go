// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"strconv"
)

// Update
//
// Updates a file with new storage information.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) Update(id int, f domain.StorageChange) error {
	const op = "FileStore.Create"

	q := s.Builder().
		Update(s.Schema()+TableName).
		Column("provider", f.Provider).
		Column("bucket", f.Bucket).
		Column("region", f.Region).
		Where("id", "=", id)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating file with the ID: " + strconv.Itoa(id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
