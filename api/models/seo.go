// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// SeoMetaRepository defines methods for Posts to interact with the database
type SeoMetaRepository interface {
	UpdateCreate(id int, p domain.PostOptions) error
}

// SeoMetaStore defines the data layer for Seo & Meta Options
type SeoMetaStore struct {
	*StoreCfgOld
}

// newSeoMeta - Construct
func newSeoMeta(cfg *StoreCfgOld) *SeoMetaStore {
	return &SeoMetaStore{
		StoreCfgOld: cfg,
	}
}

// UpdateCreate checks to see if the record exists before updating
// or creating the new record.
func (s *SeoMetaStore) UpdateCreate(id int, p domain.PostOptions) error {
	if s.exists(id) {
		if err := s.update(id, p); err != nil {
			return err
		}
	} else {
		if err := s.create(id, p); err != nil {
			return err
		}
	}
	return nil
}

// Exists Checks if a seo meta record exists by Id
func (s *SeoMetaStore) exists(id int) bool {
	var exists bool
	_ = s.DB.QueryRow("SELECT EXISTS (SELECT post_id FROM post_options WHERE post_id = ?)", id).Scan(&exists)
	return exists
}

// create a new seo meta record
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *SeoMetaStore) create(id int, p domain.PostOptions) error {
	const op = "SeoMetaRepository.create"
	_, err := s.DB.Exec("INSERT INTO post_options (post_id, seo, meta) VALUES (?, ?, ?)", id, p.Seo, p.Meta)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the seo meta options record with the post id: %d", id), Operation: op, Err: err}
	}
	return nil
}

// update a seo meta record by page Id
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *SeoMetaStore) update(id int, p domain.PostOptions) error {
	const op = "SeoMetaRepository.update"
	_, err := s.DB.Exec("UPDATE post_options SET seo = ?, meta = ? WHERE post_id = ?", p.Seo, p.Meta, id)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the seo meta options with the post id: %d", id), Operation: op, Err: err}
	}
	return nil
}
