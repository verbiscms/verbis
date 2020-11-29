package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"

	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/jmoiron/sqlx"
)

// SeoMetaRepository defines methods for Posts to interact with the database
type SeoMetaRepository interface {
	UpdateCreate(p *domain.Post) error
}

// SeoMetaStore defines the data layer for Seo & Meta Options
type SeoMetaStore struct {
	db *sqlx.DB
}

// newSeoMeta - Construct
func newSeoMeta(db *sqlx.DB) *SeoMetaStore {
	return &SeoMetaStore{
		db: db,
	}
}

// UpdateCreate checks to see if the record exists before updating
// or creating the new record.
func (s *SeoMetaStore) UpdateCreate(p *domain.Post) error {
	if s.exists(p.Id) {
		if err := s.update(p); err != nil {
			return err
		}
	} else {
		if err := s.create(p); err != nil {
			return err
		}
	}
	return nil
}

// Exists Checks if a seo meta record exists by Id
func (s *SeoMetaStore) exists(id int) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT page_id FROM post_options WHERE page_id = ?)", id).Scan(&exists)
	return exists
}

// create a new seo meta record
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *SeoMetaStore) create(p *domain.Post) error {
	const op = "SeoMetaRepository.create"
	_, err := s.db.Exec("INSERT INTO post_options (page_id, seo, meta) VALUES (?, ?, ?)", p.Id, p.SeoMeta.Seo, p.SeoMeta.Meta)
	if err != nil {
		fmt.Println(err)
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the seo meta options record for post title: %v", p.Title), Operation: op, Err: err}
	}
	return nil
}

// update a seo meta record by page Id
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *SeoMetaStore) update(p *domain.Post) error {
	const op = "SeoMetaRepository.update"
	_, err := s.db.Exec("UPDATE post_options SET seo = ?, meta = ? WHERE page_id = ?", p.SeoMeta.Seo, p.SeoMeta.Meta, p.Id)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the seo meta options for the post title: %v", p.Title), Operation: op, Err: err}
	}
	return nil
}
