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

// TODO
//  - Consider changing the router path for screenshots, looks a bit messy
// 	- Test this package
//  - Need to return the screenshot path in the API for Vue to use.
// 	- Consider moving the find screenshot path to config or even the domain package.
//  - The screenshot should come back with every request, even on configuration.
//  - Add comments to functions
//  - Check if we dont have the pass the path here, would be better not to pass anything.

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
)

// TODO, if we already have the options, can we not just construct the theme path ourselves?

// New
//
// Creates a new SiteRepository.
func New(opts *domain.Options) *Site {
	return &Site{
		config:  config.Get(),
		options: opts,
		theme:   paths.Get().Base + string(os.PathSeparator) + "themes" + string(os.PathSeparator),
	}
}
