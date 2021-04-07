// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

// Insert
//
// Checks to see if the post category record exists
// before updating or creating the new record.
func (s *Store) Insert(postID int, catID *int) error {
	if s.Exists(postID) {
		if catID == nil {
			err := s.Delete(postID)
			if err != nil {
				return err
			}
			return nil
		}

		err := s.update(postID, *catID)
		if err != nil {
			return err
		}
	} else {
		if catID == nil {
			return nil
		}

		err := s.create(postID, *catID)
		if err != nil {
			return err
		}
	}

	return nil
}
