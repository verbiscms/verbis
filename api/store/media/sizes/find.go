// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sizes

import (
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Find
//
// Returns a media sizes by searching with the given media ID.
// Returns errors.INTERNAL if there was an error executing the query.
func (s *Store) Find(mediaID int) (domain.MediaSizes, error) {
	const op = "SizesStore.Find"

	q := s.selectStmt().
		Where(TableName+".media_id", "=", mediaID)

	var sizes []domain.MediaSize
	err := s.DB().Select(&sizes, q.Build())
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	var ms = make(domain.MediaSizes)
	for _, size := range sizes {
		ms[size.SizeKey] = size
	}

	return ms, nil
}
