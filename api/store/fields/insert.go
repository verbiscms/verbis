// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/verbiscms/verbis/api/domain"
)

// Insert
//
// Insert checks to see if the record should be
// deleted before updating or creating the
// new record depending on if the field
// exists in the store.
func (s *Store) Insert(postID int, fields domain.PostFields) error {
	err := s.Delete(postID)
	if err != nil {
		return err
	}

	for _, v := range fields {
		v.PostId = postID
		_, err := s.create(v)
		if err != nil {
			return err
		}
	}

	return nil
}
