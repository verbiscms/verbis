// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	location "github.com/ainsleyclark/verbis/api/services/fields/location"
	"github.com/google/uuid"
)

// FieldsRepository defines methods for Posts to interact with the database
type FieldsRepository interface {
	GetByPost(postID int) (domain.PostFields, error)
	GetLayout(post domain.PostDatum) domain.FieldGroups
	UpdateCreate(postID int, f domain.PostFields) error
	Create(f domain.PostField) (domain.PostField, error)
	Update(f domain.PostField) (domain.PostField, error)
	Exists(postID int, uuid uuid.UUID, key string, name string) bool
}

// FieldsStore defines the data layer for Posts
type FieldsStore struct {
	*StoreCfgOld
	options domain.Options
	finder  location.Finder
}

// newFields - Construct
func newFields(cfg *StoreCfgOld) *FieldsStore {
	return &FieldsStore{
		StoreCfgOld: cfg,
		options:     cfg.Options.GetStruct(),
		finder:      location.NewLocation(cfg.Paths.Storage),
	}
}

// UpdateCreate checks to see if the record exists before updating
// or creating the new record.
func (s *FieldsStore) UpdateCreate(postID int, f domain.PostFields) error {
	fields, err := s.GetByPost(postID)
	if err != nil {
		return err
	}

	// Find fields that should be deleted (not in the array)
	for _, v := range fields {
		if !s.shouldDelete(v, f) {
			err := s.Delete(postID, v)
			if err != nil {
				return err
			}
		}
	}

	// Update or create the existing fields passed.
	for _, v := range f {
		v.PostId = postID
		if s.Exists(postID, v.UUID, v.Key, v.Name) {
			_, err := s.Update(v)
			if err != nil {
				return err
			}
		} else {
			_, err := s.Create(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetByPost fields by a post ID.
// Returns errors.NOTFOUND if there were no records found.
func (s *FieldsStore) GetByPost(postID int) (domain.PostFields, error) {
	const op = "FieldsStore.GetByPost"
	var f domain.PostFields
	if err := s.DB.Select(&f, "SELECT * FROM post_fields WHERE post_id = ?", postID); err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post field with the ID: %d", postID), Operation: op, Err: err}
	}
	return f, nil
}

// GetByPost fields by a post ID and key.
// Returns errors.NOTFOUND if there were no records found.
func (s *FieldsStore) GetByPostAndKey(key string, postID int) (domain.PostField, error) {
	const op = "FieldsRepository.GetByPostAndKey"
	var f domain.PostField
	if err := s.DB.Select(&f, "SELECT * FROM post_fields WHERE post_id = ? AND field_key = ? LIMIT = 1", postID, key); err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post field with the page ID: %d and key: %s", postID, key), Operation: op, Err: err}
	}
	return f, nil
}

// Update a post field by ID
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *FieldsStore) Create(f domain.PostField) (domain.PostField, error) {
	const op = "FieldsRepository.Create"
	q := "INSERT INTO post_fields (uuid, post_id, type, name, value, field_key) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := s.DB.Exec(q, f.UUID.String(), f.PostId, f.Type, f.Name, f.OriginalValue, f.Key)
	if err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the post field wuth the name: %s", f.Name), Operation: op, Err: err}
	}
	return f, nil
}

// Update a post field by ID
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *FieldsStore) Update(f domain.PostField) (domain.PostField, error) {
	const op = "FieldsRepository.Update"
	_, err := s.DB.Exec("UPDATE post_fields SET type = ?, name = ?, value = ?, field_key = ? WHERE uuid = ? AND post_id = ? AND field_key = ? AND name = ?", f.Type, f.Name, f.OriginalValue, f.Key, f.UUID.String(), f.PostId, f.Key, f.Name)
	if err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the post field wuth the uuid: %s", f.UUID.String()), Operation: op, Err: err}
	}
	return f, nil
}

// Update a post field by ID
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *FieldsStore) Delete(postID int, f domain.PostField) error {
	const op = "FieldsRepository.Delete"
	if _, err := s.DB.Exec("DELETE FROM post_fields WHERE uuid = ? AND field_key = ? AND post_id = ?", f.UUID, f.Key, postID); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete post field with the uuid: %v", f.UUID), Operation: op, Err: err}
	}
	return nil
}

// Exists Checks if a post field exists by the given UUID and key
func (s *FieldsStore) Exists(postID int, uniq uuid.UUID, key, name string) bool {
	var normalAmount int
	err := s.DB.QueryRow("SELECT COUNT(*) FROM `post_fields` WHERE `uuid` = ? AND `post_id` = ? AND `field_key` = ? AND `name` = ?", uniq.String(), postID, key, name).Scan(&normalAmount)
	if err != nil {
		return true
	}
	return normalAmount > 0
}

// GetLayout loops over all of the locations within the config json
// file that is defined. Produces an array of field groups that
// can be returned for the post
func (s *FieldsStore) GetLayout(post domain.PostDatum) domain.FieldGroups {
	return s.finder.Layout(post, s.options.CacheServerFields)
}

// shouldDelete
// Finds fields in the domain.PostField array that should be deleted.
func (s *FieldsStore) shouldDelete(f domain.PostField, fields domain.PostFields) bool {
	for _, v := range fields {
		if (f.Key == v.Key) && (f.UUID == v.UUID) && (f.Name == v.Name) {
			return true
		}
	}
	return false
}
