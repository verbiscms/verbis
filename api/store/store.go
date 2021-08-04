// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/store/auth"
	"github.com/verbiscms/verbis/api/store/categories"
	storeConfig "github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/store/fields"
	"github.com/verbiscms/verbis/api/store/files"
	"github.com/verbiscms/verbis/api/store/forms"
	"github.com/verbiscms/verbis/api/store/media"
	"github.com/verbiscms/verbis/api/store/options"
	"github.com/verbiscms/verbis/api/store/posts"
	"github.com/verbiscms/verbis/api/store/redirects"
	"github.com/verbiscms/verbis/api/store/roles"
	"github.com/verbiscms/verbis/api/store/users"
)

// Repository defines all of the repositories used
// to interact with the database
type Repository struct {
	Auth       auth.Repository
	Categories categories.Repository
	Fields     fields.Repository
	Files      files.Repository
	Forms      forms.Repository
	Media      media.Repository
	Options    options.Repository
	Posts      posts.Repository
	Redirects  redirects.Repository
	Roles      roles.Repository
	User       users.Repository
}

// New creates a new database instance, connect
// to database.
// TODO Change!
func New(db database.Driver, running bool) (*Repository, error) {
	p := paths.Get()
	cfg := &storeConfig.Config{
		Driver:  db,
		Paths:   p,
		Owner:   nil,
		Running: running,
		Theme:   &config.Config{ThemePath: p.Themes},
	}

	user := users.New(cfg)
	owner := user.Owner()
	cfg.Owner = &owner

	return &Repository{
		Auth:       auth.New(cfg),
		Categories: categories.New(cfg),
		Fields:     fields.New(cfg),
		Files:      files.New(cfg),
		Forms:      forms.New(cfg),
		Media:      media.New(cfg),
		Options:    options.New(cfg),
		Posts:      posts.New(cfg),
		Redirects:  redirects.New(cfg),
		Roles:      roles.New(cfg),
		User:       user,
	}, nil
}
