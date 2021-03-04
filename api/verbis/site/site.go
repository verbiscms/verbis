// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"errors"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"os"
)

// SiteRepository defines methods for the site.
type Repository interface {
	Global() domain.Site
	Templates() (domain.Templates, error)
	Layouts() (domain.Layouts, error)
	Themes() (domain.Themes, error)
	Screenshot(theme string, file string) ([]byte, string, error)
}

// Site defines the data layer for Posts
type Site struct {
	config  *domain.ThemeConfig
	options *domain.Options
	theme   string
}

var (
	// ErrNoTemplates is returned by Templates when no page
	// templates have been found by the walk matcher.
	ErrNoTemplates = errors.New("no page templates found")
	// ErrNoLayouts is returned by Layouts when no page
	// layouts have been found by the walk matcher.
	ErrNoLayouts = errors.New("no page templates found")
	// ErrNoThemes is returned by Themes when no themes
	// have been found by looping over the theme's
	// directory.
	ErrNoThemes = errors.New("no page templates found")
)

// New
//
// Creates a new SiteRepository.
func New(opts *domain.Options) *Site {
	return &Site{
		config:  config.Get(),
		options: opts,
		theme:   paths.Get().Base + string(os.PathSeparator) + "themes",
	}
}