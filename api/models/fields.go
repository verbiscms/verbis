package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/fields"
	//"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/jmoiron/sqlx"
)

// FieldsRepository defines methods for Posts to interact with the database
type FieldsRepository interface {
	//GetByPost(postId int) ([]domain.PostField, error)
	//GetByPostAndKey(key string, postId int) (domain.PostField, error)
	//GetFieldGroups() (*[]domain.FieldGroup, error)
	GetLayout(p domain.Post, a domain.User, c *domain.Category) []domain.FieldGroup
}

// FieldsStore defines the data layer for Posts
type FieldsStore struct {
	db       *sqlx.DB
	config   config.Configuration
	options  domain.Options
	jsonPath string
}

// newFields - Construct
func newFields(db *sqlx.DB, config config.Configuration) *FieldsStore {
	const op = "FieldsRepository.newFields"

	fs := FieldsStore{
		db:       db,
		config:   config,
		options:  newOptions(db, config).GetStruct(),
		jsonPath: paths.Storage() + "/fields",
	}

	return &fs
}

func (s *FieldsStore) GetByPost(postId int) ([]domain.PostField, error) {
	const op = "FieldsRepository.GetByPost"
	var f []domain.PostField
	if err := s.db.Get(&f, "SELECT * FROM post_fields WHERE page_id = ?", postId); err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post field with the key: %d", postId), Operation: op, Err: err}
	}
	return f, nil
}

func (s *FieldsStore) GetByPostAndKey(key string, postId int) (domain.PostField, error) {
	const op = "FieldsRepository.GetByPostAndKey"
	var f domain.PostField
	if err := s.db.Get(&f, "SELECT * FROM post_fields WHERE page_id = ? AND field_key = ? LIMIT = 1", postId, key); err != nil {
		return domain.PostField{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post field with the page ID: %d and key: %s", postId, key), Operation: op, Err: err}
	}
	return f, nil
}

// GetLayout loops over all of the locations within the config json
// file that is defined. Produces an array of field groups that
// can be returned for the post
func (s *FieldsStore) GetLayout(p domain.Post, a domain.User, c *domain.Category) []domain.FieldGroup {
	return fields.NewLocation().GetLayout(p, a, c, s.options.CacheServerFields)
}
