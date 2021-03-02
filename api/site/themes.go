// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gookit/color"
	"io/ioutil"
	"os"
)

// Templates
//
// Retrieves all templates stored within the templates
// directory of the theme path.
//
// Returns errors.NOTFOUND if no templates were found.
// Returns errors.INTERNAL if the template path is invalid.
func (s *Site) Themes(themePath string) (domain.Themes, error) {
	const op = "SiteRepository.Themes"

	path := "/Users/ainsley/Desktop/Reddico/apis/verbis/themes"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error finding themes", Err: err, Operation: op}
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		cfg, err := config.Config(path + string(os.PathSeparator) + f.Name())
		if err != nil {
			logger.WithError(err).Error()
		}

		color.Green.Println(cfg.Theme.Title)
	}

	return nil, nil
}
