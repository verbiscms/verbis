// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"os"
)

// Layouts
//
// Retrieves all layouts stored within the layouts
// directory of the theme path.
//
// Returns ErrNoLayouts in any error case.
// Returns errors.NOTFOUND if no layouts were found.
// Returns errors.INTERNAL if the layout path is invalid.
func (t *theme) Layouts(theme string) (domain.Layouts, error) {
	const op = "SiteRepository.GetLayouts"

	layoutDir := t.themesPath + string(os.PathSeparator) + theme + string(os.PathSeparator) + t.config.LayoutDir
	files, err := t.walkMatch(layoutDir, "*"+t.config.FileExtension)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error getting layouts with the path: %s", layoutDir), Operation: op, Err: ErrNoLayouts}
	}

	var layouts domain.Layouts
	for _, file := range files {
		layouts = append(layouts, domain.Layout{
			Key:  file,
			Name: t.fileName(file),
		})
	}

	if len(layouts) == 0 {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No layouts available", Operation: op, Err: ErrNoLayouts}
	}

	return layouts, nil
}
