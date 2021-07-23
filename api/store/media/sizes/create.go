// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sizes

import (
	"database/sql"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Create
//
// Returns nil if the media sizes were created successfully.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *Store) Create(mediaID int, sizes domain.MediaSizes) (domain.MediaSizes, error) {
	const op = "SizesStore.Create"

	for key, size := range sizes {
		q := s.Builder().
			Insert(s.Schema()+TableName).
			Column("file_id", size.FileId).
			Column("media_id", mediaID).
			Column("size_name", size.SizeName).
			Column("size_key", key).
			Column("width", size.Width).
			Column("height", size.Height).
			Column("crop", size.Crop)

		result, err := s.DB().Exec(q.Build())
		if err == sql.ErrNoRows {
			return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error creating media size with the key: " + key, Operation: op, Err: err}
		} else if err != nil {
			return nil, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
		}

		id, err := result.LastInsertId()
		if err != nil {
			return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error getting the newly created media size ID", Operation: op, Err: err}
		}

		size.Id = int(id)
		sizes[key] = size
	}

	return sizes, nil
}
