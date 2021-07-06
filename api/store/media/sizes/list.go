// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sizes

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/store/files"
	"github.com/ainsleyclark/verbis/api/store/media"
)

// List
//
// Returns a media sizes by searching with the given ID.
// Returns errors.INTERNAL if there was an error executing the query.
func (s *Store) List(mediaId int) (domain.MediaSizes, error) {
	const op = "MediaStore.Sizes"

	q := s.Builder().
		From(s.Schema()+media.TableSizesName).
		Where("media_id", "=", mediaId).
		LeftJoin(s.Schema()+files.TableName, "f", s.Schema()+"media.storage_id = "+s.Schema()+"f.id")

	var sizes []domain.MediaSize
	err := s.DB().Get(&sizes, q.Build())
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	var ms = make(domain.MediaSizes)
	for _, size := range sizes {
		ms[size.Key] = size
	}

	return ms, nil
}
