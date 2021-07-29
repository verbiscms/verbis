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
	// Config gets the themes configuration from the themes
	// path.
	// Logs errors.INTERNAL if the unmarshalling was
	// unsuccessful and returns the DefaultTheme
	// variable.
	Config() (domain.ThemeConfig, error)
	Set(theme string) (domain.ThemeConfig, error)
	Find(theme string) (domain.ThemeConfig, error)
	List() ([]domain.ThemeConfig, error)
	Exists(theme string) bool
	Templates() (domain.Templates, error)
	Layouts() (domain.Layouts, error)
	Screenshot(theme string, file string) ([]byte, domain.Mime, error)
}

// Theme defines the data layer for Verbis themes.
type Theme struct {
	config      config.Provider
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
	// ErrNoTheme is returned by Exists when no Theme
	// has been found.
	ErrNoTheme = errors.New("no Theme found")
	// ErrNoThemes is returned by List when no themes
	// have been found by looping over the theme's
	// directory.
	ErrNoThemes = errors.New("no themes found")
)

// New Creates a new Theme service.
func New(cache cache.Store, options options.Repository) *Theme {
	themePath := paths.Get().Themes
	return &Theme{
		config: &config.Config{ThemePath: themePath},
		cache:      cache,
		options:    options,
		themesPath: paths.Get().Themes,
	}
}
