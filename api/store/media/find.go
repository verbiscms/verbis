// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"database/sql"
	"fmt"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Find
//
// Returns a media item by searching with the given ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the media item was not found by the given ID.
func (s *Store) Find(id int) (domain.Media, error) {
	const op = "MediaStore.Find"

	q := s.selectStmt().
		Where(TableName+".id", "=", id).
		Limit(1)

	var media domain.Media
	err := s.DB().Get(&media, q.Build())
	if err == sql.ErrNoRows {
		return domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No media item exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	sizes, err := s.sizes.Find(id)
	if err != nil {
		return domain.Media{}, err
	}
	media.Sizes = sizes

	return media, nil
}
