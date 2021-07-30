// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/options"
	"github.com/verbiscms/verbis/api/sys"
	"github.com/verbiscms/verbis/api/version"
)

// Service defines methods for the site.
type Service interface {
	// Global returns the domain.Site struct obtained from the
	// options and retrieves the latest Verbis version.
	Global() domain.Site
}

// Site defines the data layer for set methods.
type Site struct {
	options options.Repository
	sys     sys.System
}

// Global returns the domain.Site struct obtained from the
// options and retrieves the latest Verbis version.
func (s *Site) Global() domain.Site {
	o := s.options.Struct()
	return domain.Site{
		Title:         o.SiteTitle,
		Description:   o.SiteDescription,
		Logo:          o.SiteLogo,
		Url:           o.SiteUrl,
		Version:       version.Version,
		RemoteVersion: s.sys.LatestVersion(),
		HasUpdate:     s.sys.HasUpdate(),
	}
}

// New creates a new site Service.
func New(opts options.Repository, s sys.System) Service {
	return &Site{
		options: opts,
		sys:     s,
	}
}
