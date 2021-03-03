// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/jmoiron/sqlx"
)

// Repository defines methods for categories
// to interact with the database.
type Repository interface {
	Get(meta params.Params) (domain.Categories, int, error)
	Find(id int64) (domain.Category, error)
	FindByPost(pageID int) (*domain.Category, error)
	FindBySlug(slug string) (domain.Category, error)
	FindByName(name string) (domain.Category, error)
	GetParent(id int) (domain.Category, error)
	Create(c *domain.Category) (domain.Category, error)
	Update(c *domain.Category) (domain.Category, error)
	Delete(id int) error
	Exists(id int) bool
}

// Store defines the data layer for Categories.
type Store struct {
	*database.Model
}

const (
	TableName      = "categories"
	PivotTableName = "post_categories"
)

func New(db *sqlx.DB) *Store {
	return &Store{
		Model: database.NewModel(db),
	}
}
