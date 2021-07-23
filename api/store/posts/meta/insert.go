// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"github.com/verbiscms/verbis/api/domain"
)

// Insert
//
// Checks to see if the post options record exists
// before updating or creating the new record.
func (s *Store) Insert(id int, p domain.PostOptions) error {
	if s.Exists(id) {
		err := s.update(id, p)
		if err != nil {
			return err
		}
	} else {
		err := s.create(id, p)
		if err != nil {
			return err
		}
	}
	return nil
}
