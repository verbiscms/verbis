// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seeds

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/store"
)

type Seeder struct {
	db     database.Driver
	models *store.Repository
}

// Construct
func New(db database.Driver, s *store.Repository) *Seeder {
	return &Seeder{
		db:     db,
		models: s,
	}
}

// Seed
func (s *Seeder) Seed() error {
	// IMPORTANT: Run roles before inserting the user.
	if err := s.runRoles(); err != nil {
		return err
	}
	if err := s.runOptions(); err != nil {
		return err
	}
	// if err := s.runPosts(); err != nil {
	//	return err
	//}
	return nil
}
