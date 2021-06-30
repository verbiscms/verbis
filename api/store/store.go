// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/services/theme"
	"github.com/ainsleyclark/verbis/api/store/auth"
	"github.com/ainsleyclark/verbis/api/store/categories"
	storeConfig "github.com/ainsleyclark/verbis/api/store/config"
	"github.com/ainsleyclark/verbis/api/store/fields"
	"github.com/ainsleyclark/verbis/api/store/forms"
	"github.com/ainsleyclark/verbis/api/store/media"
	"github.com/ainsleyclark/verbis/api/store/options"
	"github.com/ainsleyclark/verbis/api/store/posts"
	"github.com/ainsleyclark/verbis/api/store/redirects"
	"github.com/ainsleyclark/verbis/api/store/roles"
	"github.com/ainsleyclark/verbis/api/store/users"
	"os"
)

// Repository defines all of the repositories used
// to interact with the database
type Repository struct {
	Auth       auth.Repository
	Categories categories.Repository
	Fields     fields.Repository
	Forms      forms.Repository
	Media      media.Repository
	Options    options.Repository
	Posts      posts.Repository
	Redirects  redirects.Repository
	Roles      roles.Repository
	User       users.Repository
}

// TODO Change!
// Create a new database instance, connect
// to database.
func New(db database.Driver, running bool) (*Repository, *domain.ThemeConfig, error) {
	cfg := &storeConfig.Config{
		Driver:       db,
		Options:      nil,
		Paths:        paths.Get(),
		Owner:        nil,
		ThemeService: theme.New(),
		Running:      false,
	}

	user := users.New(cfg)
	var themeConfig *domain.ThemeConfig

	if running {
		optsStore := options.New(cfg)

		cfg.Options = optsStore.Struct()

		activeTheme, err := optsStore.GetTheme()
		if err != nil {
			return nil, nil, err
		}

		// TODO, some sanity check here
		themePath := cfg.Paths.Themes + string(os.PathSeparator) + activeTheme

		themeConfig = config.Init(themePath)
		cfg.Theme = themeConfig
		cfg.ThemePath = themePath

		cfg.ThemeService = theme.New()

		owner := user.Owner()
		cfg.Owner = &owner
	}

	return &Repository{
		Auth:       auth.New(cfg),
		Categories: categories.New(cfg),
		Fields:     fields.New(cfg),
		Forms:      forms.New(cfg),
		Media:      media.New(cfg),
		Options:    options.New(cfg),
		Posts:      posts.New(cfg),
		Redirects:  redirects.New(cfg),
		Roles:      roles.New(cfg),
		User:       user,
	}, themeConfig, nil
}
