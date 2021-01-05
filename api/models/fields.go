package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	location "github.com/ainsleyclark/verbis/api/fields/converter"
	"github.com/google/uuid"
	"github.com/gookit/color"

	//"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/jmoiron/sqlx"
)

// FieldsRepository defines methods for Posts to interact with the database
type FieldsRepository interface {
	UpdateCreate(postId int, f []domain.PostField) error
	Create(f domain.PostField) (domain.PostField, error)
	Update(f domain.PostField) (domain.PostField, error)
	Exists(uuid uuid.UUID) bool
	GetByPost(postId int) ([]domain.PostField, error)
	GetLayout(p domain.Post, a domain.User, c *domain.Category) []domain.FieldGroup
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

	fmt.Println(postId)
	for _, v := range f {
		v.PostId = postId
		if s.Exists(v.UUID) {
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

// Update a post field by Id
//Returns errors.INTERNAL if the SQL query was invalid.
func (s *FieldsStore) Create(f domain.PostField) (domain.PostField, error) {
	const op = "FieldsRepository.Update"
	q := "INSERT INTO post_fields (uuid, post_id, type, name, value, parent, layout, row_index) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := s.db.Exec(q, f.UUID.String(), f.PostId, f.Type, f.Name, f.OriginalValue, f.Parent, f.Layout, f.Index)
	if err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the post field wuth the name: %s", f.Name), Operation: op, Err: err}
	}
	return f, nil
}

// Update a post field by Id
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *FieldsStore) Update(f domain.PostField) (domain.PostField, error) {
	const op = "FieldsRepository.Update"
	_, err := s.db.Exec("UPDATE post_fields SET type = ?, name = ?, value = ?, parent = ?, layout = ?, row_index = ? WHERE uuid = ?", f.Type, f.Name, f.OriginalValue, f.Parent, f.Layout, f.Index, f.UUID.String())
	if err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the post field wuth the uuid: %s", f.UUID.String()), Operation: op, Err: err}
	}
	return f, nil
}

// Exists Checks if a post field exists by the given UUID
func (s *FieldsStore) Exists(uuid uuid.UUID) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM post_fields WHERE uuid = ?)", uuid.String()).Scan(&exists)
	return exists
}

func (s *FieldsStore) GetByPost(postId int) ([]domain.PostField, error) {
	const op = "FieldsStore.GetByPost"
	var f []domain.PostField
	if err := s.db.Select(&f, "SELECT * FROM post_fields WHERE post_id = ?", postId); err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post field with the ID: %d", postId), Operation: op, Err: err}
	}
	color.Yellow.Println(f)
	return f, nil
}

func (s *FieldsStore) GetByPostAndKey(key string, postId int) (domain.PostField, error) {
	const op = "FieldsRepository.GetByPostAndKey"
	var f domain.PostField
	if err := s.db.Select(&f, "SELECT * FROM post_fields WHERE post_id = ? AND field_key = ? LIMIT = 1", postId, key); err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post field with the page ID: %d and key: %s", postId, key), Operation: op, Err: err}
	}
	return f, nil
}

// GetLayout loops over all of the locations within the config json
// file that is defined. Produces an array of field groups that
// can be returned for the post
func (s *FieldsStore) GetLayout(p domain.Post, a domain.User, c *domain.Category) []domain.FieldGroup {
	return s.finder.GetLayout(p, a, c, s.options.CacheServerFields)
}
