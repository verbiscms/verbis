// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
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

	// Find fields that should be deleted (not in the array)
	//f, err := s.Find(postID)
	//if err != nil {
	//	return err
	//}

	//for _, v := range f {
	//	if s.shouldDelete(v, fields) {
	//		err := s.deleteField(postID, v)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//}

	// Update or create the existing fields passed.
	//for _, v := range fields {
	//	v.PostId = postID
	//	if s.Exists(v) {
	//		_, err := s.update(v)
	//		if err != nil {
	//			return err
	//		}
	//	} else {
	//		_, err := s.create(v)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//}

	return nil
}

// shouldDelete
//
// Finds fields in the domain.PostField array that should
// be deleted.
func (s *Store) shouldDelete(f domain.PostField, fields domain.PostFields) bool {
	for _, v := range fields {
		if (f.Key == v.Key) && (f.UUID == v.UUID) && (f.Name == v.Name) {
			return false
		}

	}
	return true
}
