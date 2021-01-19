package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	location "github.com/ainsleyclark/verbis/api/fields/converter"
	"github.com/google/uuid"
	//"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/jmoiron/sqlx"
)

// FieldsRepository defines methods for Posts to interact with the database
type FieldsRepository interface {
	GetByPost(postId int) ([]domain.PostField, error)
	GetLayout(p domain.Post, a domain.User, c *domain.Category) []domain.FieldGroup
	UpdateCreate(postId int, f []domain.PostField) error
	Create(f domain.PostField) (domain.PostField, error)
	Update(f domain.PostField) (domain.PostField, error)
	Exists(postId int, uuid uuid.UUID, key string) bool
}

// FieldsStore defines the data layer for Posts
type FieldsStore struct {
	db      *sqlx.DB
	config  config.Configuration
	options domain.Options
	finder  location.Finder
}

// newFields - Construct
func newFields(db *sqlx.DB, config config.Configuration) *FieldsStore {
	return &FieldsStore{
		db:      db,
		config:  config,
		options: newOptions(db, config).GetStruct(),
		finder:  location.NewLocation(),
	}
}

// UpdateCreate checks to see if the record exists before updating
// or creating the new record.
func (s *FieldsStore) UpdateCreate(postId int, f []domain.PostField) error {
	fields, err := s.GetByPost(postId)
	if err != nil {
		return err
	}

	// Find fields that should be deleted (not in the array)
	for _, v := range fields {
		if !s.shouldDelete(v, f) {
			err := s.Delete(postId, v)
			if err != nil {
				return err
			}
		}
	}

	// Update or create the existing fields passed.
	for _, v := range f {
		v.PostId = postId
		if s.Exists(postId, v.UUID, v.Key) {
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
func (s *FieldsStore) GetByPost(postId int) ([]domain.PostField, error) {
	const op = "FieldsStore.GetByPost"
	var f []domain.PostField
	if err := s.db.Select(&f, "SELECT * FROM post_fields WHERE post_id = ?", postId); err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post field with the ID: %d", postId), Operation: op, Err: err}
	}
	return f, nil
}

// GetByPost fields by a post ID and key.
// Returns errors.NOTFOUND if there were no records found.
func (s *FieldsStore) GetByPostAndKey(key string, postId int) (domain.PostField, error) {
	const op = "FieldsRepository.GetByPostAndKey"
	var f domain.PostField
	if err := s.db.Select(&f, "SELECT * FROM post_fields WHERE post_id = ? AND field_key = ? LIMIT = 1", postId, key); err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post field with the page ID: %d and key: %s", postId, key), Operation: op, Err: err}
	}
	return f, nil
}

// Update a post field by Id
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *FieldsStore) Create(f domain.PostField) (domain.PostField, error) {
	const op = "FieldsRepository.Create"
	q := "INSERT INTO post_fields (uuid, post_id, type, name, value, field_key) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := s.db.Exec(q, f.UUID.String(), f.PostId, f.Type, f.Name, f.OriginalValue, f.Key)
	if err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the post field wuth the name: %s", f.Name), Operation: op, Err: err}
	}
	return f, nil
}

// Update a post field by Id
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *FieldsStore) Update(f domain.PostField) (domain.PostField, error) {
	const op = "FieldsRepository.Update"
	_, err := s.db.Exec("UPDATE post_fields SET type = ?, name = ?, value = ?, field_key = ? WHERE uuid = ? AND post_id = ?", f.Type, f.Name, f.OriginalValue, f.Key, f.UUID.String(), f.PostId)
	if err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the post field wuth the uuid: %s", f.UUID.String()), Operation: op, Err: err}
	}
	return f, nil
}

// Update a post field by Id
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *FieldsStore) Delete(postId int, f domain.PostField) error {
	const op = "FieldsRepository.Delete"
	if _, err := s.db.Exec("DELETE FROM post_fields WHERE uuid = ? AND field_key = ? AND post_id = ?", f.UUID, f.Key, postId); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete post field with the uuid: %v", f.UUID), Operation: op, Err: err}
	}
	return nil
}

// Exists Checks if a post field exists by the given UUID and key
func (s *FieldsStore) Exists(postId int, uuid uuid.UUID, key string) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM post_fields WHERE uuid = ? AND post_id = ? AND field_key = ?)", uuid.String(), postId, key).Scan(&exists)
	return exists
}

// GetLayout loops over all of the locations within the config json
// file that is defined. Produces an array of field groups that
// can be returned for the post
func (s *FieldsStore) GetLayout(p domain.Post, a domain.User, c *domain.Category) []domain.FieldGroup {
	return s.finder.GetLayout(p, a, c, s.options.CacheServerFields)
}

// shouldDelete
// Finds fields in the domain.PostField array that should be deleted.
func (s *FieldsStore) shouldDelete(f domain.PostField, fields []domain.PostField) bool {
	for _, v := range fields {
		if (f.Key == v.Key) && (f.UUID == v.UUID) && (f.Name == v.Name) {
			return true
		}
	}
	return false
}
