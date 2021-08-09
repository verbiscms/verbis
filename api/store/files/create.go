// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"database/sql"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Create
//
// Returns a new file upon creation.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(f domain.File) (domain.File, error) {
	const op = "FileStore.Create"

	q := s.Builder().
		Insert(s.Schema()+TableName).
		Column("uuid", f.UUID.String()).
		Column("url", f.URL).
		Column("name", f.Name).
		Column("bucket_id", f.BucketID).
		Column("mime", f.Mime).
		Column("source_type", f.SourceType).
		Column("provider", f.Provider).
		Column("region", f.Region).
		Column("bucket", f.Bucket).
		Column("file_size", f.FileSize).
		Column("private", f.Private)

	result, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating file with the name: " + f.Name, Operation: op, Err: err}
	} else if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created file ID", Operation: op, Err: err}
	}
	f.ID = int(id)

	return f, nil
}
