// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"errors"
	"github.com/ainsleyclark/verbis/api/common/paths"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"os"
)

// Repository defines methods for the theme.
type Repository interface {
	List(activeTheme string) ([]*domain.ThemeConfig, error)
	Find(theme string) (*domain.ThemeConfig, error)
	Exists(theme string) bool
	Templates(theme string) (domain.Templates, error)
	Layouts(theme string) (domain.Layouts, error)
	Screenshot(theme string, file string) ([]byte, domain.Mime, error)
}

// Site defines the data layer for Posts
type theme struct {
	config     *domain.ThemeConfig
	options    *domain.Options
	themesPath string
}

var (
	// ErrNoTemplates is returned by Templates when no page
	// templates have been found by the walk matcher.
	ErrNoTemplates = errors.New("no page templates found")
	// ErrNoLayouts is returned by Layouts when no page
	// layouts have been found by the walk matcher.
	ErrNoLayouts = errors.New("no page templates found")
	// ErrNoTheme is returned by Exists when no theme
	// has been found.
	ErrNoTheme = errors.New("no theme found")
)

// New
//
// Creates a new Repository.
func New() Repository {
	return &theme{
		config:     config.Get(),
		themesPath: paths.Get().Themes,
	}
}

// List
//
// List all theme configurations.
func (t *theme) List(activeTheme string) ([]*domain.ThemeConfig, error) {
	return config.All(t.themesPath, activeTheme)
}

// Find
//
// Find a theme configuration.
func (t *theme) Find(theme string) (*domain.ThemeConfig, error) {
	return config.Find(t.themesPath + string(os.PathSeparator) + theme)
}

// Exists
//
// Checks if a theme exists by name.
func (t *theme) Exists(theme string) bool {
	_, err := os.Stat(t.themesPath + string(os.PathSeparator) + theme)
	return !os.IsNotExist(err)
}
