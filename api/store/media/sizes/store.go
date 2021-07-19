// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sizes

import (
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store/config"
	"github.com/ainsleyclark/verbis/api/store/files"
)

// Repository defines methods for media sizes
// to interact with the database.
type Repository interface {
	Find(mediaID int) (domain.MediaSizes, error)
	Create(mediaID int, sizes domain.MediaSizes) (domain.MediaSizes, error)
	Delete(mediaID int) error
}

// Store defines the data layer for media sizes.
type Store struct {
	*config.Config
}

const (
	// TableName is the database table name for media sizes.
	TableName = "media_sizes"
)

// New
//
// Creates a new media sizes store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
	}
}

// selectStmt is a helper for SELECT Statements,
// joining files by file id.
func (s *Store) selectStmt() *builder.Sqlbuilder {
	return s.Builder().
		SelectRaw(s.Schema()+TableName+".*, "+
			s.Schema()+"file.id `file.id`, "+
			s.Schema()+"file.url `file.url`, "+
			s.Schema()+"file.name `file.name`, "+
			s.Schema()+"file.bucket_id `file.bucket_id`, "+
			s.Schema()+"file.mime `file.mime`, "+
			s.Schema()+"file.source_type `file.source_type`, "+
			s.Schema()+"file.provider `file.provider`, "+
			s.Schema()+"file.region `file.region`, "+
			s.Schema()+"file.bucket `file.bucket`, "+
			s.Schema()+"file.file_size `file.file_size`, "+
			s.Schema()+"file.private `file.private`").
		From(s.Schema()+TableName).
		LeftJoin(s.Schema()+files.TableName, "file", s.Schema()+TableName+".file_id = "+s.Schema()+"file.id")
}
