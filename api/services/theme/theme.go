// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"errors"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/options"
)

// Service defines methods for the Theme.
type Service interface {
	// Config obtains the current active theme configuration from
	// the themes path. The configuration will be stored in
	// cache until another them is activated.
	//
	// Logs errors.INTERNAL if the cache item could not be cast.
	// Returns an error if the theme or configuration could not be obtained.
	Config() (domain.ThemeConfig, error)
	// Activate sets a new theme based on the theme string provided.
	// The cache will be cleaned and set if the activation was successful.
	//
	// Returns errors.INVALID if the theme could not be found within the themes directory.
	// Returns an error if the theme or configuration could not be obtained.
	Activate(theme string) (domain.ThemeConfig, error)
	// Find retrieves a theme configuration by name.
	//
	// Returns an error if the theme could not be found.
	Find(theme string) (domain.ThemeConfig, error)
	// List retrieves all theme configuration files from the themes
	// directory.
	//
	// Returns errors.INTERNAL if there was an error reading the theme directory.
	// Returns errors.NOTFOUND ErrNoThemes if there are no themes available.
	// Logs error if the the configuration is not found in a given directory.
	List() ([]domain.ThemeConfig, error)
	// Exists checks to see if a theme exists within the themes
	// directory by using os.Stat
	Exists(theme string) bool
	// Templates retrieves all templates stored within the
	// templates directory of the active theme.
	//
	// Returns ErrNoTemplates in any error case.
	// Returns errors.NOTFOUND if no templates were found.
	// Returns errors.INTERNAL if the template path is invalid.
	Templates() (domain.Templates, error)
	// Layouts retrieves all layouts stored within the
	// layouts directory of the active theme.
	//
	// Returns ErrNoLayouts in any error case.
	// Returns errors.NOTFOUND if no layouts were found.
	// Returns errors.INTERNAL if the layout path is invalid.
	Layouts() (domain.Layouts, error)
	// Screenshot finds a screenshot in the Theme directory based on
	// the Theme passed (e.g. verbis) and the file passed
	// (e.g. screenshot.png).
	//
	// Returns errors.NOTFOUND if there was not screenshot found.
	Screenshot(theme string, file string) ([]byte, domain.Mime, error)
}

// Theme defines the data layer for Verbis themes.
type Theme struct {
	config     config.Provider
	cache      cache.Store
	options    options.Repository
	themesPath string
}

var (
	// ErrNoTemplates is returned by Templates when no page
	// templates have been found by the walk matcher.
	ErrNoTemplates = errors.New("no page templates found")
	// ErrNoLayouts is returned by Layouts when no page
	// layouts have been found by the walk matcher.
	ErrNoLayouts = errors.New("no page templates found")
	// ErrNoThemes is returned by List when no themes
	// have been found by looping over the theme's
	// directory.
	ErrNoThemes = errors.New("no themes found")
)

// New Creates a new Theme service.
func New(store cache.Store, options options.Repository) *Theme {
	themePath := paths.Get().Themes
	return &Theme{
		config:     &config.Config{ThemePath: themePath},
		cache:      store,
		options:    options,
		themesPath: paths.Get().Themes,
	}
}
