// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sizes

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Delete
//
// Returns nil if the media sizes were successfully deleted.
// Returns errors.INTERNAL if there was an error executing the query.
func (s *Store) Delete(mediaID int) error {
	const op = "SizesStore.Delete"

	q := s.Builder().
		DeleteFrom(s.Schema()+TableName).
		Where("media_id", "=", mediaID)

	_, err := s.DB().Exec(q.Build())
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No media sizes exists with the media ID: %d", mediaID), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
