// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

// checkOwner
//
// checkOwner Checks if the author is set or if the author does not exist.
// Returns the owner ID under circumstances.
func (s *Store) checkOwner(id int) int {
	if id == 0 || !s.users.Exists(id) {
		return s.Owner.ID
	}
	return id
}
