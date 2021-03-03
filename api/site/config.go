// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/domain"
)

// Global
//
// Returns the domain.Site struct from the options and
// retrieves the latest Verbis version.
func (s *Site) Global() domain.Site {
	return domain.Site{
		Title:       s.options.SiteTitle,
		Description: s.options.SiteDescription,
		Logo:        s.options.SiteLogo,
		URL:         s.options.SiteURL,
		Version:     api.App.Version,
	}
}
