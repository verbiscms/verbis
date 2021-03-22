// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

// Insert
//
// Updates or creates a new option depending on if it
// already exists in the database.
func (s *Store) Insert(options domain.OptionsDBMap) error {
	for name, value := range options {
		if s.Exists(name) {
			err := s.update(name, value)
			if err != nil {
				return err
			}
		} else {
			err := s.create(name, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
