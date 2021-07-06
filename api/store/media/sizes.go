// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/store/files"
)

// sizes
//
// Returns a media item by searching with the given name.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the media was not found by the given name.
func (s *Store) sizes(mediaId int) (domain.MediaSizes, error) {
	const op = "MediaStore.Sizes"

	q := s.Builder().
		From(s.Schema()+TableSizesName).
		Where("media_id", "=", mediaId).
		LeftJoin(s.Schema()+files.TableName, "f", s.Schema()+"media.storage_id = "+s.Schema()+"f.id")

	var sizes []domain.MediaSize
	err := s.DB().Get(&sizes, q.Build())
	if err == sql.ErrNoRows {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No media size exists with the ID: %d", mediaId), Operation: op, Err: err}
	} else if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	var ms = make(domain.MediaSizes)
	for _, size := range sizes {
		ms[size.Key] = size
	}

	return ms, nil
}
