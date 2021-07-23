// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/sys"
	"github.com/verbiscms/verbis/api/version"
)

// Repository defines methods for the site.
type Repository interface {
	Global() domain.Site
}

// Site defines the data layer for Posts
type Site struct {
	options *domain.Options
	sys     sys.System
}

// Global
//
// Returns the domain.Site struct from the options and
// retrieves the latest Verbis version.
func (s *Site) Global() domain.Site {
	return domain.Site{
		Title:         s.options.SiteTitle,
		Description:   s.options.SiteDescription,
		Logo:          s.options.SiteLogo,
		Url:           s.options.SiteUrl,
		Version:       version.Version,
		RemoteVersion: s.sys.LatestVersion(),
		HasUpdate:     s.sys.HasUpdate(),
	}
}

// New
//
// Creates a new SiteRepository.
func New(opts *domain.Options, s sys.System) Repository {
	return &Site{
		options: opts,
		sys:     s,
	}
}
