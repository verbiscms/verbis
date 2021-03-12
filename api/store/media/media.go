// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/store"
	"mime/multipart"
)

// Repository defines methods for media items
// to interact with the database.
type Repository interface {
	List(meta params.Params) (domain.MediaItems, int, error) // done
	Find(id int) (domain.Media, error)                       //done
	FindByName(name string) (domain.Media, error)            //done
	FindByURL(url string) (domain.Media, error)              // private?
	Serve(uploadPath string, acceptWeb bool) ([]byte, string, error)
	Upload(file *multipart.FileHeader, token string) (domain.Media, error)
	Validate(file *multipart.FileHeader) error
	Update(m *domain.Media) error
	Delete(id int) error
	Exists(fileName string) bool
}

// Store defines the data layer for media.
type Store struct {
	*store.Config
}

const (
	// The database table name for media.
	TableName = "media"
)

// New
//
// Creates a new media store.
func New(cfg *store.Config) *Store {
	return &Store{
		Config: cfg,
	}
}

// selectStmt
//
// Helper for SELECT Statements, preventing null
// values & nil pointers for name, description
// and alt tags.
func (s *Store) selectStmt() *builder.Sqlbuilder {
	// TODO, may not support Postgres!
	return s.Builder().
		SelectRaw(`id, uuid, url, file_path, file_size, file_name, sizes, type, user_id, updated_at, created_at,
		CASE WHEN title IS NULL THEN '' ELSE title END AS 'title',
		CASE WHEN alt IS NULL THEN '' ELSE alt END AS 'alt',
		CASE WHEN description IS NULL THEN '' ELSE description END AS 'description'`).
		From(s.Schema() + TableName)
}
