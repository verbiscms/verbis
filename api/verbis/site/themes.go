// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"io/ioutil"
	"os"
)

// Themes
//
// Returns a slice of domain.Theme by reading the theme
// If the file is a directory, it will be skipped
// until a config file has been found.
//
// Returns ErrNoThemes in any error case.
// Returns errors.NOTFOUND if no themes were found.
// Returns errors.INTERNAL if the theme path is invalid.
func (s *Site) Themes() (domain.Themes, error) {
	const op = "SiteRepository.Themes"

	files, err := ioutil.ReadDir(s.theme)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error finding themes", Err: err, Operation: op}
	}

	var themes domain.Themes
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		cfg, err := config.Find(s.theme + string(os.PathSeparator) + f.Name())
		if err != nil {
			logger.WithError(err).Error()
			continue
		}
		if cfg.Theme.Name == s.options.ActiveTheme {
			cfg.Theme.Active = true
		}
		themes = append(themes, cfg.Theme)
	}

	if len(themes) == 0 {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No themes available", Operation: op, Err: ErrNoThemes}
	}

	return themes, nil
}
