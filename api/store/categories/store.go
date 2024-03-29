// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/store/config"
)

// Repository defines methods for categories
// to interact with the database.
type Repository interface {
	List(meta params.Params, cfg ListConfig) (domain.Categories, int, error)
	Find(id int) (domain.Category, error)
	FindByPost(id int) (domain.Category, error)
	FindBySlug(slug string) (domain.Category, error)
	FindByName(name string) (domain.Category, error)
	FindParent(id int) (domain.Category, error)
	Create(c domain.Category) (domain.Category, error)
	Update(c domain.Category) (domain.Category, error)
	Delete(id int) error
	Exists(id int) bool
	ExistsByName(name string) bool
	ExistsBySlug(slug string) bool
}

// Store defines the data layer for categories.
type Store struct {
	*config.Config
}

const (
	// The database table name for categories.
	TableName = "categories"
	// The database table name for the categories pivot.
	PivotTableName = "post_categories"
)

var (
	// ErrCategoryExists is returned by validate when
	// a category already exists.
	ErrCategoryExists = errors.New("category already exists")
)

// New
//
// Creates a new categories store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config: cfg,
	}
}
