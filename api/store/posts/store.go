// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/services/fields/location"
	"github.com/ainsleyclark/verbis/api/store/config"
	"github.com/ainsleyclark/verbis/api/store/fields"
	"github.com/ainsleyclark/verbis/api/store/posts/categories"
	"github.com/ainsleyclark/verbis/api/store/posts/meta"
	"github.com/ainsleyclark/verbis/api/store/users"
	"github.com/google/uuid"
)

// Repository defines methods for posts
// to interact with the database.
type Repository interface {
	List(meta params.Params, layout bool, cfg ListConfig) (domain.PostData, int, error)
	Find(id int, layout bool) (domain.PostDatum, error)
	FindBySlug(slug string) (domain.PostDatum, error)
	Create(p domain.PostCreate) (domain.PostDatum, error)
	Update(p domain.PostCreate) (domain.PostDatum, error)
	Delete(id int) error
	Exists(id int) bool
	ExistsBySlug(slug string) bool
}

// Store defines the data layer for posts.
type Store struct {
	*config.Config
	categories categories.Repository
	fields     fields.Repository
	meta       meta.Repository
	users      users.Repository
	finder     location.Finder
}

const (
	// TableName is the database table name for posts.
	TableName = "posts"
)

var (
	// ErrPostsExists is returned by validate when
	// a post already exists.
	ErrPostsExists = errors.New("post already exists")
	// ErrNoPageTemplate is returned by validate when
	// no page template has been matched with the
	// one passed.
	ErrNoPageTemplate = errors.New("no page template matched")
	// ErrNoPageLayout is returned by validate when
	// no page layout has been matched with the
	// one passed.
	ErrNoPageLayout = errors.New("no page layout matched")
)

// New
//
// Creates a new posts store.
func New(cfg *config.Config) *Store {
	return &Store{
		Config:     cfg,
		categories: categories.New(cfg),
		fields:     fields.New(cfg),
		meta:       meta.New(cfg),
		users:      users.New(cfg),
		finder:     location.NewLocation(cfg.ThemePath),
	}
}

// postsRaw
type postsRaw struct {
	domain.Post
	Author   domain.User      `db:"author"`
	Category *domain.Category `db:"category"`
	Field    postsRawFields   `db:"field"`
}

// postsRawFields
type postsRawFields struct {
	PostId        int        `db:"post_id"` //nolint
	UUID          *uuid.UUID `db:"uuid"`
	Type          string     `db:"type"`
	Name          string     `db:"name"`
	Key           string     `db:"field_key"`
	OriginalValue string     `db:"value" json:"value"`
}

// selectStmt
func selectStmt(query string) string {
	return fmt.Sprintf(`SELECT posts.*, post_options.seo 'options.seo', post_options.meta 'options.meta',
       users.id as 'author.id', users.uuid as 'author.uuid', users.first_name 'author.first_name', users.last_name 'author.last_name', users.email 'author.email', users.website 'author.website', users.facebook 'author.facebook', users.twitter 'author.twitter', users.linked_in 'author.linked_in',
       users.instagram 'author.instagram', users.biography 'author.biography', users.profile_picture_id 'author.profile_picture_id', users.updated_at 'author.updated_at', users.created_at 'author.created_at',
       roles.id 'author.roles.id', roles.name 'author.roles.name', roles.description 'author.roles.description',
       pf.uuid 'field.uuid',
       CASE WHEN categories.id IS NULL THEN 0 ELSE categories.id END AS 'category.id',
       CASE WHEN categories.name IS NULL THEN '' ELSE categories.name END AS 'category.name',
       CASE WHEN categories.resource IS NULL THEN '' ELSE categories.resource END AS 'category.resource',
       CASE WHEN pf.type IS NULL THEN "" ELSE pf.type END AS 'field.type',
       CASE WHEN pf.field_key IS NULL THEN "" ELSE pf.field_key END AS 'field.field_key',
       CASE WHEN pf.name IS NULL THEN "" ELSE pf.name END AS 'field.name',
       CASE WHEN pf.value IS NULL THEN "" ELSE pf.value END AS 'field.value'
FROM (%s) posts
      LEFT JOIN post_options ON posts.id = post_options.post_id
      LEFT JOIN users ON posts.user_id = users.id
      INNER JOIN user_roles ON users.id = user_roles.user_id
      LEFT JOIN roles ON user_roles.role_id = roles.id
      LEFT JOIN post_categories pc on posts.id = pc.post_id
      LEFT JOIN categories on pc.category_id = categories.id
      LEFT JOIN post_fields pf on posts.id = pf.post_id`, query)
}
