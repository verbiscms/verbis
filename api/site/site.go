// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"errors"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
)

// SiteRepository defines methods for the site.
type Repository interface {
	Global() domain.Site
	Templates(themePath string) (domain.Templates, error)
	Layouts(themePath string) (domain.Layouts, error)
	Themes(themePath string) (domain.Themes, error)
}

// Site defines the data layer for Posts
type Site struct {
	config  *domain.ThemeConfig
	options *domain.Options
}

var (
	// ErrNoTemplates is returned by Templates when no page
	// templates have been found by the walk matcher.
	ErrNoTemplates = errors.New("no page templates found")
	// ErrNoLayouts is returned by Layouts when no page
	// layouts have been found by the walk matcher.
	ErrNoLayouts = errors.New("no page templates found")
)

// TODO, if we already have the options, can we not just construct the theme path ourselves?

// New
//
// Creates a new SiteRepository.
func New(opts *domain.Options) *Site {
	return &Site{
		config:  config.Get(),
		options: opts,
	}
}
