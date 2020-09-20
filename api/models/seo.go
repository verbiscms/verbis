package models

import (
	"cms/api/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type SeoMetaRepository interface {
	UpdateCreate(p *domain.Post) error
}

type SeoMetaStore struct {
	db *sqlx.DB
}

//Construct
func newSeoMeta(db *sqlx.DB) *SeoMetaStore {
	return &SeoMetaStore{
		db: db,
	}
}

// Update or create the record
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

// Check if a seo meta record exists by ID
func (s *SeoMetaStore) exists(id int) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT page_id FROM seo_meta_options WHERE page_id = ?)", id).Scan(&exists)
	return exists
}

// Create the seo meta options record
func (s *SeoMetaStore) create(p *domain.Post) error {
	q := "INSERT INTO seo_meta_options (page_id, seo, meta) VALUES (?, ?, ?)"
	_, err := s.db.Exec(q, p.Id, p.SeoMeta.Seo, p.SeoMeta.Seo)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not create the seo meta options record for post title: %v", p.Title)
	}
	return nil
}

// Update the seo meta options record
func (s *SeoMetaStore) update(p *domain.Post) error {
	q := "UPDATE seo_meta_options SET seo = ?, meta = ? WHERE page_id = ?"
	_, err := s.db.Exec(q, p.SeoMeta.Seo, p.SeoMeta.Seo, p.Id)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not update the seo meta options for the post title: %v", p.Title)
	}
	return nil
}





