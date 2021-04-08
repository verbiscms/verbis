// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// update
//
// Returns nil if the meta was successfully updated.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) update(id int, p domain.PostOptions) error {
	const op = "MetaStore.Update"

	// No support for marshalling json for builder currently.
	q := "UPDATE " + s.Schema() + "post_options SET seo = ?, meta = ? WHERE post_id = ?"

	_, err := s.DB().Exec(q, p.Seo, p.Meta, id)
	if err == sql.ErrNoRows {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error updating meta with the post ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return nil
}
