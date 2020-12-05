package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/jmoiron/sqlx"
)

// FormsRepository defines methods for Posts to interact with the database
type FormsRepository interface {
	GetByUUID(uuid string) (domain.Form, error)
}

// SeoMetaStore defines the data layer for Seo & Meta Options
type FormsStore struct {
	db *sqlx.DB
}

// newSeoMeta - Construct
func newForms(db *sqlx.DB) *FormsStore {
	return &FormsStore{
		db: db,
	}
}

// GetByUUID returns a form by UUID.
//
// Returns errors.NOTFOUND if the form was not found by the given UUID.
func (s *FormsStore) GetByUUID(uuid string) (domain.Form, error) {
	const op = "FormsRepository.GetByUUID"
	var f domain.Form
	if err := s.db.Get(&f, "SELECT * FROM forms WHERE uuid = ? LIMIT 1", uuid); err != nil {
		fmt.Println(err)
		return domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the form with the UUID: %s", uuid), Operation: op}
	}
	return f, nil
}



