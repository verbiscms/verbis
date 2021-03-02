// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Layouts
//
// Retrieves all layouts stored within the layouts
// directory of the theme path.
//
// Returns errors.NOTFOUND if no layouts were found.
// Returns errors.INTERNAL if the layout path is invalid.
func (s *Site) Layouts(themePath string) (domain.Layouts, error) {
	const op = "SiteRepository.GetLayouts"

	layoutDir := themePath + s.config.TemplateDir
	files, err := s.walkMatch(layoutDir, "*"+s.config.FileExtension)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error getting templates with the path: %s", layoutDir), Operation: op, Err: ErrNoLayouts}
	}

	var layouts domain.Layouts
	for _, file := range files {
		layouts = append(layouts, domain.Layout{
			Key:  file,
			Name: s.fileName(file),
		})
	}

	if len(layouts) == 0 {
		return domain.Layouts{}, &errors.Error{Code: errors.NOTFOUND, Message: "No layouts available", Operation: op, Err: ErrNoLayouts}
	}

	return layouts, nil
}
